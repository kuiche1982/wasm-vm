package vmjsonintegrationtest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRustAdder(t *testing.T) {
	runAllTestsInFolder(t, "adder/mandos")
}

func TestRustErc20(t *testing.T) {
	runAllTestsInFolder(t, "erc20-rust/mandos")
}

func TestCErc20(t *testing.T) {
	runAllTestsInFolder(t, "erc20-c")
}

func TestMultisig(t *testing.T) {
	runAllTestsInFolder(t, "multisig/mandos")
}

func TestDnsContract(t *testing.T) {
	if testing.Short() {
		t.Skip("not a short test")
	}

	runAllTestsInFolder(t, "dns")
}

func TestDexLogs(t *testing.T) {
	err := runSingleTestReturnError("dex/mandos", "enter_farm.scen.json")
	require.Nil(t, err)
}

func TestCrowdfundingEsdt(t *testing.T) {
	runAllTestsInFolder(t, "crowdfunding-esdt")
}

func TestEgldEsdtSwap(t *testing.T) {
	runAllTestsInFolder(t, "egld-esdt-swap")
}

func TestPingPongEgld(t *testing.T) {
	runAllTestsInFolder(t, "ping-pong-egld")
}

func TestRustAttestation(t *testing.T) {
	if testing.Short() {
		t.Skip("not a short test")
	}

	runAllTestsInFolder(t, "attestation-rust")
}

func TestCAttestation(t *testing.T) {
	if testing.Short() {
		t.Skip("not a short test")
	}
	runAllTestsInFolder(t, "attestation-c")
}
