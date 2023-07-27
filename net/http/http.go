package http

import (
	httpUtil "github.com/go-resty/resty/v2"
)

const (
	ContentTypeHead = "Content-Type"
	AcceptHead      = "Accept"
	ApplicationJSON = "application/json"
)

// BuildClient
// @Description: 获取http客户端
// @return *httpUtil.Client
func BuildClient() *httpUtil.Client {
	return httpUtil.New()
}

// BuildRequest
// @Description: 获取http请求
// @return *httpUtil.Request
func BuildRequest() *httpUtil.Request {
	return BuildClient().R()
}

// Get
// @Description: Get请求
// @param: url
// @param: params
// @return T
// @return error
func Get[T comparable](url string, params map[string]string) (T, error) {
	var result T
	_, err := BuildRequest().
		SetQueryParams(params).
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetHeader(AcceptHead, ApplicationJSON).
		SetResult(&result).
		Get(url)
	return result, err
}

// GetQueryString
// @Description: Get请求
// @param: url
// @param: params productId=232&template=2这样形式
// @return T
// @return error
func GetQueryString[T comparable](url string, params string) (T, error) {
	var result T
	_, err := BuildRequest().
		SetQueryString(params).
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetHeader(AcceptHead, ApplicationJSON).
		SetResult(&result).
		Get(url)
	return result, err
}

// Post
// @Description: Post请求
// @param: url
// @param: body
// @return T
// @return error
func Post[T comparable](url string, body interface{}) (T, error) {
	var result T
	_, err := BuildRequest().
		SetBody(body).
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetResult(&result).
		Post(url)
	return result, err
}

// PostForm
// @Description: PostForm请求
// @param: url
// @param: formParam
// @return T
// @return error
func PostForm[T comparable](url string, formParam map[string]string) (T, error) {
	var result T
	_, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetFormData(formParam).
		SetResult(&result).
		Post(url)
	return result, err
}

// PostFile
// @Description: PostFile请求
// @param: url
// @param: fileName
// @param: filePath
// @return T
// @return error
func PostFile[T comparable](url, fileName, filePath string) (T, error) {
	var result T
	_, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetFile(fileName, filePath).
		SetResult(&result).
		Post(url)
	return result, err
}

// PostFiles
// @Description: PostFiles请求
// @param: url
// @param: files
// @return T
// @return error
func PostFiles[T comparable](url string, files map[string]string) (T, error) {
	var result T
	_, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetFiles(files).
		SetResult(&result).
		Post(url)
	return result, err
}

// PostFormFile
// @Description: PostFormFile请求
// @param: url
// @param: fileName
// @param: filePath
// @param: formParam
// @return T
// @return error
func PostFormFile[T comparable](url, fileName, filePath string, formParam map[string]string) (T, error) {
	var result T
	_, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetFormData(formParam).
		SetFile(fileName, filePath).
		SetResult(&result).
		Post(url)
	return result, err
}

// PostFormFiles
// @Description: PostFormFiles请求
// @param: url
// @param: formParam
// @param: files
// @return T
// @return error
func PostFormFiles[T comparable](url string, formParam, files map[string]string) (T, error) {
	var result T
	_, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetFormData(formParam).
		SetFiles(files).
		SetResult(&result).
		Post(url)
	return result, err
}

// Put
// @Description: Put请求
// @param: url
// @param: body
// @return T
// @return error
func Put[T comparable](url string, body interface{}) (T, error) {
	var result T
	_, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetResult(&result).
		SetBody(body).
		Put(url)
	return result, err
}

// Patch
// @Description: Patch请求
// @param: url
// @param: body
// @return T
// @return error
func Patch[T comparable](url string, body interface{}) (T, error) {
	var result T
	_, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetResult(&result).
		SetBody(body).
		Patch(url)
	return result, err
}

// Delete
// @Description: Delete请求
// @param: url
// @param: body
// @return T
// @return error
func Delete[T comparable](url string, body interface{}) (T, error) {
	var result T
	_, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetResult(&result).
		SetBody(body).
		Delete(url)
	return result, err
}
