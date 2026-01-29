package tag

import (
	"fmt"
)

// DataFormat represents the format version of RuuviTag data.
type DataFormat uint8

const (
	// Format2 is the URL-based format used on Kickstarter devices (obsolete).
	Format2 DataFormat = 2

	// Format3 is RAWv1, used in 1.x and 2.x firmware (deprecated but widely deployed).
	Format3 DataFormat = 3

	// Format4 is URL-based with ID, used before June 2018 (obsolete).
	Format4 DataFormat = 4

	// Format5 is RAWv2, the primary format in 2.x and 3.x firmware (in production).
	Format5 DataFormat = 5
)

// DetectFormat returns the format version from raw data.
// Returns an error if the data is too short or contains an unknown format.
func DetectFormat(data []byte) (DataFormat, error) {
	if len(data) == 0 {
		return 0, fmt.Errorf("data is empty")
	}

	format := DataFormat(data[0])

	switch format {
	case Format2, Format3, Format4, Format5:
		return format, nil
	default:
		return 0, fmt.Errorf("unknown format: 0x%02X", data[0])
	}
}

// DecodedData represents decoded RuuviTag data from any supported format.
// Only one of the format-specific fields will be populated based on the detected format.
type DecodedData struct {
	Format  DataFormat
	Format2 *Format2Data
	Format3 *Format3Data
	Format4 *Format4Data
	Format5 *Format5Data
}

// Decode automatically detects and decodes RuuviTag data from raw bytes.
// Returns a DecodedData structure with the appropriate format field populated.
func Decode(data []byte) (*DecodedData, error) {
	format, err := DetectFormat(data)
	if err != nil {
		return nil, err
	}

	result := &DecodedData{
		Format: format,
	}

	switch format {
	case Format2:
		decoded, err := DecodeFormat2(data)
		if err != nil {
			return nil, err
		}
		result.Format2 = decoded

	case Format3:
		decoded, err := DecodeFormat3(data)
		if err != nil {
			return nil, err
		}
		result.Format3 = decoded

	case Format4:
		decoded, err := DecodeFormat4(data)
		if err != nil {
			return nil, err
		}
		result.Format4 = decoded

	case Format5:
		decoded, err := DecodeFormat5(data)
		if err != nil {
			return nil, err
		}
		result.Format5 = decoded

	default:
		return nil, fmt.Errorf("unsupported format: %d", format)
	}

	return result, nil
}
