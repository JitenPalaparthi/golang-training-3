package handlers

import (
	"demo/models"
	"demo/share"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	Create(*gin.Context)
	GetByID(*gin.Context)
	DeleteByID(ctx *gin.Context)
}

type UserHandler struct {
	IUser share.IUser
}

func NewUserHandler(iuser share.IUser) IUserHandler {
	return &UserHandler{IUser: iuser}
}

func (uh *UserHandler) Create(ctx *gin.Context) {
	user := new(models.User)
	err := ctx.Bind(user) //
	// bytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, "invalid user data")
		ctx.Abort()
	}
	// err = json.Unmarshal(bytes,user)
	// if err != nil {
	// 	ctx.String(http.StatusBadRequest, "invalid user data")
	// 	ctx.Abort()
	// }

	err = user.Validate()
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}
	user.Status = "active"
	user.LastUpdated = time.Now().Unix()

	//utils.ChUser <- user
	user, err = uh.IUser.Create(user)

	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}
	ctx.JSON(201, user)
}
func (uh *UserHandler) GetByID(ctx *gin.Context) {

	id, ok := ctx.Params.Get("id")
	if !ok {
		log.Println("invalid id ")
		ctx.String(http.StatusBadRequest, "invalid id")
		ctx.Abort()
	}

	user, err := uh.IUser.GetByID(id)

	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}

	ctx.JSON(200, user)
}

func (uh *UserHandler) DeleteByID(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		log.Println("invalid id ")
		ctx.String(http.StatusBadRequest, "invalid id")
		ctx.Abort()
	}

	r, err := uh.IUser.DeleteByID(id)

	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}

	if r > 0 {
		ctx.String(202, "successfully deleted")
	} else {
		ctx.String(202, "no records to delete")
	}

}
