package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"net/http"
	"network/global/constant"
	"network/global/logger"
	"network/global/pgdb"
	"network/global/session"
	"network/model/department"
	"network/model/transaciton"
	"network/model/user"
	"time"
)

type Transaction struct {
	ID                int64     `json:"id,omitempty"`
	Publisher         int64     `json:"publisher,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	Type              int8      `json:"type,omitempty"`
	Status            int8      `json:"status,omitempty"`
	Detail            string    `json:"detail,omitempty"`
	TranType          int8      `json:"tran_type,omitempty"`
	HandlerDepartment int64     `json:"handler_department,omitempty"`
	Handler           int64     `json:"handler,omitempty"`
	ModifiedAt        time.Time `json:"modified_at,omitempty"`
}

func Add(c *gin.Context) {
	trans := Transaction{}

	if err := c.BindJSON(&trans); err != nil {
		logger.Logger().Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "事件添加失败，请检查参数"})
		return
	}

	u, err := user.OneByAccount(session.GetUser(c))
	if err != nil {
		logger.Logger().Warn("get user error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	model := transaciton.Transaction{
		Publisher:         u.ID,
		Type:              trans.Type,
		Status:            constant.StatusEventTodo,
		Detail:            trans.Detail,
		TranType:          trans.Type,
		HandlerDepartment: trans.HandlerDepartment,
	}

	if err := model.Insert(); err != nil {
		logger.Logger().Warn("insert transaction error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": model.ID})
}

func Modified(c *gin.Context) {
	trans := struct {
		ID      int64 `json:"id,omitempty" binding:"required"`
		Status  int8  `json:"status,omitempty"`
		Handler int64 `json:"handler,omitempty"`
	}{}

	if err := c.BindJSON(&trans); err != nil {
		logger.Logger().Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "事件修改失败，请检查参数"})
		return
	}

	model := transaciton.Transaction{
		ID: trans.ID,
	}
	// 修改事件状态
	if trans.Status > 0 {
		model.Status = trans.Status
	}
	// 指定单位处理事件
	if trans.Handler > 0 {
		model.Handler = trans.Handler
	}

	if err := model.Update(); err != nil {
		logger.Logger().Warn("update transaction error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func List(c *gin.Context) {
	trans := struct {
		Publisher         int64     `form:"publisher,omitempty"`
		Begin             time.Time `form:"Begin,omitempty"`
		End               time.Time `form:"end,omitempty"`
		Type              int8      `form:"type,omitempty"`
		Status            int8      `form:"status,omitempty"`
		TranType          int8      `form:"tran_type,omitempty"`
		HandlerDepartment int64     `form:"handler_department,omitempty"`
		Handler           int64     `form:"handler,omitempty"`
		Page              int       `form:"page,omitempty"`
		Limit             int       `form:"limit,omitempty"`
	}{}

	if err := c.BindQuery(&trans); err != nil {
		logger.Logger().Debug(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "查询事件失败，请检查参数"})
		return
	}

	mapID2Depart, err := department.MapID2Department()
	if err != nil {
		logger.Logger().Warn("query department error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	mapID2User, err := user.MapID2User()
	if err != nil {
		logger.Logger().Warn("query department error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	models := make([]transaciton.Transaction, 0)
	model := pgdb.DB().Model(&models)
	// 根据发布者进行查询
	if trans.Publisher > 0 {
		model.Where("publisher = ?", trans.Publisher)
	}
	// 根据事件类型进行查询
	if trans.Type > 0 {
		model.Where("publisher = ?", trans.Type)
	}
	// 根据事件状态进行查询
	if trans.Status > 0 {
		model.Where("status = ?", trans.Status)
	}
	// 根据事务类型进行查询
	if trans.TranType > 0 {
		model.Where("tran_type = ?", trans.TranType)
	}
	// 根据辖区进行查询
	if trans.HandlerDepartment > 0 {
		model.Where("handler_department = ?", trans.HandlerDepartment)
	}
	// 根据处理单位进行查询
	if trans.Handler > 0 {
		model.Where("handler = ?", trans.Handler)
	}
	// 根据事件进行查询
	if !trans.Begin.IsZero() && !trans.End.IsZero() {
		model.Where("created_at >= ? and created_at <= ?", trans.Begin, trans.End)
	}
	count, err := model.Offset((trans.Page - 1) * trans.Limit).Limit(trans.Limit).Order("id asc").SelectAndCount()
	if err != nil {
		logger.Logger().Warn("query transaction error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	type response struct {
		ID                int64     `json:"id,omitempty"`
		Publisher         string    `json:"publisher,omitempty"`
		CreatedAt         time.Time `json:"created_at,omitempty"`
		Type              int8      `json:"type,omitempty"`
		Status            int8      `json:"status,omitempty"`
		Detail            string    `json:"detail,omitempty"`
		TranType          int8      `json:"tran_type,omitempty"`
		HandlerDepartment string    `json:"handler_department,omitempty"`
		Handler           string    `json:"handler,omitempty"`
		ModifiedAt        time.Time `json:"modified_at,omitempty"`
	}

	transactions := struct {
		Count        int        `json:"count"`
		Transactions []response `json:"transactions"`
	}{}
	transactions.Count = count

	for _, tran := range models {
		transactions.Transactions = append(transactions.Transactions, response{
			ID:                tran.ID,
			Publisher:         mapID2User[tran.Publisher].Name,
			CreatedAt:         tran.CreatedAt,
			Type:              tran.Type,
			Status:            tran.Status,
			Detail:            tran.Detail,
			TranType:          tran.TranType,
			HandlerDepartment: mapID2Depart[tran.HandlerDepartment].Name,
			Handler:           mapID2User[tran.Handler].Name,
			ModifiedAt:        tran.ModifiedAt,
		})
	}

	c.JSON(http.StatusOK, &transactions)
}

func Statistic(c *gin.Context) {
	type Stat []struct {
		Status int8 `sql:"status" json:"status"`
		Count  int  `json:"count"`
	}
	stat := Stat{}

	err := transaciton.New().Model().Column("transaction.status").
		ColumnExpr("count(*)").
		Group("transaction.status").Select(&stat)
	if err != nil && err != pg.ErrNoRows {
		logger.Logger().Warn("statistic transaction error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var result struct {
		Statistic Stat `json:"statistic"`
	}
	result.Statistic = stat

	c.JSON(http.StatusOK, &result)
}
