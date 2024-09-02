package prescription

import (
	"github.com/ProtoconNet/mitum-currency/v3/types"
	mitumbase "github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util/encoder"
)

func (fact *UsePrescriptionFact) unpack(
	enc encoder.Encoder,
	sa, ta, pHash string,
	prDate uint64,
	pharmacy, cid string,
) error {
	switch sender, err := mitumbase.DecodeAddress(sa, enc); {
	case err != nil:
		return err
	default:
		fact.sender = sender
	}

	switch contract, err := mitumbase.DecodeAddress(ta, enc); {
	case err != nil:
		return err
	default:
		fact.contract = contract
	}

	fact.prescriptionHash = pHash
	fact.prepareDate = prDate
	fact.pharmacy = pharmacy
	fact.currency = types.CurrencyID(cid)

	return nil
}
