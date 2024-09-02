package prescription

import (
	"github.com/ProtoconNet/mitum-currency/v3/common"
	"github.com/ProtoconNet/mitum-currency/v3/types"
	mitumbase "github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/encoder"
)

type CreateDataFactJSONMarshaler struct {
	mitumbase.BaseFactJSONMarshaler
	Sender           mitumbase.Address `json:"sender"`
	Contract         mitumbase.Address `json:"contract"`
	PrescriptionHash string            `json:"prescription_hash"`
	PrescribeDate    uint64            `json:"prescribe_date"`
	EndDate          uint64            `json:"end_date"`
	Hospital         string            `json:"hospital"`
	Currency         types.CurrencyID  `json:"currency"`
}

func (fact RegisterPrescriptionFact) MarshalJSON() ([]byte, error) {
	return util.MarshalJSON(CreateDataFactJSONMarshaler{
		BaseFactJSONMarshaler: fact.BaseFact.JSONMarshaler(),
		Sender:                fact.sender,
		Contract:              fact.contract,
		PrescriptionHash:      fact.prescriptionHash,
		PrescribeDate:         fact.prescribeDate,
		EndDate:               fact.endDate,
		Hospital:              fact.hospital,
		Currency:              fact.currency,
	})
}

type CreateDataFactJSONUnmarshaler struct {
	mitumbase.BaseFactJSONUnmarshaler
	Sender           string `json:"sender"`
	Contract         string `json:"contract"`
	PrescriptionHash string `json:"prescription_hash"`
	PrescribeDate    uint64 `json:"prescribe_date"`
	EndDate          uint64 `json:"end_date"`
	Hospital         string `json:"hospital"`
	Currency         string `json:"currency"`
}

func (fact *RegisterPrescriptionFact) DecodeJSON(b []byte, enc encoder.Encoder) error {
	var u CreateDataFactJSONUnmarshaler
	if err := enc.Unmarshal(b, &u); err != nil {
		return common.DecorateError(err, common.ErrDecodeJson, *fact)
	}

	fact.BaseFact.SetJSONUnmarshaler(u.BaseFactJSONUnmarshaler)

	if err := fact.unpack(
		enc, u.Sender, u.Contract, u.PrescriptionHash, u.PrescribeDate,
		u.EndDate, u.Hospital, u.Currency); err != nil {
		return common.DecorateError(err, common.ErrDecodeJson, *fact)
	}

	return nil
}

func (op RegisterPrescription) MarshalJSON() ([]byte, error) {
	return util.MarshalJSON(
		op.BaseOperation.JSONMarshaler(),
	)
}

func (op *RegisterPrescription) DecodeJSON(b []byte, enc encoder.Encoder) error {
	var ubo common.BaseOperation
	if err := ubo.DecodeJSON(b, enc); err != nil {
		return common.DecorateError(err, common.ErrDecodeJson, *op)
	}

	op.BaseOperation = ubo

	return nil
}
