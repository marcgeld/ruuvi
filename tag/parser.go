package tag

import (
	"encoding/binary"
	"errors"
	"math"
)

// Measurement represents a RuuviTag sensor measurement with optional fields.
// Fields are pointers to support "not available" values as defined by the protocol.
type Measurement struct {
	Temperature         *float64 // Temperature in degrees Celsius
	Humidity            *float64 // Humidity in percentage (0-100%)
	Pressure            *uint32  // Atmospheric pressure in Pa
	AccelerationX       *float64 // Acceleration X in G
	AccelerationY       *float64 // Acceleration Y in G
	AccelerationZ       *float64 // Acceleration Z in G
	BatteryVoltage      *uint16  // Battery voltage in mV
	TxPower             *int8    // TX power in dBm
	MovementCounter     *uint8   // Movement counter (0-254)
	MeasurementSequence *uint16  // Measurement sequence number (0-65534)
	MACAddress          *[6]byte // MAC address (6 bytes)
}

// ParseManufacturerData parses RuuviTag manufacturer data.
// It validates the input length, detects the format byte, and dispatches
// to the appropriate format-specific decoder.
func ParseManufacturerData(data []byte) (*Measurement, error) {
	if len(data) == 0 {
		return nil, errors.New("empty data")
	}

	format := data[0]
	switch format {
	case 5:
		return parseFormat5(data)
	default:
		return nil, errors.New("unsupported data format")
	}
}

// parseFormat5 parses RuuviTag data format 5 (RAWv2).
// Format 5 is 24 bytes total with the following structure:
// - Byte 0: Format (5)
// - Bytes 1-2: Temperature in 0.005 degrees
// - Bytes 3-4: Humidity in 0.0025%
// - Bytes 5-6: Pressure in 1 Pa with -50000 Pa offset
// - Bytes 7-8: Acceleration X in mG
// - Bytes 9-10: Acceleration Y in mG
// - Bytes 11-12: Acceleration Z in mG
// - Bytes 13-14: Battery voltage (11 bits) and TX power (5 bits)
// - Byte 15: Movement counter
// - Bytes 16-17: Measurement sequence number
// - Bytes 18-23: MAC address
func parseFormat5(data []byte) (*Measurement, error) {
	if len(data) != 24 {
		return nil, errors.New("invalid data length for format 5")
	}

	m := &Measurement{}

	// Temperature: bytes 1-2, signed int16, 0.005 degree resolution
	tempRaw := int16(binary.BigEndian.Uint16(data[1:3]))
	if tempRaw != math.MinInt16 { // 0x8000 is invalid
		temp := float64(tempRaw) * 0.005
		m.Temperature = &temp
	}

	// Humidity: bytes 3-4, unsigned int16, 0.0025% resolution
	humidityRaw := binary.BigEndian.Uint16(data[3:5])
	if humidityRaw != 0xFFFF { // 65535 is invalid
		humidity := float64(humidityRaw) * 0.0025
		m.Humidity = &humidity
	}

	// Pressure: bytes 5-6, unsigned int16, 1 Pa resolution, offset -50000 Pa
	pressureRaw := binary.BigEndian.Uint16(data[5:7])
	if pressureRaw != 0xFFFF { // 65535 is invalid
		pressure := uint32(pressureRaw) + 50000
		m.Pressure = &pressure
	}

	// Acceleration X: bytes 7-8, signed int16, mG units
	accelXRaw := int16(binary.BigEndian.Uint16(data[7:9]))
	if accelXRaw != math.MinInt16 { // 0x8000 is invalid
		accelX := float64(accelXRaw) / 1000.0
		m.AccelerationX = &accelX
	}

	// Acceleration Y: bytes 9-10, signed int16, mG units
	accelYRaw := int16(binary.BigEndian.Uint16(data[9:11]))
	if accelYRaw != math.MinInt16 { // 0x8000 is invalid
		accelY := float64(accelYRaw) / 1000.0
		m.AccelerationY = &accelY
	}

	// Acceleration Z: bytes 11-12, signed int16, mG units
	accelZRaw := int16(binary.BigEndian.Uint16(data[11:13]))
	if accelZRaw != math.MinInt16 { // 0x8000 is invalid
		accelZ := float64(accelZRaw) / 1000.0
		m.AccelerationZ = &accelZ
	}

	// Power info: bytes 13-14
	// First 11 bits: battery voltage above 1600mV in 1mV increments
	// Last 5 bits: TX power above -40dBm in 2dBm increments
	powerInfo := binary.BigEndian.Uint16(data[13:15])

	// Battery voltage: bits 0-10 (11 bits)
	batteryRaw := (powerInfo >> 5) & 0x7FF
	if batteryRaw != 0x7FF { // 2047 is invalid
		battery := uint16(batteryRaw + 1600)
		m.BatteryVoltage = &battery
	}

	// TX power: bits 11-15 (5 bits)
	txPowerRaw := powerInfo & 0x1F
	if txPowerRaw != 0x1F { // 31 is invalid
		txPower := int8(txPowerRaw*2 - 40)
		m.TxPower = &txPower
	}

	// Movement counter: byte 15
	movementCounter := data[15]
	if movementCounter != 0xFF { // 255 is invalid
		m.MovementCounter = &movementCounter
	}

	// Measurement sequence: bytes 16-17
	measurementSeq := binary.BigEndian.Uint16(data[16:18])
	if measurementSeq != 0xFFFF { // 65535 is invalid
		m.MeasurementSequence = &measurementSeq
	}

	// MAC address: bytes 18-23
	var mac [6]byte
	copy(mac[:], data[18:24])
	// Check if MAC is all 0xFF (invalid)
	allFF := true
	for _, b := range mac {
		if b != 0xFF {
			allFF = false
			break
		}
	}
	if !allFF {
		m.MACAddress = &mac
	}

	return m, nil
}
