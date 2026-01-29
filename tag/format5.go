// Package tag provides RuuviTag Bluetooth LE advertisement data format decoders.
package tag

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"

	"github.com/marcgeld/ruuvi/common"
)

// Format5Data represents decoded RuuviTag Data Format 5 (RAWv2) sensor data.
// This is the primary format in 2.x and 3.x firmware, in production since January 2019.
type Format5Data struct {
	Temperature         *float64           // Temperature in degrees Celsius
	Humidity            *float64           // Relative humidity in percent
	Pressure            *int               // Atmospheric pressure in Pascals
	AccelerationX       *float64           // Acceleration X-axis in G
	AccelerationY       *float64           // Acceleration Y-axis in G
	AccelerationZ       *float64           // Acceleration Z-axis in G
	BatteryVoltage      *int               // Battery voltage in millivolts
	TxPower             *int               // TX power in dBm
	MovementCounter     *uint8             // Movement counter (0-254)
	MeasurementSequence *uint16            // Measurement sequence number (0-65534)
	MACAddress          *common.MACAddress // 48-bit MAC address
}

// DecodeFormat5 decodes RuuviTag Data Format 5 (RAWv2) from raw bytes.
// The input must be exactly 24 bytes: 1 byte format ID + 23 bytes data.
// Returns an error if the data is invalid or not Format 5.
func DecodeFormat5(data []byte) (*Format5Data, error) {
	if len(data) != 24 {
		return nil, fmt.Errorf("format 5 requires exactly 24 bytes, got %d", len(data))
	}

	if data[0] != 0x05 {
		return nil, fmt.Errorf("not format 5 data: format byte is 0x%02X", data[0])
	}

	result := &Format5Data{}

	// Temperature: bytes 1-2, signed 16-bit, in 0.005°C increments
	tempRaw := int16(binary.BigEndian.Uint16(data[1:3]))
	if tempRaw == -32768 { // 0x8000 = invalid
		result.Temperature = nil
	} else {
		temp := float64(tempRaw) * 0.005
		result.Temperature = &temp
	}

	// Humidity: bytes 3-4, unsigned 16-bit, in 0.0025% increments
	humRaw := binary.BigEndian.Uint16(data[3:5])
	if humRaw == 0xFFFF { // invalid
		result.Humidity = nil
	} else {
		hum := float64(humRaw) * 0.0025
		result.Humidity = &hum
	}

	// Pressure: bytes 5-6, unsigned 16-bit, offset by -50000 Pa
	pressRaw := binary.BigEndian.Uint16(data[5:7])
	if pressRaw == 0xFFFF { // invalid
		result.Pressure = nil
	} else {
		press := int(pressRaw) + 50000
		result.Pressure = &press
	}

	// Acceleration X: bytes 7-8, signed 16-bit in mG
	accXRaw := int16(binary.BigEndian.Uint16(data[7:9]))
	if accXRaw == -32768 { // 0x8000 = invalid
		result.AccelerationX = nil
	} else {
		accX := float64(accXRaw) / 1000.0 // Convert mG to G
		result.AccelerationX = &accX
	}

	// Acceleration Y: bytes 9-10, signed 16-bit in mG
	accYRaw := int16(binary.BigEndian.Uint16(data[9:11]))
	if accYRaw == -32768 { // 0x8000 = invalid
		result.AccelerationY = nil
	} else {
		accY := float64(accYRaw) / 1000.0 // Convert mG to G
		result.AccelerationY = &accY
	}

	// Acceleration Z: bytes 11-12, signed 16-bit in mG
	accZRaw := int16(binary.BigEndian.Uint16(data[11:13]))
	if accZRaw == -32768 { // 0x8000 = invalid
		result.AccelerationZ = nil
	} else {
		accZ := float64(accZRaw) / 1000.0 // Convert mG to G
		result.AccelerationZ = &accZ
	}

	// Power info: bytes 13-14, 11 bits voltage + 5 bits TX power
	powerInfo := binary.BigEndian.Uint16(data[13:15])

	// Battery voltage: first 11 bits, offset by 1600 mV
	battRaw := powerInfo >> 5
	if battRaw == 0x7FF { // 2047 = invalid
		result.BatteryVoltage = nil
	} else {
		batt := int(battRaw) + 1600
		result.BatteryVoltage = &batt
	}

	// TX power: last 5 bits, offset by -40 dBm, in 2 dBm steps
	txRaw := powerInfo & 0x1F
	if txRaw == 0x1F { // 31 = invalid
		result.TxPower = nil
	} else {
		tx := int(txRaw)*2 - 40
		result.TxPower = &tx
	}

	// Movement counter: byte 15
	movementRaw := data[15]
	if movementRaw == 0xFF { // 255 = invalid
		result.MovementCounter = nil
	} else {
		result.MovementCounter = &movementRaw
	}

	// Measurement sequence: bytes 16-17, unsigned 16-bit
	seqRaw := binary.BigEndian.Uint16(data[16:18])
	if seqRaw == 0xFFFF { // 65535 = invalid
		result.MeasurementSequence = nil
	} else {
		result.MeasurementSequence = &seqRaw
	}

	// MAC address: bytes 18-23
	var mac common.MACAddress
	copy(mac[:], data[18:24])
	if mac.IsInvalid() {
		result.MACAddress = nil
	} else {
		result.MACAddress = &mac
	}

	return result, nil
}

