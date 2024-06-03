package router

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/mkolchurin/crosspick_server/appconfig"
)

func InitS3(r *gin.Engine, cfg *appconfig.AppConfig) {
	minioClient, err := minio.New(cfg.S3.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.S3.AccessKeyID, cfg.S3.SecretAccessKey, ""),
		Secure: cfg.S3.UseSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	r.GET("/api/v1/storage/:bucket/*id", func(c *gin.Context) {
		bucket := c.Param("bucket")
		id := c.Param("id")

		if id == "/" {
			obj := minioClient.ListObjects(context.Background(), bucket, minio.ListObjectsOptions{})
			var paths []string
			for {
				info, ok := <-obj
				if !ok {
					break
				}
				paths = append(paths, info.Key)
			}
			c.JSON(http.StatusOK, gin.H{bucket: paths})
			return
		}
		id = id[1:]
		obj, err := minioClient.GetObject(context.Background(), bucket, id, minio.GetObjectOptions{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer obj.Close()
		info, err := obj.Stat()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.DataFromReader(http.StatusOK, info.Size, info.ContentType, obj, nil)
	})
}
