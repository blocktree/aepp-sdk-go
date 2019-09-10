package cmd

import (
	"github.com/aeternity/aepp-sdk-go/aeternity"
	"testing"

	"github.com/spf13/cobra"
)

func TestTxSpend(t *testing.T) {
	sender := "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
	recipient := "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"
	emptyCmd := cobra.Command{}

	err := txSpendFunc(&emptyCmd, []string{sender, recipient, "10"})
	if err != nil {
		t.Error(err)
	}
}

func TestTxVerify(t *testing.T) {
	aeternity.Config.Node.NetworkID = "ae_mainnet"
	//sender := "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
	//signedTx := "tx_+JYLAfhCuEAkhq5DuTb5s67AwoOgto9eihfvCPZrDmgDYxYLZ7hggGhp7LvzS0KDebV24R4Xnijz1LgxRKVzel/36JoLH1AIuE74TAwBoQHOp63kcMn5nZ1OQAiAqG8dSbtES2LxGp67ZLvP63P+86EBHxOjsIvwAUAGYqaLadh194A87EwIZH9u1dhMeJe9UKMKCoGaAYBc/oor"
	//sender := "ak_qcqXt6ySgRPvBkNwEpNMvaKWzrhPZsoBHLvgg68qg9vRht62y"
	//signedTx := "tx_+KULAfhCuEADxqJQ/35t5VIwLV5mSYwsFuBzQzrFZLrjsbWuXe6/ZDG62Lsa+WzypvtaRs45FMWIEYrx+wXY3LO/1bgpQ70MuF34WwwBoQFuZJC6n/o+0nYEjiPFLwmnYi4CEREk6cdw0aasEacjxqEBbuunhRwt203U1HWI8eGiBPiNbg7tMCbK87jsPkMvnYmHI4byb8EAAIYSMJzlQACDAeBOJoCRHPk5"
	sender := "ak_qrKKpCkCYej3MakQqgxZZNCfEwV1T6vYvKuEUr9j3hDY3aXYK"
	signedTx := "tx_+KULAfhCuEBtOYcTNDCmW4NAUuPKIigspbd0HKRtilwpUziouJDjQLH+pFc664947E5Wzo6wMdgQoC8Dw/PBGr6eh8yaDp8YuF34WwwBoQFu66eFHC3bTdTUdYjx4aIE+I1uDu0wJsrzuOw+Qy+diaEBbmSQup/6PtJ2BI4jxS8Jp2IuAhERJOnHcNGmrBGnI8aHI4byb8EAAIYSMJzlQACDAd/LAYBSwpYl"
	emptyCmd := cobra.Command{}

	err := txVerifyFunc(&emptyCmd, []string{sender, signedTx})
	if err != nil {
		t.Error(err)
	}
}

func TestTxDumpRaw(t *testing.T) {
	tx := "tx_+H4iAaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMFoQJei0cGdDWb3EfrY0mtZADF0LoQ4yL6z10I/3ETJ0fpKADy8Y5hY2NvdW50X3B1YmtleaEBzqet5HDJ+Z2dTkAIgKhvHUm7REti8Rqeu2S7z+tz/vMGAQXLBNnv"
	emptyCmd := cobra.Command{}

	err := txDumpRawFunc(&emptyCmd, []string{tx})
	if err != nil {
		t.Error(err)
	}
}
