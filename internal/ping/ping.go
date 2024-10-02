package ping

import "net/http"

type Ping struct {
}

func NewPing() *Ping { return &Ping{} }

func (p *Ping) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
