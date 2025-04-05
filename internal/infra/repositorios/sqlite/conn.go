package sqlite

import (
	"database/sql"
	"embed"
	"fmt"
	"os"

	"github.com/duartqx/livro-razao/internal/common"
	_ "github.com/tursodatabase/go-libsql"
)

//go:embed modelagem/*
var migracoes embed.FS

func executarMigracoes(tx *sql.Tx) error {
	modelagem, err := migracoes.ReadDir("modelagem")
	if err != nil {
		return err
	}

	for _, migracao := range modelagem {

		if migracao.IsDir() {
			continue
		}

		sql, err := migracoes.ReadFile(fmt.Sprintf("modelagem/%s", migracao.Name()))

		if err != nil {
			return err
		}

		_, err = tx.Exec(string(sql))
		if err != nil {
			return err
		}
	}
	return nil
}

func dbFilename(usuario *common.Usuario) string {
	if usuario != nil && usuario.Id != 0 {
		return fmt.Sprintf("./%d.db", usuario.Id)
	}

	return "./local.db"
}

func Connect(usuario *common.Usuario) (db *sql.DB) {

	dbFilenameStr := dbFilename(usuario)

	_, errDbFileExists := os.Stat(dbFilenameStr)

	db, err := sql.Open("libsql", fmt.Sprintf("file:%s", dbFilenameStr))

	if err != nil {
		panic(err.Error())
	}

	if os.IsNotExist(errDbFileExists) {
		tx, err := db.Begin()
		if err != nil {
			panic(err.Error())
		}

		if err := executarMigracoes(tx); err != nil {
			panic(err.Error())
		}

		if err := tx.Commit(); err != nil {
			panic(err.Error())
		}
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec("PRAGMA strict = ON;")
	if err != nil {
		panic(err.Error())
	}

	return db
}
