// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/aeternity/aepp-sdk-go/aeternity"
	"github.com/aeternity/aepp-sdk-go/utils"

	"github.com/spf13/cobra"
)

var (
	waitForTx       bool
	payload         string
	printPrivateKey bool
	walletFileName  string
)

// walletCmd represents the wallet command
var walletCmd = &cobra.Command{
	Use:   "wallet PRIVATE_KEY_PATH",
	Short: "Interact with a wallet",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Printf("%#v", args)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside walletCmd PreRun with args: %v\n", args)
	},
}

// addressCmd represents the address subcommand
var addressCmd = &cobra.Command{
	Use:   "address WALLET_PATH",
	Short: "Print the aeternity wallet address",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// ask for th keystore password
		p, err := utils.AskPassword("Enter the password to unlock the keystore: ")
		if err != nil {
			fmt.Println("Error reading the password: ", err)
			os.Exit(1)
		}
		// load the account
		account, err := aeternity.LoadAccountToKeyStoreFile(args[0], p)
		if err != nil {
			fmt.Println("Error unlocking the keystore: ", err)
			os.Exit(1)
		}
		aeternity.Pp("Account address", account.Address)
		if printPrivateKey {
			aeternity.Pp("Account private key", account.SigningKeyToHexString())
		}
	},
}

// createCmd represents the generate subcommand
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new account",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		account, _ := aeternity.NewAccount()
		// ask for password
		p, err := utils.AskPassword("Enter a password for your keystore: ")
		if err != nil {
			fmt.Println("Error reading the password: ", err)
			return
		}
		// check if a name was given
		f, err := aeternity.StoreAccountToKeyStoreFile(account, p, walletFileName)
		if err != nil {
			fmt.Println("Error saving the keystore file: ", err)
			return
		}
		aeternity.Pp(
			"Wallet path", f,
			"Account address", account.Address,
		)
	},
}

// balanceCmd represents the generate subcommand
var balanceCmd = &cobra.Command{
	Use:   "balance WALLET_PATH",
	Short: "Get the balance of an account",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// ask for th keystore password
		p, err := utils.AskPassword("Enter the password to unlock the keystore: ")
		if err != nil {
			fmt.Println("Error reading the password: ", err)
			os.Exit(1)
		}
		// load the account
		account, err := aeternity.LoadAccountToKeyStoreFile(args[0], p)
		if err != nil {
			fmt.Println("Error unlocking the keystore: ", err)
			os.Exit(1)
		}
		a, err := aeCli.APIGetAccount(account.Address)
		if err != nil {
			fmt.Println("Error retrieving the account: ", err)
			os.Exit(1)
		}
		aeternity.PrintObject(a)
	},
}

// saveCmd represents the generate subcommand
var saveCmd = &cobra.Command{
	Use:   "save ACCOUNT_HEX_STRING",
	Short: "Save an account from a hex string to a keystore file",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		account, err := aeternity.AccountFromHexString(args[0])
		if err != nil {
			fmt.Println("Error parsing the private key hex string:", err)
			return
		}
		p, err := utils.AskPassword("Enter a password for your keystore: ")
		if err != nil {
			fmt.Println("Error reading the password: ", err)
			return
		}

		f, err := aeternity.StoreAccountToKeyStoreFile(account, p, walletFileName)
		if err != nil {
			fmt.Println("Error saving the keystore file: ", err)
			return
		}
		aeternity.Pp("Keystore path ", f)
	},
}

// spendCmd represents the spend subcommand
var spendCmd = &cobra.Command{
	Use:   "spend WALLET_PATH RECIPIENT_ADDRESS AMOUNT",
	Short: "Print the aeternity wallet spend",
	Long:  ``,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		var (
			keystorePath string
			recipient    string
			amount       int64
		)

		// load variables
		for _, a := range args {
			if strings.HasPrefix(a, string(aeternity.PrefixAccount)) {
				recipient = a
				continue
			}
			if m, _ := regexp.MatchString(`^\d+$`, a); m {
				amount, _ = strconv.ParseInt(a, 10, 64)
			}
			if p, err := aeternity.GetWalletPath(a); err == nil {
				keystorePath = p
			}
		}

		// validate variables
		if len(recipient) == 0 {
			fmt.Println("Error, missing or invalid recipient address")
			os.Exit(1)
		}
		if len(keystorePath) == 0 {
			fmt.Println("Error, missing or invalid keystore path")
			os.Exit(1)
		}
		if amount <= 0 {
			fmt.Println("Error, missing or invalid amount")
			os.Exit(1)
		}
		// ask for th keystore password
		p, err := utils.AskPassword("Enter the password to unlock the keystore: ")
		if err != nil {
			fmt.Println("Error reading the password: ", err)
			os.Exit(1)
		}
		// load the account
		account, err := aeternity.LoadAccountToKeyStoreFile(keystorePath, p)
		if err != nil {
			fmt.Println("Error unlocking the keystore: ", err)
			os.Exit(1)
		}
		// run the transaction
		_, txHash, _, ttl, _, err := aeCli.WithAccount(account).Spend(recipient, amount, payload)

		// TODO: print also the ttl
		aeternity.Pp(
			"Sender Address", account.Address,
			"Recipient Address", recipient,
			"Amount", amount,
			"TransactionHash", txHash,
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if waitForTx {
			_, _, _, tx, err := aeCli.WaitForTransactionUntillHeight(ttl, txHash)
			aeternity.PrintObjectT("Transaction", tx)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(walletCmd)
	walletCmd.AddCommand(addressCmd)
	walletCmd.AddCommand(spendCmd)
	walletCmd.AddCommand(createCmd)
	walletCmd.AddCommand(saveCmd)
	walletCmd.AddCommand(balanceCmd)

	// create flags
	createCmd.Flags().StringVar(&walletFileName, "name", "", "Override the default name of a wallaet")
	// save flags
	saveCmd.Flags().StringVar(&walletFileName, "name", "", "Override the default name of a wallaet")
	// address flags
	addressCmd.Flags().BoolVar(&printPrivateKey, "private-key", false, "Print the private key as hex string")
	// spend command flags
	spendCmd.Flags().BoolVarP(&waitForTx, "wait", "w", false, "Wait for the transaction to be mined before exiting")
	spendCmd.Flags().StringVarP(&payload, "message", "m", "", "Payload to add to the spend transaction")
}
