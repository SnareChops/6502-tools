package old_asm_test

import (
	"testing"

	asm "github.com/SnareChops/6502-tools/old_asm"
	"github.com/SnareChops/6502-tools/old_asm/lang"
	"github.com/stretchr/testify/require"
)

func basic(name string, opcode byte, label bool, value ...byte) asm.Instruction {
	if value == nil {
		value = []byte{}
	}
	return asm.Instruction{
		Name:       name,
		Opcode:     opcode,
		LabelValue: label,
		Value:      value,
	}
}

func TestLDA(t *testing.T) {
	result := asm.LDA("LDA $12")
	require.Equal(t, basic(lang.LDA, 0xa5, false, 0x0c), *result)

	result = asm.LDA("lda $12")
	require.Equal(t, basic(lang.LDA, 0xa5, false, 0x0c), *result)

	result = asm.LDA("LDA $12,x")
	require.Equal(t, basic(lang.LDA, 0xb5, false, 0x0c), *result)

	result = asm.LDA("LDA ($12,x)")
	require.Equal(t, basic(lang.LDA, 0xa1, false, 0x0c), *result)

	result = asm.LDA("LDA ($12),y")
	require.Equal(t, basic(lang.LDA, 0xb1, false, 0x0c), *result)

	result = asm.LDA("LDA $0xff12")
	require.Equal(t, basic(lang.LDA, 0xad, false, 0x12, 0xff), *result)

	result = asm.LDA("LDA $0xff12,x")
	require.Equal(t, basic(lang.LDA, 0xbd, false, 0x12, 0xff), *result)

	result = asm.LDA("LDA $0xff12,y")
	require.Equal(t, basic(lang.LDA, 0xb9, false, 0x12, 0xff), *result)

	result = asm.LDA("LDA 0x12")
	require.Equal(t, basic(lang.LDA, 0xa9, false, 0x12), *result)
}

func TestLDX(t *testing.T) {
	result := asm.LDX("LDX $0xff12")
	require.Equal(t, basic(lang.LDX, 0xae, false, 0x12, 0xff), *result)

	result = asm.LDX("LDX $0xff12,y")
	require.Equal(t, basic(lang.LDX, 0xbe, false, 0x12, 0xff), *result)

	result = asm.LDX("LDX 0x12")
	require.Equal(t, basic(lang.LDX, 0xa2, false, 0x12), *result)

	result = asm.LDX("LDX $12")
	require.Equal(t, basic(lang.LDX, 0xa6, false, 0x0c), *result)

	result = asm.LDX("LDX $12,y")
	require.Equal(t, basic(lang.LDX, 0xb6, false, 0x0c), *result)
}

func TestLDY(t *testing.T) {
	result := asm.LDY("LDY $0xff12")
	require.Equal(t, basic(lang.LDY, 0xac, false, 0x12, 0xff), *result)

	result = asm.LDY("LDY $0xff12,x")
	require.Equal(t, basic(lang.LDY, 0xbc, false, 0x12, 0xff), *result)

	result = asm.LDY("LDY 0x12")
	require.Equal(t, basic(lang.LDY, 0xa0, false, 0x12), *result)

	result = asm.LDY("LDY $12")
	require.Equal(t, basic(lang.LDY, 0xa4, false, 0x0c), *result)

	result = asm.LDY("LDY $12,x")
	require.Equal(t, basic(lang.LDY, 0xb4, false, 0x0c), *result)
}

func TestSTA(t *testing.T) {
	result := asm.STA("STA $0xff12")
	require.Equal(t, basic(lang.STA, 0x8d, false, 0x12, 0xff), *result)

	result = asm.STA("STA $0xff12,x")
	require.Equal(t, basic(lang.STA, 0x9d, false, 0x12, 0xff), *result)

	result = asm.STA("STA $0xff12,y")
	require.Equal(t, basic(lang.STA, 0x99, false, 0x12, 0xff), *result)

	result = asm.STA("STA $12")
	require.Equal(t, basic(lang.STA, 0x85, false, 0x0c), *result)

	result = asm.STA("STA ($12,x)")
	require.Equal(t, basic(lang.STA, 0x81, false, 0x0c), *result)

	result = asm.STA("STA $12,x")
	require.Equal(t, basic(lang.STA, 0x95, false, 0x0c), *result)

	result = asm.STA("STA ($12),y")
	require.Equal(t, basic(lang.STA, 0x91, false, 0x0c), *result)
}

