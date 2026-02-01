package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/marcgeld/hermod/tag"
)

// captureStdoutStderr captures stdout and stderr produced by fn and returns them.
func captureStdoutStderr(fn func()) (string, string) {
	origOut := os.Stdout
	origErr := os.Stderr

	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()

	os.Stdout = wOut
	os.Stderr = wErr

	outC := make(chan string)
	errC := make(chan string)

	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, rOut)
		outC <- buf.String()
	}()

	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, rErr)
		errC <- buf.String()
	}()

	fn()

	_ = wOut.Close()
	_ = wErr.Close()

	os.Stdout = origOut
	os.Stderr = origErr

	out := <-outC
	err := <-errC

	return out, err
}

func TestRun_NoArgs_ShowsUsageAndErrors(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"ruuvi"}

	_, stderr := captureStdoutStderr(func() {
		err := run()
		if err == nil || !strings.Contains(err.Error(), "no command specified") {
			t.Fatalf("expected 'no command specified' error, got: %v", err)
		}
	})

	if !strings.Contains(stderr, "Usage: ruuvi") {
		t.Fatalf("expected usage printed to stderr, got: %q", stderr)
	}
}

func TestRun_UnknownCommand_ShowsUsageAndError(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"ruuvi", "unknowncmd"}

	_, stderr := captureStdoutStderr(func() {
		err := run()
		if err == nil || !strings.Contains(err.Error(), "unknown command") {
			t.Fatalf("expected unknown command error, got: %v", err)
		}
	})

	if !strings.Contains(stderr, "Usage: ruuvi") {
		t.Fatalf("expected usage printed to stderr for unknown command, got: %q", stderr)
	}
}

func TestRun_Decode_NoHexFlag_ReturnsError(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"ruuvi", "decode"}

	_, _ = captureStdoutStderr(func() {
		err := run()
		if err == nil || !strings.Contains(err.Error(), "--hex flag is required") {
			t.Fatalf("expected '--hex flag is required' error, got: %v", err)
		}
	})
}

func TestRun_Encode_NoJSONFlag_ReturnsError(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"ruuvi", "encode"}

	_, _ = captureStdoutStderr(func() {
		err := run()
		if err == nil || !strings.Contains(err.Error(), "--json flag is required") {
			t.Fatalf("expected '--json flag is required' error, got: %v", err)
		}
	})
}

func TestHandleDecode_InvalidHex_ReturnsError(t *testing.T) {
	_, _ = captureStdoutStderr(func() {
		err := handleDecode("nothex")
		if err == nil || !strings.Contains(err.Error(), "invalid hex string") {
			t.Fatalf("expected invalid hex string error, got: %v", err)
		}
	})
}

func TestHandleEncode_InvalidJSON_ReturnsError(t *testing.T) {
	_, _ = captureStdoutStderr(func() {
		err := handleEncode("{invalid json")
		if err == nil || !strings.Contains(err.Error(), "failed to parse JSON") {
			t.Fatalf("expected failed to parse JSON error, got: %v", err)
		}
	})
}

func TestHandleDecode_ValidFormat5(t *testing.T) {
	// Create a Format5Data and encode to bytes
	temp := 24.3
	hum := 50.0
	press := 101325
	mac := tag.Format5Data{Temperature: &temp, Humidity: &hum, Pressure: &press}
	b, err := tag.EncodeFormat5(&mac)
	if err != nil {
		t.Fatalf("failed to encode format5 test vector: %v", err)
	}
	hexStr := hex.EncodeToString(b)

	out, _ := captureStdoutStderr(func() {
		err := handleDecode(hexStr)
		if err != nil {
			t.Fatalf("handleDecode returned error: %v", err)
		}
	})

	if !strings.Contains(out, "\"Format\": 5") {
		t.Fatalf("expected decoded output to contain Format 5, got: %s", out)
	}
}

func TestHandleEncode_ValidFormat5(t *testing.T) {
	// Prepare JSON for a simple Format5Data
	temp := 21.0
	press := 100000
	data := tag.Format5Data{Temperature: &temp, Pressure: &press}
	j, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("failed to marshal JSON test data: %v", err)
	}

	expectedBytes, err := tag.EncodeFormat5(&data)
	if err != nil {
		t.Fatalf("failed to encode expected bytes: %v", err)
	}
	expectedHex := hex.EncodeToString(expectedBytes)

	out, _ := captureStdoutStderr(func() {
		err := handleEncode(string(j))
		if err != nil {
			t.Fatalf("handleEncode returned error: %v", err)
		}
	})

	out = strings.TrimSpace(out)
	if out != expectedHex {
		t.Fatalf("handleEncode output = %q; want %q", out, expectedHex)
	}
}

func TestRun_Decode_Success(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Build a valid format5 payload
	temp := 19.5
	data := tag.Format5Data{Temperature: &temp}
	b, err := tag.EncodeFormat5(&data)
	if err != nil {
		t.Fatalf("failed to encode format5: %v", err)
	}
	hexStr := hex.EncodeToString(b)

	os.Args = []string{"ruuvi", "decode", "--hex", hexStr}

	out, stderr := captureStdoutStderr(func() {
		err := run()
		if err != nil {
			t.Fatalf("run() returned error: %v", err)
		}
	})

	if !strings.Contains(out, "\"Format\": 5") {
		t.Fatalf("expected decoded output to contain Format 5, got: %s", out)
	}
	if stderr != "" {
		t.Fatalf("expected no stderr on successful run, got: %s", stderr)
	}
}

func TestRun_Encode_Success(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Prepare JSON for encode
	temp := 22.0
	press := 100500
	f := tag.Format5Data{Temperature: &temp, Pressure: &press}
	j, err := json.Marshal(f)
	if err != nil {
		t.Fatalf("failed to marshal json: %v", err)
	}

	os.Args = []string{"ruuvi", "encode", "--json", string(j)}

	out, stderr := captureStdoutStderr(func() {
		err := run()
		if err != nil {
			t.Fatalf("run() returned error: %v", err)
		}
	})

	out = strings.TrimSpace(out)
	// Verify the output is valid hex of the expected encoded bytes
	expected, err := tag.EncodeFormat5(&f)
	if err != nil {
		t.Fatalf("EncodeFormat5 error: %v", err)
	}
	expectedHex := hex.EncodeToString(expected)
	if out != expectedHex {
		t.Fatalf("run encode output = %q; want %q", out, expectedHex)
	}
	if stderr != "" {
		t.Fatalf("expected no stderr on successful run, got: %s", stderr)
	}
}
