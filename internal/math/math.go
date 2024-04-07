package math

import (
	"math/rand"

	"github.com/1kovalevskiy/math_trainer/config"
	"github.com/1kovalevskiy/math_trainer/internal/entity"
	"github.com/1kovalevskiy/math_trainer/internal/math/addition"
	"github.com/1kovalevskiy/math_trainer/internal/math/division"
	"github.com/1kovalevskiy/math_trainer/internal/math/multiplication"
	"github.com/1kovalevskiy/math_trainer/internal/math/substraction"

)

type (
	trainer interface {
		GetExample() *entity.MathExample
	}

	examplesGiver struct {
		allowTrainer 			[]trainer
	}
)

func NewTrainer(cfg *config.Config) (*examplesGiver, error) {
	allowTrainer := []trainer{}
	if cfg.Addition.Enable {
		additionTrainer := addition.NewAdditionTrainer(
			cfg.Addition.LeftDigits, cfg.Addition.RightDigits,
		)
		allowTrainer = append(allowTrainer, additionTrainer)
	}
	if cfg.Subtraction.Enable {
		substractionTrainer := substraction.NewSubstractionTrainer(
			cfg.Subtraction.LeftDigits, cfg.Subtraction.RightDigits, cfg.Subtraction.NegativeResult,
		)
		allowTrainer = append(allowTrainer, substractionTrainer)
	}
	if cfg.Multiplication.Enable {
		multiplicationTrainer := multiple.NewMultiplicationTrainer(
			cfg.Multiplication.LeftDigits, cfg.Multiplication.RightDigits,
		)
		allowTrainer = append(allowTrainer, multiplicationTrainer)
	}
	if cfg.Division.Enable {
		divisionTrainer, err := division.NewDivisionTrainer(
			cfg.Division.DividendDigits, cfg.Division.DivisorDigits,
		)
		if err != nil {
			return nil, err
		}
		allowTrainer = append(allowTrainer, divisionTrainer)
	}

	result := &examplesGiver{
		allowTrainer: allowTrainer,
	}
	return result, nil
}

func (e *examplesGiver) GetExample() *entity.MathExample {
	randomIndex := rand.Intn(len(e.allowTrainer))
	pick := e.allowTrainer[randomIndex]
	return pick.GetExample()
}