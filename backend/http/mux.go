package http

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"gitlab.com/bbrod/sideco/backend"
)

func NewHandler(backendHandler *backend.Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", evalHandle(backendHandler))
	return mux
}

func evalHandle(backendHandler *backend.Handler) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		evalReq := &backend.EvalRequest{}
		err := decoder.Decode(evalReq)
		if err != nil {
			log.WithError(err).Error("failure body decode")
			return
		}
		evalResp, err := backendHandler.Eval(evalReq)
		if err != nil {
			log.WithError(err).Error("failure eval input request")
			return
		}
		if err := json.NewEncoder(resp).Encode(evalResp); err != nil {
			log.WithError(err).Error("failure encode response")
		}
	}
}
