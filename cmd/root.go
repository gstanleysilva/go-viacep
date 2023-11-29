/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/guhstanley/go-viacep/internal/services"
	"github.com/spf13/cobra"
)

var cepString string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-viacep",
	Short: "Via CEP API",
	Long:  `Get Brazilian postal code coming from the ViaCEP API`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Loading...")

		srv, err := services.NewViaCepService()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}

		url, err := srv.GetJsonURL(cepString)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}

		viacep, err := srv.Execute(*url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}

		fmt.Println("ViaCEP Information:")
		fmt.Printf("CEP: %v \n", viacep.Cep)
		fmt.Printf("Localidade: %v \n", viacep.Localidade)
		fmt.Printf("UF: %v \n", viacep.Uf)
		fmt.Printf("Bairro: %v \n", viacep.Bairro)
		fmt.Printf("Logradouro: %v \n", viacep.Logradouro)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//Flags definition
	rootCmd.Flags().StringVarP(&cepString, "cep", "c", "", "CEP number")
	rootCmd.MarkFlagRequired("cep")
}
