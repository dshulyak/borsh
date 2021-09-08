// Code generated by github.com/dshulyak/borsh/gen. DO NOT EDIT.

package tests

import (
	"io"

	"github.com/dshulyak/borsh"
	"github.com/dshulyak/borsh/tests/go-lib"
	"github.com/dshulyak/borsh/tests/imported"
)

func (t *Local) MarshalBorsh(w io.Writer) error {
	// field LocalAlias (0)
	if t.LocalAlias == nil {
		if err := borsh.WriteBool(w, false); err != nil {
			return err
		}
	} else {
		if err := borsh.WriteBool(w, true); err != nil {
			return err
		}
		if err := borsh.WriteLength(w, len(t.LocalAlias)); err != nil {
			return err
		}
	}

	for i := range t.LocalAlias {
		if err := borsh.WriteBytes(w, t.LocalAlias[i][:]); err != nil {
			return err
		}
	}

	// field Bool (1)
	if err := borsh.WriteBool(w, t.Bool); err != nil {
		return err
	}

	// field Uint32Slice (2)
	if t.Uint32Slice == nil {
		if err := borsh.WriteBool(w, false); err != nil {
			return err
		}
	} else {
		if err := borsh.WriteBool(w, true); err != nil {
			return err
		}
		if err := borsh.WriteLength(w, len(t.Uint32Slice)); err != nil {
			return err
		}
	}

	for i := range t.Uint32Slice {
		if err := borsh.WriteUint32(w, t.Uint32Slice[i]); err != nil {
			return err
		}
	}

	// field Uint32Array (3)

	for i := range t.Uint32Array {
		if err := borsh.WriteUint32(w, t.Uint32Array[i]); err != nil {
			return err
		}
	}

	// field PtrToSelf (4)
	if t.PtrToSelf == nil {
		if err := borsh.WriteBool(w, false); err != nil {
			return err
		}
	} else {
		if err := borsh.WriteBool(w, true); err != nil {
			return err
		}
		if err := t.PtrToSelf.MarshalBorsh(w); err != nil {
			return err
		}
	}

	return nil
}

func (t *Local) UnmarshalBorsh(r io.Reader) error {
	// field LocalAlias (0)
	if exist, err := borsh.ReadBool(r); err != nil {
		return err
	} else if exist {
		if lth, err := borsh.ReadUint32(r); err != nil {
			return err
		} else {
			t.LocalAlias = make([]Alias, lth)
		}
	}

	for i := range t.LocalAlias {
		if err := borsh.ReadBytes(r, t.LocalAlias[i][:]); err != nil {
			return err
		}
	}

	// field Bool (1)
	if val, err := borsh.ReadBool(r); err != nil {
		return err
	} else {
		t.Bool = val
	}

	// field Uint32Slice (2)
	if exist, err := borsh.ReadBool(r); err != nil {
		return err
	} else if exist {
		if lth, err := borsh.ReadUint32(r); err != nil {
			return err
		} else {
			t.Uint32Slice = make([]uint32, lth)
		}
	}

	for i := range t.Uint32Slice {
		if val, err := borsh.ReadUint32(r); err != nil {
			return err
		} else {
			t.Uint32Slice[i] = val
		}
	}

	// field Uint32Array (3)

	for i := range t.Uint32Array {
		if val, err := borsh.ReadUint32(r); err != nil {
			return err
		} else {
			t.Uint32Array[i] = val
		}
	}

	// field PtrToSelf (4)
	if exist, err := borsh.ReadBool(r); err != nil {
		return err
	} else if exist {
		if t.PtrToSelf == nil {
			t.PtrToSelf = new(Local)
		}
		if err := t.PtrToSelf.UnmarshalBorsh(r); err != nil {
			return err
		}
	}

	return nil
}

