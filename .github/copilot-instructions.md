# GitHub Copilot Instructions for Ruuvi Repository

## Project Overview

This repository provides idiomatic Go packages for working with RuuviTag Bluetooth LE devices and protocols. The code decodes and encodes RuuviTag sensor data formats (2-5) according to the official Ruuvi Sensor Protocol specifications.

## Core Principles

- **Strict parsing**: Always validate data format and length
- **Type safety**: Use type-safe structs for sensor readings
- **Forward compatibility**: Support multiple format versions
- **Nil handling**: Use pointers for optional fields; nil indicates invalid/unavailable data
- **Domain-oriented layout**: Organize code by domain (`ruuvi/tag`, `ruuvi/common`)

## Code Style and Conventions

### Go Standards
- Use Go 1.25+ features and idioms
- Follow standard Go formatting (use `gofmt` and `goimports`)
- Use `golangci-lint` with the configuration in `.golangci.yml`
- Keep imports organized with local packages prefixed by `github.com/marcgeld/hermod`

### Package Organization
```
ruuvi/
├── common/          # Shared types and utilities
│   └── types.go     # Common data models (MACAddress, helper functions)
└── tag/             # RuuviTag format decoders/encoders
    ├── decoder.go   # Auto-detection and unified decoding
    ├── format2_4.go # Format 2 and 4 (URL-based, obsolete)
    ├── format3.go   # Format 3 (RAWv1, deprecated)
    └── format5.go   # Format 5 (RAWv2, production)
```

### Naming Conventions
- Use descriptive package-level comments for exported types and functions
- Format-specific types: `Format5Data`, `Format3Data`, etc.
- Decoder functions: `DecodeFormat5()`, `DecodeFormat3()`, etc.
- Encoder functions: `EncodeFormat5()`, `EncodeFormat3()`, etc.
- Use `*float64` and `*int` for optional sensor values (nil = invalid/unavailable)

### Binary Data Handling
- Use `encoding/binary` for parsing multi-byte values
- Always use `binary.BigEndian` for RuuviTag data (protocol specification)
- Validate data length before parsing
- Check for invalid values (e.g., 0x8000, 0xFFFF) and set fields to nil

### Error Handling
- Return descriptive errors using `fmt.Errorf()` with context
- Validate input data length in decoder functions
- Check format byte matches expected value
- Handle edge cases (nil pointers, invalid values, NaN)

### Testing
- Write comprehensive tests based on official Ruuvi test vectors
- Test both valid and invalid input data
- Test encoding/decoding round-trips
- Use table-driven tests for multiple test cases
- Run tests with: `go test -v -race ./...`

## Development Workflow

### Building
```bash
go build -v ./...
```

### Testing
```bash
# Run all tests with race detection and coverage
go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

# Run tests for specific package
go test -v ./tag
```

### Linting
```bash
# Using golangci-lint (see .golangci.yml for configuration)
golangci-lint run --timeout=5m
```

### Examples
Build example applications:
```bash
cd examples/basic
go build -v .
```

## Supported Data Formats

When working with format decoders/encoders, be aware of these formats:

| Format | Name | Status | Package |
|--------|------|--------|---------|
| 2 | URL | Obsolete | `tag.Format2Data` |
| 3 | RAWv1 | Deprecated (widely deployed) | `tag.Format3Data` |
| 4 | URL with ID | Obsolete | `tag.Format4Data` |
| 5 | RAWv2 | **In Production** | `tag.Format5Data` |

Focus on Format 5 (RAWv2) for new features, but maintain backward compatibility with Format 3.

## Common Patterns

### Decoding Sensor Data
```go
// Auto-detect and decode any format
decoded, err := tag.Decode(rawData)
if err != nil {
    return err
}

// Or decode specific format
data, err := tag.DecodeFormat5(rawData)
if err != nil {
    return err
}

// Check for valid fields (nil = invalid/unavailable)
if data.Temperature != nil {
    fmt.Printf("Temperature: %.2f°C\n", *data.Temperature)
}
```

### Encoding Sensor Data
```go
temp := 24.3
data := &tag.Format5Data{
    Temperature: &temp,
    // Other fields...
}

raw, err := tag.EncodeFormat5(data)
if err != nil {
    return err
}
```

### Invalid Value Handling

**Format 5 (RAWv2) Invalid Values:**
- Temperature: 0x8000 (-32768)
- Humidity: 0xFFFF (65535)
- Pressure: 0xFFFF (65535)
- Acceleration: 0x8000 (-32768)
- Battery Voltage: 0x7FF (2047) in the 11-bit field
- TX Power: 0x1F (31) in the 5-bit field
- Movement Counter: 0xFF (255)
- Measurement Sequence: 0xFFFF (65535)
- MAC Address: all bytes 0xFF

**Format 3 (RAWv1):** Check format3.go for specific invalid value handling

Always set fields to nil when invalid values are encountered.

## References

- [Official Ruuvi Sensor Protocols](https://github.com/ruuvi/ruuvi-sensor-protocols)
- [Format 5 Specification](https://github.com/ruuvi/ruuvi-sensor-protocols/blob/master/dataformat_05.md)
- [Format 3 Specification](https://github.com/ruuvi/ruuvi-sensor-protocols/blob/master/dataformat_03.md)

## CI/CD

The repository uses GitHub Actions for CI:
- **Test**: Runs on Go 1.25 with race detection
- **Lint**: Uses golangci-lint with strict configuration
- **Examples**: Builds example applications
- Exclude examples from linting (see `.golangci.yml`)

When making changes:
1. Ensure code passes `go test -race ./...`
2. Run `golangci-lint run` before committing
3. Update tests for any protocol changes
4. Maintain backward compatibility with existing formats
