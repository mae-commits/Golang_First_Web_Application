package main

// gorm.io をimport
import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// CreditCard struct を定義
type CreditCard struct {
	gorm.Model // "id", "created_at", "deleted_at" カラムの値を追加（デフォルトのパッケージを追加）
	Number     string
	UserID     int
	UserID2    string `gorm:"default:100"`  // Default の値を定義
	UserID3    int64  `gorm:"default:1000"` // Default の値を定義
}

type User struct {
	gorm.Model
	Name       string `gorm:"default:Unknown"` // Default の値を定義
	CreditCard CreditCard
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	}
	// CreditCard table を"test.db"中に存在しないならば
	// table を新たに作成、既存ならばtableを開き、新たに入力
	db.AutoMigrate(&CreditCard{})
	// 同上
	db.AutoMigrate(&User{})
	// User table に CreditCard の Number, UserID3 field の値を入力
	db.Create(&User{Name: "Naoki",
		CreditCard: CreditCard{Number: "222222222", UserID3: 3300434},
	})
}
