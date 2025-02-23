package user

import (
	"context"

	validator "github.com/diogoazevedoo/go-bid/internal/validator"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func (req CreateUserRequest) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.Username), "username", "username cannot be empty")
	eval.CheckField(validator.NotBlank(req.Email), "email", "email cannot be empty")
	eval.CheckField(validator.NotBlank(req.Bio), "bio", "bio cannot be empty")

	eval.CheckField(
		validator.Matches(req.Email, validator.EmailRX),
		"email", "must be a valid email",
	)

	eval.CheckField(
		validator.MinChars(req.Bio, 10) && validator.MaxChars(req.Bio, 255),
		"bio", "bio must have a length between 10 and 255",
	)

	eval.CheckField(
		validator.MinChars(req.Password, 8),
		"password", "password must be bigger than 8 chars",
	)

	return eval
}
