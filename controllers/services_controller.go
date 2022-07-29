package controllers

import (
	"MeLi/commons"
	singleton "MeLi/singletons"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//EndPoint type POST to return the location and message
func PostTopSecret(writer http.ResponseWriter, request *http.Request) {
	result := singleton.ResultPosition{}
	satellitesBody := singleton.SatellitesBody{}
	error := json.NewDecoder(request.Body).Decode(&satellitesBody)
	var satellites []singleton.SatelliteBody = satellitesBody.GetSatellites()

	if error != nil {
		commons.SendError(writer, http.StatusNotFound)
		return
	}

	result = GetTopSecret(satellites)
	json, _ := json.Marshal(result)

	commons.SendResponse(writer, http.StatusOK, json)
}

//EndPoint to return the location and message by one satellite
func TopSecret(writer http.ResponseWriter, request *http.Request) {
	urlVars := mux.Vars(request)
	var message string
	result := singleton.ResultPosition{}
	satelliteBody := singleton.SatelliteBody{}
	satelliteBody.Name = urlVars["satellite_name"]
	error := json.NewDecoder(request.Body).Decode(&satelliteBody)

	if error != nil {
		commons.SendError(writer, http.StatusNotFound)
		return
	}

	result, message = GetGetTopSecretOneSatellite(satelliteBody)

	if message != "" {
		json, _ := json.Marshal(message)
		commons.SendResponse(writer, http.StatusOK, json)
	} else {
		json, _ := json.Marshal(result)
		commons.SendResponse(writer, http.StatusOK, json)
	}

}
