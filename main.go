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

	log.Fatal(app.Listen(":4000"))
	//fatal is equivalent to print() followed by an equivalent os.Exit(1)

}