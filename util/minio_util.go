package util

import (
	"github.com/minio/minio-go/v6"
	"github.com/spf13/viper"
	_ "io/ioutil"
	"mime/multipart"
)

func UploadToMinio(objectName string, file *multipart.FileHeader, contentType string) error {
	// 初使化 minio client对象
	minioClient, err := minio.New(viper.GetString("minio.endpoint"), viper.GetString("minio.accessKey"), viper.GetString("minio.secretKey"), false)
	if err != nil {
		return err
	}

	src, err1 := file.Open()
	if err1 != nil {
		return err1
	}
	// 获取图片文件的信息
	defer src.Close()

	//
	_, err2 := minioClient.PutObject(viper.GetString("minio.bucketName"),
		objectName, src, file.Size, minio.PutObjectOptions{})
	if err2 != nil {
		return err2
	}

	return nil
}
