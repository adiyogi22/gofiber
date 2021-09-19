package jobs

import (
	"fmt"
	"context"
	"github.com/gofiber/fiber"
	// "time"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/adiyogi22/gofiber/database"
)

// //Issue - struct to map with mongodb documents
// type Issue struct {
// 	ID          primitive.ObjectID `bson:"_id"`
// 	CreatedAt   time.Time          `bson:"created_at"`
// 	UpdatedAt   time.Time          `bson:"updated_at"`
// 	Title       string             `bson:"title"`
// 	Code        string             `bson:"code"`
// 	Description string             `bson:"description"`
// 	Completed   bool               `bson:"completed"`
// }

type Trainer struct {
	Name string
	Age  int
	City string
}

type Response struct {
	success bool
}

func GetJobs(c *fiber.Ctx) {
	c.Send("All Job")
}

func GetJob(c *fiber.Ctx) {
	c.Send("Single Job")
}

func NewJob(c *fiber.Ctx) {

	ash := Trainer{"John Doe", 10, "Kathmandu"}
    fmt.Println(ash)
	client, err := database.GetMongoClient()
	if err != nil {
		c.Send("Error Connecting 1")
	}
	
	collection := client.Database(database.DB).Collection(database.ISSUES)
	_, err = collection.InsertOne(context.TODO(), ash)
	if err != nil {
		c.Send("Error Connecting 2")
	}

	c.Send(Response{success:true})
}

func DeleteJob(c *fiber.Ctx) {
	c.Send("Delete Job")
}


