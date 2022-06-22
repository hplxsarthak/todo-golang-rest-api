package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

type todo struct{
	ID			string		`json:"id"`
	Item		string		`json:"title"`
	Completed	bool		`json:"completed"`
}

// Defining a todos array
var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Books", Completed: false},
	{ID: "3", Item: "Record Videos", Completed: false},
}

// Function to send data to the get request as a json format
func getTodos(context *gin.Context){
	// convert go array into json format
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	// Create a server
	router := gin.Default()

	// Get the data from the server
	router.GET("/todos", getTodos)

	// Run the server on port:9090
	router.Run("localhost:9090")
}