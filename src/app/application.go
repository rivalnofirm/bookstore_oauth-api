package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rivalnofirm/bookstore_oauth-api/src/src/clients/cassandra"
	"github.com/rivalnofirm/bookstore_oauth-api/src/src/domain/access_token"
	"github.com/rivalnofirm/bookstore_oauth-api/src/src/http"
	"github.com/rivalnofirm/bookstore_oauth-api/src/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()

	atHandler := http.NewHandler(access_token.NewService(db.NewRepsitory()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080")
}
