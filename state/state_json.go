package state

import (
	"encoding/json"
	"github.com/ProtoconNet/mitum-prescription/types"

	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/encoder"
	"github.com/ProtoconNet/mitum2/util/hint"
)

type DesignStateValueJSONMarshaler struct {
	hint.BaseHinter
	Design types.Design `json:"design"`
}

func (sv DesignStateValue) MarshalJSON() ([]byte, error) {
	return util.MarshalJSON(
		DesignStateValueJSONMarshaler(sv),
	)
}

type DesignStateValueJSONUnmarshaler struct {
	Hint   hint.Hint       `json:"_hint"`
	Design json.RawMessage `json:"design"`
}

func (sv *DesignStateValue) DecodeJSON(b []byte, enc encoder.Encoder) error {
	e := util.StringError("failed to decode json of DesignStateValue")

	var u DesignStateValueJSONUnmarshaler
	if err := enc.Unmarshal(b, &u); err != nil {
		return e.Wrap(err)
	}

	sv.BaseHinter = hint.NewBaseHinter(u.Hint)

	var sd types.Design
	if err := sd.DecodeJSON(u.Design, enc); err != nil {
		return e.Wrap(err)
	}
	sv.Design = sd

	return nil
}

type PrescriptionInfoStateValueJSONMarshaler struct {
	hint.BaseHinter
	PrescriptionInfo types.PrescriptionInfo `json:"prescription_info"`
}

func (sv PrescriptionInfoStateValue) MarshalJSON() ([]byte, error) {
	return util.MarshalJSON(
		PrescriptionInfoStateValueJSONMarshaler(sv),
	)
}

type PrescriptionInfoStateValueJSONUnmarshaler struct {
	Hint             hint.Hint       `json:"_hint"`
	PrescriptionInfo json.RawMessage `json:"prescription_info"`
}

func (sv *PrescriptionInfoStateValue) DecodeJSON(b []byte, enc encoder.Encoder) error {
	e := util.StringError("decode json of PrescriptionInfoStateValue")

	var u PrescriptionInfoStateValueJSONUnmarshaler
	if err := enc.Unmarshal(b, &u); err != nil {
		return e.Wrap(err)
	}

	sv.BaseHinter = hint.NewBaseHinter(u.Hint)

	var t types.PrescriptionInfo
	if err := t.DecodeJSON(u.PrescriptionInfo, enc); err != nil {
		return e.Wrap(err)
	}
	sv.PrescriptionInfo = t

	return nil
}