func TestSTX(t *testing.T) {
	result := asm.STX("STX $0xff12")
	require.Equal(t, basic(lang.STX, 0x8e, false, 0x12, 0xff), *result)

	result = asm.STX("STX $12")
	require.Equal(t, basic(lang.STX, 0x86, false, 0x0c), *result)

	result = asm.STX("STX $12,y")
	require.Equal(t, basic(lang.STX, 0x96, false, 0x0c), *result)
}

func TestSTY(t *testing.T) {
	result := asm.STY("STY $0xff12")
	require.Equal(t, basic(lang.STY, 0x8c, false, 0x12, 0xff), *result)

	result = asm.STY("STY $12")
	require.Equal(t, basic(lang.STY, 0x84, false, 0x0c), *result)

	result = asm.STY("STY $12,x")
	require.Equal(t, basic(lang.STY, 0x94, false, 0x0c), *result)
}

func TestADC(t *testing.T) {
	result := asm.ADC("ADC $0xff12")
	require.Equal(t, basic(lang.ADC, 0x6d, false, 0x12, 0xff), *result)

	result = asm.ADC("ADC $0xff12,x")
	require.Equal(t, basic(lang.ADC, 0x7d, false, 0x12, 0xff), *result)

	result = asm.ADC("ADC $0xff12,y")
	require.Equal(t, basic(lang.ADC, 0x79, false, 0x12, 0xff), *result)

	result = asm.ADC("ADC 0x12")
	require.Equal(t, basic(lang.ADC, 0x69, false, 0x12), *result)

	result = asm.ADC("ADC $12")
	require.Equal(t, basic(lang.ADC, 0x65, false, 0x0c), *result)

	result = asm.ADC("ADC ($12,x)")
	require.Equal(t, basic(lang.ADC, 0x61, false, 0x0c), *result)

	result = asm.ADC("ADC $12,x")
	require.Equal(t, basic(lang.ADC, 0x75, false, 0x0c), *result)

	result = asm.ADC("ADC ($12),y")
	require.Equal(t, basic(lang.ADC, 0x71, false, 0x0c), *result)
}

func TestSBC(t *testing.T) {
	result := asm.SBC("SBC $0xff12")
	require.Equal(t, basic(lang.SBC, 0xed, false, 0x12, 0xff), *result)

	result = asm.SBC("SBC $0xff12,x")
	require.Equal(t, basic(lang.SBC, 0xfd, false, 0x12, 0xff), *result)

	result = asm.SBC("SBC $0xff12,y")
	require.Equal(t, basic(lang.SBC, 0xf9, false, 0x12, 0xff), *result)

	result = asm.SBC("SBC 0x12")
	require.Equal(t, basic(lang.SBC, 0xe9, false, 0x12), *result)

	result = asm.SBC("SBC $12")
	require.Equal(t, basic(lang.SBC, 0xe5, false, 0x0c), *result)

	result = asm.SBC("SBC ($12,x)")
	require.Equal(t, basic(lang.SBC, 0xe1, false, 0x0c), *result)

	result = asm.SBC("SBC $12,x")
	require.Equal(t, basic(lang.SBC, 0xf5, false, 0x0c), *result)

	result = asm.SBC("SBC ($12),y")
	require.Equal(t, basic(lang.SBC, 0xf1, false, 0x0c), *result)
}

func TestINC(t *testing.T) {
	result := asm.INC("INC $0xff12")
	require.Equal(t, basic(lang.INC, 0xee, false, 0x12, 0xff), *result)

	result = asm.INC("INC $0xff12,x")
	require.Equal(t, basic(lang.INC, 0xfe, false, 0x12, 0xff), *result)

	result = asm.INC("INC $12")
	require.Equal(t, basic(lang.INC, 0xe6, false, 0x0c), *result)

	result = asm.INC("INC $12,x")
	require.Equal(t, basic(lang.INC, 0xf6, false, 0x0c), *result)
}

