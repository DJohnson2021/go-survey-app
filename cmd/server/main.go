package main

import( 
	"github.com/DJohnson2021/go-survey-app/db"
)


func main() {
	db.InitDatabase()
	db.CloseDatabase()
}
