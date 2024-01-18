package main

import (
	"context"
	"fmt"
	"net/http"
	_ "omni-test/docs"
	"omni-test/handler"
	"omni-test/service"
	"omni-test/service/processor"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	ctx := context.Background()

	service := service.NewMockService()

	processor := processor.NewProcessor(service)

	handler := handler.NewHandler(processor)

	http.HandleFunc("/process_batch", func(w http.ResponseWriter, r *http.Request) {
		handler.ProcessBatch(ctx, w, r)
	})

	http.Handle("/docs/", httpSwagger.Handler(
		httpSwagger.URL("/static/swagger.json"),
	))

	http.Handle("/static/swagger.json", http.StripPrefix("/static/", http.FileServer(http.Dir("./docs"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("server is not started: %s\n", err)
	}
}