func TestINX(t *testing.T) {
	result := asm.INX("INX")
	require.Equal(t, basic(lang.INX, 0xe8, false), *result)
}

func TestINY(t *testing.T) {
	result := asm.INY("INY")
	require.Equal(t, basic(lang.INY, 0xc8, false), *result)
}

func TestDEC(t *testing.T) {
	result := asm.DEC("DEC $0xff12")
	require.Equal(t, basic(lang.DEC, 0xce, false, 0x12, 0xff), *result)

	result = asm.DEC("DEC $0xff12,x")
	require.Equal(t, basic(lang.DEC, 0xde, false, 0x12, 0xff), *result)

	result = asm.DEC("DEC $12")
	require.Equal(t, basic(lang.DEC, 0xc6, false, 0x0c), *result)

	result = asm.DEC("DEC $12,x")
	require.Equal(t, basic(lang.DEC, 0xd6, false, 0x0c), *result)
}

func TestDEX(t *testing.T) {
	result := asm.DEX("DEX")
	require.Equal(t, basic(lang.DEX, 0xca, false), *result)
}

func TestDEY(t *testing.T) {
	result := asm.DEY("DEY")
	require.Equal(t, basic(lang.DEY, 0x88, false), *result)
}

func TestASL(t *testing.T) {
	result := asm.ASL("ASL $0xff12")
	require.Equal(t, basic(lang.ASL, 0x0e, false, 0x12, 0xff), *result)

	result = asm.ASL("ASL $0xff12,x")
	require.Equal(t, basic(lang.ASL, 0x1e, false, 0x12, 0xff), *result)

	result = asm.ASL("ASL")
	require.Equal(t, basic(lang.ASL, 0x0a, false), *result)

	result = asm.ASL("ASL $12")
	require.Equal(t, basic(lang.ASL, 0x06, false, 0x0c), *result)

	result = asm.ASL("ASL $12,x")
	require.Equal(t, basic(lang.ASL, 0x16, false, 0x0c), *result)
}

func TestLSR(t *testing.T) {
	result := asm.LSR("LSR $0xff12")
	require.Equal(t, basic(lang.LSR, 0x4e, false, 0x12, 0xff), *result)

	result = asm.LSR("LSR $0xff12,x")
	require.Equal(t, basic(lang.LSR, 0x5e, false, 0x12, 0xff), *result)

	result = asm.LSR("LSR")
	require.Equal(t, basic(lang.LSR, 0x4a, false), *result)

	result = asm.LSR("LSR $12")
	require.Equal(t, basic(lang.LSR, 0x46, false, 0x0c), *result)

	result = asm.LSR("LSR $12,x")
	require.Equal(t, basic(lang.LSR, 0x56, false, 0x0c), *result)
}

func TestROL(t *testing.T) {
	result := asm.ROL("ROL $0xff12")
	require.Equal(t, basic(lang.ROL, 0x2e, false, 0x12, 0xff), *result)

	result = asm.ROL("ROL $0xff12,x")
	require.Equal(t, basic(lang.ROL, 0x3e, false, 0x12, 0xff), *result)

	result = asm.ROL("ROL")
	require.Equal(t, basic(lang.ROL, 0x2a, false), *result)

	result = asm.ROL("ROL $12")
	require.Equal(t, basic(lang.ROL, 0x26, false, 0x0c), *result)

	result = asm.ROL("ROL $12,x")
	require.Equal(t, basic(lang.ROL, 0x36, false, 0x0c), *result)
}

