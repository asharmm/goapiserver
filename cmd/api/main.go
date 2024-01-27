package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/asharmm/goapiserver/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()

	handlers.Handlers(r)

	fmt.Println("Starting GO API server...")


	fmt.Println(`

					             
	 ___  ___ _ ____   _____ _ __   _ __ _   _ _ __  _ __ (_)_ __   __ _ 
	/ __|/ _ \ '__\ \ / / _ \ '__| | '__| | | | '_ \| '_ \| | '_ \ / _' |
	\__ \  __/ |   \ V /  __/ |    | |  | |_| | | | | | | | | | | | (_| |
	|___/\___|_|    \_/ \___|_|    |_|   \__,_|_| |_|_| |_|_|_| |_|\__, |
									|___/    
							on :8000                                               						
	
	`)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error("Error starting GO API server : ", err)
	}


	
}