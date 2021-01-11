package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/MadhavJivrajani/halp/morse"
)

var cfgFile string

const (
	defaultLEDPath = "/sys/class/leds/input3::capslock/brightness"
	defaultMsg = ""
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "halp",
	Short: "A verryyy normal and usual application built to help you send SoS messages (is a joke)",
	Long: `
A verryyy normal and usual application built to help you send SoS messages (is a joke).

The tool will by default switch off the LED before displaying the morse code message and will restore
it back to this initial state when the message finishes displaying

Syntax:
halp -m <message>
`,

	RunE: func(cmd *cobra.Command, args []string) error {
		msg, _ := cmd.Flags().GetString("message")
		if len(msg) == 0 {
			return fmt.Errorf("PROVIDE MESSAGE TO SEND! YOUR LIFE MIGHT DEPEND ON IT!")
		}
		path, _ := cmd.Flags().GetString("path")
		return morse.SendSignal(path, msg)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.halp.yaml)")

	rootCmd.Flags().StringP(
		"message",
		"m",
		defaultMsg,
		"message to diplay in morse code",
	)

	rootCmd.Flags().StringP(
		"path",
		"p",
		defaultLEDPath,
		"/path/to/capslockLED",
	)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".halp" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".halp")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