func TestROR(t *testing.T) {
	result := asm.ROR("ROR $0xff12")
	require.Equal(t, basic(lang.ROR, 0x6e, false, 0x12, 0xff), *result)

	result = asm.ROR("ROR $0xff12,x")
	require.Equal(t, basic(lang.ROR, 0x7e, false, 0x12, 0xff), *result)

	result = asm.ROR("ROR")
	require.Equal(t, basic(lang.ROR, 0x6a, false), *result)

	result = asm.ROR("ROR $12")
	require.Equal(t, basic(lang.ROR, 0x66, false, 0x0c), *result)

	result = asm.ROR("ROR $12,x")
	require.Equal(t, basic(lang.ROR, 0x76, false, 0x0c), *result)
}

func TestAND(t *testing.T) {
	result := asm.AND("AND $0xff12")
	require.Equal(t, basic(lang.AND, 0x2d, false, 0x12, 0xff), *result)

	result = asm.AND("AND $0xff12,x")
	require.Equal(t, basic(lang.AND, 0x3d, false, 0x12, 0xff), *result)

	result = asm.AND("AND $0xff12,y")
	require.Equal(t, basic(lang.AND, 0x39, false, 0x12, 0xff), *result)

	result = asm.AND("AND 0x12")
	require.Equal(t, basic(lang.AND, 0x29, false, 0x12), *result)

	result = asm.AND("AND $12")
	require.Equal(t, basic(lang.AND, 0x25, false, 0x0c), *result)

	result = asm.AND("AND ($12,x)")
	require.Equal(t, basic(lang.AND, 0x21, false, 0x0c), *result)

	result = asm.AND("AND $12,x")
	require.Equal(t, basic(lang.AND, 0x35, false, 0x0c), *result)

	result = asm.AND("AND ($12),y")
	require.Equal(t, basic(lang.AND, 0x31, false, 0x0c), *result)
}

func TestORA(t *testing.T) {
	result := asm.ORA("ORA $0xff12")
	require.Equal(t, basic(lang.ORA, 0x0d, false, 0x12, 0xff), *result)

	result = asm.ORA("ORA $0xff12,x")
	require.Equal(t, basic(lang.ORA, 0x1d, false, 0x12, 0xff), *result)

	result = asm.ORA("ORA $0xff12,y")
	require.Equal(t, basic(lang.ORA, 0x19, false, 0x12, 0xff), *result)

	result = asm.ORA("ORA 0x12")
	require.Equal(t, basic(lang.ORA, 0x09, false, 0x12), *result)

	result = asm.ORA("ORA $12")
	require.Equal(t, basic(lang.ORA, 0x05, false, 0x0c), *result)

	result = asm.ORA("ORA ($12,x)")
	require.Equal(t, basic(lang.ORA, 0x01, false, 0x0c), *result)

	result = asm.ORA("ORA $12,x")
	require.Equal(t, basic(lang.ORA, 0x15, false, 0x0c), *result)

	result = asm.ORA("ORA ($12),y")
	require.Equal(t, basic(lang.ORA, 0x11, false, 0x0c), *result)
}

func TestXOR(t *testing.T) {
	result := asm.XOR("XOR $0xff12")
	require.Equal(t, basic(lang.XOR, 0x4d, false, 0x12, 0xff), *result)

	result = asm.XOR("XOR $0xff12,x")
	require.Equal(t, basic(lang.XOR, 0x5d, false, 0x12, 0xff), *result)

	result = asm.XOR("XOR $0xff12,y")
	require.Equal(t, basic(lang.XOR, 0x59, false, 0x12, 0xff), *result)

	result = asm.XOR("XOR 0x12")
	require.Equal(t, basic(lang.XOR, 0x49, false, 0x12), *result)

	result = asm.XOR("XOR $12")
	require.Equal(t, basic(lang.XOR, 0x45, false, 0x0c), *result)

	result = asm.XOR("XOR ($12,x)")
	require.Equal(t, basic(lang.XOR, 0x41, false, 0x0c), *result)

	result = asm.XOR("XOR $12,x")
	require.Equal(t, basic(lang.XOR, 0x55, false, 0x0c), *result)

	result = asm.XOR("XOR ($12),y")
	require.Equal(t, basic(lang.XOR, 0x51, false, 0x0c), *result)
}

