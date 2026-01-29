// Package tag provides decoders and encoders for RuuviTag Bluetooth LE advertisement data formats.
//
// This package supports decoding RuuviTag sensor data from broadcast advertisements
// according to the official Ruuvi Sensor Protocol specifications. It handles formats 2–5,
// with Format 5 (RAWv2) being the current production standard.
//
// # Basic Usage
//
// To decode data when you know the format:
//
//	raw := []byte{0x05, /* ... */}
//	data, err := tag.DecodeFormat5(raw)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if data.Temperature != nil {
//	    fmt.Printf("Temperature: %.2f°C\n", *data.Temperature)
//	}
//
// To auto-detect and decode any supported format:
//
//	decoded, err := tag.Decode(raw)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	switch decoded.Format {
//	case tag.Format5:
//	    fmt.Printf("Temp: %.2f°C\n", *decoded.Format5.Temperature)
//	case tag.Format3:
//	    fmt.Printf("Temp: %.2f°C\n", *decoded.Format3.Temperature)
//	}
//
// # Data Validation
//
// All sensor fields are pointer types. A nil value indicates that the sensor
// reading is invalid or unavailable. This follows the Ruuvi specification's
// approach to marking unavailable data with specific sentinel values.
//
// # Format Support
//
// - Format 2: URL-based (obsolete, Kickstarter devices)
// - Format 3: RAWv1 (deprecated but widely deployed)
// - Format 4: URL with ID (obsolete, pre-June 2018)
// - Format 5: RAWv2 (current production standard)
//
// Format 5 is recommended for new applications as it provides the most comprehensive
// sensor data including MAC address, movement counter, and measurement sequence for
// deduplication.
//
// # Encoding Support (Experimental)
//
// Encoding support is currently EXPERIMENTAL and limited to Data Format 5 (RAWv2) only.
// The API may change in future versions. Encoding is not supported for Formats 2, 3, or 4.
//
// The encoding functions follow the official Ruuvi Data Format 5 specification for
// scaling, rounding, and sentinel values. Fields that are nil in the input data are
// encoded using the appropriate "not available" sentinel values as defined in the spec.
//
// See EncodeFormat5 and EncodeFormat5ManufacturerData for details.
//
// # References
//
// Official specifications: https://github.com/ruuvi/ruuvi-sensor-protocols
// Format 5 specification: https://github.com/ruuvi/ruuvi-sensor-protocols/blob/master/dataformat_05.md
package tag
