package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var mysqlDB *sql.DB
var bulkDataByBlockQuery string

func mysqlPrepareTableLoad() error {
	_, err := mysqlDB.Exec("DROP TABLE IF EXISTS table_benchmark_test_bulk_load")

	if err != nil {
		return fmt.Errorf("Exec[drop]: %v", err)
	}

	_, err = mysqlDB.Exec(`
		CREATE TABLE IF NOT EXISTS table_benchmark_test_bulk_load (
			id INT NOT NULL AUTO_INCREMENT,
			user_id INT NOT NULL DEFAULT 0,
			address_id INT NOT NULL DEFAULT 0,
			name VARCHAR(255) NULL,
			validated_user_id TINYINT(1) NULL DEFAULT 0,
			validated_address_id TINYINT(1) NULL DEFAULT 0,
			PRIMARY KEY (id)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	if err != nil {
		return fmt.Errorf("Exec[create]: %v", err)
	}

	bulkDataByBlockQuery = "INSERT INTO table_benchmark_test_bulk_load (user_id, address_id, name) VALUES "
	for i := 1; i <= 100000; i++ {
		bulkDataByBlockQuery += fmt.Sprintf("(%d, %d, '%d'), ", i, i, rand.Int())
	}

	bulkDataByBlockQuery = strings.TrimRight(bulkDataByBlockQuery, ", ") + ";"
	return nil
}

func mysqlPrepareTableUser() error {
	_, err := mysqlDB.Exec("DROP TABLE IF EXISTS table_benchmark_test_user")

	if err != nil {
		return fmt.Errorf("Exec[drop]: %v", err)
	}

	_, err = mysqlDB.Exec(`
		CREATE TABLE IF NOT EXISTS table_benchmark_test_user (
			id INT NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NULL,
			PRIMARY KEY (id)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	if err != nil {
		return fmt.Errorf("Exec[create]: %v", err)
	}

	queryUserInsert := "INSERT INTO table_benchmark_test_user (name) VALUES "

	for i := 0; i < 500; i++ {
		queryUserInsert += fmt.Sprintf("('%d'), ", rand.Int())
	}

	queryUserInsert = strings.TrimRight(queryUserInsert, ", ") + ";"

	_, err = mysqlDB.Exec(queryUserInsert)

	if err != nil {
		return fmt.Errorf("Exec[insert]: %v", err)
	}

	return nil
}

func mysqlPrepareTableAddress() error {
	_, err := mysqlDB.Exec("DROP TABLE IF EXISTS table_benchmark_test_address")

	if err != nil {
		return fmt.Errorf("Exec[drop]: %v", err)
	}

	_, err = mysqlDB.Exec(`
		CREATE TABLE IF NOT EXISTS table_benchmark_test_address (
			id INT NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NULL,
			PRIMARY KEY (id)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	if err != nil {
		return fmt.Errorf("Exec[create]: %v", err)
	}

	queryUserInsert := "INSERT INTO table_benchmark_test_address (name) VALUES "

	for i := 0; i < 1000; i++ {
		queryUserInsert += fmt.Sprintf("('%d'), ", rand.Int())
	}

	queryUserInsert = strings.TrimRight(queryUserInsert, ", ") + ";"

	_, err = mysqlDB.Exec(queryUserInsert)

	if err != nil {
		return fmt.Errorf("Exec[insert]: %v", err)
	}

	return nil
}

func mysqlPrepareTableOrders() error {
	_, err := mysqlDB.Exec(`
		CREATE TABLE IF NOT EXISTS table_benchmark_test_bulk_order (
			id INT NOT NULL AUTO_INCREMENT,
			user_id INT NOT NULL DEFAULT 0,
			address_id INT NOT NULL DEFAULT 0,
			order_id VARCHAR(255) NULL,
			PRIMARY KEY (id),
			KEY fk_table_benchmark_test_user_idx (user_id),
			KEY fk_table_benchmark_test_address_idx (address_id),
			CONSTRAINT fk_table_benchmark_test_user_id
				FOREIGN KEY (user_id)
				REFERENCES table_benchmark_test_user (id)
				ON DELETE RESTRICT
				ON UPDATE RESTRICT,
			CONSTRAINT fk_table_benchmark_test_address_id
				FOREIGN KEY (address_id)
				REFERENCES table_benchmark_test_address (id)
				ON DELETE RESTRICT
				ON UPDATE RESTRICT
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	if err != nil {
		return fmt.Errorf("Exec[create]: %v", err)
	}
	return nil
}

func mysqlInit() error {
	var err error
	var conf = struct {
		URL      string
		DBName   string
		User     string
		Port     int
		Password string
	}{
		URL:      "127.0.0.1",
		DBName:   "test",
		User:     "root",
		Port:     3306,
		Password: "local-pass",
	}
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.User, conf.Password, conf.URL, strconv.Itoa(conf.Port), conf.DBName)
	mysqlDB, err = sql.Open("mysql", conn)
	if err != nil {
		log.Fatal(fmt.Sprintf("mysql error %v", err))
		return err
	}

	err = mysqlDB.Ping()
	if err != nil {
		log.Fatal(fmt.Sprintf("mysql error %v", err))
		return err
	}

	_, err = mysqlDB.Exec("DROP TABLE IF EXISTS table_benchmark_test_bulk_order")

	if err != nil {
		return fmt.Errorf("drop table_benchmark_test_bulk_order: %v", err)
	}

	err = mysqlPrepareTableUser()
	if err != nil {
		log.Fatal(fmt.Sprintf("mysqlPrepareTableUser: %v", err))
		return err
	}

	err = mysqlPrepareTableAddress()
	if err != nil {
		log.Fatal(fmt.Sprintf("mysqlPrepareTableAddress: %v", err))
		return err
	}

	err = mysqlPrepareTableLoad()
	if err != nil {
		log.Fatal(fmt.Sprintf("mysqlPrepareTableLoad: %v", err))
		return err
	}

	err = mysqlPrepareTableOrders()
	if err != nil {
		log.Fatal(fmt.Sprintf("mysqlPrepareTableOrders: %v", err))
		return err
	}

	log.Println("mysql init success")
	return nil
}

func TestMain(m *testing.M) {
	err := mysqlInit()
	if err != nil {
		log.Fatal(fmt.Sprintf("error %v", err))
	}
	m.Run()
	mysqlDB.Close()
}

func BenchmarkInsertOneByOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := mysqlDB.Exec(fmt.Sprintf("INSERT INTO table_benchmark_test_bulk_load (name) VALUES('%d')", rand.Int()))
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkInsertBlock10Stmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stmt, err := mysqlDB.Prepare("INSERT INTO table_benchmark_test_bulk_load (name) VALUES(?);")
		if err != nil {
			b.Error(err)
		}
		for i := 0; i < 10; i++ {
			_, err := stmt.Exec(rand.Int())
			if err != nil {
				b.Error(err)
			}
		}
		stmt.Close()
	}
}
func BenchmarkInsertBlock20Stmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stmt, err := mysqlDB.Prepare("INSERT INTO table_benchmark_test_bulk_load (name) VALUES(?);")
		if err != nil {
			b.Error(err)
		}
		for i := 0; i < 20; i++ {
			_, err := stmt.Exec(rand.Int())
			if err != nil {
				b.Error(err)
			}
		}
		stmt.Close()
	}
}

func BenchmarkInsertBlock100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := mysqlDB.Exec(bulkDataByBlockQuery)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkInsertValidateBlock100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// check if user_id and address_id existe
		_, err := mysqlDB.Exec(`
			UPDATE table_benchmark_test_bulk_load
			INNER JOIN table_benchmark_test_user ON table_benchmark_test_user.id = table_benchmark_test_bulk_load.user_id
			SET validated_user_id = 1;`)
		if err != nil {
			b.Error("update user_id", err)
		}
		_, err = mysqlDB.Exec(`
			UPDATE table_benchmark_test_bulk_load
			INNER JOIN table_benchmark_test_address ON table_benchmark_test_address.id = table_benchmark_test_bulk_load.address_id
			SET validated_address_id = 1;`)
		if err != nil {
			b.Error("update address_id", err)
		}
		// insert only valid registers
		_, err = mysqlDB.Exec(`
			INSERT INTO table_benchmark_test_bulk_order (user_id, address_id, order_id)
			SELECT user_id, address_id, name FROM table_benchmark_test_bulk_load
			WHERE
				validated_user_id = 1
				AND validated_address_id = 1;`)
		if err != nil {
			b.Error(err)
		}

	}
}
