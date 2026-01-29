package tag

import (
	"encoding/hex"
	"math"
	"testing"
)

func TestParseManufacturerData_EmptyData(t *testing.T) {
	_, err := ParseManufacturerData([]byte{})
	if err == nil {
		t.Error("expected error for empty data")
	}
}

func TestParseManufacturerData_UnsupportedFormat(t *testing.T) {
	data := []byte{0x03} // Format 3
	_, err := ParseManufacturerData(data)
	if err == nil {
		t.Error("expected error for unsupported format")
	}
}

func TestParseManufacturerData_Format5_InvalidLength(t *testing.T) {
	data := []byte{0x05, 0x12, 0xFC} // Too short
	_, err := ParseManufacturerData(data)
	if err == nil {
		t.Error("expected error for invalid length")
	}
}

func TestParseFormat5_ValidData(t *testing.T) {
	// Test vector from specification: valid data
	// Raw: 0x0512FC5394C37C0004FFFC040CAC364200CDCBB8334C884F
	dataHex := "0512FC5394C37C0004FFFC040CAC364200CDCBB8334C884F"
	data, err := hex.DecodeString(dataHex)
	if err != nil {
		t.Fatalf("failed to decode hex: %v", err)
	}

	m, err := ParseManufacturerData(data)
	if err != nil {
		t.Fatalf("ParseManufacturerData failed: %v", err)
	}

	// Temperature: 24.3 C
	if m.Temperature == nil {
		t.Error("temperature should not be nil")
	} else if math.Abs(*m.Temperature-24.3) > 0.0001 {
		t.Errorf("temperature = %v, want 24.3", *m.Temperature)
	}

	// Humidity: 53.49%
	if m.Humidity == nil {
		t.Error("humidity should not be nil")
	} else if math.Abs(*m.Humidity-53.49) > 0.01 {
		t.Errorf("humidity = %v, want 53.49", *m.Humidity)
	}

	// Pressure: 100044 Pa
	if m.Pressure == nil {
		t.Error("pressure should not be nil")
	} else if *m.Pressure != 100044 {
		t.Errorf("pressure = %v, want 100044", *m.Pressure)
	}

	// Acceleration X: 0.004 G
	if m.AccelerationX == nil {
		t.Error("accelerationX should not be nil")
	} else if math.Abs(*m.AccelerationX-0.004) > 0.0001 {
		t.Errorf("accelerationX = %v, want 0.004", *m.AccelerationX)
	}

	// Acceleration Y: -0.004 G
	if m.AccelerationY == nil {
		t.Error("accelerationY should not be nil")
	} else if math.Abs(*m.AccelerationY-(-0.004)) > 0.0001 {
		t.Errorf("accelerationY = %v, want -0.004", *m.AccelerationY)
	}

	// Acceleration Z: 1.036 G
	if m.AccelerationZ == nil {
		t.Error("accelerationZ should not be nil")
	} else if math.Abs(*m.AccelerationZ-1.036) > 0.0001 {
		t.Errorf("accelerationZ = %v, want 1.036", *m.AccelerationZ)
	}

	// Battery voltage: 2.977 V (2977 mV)
	if m.BatteryVoltage == nil {
		t.Error("batteryVoltage should not be nil")
	} else if *m.BatteryVoltage != 2977 {
		t.Errorf("batteryVoltage = %v, want 2977", *m.BatteryVoltage)
	}

	// TX Power: 4 dBm
	if m.TxPower == nil {
		t.Error("txPower should not be nil")
	} else if *m.TxPower != 4 {
		t.Errorf("txPower = %v, want 4", *m.TxPower)
	}

	// Movement counter: 66
	if m.MovementCounter == nil {
		t.Error("movementCounter should not be nil")
	} else if *m.MovementCounter != 66 {
		t.Errorf("movementCounter = %v, want 66", *m.MovementCounter)
	}

	// Measurement sequence: 205
	if m.MeasurementSequence == nil {
		t.Error("measurementSequence should not be nil")
	} else if *m.MeasurementSequence != 205 {
		t.Errorf("measurementSequence = %v, want 205", *m.MeasurementSequence)
	}

	// MAC address: CBB8334C884F
	if m.MACAddress == nil {
		t.Error("macAddress should not be nil")
	} else {
		expectedMAC := [6]byte{0xCB, 0xB8, 0x33, 0x4C, 0x88, 0x4F}
		if *m.MACAddress != expectedMAC {
			t.Errorf("macAddress = %v, want %v", *m.MACAddress, expectedMAC)
		}
	}
}