func TestCMP(t *testing.T) {
	result := asm.CMP("CMP $0xff12")
	require.Equal(t, basic(lang.CMP, 0xcd, false, 0x12, 0xff), *result)

	result = asm.CMP("CMP $0xff12,x")
	require.Equal(t, basic(lang.CMP, 0xdd, false, 0x12, 0xff), *result)

	result = asm.CMP("CMP $0xff12,y")
	require.Equal(t, basic(lang.CMP, 0xd9, false, 0x12, 0xff), *result)

	result = asm.CMP("CMP 0x12")
	require.Equal(t, basic(lang.CMP, 0xc9, false, 0x12), *result)

	result = asm.CMP("CMP $12")
	require.Equal(t, basic(lang.CMP, 0xc5, false, 0x0c), *result)

	result = asm.CMP("CMP ($12,x)")
	require.Equal(t, basic(lang.CMP, 0xc1, false, 0x0c), *result)

	result = asm.CMP("CMP $12,x")
	require.Equal(t, basic(lang.CMP, 0xd5, false, 0x0c), *result)

	result = asm.CMP("CMP ($12),y")
	require.Equal(t, basic(lang.CMP, 0xd1, false, 0x0c), *result)
}

func TestCPX(t *testing.T) {
	result := asm.CPX("CPX $0xff12")
	require.Equal(t, basic(lang.CPX, 0xec, false, 0x12, 0xff), *result)

	result = asm.CPX("CPX 0x12")
	require.Equal(t, basic(lang.CPX, 0xe0, false, 0x12), *result)

	result = asm.CPX("CPX $12")
	require.Equal(t, basic(lang.CPX, 0xe4, false, 0x0c), *result)
}

func TestCPY(t *testing.T) {
	result := asm.CPY("CPY $0xff12")
	require.Equal(t, basic(lang.CPY, 0xcc, false, 0x12, 0xff), *result)

	result = asm.CPY("CPY 0x12")
	require.Equal(t, basic(lang.CPY, 0xc0, false, 0x12), *result)

	result = asm.CPY("CPY $12")
	require.Equal(t, basic(lang.CPY, 0xc4, false, 0x0c), *result)
}

func TestBIT(t *testing.T) {
	result := asm.BIT("BIT $0xff12")
	require.Equal(t, basic(lang.BIT, 0x2c, false, 0x12, 0xff), *result)

	result = asm.BIT("BIT 0x12")
	require.Equal(t, basic(lang.BIT, 0x89, false, 0x12), *result)

	result = asm.BIT("BIT $12")
	require.Equal(t, basic(lang.BIT, 0x24, false, 0x0c), *result)
}

func TestBCC(t *testing.T) {
	result := asm.BCC("BCC $12")
	require.Equal(t, basic(lang.BCC, 0x90, false, 0x0c), *result)
}

func TestBCS(t *testing.T) {
	result := asm.BCS("BCS $12")
	require.Equal(t, basic(lang.BCS, 0xb0, false, 0x0c), *result)
}

func TestBNE(t *testing.T) {
	result := asm.BNE("BNE $12")
	require.Equal(t, basic(lang.BNE, 0xd0, false, 0x0c), *result)
}

func TestBEQ(t *testing.T) {
	result := asm.BEQ("BEQ $12")
	require.Equal(t, basic(lang.BEQ, 0xf0, false, 0x0c), *result)
}

func TestBPL(t *testing.T) {
	result := asm.BPL("BPL $12")
	require.Equal(t, basic(lang.BPL, 0x10, false, 0x0c), *result)
}

func TestBMI(t *testing.T) {
	result := asm.BMI("BMI $12")
	require.Equal(t, basic(lang.BMI, 0x30, false, 0x0c), *result)
}

func TestBVC(t *testing.T) {
	result := asm.BVC("BVC $12")
	require.Equal(t, basic(lang.BVC, 0x50, false, 0x0c), *result)
}

func TestBVS(t *testing.T) {
	result := asm.BVS("BVS $12")
	require.Equal(t, basic(lang.BVS, 0x70, false, 0x0c), *result)
}

