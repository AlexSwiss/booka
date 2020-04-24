package models

type Author struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Bookid    int    `json:"bookid"`
}

type Book struct {
	ID     int       `json:"id" gorm:"primary_key"`
	Name   string    `json:"name"`
	Isbn   int       `json:"isbn"`
	Author []*Author `json:"author"`
}
