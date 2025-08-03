package models

type Book struct{
	ID string `json:"id" gorm:"primaryKey"`
    Title string `json:"title" validate:"required,max=50"`
    Author string `json:"author" validate:"required"`
    Price float64 `json:"price" vaidate:"required,gt=9,lt=501"`
    Content string `json:"content" validate:"required"`  
}