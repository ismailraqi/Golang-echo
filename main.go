package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	// declare & inisialize port variable from the envirenement variables
	port := os.Getenv("MY_APP_PORT")
	// check if port is null
	if port == "" {
		port = "3030"
	}
	// declare & inisialize new instance of echo
	e := echo.New()
	// declare & inisialize slice of maps
	products := []map[int]string{{1: "TVs"}, {2: "Laptops"}, {3: "Desktops"}}

	// GET METHOD ( retreive all data )
	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, products)
	})

	// GET METHOD ( retreive only one product )
	e.GET("/product/:id", func(c echo.Context) error {
		// declare & inisialize slice of map
		var product map[int]string
		for _, p := range products {
			for k := range p {
				pID, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					return err
				}
				if pID == k {
					product = p
				}
			}
		}
		if product == nil {
			return c.JSON(http.StatusNotFound, "Product not found")
		}
		return c.JSON(http.StatusOK, product)
	})

	// POST METHOD ( APPEND ONE PRODUCT TO THE MAP )
	e.POST("/product", func(c echo.Context) error {
		type body struct {
			Name string `json:"product_name"`
		}
		var reqBody body
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		product := map[int]string{
			len(products) + 1: reqBody.Name,
		}
		products = append(products, product)
		return c.JSON(http.StatusOK, product)

	})
	e.Logger.Print(fmt.Sprintf("Listening on prot %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
