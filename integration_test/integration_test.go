package integration_test

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/ed25519"
	"math/big"
	"os"
	"testing"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/utils"
)

func TestNewPrivatekey(t *testing.T) {
	_, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return
	}
	fmt.Println(hex.EncodeToString(priv))
}

func TestSpendTxWithNode(t *testing.T) {
	//sender := "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
	//senderPrivateKey := os.Getenv("INTEGRATION_TEST_SENDER_PRIVATE_KEY")
	senderPrivateKey := "52746669f457b197486ff2926fbd9b54bcd4071deaa13ba4b4fe5aa9028158427119a3b4d99985dd029610764564d833d35a3451f003b216b8b07387a3fb3c25"
	senderAccount, err := aeternity.AccountFromHexString(senderPrivateKey)
	//ak_rozWtRmHh91aEu1Qo46wGSHtJfaGtbgRgPEezZvTmHtRu1fqe
	sender := senderAccount.Address
	fmt.Printf("sender address: %s\n", sender)
	if err != nil {
		t.Fatal(err)
	}
	recipient := "ak_mPXUBSsSCJgfu3yz2i2AiVTtLA2TzMyMJL5e6X7shM9Qa246t"
	message := "Hello World"

	aeternity.Config.Node.URL = "http://47.106.255.174:10027"
	aeternity.Config.Node.NetworkID = "ae_mainnet"
	aeCli := aeternity.NewCli(aeternity.Config.Node.URL, true)

	// In case this test has been run before, get recipient's account info. If it exists, expectedAmount = amount + 10
	var expectedAmount big.Int
	recipientAccount, err := aeCli.APIGetAccount(recipient)
	if err != nil {
		expectedAmount.SetInt64(10)
	} else {
		expectedAmount.Add(recipientAccount.Balance.Int, big.NewInt(10))
		fmt.Printf("Recipient already exists with balance %v, expectedAmount after test is %s\n", recipientAccount.Balance.String(), expectedAmount.String())
	}

	amount := utils.NewBigInt()
	amount.SetInt64(896600000000000000)
	fee := utils.NewBigInt()
	fee.SetUint64(uint64(20000000000000))
	ttl, nonce, err := aeCli.GetTTLNonce(sender, aeternity.Config.Client.TTL)
	if err != nil {
		t.Fatalf("Error in GetTTLNonce(): %v", err)
	}

	// create the SpendTransaction
	tx := aeternity.NewSpendTx(sender, recipient, *amount, *fee, message, ttl, nonce)
	base64TxMsg, err := aeternity.BaseEncodeTx(&tx)
	if err != nil {
		t.Fatalf("Base64 encoding errored out: %v", err)
	}
	fmt.Println(base64TxMsg)

	// sign the transaction, output params for debugging
	signedBase64TxMsg, _, _, err := aeternity.SignEncodeTxStr(senderAccount, base64TxMsg, aeternity.Config.Node.NetworkID)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(signedBase64TxMsg)

	// send the signed transaction to the node
	err = aeCli.BroadcastTransaction(signedBase64TxMsg)
	if err != nil {
		t.Fatalf("Error while broadcasting transaction: %v", err)
	}

	// check the recipient's balance
	recipientAccount, err = aeCli.APIGetAccount(recipient)
	if err != nil {
		t.Fatalf("Couldn't get recipient's account data: %v", err)
	}

	if recipientAccount.Balance.Cmp(&expectedAmount) != 0 {
		t.Fatalf("Recipient should have %v, but has %v instead", expectedAmount.String(), recipientAccount.Balance.String())
	}
}

func TestSpendTxLargeWithNode(t *testing.T) {
	// This is a separate test because the account may not have enough funds for this test when the node has just started.
	sender := "ak_2a1j2Mk9YSmC1gioUq4PWRm3bsv887MbuRVwyv4KaUGoR1eiKi"
	senderPrivateKey := os.Getenv("INTEGRATION_TEST_SENDER_PRIVATE_KEY")
	senderAccount, err := aeternity.AccountFromHexString(senderPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	recipient := "ak_Egp9yVdpxmvAfQ7vsXGvpnyfNq71msbdUpkMNYGTeTe8kPL3v"
	message := "Hello World"

	aeternity.Config.Node.URL = "http://localhost:3013"
	aeternity.Config.Node.NetworkID = "ae_docker"
	aeCli := aeternity.NewCli(aeternity.Config.Node.URL, false)

	amount := utils.RequireBigIntFromString("18446744073709551615") // max uint64
	fee := utils.NewBigIntFromUint64(uint64(2e13))
	var expectedAmount = utils.NewBigInt()

	// In case the recipient account already has funds, get recipient's account info. If it exists, expectedAmount = existing balance + amount + fee
	recipientAccount, err := aeCli.APIGetAccount(recipient)
	if err != nil {
		expectedAmount.Set(amount.Int)
	} else {
		expectedAmount.Add(recipientAccount.Balance.Int, amount.Int)
	}

	ttl, nonce, err := aeCli.GetTTLNonce(sender, aeternity.Config.Client.TTL)
	if err != nil {
		t.Fatalf("Error in GetTTLNonce(): %v", err)
	}

	// create the SpendTransaction
	tx := aeternity.NewSpendTx(sender, recipient, *amount, *fee, message, ttl, nonce)
	base64TxMsg, err := aeternity.BaseEncodeTx(&tx)
	if err != nil {
		t.Fatalf("Base64 encoding errored out: %v", err)
	}
	fmt.Println(base64TxMsg)

	// sign the transaction, output params for debugging
	signedBase64TxMsg, _, _, err := aeternity.SignEncodeTxStr(senderAccount, base64TxMsg, aeternity.Config.Node.NetworkID)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(signedBase64TxMsg)

	// send the signed transaction to the node
	err = aeCli.BroadcastTransaction(signedBase64TxMsg)
	if err != nil {
		t.Fatalf("Error while broadcasting transaction: %v", err)
	}

	// check the recipient's balance
	recipientAccount, err = aeCli.APIGetAccount(recipient)
	if err != nil {
		t.Fatalf("Couldn't get recipient's account data: %v", err)
	}

	if recipientAccount.Balance.Cmp(expectedAmount.Int) != 0 {
		t.Fatalf("Recipient should have %v, but has %v instead", expectedAmount.String(), recipientAccount.Balance.String())
	}
}
