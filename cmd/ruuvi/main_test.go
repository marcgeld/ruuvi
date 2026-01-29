package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
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
