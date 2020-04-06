# User_srv Service

This is the User_srv service

Generated with

```
micro new github.com/superryanguo/lightning/user_srv --namespace=micro.super.lightning --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: micro.super.lightning.srv.user_srv
- Type: srv
- Alias: user_srv

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
./user_srv-srv
```

Build a docker image
```
make docker
```