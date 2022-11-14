![example workflow](https://github.com/fpersson/gosensor/actions/workflows/go.yml/badge.svg)

# Usage

## Hardware Requirement

* Rpi [Zero W](https://www.electrokit.com/produkt/raspberry-pi-zero-w-board/)
* Temperature sensor [DS18B20](https://www.kjell.com/se/produkter/el-verktyg/utvecklingskit/arduino/tillbehor/temperatursensor-med-kabel-for-arduino-p87081)

## build
```bash
    cd sensor
    go install
```

### build for rpi-zero
```bash
    cd sensor
    GOOS=linux GOARCH=arm go build -o sensor main.go
```

### build for rpi 3/4
```bash
    cd sensor
    GOOS=linux GOARCH=arm64 go build -o sensor main.go
```

## run

### Configurations
Default path for config files is $XDG_DATA_DIRS/tempsensor

```bash
    CONFIG=. ~/go/bin/sensor
```

## Grafana dashboard
```sql
    SELECT mean("last") FROM "sensor_1" WHERE ("unit" = 'temperature') AND $timeFilter GROUP BY time($__interval) fill(null)
```

## Hardware

### Setup DS18B20
Your DS18B20 should be connected to pin 7 (BCM4), gnd, and 3v.

Edit /boot/config.txt
```bash
# Enable gpio for DS18BS20
dtoverlay=w1-gpio,pinout=4,pullup=on
```

Edit /etc/modules
```bash
# /etc/modules: kernel modules to load at boot time.
#
# This file contains the names of kernel modules that should be loaded
# at boot time, one per line. Lines beginning with "#" are ignored.
w1-gpio pullup=1
w1-therm strong_pullup=1
```

check for /sys/bus/w1/devices/28-xxxxxxxxxx

## Note
This is the go version of [tempSensor](https://github.com/fpersson/tempSensor)

## Licens (Zero Clause BSD)
```
    Copyright (C) 2019, Fredrik Persson <fpersson.se@gmail.com>

    Permission to use, copy, modify, and/or distribute this software
    for any purpose with or without fee is hereby granted.

    THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
    WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
    AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
    INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
    LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR
    OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE
    OF THIS SOFTWARE.
```
