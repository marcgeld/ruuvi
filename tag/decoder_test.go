package tag

import (
	"encoding/hex"
	"testing"
)

// TestDetectFormat tests format detection.
func TestDetectFormat(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    DataFormat
		wantErr bool
	}{
		{
			name: "Format 2",
			data: []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x00},
			want: Format2,
		},
		{
			name: "Format 3",
			data: []byte{0x03, 0x00, 0x00, 0x00, 0x00, 0x00},
			want: Format3,
		},
		{
			name: "Format 4",
			data: []byte{0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			want: Format4,
		},
		{
			name: "Format 5",
			data: []byte{0x05, 0x00, 0x00, 0x00, 0x00, 0x00},
			want: Format5,
		},
		{
			name:    "Empty data",
			data:    []byte{},
			wantErr: true,
		},
		{
			name:    "Unknown format",
			data:    []byte{0xFF, 0x00},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DetectFormat(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DetectFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("DetectFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestDecode tests the auto-detect decoder.
func TestDecode(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    DataFormat
		wantErr bool
	}{
		{
			name: "Format 5 valid",
			hex:  "0512FC5394C37C0004FFFC040CAC364200CDCBB8334C884F",
			want: Format5,
		},
		{
			name: "Format 3 valid",
			hex:  "03291A1ECE1EFC18F94202CA0B53",
			want: Format3,
		},
		{
			name:    "Empty data",
			hex:     "",
			wantErr: true,
		},
		{
			name:    "Unknown format",
			hex:     "FF00",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var raw []byte
			var err error
			if tt.hex != "" {
				raw, err = hex.DecodeString(tt.hex)
				if err != nil {
					t.Fatalf("Failed to decode hex: %v", err)
				}
			}

			got, err := Decode(raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			if got.Format != tt.want {
				t.Errorf("Decode() format = %v, want %v", got.Format, tt.want)
			}

			// Verify the correct format-specific data is populated
			switch got.Format {
			case Format2:
				if got.Format2 == nil {
					t.Error("Format2 data should not be nil")
				}
			case Format3:
				if got.Format3 == nil {
					t.Error("Format3 data should not be nil")
				}
			case Format4:
				if got.Format4 == nil {
					t.Error("Format4 data should not be nil")
				}
			case Format5:
				if got.Format5 == nil {
					t.Error("Format5 data should not be nil")
				}
			}
		})
	}
}

// TestDecode_Format5 tests decoding Format 5 data with Decode.
func TestDecode_Format5(t *testing.T) {
	raw, err := hex.DecodeString("0512FC5394C37C0004FFFC040CAC364200CDCBB8334C884F")
	if err != nil {
		t.Fatalf("Failed to decode hex: %v", err)
	}

	decoded, err := Decode(raw)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	if decoded.Format != Format5 {
		t.Errorf("Format = %v, want Format5", decoded.Format)
	}

	if decoded.Format5 == nil {
		t.Fatal("Format5 data is nil")
	}

	// Verify some fields
	if decoded.Format5.Temperature == nil {
		t.Error("Temperature should not be nil")
	} else if !floatEquals(*decoded.Format5.Temperature, 24.3, 0.01) {
		t.Errorf("Temperature = %v, want 24.3", *decoded.Format5.Temperature)
	}

	if decoded.Format5.MACAddress == nil {
		t.Error("MACAddress should not be nil")
	}
}

// TestDecode_Format3 tests decoding Format 3 data with Decode.
func TestDecode_Format3(t *testing.T) {
	raw, err := hex.DecodeString("03291A1ECE1EFC18F94202CA0B53")
	if err != nil {
		t.Fatalf("Failed to decode hex: %v", err)
	}

	decoded, err := Decode(raw)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	if decoded.Format != Format3 {
		t.Errorf("Format = %v, want Format3", decoded.Format)
	}

	if decoded.Format3 == nil {
		t.Fatal("Format3 data is nil")
	}

	// Verify some fields
	if decoded.Format3.Temperature == nil {
		t.Error("Temperature should not be nil")
	} else if !floatEquals(*decoded.Format3.Temperature, 26.3, 0.01) {
		t.Errorf("Temperature = %v, want 26.3", *decoded.Format3.Temperature)
	}
}
