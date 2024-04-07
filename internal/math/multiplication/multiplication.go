package multiple

import (
	"github.com/1kovalevskiy/math_trainer/internal/entity"
	"github.com/1kovalevskiy/math_trainer/internal/math/random"
)

type multiplicationTrainer struct {
	leftDigits uint8
	rightDigits uint8
}

func NewMultiplicationTrainer(leftDigits, rightDigits uint8) *multiplicationTrainer {
	return &multiplicationTrainer{
		leftDigits: leftDigits,
		rightDigits: rightDigits,
	}
}


func (mt *multiplicationTrainer) GetExample() *entity.MathExample {
	leftNumber := random.GetRandomValue(mt.leftDigits)
	rightNumber := random.GetRandomValue(mt.rightDigits)
	expected := leftNumber * rightNumber
	score := (int(mt.leftDigits) + int(mt.rightDigits)) * 5
	return &entity.MathExample{
		LeftNumber: int(leftNumber),
		RightNumber: int(rightNumber),
		Operator: "*",
		Expected: int(expected),
		Score: score,
	}
}