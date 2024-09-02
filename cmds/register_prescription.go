package cmds

import (
	"context"
	currencycmds "github.com/ProtoconNet/mitum-currency/v3/cmds"
	"github.com/ProtoconNet/mitum-prescription/operation/prescription"
	"github.com/ProtoconNet/mitum2/base"
	"github.com/ProtoconNet/mitum2/util"
	"github.com/pkg/errors"
)

type RegisterPrescriptionCommand struct {
	BaseCommand
	currencycmds.OperationFlags
	Sender           currencycmds.AddressFlag    `arg:"" name:"sender" help:"sender address" required:"true"`
	Contract         currencycmds.AddressFlag    `arg:"" name:"contract" help:"contract address" required:"true"`
	PrescriptionHash string                      `arg:"" name:"prescription_hash" help:"prescription hash" required:"true"`
	PrescribeDate    uint64                      `arg:"" name:"prescribe_date" help:"prescribe date" required:"true"`
	EndDate          uint64                      `arg:"" name:"end_date" help:"end date" required:"true"`
	Hospital         string                      `arg:"" name:"hospital" help:"hospital" required:"true"`
	Currency         currencycmds.CurrencyIDFlag `arg:"" name:"currency" help:"currency id" required:"true"`
	sender           base.Address
	contract         base.Address
}

func (cmd *RegisterPrescriptionCommand) Run(pctx context.Context) error { // nolint:dupl
	if _, err := cmd.prepare(pctx); err != nil {
		return err
	}

	if err := cmd.parseFlags(); err != nil {
		return err
	}

	op, err := cmd.createOperation()
	if err != nil {
		return err
	}

	currencycmds.PrettyPrint(cmd.Out, op)

	return nil
}

func (cmd *RegisterPrescriptionCommand) parseFlags() error {
	if err := cmd.OperationFlags.IsValid(nil); err != nil {
		return err
	}

	a, err := cmd.Sender.Encode(cmd.Encoders.JSON())
	if err != nil {
		return errors.Wrapf(err, "invalid sender format, %q", cmd.Sender)
	} else {
		cmd.sender = a
	}

	a, err = cmd.Contract.Encode(cmd.Encoders.JSON())
	if err != nil {
		return errors.Wrapf(err, "invalid contract format, %q", cmd.Contract)
	} else {
		cmd.contract = a
	}

	if len(cmd.PrescriptionHash) < 1 {
		return errors.Errorf("invalid prescription hash, %s", cmd.PrescriptionHash)
	}

	if len(cmd.Hospital) < 1 {
		return errors.Errorf("invalid hospital, %s", cmd.Hospital)
	}

	return nil
}

func (cmd *RegisterPrescriptionCommand) createOperation() (base.Operation, error) { // nolint:dupl
	e := util.StringError("failed to create register prescription operation")

	fact := prescription.NewRegisterPrescriptionFact(
		[]byte(cmd.Token), cmd.sender, cmd.contract, cmd.PrescriptionHash,
		cmd.PrescribeDate, cmd.EndDate, cmd.Hospital, cmd.Currency.CID,
	)

	op, err := prescription.NewRegisterPrescription(fact)
	if err != nil {
		return nil, e.Wrap(err)
	}
	err = op.Sign(cmd.Privatekey, cmd.NetworkID.NetworkID())
	if err != nil {
		return nil, e.Wrap(err)
	}

	return op, nil
}
