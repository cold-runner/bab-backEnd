package initialize

import (
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"github.com/flipped-aurora/gin-vue-admin/server/model/babAboutUs"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babApartment"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babCity"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babCompany"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babEnquiry"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babHouse"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babIdxImage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babNews"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},
		system.SysChatGptOption{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},

		babCity.BabCity{},
		babApartment.BabApartment{},
		babCompany.BabCompany{},
		babNews.BabNews{},
		babHouse.BabHouse{},
		babEnquiry.BabEnquiry{}, babIdxImage.BabIdxImage{}, babAboutUs.BabAboutUs{},
	)

	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
