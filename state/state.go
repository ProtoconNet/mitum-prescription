package state

import (
	"fmt"
	"github.com/ProtoconNet/mitum-currency/v3/common"
	"github.com/ProtoconNet/mitum-prescription/types"
	"strings"

	mitumbase "github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/hint"
	"github.com/pkg/errors"
)

var (
	DesignStateValueHint       = hint.MustNewHint("mitum-prescription-design-state-value-v0.0.1")
	PrescriptionStateKeyPrefix = "prescription"
	DesignStateKeySuffix       = "design"
)

func PrescriptionStateKey(addr mitumbase.Address) string {
	return fmt.Sprintf("%s:%s", PrescriptionStateKeyPrefix, addr.String())
}

type DesignStateValue struct {
	hint.BaseHinter
	Design types.Design
}

func NewDesignStateValue(design types.Design) DesignStateValue {
	return DesignStateValue{
		BaseHinter: hint.NewBaseHinter(DesignStateValueHint),
		Design:     design,
	}
}

func (sv DesignStateValue) Hint() hint.Hint {
	return sv.BaseHinter.Hint()
}

func (sv DesignStateValue) IsValid([]byte) error {
	e := util.ErrInvalid.Errorf("invalid DesignStateValue")

	if err := sv.BaseHinter.IsValid(DesignStateValueHint.Type().Bytes()); err != nil {
		return e.Wrap(err)
	}

	if err := sv.Design.IsValid(nil); err != nil {
		return e.Wrap(err)
	}

	return nil
}

func (sv DesignStateValue) HashBytes() []byte {
	return sv.Design.Bytes()
}

func GetDesignFromState(st mitumbase.State) (types.Design, error) {
	v := st.Value()
	if v == nil {
		return types.Design{}, errors.Errorf("state value is nil")
	}

	d, ok := v.(DesignStateValue)
	if !ok {
		return types.Design{}, errors.Errorf("expected DesignStateValue but %T", v)
	}

	return d.Design, nil
}

func IsDesignStateKey(key string) bool {
	return strings.HasPrefix(key, PrescriptionStateKeyPrefix) && strings.HasSuffix(key, DesignStateKeySuffix)
}

func DesignStateKey(addr mitumbase.Address) string {
	return fmt.Sprintf("%s:%s", PrescriptionStateKey(addr), DesignStateKeySuffix)
}

var (
	PrescriptionInfoStateValueHint = hint.MustNewHint("mitum-prescription-prescription-info-state-value-v0.0.1")
	PrescriptionInfoStateKeySuffix = "prescriptioninfo"
)

type PrescriptionInfoStateValue struct {
	hint.BaseHinter
	PrescriptionInfo types.PrescriptionInfo
}

func NewPrescriptionInfoStateValue(prescriptionInfo types.PrescriptionInfo) PrescriptionInfoStateValue {
	return PrescriptionInfoStateValue{
		BaseHinter:       hint.NewBaseHinter(PrescriptionInfoStateValueHint),
		PrescriptionInfo: prescriptionInfo,
	}
}

func (sv PrescriptionInfoStateValue) Hint() hint.Hint {
	return sv.BaseHinter.Hint()
}

func (sv PrescriptionInfoStateValue) IsValid([]byte) error {
	e := util.ErrInvalid.Errorf("invalid PrescriptionInfoStateValue")

	if err := sv.BaseHinter.IsValid(PrescriptionInfoStateValueHint.Type().Bytes()); err != nil {
		return e.Wrap(err)
	}

	if err := sv.PrescriptionInfo.IsValid(nil); err != nil {
		return e.Wrap(err)
	}

	return nil
}

func (sv PrescriptionInfoStateValue) HashBytes() []byte {
	return sv.PrescriptionInfo.Bytes()
}

func GetPrescriptionInfoFromState(st mitumbase.State) (types.PrescriptionInfo, error) {
	v := st.Value()
	if v == nil {
		return types.PrescriptionInfo{}, errors.Errorf("State value is nil")
	}

	ts, ok := v.(PrescriptionInfoStateValue)
	if !ok {
		return types.PrescriptionInfo{}, common.ErrTypeMismatch.Wrap(
			errors.Errorf("expected PrescriptionInfoStateValue found, %T", v))
	}

	return ts.PrescriptionInfo, nil
}

func IsPrescriptionInfoStateKey(key string) bool {
	return strings.HasPrefix(key, PrescriptionStateKeyPrefix) && strings.HasSuffix(key, PrescriptionInfoStateKeySuffix)
}

func PrescriptionInfoStateKey(addr mitumbase.Address, key string) string {
	return fmt.Sprintf("%s:%s:%s", PrescriptionStateKey(addr), key, PrescriptionInfoStateKeySuffix)
}
