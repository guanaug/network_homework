# 高级计算机网络结课大作业
* version: v1.0.0 
* author: 官嘉林
* created_at: 2019.12.02 

## 接口定义
* 用户相关
    > POST `/user`
    * 作用: 添加用户
    * 参数: 
        * `application/json`
        * account: string(required)，登录账号，长度1-16，数字或字母
        * name: string(required)，名字，长度1-16
        * type: int8(required)，用户类型，2:市级单位 3:市级各辖区单位 4:受监管企业单位 5:签约技术支持/安全服务单位
        * phone: string(required)，手机号码，长度11，正则规则: `^1[3-9]\d{9}$`
        * email: string(required)，邮箱地址，符合邮件命名规则
        * wechat: string(required)，微信号，符合微信号规则
        * department: int64(required)，部门号，见 `部门列表: GET /department`
        * password: string(required)，密码，长度8-32
    * 返回:
        * 200
        * 错误: 
            * 400: 用户信息填写有误
            * 500: 服务器内部错误
            
    > DELETE `/user/:id`
    * 作用: 删除用户
    * 参数: 
        * 无
    * 返回:
        * 200
        * 错误:
            * 400: 用户id有误
            * 500: 服务器内部错误
         
    > PUT `/user`
    * 作用: 修改用户信息
    * 参数: 
        * `application/json`
        * id: int64(required)，用户id
        * account: string(optional)，登录账号，长度1-16，数字或字母
        * name: string(optional)，名字，长度1-16
        * type: int8(optional)，用户类型，2:市级单位 3:市级各辖区单位 4:受监管企业单位 5:签约技术支持/安全服务单位
        * phone: string(optional)，手机号码，长度11，正则规则: `^1[3-9]\d{9}$`
        * email: string(optional)，邮箱地址，符合邮件命名规则
        * wechat: string(optional)，微信号，符合微信号规则
        * department: int64(optional)，部门号，见 `部门列表: GET /department`
        * password: string(optional)，密码，长度8-32
    * 返回:
        * 200
        * 错误: 
            * 400: 用户信息填写有误
            * 500: 服务器内部错误

    > GET `/user`
    * 作用: 获取用户列表
    * 参数:
        * `application/json`
        * offset: int(required)，分页起始值
        * limit: int(required)，每页数量，1-200
    * 返回:
        * 200
            * `application/json`
            * count: int(required)，用户总数
            * users: Array, Object, 用户信息
                * id: int64(required)，用户id
                * account: string(required)，登录账号，长度1-16，数字或字母
                * name: string(required)，名字，长度1-16
                * type: int8(required)，用户类型，2:市级单位 3:市级各辖区单位 4:受监管企业单位 5:签约技术支持/安全服务单位
                * phone: string(required)，手机号码，长度11，正则规则: `^1[3-9]\d{9}$`
                * email: string(required)，邮箱地址，符合邮件命名规则
                * wechat: string(required)，微信号，符合微信号规则
                * department: int64(required)，部门号，见 `部门列表: GET /department`
                * password: string(required)，密码，长度8-32
        * 错误:
            * 400: 分页参数有误
            * 500: 服务器内部错误

* 部门相关
     > POST `/department`
     * 作用: 添加部门
     * 参数: 
         * `application/json`
         * name: string(required)，部门名称，长度1-64
         * type: int8(required)，部门类型，2:市级单位 3:市级各辖区单位 4:受监管企业单位 5:签约技术支持/安全服务单位
         * address: string(required)，部门地址，长度1-128
         * owner: string(required)，部门负责人姓名，长度1-16
         * owner_contact: string(required)，部门负责人手机号码，长度11，正则规则: `^1[3-9]\d{9}$`
         * admin: string(require)，网安管理员姓名，长度1-16
         * admin_contact: string(required), 网安管理员手机号码，长度11，正则规则同部门负责人手机号码
     * 返回:
         * 200
         * 错误: 
             * 400: 用户信息填写有误
             * 500: 服务器内部错误
             
     > DELETE `/department/:id`
     * 作用: 删除部门
     * 参数: 
         * 无
     * 返回:
         * 200
         * 错误:
             * 400: 部门id有误
             * 500: 服务器内部错误
          
     > PUT `/department`
     * 作用: 修改部门信息
     * 参数: 
         * `application/json`
         * id: int64(required)，部门ID
         * name: string(optional)，部门名称，长度1-64
         * type: int8(optional)，部门类型，2:市级单位 3:市级各辖区单位 4:受监管企业单位 5:签约技术支持/安全服务单位
         * address: string(optional)，部门地址，长度1-128
         * owner: string(optional)，部门负责人姓名，长度1-16
         * owner_contact: string(optional)，部门负责人手机号码，长度11，正则规则: `^1[3-9]\d{9}$`
         * admin: string(optional)，网安管理员姓名，长度1-16
         * admin_contact: string(optional), 网安管理员手机号码，长度11，正则规则同部门负责人手机号码
     * 返回:
         * 200
         * 错误: 
             * 400: 用户信息填写有误
             * 500: 服务器内部错误
 
     > GET `/department`
     * 作用: 获取部门列表
     * 参数:
        * `application/json`
        * offset: int(required)，分页起始值
        * limit: int(required)，每页数量，1-200
     * 返回:
         * 200
             * `application/json`
             * count: int(required)，总部门数
             * departments: Array, Object，部门信息
                 * id: int64(required)，部门ID
                 * name: string(required)，部门名称，长度1-64
                 * type: int8(required)，部门类型，2:市级单位 3:市级各辖区单位 4:受监管企业单位 5:签约技术支持/安全服务单位
                 * address: string(required)，部门地址，长度1-128
                 * owner: string(required)，部门负责人姓名，长度1-16
                 * owner_contact: string(required)，部门负责人手机号码，长度11，正则规则: `^1[3-9]\d{9}$`
                 * admin: string(require)，网安管理员姓名，长度1-16
                 * admin_contact: string(required), 网安管理员手机号码，长度11，正则规则同部门负责人手机号码
         * 错误:
             * 400: 分页参数有误
             * 500: 服务器内部错误
     
     > GET `/department/:id`
     * 作用: 获取部门详细信息
     * 参数:
         * 无
     * 返回:
         * 200
             * `application/json`
             * id: int64(required)，部门ID
             * name: string(required)，部门名称，长度1-64
             * type: int8(required)，部门类型，2:市级单位 3:市级各辖区单位 4:受监管企业单位 5:签约技术支持/安全服务单位
             * address: string(required)，部门地址，长度1-128
             * owner: string(required)，部门负责人姓名，长度1-16
             * owner_contact: string(required)，部门负责人手机号码，长度11，正则规则: `^1[3-9]\d{9}$`
             * admin: string(require)，网安管理员姓名，长度1-16
             * admin_contact: string(required), 网安管理员手机号码，长度11，正则规则同部门负责人手机号码
         * 错误:
             * 400: 分页参数有误
             * 500: 服务器内部错误