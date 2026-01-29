// Package main demonstrates basic usage of the ruuvi/tag package for decoding
// RuuviTag Bluetooth LE advertisement data.
package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/marcgeld/ruuvi/tag"
)

func main() {
	// Example 1: Decode Format 5 (RAWv2) - Current production format
	fmt.Println("=== Example 1: Format 5 (RAWv2) ===")
	decodeFormat5Example()

	fmt.Println()

	// Example 2: Decode Format 3 (RAWv1) - Deprecated but widely deployed
	fmt.Println("=== Example 2: Format 3 (RAWv1) ===")
	decodeFormat3Example()

	fmt.Println()

	// Example 3: Auto-detect format
	fmt.Println("=== Example 3: Auto-detect Format ===")
	autoDetectExample()
}

func decodeFormat5Example() {
	// Real Format 5 data from official test vector
	// This represents a complete sensor reading from a RuuviTag
	raw, err := hex.DecodeString("0512FC5394C37C0004FFFC040CAC364200CDCBB8334C884F")
	if err != nil {
		log.Fatal(err)
	}

	data, err := tag.DecodeFormat5(raw)
	if err != nil {
		log.Fatal(err)
	}

	// Print all available sensor data
	fmt.Println("Sensor Data:")

	if data.Temperature != nil {
		fmt.Printf("  Temperature:      %.2f°C\n", *data.Temperature)
	}

	if data.Humidity != nil {
		fmt.Printf("  Humidity:         %.2f%%\n", *data.Humidity)
	}

	if data.Pressure != nil {
		fmt.Printf("  Pressure:         %d Pa (%.2f hPa)\n", *data.Pressure, float64(*data.Pressure)/100)
	}

	if data.AccelerationX != nil && data.AccelerationY != nil && data.AccelerationZ != nil {
		fmt.Printf("  Acceleration:     X=%.3f Y=%.3f Z=%.3f G\n",
			*data.AccelerationX, *data.AccelerationY, *data.AccelerationZ)
	}

	if data.BatteryVoltage != nil {
		fmt.Printf("  Battery:          %d mV (%.3f V)\n", *data.BatteryVoltage, float64(*data.BatteryVoltage)/1000)
	}

	if data.TxPower != nil {
		fmt.Printf("  TX Power:         %d dBm\n", *data.TxPower)
	}

	if data.MovementCounter != nil {
		fmt.Printf("  Movement Events:  %d\n", *data.MovementCounter)
	}

	if data.MeasurementSequence != nil {
		fmt.Printf("  Measurement #:    %d\n", *data.MeasurementSequence)
	}

	if data.MACAddress != nil {
		fmt.Printf("  MAC Address:      %s\n", data.MACAddress)
	}
}

func decodeFormat3Example() {
	// Real Format 3 data from official test vector
	raw, err := hex.DecodeString("03291A1ECE1EFC18F94202CA0B53")
	if err != nil {
		log.Fatal(err)
	}

	data, err := tag.DecodeFormat3(raw)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sensor Data:")

	if data.Temperature != nil {
		fmt.Printf("  Temperature:      %.2f°C\n", *data.Temperature)
	}

	if data.Humidity != nil {
		fmt.Printf("  Humidity:         %.1f%%\n", *data.Humidity)
	}

	if data.Pressure != nil {
		fmt.Printf("  Pressure:         %d Pa (%.2f hPa)\n", *data.Pressure, float64(*data.Pressure)/100)
	}

	if data.AccelerationX != nil && data.AccelerationY != nil && data.AccelerationZ != nil {
		fmt.Printf("  Acceleration:     X=%.3f Y=%.3f Z=%.3f G\n",
			*data.AccelerationX, *data.AccelerationY, *data.AccelerationZ)
	}

	if data.BatteryVoltage != nil {
		fmt.Printf("  Battery:          %d mV (%.3f V)\n", *data.BatteryVoltage, float64(*data.BatteryVoltage)/1000)
	}
}

func autoDetectExample() {
	// Multiple formats can be auto-detected
	examples := []struct {
		name string
		hex  string
	}{
		{"Format 5", "0512FC5394C37C0004FFFC040CAC364200CDCBB8334C884F"},
		{"Format 3", "03291A1ECE1EFC18F94202CA0B53"},
	}

	for _, ex := range examples {
		fmt.Printf("\nDecoding %s:\n", ex.name)

		raw, err := hex.DecodeString(ex.hex)
		if err != nil {
			log.Fatal(err)
		}

		decoded, err := tag.Decode(raw)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("  Detected Format: %d\n", decoded.Format)

		switch decoded.Format {
		case tag.Format5:
			if decoded.Format5.Temperature != nil {
				fmt.Printf("  Temperature:     %.2f°C\n", *decoded.Format5.Temperature)
			}
		case tag.Format3:
			if decoded.Format3.Temperature != nil {
				fmt.Printf("  Temperature:     %.2f°C\n", *decoded.Format3.Temperature)
			}
		}
	}
}
