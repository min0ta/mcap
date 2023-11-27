package mcservermanager

import (
	"os"
	"os/exec"
	"runtime"
)

func init() {
	if runtime.GOOS == "windows" {
		// windows fallback to sigkill
		signal = os.Kill
		return
	}
	signal = os.Interrupt
}

var signal os.Signal

type MinecraftServer struct {
	path string
	proc *os.Process
}

func (m *MinecraftServer) Start(args ...string) error {
	cmd := exec.Command("./"+m.path, args...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	m.proc = cmd.Process
	return nil
}

func (m *MinecraftServer) Stop() error {
	err := m.proc.Signal(signal)
	if err != nil {
		return err
	}
	return nil
}
