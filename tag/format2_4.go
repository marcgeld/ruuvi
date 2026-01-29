package tag

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

// Format2Data represents decoded RuuviTag Data Format 2 (URL-based) sensor data.
// This format is obsolete and was used on Kickstarter devices.
type Format2Data struct {
	Temperature *float64 // Temperature in degrees Celsius
	Humidity    *float64 // Relative humidity in percent
	Pressure    *int     // Atmospheric pressure in Pascals
}

// Format4Data represents decoded RuuviTag Data Format 4 (URL-based with ID) sensor data.
// This format is obsolete and was the primary format in RuuviTags shipped before June 2018.
type Format4Data struct {
	Temperature *float64 // Temperature in degrees Celsius
	Humidity    *float64 // Relative humidity in percent
	Pressure    *int     // Atmospheric pressure in Pascals
	TagID       *uint8   // Random tag identifier (6 most significant bits only)
}

// DecodeFormat2 decodes RuuviTag Data Format 2 (URL) from raw bytes.
// The input must be exactly 6 bytes: 1 byte format ID + 5 bytes data.
// Returns an error if the data is invalid or not Format 2.
func DecodeFormat2(data []byte) (*Format2Data, error) {
	if len(data) != 6 {
		return nil, fmt.Errorf("format 2 requires exactly 6 bytes, got %d", len(data))
	}

	if data[0] != 0x02 {
		return nil, fmt.Errorf("not format 2 data: format byte is 0x%02X", data[0])
	}

	result := &Format2Data{}

	// Humidity: byte 1, in 0.5% increments
	humRaw := data[1]
	if humRaw == 0 {
		result.Humidity = nil
	} else {
		hum := float64(humRaw) * 0.5
		result.Humidity = &hum
	}

	// Temperature: bytes 2-3, MSB is sign bit, fraction is always 0
	tempSign := int8(data[2])
	if tempSign == 0 {
		result.Temperature = nil
	} else {
		var temp float64
		if tempSign < 0 {
			temp = -float64(tempSign & 0x7F)
		} else {
			temp = float64(tempSign)
		}
		result.Temperature = &temp
	}

	// Pressure: bytes 4-5, unsigned 16-bit, offset by -50000 Pa
	pressRaw := binary.BigEndian.Uint16(data[4:6])
	if pressRaw == 0 {
		result.Pressure = nil
	} else {
		press := int(pressRaw) + 50000
		result.Pressure = &press
	}

	return result, nil
}

// EncodeFormat2 encodes Format2Data into raw bytes.
// Returns exactly 6 bytes: 1 byte format ID + 5 bytes data.
// Invalid/nil fields are encoded as zeros.
func EncodeFormat2(data *Format2Data) ([]byte, error) {
	if data == nil {
		return nil, errors.New("data cannot be nil")
	}

	result := make([]byte, 6)
	result[0] = 0x02 // Format ID

	// Humidity
	if data.Humidity == nil || math.IsNaN(*data.Humidity) {
		result[1] = 0
	} else {
		hum := uint8(*data.Humidity / 0.5)
		result[1] = hum
	}

	// Temperature (no fraction part in format 2)
	if data.Temperature == nil || math.IsNaN(*data.Temperature) {
		result[2] = 0
		result[3] = 0
	} else {
		temp := *data.Temperature
		if temp < 0 {
			absTemp := int8(-temp)
			result[2] = byte(int8(absTemp) | int8(-128)) // Set sign bit (0x80 = -128 in signed)
		} else {
			result[2] = byte(int8(temp))
		}
		result[3] = 0 // Fraction always 0
	}

	// Pressure
	if data.Pressure == nil {
		binary.BigEndian.PutUint16(result[4:6], 0)
	} else {
		press := uint16(*data.Pressure - 50000)
		binary.BigEndian.PutUint16(result[4:6], press)
	}

	return result, nil
}

// DecodeFormat4 decodes RuuviTag Data Format 4 (URL with ID) from raw bytes.
// The input must be exactly 7 bytes: 1 byte format ID + 6 bytes data.
// Returns an error if the data is invalid or not Format 4.
func DecodeFormat4(data []byte) (*Format4Data, error) {
	if len(data) != 7 {
		return nil, fmt.Errorf("format 4 requires exactly 7 bytes, got %d", len(data))
	}

	if data[0] != 0x04 {
		return nil, fmt.Errorf("not format 4 data: format byte is 0x%02X", data[0])
	}

	result := &Format4Data{}

	// Humidity: byte 1, in 0.5% increments
	humRaw := data[1]
	if humRaw == 0 {
		result.Humidity = nil
	} else {
		hum := float64(humRaw) * 0.5
		result.Humidity = &hum
	}

	// Temperature: bytes 2-3, MSB is sign bit, fraction is always 0
	tempSign := int8(data[2])
	if tempSign == 0 {
		result.Temperature = nil
	} else {
		var temp float64
		if tempSign < 0 {
			temp = -float64(tempSign & 0x7F)
		} else {
			temp = float64(tempSign)
		}
		result.Temperature = &temp
	}

	// Pressure: bytes 4-5, unsigned 16-bit, offset by -50000 Pa
	pressRaw := binary.BigEndian.Uint16(data[4:6])
	if pressRaw == 0 {
		result.Pressure = nil
	} else {
		press := int(pressRaw) + 50000
		result.Pressure = &press
	}

	// Tag ID: byte 6, only 6 most significant bits are readable
	tagIDRaw := data[6]
	if tagIDRaw == 0 {
		result.TagID = nil
	} else {
		result.TagID = &tagIDRaw
	}

	return result, nil
}

// EncodeFormat4 encodes Format4Data into raw bytes.
// Returns exactly 7 bytes: 1 byte format ID + 6 bytes data.
// Invalid/nil fields are encoded as zeros.
func EncodeFormat4(data *Format4Data) ([]byte, error) {
	if data == nil {
		return nil, errors.New("data cannot be nil")
	}

	result := make([]byte, 7)
	result[0] = 0x04 // Format ID

	// Humidity
	if data.Humidity == nil || math.IsNaN(*data.Humidity) {
		result[1] = 0
	} else {
		hum := uint8(*data.Humidity / 0.5)
		result[1] = hum
	}

	// Temperature (no fraction part in format 4)
	if data.Temperature == nil || math.IsNaN(*data.Temperature) {
		result[2] = 0
		result[3] = 0
	} else {
		temp := *data.Temperature
		if temp < 0 {
			absTemp := int8(-temp)
			result[2] = byte(int8(absTemp) | int8(-128)) // Set sign bit (0x80 = -128 in signed)
		} else {
			result[2] = byte(int8(temp))
		}
		result[3] = 0 // Fraction always 0
	}

	// Pressure
	if data.Pressure == nil {
		binary.BigEndian.PutUint16(result[4:6], 0)
	} else {
		press := uint16(*data.Pressure - 50000)
		binary.BigEndian.PutUint16(result[4:6], press)
	}

	// Tag ID
	if data.TagID == nil {
		result[6] = 0
	} else {
		result[6] = *data.TagID
	}

	return result, nil
}
