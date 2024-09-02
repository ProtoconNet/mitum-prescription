package prescription

import (
	"github.com/ProtoconNet/mitum-currency/v3/common"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum-prescription/types"
	mitumbase "github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/hint"
	"github.com/ProtoconNet/mitum2/util/valuehash"
	"github.com/pkg/errors"
)

var (
	RegisterPrescriptionFactHint = hint.MustNewHint("mitum-prescription-register-prescription-operation-fact-v0.0.1")
	RegisterPrescriptionHint     = hint.MustNewHint("mitum-prescription-register-prescription-operation-v0.0.1")
)

type RegisterPrescriptionFact struct {
	mitumbase.BaseFact
	sender           mitumbase.Address
	contract         mitumbase.Address
	prescriptionHash string
	prescribeDate    uint64
	endDate          uint64
	hospital         string
	currency         currencytypes.CurrencyID
}

func NewRegisterPrescriptionFact(
	token []byte, sender, contract mitumbase.Address,
	prescriptionHash string, prescribeDate, endDate uint64, hospital string,
	currency currencytypes.CurrencyID) RegisterPrescriptionFact {
	bf := mitumbase.NewBaseFact(RegisterPrescriptionFactHint, token)
	fact := RegisterPrescriptionFact{
		BaseFact:         bf,
		sender:           sender,
		contract:         contract,
		prescriptionHash: prescriptionHash,
		prescribeDate:    prescribeDate,
		endDate:          endDate,
		hospital:         hospital,
		currency:         currency,
	}

	fact.SetHash(fact.GenerateHash())
	return fact
}

func (fact RegisterPrescriptionFact) IsValid(b []byte) error {
	if fact.sender.Equal(fact.contract) {
		return common.ErrFactInvalid.Wrap(
			common.ErrSelfTarget.Wrap(errors.Errorf("sender %v is same with contract account", fact.sender)))
	}

	if len(fact.PrescriptionHash()) < 1 || len(fact.PrescriptionHash()) > types.MaxPrescriptionHashLen {
		return common.ErrFactInvalid.Wrap(
			common.ErrValOOR.Wrap(
				errors.Errorf(
					"invalid prescription hash length, %v is outside the allowed range (1 to %v)",
					len(fact.PrescriptionHash()),
					types.MaxPrescriptionHashLen,
				),
			),
		)
	}
	if len(fact.Hospital()) > types.MaxDataLen {
		return common.ErrFactInvalid.Wrap(
			common.ErrValOOR.Wrap(
				errors.Errorf(
					"invalid hospital name length, %v exeeds the maximum allowed length of %v", len(fact.Hospital()),
					types.MaxDataLen,
				),
			),
		)
	}

	if err := util.CheckIsValiders(nil, false,
		fact.BaseHinter,
		fact.sender,
		fact.contract,
		fact.currency,
	); err != nil {
		return common.ErrFactInvalid.Wrap(err)
	}

	if err := common.IsValidOperationFact(fact, b); err != nil {
		return common.ErrFactInvalid.Wrap(err)
	}

	return nil
}

func (fact RegisterPrescriptionFact) Hash() util.Hash {
	return fact.BaseFact.Hash()
}

func (fact RegisterPrescriptionFact) GenerateHash() util.Hash {
	return valuehash.NewSHA256(fact.Bytes())
}

func (fact RegisterPrescriptionFact) Bytes() []byte {
	return util.ConcatBytesSlice(
		fact.Token(),
		fact.sender.Bytes(),
		fact.contract.Bytes(),
		[]byte(fact.prescriptionHash),
		util.Uint64ToBytes(fact.prescribeDate),
		util.Uint64ToBytes(fact.endDate),
		[]byte(fact.hospital),
		fact.currency.Bytes(),
	)
}

func (fact RegisterPrescriptionFact) Token() mitumbase.Token {
	return fact.BaseFact.Token()
}

func (fact RegisterPrescriptionFact) Sender() mitumbase.Address {
	return fact.sender
}

func (fact RegisterPrescriptionFact) Contract() mitumbase.Address {
	return fact.contract
}

func (fact RegisterPrescriptionFact) PrescriptionHash() string {
	return fact.prescriptionHash
}

func (fact RegisterPrescriptionFact) PrescribeDate() uint64 {
	return fact.prescribeDate
}

func (fact RegisterPrescriptionFact) EndDate() uint64 {
	return fact.endDate
}

func (fact RegisterPrescriptionFact) Hospital() string {
	return fact.hospital
}

func (fact RegisterPrescriptionFact) Currency() currencytypes.CurrencyID {
	return fact.currency
}

type RegisterPrescription struct {
	common.BaseOperation
}

func NewRegisterPrescription(fact RegisterPrescriptionFact) (RegisterPrescription, error) {
	return RegisterPrescription{BaseOperation: common.NewBaseOperation(RegisterPrescriptionHint, fact)}, nil
}
