package processor

import (
	"fmt"
	"github.com/ProtoconNet/mitum-currency/v3/operation/currency"
	"github.com/ProtoconNet/mitum-currency/v3/operation/extension"
	currencyprocessor "github.com/ProtoconNet/mitum-currency/v3/operation/processor"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum-prescription/operation/prescription"
	mitumbase "github.com/ProtoconNet/mitum2/base"
	"github.com/pkg/errors"
)

const (
	DuplicationTypeSender       currencytypes.DuplicationType = "sender"
	DuplicationTypeCurrency     currencytypes.DuplicationType = "currency"
	DuplicationTypeContract     currencytypes.DuplicationType = "contract"
	DuplicationTypePrescription currencytypes.DuplicationType = "prescription"
)

func CheckDuplication(opr *currencyprocessor.OperationProcessor, op mitumbase.Operation) error {
	opr.Lock()
	defer opr.Unlock()

	var duplicationTypeSenderID string
	var duplicationTypeCurrencyID string
	var duplicationTypeContractID string
	var duplicationTypePrescription string
	var newAddresses []mitumbase.Address

	switch t := op.(type) {
	case currency.CreateAccount:
		fact, ok := t.Fact().(currency.CreateAccountFact)
		if !ok {
			return errors.Errorf("expected %T, not %T", currency.CreateAccountFact{}, t.Fact())
		}
		as, err := fact.Targets()
		if err != nil {
			return errors.Errorf("failed to get Addresses")
		}
		newAddresses = as
		duplicationTypeSenderID = currencyprocessor.DuplicationKey(fact.Sender().String(), DuplicationTypeSender)
	case currency.UpdateKey:
		fact, ok := t.Fact().(currency.UpdateKeyFact)
		if !ok {
			return errors.Errorf("expected %T, not %T", currency.UpdateKeyFact{}, t.Fact())
		}
		duplicationTypeSenderID = currencyprocessor.DuplicationKey(fact.Sender().String(), DuplicationTypeSender)
	case currency.Transfer:
		fact, ok := t.Fact().(currency.TransferFact)
		if !ok {
			return errors.Errorf("expected %T, not %T", currency.TransferFact{}, t.Fact())
		}
		duplicationTypeSenderID = currencyprocessor.DuplicationKey(fact.Sender().String(), DuplicationTypeSender)
	case currency.RegisterCurrency:
		fact, ok := t.Fact().(currency.RegisterCurrencyFact)
		if !ok {
			return errors.Errorf("expected RegisterCurrencyFact, not %T", t.Fact())
		}
		duplicationTypeCurrencyID = currencyprocessor.DuplicationKey(fact.Currency().Currency().String(), DuplicationTypeCurrency)
	case currency.UpdateCurrency:
		fact, ok := t.Fact().(currency.UpdateCurrencyFact)
		if !ok {
			return errors.Errorf("expected %T, not %T", currency.UpdateCurrencyFact{}, t.Fact())
		}
		duplicationTypeCurrencyID = currencyprocessor.DuplicationKey(fact.Currency().String(), DuplicationTypeCurrency)
	case currency.Mint:
	case extension.CreateContractAccount:
		fact, ok := t.Fact().(extension.CreateContractAccountFact)
		if !ok {
			return errors.Errorf("expected %T, not %T", extension.CreateContractAccountFact{}, t.Fact())
		}
		as, err := fact.Targets()
		if err != nil {
			return errors.Errorf("failed to get Addresses")
		}
		newAddresses = as
		duplicationTypeSenderID = currencyprocessor.DuplicationKey(fact.Sender().String(), DuplicationTypeSender)
		duplicationTypeContractID = currencyprocessor.DuplicationKey(fact.Sender().String(), DuplicationTypeContract)
	case extension.Withdraw:
		fact, ok := t.Fact().(extension.WithdrawFact)
		if !ok {
			return errors.Errorf("expected WithdrawFact, not %T", t.Fact())
		}
		duplicationTypeSenderID = currencyprocessor.DuplicationKey(fact.Sender().String(), DuplicationTypeSender)
	case prescription.RegisterModel:
		fact, ok := t.Fact().(prescription.RegisterModelFact)
		if !ok {
			return errors.Errorf("expected %T, not %T", prescription.RegisterModelFact{}, t.Fact())
		}
		duplicationTypeSenderID = currencyprocessor.DuplicationKey(fact.Sender().String(), DuplicationTypeSender)
		duplicationTypeContractID = currencyprocessor.DuplicationKey(fact.Contract().String(), DuplicationTypeContract)
	case prescription.RegisterPrescription:
		fact, ok := t.Fact().(prescription.RegisterPrescriptionFact)
		if !ok {
			return errors.Errorf("expected %T, not %T", prescription.RegisterPrescriptionFact{}, t.Fact())
		}
		duplicationTypeSenderID = currencyprocessor.DuplicationKey(fact.Sender().String(), DuplicationTypeSender)
		duplicationTypePrescription = currencyprocessor.DuplicationKey(
			fmt.Sprintf("%s:%s", fact.Contract().String(), fact.PrescriptionHash()), DuplicationTypePrescription)
	case prescription.UsePrescription:
		fact, ok := t.Fact().(prescription.UsePrescriptionFact)
		if !ok {
			return errors.Errorf("expected UsePrescriptionFact, not %T", t.Fact())
		}
		duplicationTypeSenderID = currencyprocessor.DuplicationKey(fact.Sender().String(), DuplicationTypeSender)
		duplicationTypePrescription = currencyprocessor.DuplicationKey(
			fmt.Sprintf("%s:%s", fact.Contract().String(), fact.PrescriptionHash()), DuplicationTypePrescription)
	default:
		return nil
	}

	if len(duplicationTypeSenderID) > 0 {
		if _, found := opr.Duplicated[duplicationTypeSenderID]; found {
			return errors.Errorf("proposal cannot have duplicated sender, %v", duplicationTypeSenderID)
		}

		opr.Duplicated[duplicationTypeSenderID] = struct{}{}
	}

	if len(duplicationTypeCurrencyID) > 0 {
		if _, found := opr.Duplicated[duplicationTypeCurrencyID]; found {
			return errors.Errorf(
				"cannot register duplicated currency id, %v within a proposal",
				duplicationTypeCurrencyID,
			)
		}

		opr.Duplicated[duplicationTypeCurrencyID] = struct{}{}
	}
	if len(duplicationTypeContractID) > 0 {
		if _, found := opr.Duplicated[duplicationTypeContractID]; found {
			return errors.Errorf(
				"cannot use a duplicated contract for registering in contract model , %v within a proposal",
				duplicationTypeSenderID,
			)
		}

		opr.Duplicated[duplicationTypeContractID] = struct{}{}
	}
	if len(duplicationTypePrescription) > 0 {
		if _, found := opr.Duplicated[duplicationTypePrescription]; found {
			return errors.Errorf(
				"cannot use a duplicated contract-hash for prescription info, %v within a proposal",
				duplicationTypePrescription,
			)
		}

		opr.Duplicated[duplicationTypePrescription] = struct{}{}
	}

	if len(newAddresses) > 0 {
		if err := opr.CheckNewAddressDuplication(newAddresses); err != nil {
			return err
		}
	}

	return nil
}

func GetNewProcessor(opr *currencyprocessor.OperationProcessor, op mitumbase.Operation) (mitumbase.OperationProcessor, bool, error) {
	switch i, err := opr.GetNewProcessorFromHintset(op); {
	case err != nil:
		return nil, false, err
	case i != nil:
		return i, true, nil
	}

	switch t := op.(type) {
	case currency.CreateAccount,
		currency.UpdateKey,
		currency.Transfer,
		extension.CreateContractAccount,
		extension.Withdraw,
		currency.RegisterCurrency,
		currency.UpdateCurrency,
		currency.Mint,
		prescription.RegisterModel,
		prescription.RegisterPrescription,
		prescription.UsePrescription:
		return nil, false, errors.Errorf("%T needs SetProcessor", t)
	default:
		return nil, false, nil
	}
}
