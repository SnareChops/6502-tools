package sim

// Registers represents the
// registers of the 6502
// processor
type Registers struct {
	Address [2]byte
	A       byte
	X       byte
	Y       byte
}

// NewRegisters returns a new Registers struct
func NewRegisters() *Registers {
	return &Registers{}
}
