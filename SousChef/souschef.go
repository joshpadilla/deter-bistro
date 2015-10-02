package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var db *sql.DB = nil
var kdb *sql.DB = nil

func bistroConnect() {
	var err error = nil
	db, err = sql.Open("mysql",
		"root:pizza@tcp(bistrofridge.deterlab.net:3306)/bistrofridge")
	if err != nil {
		fmt.Println("error opening connection to bistrofridge")
		fmt.Println(err)
		os.Exit(1)
	}
}

func kitchenConnect() {
	var err error = nil
	kdb, err = sql.Open("postgres",
		"host=kitchen.deterlab.net user=postgres dbname=fridge sslmode=disable password=pizza")
	if err != nil {
		fmt.Errorf("could not connect to the kitchen db")
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkQuantity(table, dbname string, db *sql.DB) int {

	q := fmt.Sprintf("SELECT count FROM %s", table)
	rows, err := db.Query(q)
	if err != nil {
		fmt.Println("unable to query ", dbname)
		fmt.Println(err)
	}

	if !rows.Next() {
		fmt.Println("the ", dbname, " is broken :(")
		rows.Close()
		return -1
	}

	var count int
	err = rows.Scan(&count)
	if err != nil {
		fmt.Println("the ", dbname, " is empty!")
		fmt.Println(err)
	}

	rows.Close()

	return count

}

func restock(table string, count int) {
	fmt.Printf("bistro %s supplies low, restocking\n", table)
	cmd := fmt.Sprintf("UPDATE %s SET count = %d", table, count+200)
	_, err := db.Query(cmd)
	if err != nil {
		fmt.Println("cant put stuff in the bistrofridge")
		fmt.Println(err)
	}
}

func grabIngredientPacks(count int) error {
	err := kdb.Ping()
	if err != nil {
		fmt.Println("unable to ping kitchen")
		fmt.Println(err)

		fmt.Println("trying to reconnect")
		kitchenConnect()
	}

	fmt.Println("grabbing fresh ingredients for the bistro from the kitchen fridge")
	count = checkQuantity("ingredient_packs", "kitchen fridge", kdb)
	cmd := fmt.Sprintf("UPDATE ingredient_packs SET count = %d", count-200)
	_, err = kdb.Query(cmd)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("the kitchen fridge is broke :/")
	}

	return nil
}

func main() {

	fmt.Println("I am the souschef har de har har")

	bistroConnect()
	kitchenConnect()

	for {

		err := db.Ping()
		if err != nil {
			fmt.Println("unable to ping bistrofridge")
			fmt.Println(err)

			fmt.Println("trying to reconnect")
			db, err = sql.Open("mysql",
				"root:pizza@tcp(bistrofridge.deterlab.net:3306)/bistrofridge")
			if err != nil {
				fmt.Println("error opening connection to bistrofridge")
				fmt.Println(err)
				os.Exit(1)
			}
		}

		if q := checkQuantity("ingredient_packs", "bistrofridge", db); q < 100 {
			grabIngredientPacks(200)
			restock("ingredient_packs", q)
		}

		if q := checkQuantity("doughballs", "bistrofridge", db); q < 100 {
			resp, err := http.Get("http://mixer.deterlab.net:8085/mix")
			if err != nil {
				fmt.Println("unable to make dough :(")
				fmt.Println(err)
			} else {
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("the dough has bugs")
					fmt.Println(err)
				} else if string(body) != "dough" {
					fmt.Println("the dough is not dough")
					fmt.Println(err)
				} else {
					restock("doughballs", q)
				}
			}
		}

		time.Sleep(2 * time.Second)

	}

}
