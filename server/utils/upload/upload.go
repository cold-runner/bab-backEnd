package upload

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/uploadImage"
	"mime/multipart"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// OSS 对象存储接口
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

// NewOss OSS的实例化方法
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
func NewOss() OSS {
	switch global.GVA_CONFIG.System.OssType {
	case "local":
		return &Local{}
	case "qiniu":
		return &Qiniu{}
	case "tencent-cos":
		return &TencentCOS{}
	case "aliyun-oss":
		return &AliyunOSS{}
	case "huawei-obs":
		return HuaWeiObs
	case "aws-s3":
		return &AwsS3{}
	default:
		return &Local{}
	}
}

func ProcessUpdateImages(old uploadImage.Images, newer []*multipart.FileHeader) (finalImages uploadImage.Images) {
	var needUploads []*multipart.FileHeader
	oldImages := map[string]string{}
	for _, v := range old {
		oldImages[v.Name] = v.URL
	}
	oss := NewOss()
	for _, v := range newer {
		if url, ok := oldImages[v.Filename]; ok {
			finalImages = append(finalImages, uploadImage.Image{v.Filename, url})
			delete(oldImages, v.Filename)
		} else {
			needUploads = append(needUploads, v)
		}
	}

	for i := range needUploads {
		url, name, _ := oss.UploadFile(needUploads[i])
		finalImages = append(finalImages, uploadImage.Image{name, url})
	}

	for k := range oldImages {
		oss.DeleteFile(k)
	}

	if len(finalImages) == 0 {
		finalImages = append(finalImages, uploadImage.Image{})
	}

	return
}

func ProcessUpdateImage(old *uploadImage.Image, newer *multipart.FileHeader) uploadImage.Image {
	oss := NewOss()
	if old.Name == newer.Filename {
		return *old
	}
	url, name, _ := oss.UploadFile(newer)
	oss.DeleteFile(old.Name)
	return uploadImage.Image{name, url}
}
