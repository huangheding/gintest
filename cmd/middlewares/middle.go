package middlewares

import (
	"bytes"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const DefaultMemory = 32 * 1024 * 1024

//一些中间件 备用
// GetHeaders ...
func GetHeaders(head http.Header) map[string]string {
	hdr := make(map[string]string, len(head))
	for k, v := range head {
		hdr[k] = v[0]
	}
	return hdr
}

// GetIP ...
func GetIP(c *gin.Context) string {
	ip := c.ClientIP()
	return ip
}

// GetMultiPartFormValue ...
func GetMultiPartFormValue(c *http.Request) interface{} {
	var requestBody interface{}

	multipartForm := make(map[string]interface{})
	if err := c.ParseMultipartForm(DefaultMemory); err != nil {
		// handle error
	}
	if c.MultipartForm != nil {
		for key, values := range c.MultipartForm.Value {
			multipartForm[key] = strings.Join(values, "")
		}

		for key, file := range c.MultipartForm.File {
			for k, f := range file {
				formKey := fmt.Sprintf("%s%d", key, k)
				multipartForm[formKey] = map[string]interface{}{"filename": f.Filename, "size": f.Size}
			}
		}

		if len(multipartForm) > 0 {
			requestBody = multipartForm
		}
	}
	return requestBody
}

// GetFormBody ...
func GetFormBody(c *http.Request) interface{} {
	var requestBody interface{}

	form := make(map[string]string)
	if err := c.ParseForm(); err != nil {
		// handle error
	}
	for key, values := range c.PostForm {
		form[key] = strings.Join(values, "")
	}
	if len(form) > 0 {
		requestBody = form
	}

	return requestBody
}

// GetRequestBody ...
func GetRequestBody(c *gin.Context) interface{} {
	//multiPartFormValue := GetMultiPartFormValue(c.Request)
	//if multiPartFormValue != nil {
	//	return multiPartFormValue
	//}
	//
	//formBody := GetFormBody(c.Request)
	//if formBody != nil {
	//	return formBody
	//}

	method := c.Request.Method
	if method == "GET" {
		return nil
	}
	contentType := c.ContentType()
	body := c.Request.Body
	var model interface{}
	bodyContent, err := ioutil.ReadAll(body)
	if err != nil {
		return model
	}
	// Restore the io.ReadCloser to its original state
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyContent))

	switch contentType {
	case binding.MIMEJSON:
		json.Unmarshal(bodyContent, &model)
		return model
	default:
		model = string(bodyContent)
		return model
	}
}