func TestParseFormat5_MaximumValues(t *testing.T) {
	// Test vector from specification: maximum values
	// Raw: 0x057FFFFFFEFFFE7FFF7FFF7FFFFFDEFEFFFECBB8334C884F
	dataHex := "057FFFFFFEFFFE7FFF7FFF7FFFFFDEFEFFFECBB8334C884F"
	data, err := hex.DecodeString(dataHex)
	if err != nil {
		t.Fatalf("failed to decode hex: %v", err)
	}

	m, err := ParseManufacturerData(data)
	if err != nil {
		t.Fatalf("ParseManufacturerData failed: %v", err)
	}

	// Temperature: 163.835 C
	if m.Temperature == nil {
		t.Error("temperature should not be nil")
	} else if math.Abs(*m.Temperature-163.835) > 0.0001 {
		t.Errorf("temperature = %v, want 163.835", *m.Temperature)
	}

	// Humidity: 163.8350%
	if m.Humidity == nil {
		t.Error("humidity should not be nil")
	} else if math.Abs(*m.Humidity-163.835) > 0.001 {
		t.Errorf("humidity = %v, want 163.835", *m.Humidity)
	}

	// Pressure: 115534 Pa
	if m.Pressure == nil {
		t.Error("pressure should not be nil")
	} else if *m.Pressure != 115534 {
		t.Errorf("pressure = %v, want 115534", *m.Pressure)
	}

	// Acceleration X: 32.767 G
	if m.AccelerationX == nil {
		t.Error("accelerationX should not be nil")
	} else if math.Abs(*m.AccelerationX-32.767) > 0.0001 {
		t.Errorf("accelerationX = %v, want 32.767", *m.AccelerationX)
	}

	// Acceleration Y: 32.767 G
	if m.AccelerationY == nil {
		t.Error("accelerationY should not be nil")
	} else if math.Abs(*m.AccelerationY-32.767) > 0.0001 {
		t.Errorf("accelerationY = %v, want 32.767", *m.AccelerationY)
	}

	// Acceleration Z: 32.767 G
	if m.AccelerationZ == nil {
		t.Error("accelerationZ should not be nil")
	} else if math.Abs(*m.AccelerationZ-32.767) > 0.0001 {
		t.Errorf("accelerationZ = %v, want 32.767", *m.AccelerationZ)
	}

	// Battery voltage: 3.646 V (3646 mV)
	if m.BatteryVoltage == nil {
		t.Error("batteryVoltage should not be nil")
	} else if *m.BatteryVoltage != 3646 {
		t.Errorf("batteryVoltage = %v, want 3646", *m.BatteryVoltage)
	}

	// TX Power: 20 dBm
	if m.TxPower == nil {
		t.Error("txPower should not be nil")
	} else if *m.TxPower != 20 {
		t.Errorf("txPower = %v, want 20", *m.TxPower)
	}

	// Movement counter: 254
	if m.MovementCounter == nil {
		t.Error("movementCounter should not be nil")
	} else if *m.MovementCounter != 254 {
		t.Errorf("movementCounter = %v, want 254", *m.MovementCounter)
	}

	// Measurement sequence: 65534
	if m.MeasurementSequence == nil {
		t.Error("measurementSequence should not be nil")
	} else if *m.MeasurementSequence != 65534 {
		t.Errorf("measurementSequence = %v, want 65534", *m.MeasurementSequence)
	}

	// MAC address: CBB8334C884F
	if m.MACAddress == nil {
		t.Error("macAddress should not be nil")
	} else {
		expectedMAC := [6]byte{0xCB, 0xB8, 0x33, 0x4C, 0x88, 0x4F}
		if *m.MACAddress != expectedMAC {
			t.Errorf("macAddress = %v, want %v", *m.MACAddress, expectedMAC)
		}
	}
}

