package http

import (
	"fmt"
	"net/http"
)

// HelloWorld return void
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
	return
}
