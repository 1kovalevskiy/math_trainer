package addition

import (
	"github.com/1kovalevskiy/math_trainer/internal/entity"
	"github.com/1kovalevskiy/math_trainer/internal/math/random"
)

type additionTrainer struct {
	leftDigits uint8
	rightDigits uint8
}

func NewAdditionTrainer(leftDigits, rightDigits uint8) *additionTrainer {
	return &additionTrainer{
		leftDigits: leftDigits,
		rightDigits: rightDigits,
	}
}

func (at *additionTrainer) GetExample() *entity.MathExample {
	leftNumber := random.GetRandomValue(at.leftDigits)
	rightNumber := random.GetRandomValue(at.rightDigits)
	expected := leftNumber + rightNumber
	score := int(at.leftDigits) + int(at.rightDigits)
	return &entity.MathExample{
		LeftNumber: int(leftNumber),
		RightNumber: int(rightNumber),
		Operator: "+",
		Expected: int(expected),
		Score: score,
	}
}