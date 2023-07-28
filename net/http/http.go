package http

import (
	httpUtil "github.com/go-resty/resty/v2"
	"time"
)

const (
	ContentTypeHead = "Content-Type"
	AcceptHead      = "Accept"
	ApplicationJSON = "application/json"
)

// BuildDefaultClient
// @Description: 获取http客户端
// @return *httpUtil.Client
func BuildDefaultClient() *httpUtil.Client {
	return httpUtil.New().SetTimeout(30 * time.Second)
}

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
	return BuildDefaultClient().R()
}

// Get
// @Description: Get请求
// @param: url 请求地址
// @param: params 请求参数
// @return T 返回对象
// @return error
func Get[T any](url string, params map[string]string) (T, error) {
	var result T
	_, err := BuildRequest().
		SetQueryParams(params).
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetHeader(AcceptHead, ApplicationJSON).
		SetResult(&result).
		Get(url)
	return result, err
}

// GetString
// @Description: Get请求
// @param: url 请求地址
// @param: params 请求参数
// @return string 返回字符串
// @return error
func GetString(url string, params map[string]string) (string, error) {
	resp, err := BuildRequest().
		SetQueryParams(params).
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetHeader(AcceptHead, ApplicationJSON).
		Get(url)
	return resp.String(), err
}

// GetQuery
// @Description: Get请求
// @param: url
// @param: params：productId=232&template=2 这样形式
// @return T
// @return error
func GetQuery[T any](url string, params string) (T, error) {
	var result T
	_, err := BuildRequest().
		SetQueryString(params).
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetHeader(AcceptHead, ApplicationJSON).
		SetResult(&result).
		Get(url)
	return result, err
}

// GetQueryString
// @Description: Get请求
// @param: url 请求地址
// @param: params 请求参数
// @return string 返回字符串
// @return error
func GetQueryString(url string, params string) (string, error) {
	resp, err := BuildRequest().
		SetQueryString(params).
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetHeader(AcceptHead, ApplicationJSON).
		Get(url)
	return resp.String(), err
}

// Post
// @Description: Post请求
// @param: url
// @param: body
// @return T
// @return error
func Post[T any](url string, body interface{}) (T, error) {
	var result T
	_, err := BuildRequest().
		SetBody(body).
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetResult(&result).
		Post(url)
	return result, err
}

// PostString
// @Description: Post请求
// @param: url 请求地址
// @param: body 请求参数
// @return string 返回字符串
// @return error
func PostString(url string, body interface{}) (string, error) {
	resp, err := BuildRequest().
		SetBody(body).
		SetHeader(ContentTypeHead, ApplicationJSON).
		Post(url)
	return resp.String(), err
}

// PostForm
// @Description: PostForm请求
// @param: url
// @param: formParam
// @return T
// @return error
func PostForm[T any](url string, formParam map[string]string) (T, error) {
	var result T
	_, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetFormData(formParam).
		SetResult(&result).
		Post(url)
	return result, err
}

// PostFormString
// @Description: PostForm请求
// @param: url 请求地址
// @param: formParam 请求参数
// @return string 返回字符串
// @return error
func PostFormString(url string, formParam map[string]string) (string, error) {
	resp, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetFormData(formParam).
		Post(url)
	return resp.String(), err
}

// PostFile
// @Description: PostFile请求
// @param: url 请求地址
// @param: fileName 文件名
// @param: filePath 文件路径
// @return T 返回对象
// @return error
func PostFile[T any](url, fileName, filePath string) (T, error) {
	var result T
	_, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetFile(fileName, filePath).
		SetResult(&result).
		Post(url)
	return result, err
}

// PostFileString
// @Description: PostFile请求
// @param: url 请求地址
// @param: fileName 文件名
// @param: filePath 文件路径
// @return string 返回字符串
// @return error
func PostFileString(url, fileName, filePath string) (string, error) {
	resp, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetFile(fileName, filePath).
		Post(url)
	return resp.String(), err
}

