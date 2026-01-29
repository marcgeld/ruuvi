# ruuvi

Idiomatic Go packages for working with Ruuvi devices and protocols.

## Overview

This repository provides Go packages for decoding RuuviTag Bluetooth LE advertisement (broadcast) data formats 2–5 according to the official [Ruuvi Sensor Protocol specifications](https://github.com/ruuvi/ruuvi-sensor-protocols).

The code emphasizes:
- **Strict parsing** - Validates data format and length
- **Clear data models** - Type-safe structs for sensor readings
- **Forward compatibility** - Supports multiple format versions
- **Domain-oriented layout** - Packages organized by domain (`ruuvi/tag`, `ruuvi/common`)

## Installation

```bash
go get github.com/marcgeld/ruuvi
```

## Supported Formats

| Format | Name | Status | Decoding | Encoding |
|--------|------|--------|----------|----------|
| 2 | URL | Obsolete | ✓ | ✗ |
| 3 | RAWv1 | Deprecated | ✓ | ✗ |
| 4 | URL with ID | Obsolete | ✓ | ✗ |
| 5 | RAWv2 | **In Production** | ✓ | ✓ (Experimental) |

**Note**: Encoding support is experimental and currently limited to Format 5 only.

## Quick Start

```go
package main

import (
    "encoding/hex"
    "fmt"
    "github.com/marcgeld/ruuvi/tag"
)

func main() {
    // Example: Decode Format 5 (RAWv2) data
    raw, _ := hex.DecodeString("0512FC5394C37C0004FFFC040CAC364200CDCBB8334C884F")
    
    data, err := tag.DecodeFormat5(raw)
    if err != nil {
        panic(err)
    }
    
    // Access sensor data
    if data.Temperature != nil {
        fmt.Printf("Temperature: %.2f°C\n", *data.Temperature)
    }
    if data.Humidity != nil {
        fmt.Printf("Humidity: %.2f%%\n", *data.Humidity)
    }
    if data.Pressure != nil {
        fmt.Printf("Pressure: %d Pa\n", *data.Pressure)
    }
    if data.MACAddress != nil {
        fmt.Printf("MAC: %s\n", data.MACAddress)
    }
}
```

## Usage

### Auto-detect Format

The `Decode` function automatically detects and decodes any supported format:

```go
import "github.com/marcgeld/ruuvi/tag"

raw := []byte{0x05, /* ... */}
decoded, err := tag.Decode(raw)
if err != nil {
    // Handle error
}

switch decoded.Format {
case tag.Format5:
    // Access Format 5 data
    fmt.Printf("Temp: %.2f°C\n", *decoded.Format5.Temperature)
case tag.Format3:
    // Access Format 3 data
    fmt.Printf("Temp: %.2f°C\n", *decoded.Format3.Temperature)
}
```

### Decode Specific Formats

#### Format 5 (RAWv2) - Recommended

Format 5 is the current production format with the most comprehensive sensor data:

```go
import "github.com/marcgeld/ruuvi/tag"

raw := []byte{0x05, /* 23 more bytes */}
data, err := tag.DecodeFormat5(raw)
if err != nil {
    // Handle error
}

// All fields are pointers - nil indicates unavailable/invalid
if data.Temperature != nil {
    fmt.Printf("Temperature: %.3f°C\n", *data.Temperature)
}
if data.AccelerationX != nil {
    fmt.Printf("Acceleration X: %.3f G\n", *data.AccelerationX)
}
if data.BatteryVoltage != nil {
    fmt.Printf("Battery: %d mV\n", *data.BatteryVoltage)
}
if data.MovementCounter != nil {
    fmt.Printf("Movement events: %d\n", *data.MovementCounter)
}
```

#### Format 3 (RAWv1) - Deprecated but Widely Used

```go
import "github.com/marcgeld/ruuvi/tag"

raw := []byte{0x03, /* 13 more bytes */}
data, err := tag.DecodeFormat3(raw)
if err != nil {
    // Handle error
}

// Format 3 has fewer fields than Format 5
if data.Temperature != nil {
    fmt.Printf("Temperature: %.2f°C\n", *data.Temperature)
}
```

## Encoding Data

> **⚠️ EXPERIMENTAL**: Encoding support is currently experimental and limited to Data Format 5 (RAWv2) only. The API may change in future versions. Encoding is not supported for Formats 2, 3, or 4.

### Format 5 Encoding

You can encode sensor data into Format 5 (RAWv2) raw bytes:

```go
import "github.com/marcgeld/ruuvi/tag"

temp := 24.3
hum := 53.49
pressure := 100044

data := &tag.Format5Data{
    Temperature: &temp,
    Humidity: &hum,
    Pressure: &pressure,
    // Other fields...
}

// Encode to payload (24 bytes: format ID + data)
raw, err := tag.EncodeFormat5(data)
if err != nil {
    // Handle error
}
// raw is now a 24-byte slice starting with 0x05

// Or encode with manufacturer data prefix (26 bytes: 0x9904 + payload)
manufacturerData, err := tag.EncodeFormat5ManufacturerData(data)
if err != nil {
    // Handle error
}
// manufacturerData can be used directly in BLE manufacturer data advertisement
```

**Important Notes:**

- **Quantization**: Due to fixed-point encoding, values are quantized to the resolution defined in the Format 5 specification. Decoding encoded data may not produce exactly the same values as the input (within resolution limits).
- **Resolution limits**:
  - Temperature: 0.005°C
  - Humidity: 0.0025%
  - Pressure: 1 Pa
  - Acceleration: 0.001 G (1 mG)
  - Battery Voltage: 1 mV
  - TX Power: 2 dBm
- **Missing fields**: Fields that are `nil` are encoded using "not available" sentinel values as defined in the Ruuvi specification.
- **Manufacturer data**: The manufacturer ID `0x9904` (Ruuvi Innovations Ltd.) is prepended in little-endian format as per Bluetooth specification.

## Data Formats

### Format 5 (RAWv2) Fields

- **Temperature**: -163.835°C to +163.835°C (0.005°C resolution)
- **Humidity**: 0% to 163.835% (0.0025% resolution)
- **Pressure**: 50000 Pa to 115534 Pa (1 Pa resolution)
- **Acceleration X/Y/Z**: -32.767 G to +32.767 G (0.001 G resolution)
- **Battery Voltage**: 1600 mV to 3646 mV (1 mV resolution)
- **TX Power**: -40 dBm to +20 dBm (2 dBm steps)
- **Movement Counter**: 0 to 254 movement events
- **Measurement Sequence**: 0 to 65534 (for deduplication)
- **MAC Address**: 48-bit device address

### Format 3 (RAWv1) Fields

- **Temperature**: -127.99°C to +127.99°C (0.01°C resolution)
- **Humidity**: 0% to 100% (0.5% resolution)
- **Pressure**: 50000 Pa to 115535 Pa (1 Pa resolution)
- **Acceleration X/Y/Z**: -32.767 G to +32.767 G (0.001 G resolution)
- **Battery Voltage**: 0 mV to 65535 mV (1 mV resolution)

## Package Structure

```
ruuvi/
├── common/          # Shared types and utilities
│   └── types.go     # Common data models (Temperature, Pressure, MAC, etc.)
└── tag/             # RuuviTag format decoders/encoders
    ├── decoder.go   # Auto-detection and unified decoding
    ├── format2_4.go # Format 2 and 4 (URL-based, obsolete)
    ├── format3.go   # Format 3 (RAWv1, deprecated)
    └── format5.go   # Format 5 (RAWv2, production)
```

## Testing

The package includes comprehensive tests based on official Ruuvi test vectors:

```bash
go test ./...
```

## References

- [Official Ruuvi Sensor Protocols](https://github.com/ruuvi/ruuvi-sensor-protocols)
- [Ruuvi Documentation](https://docs.ruuvi.com)
- [Format 5 Specification](https://github.com/ruuvi/ruuvi-sensor-protocols/blob/master/dataformat_05.md)
- [Format 3 Specification](https://github.com/ruuvi/ruuvi-sensor-protocols/blob/master/dataformat_03.md)

## License

MIT License - see [LICENSE](LICENSE) file for details.

