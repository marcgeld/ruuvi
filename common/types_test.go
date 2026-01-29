package common

import (
	"math"
	"testing"
)

func TestMACString(t *testing.T) {
	mac := MACAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06}
	want := "01:02:03:04:05:06"
	if s := mac.String(); s != want {
		t.Fatalf("MAC.String() = %q; want %q", s, want)
	}
}

func TestMACIsInvalid(t *testing.T) {
	allFF := MACAddress{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
	if !allFF.IsInvalid() {
		t.Fatal("expected all-FF MAC to be invalid")
	}

	valid := MACAddress{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}
	if valid.IsInvalid() {
		t.Fatal("expected non-FF MAC to be valid")
	}
}

func TestMACAddressPtr(t *testing.T) {
	allFF := MACAddress{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
	if p := MACAddressPtr(allFF); p != nil {
		t.Fatalf("MACAddressPtr(allFF) = %v; want nil", p)
	}

	mac := MACAddress{0, 1, 2, 3, 4, 5}
	p := MACAddressPtr(mac)
	if p == nil {
		t.Fatal("expected non-nil pointer for valid MAC")
	}
	if *p != mac {
		t.Fatalf("dereferenced pointer = %v; want %v", *p, mac)
	}
}

func TestFloat64Ptr(t *testing.T) {
	if p := Float64Ptr(math.NaN()); p != nil {
		t.Fatal("Float64Ptr(NaN) = non-nil; want nil")
	}

	v := 2.71828
	p := Float64Ptr(v)
	if p == nil || *p != v {
		t.Fatalf("Float64Ptr(%v) = %v; want pointer to %v", v, p, v)
	}
}

func TestIntAndUintPtrs(t *testing.T) {
	i := IntPtr(42)
	if i == nil || *i != 42 {
		t.Fatalf("IntPtr(42) = %v; want pointer to 42", i)
	}

	u8 := Uint8Ptr(7)
	if u8 == nil || *u8 != 7 {
		t.Fatalf("Uint8Ptr(7) = %v; want pointer to 7", u8)
	}

	u16 := Uint16Ptr(655)
	if u16 == nil || *u16 != 655 {
		t.Fatalf("Uint16Ptr(655) = %v; want pointer to 655", u16)
	}
}
