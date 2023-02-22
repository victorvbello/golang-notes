package main

import (
	"fmt"
	"gonotes/course/db/mariadb"
	"math/rand"
	"time"

	"log"

	"github.com/jmoiron/sqlx"
)

func getLastInserteIdUsingConcurrency(db *sqlx.DB) {
	totalTestFlows := 90

	log.Println("getLastInserteIdUsingConcurrency")

	_, err := db.Exec("DROP TABLE IF EXISTS table_concurrency_test;")
	if err != nil {
		log.Fatalf("fail on prepare drop table table_concurrency_test%v \n", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS table_concurrency_test (
			id INT NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NULL,
			PRIMARY KEY (id),
			UNIQUE KEY name (name) USING BTREE
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)
	if err != nil {
		log.Fatalf("fail on prepare create table table_concurrency_test%v \n", err)
	}

	lastIndexReturnMap := make(map[int]int)
	lastIndexChan := make(chan int)

	for i := 0; i < totalTestFlows; i++ {
		rand.Seed(time.Now().UnixNano())
		go func(index int) {
			flowName := fmt.Sprintf("FLOW %d", index)

			testQuery := fmt.Sprintf(`INSERT INTO table_concurrency_test (name) VALUE ("%d")`, rand.Int())

			log.Printf("%s Start %s\n", flowName, testQuery)

			result, err := db.Exec(testQuery)

			if err != nil {
				log.Fatalf("%s db.Exec %v \n", flowName, err)
			}

			rowsAffected, err := result.RowsAffected()
			if err != nil {
				log.Fatalf("%s result.RowsAffected() %v \n", flowName, err)
			}
			log.Printf("%s rowsAffected %d \n", flowName, rowsAffected)

			lastInsertId, err := result.LastInsertId()

			if err != nil {
				log.Fatalf("%s result.LastInsertId() %v \n", flowName, err)
			}
			log.Printf("%s LastInsertId %d \n", flowName, lastInsertId)

			lastIndexChan <- int(lastInsertId)
		}(i)
	}

	for i := 0; i < totalTestFlows; i++ {
		lId := <-lastIndexChan
		lastIndexReturnMap[lId]++
	}

	close(lastIndexChan)

	if len(lastIndexReturnMap) != totalTestFlows {
		log.Panicf("lastIndexReturnMap(%d) and totalTestFlows(%d) are different", len(lastIndexReturnMap), totalTestFlows)
	}

	// Validate
	log.Println("Init Validation")
	for k, v := range lastIndexReturnMap {
		if v > 1 {
			log.Printf("id %d was returned %v times \n", k, v)
		}
	}
}

func getLastInserteIdUsingConcurrencyAndDBStmt(db *sqlx.DB) {
	totalTestFlows := 90

	log.Println("getLastInserteIdUsingConcurrencyAndDBStmt")

	_, err := db.Exec("DROP TABLE IF EXISTS table_concurrency_test;")
	if err != nil {
		log.Fatalf("fail on prepare drop table table_concurrency_test%v \n", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS table_concurrency_test (
			id INT NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NULL,
			PRIMARY KEY (id),
			UNIQUE KEY name (name) USING BTREE
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)
	if err != nil {
		log.Fatalf("fail on prepare create table table_concurrency_test%v \n", err)
	}

	lastIndexReturnMap := make(map[int]int)
	lastIndexChan := make(chan int)

	for i := 0; i < totalTestFlows; i++ {
		rand.Seed(time.Now().UnixNano())
		go func(index int) {
			flowName := fmt.Sprintf("FLOW %d", index)

			testQuery := fmt.Sprintf(`INSERT INTO table_concurrency_test (name) VALUE ("%d")`, rand.Int())

			stmt, err := db.Prepare(testQuery)
			if err != nil {
				log.Fatalf("%s db.Prepare %v \n", flowName, err)
			}

			defer stmt.Close()

			log.Printf("%s Start %s\n", flowName, testQuery)

			result, err := stmt.Exec()

			if err != nil {
				log.Fatalf("%s stmt.Exec %v \n", flowName, err)
			}

			rowsAffected, err := result.RowsAffected()
			if err != nil {
				log.Fatalf("%s result.RowsAffected() %v \n", flowName, err)
			}
			log.Printf("%s rowsAffected %d \n", flowName, rowsAffected)

			lastInsertId, err := result.LastInsertId()

			if err != nil {
				log.Fatalf("%s result.LastInsertId() %v \n", flowName, err)
			}
			log.Printf("%s LastInsertId %d \n", flowName, lastInsertId)

			lastIndexChan <- int(lastInsertId)
		}(i)
	}

	for i := 0; i < totalTestFlows; i++ {
		lId := <-lastIndexChan
		lastIndexReturnMap[lId]++
	}

	close(lastIndexChan)

	if len(lastIndexReturnMap) != totalTestFlows {
		log.Panicf("lastIndexReturnMap(%d) and totalTestFlows(%d) are different", len(lastIndexReturnMap), totalTestFlows)
	}

	// Validate
	log.Println("Init Validation")
	for k, v := range lastIndexReturnMap {
		if v > 1 {
			log.Printf("id %d was returned %v times \n", k, v)
		}
	}
}

func main() {
	db, err := mariadb.SQLOpenConnection(mariadb.MariDBConfig{
		URL:      "127.0.0.1",
		DBName:   "test",
		User:     "root",
		Port:     3306,
		Password: "local-pass",
	})

	if err != nil {
		log.Fatalf("mariadb.SQLOpenConnection %v \n", err)
	}

	defer func() {
		log.Println("End")
		err := mariadb.SQLCloseConnection(db)
		if err != nil {
			log.Fatalf("mariadb.SQLCloseConnection %v \n", err)
		}
	}()

	getLastInserteIdUsingConcurrency(db)
	getLastInserteIdUsingConcurrencyAndDBStmt(db)
}
