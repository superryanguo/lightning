# Session_mgr Service

This is the Session_mgr service

Generated with

```
micro new github.com/superryanguo/lightning/session_mgr --namespace=micro.super.lightning --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: micro.super.lightning.srv.session_mgr
- Type: srv
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
./session_mgr-srv
```

Build a docker image
```
make docker
```