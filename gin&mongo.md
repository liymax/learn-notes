```go
package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)
type Crop struct {
	Name string `form:"name" json:"name" bson:"name"`
	GrowthCycle string `form:"growthCycle" json:"growthCycle" bson:"growthCycle"`
	SowingTime string `form:"sowingTime" json:"sowingTime" bson:"sowingTime"`
	GainTime string `form:"gainTime" json:"gainTime" bson:"gainTime"`
	Description string `form:"description" json:"description" bson:"description"`
}
type ResultCrop struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	GrowthCycle string `json:"growthCycle" bson:"growthCycle"`
	SowingTime string `json:"sowingTime" bson:"sowingTime"`
	GainTime string `json:"gainTime" bson:"gainTime"`
	Description string `json:"description" bson:"description"`
}
func main() {
	r := gin.Default()
	r.POST("/api/add/crop", func(c *gin.Context) {
		var data Crop
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, err := save(data)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("%#v\n", id)
		c.JSON(200, gin.H{
			"status": "success",
			"code": 0,
			"id": id,
		})
	})
	r.GET("/api/query/crop/:id", func(c *gin.Context) {
		id := c.Param("id")
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
			return
		}
		crop,err := getCrop(objId)
		c.JSON(200,gin.H{
			"status": "success",
			"code": 0,
			"result": crop,
		})
	})
	_ = r.Run(":9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func getCrop(id primitive.ObjectID) (ResultCrop,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://test2:test2@127.0.0.1:27017/?authSource=admin"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("test").Collection("crops")
	var result ResultCrop
	err = coll.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	fmt.Printf("%#v\n", result)
	return result, err
}

func save(data Crop) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://test2:test2@127.0.0.1:27017/?authSource=admin"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("test").Collection("crops")
	res, err := coll.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

```
