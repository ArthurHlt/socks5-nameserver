package main

import (
	"context"
	"net"
)

// CustomDNSResolver uses the system DNS to resolve host names
type CustomDNSResolver struct {
	dial func(ctx context.Context, network, address string) (net.Conn, error)
}

func factoryDial(nameserver string, inTcp bool) func(ctx context.Context, network, address string) (net.Conn, error) {
	protocol := "udp"
	if inTcp {
		protocol = "tcp"
	}
	return func(ctx context.Context, network, address string) (net.Conn, error) {
		d := net.Dialer{}
		return d.DialContext(ctx, protocol, nameserver)
	}
}

func NewCustomDNSResolver(nameserver string, inTcp bool) CustomDNSResolver {
	return CustomDNSResolver{
		dial: factoryDial(nameserver, inTcp),
	}
}

func (d CustomDNSResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	r := net.Resolver{
		PreferGo: true,
		Dial:     d.dial,
	}
	addr, err := r.LookupIPAddr(ctx, name)
	if err != nil {
		return ctx, nil, err
	}
	if len(addr) == 0 {
		return ctx, []byte{}, err
	}
	return ctx, addr[0].IP, err
}
