package main

import (
	"net"
)

type CleverTCPListener struct{
    innerListener *net.TCPListener
    sema chan struct{}
}

func MakeTCPListenerClever(inner *net.TCPListener) (ctl *CleverTCPListener) {
    const maxClients = 15

    ctl = &CleverTCPListener{
        innerListener: inner,
        sema: make(chan struct{}, maxClients),
    }
    return
}

func (l *CleverTCPListener) Accept() (conn net.Conn, err error) {
    l.sema <- struct{}{}

    raw_conn, err := l.innerListener.Accept()
    conn = CleverTCPConn{
        innerConn: &raw_conn,
        sema: l.sema,
    }
    return
}

func (l *CleverTCPListener) Addr() (addr net.Addr) {
    addr = l.innerListener.Addr()
    return
}

func (l *CleverTCPListener) Close() (err error) {
    err = l.innerListener.Close()
    return
}
