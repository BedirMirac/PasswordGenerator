package cmd

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type PasswordData struct {
	Appname  string
	Password string
}

func Save(data PasswordData) {

	insertSQL := `INSERT INTO passwords (appname, password) VALUES (?, ?)`

	_, err := DB.Exec(insertSQL, data.Appname, data.Password)
	if err != nil {
		log.Fatal("Error inserting data:", err)
	}

}

func List() {
	rows, err := DB.Query("SELECT id, appname, password FROM passwords")
	if err != nil {
		log.Fatal("Error fetching data:", err)
	}
	defer rows.Close()

	fmt.Println("ID | SITE | PASSWORD")
	fmt.Println("--------------------")

	for rows.Next() {
		var (
			id       int
			appname  string
			password string
		)
		err := rows.Scan(&id, &appname, &password)
		if err != nil {
			log.Fatal("Error during fetching data from a row:", err)
		}

		fmt.Printf("%d | %s | %s\n", id, appname, password)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal("Error during scanning:", err)
	}
}

func Delete(id int) {
	query := "DELETE FROM passwords WHERE id = ?"

	res, err := DB.Exec(query, id)
	if err != nil {
		log.Fatal("Error during deleting the password: ", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal("Error during deleting the password: ", err)
	}
	if rowsAffected == 0 {
		fmt.Printf("There is no password with %d id. Nothing deleted!\n", id)
		return
	}
	fmt.Printf("Password you chose is deleted successfully ")
}

func Update(id int, newPass string) {
	var oldPass string
	err := DB.QueryRow("SELECT password FROM passwords WHERE id = ?", id).Scan(&oldPass)

	if err == sql.ErrNoRows {
		fmt.Printf("Error: %d id caanot be found\n", id)
		return
	} else if err != nil {
		log.Fatal("Error during updating the password: ", err)
	}

	if oldPass == newPass {
		fmt.Println("The password you entered is the same as the password you tried to change! Nothing updated!")
		return
	}

	query := "UPDATE passwords SET password = ? WHERE id = ?"

	_, err = DB.Exec(query, newPass, id)
	if err != nil {
		log.Fatal("Error during updating the password: ", err)
	}

	fmt.Println("Password you chose is updated successfully.")
}
