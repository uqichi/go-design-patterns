package builder

import "testing"

func Test(t *testing.T) {
	armBuilder := newArmMuscleBuilder(20)
	armDirector := newDirector(armBuilder)
	armDirector.followTrainingSetEveryday()
	armBuilder.showTrainingResult()

	chestBuilder := newChestMuscleBuilder(30)
	chestDirector := newDirector(chestBuilder)
	chestDirector.followTrainingSetEveryday()
	chestBuilder.showTrainingResult()
}
