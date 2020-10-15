package tronics

import (
	"fmt"

	"github.com/labstack/echo/middleware"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

// declare & inisialize new instance of echo
var e = echo.New()

// declare & inisialize a new instance of Validator
var v = validator.New()

// function that read configuration from Environement before app start
func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		e.Logger.Fatal("unable to load configuration")
	}
}

//ServerMessage is a custom middleware just for testing
func ServerMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Inside middleware")
		c.Request().URL.Path = "/done"
		fmt.Printf("%+v\n", c.Request())
		return next(c)
	}
}

//Start function made to start the application
func Start() {
	/*
		// declare & inisialize port variable from the envirenement variables
		port := os.Getenv("MY_APP_PORT")
		// check if port is null
		if port == "" {
		port = "3030"
		}
	*/
	// Pre is a function that's execute with any root (using echo middleware)
	e.Pre(middleware.RemoveTrailingSlash())
	// GET METHOD ( retreive all data )
	e.GET("/products", getProducts, ServerMessage)
	// GET METHOD ( retreive only one product )
	e.GET("/product/:id", getProduct)
	// POST METHOD ( APPEND ONE PRODUCT TO THE MAP with BodyLimit using echo middleware)
	e.POST("/product", createProduct, middleware.BodyLimit("1K"))
	// PUT METHOD
	e.PUT("/product/:id", updateProduct)
	//DELETE METHOD
	e.DELETE("product/:id", deleteProduct)

	e.Logger.Print(fmt.Sprintf("Listening on prot %s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", cfg.Port)))
}
