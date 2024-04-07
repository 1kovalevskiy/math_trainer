package division

import (
	"fmt"

	"github.com/1kovalevskiy/math_trainer/internal/entity"
	"github.com/1kovalevskiy/math_trainer/internal/math/random"
)

type divisionTrainer struct {
	dividendDigits uint8
	divisorDigits uint8
}

func NewDivisionTrainer(dividendDigits, divisorDigits uint8) (*divisionTrainer, error) {
	if dividendDigits <= divisorDigits {
		return nil, fmt.Errorf("the dividend must be greater than the divisor")
	}

	result := &divisionTrainer{
		dividendDigits: dividendDigits,
		divisorDigits: divisorDigits,
	}

	return result, nil
}

func (dt *divisionTrainer) GetExample() *entity.MathExample {
	resultDigits := dt.dividendDigits - dt.divisorDigits
	result := random.GetRandomValue(resultDigits)
	divisorNumber := random.GetRandomValue(dt.divisorDigits)
	dividentNumber := result * divisorNumber
	expected := result
	score := (int(dt.dividendDigits) + int(dt.divisorDigits)) * 5
	example := &entity.MathExample{
		LeftNumber: int(dividentNumber),
		RightNumber: int(divisorNumber),
		Operator: "/",
		Expected: int(expected),
		Score: score,
	}
	return example
}