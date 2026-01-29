package tag

import (
	"encoding/hex"
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

