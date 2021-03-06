# Session_mgr Service

This is the Session_mgr service

Generated with

```
micro new session_mgr --namespace=micro.super.lightning --type=service
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: micro.super.lightning.service.session_mgr
- Type: service
- Alias: session_mgr

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./session_mgr-service
```

Build a docker image
```
make docker
```