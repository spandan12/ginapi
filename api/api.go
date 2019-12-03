package api

import(
	"github.com/gin-gonic/gin"
	"github.com/spandan12/ginapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)

func Create(c *gin.Context){
	var u models.User
	c.BindJSON(&u)
	filter := bson.M{
		"name": u.Name,
		"age": u.Age,
		"address": u.Address,
	  }
	models.InsertOne(filter)
	c.JSON(200, gin.H{
		"data inserted": u,
	})
}

func FetchAll(c *gin.Context){
	data := models.ReadAll(bson.D{{}})
	c.JSON(200, gin.H{
		"data": data,
	})
}

func Fetch(c *gin.Context){
	id := c.Param("id")
	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", oid}}
	data := models.ReadOne(filter)
	c.JSON(200, gin.H{
		"data": data,
	})
}

func Update(c *gin.Context){
	id := c.Param("id")
	var u models.User
	c.BindJSON(&u)
	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": oid}

	update := bson.M{
		`$set`: bson.M{
		  "name": u.Name,
		  "age": u.Age,
		  "address": u.Address,
		},
	  }
	models.UpdateOne(filter,update)
	c.JSON(200, gin.H{
		"updated": u,
	})
}

func Delete(c *gin.Context){
	id := c.Param("id")
	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": oid}
	models.DeleteOne(filter)
	c.JSON(200, gin.H{
		"message": "deleted",
	})
}
