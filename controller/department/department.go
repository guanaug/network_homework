package department

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"net/http"
	"network/global/logger"
	"network/model/department"
	"strconv"
	"time"
)

type Department struct {
	ID           int64     `json:"id,omitempty"`
	Name         string    `json:"name,omitempty" binding:"required,max=64"`
	Address      string    `json:"address,omitempty" binding:"required,max=128"`
	Type         int8      `json:"type,omitempty" binding:"required,departmentType"`
	Owner        string    `json:"owner,omitempty" binding:"required,max=16"`
	OwnerContact string    `json:"owner_contact,omitempty" binding:"required,phone"`
	Admin        string    `json:"admin,omitempty" binding:"required,max=16"`
	AdminContact string    `json:"admin_contact,omitempty" binding:"required,phone"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	ModifiedAt   time.Time `json:"modified_at,omitempty"`
	DeletedAt    time.Time `json:"deleted_at,omitempty"`
}

func Add(c *gin.Context) {
	departmentInfo := Department{}

	if err := c.BindJSON(&departmentInfo); err != nil {
		logger.Logger().Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "部门信息填写有误！"})
		return
	}

	depart := &department.Department{
		Name:         departmentInfo.Name,
		Address:      departmentInfo.Address,
		Type:         departmentInfo.Type,
		Owner:        departmentInfo.Owner,
		OwnerContact: departmentInfo.OwnerContact,
		Admin:        departmentInfo.Admin,
		AdminContact: departmentInfo.AdminContact,
	}

	if err := depart.Insert(); err != nil {
		logger.Logger().Warn("department add error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": depart.ID})
}

func Delete(c *gin.Context) {
	departID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "部门ID有误！"})
		return
	}

	depart := &department.Department{
		ID: departID,
	}

	if err := depart.Delete(); err != nil {
		logger.Logger().Warn("department delete error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

// TODO modified_at时间还不能自动改
func Modify(c *gin.Context) {
	departmentInfo := struct {
		ID           int64  `json:"id" binding:"required,gt=0"`
		Name         string `json:"name" binding:"max=64"`
		Address      string `json:"address" binding:"max=128"`
		Type         int8   `json:"type" binding:"departmentType"`
		Owner        string `json:"owner" binding:"max=16"`
		OwnerContact string `json:"owner_contact" binding:"phone"`
		Admin        string `json:"admin" binding:"max=16"`
		AdminContact string `json:"admin_contact" binding:"phone"`
	}{}

	if err := c.BindJSON(&departmentInfo); err != nil {
		logger.Logger().Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "部门信息填写有误！"})
		return
	}

	depart := &department.Department{
		ID: departmentInfo.ID,
	}
	// TODO ugly code, must be reconstruct
	if len(departmentInfo.Name) > 0 {
		depart.Name = departmentInfo.Name
	}
	if len(departmentInfo.Address) > 0 {
		depart.Address = departmentInfo.Address
	}
	if departmentInfo.Type > 0 {
		depart.Type = departmentInfo.Type
	}
	if len(departmentInfo.Owner) > 0 {
		depart.Owner = departmentInfo.Owner
	}
	if len(departmentInfo.OwnerContact) > 0 {
		depart.OwnerContact = departmentInfo.OwnerContact
	}
	if len(departmentInfo.Admin) > 0 {
		depart.Admin = departmentInfo.Admin
	}
	if len(departmentInfo.AdminContact) > 0 {
		depart.AdminContact = departmentInfo.AdminContact
	}

	if err := depart.Update(); err != nil {
		logger.Logger().Warn("department add error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func List(c *gin.Context) {
	page := struct {
		Type  int8 `form:"type"`
		Page  int  `form:"page" binding:"required,gt=0"`
		Limit int  `form:"limit" binding:"required,max=200"`
	}{}

	if err := c.BindQuery(&page); err != nil {
		logger.Logger().Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "分页参数有误！"})
		return
	}

	departs, count, err := department.List((page.Page-1)*page.Limit, page.Limit, page.Type)
	if err != nil {
		logger.Logger().Warn("query departments error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	departBriefInfos := struct {
		Count int          `json:"count"`
		List  []Department `json:"departments"`
	}{}

	for _, depart := range departs {
		departBriefInfos.List = append(departBriefInfos.List, Department{
			ID:           depart.ID,
			Name:         depart.Name,
			Address:      depart.Address,
			Type:         depart.Type,
			Owner:        depart.Owner,
			OwnerContact: depart.OwnerContact,
			Admin:        depart.Admin,
			AdminContact: depart.AdminContact,
			CreatedAt:    depart.CreatedAt,
			ModifiedAt:   depart.ModifiedAt,
		})
	}
	departBriefInfos.Count = count

	c.JSON(http.StatusOK, &departBriefInfos)
}

func Info(c *gin.Context) {
	departID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "部门ID有误！"})
		return
	}

	depart := &department.Department{
		ID: departID,
	}

	*depart, err = depart.Info()
	if pg.ErrNoRows == err {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error:": "未找到部门，请确认部门ID是否存在"})
		return
	}
	if err != nil {
		logger.Logger().Warn("获取部门信息失败:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	departInfo := Department{
		ID:           depart.ID,
		Name:         depart.Name,
		Address:      depart.Address,
		Type:         depart.Type,
		Owner:        depart.Owner,
		OwnerContact: depart.OwnerContact,
		Admin:        depart.Admin,
		AdminContact: depart.AdminContact,
		CreatedAt:    depart.CreatedAt,
		ModifiedAt:   depart.ModifiedAt,
	}

	c.JSON(http.StatusOK, &departInfo)
}
