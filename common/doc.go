// Package common provides shared types and utilities for Ruuvi sensor data.
//
// This package defines common data types used across different Ruuvi protocols
// and formats. It provides type-safe wrappers for sensor measurements with
// appropriate units and validation.
//
// # Data Types
//
// The package provides structured types for common sensor measurements:
//
// - Temperature: measured in degrees Celsius
// - Humidity: measured in percent
// - Pressure: measured in Pascals
// - Acceleration: measured in G (gravitational force)
// - BatteryVoltage: measured in millivolts
// - TxPower: measured in dBm
// - MACAddress: 48-bit device identifier
//
// # Invalid Values
//
// Throughout the Ruuvi protocols, nil pointers are used to indicate
// invalid or unavailable sensor readings. Helper functions are provided
// to construct pointer types safely.
package common
