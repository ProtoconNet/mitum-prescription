package types

import (
	"go.mongodb.org/mongo-driver/bson"

	bsonenc "github.com/ProtoconNet/mitum-currency/v3/digest/util/bson"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/hint"
)

func (p PrescriptionInfo) MarshalBSON() ([]byte, error) {
	return bsonenc.Marshal(bson.M{
		"_hint":             p.Hint().String(),
		"prescription_hash": p.prescriptionHash,
		"prescribe_date":    p.prescribeDate,
		"prepare_date":      p.prepareDate,
		"end_date":          p.endDate,
		"status":            p.status.String(),
		"hospital":          p.hospital,
		"pharmacy":          p.pharmacy,
	})
}

type PrescriptionInfoBSONUnmarshaler struct {
	Hint             string `bson:"_hint"`
	PrescriptionHash string `bson:"prescription_hash"`
	PrescribeDate    uint64 `bson:"prescribe_date"`
	PrepareDate      uint64 `bson:"prepare_date"`
	EndDate          uint64 `bson:"end_date"`
	Status           string `bson:"status"`
	Hospital         string `bson:"hospital"`
	Pharmacy         string `bson:"pharmacy"`
}

func (p *PrescriptionInfo) DecodeBSON(b []byte, enc *bsonenc.Encoder) error {
	e := util.StringError("decode bson of PrescriptionInfo")

	var u PrescriptionInfoBSONUnmarshaler
	if err := enc.Unmarshal(b, &u); err != nil {
		return e.Wrap(err)
	}

	ht, err := hint.ParseHint(u.Hint)
	if err != nil {
		return e.Wrap(err)
	}

	return p.encode(ht, u.PrescriptionHash, u.PrescribeDate, u.PrepareDate, u.EndDate, u.Status, u.Hospital, u.Pharmacy)
}
