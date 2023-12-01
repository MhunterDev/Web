package db

import (
	"database/sql"
	"errors"
	"fmt"

	easy "github.com/MhunterDev/Web/encryption"
	_ "github.com/lib/pq"
)

var connString, _ = easy.GetConn()

func isUser(username string) error {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return err
	}
	defer db.Close()

	var queryBase = "SELECT username FROM app.users WHERE username LIKE %s"
	formatUsername := fmt.Sprintf("'%s'", username)
	fullQuery := fmt.Sprintf(queryBase, formatUsername)

	var u string

	find := db.QueryRow(fullQuery)

	find.Scan(&u)

	if u == username {
		return nil
	}
	return errors.New("user not Found")
}

func AuthPass(username, password string) error {

	err := isUser(username)
	if err != nil {
		return err
	}

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return err
	}
	defer db.Close()

	var queryBase = "SELECT token FROM app.users WHERE username = %s"
	formatUsername := fmt.Sprintf("'%s'", username)
	fullQuery := fmt.Sprintf(queryBase, formatUsername)

	var t string

	find, err := db.Query(fullQuery)
	if err != nil {
		return err
	}

	for find.Next() {
		find.Scan(&t)
	}

	findHash := "SELECT hash FROM app.secret WHERE token = %s"
	formatHash := fmt.Sprintf("'%s'", t)
	fullHash := fmt.Sprintf(findHash, formatHash)

	var h string

	hash, err := db.Query(fullHash)
	if err != nil {
		return err
	}
	for hash.Next() {
		hash.Scan(&h)
	}

	return easy.AuthHash(h, password)

}
