package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"regexp"

	"github.com/arekkusu66/goutils/fsys"
)


type Request struct {
	Unm					any
	Output				string
	ContentType			string
}


//Makes a POST request, if ContentType isn't specified, it defaults to application/json
func (rq *Request) POST(url string, body any) ([]byte, error) {

    jsonData, err := json.Marshal(body)

    if err != nil {
		return nil, err
    }


	var contentType string

	if rq.ContentType == "" {
		contentType = "application/json"
	} else {
		contentType = rq.ContentType
	}

    response, err := http.Post(url , contentType , bytes.NewBuffer(jsonData))

    if err != nil {
		return nil, err
    }

	defer response.Body.Close()


	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}


	if rq.Output != "" {
		var file = &fsys.FileSys{}

		if regexp.MustCompile(`\.json$`).MatchString(rq.Output) {
			var mapUnm = make(map[string]interface{})

			if err = json.Unmarshal(responseData, &mapUnm); err != nil {
				return nil, err
			}

			d, err := json.MarshalIndent(mapUnm, "", "  ")

			if err != nil {
				return nil, err
			}

			if err = file.Write(d, rq.Output); err != nil {
				return nil, err
			}

		} else {
			if err = file.Write(responseData, rq.Output); err != nil {
				return nil, err
			}
		}
	}


	if rq.Unm != nil {
        if err = json.Unmarshal(responseData, &rq.Unm); err != nil {
            return nil, err
        }
    }


	return responseData, nil
}

//GET request
func (rq *Request) GET(url string) ([]byte, error) {

	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}
	


	if rq.Output != "" {
		var file = &fsys.FileSys{}

		if regexp.MustCompile(`\.json$`).MatchString(rq.Output) {
			var mapUnm = make(map[string]interface{})

			err = json.Unmarshal(body, &mapUnm)

			if err != nil {
				return nil, err
			}

			d, err := json.MarshalIndent(mapUnm, "", "  ")

			if err != nil {
				return nil, err
			}

			if err = file.Write(d, rq.Output); err != nil {
				return nil, err
			}
			
		} else {
			if err = file.Write(body, rq.Output); err != nil {
				return nil, err
			}
		}
	}



	if rq.Unm != nil {
        if err = json.Unmarshal(body, &rq.Unm); err != nil {
            return nil, err
        }
    }
	

	return body, nil
}