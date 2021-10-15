package forms

type ImageInput struct {
	FilePath string `json:"filePath" form:"filePath"'`
	Pid      string `json:"pid" form:"pid"`
}
