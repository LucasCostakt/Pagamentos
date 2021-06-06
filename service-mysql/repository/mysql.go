package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Storage {
	return &repo{
		db: db,
	}
}

func (tx *repo) CreateUserTable() error {

	query := `CREATE TABLE IF NOT EXISTS mydb.users (
	  id INT NOT NULL AUTO_INCREMENT,
	  name VARCHAR(255) NOT NULL,
	  cpf_cpnj VARCHAR(255) NOT NULL,
	  email VARCHAR(255) NOT NULL,
	  password VARCHAR(255) NOT NULL,
	  type INT NOT NULL,
	  balance FLOAT NOT NULL DEFAULT 0,
	  createdate DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	  updatedate DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	  PRIMARY KEY (id),
	  INDEX fk_users_userstype1_idx (type ASC),
	  CONSTRAINT fk_users_userstype1
		FOREIGN KEY (type)
		REFERENCES mydb.userstype (id)
		ON DELETE NO ACTION
		ON UPDATE NO ACTION)
	ENGINE = InnoDB;
	`

	err := tx.executeInsert(query)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (tx *repo) CreateUserTypeTable() error {
	query := `CREATE TABLE IF NOT EXISTS mydb.userstype (
		id INT NOT NULL,
		label VARCHAR(255) NOT NULL,
		createdate DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updatedate DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id))
	  ENGINE = InnoDB;`
	err := tx.executeInsert(query)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (tx *repo) CreateSchema() error {
	query := `CREATE SCHEMA IF NOT EXISTS mydb DEFAULT CHARACTER SET utf8 ;`
	err := tx.executeInsert(query)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (tx *repo) CreateTransationTable() error {
	query := `CREATE TABLE IF NOT EXISTS mydb.transation (
		id INT NOT NULL AUTO_INCREMENT,
		value FLOAT NOT NULL,
		payer INT NOT NULL,
		payee INT NOT NULL,
		end_reversal_date DATETIME NOT NULL,
		createdate DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updatedate DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id),
		INDEX fk_transation_users1_idx (payer ASC) ,
		INDEX fk_transation_users2_idx (payee ASC) ,
		CONSTRAINT fk_transation_users1
		  FOREIGN KEY (payer)
		  REFERENCES mydb.users (id)
		  ON DELETE NO ACTION
		  ON UPDATE NO ACTION,
		CONSTRAINT fk_transation_users2
		  FOREIGN KEY (payee)
		  REFERENCES mydb.users (id)
		  ON DELETE NO ACTION
		  ON UPDATE NO ACTION)
	  ENGINE = InnoDB;`
	err := tx.executeInsert(query)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (tx *repo) CreateReversalTable() error {
	query := `CREATE TABLE IF NOT EXISTS mydb.reversal (
		id INT NOT NULL AUTO_INCREMENT,
		transation_id INT NOT NULL,
		payer INT NOT NULL,
		payee INT NOT NULL,
		createdate DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updatedate DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id),
		INDEX fk_reversal_transation1_idx (transation_id ASC) ,
		INDEX fk_reversal_users1_idx (payer ASC) ,
		INDEX fk_reversal_users2_idx (payee ASC) ,
		CONSTRAINT fk_reversal_transation1
		  FOREIGN KEY (transation_id)
		  REFERENCES mydb.transation (id)
		  ON DELETE NO ACTION
		  ON UPDATE NO ACTION,
		CONSTRAINT fk_reversal_users1
		  FOREIGN KEY (payer)
		  REFERENCES mydb.users (id)
		  ON DELETE NO ACTION
		  ON UPDATE NO ACTION,
		CONSTRAINT fk_reversal_users2
		  FOREIGN KEY (payee)
		  REFERENCES mydb.users (id)
		  ON DELETE NO ACTION
		  ON UPDATE NO ACTION)
	  ENGINE = InnoDB;`
	err := tx.executeInsert(query)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (tx *repo) CreateUsersTypes() error {
	query := []string{`INSERT INTO mydb.userstype (id, label)
	SELECT * FROM (SELECT 1, 'comum') AS tmp
	WHERE NOT EXISTS (SELECT id, label FROM mydb.userstype WHERE id = 1 or label = 'comum') LIMIT 1;`,

		`INSERT INTO mydb.userstype (id, label)
	SELECT * FROM (SELECT 2, 'lojista') AS tmp
	WHERE NOT EXISTS (SELECT id, label FROM mydb.userstype WHERE id = 2 or label = 'lojista') LIMIT 1;`}

	for _, q := range query {
		err := tx.executeInsert(q)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

//Create users test
func (tx *repo) CreateUsers() error {
	query := []string{`INSERT INTO mydb.users (id, name, cpf_cpnj, email, password, type, balance)
	SELECT * FROM (SELECT 3, 'Tim Maia', '09228574070', 'timmaia@gmail.com', '12345', 1, 5000) AS tmp
	WHERE NOT EXISTS (SELECT cpf_cpnj, email FROM mydb.users  WHERE cpf_cpnj = '09228574070' or email = 'timmaia@gmail.com') LIMIT 1;`,

		`INSERT INTO mydb.users  (id, name, cpf_cpnj, email, password, type, balance)
	SELECT * FROM (SELECT 2, 'Darth Vader', '05593744025', 'darthvader@gmail.com', '12345', 1, 0) AS tmp
	WHERE NOT EXISTS (SELECT cpf_cpnj, email FROM mydb.users  WHERE cpf_cpnj = '05593744025' or email = 'darthvader@gmail.com') LIMIT 1;`,

		`INSERT INTO mydb.users (id, name, cpf_cpnj, email, password, type, balance)
	SELECT * FROM (SELECT 1, 'R2 D2', '61915727000193', 'r2d2@gmail.com', '12345', 2, 1000) AS tmp
	WHERE NOT EXISTS (SELECT cpf_cpnj, email FROM mydb.users  WHERE cpf_cpnj = '61915727000193' or email = 'r2d2@gmail.com') LIMIT 1;`,

		`INSERT INTO mydb.users (id, name, cpf_cpnj, email, password, type, balance)
	SELECT * FROM (SELECT 4, 'NGolo Kanté', '92531761000198', 'kante@gmail.com', '12345', 2, 500) AS tmp
	WHERE NOT EXISTS (SELECT cpf_cpnj, email FROM mydb.users  WHERE cpf_cpnj = '92531761000198' or email = 'kante@gmail.com') LIMIT 1;`}

	for _, q := range query {
		err := tx.executeInsert(q)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (tx *repo) executeInsert(query string) error {
	ls, err := tx.db.Exec(query)
	if err != nil {
		log.Println(err)
		return err
	}

	i, err := ls.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(i)
	return nil
}
