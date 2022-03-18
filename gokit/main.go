package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"myapp/pkg"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {

	svc := pkg.NewHomeService()
	homeHandler := httptransport.NewServer(
		pkg.MakeHomeEndpoint(svc),
		decodeHomeRequest,
		httptransport.EncodeJSONResponse,
	)

	r := mux.NewRouter()
	r.Handle("/", homeHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
func decodeHomeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pkg.HomeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
