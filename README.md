# 高级计算机网络结课大作业
* version: v1.0.0 
* author: dogod
* created_at: 2019.12.02 

## 接口定义
* 用户相关
    > POST `/user`
    * 作用: 添加用户
    * 参数: 
        * `application/json`
        * account: string(required)，登录账号，长度1-16，数字或字母
        * name: string(required)，名字，长度1-16
        * // 不知道有啥用，暂时弃用这个字段，type: int8(required)，用户类型，2:市级单位 3:市级各辖区单位 4:受监管企业单位 5:签约技术支持/安全服务单位
        * phone: string(required)，手机号码，长度11，正则规则: `^1[3-9]\d{9}$`
        * email: string(required)，邮箱地址，符合邮件命名规则
        * wechat: string(required)，微信号，符合微信号规则
        * department: int64(required)，部门号，见 `部门列表: GET /department`
        * password: string(required)，密码，长度8-32
    * 返回:
        * 200
            * `application/json`
            * id: int64(required)，用户ID
        * 错误: 
            * 400: 用户信息填写有误
                * `application/json`
                * error: string(required)，错误消息
            * 500: 服务器内部错误
            
    > DELETE `/user/:id`
    * 作用: 删除用户
    * 参数: 
        * 无
    * 返回:
        * 200
        * 错误:
            * 400: 用户id有误
                * `application/json`
                * error: string(required)，错误消息
            * 500: 服务器内部错误
         
    > PUT `/user`
    * 作用: 修改用户信息
    * 参数: 
        * `application/json`
        * id: int64(required)，用户id
        * name: string(optional)，名字，长度1-16
        * // type: int8(optional)，用户类型，2:市级单位 3:市级各辖区单位 4:受监管企业单位 5:签约技术支持/安全服务单位
        * phone: string(optional)，手机号码，长度11，正则规则: `^1[3-9]\d{9}$`
        * email: string(optional)，邮箱地址，符合邮件命名规则
        * wechat: string(optional)，微信号，符合微信号规则
        * department: int64(optional)，部门号，见 `部门列表: GET /department`
        * valid: bool(optional)，启用/禁用用户
        * password: string(optional)，密码，长度8-32
    * 返回:
        * 200
        * 错误: 
            * 400: 用户信息填写有误
                * `application/json`
                * error: string(required)，错误消息
            * 500: 服务器内部错误

    > GET `/user`
    * 作用: 获取用户列表
    * 参数:
        * page: int(required)，第几页
        * limit: int(required)，每页数量，1-200
    * 返回:
        * 200
            * `application/json`
            * count: int(required)，用户总数
            * users: Array, Object, 用户信息
                * id: int64(required)，用户id
                * account: string(required)，登录账号，长度1-16，数字或字母
                * name: string(required)，名字，长度1-16
                * // type: int8(required)，用户类型，2:市级单位 3:市级各辖区单位 4:受监管企业单位 5:签约技术支持/安全服务单位
                * phone: string(required)，手机号码，长度11，正则规则: `^1[3-9]\d{9}$`
                * email: string(required)，邮箱地址，符合邮件命名规则
                * wechat: string(required)，微信号，符合微信号规则
                * department_name: string(required)，部门名称，见 `部门列表: GET /department`
                * department_type: int8(required)，部门类型，见 `添加部门: POST /department`
                * password: string(required)，密码，长度8-32
        * 错误:
            * 400: 分页参数有误
                * `application/json`
                * error: string(required)，错误消息
            * 500: 服务器内部错误
            
***

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
            * `application/json`
            * id: int64(required)，部门ID
         * 错误: 
             * 400: 用户信息填写有误
                 * `application/json`
                 * error: string(required)，错误消息
             * 500: 服务器内部错误
             
     > DELETE `/department/:id`
     * 作用: 删除部门
     * 参数: 
         * 无
     * 返回:
         * 200
         * 错误:
             * 400: 部门id有误
                * `application/json`
                * error: string(required)，错误消息
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
                * `application/json`
                * error: string(required)，错误消息
             * 500: 服务器内部错误
 
     > GET `/department`
     * 作用: 获取部门列表
     * 参数:
        * page: int(required)，第几页
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
                * `application/json`
                * error: string(required)，错误消息
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
                * `application/json`
                * error: string(required)，错误消息
             * 500: 服务器内部错误
             
