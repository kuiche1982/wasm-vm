package contexts

import (
	"github.com/ElrondNetwork/arwen-wasm-vm/arwen"
	"github.com/ElrondNetwork/arwen-wasm-vm/config"
	vmcommon "github.com/ElrondNetwork/elrond-vm-common"
)

type meteringContext struct {
	gasSchedule           *config.GasCost
	blockGasLimit         uint64
	gasLockedForAsyncStep uint64
	host                  arwen.VMHost
}

func NewMeteringContext(
	host arwen.VMHost,
	gasSchedule map[string]map[string]uint64,
	blockGasLimit uint64,
) (*meteringContext, error) {

	gasCostConfig, err := config.CreateGasConfig(gasSchedule)
	if err != nil {
		return nil, err
	}

	context := &meteringContext{
		gasSchedule:           gasCostConfig,
		blockGasLimit:         blockGasLimit,
		gasLockedForAsyncStep: 0,
		host:                  host,
	}

	return context, nil
}

func (context *meteringContext) GasSchedule() *config.GasCost {
	return context.gasSchedule
}

func (context *meteringContext) UseGas(gas uint64) {
	gasUsed := context.host.Runtime().GetPointsUsed() + gas
	context.host.Runtime().SetPointsUsed(gasUsed)
}

func (context *meteringContext) FreeGas(gas uint64) {
	refund := context.host.Output().GetRefund() + gas
	context.host.Output().SetRefund(refund)
}

func (context *meteringContext) GasLeft() uint64 {
	gasProvided := context.host.Runtime().GetVMInput().GasProvided
	gasUsed := context.host.Runtime().GetPointsUsed()
	return gasProvided - gasUsed
}

func (context *meteringContext) BoundGasLimit(value int64) uint64 {
	gasLeft := context.GasLeft()
	limit := uint64(value)

	if gasLeft < limit {
		return gasLeft
	} else {
		return limit
	}
}

// deductAndLockGasIfAsyncStep will deduct the gas for an async step and also lock gas for the callback, if the execution is an asynchronous call
func (context *meteringContext) deductAndLockGasIfAsyncStep() error {
	input := context.host.Runtime().GetVMInput()

	if input.CallType == vmcommon.AsynchronousCall {
		gasSchedule := context.GasSchedule().ElrondAPICost

		gasToConsume := gasSchedule.AsyncCallStep
		gasToLock := gasSchedule.AsyncCallStep + gasSchedule.AsyncCallbackGasLock
		if input.GasProvided <= gasToConsume+gasToLock {
			return arwen.ErrNotEnoughGas
		}
		input.GasProvided -= gasToConsume + gasToLock

		context.gasLockedForAsyncStep = gasToLock
	} else {
		context.gasLockedForAsyncStep = 0
	}

	return nil
}

// UnlockGasIfAsyncStep will restore the previously locked gas, if the execution is an asynchronous call
func (context *meteringContext) UnlockGasIfAsyncStep() {
	if context.gasLockedForAsyncStep != 0 {
		input := context.host.Runtime().GetVMInput()
		input.GasProvided += context.gasLockedForAsyncStep
		context.gasLockedForAsyncStep = 0
	}
}

func (context *meteringContext) BlockGasLimit() uint64 {
	return context.blockGasLimit
}

func (context *meteringContext) DeductInitialGasForExecution(contract []byte) error {
	err := context.deductInitialGas(
		contract,
		0,
		context.gasSchedule.BaseOperationCost.CompilePerByte,
	)
	if err != nil {
		return err
	}

	err = context.deductAndLockGasIfAsyncStep()

	return err
}

func (context *meteringContext) DeductInitialGasForDirectDeployment(input *vmcommon.ContractCreateInput) error {
	err := context.deductInitialGas(
		input.ContractCode,
		context.gasSchedule.ElrondAPICost.CreateContract,
		context.gasSchedule.BaseOperationCost.StorePerByte,
	)
	return err
}

func (context *meteringContext) DeductInitialGasForIndirectDeployment(input *vmcommon.ContractCreateInput) error {
	err := context.deductInitialGas(
		input.ContractCode,
		0,
		context.gasSchedule.BaseOperationCost.StorePerByte,
	)
	return err
}

func (context *meteringContext) deductInitialGas(
	code []byte,
	baseCost uint64,
	costPerByte uint64,
) error {
	input := context.host.Runtime().GetVMInput()
	codeLength := uint64(len(code))
	codeCost := codeLength * costPerByte
	initialCost := baseCost + codeCost

	if initialCost > input.GasProvided {
		return arwen.ErrNotEnoughGas
	}

	input.GasProvided -= initialCost
	return nil
}
