package main

import (
	"net"
	"time"
)

type CleverTCPConn struct{
    innerConn *net.Conn
    sema chan struct{}
}

func (c CleverTCPConn) Close() error {
    defer func() {
        <-c.sema
    }()
    return (*(c.innerConn)).Close()
}

func (c CleverTCPConn) LocalAddr() net.Addr {
    return (*(c.innerConn)).LocalAddr()
}

func (c CleverTCPConn) RemoteAddr() net.Addr {
    return (*(c.innerConn)).RemoteAddr()
}

func (c CleverTCPConn) SetDeadline(t time.Time) error {
    return (*(c.innerConn)).SetDeadline(t)
}

func (c CleverTCPConn) SetReadDeadline(t time.Time) error {
    return (*(c.innerConn)).SetReadDeadline(t)
}

func (c CleverTCPConn) SetWriteDeadline(t time.Time) error {
    return (*(c.innerConn)).SetWriteDeadline(t)
}

func (c CleverTCPConn) Read(p []byte) (n int, err error) {
    return (*(c.innerConn)).Read(p)
}

func (c CleverTCPConn) Write(b []byte) (int, error) {
    return (*(c.innerConn)).Write(b)
}
