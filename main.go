package main

import (
	"fmt"
	"github.com/Luthfiansyah/warpin-message/app/handlers"
	"github.com/Luthfiansyah/warpin-message/config"
	"github.com/Luthfiansyah/warpin-message/database"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

var client *redis.Client
var router *gin.Engine

func init() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("toml")
	//var configuration config.Configurations
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	//Set undefined variables
	//viper.SetDefault("PORT", "8888")
}

func initRouter() {

	// INIT ROUTER
	router = gin.Default()
	initRoutes(viper.GetBool("DEBUG_MODE"))
	port := config.PORT
	appKey := viper.GetString("APP_KEY")

	// SET START TIME SERVER
	database.SetRedis(appKey, handlers.GetCurrentTimeTimeZone("Asia/Jakarta"))

	router.Run(":" + strconv.Itoa(port))
}

func Ping(c *gin.Context) {
	appKey := viper.GetString("APP_KEY")
	date := database.GetRedis(appKey)
	res := map[string]string{
		"start_time": date,
		"message":    "Warpin Message Run On " + handlers.GetRunMode() + " Mode",
	}
	c.JSON(http.StatusOK, res)
}

func main() {
	initRouter()
}
