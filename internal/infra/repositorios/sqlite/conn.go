package sqlite

import (
	"database/sql"
	"embed"
	"fmt"
	"os"
	"regexp"
	"strings"

	_ "github.com/tursodatabase/go-libsql"

	t "github.com/duartqx/livredger/internal/common/types"
)

//go:embed modelagem/*
var migracoes embed.FS

func executarMigracoes(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	modelagem, err := migracoes.ReadDir("modelagem")
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`(?m)^-- sql: .*$`)

	for _, arquivo := range modelagem {

		if arquivo.IsDir() {
			continue
		}

		sql, err := migracoes.ReadFile(fmt.Sprintf("modelagem/%s", arquivo.Name()))

		if err != nil {
			return err
		}

		for _, stmt := range re.Split(string(sql), -1) {
			stmt = strings.Trim(stmt, "\n")

			if stmt == "" {
				continue
			}

			_, err = tx.Exec(stmt + ";")

			if err != nil {
				return err
			}
		}
	}
	return tx.Commit()
}

func dbFilename(usuario *t.Usuario) string {
	if usuario != nil && usuario.Id != 0 {
		return fmt.Sprintf("./%d.db", usuario.Id)
	}

	return "./local.db"
}

func Connect(usuario *t.Usuario) (db *sql.DB) {

	dbFilenameStr := dbFilename(usuario)

	_, errDbFileExists := os.Stat(dbFilenameStr)

	db, err := sql.Open("libsql", fmt.Sprintf("file:%s", dbFilenameStr))

	if err != nil {
		panic(err.Error())
	}

	if os.IsNotExist(errDbFileExists) {
		if err := executarMigracoes(db); err != nil {
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
