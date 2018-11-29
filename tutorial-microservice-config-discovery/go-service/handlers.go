package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
	"github.com/mc0239/kumuluzee-go-discovery/discovery"
)

// returns array of customers with 200 OK code
func getCustomers(c *gin.Context) {
	c.JSON(http.StatusOK, mockDB)
	return
}

// returns user object with 200 OK code if found
// and 404 NOT FOUND code if such user doesn't exists
func getCustomerByID(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.ParseInt(sid, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			http.StatusBadRequest,
			fmt.Sprintf("ID conversion to integer failed with error: %s", err.Error()),
		})
		return
	}

	for _, e := range mockDB {
		if e.ID == id {
			c.JSON(http.StatusOK, e)
			return
		}
	}

	c.JSON(http.StatusNotFound, ErrorResponse{
		http.StatusNotFound,
		fmt.Sprintf("Customer with id %d not found.", id),
	})
	return
}

// this endpoint generates new Order Request and calls our Java service to create it
// Returns order with 201 CREATED code if successful.
func createOrder(c *gin.Context) {
	// prepare a new order
	sid := c.Param("id")
	id, err := strconv.ParseInt(sid, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			http.StatusBadRequest,
			fmt.Sprintf("ID conversion to integer failed with error: %s", err.Error()),
		})
		return
	}

	ord := OrderRequest{
		CustomerID:  id,
		Title:       "New order",
		Description: "This is a new order.",
	}

	// discover Java service to post order to
	ordAddress, err := disc.DiscoverService(discovery.DiscoverOptions{
		Value:       "java-service",
		Environment: "dev",
		Version:     "1.0.0",
		AccessType:  "direct",
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// pointer to OrderResponse, where request's response will be stored
	ordResp := &OrderResponse{}

	// perform POST request
	_, err = sling.New().Post(ordAddress).BodyJSON(ord).ReceiveSuccess(ordResp)
	if err != nil {
		// Java service returned something other than code 2xx
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, ordResp)
	return
}
