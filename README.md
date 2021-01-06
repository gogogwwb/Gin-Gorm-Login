# Gin-Gorm-Login

项目主要使用GIN和Gorm实现的登录功能

数据库使用的mysql，gorm的连接mysql方法如下：

db,err=gorm.Open("mysql","用户名:密码@(127.0.0.1:3306)/数据库名?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		panic(err)
	}
