package server

import (
	"fmt"
	"go-testing/controllers"
	"log"
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/locations/country/", controllers.GetCountry)

	fmt.Println("Listening on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", logMiddleware(mux)))
}

func logMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("%s %s %s\n", request.RemoteAddr, request.Method, request.URL)
		handler.ServeHTTP(writer, request)
	})
}
