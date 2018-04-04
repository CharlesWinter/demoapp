package main

import (
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	/* fmt.Println("Hello world")

	db, err := sql.Open("mysql", "root:White123!@/vueapp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Check that we are connected to the db.
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println(pingErr)
	}

	//Insert a user into the database.
	stmt, prepErr := db.Prepare("insert into login(name, password) values (?,?)")
	if prepErr != nil {
		log.Fatal("Prep err", prepErr)
	}
	defer stmt.Close()

	_, insertErr := stmt.Exec(generatePassword())
	if insertErr != nil {
		log.Fatal(insertErr)
	}

	rows, err := db.Query("select * from login")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var login, password string
		var id int
		err := rows.Scan(&id, &login, &password)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, login, password)
	}
	if rowErrors := rows.Err(); rowErrors != nil {
		log.Fatal(err)
		rows.Close()
	} */
}

func generatePassword() (string, int) {
	rand.Seed(time.Now().UTC().UnixNano())
	return "BillyBob", rand.Intn(100)
}
