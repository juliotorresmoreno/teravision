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
	g.PUT("", usersPUT)
}

func usersGET(c echo.Context) error {
	conn, err := db.NewConn()
	if err != nil {
		c.Logger().Debug(err)
		return c.String(http.StatusInternalServerError, "Conection failure")
	}
	defer conn.Close()
	users := make([]models.User, 0)
	err = conn.Find(&users)
	if err != nil {
		c.Logger().Debug(err)
		return c.String(http.StatusInternalServerError, "Conection failure")
	}
	return c.JSON(200, map[string]interface{}{
		"success": true,
		"data":    users,
	})
}

func usersPUT(c echo.Context) error {
	conn, err := db.NewConn()
	if err != nil {
		c.Logger().Debug(err)
		return c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"success": "error",
			"message": "Conection failure",
		})
	}
	defer conn.Close()

	user := new(models.User)
	if err = c.Bind(user); err != nil {
		c.Logger().Debug(err)
		return c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"success": "error",
			"message": err.Error(),
		})
	}
	user.CleanForInsert()
	_, err = conn.InsertOne(user)
	if err != nil {
		if err.Error()[:17] == "pq: duplicate key" {
			c.Logger().Debug(err)
			return c.JSON(http.StatusNotAcceptable, map[string]interface{}{
				"success": "error",
				"message": "The DNI already exists.",
			})
		}
		c.Logger().Debug(err)
		return c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"success": "error",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted, user)
}
