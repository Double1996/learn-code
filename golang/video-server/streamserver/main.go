package streamserver

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHandler)
	router.POST("/videos/:vid-id", updaloadHandler)
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":9000", r)
}
