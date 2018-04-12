package httputil

import (
	"net/url"
	"net/http"
	"fmt"
	"net/http/cookiejar"
	"strings"
	"io/ioutil"
	"bytes"
	"errors"
	"os"
	"mime/multipart"
)

type AuthContext struct {
	authCookie *http.Cookie
}

var authContext = AuthContext{}

func PostRequest(postUrl string, postData url.Values) error {
	if nil == authContext.authCookie {
		return errors.New("Auth cookie is empty, probably you haven't logged in NIDD")
	}

	var cookies []*http.Cookie
	cookies = append(cookies, authContext.authCookie)

	u, err := url.Parse(postUrl)
	if nil != err {
		return errors.New(fmt.Sprintln("Error while parsing post url:", err.Error()))
	}
	jar, err := cookiejar.New(nil)
	if nil != err {
		return errors.New(fmt.Sprintln("Error while creating cookiejar:", err.Error()))
	}
	jar.SetCookies(u, cookies)

	client := http.Client{Jar: jar}

	req, err := http.NewRequest("POST", postUrl, strings.NewReader(postData.Encode()))
	if nil != err {
		return errors.New(fmt.Sprintln("Error while creating HTTP request:", err.Error()))
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	printRequestBody(req)

	resp, err := client.Do(req)

	if err != nil {
		return errors.New(fmt.Sprintln("Error while doing HTTP request:", err.Error()))
	}

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintln("Post request failed. Status:", resp.Status))
	}

	err = SetAuthContext(resp)
	if err != nil {
		return errors.New(fmt.Sprintln("Error while setting auth context:", err.Error()))
	}
	return nil
}

func SetAuthContext(resp *http.Response) error {
	cookie := getAuthCookie(resp)
	if nil == cookie {
		return errors.New(".ASPXAUTH cookie is not found from HTTP response")
	}
	authContext.authCookie = cookie
	return nil
}

func getAuthCookie(resp *http.Response) *http.Cookie {
	if len(resp.Cookies()) > 0 {
		for _, cookie := range resp.Cookies() {
			if cookie.Name == ".ASPXAUTH" {
				return cookie
			}
		}
	}
	return nil
}

func printRequestBody(req *http.Request) {
	if nil == req.Body {
		fmt.Println("Request body is empty")
		return
	}
	buf, bodyErr := ioutil.ReadAll(req.Body)
	if bodyErr != nil {
		fmt.Println("read request body error: ", bodyErr.Error())
		return
	}

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	fmt.Printf("BODY: %q \n", rdr1)
	req.Body = rdr2
}

func newFileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return http.NewRequest("POST", uri, body)
}

func UploadFile(postUrl string, params map[string]string, paramName, path string) error {
	if nil == authContext.authCookie {
		return errors.New("Auth cookie is empty, probably you haven't logged in NIDD")
	}

	var cookies []*http.Cookie
	cookies = append(cookies, authContext.authCookie)

	u, err := url.Parse(postUrl)
	if nil != err {
		return errors.New(fmt.Sprintln("Error while parsing post url:", err.Error()))
	}
	jar, err := cookiejar.New(nil)
	if nil != err {
		return errors.New(fmt.Sprintln("Error while creating cookiejar:", err.Error()))
	}
	jar.SetCookies(u, cookies)

	client := http.Client{Jar: jar}

	req, err := newFileUploadRequest(postUrl, params, paramName, path)
	if nil != err {
		return errors.New(fmt.Sprintln("Error while creating HTTP request:", err.Error()))
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	printRequestBody(req)

	resp, err := client.Do(req)

	if err != nil {
		return errors.New(fmt.Sprintln("Error while doing HTTP request:", err.Error()))
	}

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintln("Post request failed. Status:", resp.Status))
	}

	err = SetAuthContext(resp)
	if err != nil {
		return errors.New(fmt.Sprintln("Error while setting auth context:", err.Error()))
	}
	return nil
}