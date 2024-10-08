package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load() // envの時GO＿ENV読み込み、本番では実行されない
		if err != nil {
			log.Fatalln(err)
		}
	}

	// 環境変数の値をログに出力して確認
	log.Println("POSTGRES_USER:", os.Getenv("POSTGRES_USER"))
	log.Println("POSTGRES_PW:", os.Getenv("POSTGRES_PW"))
	log.Println("POSTGRES_HOST:", os.Getenv("POSTGRES_HOST"))
	log.Println("POSTGRES_PORT:", os.Getenv("POSTGRES_PORT"))
	log.Println("POSTGRES_DB:", os.Getenv("POSTGRES_DB"))

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Println("Database connection URL:", url)
		log.Fatalln(err)
	}

	fmt.Println("Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
