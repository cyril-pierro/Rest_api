package helpers

import (
	"strconv"

	"github.com/cyril_pierro/Rest_api_2/model"
	sql "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db_value *gorm.DB

// start database up
func StartDb() {
	db, err := gorm.Open(sql.Open("../test.db"), &gorm.Config{})
	CheckError(err)
	setDbValue(db)
	db.AutoMigrate(&model.Book{})
}

func setDbValue(db *gorm.DB) {
	db_value = db
}

// Operations used on database

func GetAllBooks(books []model.Book) []model.Book {
	data := db_value.Find(&books)
	CheckError(data.Error)
	return books

}
func GetBookWithId(id int) model.Book {
	var book model.Book
	db_value.First(&book, id)
	return book
}

func AddBook(book *model.Book) {
	db_value.Create(book)
}

func ParseInt(v string) int {
	value, _ := strconv.Atoi(v)
	return value
}

func UpdateBook(book *model.Book, new_book *model.Book) {
	db_value.Model(&book).Updates(new_book)
}

func DeleteBook(book *model.Book, id int) {
	db_value.Delete(&book, id)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
