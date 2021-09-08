// Code generated by github.com/spacemeshos/borsh/gen. DO NOT EDIT.

package lib

import (
	"io"

	"github.com/spacemeshos/borsh"
)

func (t *Struct) MarshalBorsh(w io.Writer) error {
	// field Bytes (0)
	if t.Bytes == nil {
		if err := borsh.WriteBool(w, false); err != nil {
			return err
		}
	} else {
		if err := borsh.WriteBool(w, true); err != nil {
			return err
		}
		if err := borsh.WriteLength(w, len(t.Bytes)); err != nil {
			return err
		}
	}
	if err := borsh.WriteBytes(w, t.Bytes[:]); err != nil {
		return err
	}

	// field Bool (1)
	if err := borsh.WriteBool(w, t.Bool); err != nil {
		return err
	}

	// field Ptr (2)
	if t.Ptr == nil {
		if err := borsh.WriteBool(w, false); err != nil {
			return err
		}
	} else {
		if err := borsh.WriteBool(w, true); err != nil {
			return err
		}
		if err := t.Ptr.MarshalBorsh(w); err != nil {
			return err
		}
	}

	return nil
}

func (t *Struct) UnmarshalBorsh(r io.Reader) error {
	// field Bytes (0)
	if exist, err := borsh.ReadBool(r); err != nil {
		return err
	} else if exist {
		if lth, err := borsh.ReadUint32(r); err != nil {
			return err
		} else {
			t.Bytes = make([]uint8, lth)
		}
	}
	if err := borsh.ReadBytes(r, t.Bytes[:]); err != nil {
		return err
	}

	// field Bool (1)
	if val, err := borsh.ReadBool(r); err != nil {
		return err
	} else {
		t.Bool = val
	}

	// field Ptr (2)
	if exist, err := borsh.ReadBool(r); err != nil {
		return err
	} else if exist {
		if t.Ptr == nil {
			t.Ptr = new(Struct)
		}
		if err := t.Ptr.UnmarshalBorsh(r); err != nil {
			return err
		}
	}

	return nil
}
