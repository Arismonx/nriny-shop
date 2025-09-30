package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App      App
		Db       DB
		Jwt      Jwt
		Kafka    Kafka
		Grpc     Grpc
		Paginate Paginate
	}

	App struct {
		AppStage string
		AppName  string
		AppUrl   string
	}

	DB struct {
		DB_Url string
	}

	Jwt struct {
		JwtAccessSecretKey  string
		JwtAccessDuration   int64
		JwtRefreshSecretKey string
		JwtRefreshDuration  int64
		JwtApiSecretKey     string
		JwtApiDuration      int64
	}

	Kafka struct {
		KafkaUrl       string
		KafkaApiKey    string
		KafkaApiSecret string
	}

	Grpc struct {
		GrpcAuthUrl      string
		GrpcItemUrl      string
		GrpcPlayerUrl    string
		GrpcInventoryUrl string
		GrpcPaymentUrl   string
	}

	Paginate struct {
		PaginateItemNextPageBaseUrl      string
		PaginateInventoryNextPageBaseUrl string
	}
)

func LoadConfig(path string) Config {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error loading .env file!")
	}

	cfg := Config{
		App: App{
			AppStage: os.Getenv("APP_STAGE"),
			AppName:  os.Getenv("APP_NAME"),
			AppUrl:   os.Getenv("APP_URL"),
		},
		Db: DB{
			DB_Url: os.Getenv("DB_URL"),
		},
		Jwt: Jwt{
			JwtAccessSecretKey: os.Getenv("JWT_ACCESS_SECRET_KEY"),
			JwtAccessDuration: func() int64 {
				result, err := strconv.ParseInt(os.Getenv("JWT_ACCESS_DURATION"), 10, 64)
				if err != nil {
					log.Fatalf("Failed to convert string to integer: %v", err)
				}
				return result
			}(),
			JwtRefreshSecretKey: os.Getenv("JWT_REFRESH_SECRET_KEY"),
			JwtRefreshDuration: func() int64 {
				result, err := strconv.ParseInt(os.Getenv("JWT_REFRESH_DURATION"), 10, 64)
				if err != nil {
					log.Fatalf("Failed to convert string to integer: %v", err)
				}
				return result
			}(),
			JwtApiSecretKey: os.Getenv("JWT_API_SECRET_KEY"),
		},
		Kafka: Kafka{
			KafkaUrl:       os.Getenv("KAFKA_URL"),
			KafkaApiKey:    os.Getenv("KAFKA_API_KEY"),
			KafkaApiSecret: os.Getenv("KAFKA_API_SECRET"),
		},
		Grpc: Grpc{
			GrpcAuthUrl:      os.Getenv("GRPC_AUTH_URL"),
			GrpcItemUrl:      os.Getenv("GRPC_ITEM_URL"),
			GrpcPlayerUrl:    os.Getenv("GRPC_PLAYER_URL"),
			GrpcInventoryUrl: os.Getenv("GRPC_INVENTORY_URL"),
			GrpcPaymentUrl:   os.Getenv("GRPC_PAYMENT_URL"),
		},
		Paginate: Paginate{
			PaginateItemNextPageBaseUrl:      os.Getenv("PAGINATE_ITEM_NEXT_PAGE_BASED_URL"),
			PaginateInventoryNextPageBaseUrl: os.Getenv("PAGINATE_INVENTORY_NEXT_PAGE_BASED_URL"),
		},
	}

	return cfg
}
