// Code generated by github.com/tomisetsu/gencodec. DO NOT EDIT.

package arrayconv

import (
	"encoding/json"
	"errors"
)

var _ = (*Xo)(nil)

// MarshalJSON marshals as JSON.
func (x X) MarshalJSON() ([]byte, error) {
	type X struct {
		A         []int
		A2        []MyInt
		RequiredA []int `gencodec:"required"`
		S         [16]int64
		RequiredS [16]int64 `gencodec:"required"`
	}
	var enc X
	enc.A = x.A[:]
	enc.A2 = make([]MyInt, len(x.A2))
	for k, v := range x.A2 {
		enc.A2[k] = MyInt(v)
	}
	enc.RequiredA = x.RequiredA[:]
	copy(enc.S[:], x.S)
	copy(enc.RequiredS[:], x.RequiredS)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *X) UnmarshalJSON(input []byte) error {
	type X struct {
		A         []int
		A2        []MyInt
		RequiredA []int `gencodec:"required"`
		S         *[16]int64
		RequiredS *[16]int64 `gencodec:"required"`
	}
	var dec X
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.A != nil {
		if len(dec.A) != len(x.A) {
			return errors.New("field 'a' has wrong length, need 32 items")
		}
		copy(x.A[:], dec.A)
	}
	if dec.A2 != nil {
		if len(dec.A2) != len(x.A2) {
			return errors.New("field 'a2' has wrong length, need 32 items")
		}
		for k, v := range dec.A2 {
			x.A2[k] = int(v)
		}
	}
	if dec.RequiredA == nil {
		return errors.New("missing required field 'requiredA' for X")
	}
	if len(dec.RequiredA) != len(x.RequiredA) {
		return errors.New("field 'requiredA' has wrong length, need 32 items")
	}
	copy(x.RequiredA[:], dec.RequiredA)
	if dec.S != nil {
		x.S = (*dec.S)[:]
	}
	if dec.RequiredS == nil {
		return errors.New("missing required field 'requiredS' for X")
	}
	x.RequiredS = (*dec.RequiredS)[:]
	return nil
}
