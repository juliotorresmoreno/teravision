package controllers

import (
	"net/http"

	"github.com/juliotorresmoreno/teravision/db"
	"github.com/juliotorresmoreno/teravision/models"
	"github.com/labstack/echo"
)

// RouteUsers configure endoint /users
func RouteUsers(g *echo.Group) {
	g.GET("", usersGET)
	g.POST("", usersPOST)
	g.PATCH("", usersPOST)
	g.PUT("", usersPUT)
	g.DELETE("", usersDELETE)
}

func usersGET(c echo.Context) error {
	conn, err := db.NewConn()
	if err != nil {
		c.Logger().Debug(err)
		return c.String(http.StatusInternalServerError, "Conection failure")
	}
	defer conn.Close()
	users := make([]models.User, 0)
	err = conn.Find(users)
	if err != nil {
		c.Logger().Debug(err)
		return c.String(http.StatusInternalServerError, "Conection failure")
	}
	return c.JSON(200, map[string]interface{}{
		"success": true,
		"data":    users,
	})
}

func usersPOST(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"success": true,
	})
}

func usersPUT(c echo.Context) error {
	conn, err := db.NewConn()
	if err != nil {
		c.Logger().Debug(err)
		return c.String(http.StatusInternalServerError, "Conection failure")
	}
	defer conn.Close()
	users := models.User{}
	_, err = conn.InsertOne(users)
	if err != nil {
		c.Logger().Debug(err)
		return c.String(http.StatusInternalServerError, "Conection failure")
	}
	return c.JSON(200, map[string]interface{}{
		"success": true,
		"data":    users,
	})
}

func usersDELETE(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"success": true,
	})
}
