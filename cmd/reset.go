package cmd

import (
	"fmt"

	"github.com/MadhavJivrajani/halp/morse"
	"github.com/spf13/cobra"
)

const (
	defaultValue = 0
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the resource state to 'value'",
	Long: `
It will default to a capslock LED path, you can try using the -s and the -k flags,
to reset to value for screen and keyborad resources respectively

Syntax:
halp reset [--keyboard][--screen][--path][--value]`,

	RunE: func(cmd *cobra.Command, args []string) error {
		path, _ := cmd.Flags().GetString("path")
		value, _ := cmd.Flags().GetInt("value")

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

		return morse.UpdateState(path, value)
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
	resetCmd.Flags().IntP(
		"value",
		"v",
		defaultValue,
		"reset the resource to this value",
	)

	resetCmd.Flags().StringP(
		"path",
		"p",
		defaultLEDPathRegex,
		"/path/to/resource",
	)

	resetCmd.Flags().BoolP(
		"screen",
		"s",
		false,
		"reset screen backlight",
	)

	resetCmd.Flags().BoolP(
		"keyboard",
		"k",
		false,
		"reset keyboard backlight",
	)
}
