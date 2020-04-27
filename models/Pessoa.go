package models

// https://www.calhoun.io/inserting-records-into-a-postgresql-database-with-gos-database-sql-package/

import (

	"fmt"
	"database/sql"
	db "github.com/axeldeveloper/go.importar.dados.postgresql/db"
	"time"

)

type Pessoa struct {
	ID 		           sql.NullInt64
	Nome  		       sql.NullString
	Apelido 	       sql.NullString
	Cpf 		       sql.NullString
	Pai			       sql.NullString
	Mae                sql.NullString
	Nascimento         sql.NullString
	Sexo               sql.NullString
	Status             bool 
	Date_joined        time.Time
	Estabelecimento_id sql.NullInt64
	Religiao_id        sql.NullInt64 
	Tratamento_id      sql.NullInt64
	Usuario_id         sql.NullInt64 
	Vinculo_id         sql.NullInt64 
}

func BuscaTodos() ([]Pessoa, error) {
	
	db := db.Conecta()
	rows, err := db.Query(`
		select id, nome, apelido, cpf, pai, mae, nascimento, 
		   sexo, status, date_joined, estabelecimento_id, religiao_id, 
		   tratamento_id, usuario_id, vinculo_id 
		from pessoa_pessoafisica order by id asc ;`)

	if err != nil {
		panic(err.Error())
		return nil, err
	}

	defer rows.Close()

	pessoas := []Pessoa{}

	
	for rows.Next() {
		p := Pessoa{}
		err := rows.Scan(
			&p.ID,
			&p.Nome,
			&p.Apelido,
			&p.Cpf,
			&p.Pai,
			&p.Mae,
			&p.Nascimento,
			&p.Sexo,    
			&p.Status, 
			&p.Date_joined,
			&p.Estabelecimento_id,
			&p.Religiao_id, 
			&p.Tratamento_id,
			&p.Usuario_id,
			&p.Vinculo_id)
		
			pessoas = append(pessoas, p)

		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return nil, err
		}
	}
	//defer db.Close()
	
	return pessoas, nil
}