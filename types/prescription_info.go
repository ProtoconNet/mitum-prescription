package types

import (
	"github.com/ProtoconNet/mitum-currency/v3/common"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/hint"
	"github.com/pkg/errors"
)

var (
	MaxPrescriptionHashLen = 100
	MaxDataLen             = 100
)

var PrescriptionInfoHint = hint.MustNewHint("mitum-prescription-prescription-info-v0.0.1")

type PrescriptionInfo struct {
	hint.BaseHinter
	prescriptionHash string
	prescribeDate    uint64
	prepareDate      uint64
	endDate          uint64
	status           PrescriptionStatus
	hospital         string
	pharmacy         string
}

func NewPrescriptionInfo(
	prescriptionHash string,
	prescribeDate, prepareDate, endDate uint64,
	status PrescriptionStatus,
	hospital, pharmacy string,
) PrescriptionInfo {
	return PrescriptionInfo{
		BaseHinter:       hint.NewBaseHinter(PrescriptionInfoHint),
		prescriptionHash: prescriptionHash,
		prescribeDate:    prescribeDate,
		prepareDate:      prepareDate,
		endDate:          endDate,
		status:           status,
		hospital:         hospital,
		pharmacy:         pharmacy,
	}
}

func (p PrescriptionInfo) IsValid([]byte) error {
	if len(p.prescriptionHash) < 1 || len(p.prescriptionHash) > MaxPrescriptionHashLen {
		return errors.Errorf("invalid prescription hash length, %v is outside the allowed range (1 to %v)", len(p.prescriptionHash), MaxPrescriptionHashLen)
	}
	if !currencytypes.ReValidSpcecialCh.Match([]byte(p.prescriptionHash)) {
		return common.ErrValueInvalid.Wrap(errors.Errorf("prescription hash %s, must match regex `^[^\\s:/?#\\[\\]$@]*$`", p.prescriptionHash))
	}
	if len(p.hospital) < 1 || len(p.hospital) > MaxDataLen {
		return errors.Errorf(
			"invalid hospital name length, %v is outside the allowed range (1 to %v)", len(p.prescriptionHash),
			MaxDataLen,
		)
	}

	if len(p.pharmacy) > MaxDataLen {
		return errors.Errorf(
			"invalid pharmacy name length, %v exeeds the maximum allowed length of %v", len(p.prescriptionHash),
			MaxDataLen,
		)
	}
	return nil
}

func (p PrescriptionInfo) Bytes() []byte {
	return util.ConcatBytesSlice(
		[]byte(p.prescriptionHash),
		util.Uint64ToBytes(p.prepareDate),
		util.Uint64ToBytes(p.prepareDate),
		util.Uint64ToBytes(p.endDate),
		p.status.Bytes(),
		[]byte(p.hospital),
		[]byte(p.pharmacy),
	)
}

func (p PrescriptionInfo) PrescriptionHash() string {
	return p.prescriptionHash
}

func (p PrescriptionInfo) PrescribeDate() uint64 {
	return p.prescribeDate
}

func (p PrescriptionInfo) PrepareDate() uint64 {
	return p.prepareDate
}

func (p PrescriptionInfo) EndDate() uint64 {
	return p.endDate
}

func (p PrescriptionInfo) Status() PrescriptionStatus {
	return p.status
}

func (p PrescriptionInfo) Hospital() string {
	return p.hospital
}

func (p PrescriptionInfo) Pharmacy() string {
	return p.pharmacy
}

func (p PrescriptionInfo) Equal(ct PrescriptionInfo) bool {
	if p.prescriptionHash != ct.prescriptionHash {
		return false
	}

	if p.prescribeDate != ct.prescribeDate {
		return false
	}

	if p.prepareDate != ct.prepareDate {
		return false
	}

	if p.endDate != ct.endDate {
		return false
	}

	if p.status != ct.status {
		return false
	}

	if p.hospital != ct.hospital {
		return false
	}

	if p.pharmacy != ct.pharmacy {
		return false
	}

	return true
}
