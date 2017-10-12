package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/alexiskoss/info344-in-class/zipsvr/models"
)

type CityHandler struct {
	PathPrefix string
	Index      models.ZipIndex
}

//receiver function ch *CityHandler sits on the left side of the function
//allows you to refer to the struct of the handler, .this pointer in JS
//will be called from the mux of the server
func (ch *CityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// URL: /zips/city-name need to grab last token off the URL
	//starts after the path prefix /zips/ to get city name, doesn't end
	cityName := r.URL.Path[len(ch.PathPrefix):]
	cityName = strings.ToLower(cityName)
	if len(cityName) == 0 {
		http.Error(w, "please provide a city name", http.StatusBadRequest)
		return
	}

	w.Header().Add(headerContentType, contentTypeJSON)
	w.Header().Add(accessControlAllowOrigin, "*")
	zips := ch.Index[cityName]
	json.NewEncoder(w).Encode(zips)
}
