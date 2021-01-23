# halp

<p align="center">
    <img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/WonderWomanGopher.png" width="300" height="300">
</p>

`halp` is a command line utility that lets you display messages in morse code on your capslock LED, keyboard backlight (if you have it) or using your screen's brightness! 

### Note:
`halp` will only work if you're running Linux. The original purpose of developing this was to learn more about the concept of 'everyhting in linux is a file' and learn more about how permissions in linux work. The idea originated after watching [this](https://www.youtube.com/watch?v=Z56Jmr9Z34Q&feature=emb_title) video.

## Installation

Use the executable:
```
curl -sSL https://github.com/MadhavJivrajani/halp/releases/download/<tag>/halp --output halp
chmod +x halp
mv halp /usr/bin # to make the executable available system wide
```

Installation can also done using the `go get` command which will take care of installation of any libraries and dependencies nescessary. This will also install the `halp` executable which can be used anywhere in the termnial provided `$GOPATH/bin` is in your `PATH`.

```
go get -u github.com/MadhavJivrajani/halp
```

Or you can clone this repository and work directly with the `halp` executable!
```
git clone https://github.com/MadhavJivrajani/halp.git
cd halp
go build -o halp main.go
mv halp /usr/bin # to make the executable available system wide
```

## Usage:
```
A verryyy normal and usual application built to help you send SoS messages.

The tool will by default switch off the LED before displaying the morse code
message and will restore it back to this initial state when the message finishes displaying

Syntax:
halp -m <message>

Usage:
  halp [flags]
  halp [command]

Available Commands:
  help        Help about any command
  reset       Reset the resource state to 'value'

Flags:
      --config string    config file (default is $HOME/.halp.yaml)
  -h, --help             help for halp
  -k, --keyboard         use keyboard backlight
  -m, --message string   message to diplay in morse code
  -p, --path string      /path/to/capslockLED (default "/sys/class/leds/input3::capslock")
  -s, --screen           use screen backlight

Use "halp [command] --help" for more information about a command.
```

## Commands:
### `reset`
```
It will default to a capslock LED path, you can try using the -s and the -k flags,
to reset to value for screen and keyborad resources respectively

Syntax:
halp reset [--keyboard][--screen][--path][--value]

Usage:
  halp reset [flags]

Flags:
  -h, --help          help for reset
  -k, --keyboard      reset keyboard backlight
  -p, --path string   /path/to/resource (default "/sys/class/leds/input3::capslock")
  -s, --screen        reset screen backlight
  -v, --value int     reset the resource to this value
```

## Disclaimer

### I am not responsible for anything going wrong with the system `halp` is run on, this is an educational project.