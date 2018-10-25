package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mc0239/kumuluzee-go-discovery/discovery"

	"github.com/gin-gonic/gin"
	"github.com/mc0239/kumuluzee-go-config/config"
)

type Customer struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type OrderRequest struct {
	CustomerID  int64  `json:"customerId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type OrderResponse struct {
	ID          int64  `json:"id"`
	CustomerID  int64  `json:"customerId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var mockDB []Customer
var conf config.Util
var disc discovery.Util

func main() {

	initDB()
	initConfig()
	initDiscovery()
	disc.RegisterService(discovery.RegisterOptions{})

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		maintenanceMode, _ := conf.GetBool("rest-config.maintenance")
		if maintenanceMode {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, ErrorResponse{
				http.StatusServiceUnavailable,
				"Service is undergoing maintenance, check back in a minute.",
			})
		} else {
			c.Next()
		}
	})

	v1c := router.Group("/v1/customers")
	{
		v1c.GET("/", getCustomers)
		v1c.GET("/:id", getCustomerByID)
		v1c.GET("/:id/order")
	}

	router.Run(":9000")
}

func getCustomers(c *gin.Context) {
	c.JSON(http.StatusOK, mockDB)
	return
}

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

func initDB() {
	mockDB = make([]Customer, 0)
	mockDB = append(mockDB,
		Customer{100, "John", "Carlile", "john.ca@mail.com", "053347863"},
		Customer{101, "Ann", "Lockwood", "lockwood_ann@mail.com", "023773123"},
		Customer{102, "Elizabeth", "Mathews", "eli23@mail.com", "043343403"},
		Customer{103, "Isaac", "Anderson", "isaac.anderson@mail.com", "018743831"},
		Customer{104, "Barret", "Peyton", "barretp@mail.com", "063343148"},
		Customer{105, "Terry", "Cokes", "terry_cokes@mail.com", "053339123"},
	)
}

func initConfig() {
	conf = config.NewUtil(config.Options{
		Extension:  "consul",
		ConfigPath: "config.yaml",
	})
}

func initDiscovery() {
	disc = discovery.New(discovery.Options{
		Extension:  "consul",
		ConfigPath: "config.yaml",
	})
}
