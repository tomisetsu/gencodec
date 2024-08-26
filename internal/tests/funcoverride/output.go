// Code generated by github.com/AlexanderMint/gencodec. DO NOT EDIT.

package funcoverride

import (
	"encoding/json"
)

var _ = (*Zo)(nil)

// MarshalJSON marshals as JSON.
func (z Z) MarshalJSON() ([]byte, error) {
	type Z struct {
		S              string `json:"s"`
		I              int32  `json:"iVal"`
		Hash           string
		MultiplyIByTwo int64 `json:"multipliedByTwo"`
	}
	var enc Z
	enc.S = z.S
	enc.I = z.I
	enc.Hash = z.Hash()
	enc.MultiplyIByTwo = int64(z.MultiplyIByTwo())
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (z *Z) UnmarshalJSON(input []byte) error {
	type Z struct {
		S *string `json:"s"`
		I *int32  `json:"iVal"`
	}
	var dec Z
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.S != nil {
		z.S = *dec.S
	}
	if dec.I != nil {
		z.I = *dec.I
	}
	return nil
}

// MarshalYAML marshals as YAML.
func (z Z) MarshalYAML() (interface{}, error) {
	type Z struct {
		S              string `json:"s"`
		I              int32  `json:"iVal"`
		Hash           string
		MultiplyIByTwo int64 `json:"multipliedByTwo"`
	}
	var enc Z
	enc.S = z.S
	enc.I = z.I
	enc.Hash = z.Hash()
	enc.MultiplyIByTwo = int64(z.MultiplyIByTwo())
	return &enc, nil
}

// UnmarshalYAML unmarshals from YAML.
func (z *Z) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type Z struct {
		S *string `json:"s"`
		I *int32  `json:"iVal"`
	}
	var dec Z
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.S != nil {
		z.S = *dec.S
	}
	if dec.I != nil {
		z.I = *dec.I
	}
	return nil
}

// MarshalTOML marshals as TOML.
func (z Z) MarshalTOML() (interface{}, error) {
	type Z struct {
		S              string `json:"s"`
		I              int32  `json:"iVal"`
		Hash           string
		MultiplyIByTwo int64 `json:"multipliedByTwo"`
	}
	var enc Z
	enc.S = z.S
	enc.I = z.I
	enc.Hash = z.Hash()
	enc.MultiplyIByTwo = int64(z.MultiplyIByTwo())
	return &enc, nil
}

// UnmarshalTOML unmarshals from TOML.
func (z *Z) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Z struct {
		S *string `json:"s"`
		I *int32  `json:"iVal"`
	}
	var dec Z
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.S != nil {
		z.S = *dec.S
	}
	if dec.I != nil {
		z.I = *dec.I
	}
	return nil
}
