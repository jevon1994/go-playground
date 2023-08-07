package event

type StreamPushRequest struct {
	Target     string                 `json:"url"`
	Type       string                 `json:"type"`
	Source     string                 `json:"source"`
	Extensions map[string]interface{} `json:"extensions"`
	Data       []byte                 `json:"data"`
}

//type StreamCallBackRequest struct {
//	Idc       string `json:"idc"`
//	Env       string `json:"env"`
//	Namespace string `json:"namespace"`
//}

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
