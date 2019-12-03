package main

import(
	
	// "fmt"
	// "reflect"
	// "github.com/spandan12/ginapi/models"
	"github.com/spandan12/ginapi/router"
)


func main(){
	
	r := router.New()
	// r.GET("/ping", getController)
	r.Run()
	// models.InsertOne("spandan",23,"hetauda")
	// fmt.Println(models.ReadAll())
}