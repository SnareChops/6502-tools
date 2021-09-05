package asm

type Program struct {
	Labels       map[string]uint16
	Instructions []Instruction
}

type Instruction struct {
	Name       string
	Opcode     byte
	LabelValue bool
	Value      []byte
}

func (i *Instruction) Bytes() []byte {
	output := []byte{i.Opcode}
	if i.Value != nil {
		output = append(output, i.Value...)
	}
	return output
}

func (i *Instruction) Size() uint16 {
	if i.Value != nil {
		return uint16(len(i.Value) + 1)
	}
	return 1
}

func (p *Program) Size() uint16 {
	size := uint16(0)
	for _, i := range p.Instructions {
		size += i.Size()
	}
	return size
}

func (p *Program) Compile() []byte {
	output := []byte{}
	for _, i := range p.Instructions {
		output = append(output, i.Bytes()...)
	}
	return output
}

func (p *Program) HasLabel(name string) bool {
	if _, ok := p.Labels[name]; ok {
		return true
	}
	return false
}
