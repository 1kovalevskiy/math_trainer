package substraction

import (
	"math/rand"

	"github.com/1kovalevskiy/math_trainer/internal/entity"
	"github.com/1kovalevskiy/math_trainer/internal/math/random"
)

type substractionTrainer struct {
	leftDigits uint8
	rightDigits uint8
	negativeResult bool
}

func NewSubstractionTrainer(leftDigits, rightDigits uint8, negativeResult bool) *substractionTrainer {
	return &substractionTrainer{
		leftDigits: leftDigits,
		rightDigits: rightDigits,
		negativeResult: negativeResult,
	}
}

func (st *substractionTrainer) GetExample() *entity.MathExample {
	leftNumber := random.GetRandomValue(st.leftDigits)
	rightNumber := random.GetRandomValue(st.rightDigits)
	if !st.negativeResult && rightNumber > leftNumber{
		rightNumber = rand.Int31n(leftNumber - 1) + 1
	}
	expected := leftNumber - rightNumber
	score := int(st.leftDigits) + int(st.rightDigits)
	return &entity.MathExample{
		LeftNumber: int(leftNumber),
		RightNumber: int(rightNumber),
		Operator: "-",
		Expected: int(expected),
		Score: score,
	}
}