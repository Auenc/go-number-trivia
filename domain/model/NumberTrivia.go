package model

//NumberTrivia is the data object containing the trivia of the number
type NumberTrivia struct{
	Number int `json:"number"`
	Text string `json:"text"`
}
