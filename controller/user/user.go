package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"net/http"
	"network/global/logger"
	"network/model/department"
	"network/model/user"
	"network/util/password"
	"strconv"
	"time"
)

// TODO tag检验
type User struct {
	ID      int64  `json:"id,omitempty"`
	Account string `json:"account,omitempty" binding:"required"`
	Name    string `json:"name,omitempty"`
	// 这个Type暂时想不到有什么作用，先不用
	//Type       int8      `json:"type, notnull"` // 1: 管理员 2:市级单位 3:市级各辖区单位 4:受监管企业单位 5:签约技术支持/安全服务单位
	Phone          string    `json:"phone,omitempty"`
	Email          string    `json:"email,omitempty"`
	WeChat         string    `json:"wechat,omitempty"`
	Department     int64     `json:"department,omitempty"`
	DepartmentName string    `json:"department_name,omitempty"`
	DepartmentType int8      `json:"department_type,omitempty"`
	Valid          bool      `json:"valid,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	ModifiedAt     time.Time `json:"modified_at,omitempty"`
	DeletedAt      time.Time `json:"deleted_at,omitempty"`
	Password       string    `json:"password,omitempty"`
}

func Add(c *gin.Context) {
	u := User{}

	if err := c.BindJSON(&u); err != nil {
		logger.Logger().Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "用户信息填写有误！"})
		return
	}

	existUser, err := user.OneWithDeletedByAccount(u.Account)
	if err != nil && err != pg.ErrNoRows {
		logger.Logger().Warn("query user by account error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	model := &user.User{
		Account: u.Account,
		Name:    u.Name,
		//Type:       u.Type,
		Type:       -1,
		Phone:      u.Phone,
		Email:      u.Email,
		WeChat:     u.WeChat,
		Department: u.Department,
		Valid:      true,
		Password:   password.New(u.Password),
	}

	if existUser.DeletedAt.IsZero() {
		// 之前没有软删除过，直接插入
		err = model.Insert()
	} else {
		// 软删除过，则直接更新
		model.ID = existUser.ID
		err = model.Restore()
		err = model.Update()
	}
	if err != nil {
		logger.Logger().Warn("insert or user error:", err)
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

// TODO 无法禁用用户
func Modify(c *gin.Context) {
	u := struct {
		// TODO tag检验
		ID         int64  `json:"id,omitempty" binding:"required"`
		Name       string `json:"name,omitempty"`
		Phone      string `json:"phone,omitempty"`
		Email      string `json:"email,omitempty"`
		WeChat     string `json:"wechat,omitempty"`
		Department int64  `json:"department,omitempty"`
		Valid      *bool  `json:"valid,omitempty"`
		Password   string `json:"password,omitempty"`
	}{}

	if err := c.BindJSON(&u); err != nil {
		logger.Logger().Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "用户信息填写有误！"})
		return
	}

	model := &user.User{
		ID: u.ID,
	}
	// TODO ugly code
	if len(u.Name) > 0 {
		model.Name = u.Name
	}
	if len(u.Phone) > 0 {
		model.Phone = u.Phone
	}
	if len(u.Email) > 0 {
		model.Email = u.Email
	}
	if len(u.WeChat) > 0 {
		model.WeChat = u.WeChat
	}
	if u.Department > 0 {
		model.Department = u.Department
	}
	if u.Valid != nil {
		model.Valid = *u.Valid
	}
	if len(u.Password) > 0 {
		model.Password = password.New(u.Password)
	}

	if err := model.Update(); err != nil {
		logger.Logger().Warn("user update error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func List(c *gin.Context) {
	page := struct {
		Page  int `form:"page" binding:"required,gt=0"`
		Limit int `form:"limit" binding:"required,max=200"`
	}{}

	if err := c.BindQuery(&page); err != nil {
		logger.Logger().Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "分页参数有误！"})
		return
	}

	types := make([]int8, 0)
	for _, v := range c.QueryArray("type") {
		t, _ := strconv.Atoi(v)
		types = append(types, int8(t))
	}

	mapID2Department, err := department.MapID2Department()
	if err != nil {
		logger.Logger().Warn("获取部门ID到部门名字的映射失败:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	users, count, err := user.List((page.Page-1)*page.Limit, page.Limit, types...)
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
			ID:      u.ID,
			Account: u.Account,
			Name:    u.Name,
			//Type:       u.Type,
			Phone:          u.Phone,
			Email:          u.Email,
			WeChat:         u.WeChat,
			DepartmentName: mapID2Department[u.Department].Name,
			DepartmentType: mapID2Department[u.Department].Type,
			Valid:          u.Valid,
			CreatedAt:      u.CreatedAt,
			ModifiedAt:     u.ModifiedAt,
		})
	}
	userBriefInfos.Count = count

	c.JSON(http.StatusOK, &userBriefInfos)
}

func Info(c *gin.Context) {

}

func ListWithType(c *gin.Context) {
	page := struct {
		Type []int8 `json:"type"`
	}{}

	if err := c.BindJSON(&page); err != nil {
		logger.Logger().Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "分页参数有误！"})
		return
	}

	users, _, err := user.List(0, 0, page.Type...)
	if err != nil {
		logger.Logger().Warn("query users error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	type ub struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
	ubs := make([]ub, 0)

	for _, u := range users {
		ubs = append(ubs, ub{
			ID:   u.ID,
			Name: u.Name,
		})
	}

	c.JSON(http.StatusOK, &ubs)
}
