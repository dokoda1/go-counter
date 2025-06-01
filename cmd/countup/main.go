package main

import (
	"fmt"
	"os"

	"gorm.io/gen"

	"mymodule/interface/db"

	"github.com/dokoda1/go-counter"

	"github.com/joho/godotenv"
)

func main() {
	// ...
	counter.Count(1, 10)

	//func loadEnvを呼び出します。
	loadEnv()

	g := gen.NewGenerator(gen.Config{
		OutPath:           "../../interface/db/query",
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		FieldWithIndexTag: true,                                                               // modelのgormタグのindex節を生成する
		FieldWithTypeTag:  true,                                                               // modelのgormタグのtype節を生成する
		FieldNullable:     true,                                                               // null許容カラムのときポインタ型にする
	})

	//
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
	g.UseDB(sqlHandler.Conn) // reuse your gorm db
}

// .envを呼び出します。
func loadEnv() {
	// ここで.envファイル全体を読み込みます。
	// この読み込み処理がないと、個々の環境変数が取得出来ません。
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load("../../../go-counter.env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	// .envの SAMPLE_MESSAGEを取得して、messageに代入します。
	message := os.Getenv("SAMPLE_MESSAGE")
	fmt.Println(message)
}
