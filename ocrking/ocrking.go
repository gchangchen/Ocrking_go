package ocrking

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type Ocrking struct {
	Service, Language, Type, Charset, Apikey string
}

func (p *Ocrking) Parse(f io.Reader) (result string, err error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("image", "code.jpg")
	if err != nil {
		return
	}
	if _, err = io.Copy(fw, f); err != nil {
		return
	}
	// Add the other fields
	//Service, Language, Type, Charset, Apikey string
	w.WriteField("service", p.Service)
	w.WriteField("language", p.Language)
	w.WriteField("type", p.Type)
	w.WriteField("charset", p.Charset)
	w.WriteField("apiKey", p.Apikey)
	w.Close()

	req, err := http.NewRequest("POST", "http://lab.ocrking.com/ok.html", &b)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("Http : %s", res.Status)
		return
	}

	body := xml.NewDecoder(res.Body)
	var resultcode struct {
		ResultList struct {
			Item struct {
				Result string
				Status bool
			}
		}
	}
	err = body.Decode(&resultcode)
	res.Body.Close()
	if err != nil {
		return
	}

	if !resultcode.ResultList.Item.Status {
		err = fmt.Errorf("Xml : %s", resultcode.ResultList.Item.Result)
		return
	}
	return resultcode.ResultList.Item.Result, nil
}
