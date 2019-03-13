# socks5-nameserver [![Build Status](https://travis-ci.com/ArthurHlt/socks5-nameserver.svg?branch=master)](https://travis-ci.com/ArthurHlt/socks5-nameserver)

A socks5 server to change dns nameserver.

## install

### On *nix system

You can install this via the command-line with either `curl` or `wget`.

#### via curl

```bash
$ sh -c "$(curl -fsSL https://raw.github.com/ArthurHlt/socks5-nameserver/master/bin/install.sh)"
```

#### via wget

```bash
$ sh -c "$(wget https://raw.github.com/ArthurHlt/socks5-nameserver/master/bin/install.sh -O -)"
```

### On windows

You can install it by downloading the `.exe` corresponding to your cpu from releases page: https://github.com/ArthurHlt/socks5-nameserver/releases .
Alternatively, if you have terminal interpreting shell you can also use command line script above, it will download file in your current working dir.

### From go command line

Simply run in terminal:

```bash
$ go get github.com/socks5-nameserver/notifslack
```

## Usage

```
NAME:
   socks5-nameserver - socks5-nameserver change your nameserver and resolve on sock5

USAGE:
   socks5-nameserver [global options] command [command options] [arguments...]

VERSION:
   dev

COMMANDS:
     start    launch socks5-nameserver
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --host value                  Set host for listen valid are 127.0.0.1 and 0.0.0.0 (default: "0.0.0.0")
   --port value, -p value        Listen port (default: 8000)
   --nameserver value, -n value  Set dns nameserver
   --in-tcp                      Set dns nameserver to be called in tcp instead of udp
   --help, -h                    show help
   --version, -v                 print the version
```