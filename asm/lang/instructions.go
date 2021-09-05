package lang

const (
	// LDA Load Accumulator with Memory
	LDA = `LDA`
	// LDX Load Index X with Memory
	LDX = `LDX`
	// LDY Load Index Y with Memory
	LDY = `LDY`
	// STA Store Accumulator in Memory
	STA = `STA`
	// STX Store Index X in Memory
	STX = `STX`
	// STY Store Index Y in Memory
	STY = `STY`
	// ADC Add Memory to Accumulator with Carry
	ADC = `ADC`
	// SBC Subtract Memory from Accumulator with Borrow
	SBC = `SBC`
	// INC Increment Memory by One
	INC = `INC`
	// INX Increment Index X by One
	INX = `INX`
	// INY Increment Index Y by One
	INY = `INY`
	// DEC Decrement Memory by One
	DEC = `DEC`
	// DEX Decrement Index X by One
	DEX = `DEX`
	// DEY Decrement Index Y by One
	DEY = `DEY`
	// ASL Arithmetic Shift Left One Bit
	ASL = `ASL`
	// LSR Logical Shift Right One Bit
	LSR = `LSR`
	// ROL Rotate Left One Bit
	ROL = `ROL`
	// ROR Rotate Right One Bit
	ROR = `ROR`
	// AND AND Memory with Accumulator
	AND = `AND`
	// ORA OR Memory with Accumulator
	ORA = `ORA`
	// XOR XOR Memory with Accumulator
	XOR = `XOR`
	// CMP Compare Memory and Accumulator
	CMP = `CMP`
	// CPX Compare Memory and Index X
	CPX = `CPX`
	// CPY Compare Memory with Index Y
	CPY = `CPY`
	// BIT Test Bits in Memory with Accumulator
	BIT = `BIT`
	// BCC Branch on Carry Clear
	BCC = `BCC`
	// BCS Branch on Carry Set
	BCS = `BCS`
	// BNE Branch on Result not Zero
	BNE = `BNE`
	// BEQ Branch on Result Zero
	BEQ = `BEQ`
	// BPL Branch on Result Plus
	BPL = `BPL`
	// BMI Branch on Result Minus
	BMI = `BMI`
	// BVC Branch on Overflow Clear
	BVC = `BVC`
	// BVS Branch on Overflow Set
	BVS = `BVS`
	// TAX Transfer Accumulator to Index X
	TAX = `TAX`
	// TXA Transfer Index X to Accumulator
	TXA = `TXA`
	// TAY Transfer Accumulator to Index Y
	TAY = `TAY`
	// TYA Transfer Index Y to Accumulator
	TYA = `TYA`
	// TSX Transfer Stack Pointer to Index X
	TSX = `TSX`
	// TXS Transfer Index X to Stack Pointer
	TXS = `TXS`
	// PHA Push Accumulator on Stack
	PHA = `PHA`
	// PLA Pull Accumulator from Stack
	PLA = `PLA`
	// PHP Push Processor Status on Stack
	PHP = `PHP`
	// PLP Pull Processor Status from Stack
	PLP = `PLP`
	// JMP Jump to New Location
	JMP = `JMP`
	// JSR Jump to New Location Saving Return Address
	JSR = `JSR`
	// RTS Return from Subroutine
	RTS = `RTS`
	// RTI Return from Interrupt
	RTI = `RTI`
	// CLC Clear Carry Flag
	CLC = `CLC`
	// SEC Set Carry Flag
	SEC = `SEC`
	// CLD Clear Decimal Mode
	CLD = `CLD`
	// SED Set Decimal Mode
	SED = `SED`
	// CLI Clear Interrupt Disable Status
	CLI = `CLI`
	// SEI Set Interrupt Disable Status
	SEI = `SEI`
	// CLV Clear Overflow Flag
	CLV = `CLV`
	// BRK Break
	BRK = `BRK`
	// NOP No Operation
	NOP = `NOP`
)
