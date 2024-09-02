package cmds

type PrescriptionCommand struct {
	RegisterPrescription RegisterPrescriptionCommand `cmd:"" name:"register-prescription" help:"register prescription"`
	UsePrescription      UsePrescriptionCommand      `cmd:"" name:"use-prescription" help:"use prescription"`
	RegisterModel        RegisterModelCommand        `cmd:"" name:"register-model" help:"register prescription model"`
}
