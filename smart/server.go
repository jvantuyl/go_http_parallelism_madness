package main

import (
	"net"
	"net/http"
)

type CleverServer struct {
    http.Server
}

func MakeCleverServer(addr string) (srv *CleverServer, mux *http.ServeMux) {
    mux = &http.ServeMux{}
    srv = &CleverServer{
        http.Server{
            Addr: addr,
            Handler: mux,
        },
    }
	return
}

func (srv *CleverServer) ListenAndServe() error {
    // Does not implement shutdown behavior due to private fields
	// if srv.shuttingDown() {
	//	return http.ErrServerClosed
	//}
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
    ln, err := net.Listen("tcp", addr)
    cln := MakeTCPListenerClever(ln.(*net.TCPListener))

	if err != nil {
		return err
	}
	return srv.Serve(cln)
}
