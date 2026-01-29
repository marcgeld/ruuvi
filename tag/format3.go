package tag

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

// Format3Data represents decoded RuuviTag Data Format 3 (RAWv1) sensor data.
// This format was the primary format in 1.x and 2.x firmware.
// It is deprecated but still in use on many deployed RuuviTags.
type Format3Data struct {
	Temperature    *float64 // Temperature in degrees Celsius
	Humidity       *float64 // Relative humidity in percent
	Pressure       *int     // Atmospheric pressure in Pascals
	AccelerationX  *float64 // Acceleration X-axis in G
	AccelerationY  *float64 // Acceleration Y-axis in G
	AccelerationZ  *float64 // Acceleration Z-axis in G
	BatteryVoltage *int     // Battery voltage in millivolts
}

// DecodeFormat3 decodes RuuviTag Data Format 3 (RAWv1) from raw bytes.
// The input must be exactly 14 bytes: 1 byte format ID + 13 bytes data.
// Returns an error if the data is invalid or not Format 3.
func DecodeFormat3(data []byte) (*Format3Data, error) {
	if len(data) != 14 {
		return nil, fmt.Errorf("format 3 requires exactly 14 bytes, got %d", len(data))
	}

	if data[0] != 0x03 {
		return nil, fmt.Errorf("not format 3 data: format byte is 0x%02X", data[0])
	}

	result := &Format3Data{}

	// Humidity: byte 1, in 0.5% increments
	humRaw := data[1]
	if humRaw == 0 {
		result.Humidity = nil
	} else {
		hum := float64(humRaw) * 0.5
		result.Humidity = &hum
	}

	// Temperature: bytes 2-3, MSB is sign bit, in 0.01°C increments
	tempSign := int8(data[2])
	tempFraction := data[3]
	if tempSign == 0 && tempFraction == 0 {
		result.Temperature = nil
	} else {
		var temp float64
		if tempSign < 0 {
			// Negative: extract absolute value (clear sign bit)
			temp = -float64(tempSign&0x7F) - float64(tempFraction)*0.01
		} else {
			temp = float64(tempSign) + float64(tempFraction)*0.01
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

	// Acceleration X: bytes 6-7, signed 16-bit in mG
	// Format 3 spec: "There is no specific value for invalid/not available sensor readings"
	// All values should be treated as valid (0 is a legitimate reading)
	accXRaw := int16(binary.BigEndian.Uint16(data[6:8]))
	accX := float64(accXRaw) / 1000.0 // Convert mG to G
	result.AccelerationX = &accX

	// Acceleration Y: bytes 8-9, signed 16-bit in mG
	accYRaw := int16(binary.BigEndian.Uint16(data[8:10]))
	accY := float64(accYRaw) / 1000.0 // Convert mG to G
	result.AccelerationY = &accY

	// Acceleration Z: bytes 10-11, signed 16-bit in mG
	accZRaw := int16(binary.BigEndian.Uint16(data[10:12]))
	accZ := float64(accZRaw) / 1000.0 // Convert mG to G
	result.AccelerationZ = &accZ

	// Battery voltage: bytes 12-13, unsigned 16-bit in mV
	battRaw := binary.BigEndian.Uint16(data[12:14])
	if battRaw == 0 {
		result.BatteryVoltage = nil
	} else {
		batt := int(battRaw)
		result.BatteryVoltage = &batt
	}

	return result, nil
}

// EncodeFormat3 encodes Format3Data into raw bytes suitable for BLE advertisement.
// Returns exactly 14 bytes: 1 byte format ID + 13 bytes data.
// Invalid/nil fields are encoded as zeros per the spec.
func EncodeFormat3(data *Format3Data) ([]byte, error) {
	if data == nil {
		return nil, errors.New("data cannot be nil")
	}

	result := make([]byte, 14)
	result[0] = 0x03 // Format ID

	// Humidity
	if data.Humidity == nil || math.IsNaN(*data.Humidity) {
		result[1] = 0
	} else {
		hum := uint8(*data.Humidity / 0.5)
		result[1] = hum
	}

	// Temperature
	if data.Temperature == nil || math.IsNaN(*data.Temperature) {
		result[2] = 0
		result[3] = 0
	} else {
		temp := *data.Temperature
		if temp < 0 {
			// Negative temperature
			absTemp := -temp
			wholePart := int(absTemp)
			fracPart := uint8(math.Round((absTemp - float64(wholePart)) * 100))
			result[2] = byte(int8(wholePart) | int8(-128)) // Set sign bit (0x80 = -128 in signed)
			result[3] = fracPart
		} else {
			// Positive temperature
			wholePart := int(temp)
			fracPart := uint8(math.Round((temp - float64(wholePart)) * 100))
			result[2] = byte(wholePart)
			result[3] = fracPart
		}
	}

	// Pressure
	if data.Pressure == nil {
		binary.BigEndian.PutUint16(result[4:6], 0)
	} else {
		press := uint16(*data.Pressure - 50000)
		binary.BigEndian.PutUint16(result[4:6], press)
	}

	// Acceleration X
	if data.AccelerationX == nil || math.IsNaN(*data.AccelerationX) {
		binary.BigEndian.PutUint16(result[6:8], 0) // Use 0 for unavailable per spec
	} else {
		accX := int16(*data.AccelerationX * 1000)
		binary.BigEndian.PutUint16(result[6:8], uint16(accX))
	}

	// Acceleration Y
	if data.AccelerationY == nil || math.IsNaN(*data.AccelerationY) {
		binary.BigEndian.PutUint16(result[8:10], 0) // Use 0 for unavailable per spec
	} else {
		accY := int16(*data.AccelerationY * 1000)
		binary.BigEndian.PutUint16(result[8:10], uint16(accY))
	}

	// Acceleration Z
	if data.AccelerationZ == nil || math.IsNaN(*data.AccelerationZ) {
		binary.BigEndian.PutUint16(result[10:12], 0) // Use 0 for unavailable per spec
	} else {
		accZ := int16(*data.AccelerationZ * 1000)
		binary.BigEndian.PutUint16(result[10:12], uint16(accZ))
	}

	// Battery voltage
	if data.BatteryVoltage == nil {
		binary.BigEndian.PutUint16(result[12:14], 0)
	} else {
		batt := uint16(*data.BatteryVoltage)
		binary.BigEndian.PutUint16(result[12:14], batt)
	}

	return result, nil
}
