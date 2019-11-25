package department

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"network/global/logger"
	"network/model/department"
	"strconv"
	"time"
)

type Department struct {
	ID           int64     `json:"id, omitempty"`
	Name         string    `json:"name, omitempty" binding:"required, max=64"`
	Address      string    `json:"address, omitempty" binding:"required, max=128"`
	Type         int8      `json:"type, omitempty" binding:"required, departmentType"`
	Owner        string    `json:"owner, omitempty" binding:"required, max=16"`
	OwnerContact string    `json:"owner_contact, omitempty" binding:"required, phone, len=11"`
	Admin        string    `json:"admin, omitempty" binding:"required, max=16"`
	AdminContact string    `json:"admin_contact, omitempty" binding:"required, phone, len=11"`
	CreatedAt    time.Time `json:"created_at, omitempty"`
	ModifiedAt   time.Time `json:"modified_at, omitempty"`
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

	if err := depart.Add(); err != nil {
		logger.Logger().Warn("department add error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
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

func Modified(c *gin.Context) {
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

	if err := depart.Update(); err != nil {
		logger.Logger().Warn("department add error:", err)
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

	departs, count, err := department.New().List(page.Offset, page.Limit)
	if err != nil {
		logger.Logger().Warn("query departments error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
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
