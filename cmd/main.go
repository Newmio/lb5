package main

import (
	handler "lb5/pkg/handler"
	"lb5/pkg/postgres/database"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.ReadInConfig()

	db, err := database.Init(database.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Name:     viper.GetString("db.name"),
		DbName:   viper.GetString("db.dbname"),
		Password: viper.GetString("db.password"),
		SslMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	r := gin.Default()

	handler.InitRoutes(db, r)
	r.Run(":8081")
}
