package template_method

import "testing"

func Test(t *testing.T) {
	myWoodCutPrint := newWoodCutPrint(newMyWoodCutPrint())
	myWoodCutPrint.process()

	yourWoodCutPrint := newWoodCutPrint(newYourWoodCutPrint())
	yourWoodCutPrint.process()
}
