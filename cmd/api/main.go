package main

import (
	"fmt"
	"net/http"

	"github.com/fedemiodo/Crabi-code-challenge/internal/api"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
)

/*

CHANGELOG
--- v0.1 ---
FEATURES:
1) Endpoint CreateUser. No soporta el servicio externo de PLD
2) Endpoint GetUserInformation. No hay auth, solo username por parametro.

TO-DO
1) consumir PLD externo
2) Endpoint login - respuesta de token
3) usar el token de respuesta como autenticacion de GetUserInformation
4) Refactor - organizar en paquetes/modularizar.

*/

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	api.Handler(r)
	fmt.Println(`

    ______           __    _    ______          __        ________          ____                    
   / ____/________ _/ /_  (_)  / ____/___  ____/ /__     / ____/ /_  ____ _/ / /__  ____  ____ ____ 
  / /   / ___/ __  / __ \/ /  / /   / __ \/ __  / _ \   / /   / __ \/ __  / / / _ \/ __ \/ __  / _ \
 / /___/ /  / /_/ / /_/ / /  / /___/ /_/ / /_/ /  __/  / /___/ / / / /_/ / / /  __/ / / / /_/ /  __/
 \____/_/   \__,_/_.___/_/   \____/\____/\__,_/\___/   \____/_/ /_/\__,_/_/_/\___/_/ /_/\__, /\___/ 
                                                                                         ___/	`)
	fmt.Println("Initializing API Rest Server on port 8000...")
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
