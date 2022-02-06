/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nest-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var fSet token.FileSet
		absDirPath, err := filepath.Abs("./")
		if err != nil {
			panic(fmt.Errorf("abs: %w",err))
		}

		parsedPackages, err := parser.ParseDir(&fSet, absDirPath, nil, parser.ParseComments)
		if err != nil {
			panic(fmt.Errorf("parse: %w",err))
		}

		for _, v := range parsedPackages {
			for _, file := range v.Files {
				ast.Inspect(file, func(n ast.Node) bool {
					switch x := n.(type) {
					case *ast.FuncDecl:
						fmt.Println(x.Name)
						fmt.Println("results:")
						fmt.Printf("%+v", x.Doc.Text())
						fmt.Println(x.Name.Name)
						fmt.Println(file.Name.Name)
					}
					return true
				})
			}
		}

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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nest-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


