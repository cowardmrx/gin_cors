#gin_cors

##说明
gin框架的cors跨域处理中间件

##使用
go get github.com/cowardmrx/gin_cors

```go
r := gin.Default()

// 创建cors配置 只要是string类型的一律按照 , 分割，只有AccessControlAllowCredentials是字符串类型的bool值
cors := &Cors{
	// 允许的域名
    AccessControlAllowOrigins: []string{
    "http://localhost",
    },
	// 允许的请求头 使用 ',' 分割
    AccessControlAllowHeaders: "Authorization",
	// 允许的请求方式 使用 ',' 分割
    AccessControlAllowMethods: "PUT,POST,GET,DELETE",
	// 允许暴露的请求头 使用 ',' 分割 
    AccessControlExposeHeaders: "Authorization",
    // 允许凭证记录 true | false 
	AccessControlAllowCredentials: 'true'
}

    

r.POST("/te", func(c *gin.Context) {
c.JSON(200, gin.H{
    "data": "this is data",
    })
})

r.Run(":7878")

}
```