package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// const (
// 	ContentTypeBinary = "application/octet-stream"
// 	ContentTypeForm   = "application/x-www-form-urlencoded"
// 	ContentTypeJSON   = "application/json"
// 	ContentTypeHTML   = "text/html; charset=utf-8"
// 	ContentTypeText   = "text/plain; charset=utf-8"
// )

type loginUserRequest struct {
	Login    string `form:"login"`
	Password string `form:"password"`
}

type loginUserResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
	User                 string    `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Login)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// err = util.CheckPassword(req.Password, user.Password)
	// if err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	// 	return
	// }

	if req.Password != user.Password {
		ctx.JSON(http.StatusUnauthorized, "passwords dont match")
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Login, server.config.AccessTokenDuration)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	_ = loginUserResponse{
		AccessToken: accessToken,
		User:        user.Login,
	}

	//ctx.JSON(http.StatusOK, rsp)

	fmt.Println("loginUser")

	ctx.HTML(http.StatusOK, "admin.html", gin.H{"title": "Админ панель"})

}
