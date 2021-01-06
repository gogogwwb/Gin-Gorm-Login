package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

/*type User struct {
	Name string
	Age int
}
func Say(w http.ResponseWriter,r *http.Request)  {
	t,err:=template.ParseFiles("./test.html")
	u:=User{
		"wwb",
		15,
	}
	if err!=nil{
		fmt.Print("err")
		return
	}
	t.Execute(w,u)
}

func main() {
	http.HandleFunc("/",Say)
	err :=http.ListenAndServe(":9000",nil)
	if err!=nil{
		fmt.Print("err")
		return
	}
*/
type Users struct{
	Id  uint  `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Pwd string `json:"pwd"`
}
var db *gorm.DB
func init()  {
	var err error
	db,err=gorm.Open("mysql","root:123456@(127.0.0.1:3306)/user?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		panic(err)
	}
}

func main() {
	r:=gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.tmpl",gin.H{
			"err":nil,
		})
	})
	r.POST("/login", func(c *gin.Context) {
		var name,pwd string
		name=c.PostForm("name")
		fmt.Print(name)
		pwd=c.PostForm("pwd")
		users:=new(Users)
		db.Where("name=? and pwd=?",name,pwd).Find(&users)
		fmt.Printf("%#v\n", users)
		if users.Name!=""{
			c.HTML(http.StatusOK,"wel.tmpl",gin.H{
				"name":name,
			})
		}else{
			c.HTML(http.StatusOK,"login.tmpl",gin.H{
				"err":"账号或密码错误",
			})
		}
	})
	defer db.Close()
	r.Run()
}



