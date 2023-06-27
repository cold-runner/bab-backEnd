package public

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babAboutUs"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babEnquiry"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babHouse"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babIdxImage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/public"
)

type PublicService struct {
}

func (p *PublicService) HomeInfo() (res []public.ApartmentCountCity, err error) {
	err = global.GVA_DB.Raw(`
SELECT 
    c.name AS city,
    c.chinese_name AS city_chinese_name,
    COUNT(a.id) AS apartment_count,
    c.image AS image
FROM 
    bab_city c
LEFT JOIN 
    bab_apartment a ON c.name = a.city
GROUP BY 
    c.name, c.image, c.chinese_name
`).Scan(&res).Error
	return
}

func (p *PublicService) ApartmentsByCity(city string) (res []public.ApartmentsByCity, err error) {
	err = global.GVA_DB.Raw(`SELECT
    bab_apartment.name apartment_name,
    bab_apartment.company company,
    bab_apartment.introduction introduction,
    bab_apartment.type type,
    bab_apartment.image image,
    MIN(bab_house.price) AS price
FROM
    bab_apartment
        JOIN bab_house ON bab_apartment.name = bab_house.apartment AND bab_apartment.city = bab_house.city
WHERE
        bab_apartment.city = ? AND bab_apartment.deleted_at IS NULL AND bab_house.deleted_at IS NULL
GROUP BY
    bab_apartment.id;
`, city).Scan(&res).Error
	return
}

func (p *PublicService) HousesByApartment(apartment string) (res []public.Houses, err error) {
	err = global.GVA_DB.Raw(`
SELECT id, name, chinese_name, price, image
FROM bab_house
WHERE apartment = ? 
  AND deleted_at IS NULL;
`, apartment).Scan(&res).Error
	return
}

func (p *PublicService) News() (res []public.News, err error) {
	err = global.GVA_DB.Raw(`
SELECT id, title, image, content
FROM bab_news
WHERE deleted_at IS NULL;
`).Scan(&res).Error
	return
}

func (p *PublicService) SingleHouse(id string) (house babHouse.BabHouse, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&house).Error
	return
}

func (p *PublicService) AboutUs() (res babAboutUs.BabAboutUs, err error) {
	err = global.GVA_DB.First(&res).Error
	return
}

func (p *PublicService) HomePageUrl() (urls []string, err error) {
	res := babIdxImage.BabIdxImage{}
	err = global.GVA_DB.First(&res).Error
	for i := range res.Hp {
		urls = append(urls, res.Hp[i].URL)
	}
	return
}

func (p *PublicService) EnquiryImageUrl() (urls []string, err error) {
	res := babIdxImage.BabIdxImage{}
	err = global.GVA_DB.First(&res).Error
	for i := range res.Enq {
		urls = append(urls, res.Enq[i].URL)
	}
	return
}

func (p *PublicService) NewsUrl() (urls []string, err error) {
	res := babIdxImage.BabIdxImage{}
	err = global.GVA_DB.First(&res).Error
	for i := range res.Ne {
		urls = append(urls, res.Ne[i].URL)
	}
	return
}

func (p *PublicService) CreateUserEnquiry(enquiry *babEnquiry.BabEnquiry) (err error) {
	err = global.GVA_DB.Create(enquiry).Error
	return err
}
