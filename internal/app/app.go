package app

import (
	"fmt"
	"log/slog"

	"github.com/1kovalevskiy/math_trainer/config"
	"github.com/1kovalevskiy/math_trainer/internal/math"
)

func Run(cfg *config.Config) {
	fmt.Println(*cfg)
	trainer, err := math.NewTrainer(cfg)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	for i := 0; i < 10; i++ {
		example := trainer.GetExample()
		fmt.Println(example)
	} 
}
