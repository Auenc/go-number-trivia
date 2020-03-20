package usercases

import (
	"errors"
	"github.com/auenc/go-number-trivia/domain/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

//mockNumberTriviaRepository is the fake repository to test the interface
type mockNumberTriviaRepository struct{
}
//Specific is the mocked method of NumberTriviaRepository.Specific
func (r *mockNumberTriviaRepository) Specific(number int) (*model.NumberTrivia, error){
	if number < 0{
		return nil, errors.New(ErrorNumberIsNegative)
	}
	return &model.NumberTrivia{
		Number: number,
		Text:   "some test trivia",
	}, nil
}

func (r *mockNumberTriviaRepository) Random() (*model.NumberTrivia, error){
	return &model.NumberTrivia{
		Number: 1,
		Text:  "some random test trivia",
	}, nil
}

func TestNumberTriviaRepository_Specific(t *testing.T){
	t.Run("positive number", func(t *testing.T) {
		repo := new(mockNumberTriviaRepository)

		expected := &model.NumberTrivia{
			Number: 1,
			Text:   "some test trivia",
		}

		result, err := repo.Specific(1)
		assert.Nil(t, err)
		assert.Equal(t, expected,result)
	})
	t.Run("negative number", func(t *testing.T) {
		repo := new(mockNumberTriviaRepository)
		expected := errors.New(ErrorNumberIsNegative)
		shouldBeNil, result := repo.Specific(-1)
		assert.Nil(t, shouldBeNil)
		assert.NotNil(t, result)
		assert.Equal(t, expected, result)
	})
}

func TestNumberTriviaRepository_Random(t *testing.T){
	t.Run("random number", func(t *testing.T) {
		repo := new(mockNumberTriviaRepository)

		expected := &model.NumberTrivia{
			Number: 1,
			Text:   "some random test trivia",
		}

		result, err := repo.Random()
		assert.NotNil(t,result)
		assert.Equal(t, expected, result)
		assert.Nil(t, err)
	})
}

