package mcservermanager

import (
	"mcap/internal/rcon"
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
	Config *ServerConfig
	proc   *os.Process
	logs   chan string
	Rcon   *rcon.RconClient
}

func New(s *ServerConfig) *MinecraftServer {
	rcon := rcon.New(s.Address, s.Name)
	return &MinecraftServer{
		Config: s,
		Rcon:   rcon,
	}
}

func (m *MinecraftServer) Start(args ...string) error {
	cmd := exec.Command(m.Config.RunCommand, args...)
	m.proc = cmd.Process
	ch := make(chan string)
	m.logs = ch
	m.Rcon.Dial()
	logReader := &StringReader{
		OutputBroadcast: ch,
	}
	cmd.Stdout = logReader
	cmd.Stderr = logReader
	return cmd.Start()
}

/* BLOCKING! use only in goroutines*/
func (m *MinecraftServer) ReadLogs() string {
	return <-m.logs
}

func (m *MinecraftServer) Stop() error {
	return m.proc.Signal(signal)
}

func writeS(s string, c chan string) {
	c <- s
}

type StringReader struct {
	OutputBroadcast chan string
}

func (sr *StringReader) Write(p []byte) (n int, err error) {
	s := string(p)
	go writeS(s, sr.OutputBroadcast)
	return len(p), nil
}
