package integrationtest

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/aeternity/aepp-sdk-go/aeternity"
)

var sender = "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
var senderPrivateKey = os.Getenv("INTEGRATION_TEST_SENDER_PRIVATE_KEY")
var recipientPrivateKey = os.Getenv("INTEGRATION_TEST_RECEIVER_PRIVATE_KEY")
var nodeURL = "http://localhost:3013"
var networkID = "ae_docker"

func setupNetwork(t *testing.T) *aeternity.Client {
	aeternity.Config.Node.NetworkID = networkID
	client := aeternity.NewClient(nodeURL, false)
	t.Logf("nodeURL: %s, networkID: %s", nodeURL, aeternity.Config.Node.NetworkID)
	return client
}

func setupAccounts(t *testing.T) (*aeternity.Account, *aeternity.Account) {
	alice, err := aeternity.AccountFromHexString(senderPrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	bob, err := aeternity.AccountFromHexString(recipientPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Alice: %s, Bob: %s", alice.Address, bob.Address)
	return alice, bob
}

func signBroadcast(t *testing.T, tx aeternity.Tx, acc *aeternity.Account, aeClient *aeternity.Client) (hash string) {
	txB64, _ := aeternity.BaseEncodeTx(tx)
	// t.Log(txB64)

	signedTxStr, hash, _, err := aeternity.SignEncodeTxStr(acc, txB64, aeternity.Config.Node.NetworkID)
	if err != nil {
		t.Fatal(err)
	}

	err = aeClient.BroadcastTransaction(signedTxStr)
	if err != nil {
		t.Fatal(err)
	}

	return hash

}

type delayableCode func()

func delay(f delayableCode) {
	time.Sleep(2000 * time.Millisecond)
	f()
}

func getHeight(aeClient *aeternity.Client) (h uint64) {
	h, err := aeClient.APIGetHeight()
	if err != nil {
		fmt.Println("Could not retrieve chain height")
		return
	}
	// fmt.Println("Current Height:", h)
	return
}

func waitForTransaction(aeClient *aeternity.Client, hash string) (err error) {
	height := getHeight(aeClient)
	// fmt.Println("Waiting for", hash)
	height, _, _, _, err = aeClient.WaitForTransactionUntilHeight(height+10, hash)
	if err != nil {
		// Sometimes, the tests want the tx to fail. Return the err to let them know.
		return err
	}
	// fmt.Println("Transaction was found at", height, "blockhash", blockHash, "microBlockHash", microBlockHash, "err", err)
	return nil
}