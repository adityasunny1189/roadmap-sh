package cachingproxy

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"sync"
)

var Addr string
var BaseUrl string

type HttpResponse struct {
	Data   interface{}
	ReqUrl string
}

type ProxyServer interface {
	Get(req string) (HttpResponse, bool)
	Cache(path string, res HttpResponse)
	Clear()
}

type ProxyStore struct {
	*sync.Mutex
	Store map[string]HttpResponse
}

func NewProxyServer() *ProxyStore {
	return &ProxyStore{
		Store: make(map[string]HttpResponse),
		Mutex: &sync.Mutex{},
	}
}

func (ps *ProxyStore) Get(req string) (HttpResponse, bool) {
	ps.Mutex.Lock()
	defer ps.Mutex.Unlock()
	if ps.Store[req].Data != nil {
		return ps.Store[req], true
	}
	return HttpResponse{
		Data:   nil,
		ReqUrl: req,
	}, false
}

func (ps *ProxyStore) Cache(path string, res HttpResponse) {
	ps.Mutex.Lock()
	defer ps.Mutex.Unlock()
	ps.Store[path] = res
}

func (ps *ProxyStore) Clear() {
	for key := range ps.Store {
		delete(ps.Store, key)
	}
}

type ProxyHandler struct {
	proxyServer ProxyServer
}

func NewProxyHandler(proxyServer ProxyServer) *ProxyHandler {
	return &ProxyHandler{
		proxyServer: proxyServer,
	}
}

func MainServerCall(path string) interface{} {
	reqUrl := BaseUrl + path
	res, err := http.Get(reqUrl)
	if err != nil {
		log.Println("Err: ", err)
		return nil
	}
	defer res.Body.Close()
	resBody := res.Body
	return resBody
}

func (h *ProxyHandler) HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.RawPath
	res, ok := h.proxyServer.Get(path)
	if !ok {
		w.Header().Set("X-Cache", "MISS")
		log.Println("X-Cache: MISS")
		// call main server
		response := MainServerCall(path)
		newRes := HttpResponse{
			Data:   response,
			ReqUrl: path,
		}
		// cache the response
		h.proxyServer.Cache(path, newRes)

		// return the response
		if err := json.NewEncoder(w).Encode(newRes); err != nil {
			log.Println("Err: ", err)
		}
		return
	}
	w.Header().Set("X-Cache", "HIT")
	log.Println("X-Cache: HIT")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println("Err: ", err)
	}
}

func ProxyServerStart() {
	port := flag.Int("port", 0, "Port number for the proxy server")
	origin := flag.String("origin", "", "Origin URL for the proxy server")
	cacheClear := flag.Bool("clear-cache", false, "Field to clear out cache")

	flag.Parse()

	if *port == 0 || *origin == "" {
		log.Println("Usage: caching-proxy --port <number> --origin <url>")
		os.Exit(1)
	}

	Addr = ":" + string(rune(*port))

	proxyServer := NewProxyServer()
	proxyHandler := NewProxyHandler(proxyServer)

	if *cacheClear {
		proxyHandler.proxyServer.Clear()
	}

	http.HandleFunc("/", proxyHandler.HandleHttpRequest)
	http.ListenAndServe(Addr, nil)
}
