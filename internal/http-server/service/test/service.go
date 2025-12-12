package testservice

import (
	req "project-go/internal/http-server/dto/request"
	"project-go/internal/models"
)

type TestCreate interface {
	CreateTest(test *models.Test) (*models.Test, error)
}

type QuestionAdd interface {
	CreateQuestion(question *models.TestQuestion) (*models.TestQuestion, error)
}

type Service struct {
	TestRepo     TestCreate
	QuestionRepo QuestionAdd
}

func New(TestRepo TestCreate, QuestionRepo QuestionAdd) *Service {
	return &Service{
		TestRepo:     TestRepo,
		QuestionRepo: QuestionRepo,
	}
}

func (s *Service) TestCreate(test req.TestRequest) (*models.Test, error) {
	testModel := mapTestRequestToModel(test)
	createdTest, err := s.TestRepo.CreateTest(testModel)
	if err != nil {
		return nil, err
	}

	return createdTest, nil
}
func mapTestRequestToModel(req req.TestRequest) *models.Test {
	test := &models.Test{
		Title:       req.Title,
		Description: req.Description,
	}

	for _, q := range req.Questions {
		question := models.TestQuestion{
			Question: q.Question,
		}
		for _, o := range q.Options {
			option := models.TestOption{
				OptionText: o.OptionText,
				IsCorrect:  o.IsCorrect,
			}
			question.Options = append(question.Options, option)
		}
		test.Questions = append(test.Questions, question)
	}

	return test
}
