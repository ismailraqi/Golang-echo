package tronics

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)
//WriteCookie 
func WriteCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "Products"
	cookie.Value = "dataCookies"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}
