package constant

// 用户/部门类型
const (
	TypeUserPlaceholder   = iota
	TypeUserAdministrator // 1:管理员用户，系统只有一个，不能添加
	TypeUserCity          // 2:市级单位
	TypeUserDistrict      // 3:市级各辖区单位
	TypeUserSupervised    // 4:受监管企业单位
	TypeUserSupport       // 签约技术支持/安全服务单位
	TypeUserMax
)
