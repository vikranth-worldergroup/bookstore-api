package models

type Author struct{
    BookID string `json:"-" gorm:"primaryKey"`
    Name string `json:"name" gorm:"primaryKey" validate:"required"`

}

type Book struct{
	ID string `json:"id" gorm:"primaryKey"`
    Title string `json:"title" validate:"required,max=50"`
    Price float64 `json:"price" vaidate:"required,gt=9,lt=501"`
    Content string `json:"content" validate:"required"`  
    Authors []Author `json:"authors" gorm:"foreignKey:BookID"`
}