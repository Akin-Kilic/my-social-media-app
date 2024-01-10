package utils

type ResponseStr struct {
	Data interface{} `json:"data"`
}

func Response(data interface{}) ResponseStr {
	return ResponseStr{
		Data: data,
	}
}
