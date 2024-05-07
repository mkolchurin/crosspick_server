package siteRouter

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	endpoint        = "192.168.0.111:9000"
	accessKeyID     = "j5oJsXZ0B9smNeeMUDLY"
	secretAccessKey = "tJyPtxrRs3JPnl3NnFyfXz0iYYKTCGdwDQ9KwB0N"
	useSSL          = false
)

func AddMinio(r *gin.Engine) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
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
