package mcservermanager

import (
	"fmt"
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
	// proc     *os.Process
	cmd      *exec.Cmd
	newLogs  chan string
	logs     string
	Rcon     *rcon.RconClient
	IsOnline bool
}

func New(s *ServerConfig) *MinecraftServer {
	rcon := rcon.New(s.Address, s.Name)
	return &MinecraftServer{
		Config: s,
		Rcon:   rcon,
	}
}

func (m *MinecraftServer) Start() error {
	cmd := exec.Command(m.Config.RunCommand, m.Config.Args...)
	ch := make(chan string, 1)
	m.newLogs = ch
	m.Rcon.Dial()
	logReader := &StringReader{
		OutputBroadcast: ch,
		logs:            &m.logs,
	}
	cmd.Stdout = logReader
	cmd.Stderr = logReader
	err := cmd.Start()
	if err != nil {
		return err
	}
	fmt.Println(cmd.Process.Pid)
	m.cmd = cmd

	m.IsOnline = true
	return nil
}
func (m *MinecraftServer) ReadLogs() string {
	return m.logs
}

/* BLOCKING! use only in goroutines*/
func (m *MinecraftServer) GetUpdates() string {
	return <-m.newLogs
}

func (m *MinecraftServer) Stop() error {
	m.IsOnline = false
	fmt.Println(m.cmd.Process.Pid)
	err := m.cmd.Process.Kill()
	if err != nil {
		fmt.Println("sigerr", err)
		return err
	}
	// err = exec.Command("kill", "-KILL", fmt.Sprint(m.cmd.Process.Pid)).Run()
	// if err != nil {
	// 	fmt.Println("cannot kill procc")
	// 	return err
	// }
	return nil
}

func (sr *StringReader) writeS(s string) {
	*sr.logs += s
	sr.OutputBroadcast <- s
}

type StringReader struct {
	OutputBroadcast chan string
	logs            *string
}

func (sr *StringReader) Write(p []byte) (n int, err error) {
	s := string(p)
	go sr.writeS(s)
	fmt.Println(s)
	return len(p), nil
}
