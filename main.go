package main

func main() {
	app := App{}
	app.Initialise(DbUser, DbPassword, DbName)
	app.Run("localhost:10000")
}
