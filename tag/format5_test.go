package tag

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/marcgeld/ruuvi/common"
)

// TestDecodeFormat5_ValidData tests decoding with valid data from official spec.
func TestDecodeFormat5_ValidData(t *testing.T) {
	// Test vector from official specification
	// Raw: 0x0512FC5394C37C0004FFFC040CAC364200CDCBB8334C884F
	raw, err := hex.DecodeString("0512FC5394C37C0004FFFC040CAC364200CDCBB8334C884F")
	if err != nil {
		t.Fatalf("Failed to decode hex: %v", err)
	}

	data, err := DecodeFormat5(raw)
	if err != nil {
		t.Fatalf("DecodeFormat5 failed: %v", err)
	}

	// Temperature: 24.3°C
	if data.Temperature == nil {
		t.Error("Temperature should not be nil")
	} else if !floatEquals(*data.Temperature, 24.3, 0.001) {
		t.Errorf("Temperature = %v, want 24.3", *data.Temperature)
	}

	// Humidity: 53.49%
	if data.Humidity == nil {
		t.Error("Humidity should not be nil")
	} else if !floatEquals(*data.Humidity, 53.49, 0.01) {
		t.Errorf("Humidity = %v, want 53.49", *data.Humidity)
	}

	// Pressure: 100044 Pa
	if data.Pressure == nil {
		t.Error("Pressure should not be nil")
	} else if *data.Pressure != 100044 {
		t.Errorf("Pressure = %v, want 100044", *data.Pressure)
	}

	// Acceleration X: 0.004 G
	if data.AccelerationX == nil {
		t.Error("AccelerationX should not be nil")
	} else if !floatEquals(*data.AccelerationX, 0.004, 0.001) {
		t.Errorf("AccelerationX = %v, want 0.004", *data.AccelerationX)
	}

	// Acceleration Y: -0.004 G
	if data.AccelerationY == nil {
		t.Error("AccelerationY should not be nil")
	} else if !floatEquals(*data.AccelerationY, -0.004, 0.001) {
		t.Errorf("AccelerationY = %v, want -0.004", *data.AccelerationY)
	}

	// Acceleration Z: 1.036 G
	if data.AccelerationZ == nil {
		t.Error("AccelerationZ should not be nil")
	} else if !floatEquals(*data.AccelerationZ, 1.036, 0.001) {
		t.Errorf("AccelerationZ = %v, want 1.036", *data.AccelerationZ)
	}

	// Battery Voltage: 2977 mV
	if data.BatteryVoltage == nil {
		t.Error("BatteryVoltage should not be nil")
	} else if *data.BatteryVoltage != 2977 {
		t.Errorf("BatteryVoltage = %v, want 2977", *data.BatteryVoltage)
	}

	// TX Power: 4 dBm
	if data.TxPower == nil {
		t.Error("TxPower should not be nil")
	} else if *data.TxPower != 4 {
		t.Errorf("TxPower = %v, want 4", *data.TxPower)
	}

	// Movement Counter: 66
	if data.MovementCounter == nil {
		t.Error("MovementCounter should not be nil")
	} else if *data.MovementCounter != 66 {
		t.Errorf("MovementCounter = %v, want 66", *data.MovementCounter)
	}

	// Measurement Sequence: 205
	if data.MeasurementSequence == nil {
		t.Error("MeasurementSequence should not be nil")
	} else if *data.MeasurementSequence != 205 {
		t.Errorf("MeasurementSequence = %v, want 205", *data.MeasurementSequence)
	}

	// MAC Address: CB:B8:33:4C:88:4F
	if data.MACAddress == nil {
		t.Error("MACAddress should not be nil")
	} else {
		expected := common.MACAddress{0xCB, 0xB8, 0x33, 0x4C, 0x88, 0x4F}
		if *data.MACAddress != expected {
			t.Errorf("MACAddress = %v, want %v", data.MACAddress, expected)
		}
	}
}

