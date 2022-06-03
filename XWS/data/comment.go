package data

type Comment struct {
	Author  string `json:"username,omitempty"`
	Content string `json:"text,omitempty"`
}

type Comments []*Comment

func getCommentByCompany(id string) (Comments, error) {

	return nil, nil
}
