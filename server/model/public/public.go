package public

import "github.com/flipped-aurora/gin-vue-admin/server/model/uploadImage"

type ApartmentCountCity struct {
	City            string              `json:"city" gorm:"column:city"`
	CityChineseName string              `json:"cityChineseName" gorm:"column:city_chinese_name"`
	Count           int                 `json:"count" gorm:"column:apartment_count"`
	Image           []uploadImage.Image `json:"image" gorm:"type:json;serializer:json;column:image"`
}

type ApartmentsByCity struct {
	ApartmentName string              `json:"apartmentName" gorm:"column:apartment_name"`
	Price         float64             `json:"price" gorm:"column:price"`
	Company       string              `json:"company" gorm:"column:company"`
	Introduction  string              `json:"introduction" gorm:"column:introduction"`
	Type          string              `json:"type" gorm:"column:type"`
	Image         []uploadImage.Image `json:"image" gorm:"serializer:json;column:image"`
}

type Houses struct {
	Id          int                 `json:"id" gorm:"primarykey"`
	Price       float64             `json:"price" gorm:"column:price"`
	Name        string              `json:"name" gorm:"column:name"`
	ChineseName string              `json:"chineseName" gorm:"column:chinese_name"`
	Image       []uploadImage.Image `json:"image" gorm:"serializer:json;column:image"`
}

type News struct {
	Id      int                 `json:"id" gorm:"primarykey"`
	Title   string              `json:"title" gorm:"column:title"`
	Image   []uploadImage.Image `json:"image" gorm:"serializer:json;column:image"`
	Content string              `json:"content" gorm:"column:content"`
}
