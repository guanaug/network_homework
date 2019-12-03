package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"network/global/logger"
	"network/model/user"
	"network/util/password"
	"strconv"
	"time"
)

// TODO tag检验
type User struct {
	ID         int64     `json:"id,omitempty"`
	Account    string    `json:"account,omitempty" binding:"required"`
	Name       string    `json:"name,omitempty"`
	// 这个Type暂时想不到有什么作用，先不用
	//Type       int8      `json:"type, notnull"` // 1: 管理员 2:市级单位 3:市级各辖区单位 4:受监管企业单位 5:签约技术支持/安全服务单位
	Phone      string    `json:"phone,omitempty"`
	Email      string    `json:"email,omitempty"`
	WeChat     string    `json:"wechat,omitempty"`
	Department int64     `json:"department,omitempty"`
	Valid      bool      `json:"valid,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	DeletedAt  time.Time `json:"deleted_at,omitempty"`
	Password   string    `json:"password,omitempty"`
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
		//Type:       u.Type,
		Type:		-1,
		Phone:      u.Phone,
		Email:      u.Email,
		WeChat:     u.WeChat,
		Department: u.Department,
		Valid:      true,
		Password:   password.New(u.Password),
	}

	if err := model.Insert(); err != nil {
		logger.Logger().Warn("insert user error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": model.ID})
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
		//Type:       u.Type,
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
		Offset *int `form:"offset" binding:"exists,gte=0"`	// 这里不能用required，因为 offset=0的时候校验不能通过
		Limit  int `form:"limit" binding:"required,max=200"`
	}{}

	if err := c.BindQuery(&page); err != nil {
		logger.Logger().Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "分页参数有误！"})
		return
	}

	users, count, err := user.List(*page.Offset, page.Limit)
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
			//Type:       u.Type,
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
