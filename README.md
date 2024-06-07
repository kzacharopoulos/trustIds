# Capture

Captures traffic for real time or near real time analysis.

It targets low end edge devices that support linux and the dependencies
described in the next section.

It is developed around three interfaces:

- api  , which handles the analysis reporting
- trust, which handles the trust evaluation of communicating nodes and uses the 'api' for reporting
- kit  , which handles network metrics and uses 'trust' to reward/penalize node behavior

It supports:

- Sniffing traffic from network interface
  - *needs elevated rights*
- Reading pcap files
- The following 'api' backends:
  - amqp
  - log
- The following 'trust' backends:
  - thresholds
- The following 'kit' backends:
  - trustIds (based on packet rate and throughput)
  - log

To do:

- Add ability to save network traffic from interface to file

## Dependencies

GoLang Version
```
go1.19
```

The following are only needed if you want to cross compile. They do not
show all available architectures, only aarch64.

Example for debian based systems:

```
sudo apt install \
    qemu-user \
    qemu-user-static \
    gcc-aarch64-linux-gnu \
    binutils-aarch64-linux-gnu \
    binutils-aarch64-linux-gnu-dbg \
    build-essential
```

Examples for redhat based systems:

```
sudo dnf group install \
    "C Development Tools and Libraries" \
    "Development Tools"

sudo dnf install \
    qemu-user \
    qemu-user-static \
    gcc-aarch64-linux-gnu \
    binutils-aarch64-linux-gnu \
```

## Build

```
make
make setcap
make clean
```

```
make cross PLATFORM=debian ARCHITECTURE=386
make cross PLATFORM=debian ARCHITECTURE=arm/v5
make cross PLATFORM=debian ARCHITECTURE=arm/v7
make cross PLATFORM=debian ARCHITECTURE=arm64/v8
make cross PLATFORM=debian ARCHITECTURE=mips64le
make cross PLATFORM=debian ARCHITECTURE=ppc64le
make cross PLATFORM=debian ARCHITECTURE=s390x
```

```
make cross PLATFORM=alpine ARCHITECTURE=386
make cross PLATFORM=alpine ARCHITECTURE=arm/v6
make cross PLATFORM=alpine ARCHITECTURE=arm/v7
make cross PLATFORM=alpine ARCHITECTURE=arm64/v8
make cross PLATFORM=alpine ARCHITECTURE=ppc64le
make cross PLATFORM=alpine ARCHITECTURE=s390x
```

## Run

```
./capture
```

Available Commands:
- completion (Generate the autocompletion script for the specified shell)
- conf (Shows effective configuration)
- help (Help about any command)
- info (Shows release info and exits)
- open (Opens a packet source to sniff from)
- Flags:
    - -h, --help





# Links

https://www.netresec.com/?page=PcapFiles
https://xenanetworks.com/?knowledge-base=knowledge-base/valkyrie/downloads/pcap-samples
