package main

func main() {
	db := Database()
	db.AutoMigrate(&Product{})

	var test1 = Product{ID: 111, Name: "Sony Playstation 4", Description: "Home video game console", Price: "500$"}
	var test2 = Product{ID: 222, Name: "Xbox One S", Description: "Home video game console", Price: "420$"}
	var test3 = Product{ID: 333, Name: "Nintendo Switch", Description: "Home video game console", Price: "370$"}

	db.Create(&test1)
	db.Create(&test2)
	db.Create(&test3)

	defer db.Close()
}
