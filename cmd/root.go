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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-viacep",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		srv, err := services.NewViaCepService()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		url := srv.GetJsonURL("30575460")
		viacep, err := srv.Execute(url)
		if err != nil {
			fmt.Println(err.Error())
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-viacep.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
