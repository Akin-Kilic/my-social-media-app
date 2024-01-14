package dtos

type AddPostDTO struct {
	Caption string `json:"caption" form:"caption"`
	// ShortLink string `json:"-"`
	// File string `json:"file" form:"file"`
	// Images    []string `json:"-"`
}

type UpdatePostDTO struct {
	Id      string `json:"id"`
	Caption string `json:"caption"`
}
