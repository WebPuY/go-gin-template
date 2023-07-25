# go-gin-template

## 项目开发

### 启动服务
```bash
go run ./src/main.go
```

## 目录说明

- 入口

  * `main.go`：引入配置，启动主程序，引入各种全局服务
  * `app.controller.go`：主程序根控制器
  * `app.service.go`: 主程序根服务
  * `app.config.go`：配置文件
  * `app.environment.go：`全局环境变量

- 目录

  * `constants`：系统常量
  * `errors`: 自定义的错误状态
  * `filters`：异常处理
  * `interfaces`：接口
  * `middlewares`：自定义中间件       路由处理程序 之前 调用的函数
  * `guards`：自定义的守卫            用于给特定请求 授权
  * `interceptors`: 自定义的拦截器    用于在函数执行之前/之后绑定额外的逻辑
  * `modules`: 应用模块
      * module文件：                  注册当前模块
      * controller文件：              业务层，路由文件，定义接口，处理业务，处理传入的请求 和 返回响应
      * service文件：                 数据层，负责数据存储和检索
      * entity文件：                  访问对象层，数据库字段文件
      * DTO文件：校验文件，            如校验参数类型和参数长度（可选）
  * `pipes`：管道                     用于数据转换 和 数据验证，如果验证成功继续传递; 验证失败则抛出异常;
  * `processors`：三方库连接
  * `transforms`：转换工具

- 环境变量

  `.env` 

- 请求处理流程

  1. `request`：收到请求
  2. `middleware`：中间件过滤（跨域、来源校验等处理）
  3. `guard`：守卫过滤（鉴权）
  4. `interceptor:before`：数据流拦截器
  5. `pipe`：参数提取（校验）器
  6. `controller`：业务控制器
  7. `service`：业务服务
  8. `interceptor:after`：数据流拦截器（格式化数据、错误）
  9. `filter`：捕获以上所有流程中出现的异常，如果任何一个环节抛出异常，则返回错误



## 常见问题

1. 

2. 
  
3. 