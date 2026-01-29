// Package common provides shared types and utilities for Ruuvi sensor data.
package common

import (
	"fmt"
	"math"
)

// Temperature represents a temperature measurement in degrees Celsius.
// A nil pointer indicates an invalid or unavailable reading.
type Temperature struct {
	Celsius float64
}

// Humidity represents a relative humidity measurement in percent.
// A nil pointer indicates an invalid or unavailable reading.
type Humidity struct {
	Percent float64
}

// Pressure represents an atmospheric pressure measurement in Pascals.
// A nil pointer indicates an invalid or unavailable reading.
type Pressure struct {
	Pascals int
}

// Acceleration represents an acceleration measurement in G (gravitational force).
// A nil pointer indicates an invalid or unavailable reading.
type Acceleration struct {
	X float64
	Y float64
	Z float64
}

// BatteryVoltage represents a battery voltage measurement in millivolts.
// A nil pointer indicates an invalid or unavailable reading.
type BatteryVoltage struct {
	Millivolts int
}

// TxPower represents transmit power in dBm.
// A nil pointer indicates an invalid or unavailable reading.
type TxPower struct {
	DBm int
}

// MovementCounter tracks movement detection events.
// A nil pointer indicates an invalid or unavailable reading.
type MovementCounter struct {
	Count uint8
}

// MeasurementSequence tracks measurement sequence numbers for deduplication.
// A nil pointer indicates an invalid or unavailable reading.
type MeasurementSequence struct {
	Number uint16
}

// MACAddress represents a 48-bit MAC address.
// A nil pointer or all 0xFF bytes indicates an invalid or unavailable MAC.
type MACAddress [6]byte

// String returns the MAC address in standard colon-separated hex format.
func (m MACAddress) String() string {
	return fmt.Sprintf("%02X:%02X:%02X:%02X:%02X:%02X",
		m[0], m[1], m[2], m[3], m[4], m[5])
}

// IsInvalid checks if the MAC address is invalid (all 0xFF).
func (m MACAddress) IsInvalid() bool {
	for _, b := range m {
		if b != 0xFF {
			return false
		}
	}
	return true
}

// Float64Ptr returns a pointer to the given float64 value.
// Returns nil if the value is NaN.
func Float64Ptr(v float64) *float64 {
	if math.IsNaN(v) {
		return nil
	}
	return &v
}

// IntPtr returns a pointer to the given int value.
func IntPtr(v int) *int {
	return &v
}

// Uint8Ptr returns a pointer to the given uint8 value.
func Uint8Ptr(v uint8) *uint8 {
	return &v
}

// Uint16Ptr returns a pointer to the given uint16 value.
func Uint16Ptr(v uint16) *uint16 {
	return &v
}

// MACAddressPtr returns a pointer to the given MAC address.
// Returns nil if the MAC address is invalid.
func MACAddressPtr(m MACAddress) *MACAddress {
	if m.IsInvalid() {
		return nil
	}
	return &m
}
