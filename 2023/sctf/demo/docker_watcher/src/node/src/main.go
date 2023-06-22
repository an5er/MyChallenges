package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func CORS(ctx *context.Context) {
	ctx.Output.Header("Access-Control-Allow-Origin", "*")
	ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Output.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	if ctx.Request.Method == "OPTIONS" {
		ctx.ResponseWriter.WriteHeader(http.StatusNoContent)
	}
}

type UploadController struct {
	beego.Controller
}
type HomeController struct {
	beego.Controller
}

func (c *UploadController) Post() {
	file, header, err := c.GetFile("file")
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Failed to get file",
		}
		c.ServeJSON()
		return
	}
	defer file.Close()

	uploadDir := "./upload"
	customPath := c.GetString("path")
	uploadDir = filepath.Join(uploadDir, customPath)
	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Failed to create upload directory",
		}
		c.ServeJSON()
		return
	}

	filePath := filepath.Join(uploadDir, header.Filename)

	saveFile, err := os.Create(filePath)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Failed to create save file",
		}
		c.ServeJSON()
		return
	}
	defer saveFile.Close()

	_, err = io.Copy(saveFile, file)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Failed to save file",
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("File %s uploaded successfully", header.Filename),
	}
	c.ServeJSON()
}

func (c HomeController) Get() {
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Hello SCTF"),
	}
	c.ServeJSON()
}
func main() {
	beego.InsertFilter("*", beego.BeforeRouter, CORS)
	beego.Router("/api/node", &HomeController{})
	beego.Router("/api/node/upload", &UploadController{})
	beego.Run("")
}
