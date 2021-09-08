package gen

import (
	"io"
	"text/template"
)

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
	marshalSlice = `if t.{{ .Name }} == nil {
		if err := borsh.WriteBool(w, false); err != nil {
			return err
		}
	} else {
		if err := borsh.WriteBool(w, true); err != nil {
			return err
		}
		if err := borsh.WriteLength(w, len(t.{{ .Name }})); err != nil {
			return err
		}
	}
	`
	marshalBytes = `if err := borsh.WriteBytes(w, t.{{ .Name }}[:]); err != nil {
		return err
	}
	`
	marshalString = `if _, err := borsh.WriteString(t.{{ .Name }}); err != nil {
		return err
	}
	`
	marshalPtr = `if t.{{ .Name }} == nil {
		if err := borsh.WriteBool(w, false); err != nil {
			return err
		}
	} else {
		if err := borsh.WriteBool(w, true); err != nil {
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

	addLoop = `
	for {{ .Index }} := range t.{{ .Name }} {
	`

	unmarshalStart = `func (t *{{ .Name }}) UnmarshalBorsh(r io.Reader) error {
	`
	unmarshalBool = `if val, err := borsh.ReadBool(r); err != nil {
		return err
	} else {
		t.{{ .Name }} = val
	}
	`
	unmarshalUint32 = `if val, err := borsh.ReadUint32(r); err != nil {
		return err
	} else {
		t.{{ .Name }} = val
	}
	`
	unmarshalBytes = `if err := borsh.ReadBytes(r, t.{{ .Name }}[:]); err != nil {
		return err
	}
	`
	unmarshalSlice = `if exist, err := borsh.ReadBool(r); err != nil {
		return err
	} else if exist {	
		if lth, err := borsh.ReadUint32(r); err != nil {
			return err
		} else {
			t.{{ .Name }} = make({{ .TypeName }}, lth)
		}
	}
	`
	unmarshalStruct = `if err := t.{{ .Name }}.UnmarshalBorsh(r); err != nil {
		return err
	}
	`
	unmarshalPtr = `if exist, err := borsh.ReadBool(r); err != nil {
		return err
	} else if exist {
		if t.{{ .Name }} == nil {
			t.{{ .Name }} = new({{ .TypeName }})
		}
		if err := t.{{ .Name }}.UnmarshalBorsh(r); err != nil {
			return err
		} 
	}
	`
)

func executeTemplate(w io.Writer, text string, ctx interface{}) error {
	tpl, err := template.New("").Parse(text)
	if err != nil {
		return err
	}
	return tpl.Execute(w, ctx)
}
