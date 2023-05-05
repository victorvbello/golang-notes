package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	bolt "go.etcd.io/bbolt"
)

var mysqlDB *sql.DB
var redisCli *redis.Client
var boltDB *bolt.DB

type dbTest struct {
	name string
	fn   func(*testing.B)
}

func TestMain(m *testing.M) {
	errs := make(chan error)
	go func() {
		for e := range errs {
			if e != nil {
				log.Fatal(fmt.Sprintf("error %v", e))
			}
		}
	}()
	errs <- mysqlInit()
	errs <- redisInit()
	errs <- boltDBInit()
	close(errs)
	m.Run()
	mysqlDB.Close()
	os.Remove(boltDB.Path())
	boltDB.Close()
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

	_, err = mysqlDB.Exec("DROP TABLE IF EXISTS table_benchmark_test_db")

	if err != nil {
		log.Fatalf("mysql error fail on drop table table_benchmark_test_db %v \n", err)
		return err
	}

	_, err = mysqlDB.Exec(`
		CREATE TABLE IF NOT EXISTS table_benchmark_test_db (
			id INT NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NULL,
			PRIMARY KEY (id)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)

	if err != nil {
		log.Fatalf("mysql error fail on prepare create table table_benchmark_test_db %v \n", err)
		return err
	}
	log.Println("mysql init success")
	return nil
}

func redisInit() error {
	var conf = struct {
		Address  string
		Port     int
		Password string
		DB       int
	}{
		Address:  "127.0.0.1",
		Port:     6379,
		Password: "",
		DB:       0,
	}
	address := fmt.Sprintf("%s:%s", conf.Address, strconv.Itoa(conf.Port))
	redisCli = redis.NewClient(&redis.Options{
		Addr:         address,
		Password:     conf.Password,
		DB:           conf.DB,
		MinIdleConns: 10,
		IdleTimeout:  240 * time.Second,
	})
	_, err := redisCli.Ping().Result()
	if err != nil {
		log.Fatalf("redis error fail ping %v \n", err)
		return err
	}
	err = redisCli.FlushDB().Err()
	if err != nil {
		log.Fatalf("redis error fail on FlushDB %v \n", err)
		return err
	}
	log.Println("redis init success")
	return nil
}

func boltDBInit() error {
	var err error
	boltDB, err = bolt.Open("./test.db", 0600, nil)
	if err != nil {
		log.Fatalf("boltDB error fail open %v \n", err)
		return err
	}
	err = boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("test-bucket"))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatalf("boltDB error fail on create bucket %v \n", err)
		return err
	}
	log.Println("boltDB init success")
	return nil
}

func BenchmarkInsert(bg *testing.B) {
	ts := []dbTest{
		{name: "mysql", fn: func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := mysqlDB.Exec(fmt.Sprintf("INSERT INTO table_benchmark_test_db (name) VALUES('%d')", rand.Int()))
				if err != nil {
					b.Error(err)
				}
			}
		}},
		{name: "redis", fn: func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err := redisCli.Set(fmt.Sprintf("test-%d", i), rand.Int(), 0).Err()
				if err != nil {
					b.Error(err)
				}
			}
		}},
		{name: "boltDB", fn: func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				boltDB.Update(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte("test-bucket"))
					err := b.Put([]byte(fmt.Sprintf("test-%d", i)), []byte(strconv.Itoa(rand.Int())))
					return err
				})
			}
		}},
	}
	for _, t := range ts {
		bg.Run(t.name, t.fn)
	}
}

func BenchmarkGet(bg *testing.B) {
	ts := []dbTest{
		{name: "mysql", fn: func(b *testing.B) {
			var name string
			for i := 0; i < b.N; i++ {
				row := mysqlDB.QueryRow("SELECT name FROM table_benchmark_test_db WHERE id =1")
				err := row.Scan(&name)
				if err != nil && err != sql.ErrNoRows {
					b.Error(err)
				}
			}
		}},
		{name: "redis", fn: func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err := redisCli.Get(fmt.Sprintf("test-%d", i)).Err()
				if err != nil && err.Error() != "redis: nil" {
					b.Error(err)
				}
			}
		}},
		{name: "boltDB", fn: func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				boltDB.View(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte("test-bucket"))
					b.Get([]byte(fmt.Sprintf("test-%d", i)))
					return nil
				})
			}
		}},
	}
	for _, t := range ts {
		bg.Run(t.name, t.fn)
	}
}

func BenchmarkUpdate(bg *testing.B) {
	ts := []dbTest{
		{name: "mysql", fn: func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := mysqlDB.Exec(fmt.Sprintf(`UPDATE table_benchmark_test_db SET name="%d"`, rand.Int()))
				if err != nil {
					b.Error(err)
				}
			}
		}},
		{name: "redis", fn: func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err := redisCli.Set(fmt.Sprintf("test-%d", i), rand.Int(), 0).Err()
				if err != nil {
					b.Error(err)
				}
			}
		}},
		{name: "boltDB", fn: func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				boltDB.Update(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte("test-bucket"))
					err := b.Put([]byte(fmt.Sprintf("test-%d", i)), []byte(strconv.Itoa(rand.Int())))
					return err
				})
			}
		}},
	}
	for _, t := range ts {
		bg.Run(t.name, t.fn)
	}
}

func BenchmarkDelete(bg *testing.B) {
	ts := []dbTest{
		{name: "mysql", fn: func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := mysqlDB.Exec(fmt.Sprintf(`DELETE FROM table_benchmark_test_db WHERE id=%d`, 2))
				if err != nil {
					b.Error(err)
				}
			}
		}},
		{name: "redis", fn: func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err := redisCli.Del(fmt.Sprintf("test-%d", i)).Err()
				if err != nil {
					b.Error(err)
				}
			}
		}},
		{name: "boltDB", fn: func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				boltDB.Update(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte("test-bucket"))
					err := b.Delete([]byte(fmt.Sprintf("test-%d", i)))
					return err
				})
			}
		}},
	}
	for _, t := range ts {
		bg.Run(t.name, t.fn)
	}
}
