package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"errors"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/MadhavJivrajani/halp/morse"
	"regexp"
)

var cfgFile string

const (
	defaultLEDPath = "/sys/class/leds/input3::capslock"
	defaultMsg     = ""
	keyboardBacklightRegex = ".+::kbd_backlight"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "halp",
	Short: "A verryyy normal and usual application built to help you send SoS messages",
	Long: `
A verryyy normal and usual application built to help you send SoS messages.

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
		screen, _ := cmd.Flags().GetBool("screen")
		keyboard, _ := cmd.Flags().GetBool("keyboard")
				
		if keyboard == true && screen == true {			
			return fmt.Errorf("Use only one of keyboard or screen options")
		}

		if keyboard == true {
			var err error
			path, err = getKeyboardBacklightPath()
			if err != nil {
				return fmt.Errorf(err.Error())
			}			
		} else if screen == true {
			var err error
			path, err = getScreenBacklightPath()
			if err != nil {
				return fmt.Errorf(err.Error())
			}
		}

		return morse.SendSignal(path, msg)
	},
}

func getKeyboardBacklightPath() (string, error) {
	rootPath := "/sys/class/leds/"
	re := regexp.MustCompile(keyboardBacklightRegex)

	requiredDir := ""
	
	walk := func(fn string, fi os.FileInfo, err error) error {		
		if re.MatchString(fn) {
			requiredDir = fn
		}
		return nil
	}
	filepath.Walk(rootPath, walk)

	if requiredDir == "" {
		return "", errors.New("couldn't find the keyboard file")
	}
		
	return requiredDir, nil
}

func getScreenBacklightPath() (string, error) {
	return "/sys/class/backlight/intel_backlight", nil
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

	rootCmd.Flags().BoolP(
		"screen",
		"s",
		false,
		"use screen backlight",
	)

	rootCmd.Flags().BoolP(
		"keyboard",
		"k",
		false,
		"use keyboard backlight",
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
