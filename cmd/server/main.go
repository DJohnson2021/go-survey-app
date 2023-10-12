package main

import( 
	"github.com/DJohnson2021/go-survey-app/db"
	"github.com/DJohnson2021/go-survey-app/utils"
)


func main() {
	utils.LoadEnv()
	db.InitDatabase()
	db.CloseDatabase()
}
