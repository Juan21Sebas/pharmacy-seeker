package main

import (
	"fmt"
	"log"
	"os"

	"Pharmacy/internal/infrastructure/api/pharmacy"

	"Pharmacy/internal/infrastructure/repositories"
	"Pharmacy/internal/pkg/usecase"

	"Pharmacy/cmd/config"

	"github.com/gin-gonic/gin"
)

const (
	localesURL = "https://farmanet.minsal.cl/index.php/ws/getLocales"
)

func main() {
	// Inicializar la aplicación de Gin
	r := gin.Default()

	//Configurar la dependencia HTTPClient con el cliente HTTP de la biblioteca estándar
	// Configurar la dependencia LocalRepository con el repositorio de locales de Farmanet
	localRepository := repositories.NewLocalRepository(localesURL)

	// Configurar la dependencia LocalUsecase con el caso de uso de locales
	localUsecase := usecase.NewLocalUseCase(localRepository)

	// Convertir localUsecase a usecase.LocalUseCase
	localUsecaseImpl := usecase.LocalUseCase(*localUsecase)

	// Configurar el controlador de locales con el caso de uso de locales
	localHandler := pharmacy.NewLocalHandler(localUsecaseImpl)

	pharmacy.Routers(r, localHandler)

	fmt.Println(localHandler)
	// Configurar las rutas de la aplicación
	//setupRouters(r, localHandler)
	config.LoadConfig()
	//Iniciar la aplicación
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Iniciando la aplicación en el puerto %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error al iniciar la aplicación: %v", err)
	}
}
