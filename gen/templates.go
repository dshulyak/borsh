package gen

const (
	marshalStart = `func (t *{{ .Name }}) MarshalBorsh(w io.Writer) error {
	`
	marshalBool = `if err := borsh.WriteBool(w, t.{{ .Name }}); err != nil {
		return err
	}
	`
	marshalUint8 = `if err := borsh.WriteUint8(w, t.{{ .Name }}); err != nil {
		return err
	}
	`
	marshalUint16 = `if err := borsh.WriteUint16(w, t.{{ .Name }}); err != nil {
		return err
	}
	`
	marshalUint32 = `if err := borsh.WriteUint32(w, t.{{ .Name }}); err != nil {
		return err
	}
	`
	marshalUint64 = `if err := borsh.WriteUint64(w, t.{{ .Name }}); err != nil {
		return er r
	}
	`
	marshalLength = `if err := borsh.WriteLength(w, len(t.{{ .Name }})); err != nil {
		return err
	}
	`
	marshalBytes = `if _, err := borsh.WriteBytes(t.{{ .Name }}[:]); err != nil {
		return err
	}
	`
	marshalString = `if _, err := borsh.WriteString(t.{{ .Name }}); err != nil {
		return err
	}
	`
	marshalPtr = `if t.{{ .Name }} == nil {
		if err := borsh.WriteBool(false); err != nil {
			return err
		}
	} else {
		if err := borsh.WriteBool(true); err != nil {
			return err
		}
		if err := t.{{ .Name }}.MarshalBorsh(w); err != nil {
			return err
		}
	}
	`
	marshalStruct = `if err := t.{{ .Name }}.MarshalBorsh(w); err != nil {
		return err
	}
	`
	marshalLoop = `
	for {{ .Index }} := range t.{{ .Name }} {
	`

	unmarshalStart = `func (t *{{ .Name }}) UnmarshalBorsh(r io.Reader) error {
	`
)
