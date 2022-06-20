// package main

// import (
// 	"database/sql"

// 	_ "github.com/mattn/go-sqlite3"
// )

// func main() {
// 	db, _ := sql.Open("sqlite3", "./crud.db")

// 	stmt, _ := db.Prepare(`
// 	CREATE TABLE "crud" (
// 		"ID"	INTEGER NOT NULL,
// 		"CONTENT"	TEXT,
// 		PRIMARY KEY("ID" AUTOINCREMENT)
// 	);
// 	`)
// 	stmt.Exec()
// }
package main

import (
	"crud/platform/policy"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

/*
CREATE TABLE IF NOT EXISTS "crud" (
	"ID"	INTEGER NOT NULL,
	"content"	TEXT,
	PRIMARY KEY("ID" AUTOINCREMENT)
);
*/

func main() {
	db, _ := sql.Open("sqlite3", "./crud.db")

	pol := policy.Policyval(db)
	var choice int
	fmt.Println("BASIC CRUD OPERATIONS:")
	fmt.Println("1. Inserting into table")
	fmt.Println("2. Updating table")
	fmt.Println("3. Delete from table")
	fmt.Println("4. Diplay table")
	fmt.Println("5. Exit")
	fmt.Println("Enter your choice:")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		var st string
		fmt.Println("Enter content to be inserted:")
		fmt.Scanln(&st)
		pol.Add(policy.Item{
			Content: st,
		})

	case 2:
		var id int
		var cont string
		fmt.Println("Enter content to be Updated in the given specific ID:")
		fmt.Scanln(&id)
		fmt.Scanln(&cont)
		pol.Update(policy.Item{
			Content: cont,
		}, id)

	case 3:
		var id int
		fmt.Println("Enter the row ID to be Deleted:")
		fmt.Scanln(&id)
		pol.Delete(id)

	case 4:
		items := pol.Get()
		fmt.Println(items)

	case 5:
		break
	default:
		fmt.Println("Invalid Option")
	}
	/*pol.Update(policy.Item{
		Content: "Policy 3",
	}, 1)*/
	defer db.Close()
}
