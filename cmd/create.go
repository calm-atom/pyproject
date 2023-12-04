/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// Logo generated from https://www.asciiart.eu/text-to-ascii-art
const logo = `
 ______   __  __     ______   ______     ______       __     ______     ______     ______ 
/\  == \ /\ \_\ \   /\  == \ /\  == \   /\  __ \     /\ \   /\  ___\   /\  ___\   /\__  _\
\ \  _-/ \ \____ \  \ \  _-/ \ \  __<   \ \ \/\ \   _\_\ \  \ \  __\   \ \ \____  \/_/\ \/
 \ \_\    \/\_____\  \ \_\    \ \_\ \_\  \ \_____\ /\_____\  \ \_____\  \ \_____\    \ \_\
  \/_/     \/_____/   \/_/     \/_/ /_/   \/_____/ \/_____/   \/_____/   \/_____/     \/_/
`

var logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0097")).Bold(true)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", logoStyle.Render(logo))

		installed, version := checkIfPythonInstalled()
		if installed == false {
			log.Fatal("Python not installed!")
		}
		fmt.Printf("%s", version)

		fmt.Println("What is the name of your project? ")
		var input string
		fmt.Scanln(&input)
		fmt.Printf("You entered %s\n", input)
		if checkIfProjectExists(input) {
			log.Fatal("Project already exists!")
		}

		err := os.MkdirAll(input, 0751)
		if err != nil {
			panic("Error creating directory!")
		}

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// High level steps:
// Check if python is installed
// Ask for project name
// Ask for make file or pyproject.toml

func checkIfPythonInstalled() (bool, string) {
	out, err := exec.Command("python", "--version").Output()
	if err != nil {
		out, err = exec.Command("python3", "--version").Output()
		if err != nil {
			log.Println(err)
			return false, "not installed"
		}
	}

	return true, string(out)
}

func checkIfProjectExists(name string) bool {
	if _, err := os.Stat(name); err == nil {
		dirEntries, err := os.ReadDir(name)
		if err != nil {
			log.Printf("Could not read directory %v", err)
		}
		if len(dirEntries) > 0 {
			return true
		}
	}
	return false
}
