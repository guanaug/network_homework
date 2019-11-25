package session

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"sync"
)

const(
	secretKey = "example-secret-key"
	sessionName = "network-homework"
	loginName = "username"
)

var sess *sessions.Session
var store * sessions.CookieStore

func init()  {
	once := sync.Once{}

	once.Do(func() {
		store = sessions.NewCookieStore([]byte(secretKey))
	})
}

func Session(c *gin.Context) *sessions.Session {
	var err error
	sess, err = store.Get(c.Request, sessionName)
	if err != nil {
		panic(err)
	}

	return sess
}

func IsLogin(c *gin.Context) bool {
	if _, ok := Session(c).Values[loginName]; ok {
		return true
	}

	return false
}

func Login(c *gin.Context, value string) error {
	Session(c).Values[loginName] = value
	return sessions.Save(c.Request, c.Writer)
}

func Logout(c *gin.Context) error {
	delete(Session(c).Values, loginName)
	return sessions.Save(c.Request, c.Writer)
}