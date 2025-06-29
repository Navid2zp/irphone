# irphone

A simple Go package for validating, normalizing, formatting, and extracting Iranian mobile phone numbers.

## Features
- **Validation**: Check if a string is a valid Iranian mobile phone number in various formats.
- **Normalization**: Convert phone numbers to standard formats:
  - `09xxxxxxxxx` (local)
  - `98xxxxxxxxxx` (international, no plus)
  - `+98xxxxxxxxxx` (international, with plus)
- **Masking**: Hide the middle digits of a phone number for privacy.
- **Formatting**: Add dashes for readability (e.g., `0912-345-6789`).
- **Extraction**: Find and extract all valid Iranian mobile numbers from a text.

## Installation

```
go get github.com/navid2zp/irphone
```

## Usage

```go
import "github.com/navid2zp/irphone"

err := irphone.Validate("09123456789")
normalized, err := irphone.To09("+989123456789")
masked, err := irphone.MaskPhone("09123456789")
withDash, err := irphone.FormatWithDash("09123456789")
phones := irphone.ExtractPhoneNumbers("Call me at 09123456789 or +989123456789.")
```

## Supported Formats
- `09123456789`
- `9123456789`
- `+989123456789`
- `989123456789`
- `00989123456789`

## License
MIT

