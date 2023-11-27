package db

import (
	"fmt"

	"database/sql"

	easy "github.com/MhunterDev/Web/encryption"
	_ "github.com/lib/pq"
)

var connString, _ = easy.GetConn()

func BuildTables() error {

	userTable := "CREATE TABLE app.users(ID serial, username varchar(32) not null, token text not null, status varchar(20))"
	hashTable := "CREATE TABLE app.secret(ID serial,token text not null,hash text not null)"

	db, err := sql.Open("postgres", connString)
	if err != nil {
		fmt.Println("Error opening database")
		return err
	}
	defer db.Close()

	db.Exec(userTable)
	db.Exec(hashTable)

	token, hash, err := easy.HashAndToken("admin")
	if err != nil {
		fmt.Println("Error hashing password")
		return err
	}

	insertTokenBase := "INSERT INTO app.users(username,token,status) VALUES('admin',%s,'yes')"
	formatToken := fmt.Sprintf("'%s'", token)
	fullInsert := fmt.Sprintf(insertTokenBase, formatToken)

	db.Exec(fullInsert)

	hashInsertBase := "INSERT INTO app.secret(token,hash) VALUES(%s)"
	formatHashInsert := fmt.Sprintf("'%s','%s'", token, hash)
	fullHashInsert := fmt.Sprintf(hashInsertBase, formatHashInsert)

	db.Exec(fullHashInsert)
	return nil
}