func TestParseFormat5_MinimumValues(t *testing.T) {
	// Test vector from specification: minimum values
	// Raw: 0x058001000000008001800180010000000000CBB8334C884F
	dataHex := "058001000000008001800180010000000000CBB8334C884F"
	data, err := hex.DecodeString(dataHex)
	if err != nil {
		t.Fatalf("failed to decode hex: %v", err)
	}

	m, err := ParseManufacturerData(data)
	if err != nil {
		t.Fatalf("ParseManufacturerData failed: %v", err)
	}

	// Temperature: -163.835 C
	if m.Temperature == nil {
		t.Error("temperature should not be nil")
	} else if math.Abs(*m.Temperature-(-163.835)) > 0.0001 {
		t.Errorf("temperature = %v, want -163.835", *m.Temperature)
	}

	// Humidity: 0.000%
	if m.Humidity == nil {
		t.Error("humidity should not be nil")
	} else if math.Abs(*m.Humidity-0.0) > 0.0001 {
		t.Errorf("humidity = %v, want 0.0", *m.Humidity)
	}

	// Pressure: 50000 Pa
	if m.Pressure == nil {
		t.Error("pressure should not be nil")
	} else if *m.Pressure != 50000 {
		t.Errorf("pressure = %v, want 50000", *m.Pressure)
	}

	// Acceleration X: -32.767 G
	if m.AccelerationX == nil {
		t.Error("accelerationX should not be nil")
	} else if math.Abs(*m.AccelerationX-(-32.767)) > 0.0001 {
		t.Errorf("accelerationX = %v, want -32.767", *m.AccelerationX)
	}

	// Acceleration Y: -32.767 G
	if m.AccelerationY == nil {
		t.Error("accelerationY should not be nil")
	} else if math.Abs(*m.AccelerationY-(-32.767)) > 0.0001 {
		t.Errorf("accelerationY = %v, want -32.767", *m.AccelerationY)
	}

	// Acceleration Z: -32.767 G
	if m.AccelerationZ == nil {
		t.Error("accelerationZ should not be nil")
	} else if math.Abs(*m.AccelerationZ-(-32.767)) > 0.0001 {
		t.Errorf("accelerationZ = %v, want -32.767", *m.AccelerationZ)
	}

	// Battery voltage: 1.600 V (1600 mV)
	if m.BatteryVoltage == nil {
		t.Error("batteryVoltage should not be nil")
	} else if *m.BatteryVoltage != 1600 {
		t.Errorf("batteryVoltage = %v, want 1600", *m.BatteryVoltage)
	}

	// TX Power: -40 dBm
	if m.TxPower == nil {
		t.Error("txPower should not be nil")
	} else if *m.TxPower != -40 {
		t.Errorf("txPower = %v, want -40", *m.TxPower)
	}

	// Movement counter: 0
	if m.MovementCounter == nil {
		t.Error("movementCounter should not be nil")
	} else if *m.MovementCounter != 0 {
		t.Errorf("movementCounter = %v, want 0", *m.MovementCounter)
	}

	// Measurement sequence: 0
	if m.MeasurementSequence == nil {
		t.Error("measurementSequence should not be nil")
	} else if *m.MeasurementSequence != 0 {
		t.Errorf("measurementSequence = %v, want 0", *m.MeasurementSequence)
	}

	// MAC address: CBB8334C884F
	if m.MACAddress == nil {
		t.Error("macAddress should not be nil")
	} else {
		expectedMAC := [6]byte{0xCB, 0xB8, 0x33, 0x4C, 0x88, 0x4F}
		if *m.MACAddress != expectedMAC {
			t.Errorf("macAddress = %v, want %v", *m.MACAddress, expectedMAC)
		}
	}
}

func TestParseFormat5_InvalidValues(t *testing.T) {
	// Test vector from specification: invalid values
	// Raw: 0x058000FFFFFFFF800080008000FFFFFFFFFFFFFFFFFFFFFF
	dataHex := "058000FFFFFFFF800080008000FFFFFFFFFFFFFFFFFFFFFF"
	data, err := hex.DecodeString(dataHex)
	if err != nil {
		t.Fatalf("failed to decode hex: %v", err)
	}

	m, err := ParseManufacturerData(data)
	if err != nil {
		t.Fatalf("ParseManufacturerData failed: %v", err)
	}

	// All fields should be nil (not available)
	if m.Temperature != nil {
		t.Error("temperature should be nil")
	}
	if m.Humidity != nil {
		t.Error("humidity should be nil")
	}
	if m.Pressure != nil {
		t.Error("pressure should be nil")
	}
	if m.AccelerationX != nil {
		t.Error("accelerationX should be nil")
	}
	if m.AccelerationY != nil {
		t.Error("accelerationY should be nil")
	}
	if m.AccelerationZ != nil {
		t.Error("accelerationZ should be nil")
	}
	if m.BatteryVoltage != nil {
		t.Error("batteryVoltage should be nil")
	}
	if m.TxPower != nil {
		t.Error("txPower should be nil")
	}
	if m.MovementCounter != nil {
		t.Error("movementCounter should be nil")
	}
	if m.MeasurementSequence != nil {
		t.Error("measurementSequence should be nil")
	}
	if m.MACAddress != nil {
		t.Error("macAddress should be nil")
	}
}
