# mp707_exporter

Prometheus exporter for MP707 USB thermometer.
See [nradchenko/mp707](https://github.com/nradchenko/mp707) project for more details.

## Build

Run `make` to compile `mp707_exporter` binary. Copy the binary to one of directories listed in your `$PATH`
environment variable.

## Usage

Launching `mp707_exporter` without arguments is fine, however consider using the following command line flags:
* `--listen-address`, `-l`: address to listen on (default is `:9774`)
* `--sensors-description`, `-d`: temperature sensor description which is filled in metric `description` label
  in `rom=text` format, e.g. `1800000c911b9228='Room #1'`. The flag can be provided multiple times.
* `--verbose`, `-v`: enable verbose logging (useful for debugging)

## Example metrics output
```
# HELP mp707_sensor_temp_celsius Sensor temperature
# TYPE mp707_sensor_temp_celsius gauge
mp707_sensor_temp_celsius{address="1800000c911b9228",description="Room #1"} 19.375
mp707_sensor_temp_celsius{address="7a00000c90526528",description="Room #2"} 19.25
mp707_sensor_temp_celsius{address="ea011933877d9828",description="Outdoor"} 7.625
```