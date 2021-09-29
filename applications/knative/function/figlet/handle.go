package function

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/thunderboltsid/cli-tools-go/figlet"
	"io/ioutil"
	"net/http"
)

// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	res.Header().Add("Content-Type", "text/plain")

	if req.Method == http.MethodPost {
		if req.Body == nil {
			http.Error(res, "Please send a request body", 400)
			return
		}

		body, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			http.Error(res, fmt.Sprintf("unable to read request body: %v", err), 500)
			return
		}


		type bodyData struct {
			Text string `json:"text"`
		}
		bd := bodyData{}

		if err := json.Unmarshal(body, &bd); err != nil {
			http.Error(res, fmt.Sprintf("unable to unmarshal request body: %v", err), 500)
			return
		}

		if err := figlet.New(figlet.WithWriter(res), figlet.WithMsg(bd.Text)); err != nil {
			http.Error(res, fmt.Sprintf("unable to create a figlet: %v", err), 500)
			return
		}
		return
	}

	if err := figlet.New(figlet.WithWriter(res), figlet.WithMsg("ok")); err != nil {
		http.Error(res, fmt.Sprintf("unable to return OK: %v", err), 500)
		return
	}
}
