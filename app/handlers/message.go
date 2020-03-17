package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Luthfiansyah/warpin-message/app/types"
	"github.com/Luthfiansyah/warpin-message/config"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"net/http"
)

func GetAllMessage(c *gin.Context) {

	// GET TIME
	_ = GetCurrentTimestampTimeZone("Asia/Jakarta")

	var messages []types.MessageRequest

	// GET CACHE DATA COUNTRY INFO
	cache, err := GetCache(CacheKeyMessages())

	// CHECK IF CACHE EXIST DO THIS
	if cache != "" {
		// BIND CACHE TO STRUCT
		err = json.Unmarshal([]byte(cache), &messages)
		if err != nil {
			fmt.Println(err)
		}
	}

	// BUILD RESPONSE
	gr := GeneralResponseSuccessBuild(true,
		0, "Success")
	c.JSON(http.StatusOK, gin.H{
		"result":           messages,
		"general_response": gr,
	})
	return
}

func AddMessage(c *gin.Context) {

	// GET TIME
	reqTime := GetCurrentTimestampTimeZone("Asia/Jakarta")
	fmt.Println(reqTime)

	// BIND JSON BODY REQUEST
	var param types.MessageRequest
	if err := c.BindJSON(&param); err != nil {
		fmt.Println(err.Error())
		ShowResponseError(http.StatusBadRequest, c, 4016, err.Error())
		return
	}

	// CLEAR MEMORY PREVENTIVE MEMORY LEAK
	defer c.Request.Body.Close()

	conn, err :=  config.RabbitMqConnection()
	if err != nil {
		fmt.Println(err.Error())
		ShowResponseError(http.StatusInternalServerError, c, 5002, err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err.Error())
		ShowResponseError(http.StatusInternalServerError, c, 5002, err.Error())
	}

	// WILL BE EXECUTED LAST
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"messages",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	//failOnError(err, "Failed to declare an exchange")
	if err != nil {
		fmt.Println(err.Error())
		ShowResponseError(http.StatusInternalServerError, c, 5002, err.Error())
	}

	// DEFINE STRUCT
	var message types.MessageRequest
	var messages []types.MessageRequest

	// ASSIGN VALUE TO STRUCT
	message.Text = param.Text

	// GET CACHE DATA COUNTRY INFO
	cache, err := GetCache(CacheKeyMessages())

	// CHECK IF CACHE EXIST DO THIS
	if cache != "" {
		// BIND CACHE TO STRUCT
		err = json.Unmarshal([]byte(cache), &messages)
		if err != nil {
			fmt.Println(err)
		}
	}

	// APPEND TEXT TO ARRAY MESSAGE
	messages = append(messages, message)
	SetCache(CacheKeyMessages(), messages)

	// PUBLISH MESSAGE
	body := param.Text
	err = ch.Publish(
		"messages", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	//failOnError(err, "Failed to publish a message")
	if err != nil {
		fmt.Println(err.Error())
		ShowResponseError(http.StatusInternalServerError, c, 5002, err.Error())
	}

	// BUILD RESPONSE
	var data = types.MessageResponse{}
	data.Text = "Sent" + param.Text
	gr := GeneralResponseSuccessBuild(true,
		0, "Success")

	c.JSON(http.StatusOK, gin.H{
		"result":           &data,
		"general_response": gr,
	})

	return
}
