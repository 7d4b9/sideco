package backend

import "go.mongodb.org/mongo-driver/mongo"

type EvalRequest struct {
	Applicants []struct {
		SiderID string `json:"taskId"`
	} `json:"taskId"`
}

type EvalResponse struct {
	TaskID     string `json:"taskId"`
	Applicants []struct {
		SiderID   string `json:"siderId"`
		Score     string `json:"score"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"applicants"`
	Description string   `json:"description"`
	Country     string   `json:"country"`
	Tags        []string `json:"tags"`
}

type Handler struct {
	*mongo.Database
}

func (s *Handler) Eval(req *EvalRequest) (resp *EvalResponse, err error) {
	return &EvalResponse{}, nil
}