// TestDecodeFormat5_MaximumValues tests decoding with maximum values from official spec.
func TestDecodeFormat5_MaximumValues(t *testing.T) {
	// Test vector: maximum values
	// Raw: 0x057FFFFFFEFFFE7FFF7FFF7FFFFFDEFEFFFECBB8334C884F
	raw, err := hex.DecodeString("057FFFFFFEFFFE7FFF7FFF7FFFFFDEFEFFFECBB8334C884F")
	if err != nil {
		t.Fatalf("Failed to decode hex: %v", err)
	}

	data, err := DecodeFormat5(raw)
	if err != nil {
		t.Fatalf("DecodeFormat5 failed: %v", err)
	}

	// Temperature: 163.835°C
	if data.Temperature == nil {
		t.Error("Temperature should not be nil")
	} else if !floatEquals(*data.Temperature, 163.835, 0.001) {
		t.Errorf("Temperature = %v, want 163.835", *data.Temperature)
	}

	// Humidity: 163.835%
	if data.Humidity == nil {
		t.Error("Humidity should not be nil")
	} else if !floatEquals(*data.Humidity, 163.835, 0.001) {
		t.Errorf("Humidity = %v, want 163.835", *data.Humidity)
	}

	// Pressure: 115534 Pa
	if data.Pressure == nil {
		t.Error("Pressure should not be nil")
	} else if *data.Pressure != 115534 {
		t.Errorf("Pressure = %v, want 115534", *data.Pressure)
	}

	// Acceleration X: 32.767 G
	if data.AccelerationX == nil {
		t.Error("AccelerationX should not be nil")
	} else if !floatEquals(*data.AccelerationX, 32.767, 0.001) {
		t.Errorf("AccelerationX = %v, want 32.767", *data.AccelerationX)
	}

	// Acceleration Y: 32.767 G
	if data.AccelerationY == nil {
		t.Error("AccelerationY should not be nil")
	} else if !floatEquals(*data.AccelerationY, 32.767, 0.001) {
		t.Errorf("AccelerationY = %v, want 32.767", *data.AccelerationY)
	}

	// Acceleration Z: 32.767 G
	if data.AccelerationZ == nil {
		t.Error("AccelerationZ should not be nil")
	} else if !floatEquals(*data.AccelerationZ, 32.767, 0.001) {
		t.Errorf("AccelerationZ = %v, want 32.767", *data.AccelerationZ)
	}

	// Battery Voltage: 3646 mV
	if data.BatteryVoltage == nil {
		t.Error("BatteryVoltage should not be nil")
	} else if *data.BatteryVoltage != 3646 {
		t.Errorf("BatteryVoltage = %v, want 3646", *data.BatteryVoltage)
	}

	// TX Power: 20 dBm
	if data.TxPower == nil {
		t.Error("TxPower should not be nil")
	} else if *data.TxPower != 20 {
		t.Errorf("TxPower = %v, want 20", *data.TxPower)
	}

	// Movement Counter: 254
	if data.MovementCounter == nil {
		t.Error("MovementCounter should not be nil")
	} else if *data.MovementCounter != 254 {
		t.Errorf("MovementCounter = %v, want 254", *data.MovementCounter)
	}

	// Measurement Sequence: 65534
	if data.MeasurementSequence == nil {
		t.Error("MeasurementSequence should not be nil")
	} else if *data.MeasurementSequence != 65534 {
		t.Errorf("MeasurementSequence = %v, want 65534", *data.MeasurementSequence)
	}
}

// TestDecodeFormat5_MinimumValues tests decoding with minimum values from official spec.
func TestDecodeFormat5_MinimumValues(t *testing.T) {
	// Test vector: minimum values
	// Raw: 0x058001000000008001800180010000000000CBB8334C884F
	raw, err := hex.DecodeString("058001000000008001800180010000000000CBB8334C884F")
	if err != nil {
		t.Fatalf("Failed to decode hex: %v", err)
	}

	data, err := DecodeFormat5(raw)
	if err != nil {
		t.Fatalf("DecodeFormat5 failed: %v", err)
	}

	// Temperature: -163.835°C
	if data.Temperature == nil {
		t.Error("Temperature should not be nil")
	} else if !floatEquals(*data.Temperature, -163.835, 0.001) {
		t.Errorf("Temperature = %v, want -163.835", *data.Temperature)
	}

	// Humidity: 0.0%
	if data.Humidity == nil {
		t.Error("Humidity should not be nil")
	} else if !floatEquals(*data.Humidity, 0.0, 0.001) {
		t.Errorf("Humidity = %v, want 0.0", *data.Humidity)
	}

	// Pressure: 50000 Pa
	if data.Pressure == nil {
		t.Error("Pressure should not be nil")
	} else if *data.Pressure != 50000 {
		t.Errorf("Pressure = %v, want 50000", *data.Pressure)
	}

	// Acceleration X: -32.767 G
	if data.AccelerationX == nil {
		t.Error("AccelerationX should not be nil")
	} else if !floatEquals(*data.AccelerationX, -32.767, 0.001) {
		t.Errorf("AccelerationX = %v, want -32.767", *data.AccelerationX)
	}

	// Acceleration Y: -32.767 G
	if data.AccelerationY == nil {
		t.Error("AccelerationY should not be nil")
	} else if !floatEquals(*data.AccelerationY, -32.767, 0.001) {
		t.Errorf("AccelerationY = %v, want -32.767", *data.AccelerationY)
	}

	// Acceleration Z: -32.767 G
	if data.AccelerationZ == nil {
		t.Error("AccelerationZ should not be nil")
	} else if !floatEquals(*data.AccelerationZ, -32.767, 0.001) {
		t.Errorf("AccelerationZ = %v, want -32.767", *data.AccelerationZ)
	}

	// Battery Voltage: 1600 mV
	if data.BatteryVoltage == nil {
		t.Error("BatteryVoltage should not be nil")
	} else if *data.BatteryVoltage != 1600 {
		t.Errorf("BatteryVoltage = %v, want 1600", *data.BatteryVoltage)
	}

	// TX Power: -40 dBm
	if data.TxPower == nil {
		t.Error("TxPower should not be nil")
	} else if *data.TxPower != -40 {
		t.Errorf("TxPower = %v, want -40", *data.TxPower)
	}

	// Movement Counter: 0
	if data.MovementCounter == nil {
		t.Error("MovementCounter should not be nil")
	} else if *data.MovementCounter != 0 {
		t.Errorf("MovementCounter = %v, want 0", *data.MovementCounter)
	}

	// Measurement Sequence: 0
	if data.MeasurementSequence == nil {
		t.Error("MeasurementSequence should not be nil")
	} else if *data.MeasurementSequence != 0 {
		t.Errorf("MeasurementSequence = %v, want 0", *data.MeasurementSequence)
	}
}

