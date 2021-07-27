package spec

// Detect the host architecture
func Detect() Microarchitecture {
	return CpuArches["x86"]
}
