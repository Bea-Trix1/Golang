package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Produto struct {
	Id    string
	Nome  string
	Preco float64
}

func NewProduto(id string, nome string, preco float64) *Produto {
	return &Produto{
		Id:    uuid.New().String(),
		Nome:  nome,
		Preco: preco,
	}
}

func main() {
	db, err := sql.Open("mysql", "credenciais do banco de dados")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	produto := NewProduto("", "Produto 1", 12.34)
	err = InsertProduct(db, *produto)
	if err != nil {
		panic(err)
	}
}

func InsertProduct(db *sql.DB, p Produto) error {
	stmt, err := db.Prepare("insert into produtos(id, nome, preco) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Id, p.Nome, p.Preco)
	if err != nil {
		return err
	}

	return nil
}
