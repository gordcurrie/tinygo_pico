# TinyGo with Pico

Playing around with [TinyGo](https://tinygo.org/) and a Raspberry Pi [Pico](https://www.raspberrypi.com/products/raspberry-pi-pico/).

# Setup

# Prerequisites

TinyGo installed and configured - https://tinygo.org/getting-started/install/

If comfortable using Vim and [direnv](https://direnv.net/) for development install and use [TinyHelper](https://github.com/gordcurrie/tinyhelper) for configuring environment variables for easy use.

Otherwise https://tinygo.org/docs/guides/ide-integration/

# Debugging

Debugging output can be sent to serial via `println(val)` command.

Serial output can be viewed via screen.

Screen can be installed via the following command on Debian based systems.

```
sudo apt-get install screen
```

Once pico is flashed with program and connected via USB run the following command to view serial output.

```
screen /dev/ttyACM0
```

* `screen` is the command
* `/dev/ttyACM0` is the device port where the pico is connected, adjust if needed.

NOTE:

Exit screen with `ctrl + a :quit`

