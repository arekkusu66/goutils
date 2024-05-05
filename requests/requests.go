package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"regexp"

	"github.com/arekkusu66/goutils/fsys"
);


type Request struct {
	Unm			any
	Output		string
};


//POST request
func (rq *Request) POST(url string, body any) ([]byte, error) {

    json_data, err := json.Marshal(body);

    if err != nil {
		return nil, err;
    };


    response, err := http.Post(url , "application/json" , bytes.NewBuffer(json_data));

    if err != nil {
		return nil, err;
    };

	defer response.Body.Close();


	response_data, err := io.ReadAll(response.Body);

	if err != nil {
		return nil, err;
	};


	if rq.Output != "" {
		var file = &fsys.FileSys{};

		if regexp.MustCompile(`\.json$`).MatchString(rq.Output) {
			var mapUnm = make(map[string]interface{});

			if err = json.Unmarshal(response_data, &mapUnm); err != nil {
				return nil, err;
			};

			d, err := json.MarshalIndent(mapUnm, "", "  ");

			if err != nil {
				return nil, err;
			};

			if err = file.Write(d, rq.Output); err != nil {
				return nil, err;
			};

		} else {
			if err = file.Write(response_data, rq.Output); err != nil {
				return nil, err;
			};
		};
	};


	if rq.Unm != nil {
        if err = json.Unmarshal(response_data, &rq.Unm); err != nil {
            return nil, err;
        };
    };


	return response_data, nil;
};

//GET request
func (rq *Request) GET(url string) ([]byte, error) {

	response, err := http.Get(url);

	if err != nil {
		return nil, err;
	};

	defer response.Body.Close();
	

	body, err := io.ReadAll(response.Body);

	if err != nil {
		return nil, err;
	};
	


	if rq.Output != "" {
		var file = &fsys.FileSys{};

		if regexp.MustCompile(`\.json$`).MatchString(rq.Output) {
			var mapUnm = make(map[string]interface{});

			err = json.Unmarshal(body, &mapUnm);

			if err != nil {
				return nil, err;
			};

			d, err := json.MarshalIndent(mapUnm, "", "  ");

			if err != nil {
				return nil, err;
			};

			if err = file.Write(d, rq.Output); err != nil {
				return nil, err;
			};
			
		} else {
			if err = file.Write(body, rq.Output); err != nil {
				return nil, err;
			};
		};
	};



	if rq.Unm != nil {
        if err = json.Unmarshal(body, &rq.Unm); err != nil {
            return nil, err;
        };
    };
	

	return body, nil;
};