***

* 会话相关
    > POST `/session`
    * 作用：用户登录
    * 参数：
        * `application/json`
        * account: string(required)，用户账号，规则同`POST /user`
        * password: string(required)，用户密码，规则同`POST /user`
    * 返回:
        * 200：
            * `application/json`
            * id: int64(required)，用户ID，规则同`POST /user`
            * name: string(required)，用户名称
            * type: int8(required)，用户类型 1:管理员，-1:其他
        * 错误：
            * 400: 请输入正确的账号或密码
                * `application/json`
                * error: string(required)，错误消息
            * 401: 账号或密码错误
                * `application/json`
                * error: string(required)，错误消息
            * 500: 服务器内部错误
        
    > DELETE `/session/:id`
    * 作用：注销登录
    * 参数：
        * 无
    * 返回:
        * 200：
            * `application`
            * id: string(required)，用户ID
            * name: string(required)，用户名称
        * 错误：
            * 500: 服务器内部错误
            
    > GET `/session/log`
    * 作用: 获取用户登录日志
    * 参数:
        * page: int(required)，第几页
        * limit: int(required)，每页数量，1-200
    * 返回:
        * 200:
            * `application/json`
            * id: int64(required)，编号
            * user_account: string(required)，用户账号
            * user_name: string(required)，用户姓名
            * ip: string(required)，登录IP
            * created_at: time(required)，登录时间
            
***

* 事务相关
    > POST `/transaction`
    * 作用：添加事务
    * 参数:
        * `application/json`
        * type: int8(required)，事件类型: 1:违处信息、2:网络攻击、3:恶意软件、4:信息泄露、5:安全威胁/漏洞
        * detail: string(required)，具体事件描述
        * tranType: string(required)，事务类型：1:安全事件通报、2:热点事件发布
        * handler_department: int64(required)，辖区，见`添加部门 POST /department`
    * 返回:
        * 200:
            * `application/json`
            * id: int64(required)，事件编号
         * 错误:
            * 400:
                * `application/json`
                * error: string(required)，错误消息
            * 500:
            
    > PUT `/transaction`
    * 作用：修改事务信息
    * 参数:
        * `application/json`
        * id: int64(required)，事件编号
        * status: int8(optional)，事件状态，1:待办事项、2:处理中事项、3:已完成事项
        * handler: int64(optional)，处理事件用户
    * 返回：
        * 200
        * 错误:
           * 400:
               * `application/json`
               * error: string(required)，错误消息
           * 500:
           
    > GET '/transaction'
    * 作用:根据查询条件查询事务
    * 参数:
        * `application/json`
        * publisher: int64(optional)，根据发布者查询事务
        * begin: time(optional)，根据时间段查询事务，必须和end成对出现
        * end: time(optional)
        * type: int8(optional)，根据事件类型查询事务
        * status: int8(optional)，根据事件状态查询事务
        * tran_type: int8(optional)，根据事务类型查询事务
        * handler_department：int64(optional)，根据辖区查询事务
        * handler: int64(optional)，根据事件处理用户查询事务
        * page: int(required)，第几页
        * limit: int(required)，分页每页大小
    * 返回:
        * 200
            * `application/json`
            * id: int64(required)，事件编号
            * publisher: int64(required)，发布者ID
            * crated_at: time(required)，事务创建时间
            * type: int8(required)，事件类型: 1:违处信息、2:网络攻击、3:恶意软件、4:信息泄露、5:安全威胁/漏洞
            * status: int8(required)，事件状态，1:待办事项、2:处理中事项、3:已完成事项
            * detail: string(required)，具体事件描述
            * tranType: string(required)，事务类型：1:安全事件通报、2:热点事件发布
            * handler_department: int64(required)，辖区，见`添加部门 POST /department`
            * handler: int64(optional)，处理事件用户

    > GET `/transaction/statistic`
    * 作用: 获取事务统计信息
    * 参数:
        无
    * 返回：
        * 200
            * `application/json`
            * status: int64(required)，事务状态
            * count: int(required)，事务状态对应的事务数量