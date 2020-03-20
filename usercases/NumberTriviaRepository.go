package usercases

import "github.com/auenc/go-number-trivia/domain/model"

//NumberTriviaRepository is the interface that defines the two ways of obtaining NumberTrivia
type NumberTriviaRepository interface{
	Specific(int) (model.NumberTrivia, error)
	Random() (model.NumberTrivia, error)
}