// TestDecodeFormat5_InvalidValues tests decoding with invalid/unavailable values.
func TestDecodeFormat5_InvalidValues(t *testing.T) {
	// Test vector: invalid values
	// Raw: 0x058000FFFFFFFF800080008000FFFFFFFFFFFFFFFFFFFFFF
	raw, err := hex.DecodeString("058000FFFFFFFF800080008000FFFFFFFFFFFFFFFFFFFFFF")
	if err != nil {
		t.Fatalf("Failed to decode hex: %v", err)
	}

	data, err := DecodeFormat5(raw)
	if err != nil {
		t.Fatalf("DecodeFormat5 failed: %v", err)
	}

	// All fields should be nil
	if data.Temperature != nil {
		t.Error("Temperature should be nil for invalid value")
	}
	if data.Humidity != nil {
		t.Error("Humidity should be nil for invalid value")
	}
	if data.Pressure != nil {
		t.Error("Pressure should be nil for invalid value")
	}
	if data.AccelerationX != nil {
		t.Error("AccelerationX should be nil for invalid value")
	}
	if data.AccelerationY != nil {
		t.Error("AccelerationY should be nil for invalid value")
	}
	if data.AccelerationZ != nil {
		t.Error("AccelerationZ should be nil for invalid value")
	}
	if data.BatteryVoltage != nil {
		t.Error("BatteryVoltage should be nil for invalid value")
	}
	if data.TxPower != nil {
		t.Error("TxPower should be nil for invalid value")
	}
	if data.MovementCounter != nil {
		t.Error("MovementCounter should be nil for invalid value")
	}
	if data.MeasurementSequence != nil {
		t.Error("MeasurementSequence should be nil for invalid value")
	}
	if data.MACAddress != nil {
		t.Error("MACAddress should be nil for invalid value")
	}
}

