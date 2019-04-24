package main

import (
    "io"
    "log"
    "net/http"
    "time"
    "sync"
)

func main() {
    var num int
    var num_mtx sync.Mutex

    srv, mux := MakeCleverServer(":3000")

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var id int;
        func() {
            num_mtx.Lock()
            defer num_mtx.Unlock()
            id = num
            num += 1
        }()

        log.Println("enter", id)

        log.Println("start", id)

        res, err := getExpensiveResource()
        if err != nil {
            http.Error(w, "failed to get resource", http.StatusInternalServerError)
            return
        }
        io.WriteString(w, res.String())

        log.Println("finish", id)
    })

    log.Println("starting...")
    srv.ListenAndServe()
}

type resource struct {}

func (r *resource) String() string {
    return "ping"
}

func getExpensiveResource() (r resource, err error) {
    time.Sleep(5 * time.Second)
    return resource{}, nil
}
