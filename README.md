# Lmsensorsbeat

Welcome to Lmsensorsbeat.  This Beat strives to pull data from lm-sensors so that you can monitor a variety of I2C/SMBus sensors, such as CPU/motherboard temperatures, fan speeds, voltages, etc.  To run this, you may need to:

 1. Install the lm-sensors library and headers (e.g. libsensors4-dev) for your distribution
 2. Load (modprobe) various kernel modules for your sensors if you don't already.  Once you install lm-sensors, you can usually do this automatically by running the "sensors-detect" command

Ensure that this folder is at the following location:
`${GOPATH}/github.com/eskibars`

## Getting Started with Lmsensorsbeat

### Init Project
To get running with Lmsensorsbeat, run the following commands:

```
glide update --no-recursive
make update
```


To push Lmsensorsbeat in the git repository, run the following commands:

```
git init
git add .
git commit
git remote set-url origin https://github.com/eskibars/lmsensorsbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Lmsensorsbeat run the command below. This will generate a binary
in the same directory with the name lmsensorsbeat.

```
make
```


### Run

To run Lmsensorsbeat with debugging output enabled, run:

```
./lmsensorsbeat -c lmsensorsbeat.yml -e -d "*"
```


### Test

To test Lmsensorsbeat, run the following commands:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`


### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/lmsensorsbeat.template.json and etc/lmsensorsbeat.asciidoc

```
make update
```


### Cleanup

To clean  Lmsensorsbeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Lmsensorsbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/eskibars
cd ${GOPATH}/github.com/eskibars
git clone https://github.com/eskibars/lmsensorsbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).
