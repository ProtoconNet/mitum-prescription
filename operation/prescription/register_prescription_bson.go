package prescription

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/ProtoconNet/mitum-currency/v3/common"
	bsonenc "github.com/ProtoconNet/mitum-currency/v3/digest/util/bson"
	"github.com/ProtoconNet/mitum2/util/hint"
	"github.com/ProtoconNet/mitum2/util/valuehash"
)

func (fact RegisterPrescriptionFact) MarshalBSON() ([]byte, error) {
	return bsonenc.Marshal(
		bson.M{
			"_hint":             fact.Hint().String(),
			"hash":              fact.BaseFact.Hash().String(),
			"token":             fact.BaseFact.Token(),
			"sender":            fact.sender,
			"contract":          fact.contract,
			"prescription_hash": fact.prescriptionHash,
			"prescribe_date":    fact.prescribeDate,
			"end_date":          fact.endDate,
			"hospital":          fact.hospital,
			"currency":          fact.currency,
		},
	)
}

type RegisterPrescriptionFactBSONUnmarshaler struct {
	Hint             string `bson:"_hint"`
	Sender           string `bson:"sender"`
	Contract         string `bson:"contract"`
	PrescriptionHash string `bson:"prescription_hash"`
	PrescribeDate    uint64 `bson:"prescribe_date"`
	EndDate          uint64 `bson:"end_date"`
	Hospital         string `bson:"hospital"`
	Currency         string `bson:"currency"`
}

func (fact *RegisterPrescriptionFact) DecodeBSON(b []byte, enc *bsonenc.Encoder) error {
	var u common.BaseFactBSONUnmarshaler

	err := enc.Unmarshal(b, &u)
	if err != nil {
		return common.DecorateError(err, common.ErrDecodeBson, *fact)
	}

	fact.BaseFact.SetHash(valuehash.NewBytesFromString(u.Hash))
	fact.BaseFact.SetToken(u.Token)

	var uf RegisterPrescriptionFactBSONUnmarshaler
	if err := bson.Unmarshal(b, &uf); err != nil {
		return common.DecorateError(err, common.ErrDecodeBson, *fact)
	}

	ht, err := hint.ParseHint(uf.Hint)
	if err != nil {
		return common.DecorateError(err, common.ErrDecodeBson, *fact)
	}
	fact.BaseHinter = hint.NewBaseHinter(ht)

	if err := fact.unpack(
		enc, uf.Sender, uf.Contract, uf.PrescriptionHash, uf.PrescribeDate,
		uf.EndDate, uf.Hospital, uf.Currency,
	); err != nil {
		return common.DecorateError(err, common.ErrDecodeBson, *fact)
	}

	return nil
}

func (op RegisterPrescription) MarshalBSON() ([]byte, error) {
	return bsonenc.Marshal(
		bson.M{
			"_hint": op.Hint().String(),
			"hash":  op.Hash().String(),
			"fact":  op.Fact(),
			"signs": op.Signs(),
		})
}

func (op *RegisterPrescription) DecodeBSON(b []byte, enc *bsonenc.Encoder) error {
	var ubo common.BaseOperation
	if err := ubo.DecodeBSON(b, enc); err != nil {
		return common.DecorateError(err, common.ErrDecodeBson, *op)
	}

	op.BaseOperation = ubo

	return nil
}
