package irphone

import (
	"testing"
)

func TestValidate(t *testing.T) {
	validNumbers := []string{
		"09123456789",
		"9123456789",
		"+989123456789",
		"989123456789",
		"00989123456789",
	}
	invalidNumbers := []string{
		"123456789",
		"08123456789",
		"+97123456789",
		"98912345678",
		"0098912345678",
		"abcdefghijk",
	}
	for _, num := range validNumbers {
		err := Validate(num)
		if err != nil {
			t.Errorf("expected valid, got error for %s: %v", num, err)
		}
	}
	for _, num := range invalidNumbers {
		err := Validate(num)
		if err == nil {
			t.Errorf("expected error, got valid for %s", num)
		}
	}
}

func TestTo09(t *testing.T) {
	cases := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{"+989123456789", "09123456789", false},
		{"989123456789", "09123456789", false},
		{"00989123456789", "09123456789", false},
		{"09123456789", "09123456789", false},
		{"9123456789", "09123456789", false},
		{"123456789", "", true},
	}
	for _, c := range cases {
		out, err := To09(c.input)
		if (err != nil) != c.wantErr {
			t.Errorf("To09(%s) error = %v, wantErr %v", c.input, err, c.wantErr)
		}
		if out != c.expected {
			t.Errorf("To09(%s) = %s, want %s", c.input, out, c.expected)
		}
	}
}

func TestTo98(t *testing.T) {
	cases := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{"09123456789", "989123456789", false},
		{"9123456789", "989123456789", false},
		{"+989123456789", "989123456789", false},
		{"00989123456789", "989123456789", false},
		{"123456789", "", true},
	}
	for _, c := range cases {
		out, err := To98(c.input)
		if (err != nil) != c.wantErr {
			t.Errorf("To98(%s) error = %v, wantErr %v", c.input, err, c.wantErr)
		}
		if out != c.expected {
			t.Errorf("To98(%s) = %s, want %s", c.input, out, c.expected)
		}
	}
}

func TestToPlus98(t *testing.T) {
	cases := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{"09123456789", "+989123456789", false},
		{"9123456789", "+989123456789", false},
		{"+989123456789", "+989123456789", false},
		{"989123456789", "+989123456789", false},
		{"00989123456789", "+989123456789", false},
		{"123456789", "", true},
	}
	for _, c := range cases {
		out, err := ToPlus98(c.input)
		if (err != nil) != c.wantErr {
			t.Errorf("ToPlus98(%s) error = %v, wantErr %v", c.input, err, c.wantErr)
		}
		if out != c.expected {
			t.Errorf("ToPlus98(%s) = %s, want %s", c.input, out, c.expected)
		}
	}
}

func TestSMobile(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"09123456789", true},
		{"9123456789", true},
		{"+989123456789", true},
		{"08123456789", false},
	}
	for _, c := range cases {
		result, _ := sMobile(c.input)
		if result != c.expected {
			t.Errorf("sMobile(%q) == %v, expected %v", c.input, result, c.expected)
		}
	}
}

func TestMaskPhone(t *testing.T) {
	input := "09123456789"
	expected := "0912****6789"
	result, _ := MaskPhone(input)
	if result != expected {
		t.Errorf("MaskPhone(%q) == %q, expected %q", input, result, expected)
	}
}

func TestFormatWithDash(t *testing.T) {
	input := "09123456789"
	expected := "0912-345-6789"
	result, _ := FormatWithDash(input)
	if result != expected {
		t.Errorf("FormatWithDash(%q) == %q, expected %q", input, result, expected)
	}
}

func TestExtractPhoneNumbers(t *testing.T) {
	text := "Call me at 09123456789 or +989123456789."
	expected := []string{"09123456789", "+989123456789"}
	result := ExtractPhoneNumbers(text)
	if len(result) != len(expected) {
		t.Errorf("ExtractPhoneNumbers(%q) == %v, expected %v", text, result, expected)
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("ExtractPhoneNumbers(%q)[%d] == %q, expected %q", text, i, v, expected[i])
		}
	}
}
