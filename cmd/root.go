package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"crypto/sha256"
	"fmt"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shagen [text to hash]",
	Short: "Generates a SHA256 HASH for any string value or set of string values provided.",
	Long: `Generates a SHA256 HASH for any string value or set of string values provided. 
This program accepts at least one string argument. Single string values with spaces should be placed in quotes. 
Multiple strings to be hashed must be passed as arguments to the command seperated by a space.`,
	Run: func(cmd *cobra.Command, args []string) {
		//For each arg passed, generate a SHA256 Hash
		for _, arg := range args {
			hasher := sha256.New()
			hasher.Write([]byte(arg))
			sha := hasher.Sum(nil)
			fmt.Printf("Original Text: %s   \nHashed Text: %x \n--------\n", arg, sha)		}
	},
	//ensure that atleast 1 argument is passed by user
	Args: cobra.MinimumNArgs(1),
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

