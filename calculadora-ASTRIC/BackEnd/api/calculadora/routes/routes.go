package routes

import (
	"ASTRIC/BackEnd/api/calculadora/controllers"
	"ASTRIC/BackEnd/shared/routes"
)

func rutas(ruta routes.TipoRuta) {
	ruta("/history", controllers.GetHistory, "GET", "Ruta para obtener el historial")
	ruta("/history", controllers.PostHistory, "POST", "Ruta para agregar elementos al historial")
}
