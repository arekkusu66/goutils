package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"net/http"
	"regexp"

	"github.com/arekkusu66/goutils/fsys"
)


type Request struct {
	Unm					any
	Output				string
	FilePerms			fs.FileMode
}


//Makes a POST request, if ContentType isn't specified, it defaults to application/json
func (rq *Request) POST(url string, contentType string, body any) ([]byte, error) {

    jsonData, err := json.Marshal(body)

    if err != nil {
		return nil, err
    }

    response, err := http.Post(url, contentType, bytes.NewBuffer(jsonData))

    if err != nil {
		return nil, err
    }

	defer response.Body.Close()


	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}


	if rq.Output != "" && rq.FilePerms != 0 {
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

			if err = file.Write(d, rq.Output, rq.FilePerms); err != nil {
				return nil, err
			}

		} else {
			if err = file.Write(responseData, rq.Output, rq.FilePerms); err != nil {
				return nil, err
			}
		}
	} else {
		return nil, errors.New("One of the Output or FilePerms fields are empty")
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
	

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}
	


	if rq.Output != "" && rq.FilePerms != 0 {
		var file = &fsys.FileSys{}

		if regexp.MustCompile(`\.json$`).MatchString(rq.Output) {
			var mapUnm = make(map[string]interface{})

			err = json.Unmarshal(responseData, &mapUnm)

			if err != nil {
				return nil, err
			}

			d, err := json.MarshalIndent(mapUnm, "", "  ")

			if err != nil {
				return nil, err
			}

			if err = file.Write(d, rq.Output, rq.FilePerms); err != nil {
				return nil, err
			}
			
		} else {
			if err = file.Write(responseData, rq.Output, rq.FilePerms); err != nil {
				return nil, err
			}
		}
	} else {
		return nil, errors.New("One of the Output or FilePerms fields are empty")
	}



	if rq.Unm != nil {
        if err = json.Unmarshal(responseData, &rq.Unm); err != nil {
            return nil, err
        }
    }
	

	return responseData, nil
}