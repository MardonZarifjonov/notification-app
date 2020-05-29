package controllers

import (
	"math/rand"
	"net/http"
	"notification-app/api/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// SendNotification func to create a notification for sending
func SendNotification(c *gin.Context) {
	// Initializing required variables
	var user models.User
	var order models.Order
	var notification models.Notification
	var mailMessage string
	var smsMessage string

	var notificationMessage string
	var emailStatus int
	var smsStatus int

	// Extracting values from the request body
	goods := c.Request.FormValue("goods")
	email := c.Request.FormValue("email")
	phoneNumber := c.Request.FormValue("phone_number")
	totalPrice := c.Request.FormValue("total_price")

	typeEmail, err := strconv.Atoi(c.Request.FormValue("send_by_email"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err,
		})
		return
	}
	typeSms, err := strconv.Atoi(c.Request.FormValue("send_by_sms"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err,
		})
		return
	}
	user.Email = email
	user.Phone = phoneNumber

	// Storing user data into database
	if err := user.Create(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		return
	}

	tPrice, err := strconv.ParseFloat(totalPrice, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err,
		})
		return
	}
	order.Goods = goods
	order.TotalPrice = tPrice

	// Storing order data into database
	if err := order.Create(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		return
	}
	// Initializing notification message
	notificationMessage = "Your purchase is completed. Goods you purchased: " + goods + ". Total price: " + totalPrice + "$. Thank you for your purchase"

	if typeEmail == 1 {
		mailMessage, emailStatus = sendMessageByEmail(notificationMessage)
	}
	if typeSms == 1 {
		smsMessage, smsStatus = sendMessageBySMS(notificationMessage)
	}

	notification.OrderID = order.ID
	notification.UserID = user.ID
	notification.Message = notificationMessage
	notification.SendByEmail = &emailStatus
	notification.SendBySMS = &smsStatus

	// Storing notification data into database
	if err := notification.Create(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": mailMessage + smsMessage,
		"message":  notificationMessage,
	})
}

// SendMessageByEmail func to send notification by email
func sendMessageByEmail(messageToSend string) (string, int) {
	var statusMessage string
	// Simulation of sending the notification by email
	randomStatus := sendSimulation()

	if randomStatus == -1 {
		statusMessage = "Error occured during notification sending by email. "
	} else if randomStatus == 0 {
		statusMessage = "Sending notification by email in waiting status. "
	} else if randomStatus == 1 {
		statusMessage = "Notification sent successfully by email. "
	}
	return statusMessage, randomStatus
}

// SendMessageBySMS func to send notification by sms
func sendMessageBySMS(messageToSend string) (string, int) {
	var statusMessage string
	// Simulation of sending the notification by SMS
	randomStatus := sendSimulation()

	if randomStatus == -1 {
		statusMessage = "Error occured during notification sending by SMS. "
	} else if randomStatus == 0 {
		statusMessage = "Sending notification by SMS in waiting status. "
	} else if randomStatus == 1 {
		statusMessage = "Notification sent successfully by SMS. "
	}
	return statusMessage, randomStatus
}

// SendSimulation func to simulate sending the notification by generating random status
func sendSimulation() int {
	rand.Seed(time.Now().UnixNano())
	min := -1
	max := 1
	randomStatus := rand.Intn(max-min+1) + min
	return randomStatus
}
