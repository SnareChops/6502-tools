package old_asm

type Token interface {
	Bytes() []byte
}

type InstructionToken struct {
}

type CommentToken struct {
}

type LabelToken struct {
}

type ImportToken struct {
}
