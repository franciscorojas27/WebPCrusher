package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestBuildAndHelp(t *testing.T) {
	binName := "webp-crusher-testbin"
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}
	bin := filepath.Join(os.TempDir(), binName)

	cmd := exec.Command("go", "build", "-o", bin)
	cmd.Dir = ".."
	cmd.Env = os.Environ()
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("go build failed: %v\n%s", err, string(out))
	}
	defer os.Remove(bin)

	cmd = exec.Command(bin, "-h")
	out, _ := cmd.CombinedOutput()
	outStr := string(out)
	if strings.TrimSpace(outStr) == "" {
		t.Fatalf("expected help output, got empty")
	}
	if !strings.Contains(outStr, "Path to input images directory") {
		t.Fatalf("help output did not contain expected text: %s", outStr)
	}
}
