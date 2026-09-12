package main

import (
	"log"
	"net/http"

	"goshop/internal/cache"
	"goshop/internal/config"
	"goshop/internal/db"
	"goshop/internal/handlers"
	appkafka "goshop/internal/kafka"
	"goshop/internal/services"
)

func main() {
	cfg := config.Load()

	dbConn, err := db.NewPostgres(cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	redisCache := cache.NewProductCache(
		cfg.RedisAddr,
	)

	producer := appkafka.NewProducer(
		cfg.KafkaBroker,
		cfg.KafkaTopic,
	)
	defer producer.Close()

	productRepo := db.NewProductRepo(dbConn)
	orderRepo := db.NewOrderRepo(dbConn)

	productService := services.NewProductService(
		productRepo,
		redisCache,
	)

	orderService := services.NewOrderService(
		orderRepo,
		producer,
	)

	productHandler := handlers.NewProductHandler(
		productService,
	)

	orderHandler := handlers.NewOrderHandler(
		orderService,
	)

	http.HandleFunc(
		"/products/",
		productHandler.GetProduct,
	)

	http.HandleFunc(
		"/orders",
		orderHandler.CreateOrder,
	)

	log.Println("Server started on :8080, environment:", cfg.AppEnv)

	if err := http.ListenAndServe(
		":8080",
		nil,
	); err != nil {
		log.Fatal(err)
	}
}
