#gin_cors

##说明
gin框架的cors跨域处理中间件

##使用
go get git@github.com:cowardmrx/gin_cors.git

```go
r := gin.Default()

// 创建cors配置 只要是string类型的一律按照 , 分割，只有AccessControlAllowCredentials是字符串类型的bool值
cors := &Cors{
    AccessControlAllowOrigins: []string{
    "http://localhost",
    },
    AccessControlAllowMethods: "PUT,POST,GET,DELETE",
}

    

r.POST("/te", func(c *gin.Context) {
c.JSON(200, gin.H{
    "data": "this is data",
    })
})

r.Run(":7878")

}
```