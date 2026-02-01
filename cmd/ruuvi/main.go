// Package main provides a command-line tool for decoding and encoding RuuviTag data.
package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/marcgeld/ruuvi/tag"
)

const (
	exitSuccess = 0
	exitError   = 1
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(exitError)
	}
	os.Exit(exitSuccess)
}

func run() error {
	// Define subcommands
	decodeCmd := flag.NewFlagSet("decode", flag.ExitOnError)
	encodeCmd := flag.NewFlagSet("encode", flag.ExitOnError)

	// Decode flags
	decodeHex := decodeCmd.String("hex", "", "Hex-encoded RuuviTag data to decode (required)")

	// Encode flags
	encodeJSON := encodeCmd.String("json", "", "JSON-encoded Format5Data to encode (required)")

	// Check if a subcommand was provided
	if len(os.Args) < 2 {
		printUsage()
		return fmt.Errorf("no command specified")
	}

	// Parse subcommand
	switch os.Args[1] {
	case "decode":
		if err := decodeCmd.Parse(os.Args[2:]); err != nil {
			return err
		}
		return handleDecode(*decodeHex)

	case "encode":
		if err := encodeCmd.Parse(os.Args[2:]); err != nil {
			return err
		}
		return handleEncode(*encodeJSON)

	default:
		printUsage()
		return fmt.Errorf("unknown command: %s", os.Args[1])
	}
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage: ruuvi <command> [flags]")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Commands:")
	fmt.Fprintln(os.Stderr, "  decode    Decode RuuviTag data from hex to JSON")
	fmt.Fprintln(os.Stderr, "  encode    Encode Format 5 data from JSON to hex")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Decode flags:")
	fmt.Fprintln(os.Stderr, "  --hex string    Hex-encoded RuuviTag data (required)")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Encode flags:")
	fmt.Fprintln(os.Stderr, "  --json string   JSON-encoded Format5Data (required)")
}

func handleDecode(hexStr string) error {
	if hexStr == "" {
		return fmt.Errorf("--hex flag is required")
	}

	// Decode hex string to bytes
	data, err := hex.DecodeString(hexStr)
	if err != nil {
		return fmt.Errorf("invalid hex string: %w", err)
	}

	// Decode RuuviTag data
	decoded, err := tag.Decode(data)
	if err != nil {
		return fmt.Errorf("failed to decode data: %w", err)
	}

	// Convert to JSON and print
	output, err := json.MarshalIndent(decoded, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	fmt.Println(string(output))
	return nil
}

func handleEncode(jsonStr string) error {
	if jsonStr == "" {
		return fmt.Errorf("--json flag is required")
	}

	// Parse JSON input
	var data tag.Format5Data
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Encode to Format 5
	encoded, err := tag.EncodeFormat5(&data)
	if err != nil {
		return fmt.Errorf("failed to encode data: %w", err)
	}

	// Output as hex
	fmt.Println(hex.EncodeToString(encoded))
	return nil
}
