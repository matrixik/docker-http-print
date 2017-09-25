package main

import (
	"fmt"
	"log"
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

	log.Fatal(http.ListenAndServe(getPort("PORT", ":8080"), router))
}

func getPort(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		port := value
		if !strings.HasPrefix(value, ":") {
			port = ":" + value
		}
		fmt.Printf("Running on port: %s\n", port)
		return port
	}
	return fallback
}
