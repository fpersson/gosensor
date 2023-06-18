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

With test device, the DEVICE is a fake device for testing
```bash
    CONFIG=./testdata DEVICE=./fejksensor/ go run ./sensor
```

With logging to file
```bash
    LOGDIR=./testlog.log CONFIG=./testdata DEVICE=./fejksensor/ go run ./sensor
```

## Grafana dashboard
```sql
    SELECT mean("last") FROM "sensor_1" WHERE ("unit" = 'temperature') AND $timeFilter GROUP BY time($__interval) fill(null)
```

## Health Check
```
    http://localhost:8081/health_check
```

## Hardware

### Setup DS18B20
Your DS18B20 should be connected to pin 7 (BCM4), gnd, and 3v.

Edit config.txt on raspbian it can be found in /boot nad opensuse it can be found in /boot/efi/
config.txt
```bash
# Enable gpio for DS18BS20
dtoverlay=w1-gpio,pinout=4,pullup=on
```

Reastart your system, check for /sys/bus/w1/devices/28-xxxxxxxxxx. If your device does not show up try:
```bash
    sudo modprobe w1-gpio
    sudo modprobe w1-therm
```

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
