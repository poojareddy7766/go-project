// package main
// import 
// (
//     "fmt"
// 	"log"
// 	"os"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/joho/godotenv"
// )

// type Todo struct {
// 	ID        int    `json:"id"`
// 	Completed bool   `json:"completed"`
// 	Body      string `json:"body"`
// }

// var todos []Todo 
// func main(){
	
// 	// var myName string= "Pooja reddy"
// 	// var secondName string="Pooja reddy 2"
// 	// thirdName := "Pooja reddy 3"
// 	// fmt.Println(myName)
// 	// fmt.Println(secondName)
// 	// fmt.Println(thirdName)
//     fmt.Println("Hello World")
// 	app :=fiber.New()

// 	// trying to make port dynamic via .env file
// 	err := godotenv.Load(".env")
// 	if err != nil{
// 		log.Fatal("Error loading .env file")
// 	}

// 	PORT := os.Getenv("PORT")
// 	//fatal is equivalent to print() followed by an equivalent os.Exit(1)
//    	///----- All about get request
// 	//    app.Get("/",func(c *fiber.Ctx) error{
// 	// 	return c.Status(200).JSON(fiber.Map{"msg": "Start Understanding get request"})
// 	// })

// 	// get request
// 	app.Get("/api/todos",func(c *fiber.Ctx) error{
// 		return c.Status(200).JSON(todos)
// 	})
// 	////--- All about post request - create a todo
//     app.Post("/api/todos",func(c *fiber.Ctx) error{
// 		todo := &Todo{}
// 		if err := c.BodyParser(todo); err != nil {
// 			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body must not be a string"})
// 		}
// 		if todo.Body == "" {
// 			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
// 		}
// 		todo.ID = len(todos) + 1
// 		todos = append(todos, *todo)
// 		// var x int=5
// 		// var p *int=&X
// 		// fmt.println(p)
// 		// fmt.println(*p)
// 		return c.Status(201).JSON(todo)
// 	})

// 	////--- All about updating Todo 
// 	app.Patch("/api/todos/:id",func(c *fiber.Ctx) error{
//         id :=c.Params("id")
		
// 		for i, todo:=range todos{
// 			if fmt.Sprint(todo.ID) == id{
// 				todos[i].Completed = true
// 				return c.Status(200).JSON(todos[i])
// 			}
// 		}
// 		return c.Status(404).JSON(fiber.Map{"error": "Todo not Found"})
// 	})
	
// 	 // -- Delete a Todo 
// 	 app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")
// 		for i, todo := range todos {
// 			if fmt.Sprint(todo.ID) == id {
// 				todos = append(todos[:i], todos[i+1:]...) // Fixed closing parenthesis here
// 				return c.Status(200).JSON(fiber.Map{"Success": "true"})
// 			}
// 		}
// 		return c.Status(404).JSON(fiber.Map{"error": "Todo not Found"})
// 	})
   
// 	log.Fatal(app.Listen(":"+PORT))

// }

package main
import 
(
	"context"
    "fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
//mongoDb stores data in the form of Binary JSON
type Todo struct {
	ID        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}
var collection *mongo.Collection
func main(){
	fmt.Println("Hello World");
		err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading .env file")
	}
	//load and take the uri
	MONGODB_URI :=os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client,err := mongo.Connect(context.Background(),clientOptions)
	if err != nil{
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	err = client.Ping(context.Background(),nil)

	if err != nil{
		log.Fatal(err)
	}
	
	fmt.Println("CONNECTED TO MONGODB ATLAS")
 

	collection=client.Database("golang_db").Collection("todos")
	app :=fiber.New()	
	app.Get("api/todos",getTodos)
	app.Post("api/todos",createTodo)
    app.Patch("api/todos/:id",updateTodo)
	app.Delete("api/todos/:id",deleteTodo)
	
	port := os.Getenv("PORT")
    if port == ""{
		port ="5000"
	}

	log.Fatal(app.Listen("0.0.0.0:"+port))
}

// Get request
func getTodos(c *fiber.Ctx) error{
	var todos []Todo
	cursor,err := collection.Find(context.Background(),bson.M{})
	if err != nil{
		return err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()){
		var todo Todo
		if err := cursor.Decode(&todo);err!=nil{
			return err
		}
		todos = append(todos,todo)
	}
	return c.JSON(todos)
}

// Post Request

func createTodo(c *fiber.Ctx) error{
	todo := new(Todo)
	if err := c.BodyParser(todo); err!=nil{
		return err
	}
	if todo.Body==""{
		return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
	}
	insertResult,err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}
	todo.ID=insertResult.InsertedID.(primitive.ObjectID)
	return c.Status(201).JSON(todo)
 }

 // Patch Request
 func updateTodo(c *fiber.Ctx) error{
	id:=c.Params("id")
	objectID,err:=primitive.ObjectIDFromHex(id)
	if err!=nil{
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}
	filter :=bson.M{"_id":objectID}
	update :=bson.M{"$set":bson.M{"completed":true}}
	_,err=collection.UpdateOne(context.Background(),filter,update)
	if err!=nil{
		return err
	}
	return c.Status(200).JSON(fiber.Map{"success": true})
}






