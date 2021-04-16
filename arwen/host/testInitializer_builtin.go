package host

import (
	"math/big"

	"github.com/ElrondNetwork/elrond-go/core/vmcommon"
	"github.com/ElrondNetwork/elrond-go/data/state"
	"github.com/ElrondNetwork/elrond-go/process"
)

type MockClaimBuiltin struct {
	AmountToGive *big.Int
	GasCost      uint64
}

func (m *MockClaimBuiltin) ProcessBuiltinFunction(acntSnd, _ state.UserAccountHandler, vmInput *vmcommon.ContractCallInput) (*vmcommon.VMOutput, error) {
	vmOutput := MakeVMOutput()
	AddNewOutputAccount(
		vmOutput,
		acntSnd.AddressBytes(),
		nil,
		42,
		nil)

	vmOutput.GasRemaining = vmInput.GasProvided - m.GasCost + vmInput.GasLocked
	return vmOutput, nil
}

func (m *MockClaimBuiltin) SetNewGasConfig(_ *process.GasCost) {
}

func (m *MockClaimBuiltin) IsInterfaceNil() bool {
	return m == nil
}