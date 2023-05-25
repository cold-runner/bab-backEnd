package private

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babApartment"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babCity"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babCompany"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"mime/multipart"
)

type PrivateService struct {
}

func (p *PrivateService) Upload(f *multipart.FileHeader) (string, string, error) {
	oss := upload.NewOss()
	return oss.UploadFile(f)
}

func (p *PrivateService) Delete(fileName string) error {
	oss := upload.NewOss()
	return oss.DeleteFile(fileName)
}

func (p *PrivateService) GetApartmentList() (apartments []babApartment.BabApartment, err error) {
	err = global.GVA_DB.Find(&apartments).Error
	return
}

func (p *PrivateService) GetCityList() (cities []babCity.BabCity, err error) {
	err = global.GVA_DB.Find(&cities).Error
	return
}

func (p *PrivateService) GetCompaniesList() (companies []babCompany.BabCompany, err error) {
	err = global.GVA_DB.Find(&companies).Error
	return
}
