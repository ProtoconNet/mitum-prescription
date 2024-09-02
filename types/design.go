package types

import (
	"github.com/ProtoconNet/mitum2/util"
	"github.com/ProtoconNet/mitum2/util/hint"
	"github.com/ProtoconNet/mitum2/util/valuehash"
)

var DesignHint = hint.MustNewHint("mitum-prescription-design-v0.0.1")

type Design struct {
	hint.BaseHinter
}

func NewDesign() Design {
	return Design{
		BaseHinter: hint.NewBaseHinter(DesignHint),
	}
}

func (de Design) IsValid([]byte) error {
	if err := util.CheckIsValiders(nil, false,
		de.BaseHinter,
	); err != nil {
		return err
	}

	return nil
}

func (de Design) Bytes() []byte {
	return util.ConcatBytesSlice([]byte{})
}

func (de Design) Hash() util.Hash {
	return de.GenerateHash()
}

func (de Design) GenerateHash() util.Hash {
	return valuehash.NewSHA256(de.Bytes())
}
