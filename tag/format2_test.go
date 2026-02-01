package tag

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func float64Ptr(v float64) *float64 { return &v }
func intPtr(v int) *int             { return &v }

func TestDecodeFormat2_Valid(t *testing.T) {
	// Build a valid format 2 payload:
	// format byte = 0x02
	// humidity = 50% -> raw = 50 / 0.5 = 100
	// temperature = 25°C -> byte value 25
	// pressure = 101325 Pa -> raw = 101325 - 50000 = 51325
	pressRaw := uint16(101325 - 50000)
	data := make([]byte, 6)
	data[0] = 0x02
	data[1] = 100
	data[2] = byte(int8(25))
	data[3] = 0
	binary.BigEndian.PutUint16(data[4:6], pressRaw)

	got, err := DecodeFormat2(data)
	if err != nil {
		t.Fatalf("DecodeFormat2 returned error: %v", err)
	}

	if got.Humidity == nil || *got.Humidity != 50.0 {
		t.Fatalf("Humidity = %v; want 50.0", got.Humidity)
	}
	if got.Temperature == nil || *got.Temperature != 25.0 {
		t.Fatalf("Temperature = %v; want 25.0", got.Temperature)
	}
	if got.Pressure == nil || *got.Pressure != 101325 {
		t.Fatalf("Pressure = %v; want 101325", got.Pressure)
	}
}

func TestEncodeDecodeFormat2_RoundTrip(t *testing.T) {
	in := &Format2Data{
		Temperature: float64Ptr(18.5),
		Humidity:    float64Ptr(42.5),
		Pressure:    intPtr(100000),
	}

	b, err := EncodeFormat2(in)
	if err != nil {
		t.Fatalf("EncodeFormat2 error: %v", err)
	}

	out, err := DecodeFormat2(b)
	if err != nil {
		t.Fatalf("DecodeFormat2 error: %v", err)
	}

	// Note: Format 2 does not encode fractional temperature or fraction for humidity beyond 0.5 steps.
	if out.Temperature == nil || *out.Temperature != 18.0 {
		t.Fatalf("round-trip Temperature = %v; want 18.0 (fractional parts truncated)", out.Temperature)
	}
	if out.Humidity == nil || *out.Humidity != 42.5 {
		t.Fatalf("round-trip Humidity = %v; want 42.5", out.Humidity)
	}
	if out.Pressure == nil || *out.Pressure != 100000 {
		t.Fatalf("round-trip Pressure = %v; want 100000", out.Pressure)
	}
}

func TestDecodeFormat2_InvalidInputs(t *testing.T) {
	// Wrong length
	if _, err := DecodeFormat2([]byte{0x02, 1}); err == nil {
		t.Fatalf("expected error for invalid length")
	}

	// Wrong format byte
	if _, err := DecodeFormat2([]byte{0x03, 0, 0, 0, 0, 0}); err == nil {
		t.Fatalf("expected error for wrong format byte")
	}
}

func TestEncodeFormat2_NilAndZeroFields(t *testing.T) {
	// Nil data pointer
	if _, err := EncodeFormat2(nil); err == nil {
		t.Fatalf("expected error when encoding nil data")
	}

	// All nil fields should produce zeros after the format byte
	in := &Format2Data{}
	b, err := EncodeFormat2(in)
	if err != nil {
		t.Fatalf("EncodeFormat2 error: %v", err)
	}
	want := []byte{0x02, 0, 0, 0, 0, 0}
	if !reflect.DeepEqual(b, want) {
		t.Fatalf("encoded bytes = %v; want %v", b, want)
	}
}

func TestDecodeFormat2_ZeroFields_YieldNilPointers(t *testing.T) {
	b := []byte{0x02, 0, 0, 0, 0, 0}
	out, err := DecodeFormat2(b)
	if err != nil {
		t.Fatalf("DecodeFormat2 error: %v", err)
	}
	if out.Humidity != nil {
		t.Fatalf("expected nil Humidity for zero field, got %v", out.Humidity)
	}
	if out.Temperature != nil {
		t.Fatalf("expected nil Temperature for zero field, got %v", out.Temperature)
	}
	if out.Pressure != nil {
		t.Fatalf("expected nil Pressure for zero field, got %v", out.Pressure)
	}
}
