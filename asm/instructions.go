package asm

import (
	"github.com/SnareChops/6502-tools/asm/lang"
)

var (
	// LDA parses an LDA instruction
	LDA = inst(lang.LDA, ZP, ZPX, AX, AY, A, I, ZPIX, ZPIY)
	// LDX parses an LDX instruction
	LDX = inst(lang.LDX, ZP, ZPY, AY, A, I)
	// LDY parses an LDY instruction
	LDY = inst(lang.LDY, ZP, ZPX, AX, A, I)
	// STA parses an STA instruction
	STA = inst(lang.STA, ZP, ZPX, AX, AY, A, ZPIX, ZPIY)
	// STX parses an STX instruction
	STX = inst(lang.STX, ZP, ZPY, A)
	// STY parses an STY instruction
	STY = inst(lang.STY, ZP, ZPX, A)
	// ADC parses an ADC instruction
	ADC = inst(lang.ADC, ZP, ZPX, ZPIX, ZPIY, AX, AY, A, I)
	// SBC parses an SBC instruction
	SBC = inst(lang.SBC, ZP, ZPX, ZPIX, ZPIY, AX, AY, A, I)
	// INC parses an INC instruction
	INC = inst(lang.INC, ZP, ZPX, AX, A)
	// INX parses an INX instruction
	INX = implied(lang.INX)
	// INY parses an INY instruction
	INY = implied(lang.INY)
	// DEC parses a DEC instruction
	DEC = inst(lang.DEC, ZP, ZPX, AX, A)
	// DEX parses a DEX instructiom
	DEX = implied(lang.DEX)
	// DEY parses a DEY instruction
	DEY = implied(lang.DEY)
	// ASL parses an ASL instruction
	ASL = instOrImp(lang.ASL, ZP, ZPX, AX, A)
	// LSR parses an LSR instruction
	LSR = instOrImp(lang.LSR, ZP, ZPX, AX, A)
	// ROL parses an ROL instruction
	ROL = instOrImp(lang.ROL, ZP, ZPX, AX, A)
	// ROR parses an ROR instruction
	ROR = instOrImp(lang.ROR, ZP, ZPX, AX, A)
	// AND parses an AND instruction
	AND = inst(lang.AND, ZP, ZPX, ZPIX, ZPIY, AX, AY, A, I)
	// ORA parses an ORA instruction
	ORA = inst(lang.ORA, ZP, ZPX, ZPIX, ZPIY, AX, AY, A, I)
	// XOR parses an XOR instruction
	XOR = inst(lang.XOR, ZP, ZPX, ZPIX, ZPIY, AX, AY, A, I)
	// CMP parses a CMP instruction
	CMP = inst(lang.CMP, ZP, ZPX, ZPIX, ZPIY, AX, AY, A, I)
	// CPX parses a CPX instruction
	CPX = inst(lang.CPX, ZP, A, I)
	// CPY parses a CPY instruction
	CPY = inst(lang.CPY, ZP, A, I)
	// BIT parses a BIT instruction
	BIT = inst(lang.BIT, ZP, A, I)
	// BCC parses a BCC instruction
	BCC = inst(lang.BCC, R)
	// BCS parses a BCS instruction
	BCS = inst(lang.BCS, R)
	// BNE parses a BNE instruction
	BNE = inst(lang.BNE, R)
	// BEQ parses a BEQ instruction
	BEQ = inst(lang.BEQ, R)
	// BPL parses a BPL instruction
	BPL = inst(lang.BPL, R)
	// BMI parses a BMI instruction
	BMI = inst(lang.BMI, R)
	// BVC parses a BVC instruction
	BVC = inst(lang.BVC, R)
	// BVS parses a BVS instruction
	BVS = inst(lang.BVS, R)
	// TAX parses a TAX instruction
	TAX = implied(lang.TAX)
	// TXA parses a TXA instruction
	TXA = implied(lang.TXA)
	// TAY parses a TAY instruction
	TAY = implied(lang.TAY)
	// TYA parses a TYA instruction
	TYA = implied(lang.TYA)
	// TSX parses a TSX instruction
	TSX = implied(lang.TSX)
	// TXS parses a TXS instruction
	TXS = implied(lang.TXS)
	// PHA parses a PHA instruction
	PHA = implied(lang.PHA)
	// PLA parses a PLA instruction
	PLA = implied(lang.PLA)
	// PHP parses a PHP instruction
	PHP = implied(lang.PHP)
	// PLP parses a PLP instruction
	PLP = implied(lang.PLP)
	// JMP parses a JMP instruction
	JMP = inst(lang.JMP, AI, A)
	// JSR parses a JSR instruction
	JSR = inst(lang.JSR, A)
	// RTS parses an RTS instruction
	RTS = implied(lang.RTS)
	// RTI parses an RTI instruction
	RTI = implied(lang.RTI)
	// CLC parses a CLC instruction
	CLC = implied(lang.CLC)
	// SEC parses an SEC instruction
	SEC = implied(lang.SEC)
	// CLD parses a CLD instruction
	CLD = implied(lang.CLD)
	// SED parses an SED instruction
	SED = implied(lang.SED)
	// CLI parses a CLI instruction
	CLI = implied(lang.CLI)
	// SEI parses an SEI instruction
	SEI = implied(lang.SEI)
	// CLV parses a CLV instruction
	CLV = implied(lang.CLV)
	// BRK parses a BRK instruction
	BRK = implied(lang.BRK)
	// NOP parses a NOP instruction
	NOP = implied(lang.NOP)
	// Instructions List of all instruction parsers
	Instructions = []InstructionParser{LDA, LDX, LDY, STA, STX, STY, ADC, SBC, INC, INX, INY, DEC, DEX, DEY, ASL, LSR, ROL, ROR, AND, ORA, XOR, CMP, CPX, CPY, BIT, BCC, BCS, BNE, BEQ, BPL, BMI, BVC, BVS, TAX, TXA, TAY, TYA, TSX, TXS, PHA, PLA, PHP, PLP, JMP, JSR, RTS, RTI, CLC, SEC, CLD, SED, CLI, SEI, CLV, BRK, NOP}
)

type InstructionParser = func(string) *Instruction

func inst(acronym string, parsers ...ModeParser) InstructionParser {
	matcher := Matcher("(?i)^" + acronym + "\\s+([\\w$,()]*)(?:\\s*//.*)*$")
	return func(inp string) *Instruction {
		if match := matcher(inp); match != nil {
			match[1] = ReplaceFromLabel(match[1])
			if mode, value := EitherMode(match[1], parsers...); mode != "" {
				instruction := &Instruction{
					Name:   acronym,
					Opcode: lang.Opcode(acronym, mode),
					Value:  value,
				}
				if mode == lang.Label {
					instruction.LabelValue = true
				}
				return instruction
			}
		}
		return nil
	}
}

func implied(acronym string) InstructionParser {
	matcher := Matcher("(?i)^" + acronym + "(?:\\s*//.*)*$")
	return func(inp string) *Instruction {
		if match := matcher(inp); match != nil {
			return &Instruction{
				Name:       acronym,
				Opcode:     lang.Opcode(acronym, lang.IMP),
				LabelValue: false,
				Value:      []byte{},
			}
		}
		return nil
	}
}

func instOrImp(acronym string, parsers ...ModeParser) InstructionParser {
	in := inst(acronym, parsers...)
	imp := implied(acronym)
	return func(inp string) *Instruction {
		return EitherInstruction(inp, imp, in)
	}
}
