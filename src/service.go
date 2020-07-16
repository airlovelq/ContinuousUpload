package main

import (
	"fmt"
	"io"
	"strconv"

	//	"io"
	"net/http"
	"os"
	"syscall"

	//	"strconv"

	//"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func GetEnvDefault(key string, default_value string) string {
	if res := os.Getenv(key); res != "" {
		return res
	} else {
		return default_value
	}
}

func formatResponse(c *gin.Context, success int, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"message": message,
		"data":    data,
	})
}

func status(c *gin.Context) {
	file_name := c.Query("file_name")
	if file_name == "" {
		file_name = uuid.NewV4().String()
	}
	file, err := os.OpenFile(GetEnvDefault("DOCKER_DATA_DIR", "/home/luqin2/tmp")+"/"+file_name, os.O_RDWR|os.O_CREATE, 0766)
	defer file.Close()
	if err != nil {
		formatResponse(c, 1, "Create file error", nil)
		// c.JSON(http.StatusOK, gin.H{
		// 	"success": 1,
		// 	"message": "Create file error",
		// 	"data":    nil,
		// })
		return
	}
	syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
	file_size, _ := file.Seek(0, os.SEEK_END)
	formatResponse(c, 0, "Success", gin.H{
		"file_size": file_size,
		"file_name": file_name,
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"success": 0,
	// 	"message": "Success",
	// 	"data": gin.H{
	// 		"file_size": file_size,
	// 		"file_name": file_name,
	// 	},
	// })
}

func upload(c *gin.Context) {
	file_size, ok := c.GetPostForm("file_size")
	if !ok {
		formatResponse(c, 1, "Param Error", nil)
		// c.String(http.StatusOK, "param error!")
		return
	}
	nfile_size, err := strconv.Atoi(file_size)
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		formatResponse(c, 1, "Param Error", nil)
		// c.JSON(http.StatusOK, gin.H{
		// 	"success": 1,
		// 	"message": "Param Error",
		// 	"data":    nil,
		// })
		return
	}

	pfile, err := os.OpenFile(GetEnvDefault("DOCKER_DATA_DIR", "/home/luqin2/tmp")+"/"+header.Filename, os.O_RDWR|os.O_APPEND, 0766)
	defer pfile.Close()
	if err != nil {
		formatResponse(c, 1, "Open File Error", nil)
		// c.JSON(http.StatusOK, gin.H{
		// 	"success": 1,
		// 	"message": "Open File Error",
		// 	"data":    nil,
		// })
		return
	}
	// length, _ := strconv.Atoi(header.Header.Get("Content-length"))
	// fmt.Println(length)
	// by := make([]byte, length)
	// file.Read(by)
	syscall.Flock(int(pfile.Fd()), syscall.LOCK_EX)
	pfile.Seek(0, os.SEEK_END)
	st, err := os.Stat(GetEnvDefault("DOCKER_DATA_DIR", "/home/luqin2/tmp") + "/" + header.Filename)
	if err != nil {
		formatResponse(c, 1, "Get File Stat Error", nil)
		// c.JSON(http.StatusOK, gin.H{
		// 	"success": 1,
		// 	"message": "Get File Stat Error",
		// 	"data":    nil,
		// })
		return
	}
	if st.Size() != int64(nfile_size) {
		formatResponse(c, 1, "Start pos not same", gin.H{
			"file_size": st.Size(),
		})
		// c.JSON(http.StatusOK, gin.H{
		// 	"success": 1,
		// 	"message": "Start pos not same",
		// 	"data": gin.H{
		// 		"file_size": st.Size(),
		// 	},
		// })
		return
	}
	_, err = io.Copy(pfile, file)
	if err != nil {
		formatResponse(c, 1, "Write File Error", gin.H{
			"file_size": st.Size(),
		})
		// c.JSON(http.StatusOK, gin.H{
		// 	"success": 1,
		// 	"message": "Write File Error",
		// 	"data": gin.H{
		// 		"file_size": st.Size(),
		// 	},
		// })
		return
	}

	st, err = os.Stat(GetEnvDefault("DOCKER_DATA_DIR", "/home/luqin2/tmp") + "/" + header.Filename)
	if err != nil {
		formatResponse(c, 1, "Get File Stat Error", gin.H{
			"file_size": nfile_size,
		})
		// c.JSON(http.StatusOK, gin.H{
		// 	"success": 1,
		// 	"message": "Get File Stat Error",
		// 	"data": gin.H{
		// 		"file_size": nfile_size,
		// 	},
		// })
		return
	}

	//log.Println(file.Filename)
	// fmt.Println(file_name)
	// fmt.Println(file)
	// 上传文件到指定的路径
	//c.SaveUploadedFile(file, dst)
	formatResponse(c, 0, "Success", gin.H{
		"file_size": st.Size(),
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"success": 0,
	// 	"message": "Success",
	// 	"data": gin.H{
	// 		"file_size": st.Size(),
	// 	},
	// })

	// c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", header.Filename))
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func StartService() {
	router := gin.Default()
	router.Use(Cors())
	router.POST("/upload", upload)
	router.GET("/upload/status", status)
	fmt.Println("Service Started!")
	router.Run(":" + GetEnvDefault("DOCKER_PORT", "8908"))
}
