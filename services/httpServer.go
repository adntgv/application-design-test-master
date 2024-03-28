package services

import "net/http"

type HTTPServer struct {
	host string
	port string
	mux  *http.ServeMux
}

func NewHTTPServer(host, port string) *HTTPServer {
	mux := http.NewServeMux()
	return &HTTPServer{
		mux:  mux,
		host: host,
		port: port,
	}
}

func (s *HTTPServer) Register(app Registerable) {
	for route, handler := range app.GetRouteHandlers() {
		s.mux.HandleFunc(route, handler)
	}
}

func (s *HTTPServer) Serve() error {
	return http.ListenAndServe(s.host+":"+s.port, s.mux)
}

type Registerable interface {
	GetRouteHandlers() map[string]http.HandlerFunc
}
