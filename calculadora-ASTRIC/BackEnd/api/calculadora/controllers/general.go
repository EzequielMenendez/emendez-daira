package controllers

import (
	"ASTRIC/BackEnd/api/calculadora/models"
	"ASTRIC/BackEnd/shared/ep"
	"net/http"
	"encoding/json"
)

// GetHistory es una función para obtener el historial
func GetHistory(w http.ResponseWriter, r *http.Request) {

	defer ep.ErrorControlResponse("/calculadora/history", w, r)

	res := ep.NewResponse("Obteniendo Historial", w)

	res.DatoSend(models.Historys)
}

// PostHistory es una función para agregar elementos al historial
func PostHistory(w http.ResponseWriter, r *http.Request) {

	defer ep.ErrorControlResponse("/calculadora/history", w, r)

	res := ep.NewResponse("Agregando al Historial", w)

	var newHistory models.History

	err := json.NewDecoder(r.Body).Decode(&newHistory)
    if err != nil {
        res.Err("Error al deserializar el cuerpo de la solicitud").DatoSend(models.Historys)
        return
    }

    models.Historys = append(models.Historys, newHistory)

    res.DatoSend(models.Historys)
}