// PostFiles
// @Description: PostFiles请求
// @param: url 请求地址
// @param: files 文件
// @return T 返回对象
// @return error
func PostFiles[T any](url string, files map[string]string) (T, error) {
	var result T
	_, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetFiles(files).
		SetResult(&result).
		Post(url)
	return result, err
}

// PostFilesString
// @Description: PostFiles请求
// @param: url 请求地址
// @param: files 文件
// @return string 返回字符串
// @return error
func PostFilesString(url string, files map[string]string) (string, error) {
	resp, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetFiles(files).
		Post(url)
	return resp.String(), err
}

// PostFormFile
// @Description: PostFormFile请求
// @param: url
// @param: fileName
// @param: filePath
// @param: formParam
// @return T
// @return error
func PostFormFile[T any](url, fileName, filePath string, formParam map[string]string) (T, error) {
	var result T
	_, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetFormData(formParam).
		SetFile(fileName, filePath).
		SetResult(&result).
		Post(url)
	return result, err
}

// PostFormFileString
// @Description: PostFormFile请求
// @param: url 请求地址
// @param: fileName 文件名
// @param: filePath 文件路径
// @param: formParam 请求参数
// @return string 返回字符串
// @return error
func PostFormFileString(url, fileName, filePath string, formParam map[string]string) (string, error) {
	resp, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetFormData(formParam).
		SetFile(fileName, filePath).
		Post(url)
	return resp.String(), err
}

// PostFormFiles
// @Description: PostFormFiles请求
// @param: url
// @param: formParam
// @param: files
// @return T
// @return error
func PostFormFiles[T any](url string, formParam, files map[string]string) (T, error) {
	var result T
	_, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetFormData(formParam).
		SetFiles(files).
		SetResult(&result).
		Post(url)
	return result, err
}

// PostFormFilesString
// @Description: PostFormFiles请求
// @param: url 请求地址
// @param: formParam 请求参数
// @param: files 文件
// @return string 返回字符串
// @return error
func PostFormFilesString(url string, formParam, files map[string]string) (string, error) {
	resp, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetFormData(formParam).
		SetFiles(files).
		Post(url)
	return resp.String(), err
}

// Put
// @Description: Put请求
// @param: url
// @param: body
// @return T
// @return error
func Put[T any](url string, body interface{}) (T, error) {
	var result T
	_, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetResult(&result).
		SetBody(body).
		Put(url)
	return result, err
}

// PutString
// @Description: Put请求
// @param: url 请求地址
// @param: body 请求参数
// @return string 返回字符串
// @return error
func PutString(url string, body interface{}) (string, error) {
	resp, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetBody(body).
		Put(url)
	return resp.String(), err
}

// Patch
// @Description: Patch请求
// @param: url
// @param: body
// @return T
// @return error
func Patch[T any](url string, body interface{}) (T, error) {
	var result T
	_, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetResult(&result).
		SetBody(body).
		Patch(url)
	return result, err
}

// PatchString
// @Description: Patch请求
// @param: url 请求地址
// @param: body 请求参数
// @return string 返回字符串
// @return error
func PatchString(url string, body interface{}) (string, error) {
	resp, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetBody(body).
		Patch(url)
	return resp.String(), err
}

// Delete
// @Description: Delete请求
// @param: url
// @param: body
// @return T
// @return error
func Delete[T any](url string, body interface{}) (T, error) {
	var result T
	_, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetResult(&result).
		SetBody(body).
		Delete(url)
	return result, err
}

// DeleteString
// @Description: Delete请求
// @param: url 请求地址
// @param: body 请求参数
// @return string 返回字符串
// @return error
func DeleteString(url string, body interface{}) (string, error) {
	resp, err := BuildRequest().
		SetHeader(ContentTypeHead, ApplicationJSON).
		SetBody(body).
		Delete(url)
	return resp.String(), err
}
