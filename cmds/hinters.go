package cmds

import (
	currencycmds "github.com/ProtoconNet/mitum-currency/v3/cmds"
	"github.com/ProtoconNet/mitum-prescription/operation/prescription"
	"github.com/ProtoconNet/mitum-prescription/state"
	"github.com/ProtoconNet/mitum-prescription/types"
	"github.com/ProtoconNet/mitum2/launch"
	"github.com/ProtoconNet/mitum2/util/encoder"
	"github.com/pkg/errors"
)

var Hinters []encoder.DecodeDetail
var SupportedProposalOperationFactHinters []encoder.DecodeDetail

var AddedHinters = []encoder.DecodeDetail{
	// revive:disable-next-line:line-length-limit

	{Hint: types.DesignHint, Instance: types.Design{}},
	{Hint: types.PrescriptionInfoHint, Instance: types.PrescriptionInfo{}},

	{Hint: prescription.RegisterPrescriptionHint, Instance: prescription.RegisterPrescription{}},
	{Hint: prescription.UsePrescriptionHint, Instance: prescription.UsePrescription{}},
	{Hint: prescription.RegisterModelHint, Instance: prescription.RegisterModel{}},

	{Hint: state.PrescriptionInfoStateValueHint, Instance: state.PrescriptionInfoStateValue{}},
	{Hint: state.DesignStateValueHint, Instance: state.DesignStateValue{}},
}

var AddedSupportedHinters = []encoder.DecodeDetail{
	{Hint: prescription.RegisterPrescriptionFactHint, Instance: prescription.RegisterPrescriptionFact{}},
	{Hint: prescription.UsePrescriptionFactHint, Instance: prescription.UsePrescriptionFact{}},
	{Hint: prescription.RegisterModelFactHint, Instance: prescription.RegisterModelFact{}},
}

func init() {
	defaultLen := len(launch.Hinters)
	currencyExtendedLen := defaultLen + len(currencycmds.AddedHinters)
	allExtendedLen := currencyExtendedLen + len(AddedHinters)

	Hinters = make([]encoder.DecodeDetail, allExtendedLen)
	copy(Hinters, launch.Hinters)
	copy(Hinters[defaultLen:currencyExtendedLen], currencycmds.AddedHinters)
	copy(Hinters[currencyExtendedLen:], AddedHinters)

	defaultSupportedLen := len(launch.SupportedProposalOperationFactHinters)
	currencySupportedExtendedLen := defaultSupportedLen + len(currencycmds.AddedSupportedHinters)
	allSupportedExtendedLen := currencySupportedExtendedLen + len(AddedSupportedHinters)

	SupportedProposalOperationFactHinters = make(
		[]encoder.DecodeDetail,
		allSupportedExtendedLen)
	copy(SupportedProposalOperationFactHinters, launch.SupportedProposalOperationFactHinters)
	copy(SupportedProposalOperationFactHinters[defaultSupportedLen:currencySupportedExtendedLen], currencycmds.AddedSupportedHinters)
	copy(SupportedProposalOperationFactHinters[currencySupportedExtendedLen:], AddedSupportedHinters)
}

func LoadHinters(encs *encoder.Encoders) error {
	for i := range Hinters {
		if err := encs.AddDetail(Hinters[i]); err != nil {
			return errors.Wrap(err, "add hinter to encoder")
		}
	}

	for i := range SupportedProposalOperationFactHinters {
		if err := encs.AddDetail(SupportedProposalOperationFactHinters[i]); err != nil {
			return errors.Wrap(err, "add supported proposal operation fact hinter to encoder")
		}
	}

	return nil
}
