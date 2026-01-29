package tag

import (
	"encoding/hex"
	"testing"
)

// TestDecodeFormat3_ValidData tests decoding with valid data from official spec.
func TestDecodeFormat3_ValidData(t *testing.T) {
	// Test vector from official specification
	// Raw: 0x03291A1ECE1EFC18F94202CA0B53
	raw, err := hex.DecodeString("03291A1ECE1EFC18F94202CA0B53")
	if err != nil {
		t.Fatalf("Failed to decode hex: %v", err)
	}

	data, err := DecodeFormat3(raw)
	if err != nil {
		t.Fatalf("DecodeFormat3 failed: %v", err)
	}

	// Humidity: 20.5%
	if data.Humidity == nil {
		t.Error("Humidity should not be nil")
	} else if !floatEquals(*data.Humidity, 20.5, 0.1) {
		t.Errorf("Humidity = %v, want 20.5", *data.Humidity)
	}

	// Temperature: 26.3°C
	if data.Temperature == nil {
		t.Error("Temperature should not be nil")
	} else if !floatEquals(*data.Temperature, 26.3, 0.01) {
		t.Errorf("Temperature = %v, want 26.3", *data.Temperature)
	}

	// Pressure: 102766 Pa
	if data.Pressure == nil {
		t.Error("Pressure should not be nil")
	} else if *data.Pressure != 102766 {
		t.Errorf("Pressure = %v, want 102766", *data.Pressure)
	}

	// Acceleration X: -1.000 G
	if data.AccelerationX == nil {
		t.Error("AccelerationX should not be nil")
	} else if !floatEquals(*data.AccelerationX, -1.000, 0.001) {
		t.Errorf("AccelerationX = %v, want -1.000", *data.AccelerationX)
	}

	// Acceleration Y: -1.726 G
	if data.AccelerationY == nil {
		t.Error("AccelerationY should not be nil")
	} else if !floatEquals(*data.AccelerationY, -1.726, 0.001) {
		t.Errorf("AccelerationY = %v, want -1.726", *data.AccelerationY)
	}

	// Acceleration Z: 0.714 G
	if data.AccelerationZ == nil {
		t.Error("AccelerationZ should not be nil")
	} else if !floatEquals(*data.AccelerationZ, 0.714, 0.001) {
		t.Errorf("AccelerationZ = %v, want 0.714", *data.AccelerationZ)
	}

	// Battery Voltage: 2899 mV
	if data.BatteryVoltage == nil {
		t.Error("BatteryVoltage should not be nil")
	} else if *data.BatteryVoltage != 2899 {
		t.Errorf("BatteryVoltage = %v, want 2899", *data.BatteryVoltage)
	}
}

// TestDecodeFormat3_MaximumValues tests decoding with maximum values from official spec.
func TestDecodeFormat3_MaximumValues(t *testing.T) {
	// Test vector: maximum values
	// Raw: 0x03FF7F63FFFF7FFF7FFF7FFFFFFF
	raw, err := hex.DecodeString("03FF7F63FFFF7FFF7FFF7FFFFFFF")
	if err != nil {
		t.Fatalf("Failed to decode hex: %v", err)
	}

	data, err := DecodeFormat3(raw)
	if err != nil {
		t.Fatalf("DecodeFormat3 failed: %v", err)
	}

	// Humidity: 127.5%
	if data.Humidity == nil {
		t.Error("Humidity should not be nil")
	} else if !floatEquals(*data.Humidity, 127.5, 0.1) {
		t.Errorf("Humidity = %v, want 127.5", *data.Humidity)
	}

	// Temperature: 127.99°C
	if data.Temperature == nil {
		t.Error("Temperature should not be nil")
	} else if !floatEquals(*data.Temperature, 127.99, 0.01) {
		t.Errorf("Temperature = %v, want 127.99", *data.Temperature)
	}

	// Pressure: 115535 Pa
	if data.Pressure == nil {
		t.Error("Pressure should not be nil")
	} else if *data.Pressure != 115535 {
		t.Errorf("Pressure = %v, want 115535", *data.Pressure)
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

	// Battery Voltage: 65535 mV
	if data.BatteryVoltage == nil {
		t.Error("BatteryVoltage should not be nil")
	} else if *data.BatteryVoltage != 65535 {
		t.Errorf("BatteryVoltage = %v, want 65535", *data.BatteryVoltage)
	}
}

// TestDecodeFormat3_MinimumValues tests decoding with minimum values from official spec.
func TestDecodeFormat3_MinimumValues(t *testing.T) {
	// Test vector: minimum values
	// Raw: 0x0300FF6300008001800180010000
	raw, err := hex.DecodeString("0300FF6300008001800180010000")
	if err != nil {
		t.Fatalf("Failed to decode hex: %v", err)
	}

	data, err := DecodeFormat3(raw)
	if err != nil {
		t.Fatalf("DecodeFormat3 failed: %v", err)
	}

	// Humidity: 0.0%
	if data.Humidity != nil {
		t.Error("Humidity should be nil (zero value indicates unavailable)")
	}

	// Temperature: -127.99°C
	if data.Temperature == nil {
		t.Error("Temperature should not be nil")
	} else if !floatEquals(*data.Temperature, -127.99, 0.01) {
		t.Errorf("Temperature = %v, want -127.99", *data.Temperature)
	}

	// Pressure: 50000 Pa
	if data.Pressure != nil {
		t.Error("Pressure should be nil (zero value indicates unavailable)")
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

	// Battery Voltage: 0 mV (unavailable)
	if data.BatteryVoltage != nil {
		t.Error("BatteryVoltage should be nil (zero value indicates unavailable)")
	}
}

// TestDecodeFormat3_Errors tests error conditions.
func TestDecodeFormat3_Errors(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{
			name:    "too short",
			data:    []byte{0x03, 0x29, 0x1A},
			wantErr: true,
		},
		{
			name:    "too long",
			data:    make([]byte, 15),
			wantErr: true,
		},
		{
			name:    "wrong format",
			data:    make([]byte, 14),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DecodeFormat3(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeFormat3() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestEncodeFormat3_RoundTrip tests that encode-decode is bidirectional.
func TestEncodeFormat3_RoundTrip(t *testing.T) {
	tests := []struct {
		name string
		hex  string
	}{
		{
			name: "valid data",
			hex:  "03291A1ECE1EFC18F94202CA0B53",
		},
		{
			name: "maximum values",
			hex:  "03FF7F63FFFF7FFF7FFF7FFFFFFF",
		},
		{
			name: "minimum values",
			hex:  "0300FF6300008001800180010000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw, err := hex.DecodeString(tt.hex)
			if err != nil {
				t.Fatalf("Failed to decode hex: %v", err)
			}

			// Decode
			decoded, err := DecodeFormat3(raw)
			if err != nil {
				t.Fatalf("DecodeFormat3 failed: %v", err)
			}

			// Encode
			encoded, err := EncodeFormat3(decoded)
			if err != nil {
				t.Fatalf("EncodeFormat3 failed: %v", err)
			}

			// Compare (case insensitive)
			gotHex := hex.EncodeToString(encoded)
			wantHex := tt.hex
			if gotHex != wantHex && gotHex != hex.EncodeToString(raw) {
				t.Errorf("Round trip failed:\ngot  %s\nwant %s", gotHex, wantHex)
			}
		})
	}
}