func TestTAX(t *testing.T) {
	result := asm.TAX("TAX")
	require.Equal(t, basic(lang.TAX, 0xaa, false), *result)
}

func TestTXA(t *testing.T) {
	result := asm.TXA("TXA")
	require.Equal(t, basic(lang.TXA, 0x8a, false), *result)
}

func TestTAY(t *testing.T) {
	result := asm.TAY("TAY")
	require.Equal(t, basic(lang.TAY, 0xa8, false), *result)
}

func TestTYA(t *testing.T) {
	result := asm.TYA("TYA")
	require.Equal(t, basic(lang.TYA, 0x98, false), *result)
}

func TestTSX(t *testing.T) {
	result := asm.TSX("TSX")
	require.Equal(t, basic(lang.TSX, 0xba, false), *result)
}

func TestTXS(t *testing.T) {
	result := asm.TXS("TXS")
	require.Equal(t, basic(lang.TXS, 0x9a, false), *result)
}

func TestPHA(t *testing.T) {
	result := asm.PHA("PHA")
	require.Equal(t, basic(lang.PHA, 0x48, false), *result)
}

func TestPLA(t *testing.T) {
	result := asm.PLA("PLA")
	require.Equal(t, basic(lang.PLA, 0x68, false), *result)
}

func TestPHP(t *testing.T) {
	result := asm.PHP("PHP")
	require.Equal(t, basic(lang.PHP, 0x08, false), *result)
}

func TestPLP(t *testing.T) {
	result := asm.PLP("PLP")
	require.Equal(t, basic(lang.PLP, 0x28, false), *result)
}

func TestJMP(t *testing.T) {
	result := asm.JMP("JMP $0xff12")
	require.Equal(t, basic(lang.JMP, 0x4c, false, 0x12, 0xff), *result)

	result = asm.JMP("JMP ($0xff12)")
	require.Equal(t, basic(lang.JMP, 0x6c, false, 0x12, 0xff), *result)

	// result = asm.JMP("JMP label")
	// require.Equal(t, basic(lang.JMP, 0x4c, true, []byte("label")...), *result)
}

func TestJSR(t *testing.T) {
	result := asm.JSR("JSR $0xff12")
	require.Equal(t, basic(lang.JSR, 0x20, false, 0x12, 0xff), *result)
}

func TestRTS(t *testing.T) {
	result := asm.RTS("RTS")
	require.Equal(t, basic(lang.RTS, 0x60, false), *result)
}

func TestRTI(t *testing.T) {
	result := asm.RTI("RTI")
	require.Equal(t, basic(lang.RTI, 0x40, false), *result)
}

func TestCLC(t *testing.T) {
	result := asm.CLC("CLC")
	require.Equal(t, basic(lang.CLC, 0x18, false), *result)
}

func TestSEC(t *testing.T) {
	result := asm.SEC("SEC")
	require.Equal(t, basic(lang.SEC, 0x38, false), *result)
}

func TestCLD(t *testing.T) {
	result := asm.CLD("CLD")
	require.Equal(t, basic(lang.CLD, 0xd8, false), *result)
}

func TestSED(t *testing.T) {
	result := asm.SED("SED")
	require.Equal(t, basic(lang.SED, 0xf8, false), *result)
}

func TestCLI(t *testing.T) {
	result := asm.CLI("CLI")
	require.Equal(t, basic(lang.CLI, 0x58, false), *result)
}

func TestSEI(t *testing.T) {
	result := asm.SEI("SEI")
	require.Equal(t, basic(lang.SEI, 0x78, false), *result)
}

func TestCLV(t *testing.T) {
	result := asm.CLV("CLV")
	require.Equal(t, basic(lang.CLV, 0xb8, false), *result)
}

func TestBRK(t *testing.T) {
	result := asm.BRK("BRK")
	require.Equal(t, basic(lang.BRK, 0x00, false), *result)
}

func TestNOP(t *testing.T) {
	result := asm.NOP("NOP")
	require.Equal(t, basic(lang.NOP, 0xea, false), *result)
}
