/*
Copyright © 2022 Philippe Moneda <philippe.moneda@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/PMoneda/gcoin/wallet"
	"github.com/spf13/cobra"
)

// walletCmd represents the wallet command
var walletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Wallet cli to gcoin",
	Long: `Wallet cli to gcoin
	* Create new wallet
	* Check funds of wallet
	* Transfer money to different wallets`,
	Run: func(cmd *cobra.Command, args []string) {
		for i, k := range args {
			fmt.Printf("%d, %s\n", i, k)
		}
		fmt.Println()
		if len(args) > 0 && args[0] == "create" {
			w := wallet.CreateNewWallet()
			fmt.Printf("Private Key: %s\nPublic Key: %s\n\n*** Keep your private key secure ***\n\n", w.PrivateKey, w.PublicKey)
		} else if len(args) >= 2 && args[0] == "signin" {
			wallet.SignIn(args[1])
		} else if args[0] == "add" {
			alias := cmd.Flag("alias").Value.String()
			pk := cmd.Flag("pubkey").Value.String()
			fmt.Printf("saving pubkey %s to alias %s\n", pk, alias)
			err := wallet.AddTrustedContact(pk, alias)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("contacted added")
			}
		} else if args[0] == "list" {
			list, err := wallet.ListAllContacts()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			for _, k := range list {
				fmt.Printf("%s\n", k)
			}
		} else if args[0] == "status" {
			status, err := wallet.Status()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("Public Key: %s", status.PublicKey)
			}
		} else if args[0] == "delete" {
			alias := cmd.Flag("alias").Value.String()
			err := wallet.DeleteAliasContact(alias)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("contact deleted")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(walletCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// walletCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	walletCmd.Flags().String("alias", "", "Contact's alias")
	walletCmd.Flags().String("pubkey", "", "Public Key from contact")
}
