# Center Service

This is the Center service

Generated with

```
micro new github.com/fidelfly/fxms/center --namespace=go.micro --fqdn=com.fxms.web.center --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: com.fxms.web.center
- Type: web
- Alias: center

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./center-web
```

Build a docker image
```
make docker
```