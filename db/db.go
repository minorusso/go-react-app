package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// *gorm.DB という部分が戻り値の型（アドレス）
// * の前にある型は、その型へのポインタを格納する変数を宣言することを意味します。ポインタ型は、データの実際の値ではなく、
// データへの参照を保持します。この場合、*gorm.DB は gorm.DB 型へのポインタを示しています。
func NewDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "dev" {
		// ローカルの環境変数を参照
		err := godotenv.Load()
		if err != nil {
			// エラー出力後強制終了
			log.Fatalln(err)
		}
	}
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		// エラー出力後強制終了
		log.Fatalln(err)
	}
	fmt.Println("Connceted")
	return db
}

// 2つ目の戻り値はエラーですが、提供されたコードでは _ としてアンダースコアが使われています。
// アンダースコアは通常、ゴミ値（discard value）として扱われ、変数の値を無視するために使用されます。
func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}

// migrateから
