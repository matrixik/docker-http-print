package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Fprint(w, err)
		fmt.Println(err)
	}
	fmt.Fprint(w, string(requestDump))
	fmt.Println(string(requestDump))
}

func main() {
	router := httprouter.New()
	router.GET("/*all", index)
	router.POST("/*all", index)
	router.DELETE("/*all", index)
	router.HEAD("/*all", index)
	router.PUT("/*all", index)
	router.PATCH("/*all", index)

	port := getPort("PORT", ":8080")
	fmt.Printf("Running on: http://%s%s\n", getIP(), port)
	log.Fatal(http.ListenAndServe(port, router))
}

func getIP() string {
	inter, err := net.Interfaces()
	if err != nil {
		return "127.0.0.1"
	}
	var ip net.IP
	for _, i := range inter {
		addrs, err := i.Addrs()
		if err != nil {
			return "127.0.0.1"
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
		}
	}
	return ip.String()
}

func getPort(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		port := value
		if !strings.HasPrefix(value, ":") {
			port = ":" + value
		}
		return port
	}
	return fallback
}
