package audio

type AudioProcess struct {
	Name string
	Pid  uint32 // It is not a real Pid on linux
}