package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

// ProductValidator echo validator for product
type ProductValidator struct {
	validator *validator.Validate
}

// Validate product request body
func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

func main() {
	// declare & inisialize port variable from the envirenement variables
	port := os.Getenv("MY_APP_PORT")
	// check if port is null
	if port == "" {
		port = "3030"
	}
	// declare & inisialize new instance of echo
	e := echo.New()
	// declare & inisialize a new instance of Validator
	v := validator.New()
	// declare & inisialize slice of maps
	products := []map[int]string{{1: "TVs"}, {2: "Laptops"}, {3: "Desktops"}, {4: "test"}}

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
		// declare & inisialize request body
		type body struct {
			Name string `json:"product_name" validate:"min=4"`
		}
		// declare & inisialize variable typeof body (struct)
		var reqBody body
		e.Validator = &ProductValidator{validator: v}
		// bind data given fro user & check if there is an error
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		// validates a structs exposed fields & check if there is an error
		if err := c.Validate(reqBody); err != nil {
			return err
		}
		// declare & inisialize a map to store the given product
		product := map[int]string{
			len(products) + 1: reqBody.Name,
		}
		// append the given product to the PRODUCTS slice maps
		products = append(products, product)
		return c.JSON(http.StatusOK, product)

	})
	// PUT METHOD
	e.PUT("/product/:id", func(c echo.Context) error {
		var product map[int]string
		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		for _, p := range products {
			for k := range p {

				if pID == k {
					product = p
				}
			}
		}
		if product == nil {
			return c.JSON(http.StatusNotFound, "Product not found")
		}
		type body struct {
			Name string `json:"product_name" validate:"min=4"`
		}
		var reqBody body
		e.Validator = &ProductValidator{validator: v}
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		if err := c.Validate(reqBody); err != nil {
			return err
		}
		product[pID] = reqBody.Name
		return c.JSON(http.StatusOK, product)
	})

	//DELETE METHOD
	e.DELETE("product/:id", func(c echo.Context) error {
		var product map[int]string
		var index int
		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		for i, p := range products {
			for k := range p {

				if pID == k {
					product = p
					index = i
				}
			}
		}
		if product == nil {
			return c.JSON(http.StatusNotFound, "Product not found")
		}
		splice := func(s []map[int]string, index int) []map[int]string {
			return append(s[index:], s[index+1:]...)
		}
		products = splice(products, index)
		return c.NoContent(http.StatusNoContent)
	})

	e.Logger.Print(fmt.Sprintf("Listening on prot %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
