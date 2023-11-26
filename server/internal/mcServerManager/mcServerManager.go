package mcservermanager

import (
	"os"
	"os/exec"
)

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
	// m.proc.Release()
	err := m.proc.Signal(os.Interrupt)
	if err != nil {
		return err
	}
	return nil
}
