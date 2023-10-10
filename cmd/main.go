package main

import (
	"demo/database"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	Id       int    `json:"id"`
	Todo     string `json:"todo"`
	Complete bool   `json:"complete"`
}

func main() {
	dns := "user:password@tcp(127.0.0.1:3306)/mydatabase"
	db, err := database.OpenDB(dns)
	if err != nil {
		log.Fatal(err)
	}

	// var sql = `CREATE TABLE IF NOT EXISTS todos (
	// 	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	// 	todo TEXT NOT NULL,
	// 	complete tinyint(1),
	// 	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	// 	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	// 	deleted_at DATETIME DEFAULT CURRENT_TIMESTAMP
	// );`

	// var input = `
	// 	INSERT INTO todos (todo, complete)
	// 	VALUES
	// 	("DB erstellen", 0),
	// 	("Schichtbuch erstellen", 0),
	// 	("Ausbessern Verbesserungen", 1),
	// 	("Rezepte erstellen", 0),
	// 	("Kaffewasser kochen", 1),
	// 	("Tasse reinigen", 0),
	// 	("Feierabend machen", 0)

	// `

	// res, err := db.Exec(input)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// id, _ := res.RowsAffected()
	// log.Println(id)

	var todos []Todo
	rows, err := db.Query("SELECT id, todo, complete FROM todos")

	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.Id, &todo.Todo, &todo.Complete)
		fmt.Println(todo)
		todos = append(todos, todo)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", len(todos))

}
