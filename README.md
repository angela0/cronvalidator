# cronvalidator

cronvalidator is a simple cron expression validator lib.

## Manual Installation

### Requirements

It requires Go 1.11 or later due to usage of Go Modules.

```
git clone https://github.com/angela0/cronvalidator && cd cronvalidator
make install
```

## Usage

```
cronnext
usage: cronnext <cron> <times>

example: cronnext "* * * * *" 10
```

```
cronvalidate
usage: cronvalidate <cron>

example: cronvalidate "* * * * *"
```