// EncodeFormat5 encodes Format5Data into raw bytes suitable for BLE advertisement.
// Returns exactly 24 bytes: 1 byte format ID + 23 bytes data.
// Invalid/nil fields are encoded using their respective invalid values.
func EncodeFormat5(data *Format5Data) ([]byte, error) {
	if data == nil {
		return nil, errors.New("data cannot be nil")
	}

	result := make([]byte, 24)
	result[0] = 0x05 // Format ID

	// Temperature
	if data.Temperature == nil || math.IsNaN(*data.Temperature) {
		binary.BigEndian.PutUint16(result[1:3], 0x8000)
	} else {
		temp := int16(*data.Temperature / 0.005)
		binary.BigEndian.PutUint16(result[1:3], uint16(temp))
	}

	// Humidity
	if data.Humidity == nil || math.IsNaN(*data.Humidity) {
		binary.BigEndian.PutUint16(result[3:5], 0xFFFF)
	} else {
		hum := uint16(*data.Humidity / 0.0025)
		binary.BigEndian.PutUint16(result[3:5], hum)
	}

	// Pressure
	if data.Pressure == nil {
		binary.BigEndian.PutUint16(result[5:7], 0xFFFF)
	} else {
		press := uint16(*data.Pressure - 50000)
		binary.BigEndian.PutUint16(result[5:7], press)
	}

	// Acceleration X
	if data.AccelerationX == nil || math.IsNaN(*data.AccelerationX) {
		binary.BigEndian.PutUint16(result[7:9], 0x8000)
	} else {
		accX := int16(*data.AccelerationX * 1000)
		binary.BigEndian.PutUint16(result[7:9], uint16(accX))
	}

	// Acceleration Y
	if data.AccelerationY == nil || math.IsNaN(*data.AccelerationY) {
		binary.BigEndian.PutUint16(result[9:11], 0x8000)
	} else {
		accY := int16(*data.AccelerationY * 1000)
		binary.BigEndian.PutUint16(result[9:11], uint16(accY))
	}

	// Acceleration Z
	if data.AccelerationZ == nil || math.IsNaN(*data.AccelerationZ) {
		binary.BigEndian.PutUint16(result[11:13], 0x8000)
	} else {
		accZ := int16(*data.AccelerationZ * 1000)
		binary.BigEndian.PutUint16(result[11:13], uint16(accZ))
	}

	// Power info: 11 bits voltage + 5 bits TX power
	var powerInfo uint16

	if data.BatteryVoltage == nil {
		powerInfo |= 0x7FF << 5 // 2047 shifted left 5 bits
	} else {
		batt := uint16(*data.BatteryVoltage - 1600)
		powerInfo |= (batt & 0x7FF) << 5
	}

	if data.TxPower == nil {
		powerInfo |= 0x1F // 31
	} else {
		tx := uint16((*data.TxPower + 40) / 2)
		powerInfo |= tx & 0x1F
	}

	binary.BigEndian.PutUint16(result[13:15], powerInfo)

	// Movement counter
	if data.MovementCounter == nil {
		result[15] = 0xFF
	} else {
		result[15] = *data.MovementCounter
	}

	// Measurement sequence
	if data.MeasurementSequence == nil {
		binary.BigEndian.PutUint16(result[16:18], 0xFFFF)
	} else {
		binary.BigEndian.PutUint16(result[16:18], *data.MeasurementSequence)
	}

	// MAC address
	if data.MACAddress == nil {
		for i := 18; i < 24; i++ {
			result[i] = 0xFF
		}
	} else {
		copy(result[18:24], (*data.MACAddress)[:])
	}

	return result, nil
}
