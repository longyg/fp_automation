package nidd

import (
	"net/url"
	"net/http"
	"fmt"
	"errors"
	"../httputil"
)

func Login(loginUrl string, username string, password string) error  {
	values := make(url.Values)
	values.Set("login", username)
	values.Set("passwd", password)
	resp, err := http.PostForm(loginUrl, values)

	if nil != err {
		return errors.New(fmt.Sprintln("Error while logging to NIDD:", err.Error()))
	}

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintln("Failed to login NIDD, return code:", resp.StatusCode))
	}

	err = httputil.SetAuthContext(resp)
	if err != nil {
		return errors.New(fmt.Sprintln("Error while setting auth context:", err.Error()))
	}
	return nil
}
