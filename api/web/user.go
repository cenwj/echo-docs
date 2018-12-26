package web

import (
	"echo-docs/db"
	"echo-docs/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	db := db.MysqlConn()
	users := []model.User{}
	db.Find(&users)
	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)
	return c.JSON(http.StatusOK, users)
}

func Ab(c echo.Context) error {
	redis := db.RedisConn(1)
	err := redis.Set("a", "value", 0).Err()
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func CreateUsers(c echo.Context) error {
	rdb := db.RedisConn(2)
	rdb.Set("ket", "v", 0).Err()

	db := db.MysqlConn()

	user := &model.User{Name: c.QueryParam("name"),
		Password: c.QueryParam("password"),
		Username: c.QueryParam("username")}

	res := db.Create(user)

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, &user.ID)

}
