package libinstall

import (
	"github.com/4thel00z/libhttp"
	"github.com/mileusna/useragent"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

const (
	Windows      = "windows"
	WindowsPhone = "windows phone"
	Android      = "android"
	MacOS        = "macos"
	IOS          = "ios"
	Linux        = "linux"
)

func TemplatesFromPath(path string) (map[string]string, error) {
	result := map[string]string{}
	files, err := filepath.Glob(path + "/*")

	if err != nil {
		return nil, err
	}

	for _, p := range files {
		_, file := filepath.Split(p)
		loweredFile := strings.ToLower(file)
		if strings.Contains(loweredFile, Linux) {
			content, err := ioutil.ReadFile(p)
			if err != nil {
				log.Println("Could not read file:", p)
				continue
			}
			result[Linux] = string(content)
		}

		if strings.Contains(loweredFile, Windows) {
			content, err := ioutil.ReadFile(p)
			if err != nil {
				log.Println("Could not read file:", p)
				continue
			}
			result[Windows] = string(content)
		}

		if strings.Contains(loweredFile, MacOS) {
			content, err := ioutil.ReadFile(p)
			if err != nil {
				log.Println("Could not read file:", p)
				continue
			}
			result[MacOS] = string(content)
		}

		if strings.Contains(loweredFile, WindowsPhone) {
			content, err := ioutil.ReadFile(p)
			if err != nil {
				log.Println("Could not read file:", p)
				continue
			}
			result[WindowsPhone] = string(content)
		}

		if strings.Contains(loweredFile, Android) {
			content, err := ioutil.ReadFile(p)
			if err != nil {
				log.Println("Could not read file:", p)
				continue
			}
			result[Android] = string(content)
		}

		if strings.Contains(loweredFile, IOS) {
			content, err := ioutil.ReadFile(p)
			if err != nil {
				log.Println("Could not read file:", p)
				continue
			}
			result[IOS] = string(content)
		}

	}
	return result, nil
}

func ParseOS(req libhttp.Request) string {
	return ua.Parse(req.Header.Get("User-Agent")).OS
}

func OSHandler(templates map[string]string) libhttp.Service {
	return func(req libhttp.Request) libhttp.Response {
		os := strings.ToLower(ParseOS(req))

		t, ok := templates[os]
		if !ok {
			res := StringResponse(req, "# We do not support your operation system at this point!")
			res.StatusCode = http.StatusNotFound
			return res
		}

		return StringResponse(req, t)
	}

}

func StringResponse(req libhttp.Request, t string) libhttp.Response {
	res := req.Response(nil)
	res.Header.Set("Content-Type", "text/plain")
	res.Body = io.NopCloser(strings.NewReader(t))
	return res
}
