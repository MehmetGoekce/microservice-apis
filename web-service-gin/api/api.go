package api

import (
    "net/http"
	"time"
    "github.com/gin-gonic/gin"
	"web-service-gin/models"
)

var order = models.GetOrder{
    ID:      "cb9082f7-3aa3-4278-a602-cc3b125733a8",
    Status:  models.StatusDelivered,
    Created: time.Now(),
    OrderItems: []models.OrderItem{
	{
	    Product: "Kaffee",
	    Size:    models.SizeMedium,
	    Quantity: 1,
	},
    },
}

func MountRoutes(router *gin.Engine) {
    router.GET("/orders", getOrders)
    router.POST("/orders", createOrder)
    router.GET("/orders/:order_id", getOrder)
    router.PUT("/orders/:order_id", updateOrder)
    router.DELETE("/orders/:order_id", deleteOrder)
    router.POST("/orders/:order_id/cancel", cancelOrder)
    router.POST("/orders/:order_id/pay", payOrder)
}

func getOrders(c *gin.Context) {
    c.JSON(http.StatusOK, []models.GetOrder{order})
}

func createOrder(c *gin.Context) {
    // Currently returns the pre-defined order (no actual creation)
    c.JSON(http.StatusCreated, order)
}

func getOrder(c *gin.Context) {
    c.JSON(http.StatusOK, order)
}

func updateOrder(c *gin.Context) {
    // Currently returns the pre-defined order (no actual update)
    c.JSON(http.StatusOK, order)
}

func deleteOrder(c *gin.Context) {
    // Cannot delete with pre-defined data, consider returning an error
    c.JSON(http.StatusOK, order)
}

func cancelOrder(c *gin.Context) {
    // Currently returns the pre-defined order (no actual cancellation)
    c.JSON(http.StatusOK, order)
}

func payOrder(c *gin.Context) {
    // Currently returns the pre-defined order (no actual payment processing)
    c.JSON(http.StatusOK, order)
}
