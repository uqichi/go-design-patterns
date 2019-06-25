package builder

import (
	"fmt"
	"strconv"
)

type muscleBuilder interface {
	train()
	showTrainingResult()
}

type armMuscleBuilder struct {
	setCount          int
	totalTrainedCount int
}

func newArmMuscleBuilder(setCount int) *armMuscleBuilder {
	return &armMuscleBuilder{
		setCount: setCount,
	}
}

func (b *armMuscleBuilder) train() {
	b.totalTrainedCount += b.setCount
	fmt.Println("trained arm " + strconv.Itoa(b.setCount) + " times")
}

func (b *armMuscleBuilder) showTrainingResult() {
	fmt.Println("trained arm " + strconv.Itoa(b.totalTrainedCount) + " times in total! GJ")
}

type chestMuscleBuilder struct {
	setCount          int
	totalTrainedCount int
}

func newChestMuscleBuilder(setCount int) *chestMuscleBuilder {
	return &chestMuscleBuilder{
		setCount: setCount,
	}
}

func (b *chestMuscleBuilder) train() {
	b.totalTrainedCount += b.setCount
	fmt.Println("trained chest " + strconv.Itoa(b.setCount) + " times")
}

func (b *chestMuscleBuilder) showTrainingResult() {
	fmt.Println("trained chest " + strconv.Itoa(b.totalTrainedCount) + " times in total! GJ")
}

type director struct {
	builder muscleBuilder
}

func newDirector(builder muscleBuilder) *director {
	return &director{
		builder: builder,
	}
}

func (dir *director) followTrainingSetEveryday() {
	// do 3 set training
	dir.builder.train()
	dir.builder.train()
	dir.builder.train()
}
