package tronics

import (
	"fmt"
	"net/http"
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

//Products declare & inisialize slice of maps
var Products = []map[int]string{{1: "TVs"}, {2: "Laptops"}, {3: "Desktops"}, {4: "test"}}

// handler to get all products
func getProducts(c echo.Context) error {
	fmt.Printf("getproduct : %v\n", c.Request())
	return c.JSON(http.StatusOK, Products)
}

// handler to get only one product filtering with id
func getProduct(c echo.Context) error {
	// declare & inisialize slice of map
	var product map[int]string
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for _, p := range Products {
		for k := range p {
			if pID == k {
				product = p
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	fmt.Println("test")
	return c.JSON(http.StatusOK, product)

}

// handler to create a new product
func createProduct(c echo.Context) error {
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
		len(Products) + 1: reqBody.Name,
	}
	// append the given product to the PRODUCTS slice maps
	Products = append(Products, product)
	return c.JSON(http.StatusOK, product)
}

// handler to update one product filtering with id
func updateProduct(c echo.Context) error {
	var product map[int]string
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for _, p := range Products {
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
}

// handler to delete a product filtering with id
func deleteProduct(c echo.Context) error {
	var product map[int]string
	var index int
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for i, p := range Products {
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
		return append(s[:index], s[index+1:]...)
	}
	Products = splice(Products, index)
	return c.NoContent(http.StatusNoContent)
}
