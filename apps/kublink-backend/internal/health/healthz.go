package health

import (
	"fmt"
	"net/http"
)

// HealthCheck godoc
//
//	@Summary		Server health check
//	@Description	Returns ok if the server is up and running
//	@Tags			system
//	@Success		200	{string}	string	"ok"
//	@Router			/healthz [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "ok\n")
}
