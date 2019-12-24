package constant

// TODO 这里应该可以整合到一起
// 用户/部门类型
const (
	TypeUserPlaceholder   = iota
	TypeUserAdministrator // 1:管理员用户，系统只有一个，不能添加
	TypeUserCity          // 2:市级单位
	TypeUserDistrict      // 3:市级各辖区单位
	TypeUserSupervised    // 4:受监管企业单位
	TypeUserSupport       // 5:签约技术支持/安全服务单位
	TypeUserMax
)

var departmentType = []string{"", "管理员", "市级单位", "市级各辖区单位", "受监管企业单位", "签约技术支持/安全服务单位"}

func GetDepartmentTypeName(id int8) string {
	if id <= TypeUserAdministrator {
		return ""
	}
	return departmentType[id]
}

// 用户登录相关
const (
	KeyUserID = "user_id"
)

// 数据库相关
const (
	PageMaxItem = 200
)

// 事件类型: 1:违处信息、2:网络攻击、3:恶意软件、4:信息泄露、5:安全威胁/漏洞
const (
	TypeEventPlaceholder = iota
	TypeEventIllegal
	TypeEventAttack
	TypeEventMalware
	TypeEventLeakage
	TypeEventFlaw
	TypeEventMax
)

// 事件处理状态: 1:待办事项、2:处理中事项、3:已完成事项
const (
	StatusPlaceholder = iota
	StatusEventTodo
	StatusEventDealing
	StatusEventFinished
)

// 事务类型：1:安全事件通报、2:热点事件发布
const (
	TypeTransPlaceholder = iota
	TypeTransReport
	TypeTransHot
	TypeTransMax
)
