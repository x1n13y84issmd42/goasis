package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/x1n13y84issmd42/goasis/src/api"
	"github.com/x1n13y84issmd42/goasis/src/log"
	"github.com/x1n13y84issmd42/goasis/src/utility"
)

// JSONResponseObject --
type JSONResponseObject struct {
	Log         log.ILogger
	APIResponse *api.Response
}

// JSONMap is a map to unmarshal JSONs into.
type JSONMap = map[string]interface{}

// Test tests.
func (tResp JSONResponseObject) Test(response *http.Response) bool {
	//TODO: HTTPReponse{}.Test(response)
	responseBody, _ := ioutil.ReadAll(response.Body)
	mJSON := make(JSONMap)
	errJSON := json.Unmarshal(responseBody, &mJSON)

	if errJSON != nil {
		tResp.Log.ResponseExpectedObject(tResp.APIResponse)
		tResp.Log.Error(errJSON)
		return false
	}

	ctx := &utility.Context{
		Path: []string{"Response"},
	}

	return Schema{tResp.APIResponse.Schema, tResp.Log}.Test(mJSON, ctx)
}