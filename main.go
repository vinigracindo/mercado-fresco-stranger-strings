package main

import (
	"database/sql"

	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/config"
)

// @title   Swagger Mercado Fresco
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos_y_condiciones

// @contact.name API Support
// @contact.url http://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @version         1.0
// @description     This is a sample server celler server.

// @host      localhost:8080
// @BasePath  /api/v1

var DB *sql.DB

func init() {
	DB = config.ConnectDb("mysql")

}

func main() {
	port := 8080
	server := server.NewAPIServer(DB)
	server.Run(port)
}
