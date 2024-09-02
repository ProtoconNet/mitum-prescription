package prescription

import (
	"github.com/ProtoconNet/mitum-currency/v3/common"
	"github.com/ProtoconNet/mitum-currency/v3/types"
	mitumbase "github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/encoder"
)

type UpdateDataFactJSONMarshaler struct {
	mitumbase.BaseFactJSONMarshaler
	Sender           mitumbase.Address `json:"sender"`
	Contract         mitumbase.Address `json:"contract"`
	PrescriptionHash string            `json:"prescription_hash"`
	PrepareDate      uint64            `json:"prepare_date"`
	Pharmacy         string            `json:"pharmacy"`
	Currency         types.CurrencyID  `json:"currency"`
}

func (fact UsePrescriptionFact) MarshalJSON() ([]byte, error) {
	return util.MarshalJSON(UpdateDataFactJSONMarshaler{
		BaseFactJSONMarshaler: fact.BaseFact.JSONMarshaler(),
		Sender:                fact.sender,
		Contract:              fact.contract,
		PrescriptionHash:      fact.prescriptionHash,
		PrepareDate:           fact.prepareDate,
		Pharmacy:              fact.pharmacy,
		Currency:              fact.currency,
	})
}

type UpdateDataFactJSONUnmarshaler struct {
	mitumbase.BaseFactJSONUnmarshaler
	Sender           string `json:"sender"`
	Contract         string `json:"contract"`
	PrescriptionHash string `json:"prescription_hash"`
	PrepareDate      uint64 `json:"prepare_date"`
	Pharmacy         string `json:"pharmacy"`
	Currency         string `json:"currency"`
}

func (fact *UsePrescriptionFact) DecodeJSON(b []byte, enc encoder.Encoder) error {
	var u UpdateDataFactJSONUnmarshaler
	if err := enc.Unmarshal(b, &u); err != nil {
		return common.DecorateError(err, common.ErrDecodeJson, *fact)
	}

	fact.BaseFact.SetJSONUnmarshaler(u.BaseFactJSONUnmarshaler)

	if err := fact.unpack(
		enc, u.Sender, u.Contract, u.PrescriptionHash, u.PrepareDate,
		u.Pharmacy, u.Currency); err != nil {
		return common.DecorateError(err, common.ErrDecodeJson, *fact)
	}

	return nil
}

func (op UsePrescription) MarshalJSON() ([]byte, error) {
	return util.MarshalJSON(
		op.BaseOperation.JSONMarshaler(),
	)
}

func (op *UsePrescription) DecodeJSON(b []byte, enc encoder.Encoder) error {
	var ubo common.BaseOperation
	if err := ubo.DecodeJSON(b, enc); err != nil {
		return common.DecorateError(err, common.ErrDecodeJson, *op)
	}

	op.BaseOperation = ubo

	return nil
}
