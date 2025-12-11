package test_create

import (
	"log/slog"
	"net/http"
	req "project-go/internal/http-server/dto/request"
	res "project-go/internal/http-server/dto/response"
	"project-go/internal/lib/api/response"
	"project-go/internal/models"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type TestCreate interface {
	CreateTest(test *models.Test) (*models.Test, error)
}

type QuestionAdd interface {
	CreateQuestion(question *models.TestQuestion) (*models.TestQuestion, error)
}

type Response struct {
	response.Response
	Test res.TestResponse
}

func New(log *slog.Logger, TestCreate TestCreate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.test.test.New"
		log = log.With(
			slog.String("op", op),
			slog.String("req_id", middleware.GetReqID(r.Context())),
		)
		var req req.TestRequest
		err := render.DecodeJSON(r.Body, &req)

		if err != nil {
			log.Error("failed to decode req", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("failed to decode req"))
			return
		}
		if req.Title == "" {
			log.Error("text field is empty")
			render.JSON(w, r, response.Error("text field is empty"))
			return
		}
		testModel := mapTestRequestToModel(req)
		createdTest, err := TestCreate.CreateTest(testModel)
		if err != nil {
			log.Error("failed to create test", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("failed to create test"))
			return
		}

		var questionsResp []res.QuestionResponse
		for _, q := range createdTest.Questions {
			var optionsResp []res.OptionResponse
			for _, o := range q.Options {
				optionsResp = append(optionsResp, res.OptionResponse{
					OptionText: o.OptionText,
					IsCorrect:  o.IsCorrect,
				})
			}

			// 2. Маппим сам вопрос
			questionsResp = append(questionsResp, res.QuestionResponse{
				Question: q.Question,
				Options:  optionsResp,
			})
		}

		render.JSON(w, r, Response{
			Response: response.OK(),
			Test: res.TestResponse{
				ID:          createdTest.ID,
				Title:       createdTest.Title,
				Description: createdTest.Description,
				CreatedAt:   createdTest.CreatedAt,
				Questions:   questionsResp,
			},
		})

	}
}

// маппим TestRequest в models.Test
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
