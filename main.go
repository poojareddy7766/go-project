package main
import 
(
    "fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var todos []Todo 
func main(){
	
	// var myName string= "Pooja reddy"
	// var secondName string="Pooja reddy 2"
	// thirdName := "Pooja reddy 3"
	// fmt.Println(myName)
	// fmt.Println(secondName)
	// fmt.Println(thirdName)
    fmt.Println("Hello World")
	app :=fiber.New()

	// trying to make port dynamic via .env file
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	//fatal is equivalent to print() followed by an equivalent os.Exit(1)
   	///----- All about get request
	//    app.Get("/",func(c *fiber.Ctx) error{
	// 	return c.Status(200).JSON(fiber.Map{"msg": "Start Understanding get request"})
	// })

	// get request
	app.Get("/api/todos",func(c *fiber.Ctx) error{
		return c.Status(200).JSON(todos)
	})
	////--- All about post request - create a todo
    app.Post("/api/todos",func(c *fiber.Ctx) error{
		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body must not be a string"})
		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)
		// var x int=5
		// var p *int=&X
		// fmt.println(p)
		// fmt.println(*p)
		return c.Status(201).JSON(todo)
	})

	////--- All about updating Todo 
	app.Patch("/api/todos/:id",func(c *fiber.Ctx) error{
        id :=c.Params("id")
		
		for i, todo:=range todos{
			if fmt.Sprint(todo.ID) == id{
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not Found"})
	})
	
	 // -- Delete a Todo 
	 app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...) // Fixed closing parenthesis here
				return c.Status(200).JSON(fiber.Map{"Success": "true"})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not Found"})
	})
   
	log.Fatal(app.Listen(":"+PORT))

}