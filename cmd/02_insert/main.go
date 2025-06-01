package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dokoda1/go-counter/interface/db"
	"github.com/dokoda1/go-counter/interface/db/model"
	"github.com/dokoda1/go-counter/interface/db/query"
	"github.com/joho/godotenv"

	"github.com/dokoda1/go-counter"
)

func main() {
	// ...
	counter.Count(1, 10)

	loadEnv()

	sqlHandler, err := db.NewSqlHandler(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	ctx := context.Background()
	qu := query.Use(sqlHandler.Conn)
	domainQuery := qu.Domain

	// Insert
	domain := model.Domain{
		Name: "Name_sample",
	}
	err = qu.Transaction(func(tx *query.Query) error {
		if err := domainQuery.WithContext(ctx).Create(&domain); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

}

func loadEnv() {
	// ここで.envファイル全体を読み込みます。
	// この読み込み処理がないと、個々の環境変数が取得出来ません。
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load("../../../go-counter.env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
}