// TestDecodeFormat5_Errors tests error conditions.
func TestDecodeFormat5_Errors(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{
			name:    "too short",
			data:    []byte{0x05, 0x12, 0xFC},
			wantErr: true,
		},
		{
			name:    "too long",
			data:    make([]byte, 25),
			wantErr: true,
		},
		{
			name:    "wrong format",
			data:    make([]byte, 24),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DecodeFormat5(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeFormat5() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestEncodeFormat5_RoundTrip tests that encode-decode is bidirectional.
func TestEncodeFormat5_RoundTrip(t *testing.T) {
	tests := []struct {
		name string
		hex  string
	}{
		{
			name: "valid data",
			hex:  "0512FC5394C37C0004FFFC040CAC364200CDCBB8334C884F",
		},
		{
			name: "maximum values",
			hex:  "057FFFFFFEFFFE7FFF7FFF7FFFFFDEFEFFFECBB8334C884F",
		},
		{
			name: "minimum values",
			hex:  "058001000000008001800180010000000000CBB8334C884F",
		},
		{
			name: "invalid values",
			hex:  "058000FFFFFFFF800080008000FFFFFFFFFFFFFFFFFFFFFF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw, err := hex.DecodeString(tt.hex)
			if err != nil {
				t.Fatalf("Failed to decode hex: %v", err)
			}

			// Decode
			decoded, err := DecodeFormat5(raw)
			if err != nil {
				t.Fatalf("DecodeFormat5 failed: %v", err)
			}

			// Encode
			encoded, err := EncodeFormat5(decoded)
			if err != nil {
				t.Fatalf("EncodeFormat5 failed: %v", err)
			}

			// Compare (case insensitive by comparing with original raw bytes)
			if hex.EncodeToString(encoded) != hex.EncodeToString(raw) {
				t.Errorf("Round trip failed:\ngot  %s\nwant %s", hex.EncodeToString(encoded), hex.EncodeToString(raw))
			}
		})
	}
}

// TestEncodeFormat5_FullyPopulated tests encoding with all fields populated.
func TestEncodeFormat5_FullyPopulated(t *testing.T) {
	temp := 24.3
	hum := 53.49
	pressure := 100044
	accX := 0.004
	accY := -0.004
	accZ := 1.036
	batt := 2977
	tx := 4
	movement := uint8(66)
	sequence := uint16(205)
	mac := common.MACAddress{0xCB, 0xB8, 0x33, 0x4C, 0x88, 0x4F}

	data := &Format5Data{
		Temperature:         &temp,
		Humidity:            &hum,
		Pressure:            &pressure,
		AccelerationX:       &accX,
		AccelerationY:       &accY,
		AccelerationZ:       &accZ,
		BatteryVoltage:      &batt,
		TxPower:             &tx,
		MovementCounter:     &movement,
		MeasurementSequence: &sequence,
		MACAddress:          &mac,
	}

	encoded, err := EncodeFormat5(data)
	if err != nil {
		t.Fatalf("EncodeFormat5 failed: %v", err)
	}

	// Should match the official test vector
	expected := "0512FC5394C37C0004FFFC040CAC364200CDCBB8334C884F"
	if hex.EncodeToString(encoded) != strings.ToLower(expected) {
		t.Errorf("Encoded data mismatch:\ngot  %s\nwant %s", hex.EncodeToString(encoded), strings.ToLower(expected))
	}

	// Round trip test
	decoded, err := DecodeFormat5(encoded)
	if err != nil {
		t.Fatalf("DecodeFormat5 failed: %v", err)
	}

	// Verify all fields match (within quantization tolerance)
	if decoded.Temperature == nil || !floatEquals(*decoded.Temperature, temp, 0.005) {
		t.Errorf("Temperature = %v, want %v", decoded.Temperature, temp)
	}
	if decoded.Humidity == nil || !floatEquals(*decoded.Humidity, hum, 0.01) {
		t.Errorf("Humidity = %v, want %v", decoded.Humidity, hum)
	}
	if decoded.Pressure == nil || *decoded.Pressure != pressure {
		t.Errorf("Pressure = %v, want %v", decoded.Pressure, pressure)
	}
}

// TestEncodeFormat5_MissingFields tests encoding with some fields missing (nil).
func TestEncodeFormat5_MissingFields(t *testing.T) {
	// Only populate some fields
	temp := 20.5
	pressure := 101325

	data := &Format5Data{
		Temperature: &temp,
		Pressure:    &pressure,
		// Other fields are nil
	}

	encoded, err := EncodeFormat5(data)
	if err != nil {
		t.Fatalf("EncodeFormat5 failed: %v", err)
	}

	if len(encoded) != 24 {
		t.Errorf("Encoded length = %d, want 24", len(encoded))
	}

	// Decode and verify
	decoded, err := DecodeFormat5(encoded)
	if err != nil {
		t.Fatalf("DecodeFormat5 failed: %v", err)
	}

	// Check that populated fields are preserved
	if decoded.Temperature == nil {
		t.Error("Temperature should not be nil")
	} else if !floatEquals(*decoded.Temperature, temp, 0.005) {
		t.Errorf("Temperature = %v, want %v", *decoded.Temperature, temp)
	}

	if decoded.Pressure == nil {
		t.Error("Pressure should not be nil")
	} else if *decoded.Pressure != pressure {
		t.Errorf("Pressure = %v, want %v", *decoded.Pressure, pressure)
	}

	// Check that nil fields are properly encoded as invalid/unavailable
	if decoded.Humidity != nil {
		t.Errorf("Humidity should be nil, got %v", *decoded.Humidity)
	}
	if decoded.AccelerationX != nil {
		t.Errorf("AccelerationX should be nil, got %v", *decoded.AccelerationX)
	}
	if decoded.BatteryVoltage != nil {
		t.Errorf("BatteryVoltage should be nil, got %v", *decoded.BatteryVoltage)
	}
	if decoded.TxPower != nil {
		t.Errorf("TxPower should be nil, got %v", *decoded.TxPower)
	}
	if decoded.MovementCounter != nil {
		t.Errorf("MovementCounter should be nil, got %v", *decoded.MovementCounter)
	}
	if decoded.MeasurementSequence != nil {
		t.Errorf("MeasurementSequence should be nil, got %v", *decoded.MeasurementSequence)
	}
	if decoded.MACAddress != nil {
		t.Errorf("MACAddress should be nil, got %v", decoded.MACAddress)
	}
}

// TestEncodeFormat5ManufacturerData tests encoding with manufacturer data prefix.
func TestEncodeFormat5ManufacturerData(t *testing.T) {
	temp := 24.3
	hum := 53.49
	pressure := 100044
	mac := common.MACAddress{0xCB, 0xB8, 0x33, 0x4C, 0x88, 0x4F}

	data := &Format5Data{
		Temperature: &temp,
		Humidity:    &hum,
		Pressure:    &pressure,
		MACAddress:  &mac,
	}

	encoded, err := EncodeFormat5ManufacturerData(data)
	if err != nil {
		t.Fatalf("EncodeFormat5ManufacturerData failed: %v", err)
	}

	// Should be 26 bytes: 2 manufacturer ID + 24 payload
	if len(encoded) != 26 {
		t.Errorf("Encoded length = %d, want 26", len(encoded))
	}

	// Check manufacturer ID (little-endian: 0x0499 -> 0x99 0x04)
	if encoded[0] != 0x99 || encoded[1] != 0x04 {
		t.Errorf("Manufacturer ID = %02X%02X, want 9904", encoded[0], encoded[1])
	}

	// Check format byte
	if encoded[2] != 0x05 {
		t.Errorf("Format byte = %02X, want 05", encoded[2])
	}

	// Verify the payload matches EncodeFormat5 output
	payload, err := EncodeFormat5(data)
	if err != nil {
		t.Fatalf("EncodeFormat5 failed: %v", err)
	}

	if !bytesEqual(encoded[2:], payload) {
		t.Error("Manufacturer data payload does not match EncodeFormat5 output")
	}
}

// TestEncodeFormat5ManufacturerData_RoundTrip tests that manufacturer data
// can be decoded using ParseManufacturerData (if implemented).
func TestEncodeFormat5ManufacturerData_RoundTrip(t *testing.T) {
	temp := 20.0
	hum := 50.0
	pressure := 100000

	data := &Format5Data{
		Temperature: &temp,
		Humidity:    &hum,
		Pressure:    &pressure,
	}

	encoded, err := EncodeFormat5ManufacturerData(data)
	if err != nil {
		t.Fatalf("EncodeFormat5ManufacturerData failed: %v", err)
	}

	// Remove manufacturer ID prefix and decode the payload
	payload := encoded[2:]
	decoded, err := DecodeFormat5(payload)
	if err != nil {
		t.Fatalf("DecodeFormat5 failed: %v", err)
	}

	// Verify fields (within quantization tolerance)
	if decoded.Temperature == nil || !floatEquals(*decoded.Temperature, temp, 0.01) {
		t.Errorf("Temperature = %v, want %v", decoded.Temperature, temp)
	}
	if decoded.Humidity == nil || !floatEquals(*decoded.Humidity, hum, 0.01) {
		t.Errorf("Humidity = %v, want %v", decoded.Humidity, hum)
	}
	if decoded.Pressure == nil || *decoded.Pressure != pressure {
		t.Errorf("Pressure = %v, want %v", decoded.Pressure, pressure)
	}
}

// TestEncodeFormat5_NilData tests error handling for nil input.
func TestEncodeFormat5_NilData(t *testing.T) {
	_, err := EncodeFormat5(nil)
	if err == nil {
		t.Error("EncodeFormat5 should return error for nil data")
	}
}

// TestEncodeFormat5ManufacturerData_NilData tests error handling for nil input.
func TestEncodeFormat5ManufacturerData_NilData(t *testing.T) {
	_, err := EncodeFormat5ManufacturerData(nil)
	if err == nil {
		t.Error("EncodeFormat5ManufacturerData should return error for nil data")
	}
}

// bytesEqual compares two byte slices for equality.
func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

