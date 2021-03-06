package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/weishi258/http-log-collector/log"
	"github.com/weishi258/http-log-collector/rest"
	"github.com/weishi258/http-log-collector/rest/error_code"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetLogRoutes() []rest.Route {
	return []rest.Route{

		{
			"logPost",
			strings.ToUpper("post"),
			"/log",
			LogPost,
			nil,
		},
	}

}

func LogPost(w http.ResponseWriter, r *http.Request) ([]byte, int, error_code.ResponseError) {
	logger := log.GetLogger()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("read from body failed", zap.String("error", err.Error()))
		return nil, http.StatusBadRequest, error_code.NewResponseError(error_code.InvalidParameter)
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "\t")
	if err != nil {
		logger.Error("unmarshal log json failed", zap.String("error", err.Error()))
		return nil, http.StatusInternalServerError, error_code.NewResponseError(error_code.InternalError)
	}
	logger.Info(fmt.Sprintf("%s", string(prettyJSON.Bytes())))

	return nil, http.StatusOK, nil
}