func (t *Hello) MarshalBorsh(w io.Writer) error {
	// field ID (0)
	if err := borsh.WriteBytes(w, t.ID[:]); err != nil {
		return err
	}

	// field Ptr (1)
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

	// field Ptrs (2)
	if t.Ptrs == nil {
		if err := borsh.WriteBool(w, false); err != nil {
			return err
		}
	} else {
		if err := borsh.WriteBool(w, true); err != nil {
			return err
		}
		if err := borsh.WriteLength(w, len(t.Ptrs)); err != nil {
			return err
		}
	}

	for i := range t.Ptrs {
		if t.Ptrs[i] == nil {
			if err := borsh.WriteBool(w, false); err != nil {
				return err
			}
		} else {
			if err := borsh.WriteBool(w, true); err != nil {
				return err
			}
			if err := t.Ptrs[i].MarshalBorsh(w); err != nil {
				return err
			}
		}
	}

	// field NestedPtrs (3)
	if t.NestedPtrs == nil {
		if err := borsh.WriteBool(w, false); err != nil {
			return err
		}
	} else {
		if err := borsh.WriteBool(w, true); err != nil {
			return err
		}
		if err := borsh.WriteLength(w, len(t.NestedPtrs)); err != nil {
			return err
		}
	}

	for i := range t.NestedPtrs {
		if t.NestedPtrs[i] == nil {
			if err := borsh.WriteBool(w, false); err != nil {
				return err
			}
		} else {
			if err := borsh.WriteBool(w, true); err != nil {
				return err
			}
			if err := borsh.WriteLength(w, len(t.NestedPtrs[i])); err != nil {
				return err
			}
		}

		for ii := range t.NestedPtrs[i] {
			if t.NestedPtrs[i][ii] == nil {
				if err := borsh.WriteBool(w, false); err != nil {
					return err
				}
			} else {
				if err := borsh.WriteBool(w, true); err != nil {
					return err
				}
				if err := t.NestedPtrs[i][ii].MarshalBorsh(w); err != nil {
					return err
				}
			}
		}
	}

	// field BytesIDs (4)
	if t.BytesIDs == nil {
		if err := borsh.WriteBool(w, false); err != nil {
			return err
		}
	} else {
		if err := borsh.WriteBool(w, true); err != nil {
			return err
		}
		if err := borsh.WriteLength(w, len(t.BytesIDs)); err != nil {
			return err
		}
	}

	for i := range t.BytesIDs {
		if err := borsh.WriteBytes(w, t.BytesIDs[i][:]); err != nil {
			return err
		}
	}

	// field BytesCollection (5)
	if t.BytesCollection == nil {
		if err := borsh.WriteBool(w, false); err != nil {
			return err
		}
	} else {
		if err := borsh.WriteBool(w, true); err != nil {
			return err
		}
		if err := borsh.WriteLength(w, len(t.BytesCollection)); err != nil {
			return err
		}
	}

	for i := range t.BytesCollection {
		if t.BytesCollection[i] == nil {
			if err := borsh.WriteBool(w, false); err != nil {
				return err
			}
		} else {
			if err := borsh.WriteBool(w, true); err != nil {
				return err
			}
			if err := borsh.WriteLength(w, len(t.BytesCollection[i])); err != nil {
				return err
			}
		}
		if err := borsh.WriteBytes(w, t.BytesCollection[i][:]); err != nil {
			return err
		}
	}

	// field AnotherImport (6)
	if t.AnotherImport == nil {
		if err := borsh.WriteBool(w, false); err != nil {
			return err
		}
	} else {
		if err := borsh.WriteBool(w, true); err != nil {
			return err
		}
		if err := t.AnotherImport.MarshalBorsh(w); err != nil {
			return err
		}
	}

	return nil
}

func (t *Hello) UnmarshalBorsh(r io.Reader) error {
	// field ID (0)
	if err := borsh.ReadBytes(r, t.ID[:]); err != nil {
		return err
	}

	// field Ptr (1)
	if exist, err := borsh.ReadBool(r); err != nil {
		return err
	} else if exist {
		if t.Ptr == nil {
			t.Ptr = new(imported.Struct)
		}
		if err := t.Ptr.UnmarshalBorsh(r); err != nil {
			return err
		}
	}

	// field Ptrs (2)
	if exist, err := borsh.ReadBool(r); err != nil {
		return err
	} else if exist {
		if lth, err := borsh.ReadUint32(r); err != nil {
			return err
		} else {
			t.Ptrs = make([]*imported.Struct, lth)
		}
	}

	for i := range t.Ptrs {
		if exist, err := borsh.ReadBool(r); err != nil {
			return err
		} else if exist {
			if t.Ptrs[i] == nil {
				t.Ptrs[i] = new(imported.Struct)
			}
			if err := t.Ptrs[i].UnmarshalBorsh(r); err != nil {
				return err
			}
		}
	}

	// field NestedPtrs (3)
	if exist, err := borsh.ReadBool(r); err != nil {
		return err
	} else if exist {
		if lth, err := borsh.ReadUint32(r); err != nil {
			return err
		} else {
			t.NestedPtrs = make([][]*imported.Struct, lth)
		}
	}

	for i := range t.NestedPtrs {
		if exist, err := borsh.ReadBool(r); err != nil {
			return err
		} else if exist {
			if lth, err := borsh.ReadUint32(r); err != nil {
				return err
			} else {
				t.NestedPtrs[i] = make([]*imported.Struct, lth)
			}
		}

		for ii := range t.NestedPtrs[i] {
			if exist, err := borsh.ReadBool(r); err != nil {
				return err
			} else if exist {
				if t.NestedPtrs[i][ii] == nil {
					t.NestedPtrs[i][ii] = new(imported.Struct)
				}
				if err := t.NestedPtrs[i][ii].UnmarshalBorsh(r); err != nil {
					return err
				}
			}
		}
	}

	// field BytesIDs (4)
	if exist, err := borsh.ReadBool(r); err != nil {
		return err
	} else if exist {
		if lth, err := borsh.ReadUint32(r); err != nil {
			return err
		} else {
			t.BytesIDs = make([]imported.ByteAlias, lth)
		}
	}

	for i := range t.BytesIDs {
		if err := borsh.ReadBytes(r, t.BytesIDs[i][:]); err != nil {
			return err
		}
	}

	// field BytesCollection (5)
	if exist, err := borsh.ReadBool(r); err != nil {
		return err
	} else if exist {
		if lth, err := borsh.ReadUint32(r); err != nil {
			return err
		} else {
			t.BytesCollection = make(imported.BytesCollection, lth)
		}
	}

	for i := range t.BytesCollection {
		if exist, err := borsh.ReadBool(r); err != nil {
			return err
		} else if exist {
			if lth, err := borsh.ReadUint32(r); err != nil {
				return err
			} else {
				t.BytesCollection[i] = make([]uint8, lth)
			}
		}
		if err := borsh.ReadBytes(r, t.BytesCollection[i][:]); err != nil {
			return err
		}
	}

	// field AnotherImport (6)
	if exist, err := borsh.ReadBool(r); err != nil {
		return err
	} else if exist {
		if t.AnotherImport == nil {
			t.AnotherImport = new(lib.Struct)
		}
		if err := t.AnotherImport.UnmarshalBorsh(r); err != nil {
			return err
		}
	}

	return nil
}
