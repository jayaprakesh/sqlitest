package policy

import "database/sql"

type pol struct {
	DB *sql.DB
}

func Policyval(db *sql.DB) *pol {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "policy" (
			"ID" INTEGER PRIMARY KEY AUTOINCREMENT,
			"content" TEXT
		);
	`)
	stmt.Exec()
	return &pol{
		DB: db,
	}
}

func (pol *pol) Get() []Item {
	items := []Item{}
	rows, _ := pol.DB.Query(`
		SELECT * FROM policy
	`)
	var id int
	var content string
	for rows.Next() {
		rows.Scan(&id, &content)
		items = append(items, Item{
			ID:      id,
			Content: content,
		})
	}
	defer rows.Close()
	return items
}

func (pol *pol) Update(item Item, id int) {
	stmt, _ := pol.DB.Prepare(`
		UPDATE policy SET content = ? WHERE ID = ?
	`)
	stmt.Exec(item.Content, id)
}

func (pol *pol) Delete(id int) {
	stmt, _ := pol.DB.Prepare(`
		DELETE FROM policy WHERE ID = ?
	`)
	stmt.Exec(id)
}

func (pol *pol) Add(item Item) {
	stmt, _ := pol.DB.Prepare(`
		INSERT INTO policy (content) values (?)
	`)
	stmt.Exec(item.Content)
}
