package main
import 
(
    "fmt"
	"log"
	"github.com/gofiber/fiber/v2"
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
	//fatal is equivalent to print() followed by an equivalent os.Exit(1)
   	///----- All about get request
	   app.Get("/",func(c *fiber.Ctx) error{
		return c.Status(200).JSON(fiber.Map{"msg": "Start Understanding get request"})
	})

	////--- All about post request - create a todo
    app.Post("/api/todos",func(c *fiber.Ctx) error{
		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)
		return c.Status(201).JSON(todo)
	})

	log.Fatal(app.Listen(":4000"))

}