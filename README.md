# CIDR Info

Command line tool for printing the CIDR information. Example:

```
$ cidrinfo 10.128.0.0/9

IP:          10.128.0.0
Network:     10.128.0.0/9
Netmask:     255.128.0.0
CIDR Range:  10.128.0.0  <-to->  10.255.255.255
Count:       8,388,608
```

## Install

```
go install github.com/wiztools/cidrinfo
```
