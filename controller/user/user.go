package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"network/global/logger"
	"network/global/session"
	"network/model/user"
	"strconv"
	"strings"
	"time"
)

// TODO tag检验
type User struct {
	ID         int64     `json:"id, omitempty"`
	Account    string    `json:"account, omitempty" binding:"required"`
	Name       string    `json:"name, notnull"`
	Type       int8      `json:"type, notnull"` // 1: 管理员 2:市级单位 3:市级各辖区单位 4:受监管企业单位 5:签约技术支持/安全服务单位
	Phone      string    `json:"phone" binding:"phone"`
	Email      string    `json:"email" binding:"email"`
	WeChat     string    `json:"wechat"`
	Department int64     `json:"department"`
	Valid      bool      `json:"valid, notnull"`
	CreatedAt  time.Time `json:"created_at, notnull"`
	ModifiedAt time.Time `json:"modified_at, notnull"`
	DeletedAt  time.Time `json:"deleted_at"`
	Password   string    `json:"password"`
}

func Login(c *gin.Context) {
	loginInfo := struct {
		Account  string `binding:"required"`
		Password string `binding:"required"`
	}{}

	if err := c.BindQuery(&loginInfo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "账号和密码不能为空！"})
		return
	}

	if session.IsLogin(c) {
		c.Status(http.StatusOK)
		return
	}

	u := &user.User{
		Account:  strings.TrimSpace(loginInfo.Account),
		Password: strings.TrimSpace(loginInfo.Password),
	}

	ok, err := u.Login()
	if err != nil {
		logger.Logger().Warn("query login info error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "账号或密码错误！"})
		return
	}

	if err := session.Login(c, loginInfo.Account); err != nil {
		logger.Logger().Warn("add session error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	userBriefInfo := struct {
		Name string `json:"name"`
	}{u.Name}

	c.JSON(http.StatusOK, &userBriefInfo)
}

func Logout(c *gin.Context) {
	if err := session.Logout(c); err != nil {
		logger.Logger().Warn("delete session error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func Add(c *gin.Context) {
	u := User{}

	if err := c.BindJSON(&u); err != nil {
		logger.Logger().Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "用户信息填写有误！"})
		return
	}

	model := &user.User{
		Account:    u.Account,
		Name:       u.Name,
		Type:       u.Type,
		Phone:      u.Phone,
		Email:      u.Email,
		WeChat:     u.WeChat,
		Department: u.Department,
		Valid:      true,
		Password:   u.Password,
	}

	if err := model.Insert(); err != nil {
		logger.Logger().Warn("insert user error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func Delete(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "用户ID有误！"})
		return
	}

	model := &user.User{
		ID: userID,
	}

	if err := model.Delete(); err != nil {
		logger.Logger().Warn("user delete error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func Modify(c *gin.Context) {
	u := User{}

	if err := c.BindJSON(&u); err != nil {
		logger.Logger().Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "用户信息填写有误！"})
		return
	}

	model := &user.User{
		ID:         u.ID,
		Account:    u.Account,
		Name:       u.Name,
		Type:       u.Type,
		Phone:      u.Phone,
		Email:      u.Email,
		WeChat:     u.Email,
		Department: u.Department,
		Valid:      true,
		Password:   u.Password,
	}

	if err := model.Update(); err != nil {
		logger.Logger().Warn("user add error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func List(c *gin.Context) {
	page := struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	}{}

	if err := c.BindQuery(&page); err != nil {
		logger.Logger().Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "分页参数有误！"})
		return
	}

	users, count, err := user.List(page.Offset, page.Limit)
	if err != nil {
		logger.Logger().Warn("query users error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	userBriefInfos := struct {
		Count int    `json:"count"`
		List  []User `json:"users"`
	}{}

	for _, u := range users {
		userBriefInfos.List = append(userBriefInfos.List, User{
			ID:         u.ID,
			Account:    u.Account,
			Name:       u.Name,
			Type:       u.Type,
			Phone:      u.Phone,
			Email:      u.Email,
			WeChat:     u.WeChat,
			Department: u.Department,
			Valid:      u.Valid,
			CreatedAt:  u.CreatedAt,
			ModifiedAt: u.ModifiedAt,
			DeletedAt:  u.DeletedAt,
		})
	}
	userBriefInfos.Count = count

	c.JSON(http.StatusOK, &userBriefInfos)
}

func Info(c *gin.Context) {

}
