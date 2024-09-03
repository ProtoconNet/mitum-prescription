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
	UsePrescriptionFactHint = hint.MustNewHint("mitum-prescription-use-prescription-operation-fact-v0.0.1")
	UsePrescriptionHint     = hint.MustNewHint("mitum-prescription-use-prescription-operation-v0.0.1")
)

type UsePrescriptionFact struct {
	mitumbase.BaseFact
	sender           mitumbase.Address
	contract         mitumbase.Address
	prescriptionHash string
	prepareDate      uint64
	pharmacy         string
	currency         currencytypes.CurrencyID
}

func NewUsePrescriptionFact(
	token []byte, sender, contract mitumbase.Address,
	prescriptionHash string, prepareDate uint64, pharmacy string,
	currency currencytypes.CurrencyID) UsePrescriptionFact {
	bf := mitumbase.NewBaseFact(UsePrescriptionFactHint, token)
	fact := UsePrescriptionFact{
		BaseFact:         bf,
		sender:           sender,
		contract:         contract,
		prescriptionHash: prescriptionHash,
		prepareDate:      prepareDate,
		pharmacy:         pharmacy,
		currency:         currency,
	}

	fact.SetHash(fact.GenerateHash())
	return fact
}

func (fact UsePrescriptionFact) IsValid(b []byte) error {
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
	if !currencytypes.ReValidSpcecialCh.Match([]byte(fact.PrescriptionHash())) {
		return common.ErrValueInvalid.Wrap(errors.Errorf("prescription hash %s, must match regex `^[^\\s:/?#\\[\\]$@]*$`", fact.PrescriptionHash()))
	}
	if len(fact.Pharmacy()) < 1 || len(fact.Pharmacy()) > types.MaxDataLen {
		return common.ErrFactInvalid.Wrap(
			common.ErrValOOR.Wrap(
				errors.Errorf(
					"invalid pharmacy name length, %v is outside the allowed range (1 to %v)", len(fact.Pharmacy()),
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

func (fact UsePrescriptionFact) Hash() util.Hash {
	return fact.BaseFact.Hash()
}

func (fact UsePrescriptionFact) GenerateHash() util.Hash {
	return valuehash.NewSHA256(fact.Bytes())
}

func (fact UsePrescriptionFact) Bytes() []byte {
	return util.ConcatBytesSlice(
		fact.Token(),
		fact.sender.Bytes(),
		fact.contract.Bytes(),
		[]byte(fact.prescriptionHash),
		util.Uint64ToBytes(fact.prepareDate),
		[]byte(fact.pharmacy),
		fact.currency.Bytes(),
	)
}

func (fact UsePrescriptionFact) Token() mitumbase.Token {
	return fact.BaseFact.Token()
}

func (fact UsePrescriptionFact) Sender() mitumbase.Address {
	return fact.sender
}

func (fact UsePrescriptionFact) Contract() mitumbase.Address {
	return fact.contract
}

func (fact UsePrescriptionFact) PrescriptionHash() string {
	return fact.prescriptionHash
}

func (fact UsePrescriptionFact) PrepareDate() uint64 {
	return fact.prepareDate
}

func (fact UsePrescriptionFact) Pharmacy() string {
	return fact.pharmacy
}

func (fact UsePrescriptionFact) Currency() currencytypes.CurrencyID {
	return fact.currency
}

type UsePrescription struct {
	common.BaseOperation
}

func NewUsePrescription(fact UsePrescriptionFact) (UsePrescription, error) {
	return UsePrescription{BaseOperation: common.NewBaseOperation(UsePrescriptionHint, fact)}, nil
}
