#Creeper 苦力怕

起个名字很难

------------
![creeper](https://github.com/lvxin0315/creeper/blob/master/creeper.jpeg "creeper")
------------

##### **接口中继**

- 1.应用管理-~~创建~~，删除，~~secret重置~~
- 2.平台接口接入鉴权
- 3.设置access_token与app关系
- 4.用户登录认证方式
- 5.平台api管理
- 6.平台api与应用关联（操作权限）
- 7.~~gin做http服务，独立控制器目录~~
- 8.http转发
- 9.对外http api
- 10.管理http api
- 11.表单验证

------------

##### **内容记录**
- app secret 32位
- 对称加密使用 AES/CBC/PKCS7Padding
- access_token 32位
- swagger文档 https://github.com/swaggo/gin-swagger
- validate 数据验证 https://github.com/go-playground/validator 

------------

##### **接口文档使用 gin-swagger**
1. 对应目录 creeper_http_api:
		使用 swag init -g creeper_http_api/doc.go -o creeper_http_docs
		访问地址 http://127.0.0.1:8081/swagger/index.html

------------

##### **目录结构**
- creeper_http_api 平台接口
- runner 项目启动项文件
- orm_model gorm模拟目录，不建议放业务代码
- etc 配置文件
- db_conn 处理db连接函数
- service 纯业务逻辑代码，自己控制db处理，返回值统一 output
- cache 缓存或全局存储内容目录