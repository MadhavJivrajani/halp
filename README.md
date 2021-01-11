# halp

<p align="center">
    <img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/WonderWomanGopher.png" width="300" height="300">
</p>

`halp` is a command line utility that lets you display messages in morse code on your capslock LED, or any LED for that matter as long the appearance of the states of the LEDs are distinct for values of `0` and `1`. 

### Note:
`halp` will only work if you're running Linux. The original purpose of developing this was to learn more about the concept of 'everyhting in linux is a file' and learn more about how permissions in linux work. The idea originated after watching [this](https://www.youtube.com/watch?v=Z56Jmr9Z34Q&feature=emb_title) video.

## Installation
Installation can be done using the `go get` command which will take care of installation of any libraries and dependencies nescessary. This will also install the `halp` executable which can be used anywhere in the termnial provided `$GOPATH/bin` is in your `PATH`.

```
go get -u github.com/MadhavJivrajani/halp
```

Or you can clone this repository and work directly with the `halp` executable!
```
git clone https://github.com/MadhavJivrajani/halp.git
cd halp
go build -o halp main.go
chmod +x halp
mv halp /usr/bin # to make the executable available system wide
```

## Usage:
```
A verryyy normal and usual application built to help you send SoS messages.

The tool will by default switch off the LED before displaying the morse code message and will restore
it back to this initial state when the message finishes displaying

Syntax:
halp -m <message>

Usage:
  halp [flags]
  halp [command]

Available Commands:
  help        Help about any command
  reset       Reset the capslock LED to an OFF state

Flags:
      --config string    config file (default is $HOME/.halp.yaml)
  -h, --help             help for halp
  -m, --message string   message to diplay in morse code
  -p, --path string      /path/to/capslockLED (default "/sys/class/leds/input3::capslock/brightness")

Use "halp [command] --help" for more information about a command.
```

## Commands:
```
halp -m <message> [--path] # display message in morse code on your capslock LED
```

In case the program is stopped mid-way while displaying the morse code or it crashes due to whatever reason, the state of the capslock LED can be reset to `off` by running the following:
```
halp reset [--path]
```

## Disclaimer

### I am not responsible for anything going wrong with the system `halp` is run on, this is an educational project.