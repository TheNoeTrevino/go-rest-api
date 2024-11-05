package main

// TODO: add handlers once we add them in
import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"rest-api-project/internal/handlers"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()

	handlers.Handler(r)
	fmt.Printf("Starting Go API")
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
