package factory_method

import "fmt"

type woodCutPrintProcessor interface {
	draw(cuttable cuttable)
	cut(cuttable cuttable)
	print(cuttable cuttable)
	createCuttable() cuttable
}

type woodCutPrint struct {
	processor woodCutPrintProcessor
}

func newWoodCutPrint(processor woodCutPrintProcessor) *woodCutPrint {
	return &woodCutPrint{
		processor: processor,
	}
}

func (wcp *woodCutPrint) process() {
	c := wcp.processor.createCuttable()
	wcp.processor.draw(c)
	wcp.processor.cut(c)
	wcp.processor.print(c)
}

type cuttable interface {
	getName() string
}

type wood struct{}

func (*wood) getName() string {
	return "wood"
}

type plastic struct{}

func (*plastic) getName() string {
	return "plastic"
}

type myWoodCutPrint struct{}

func newMyWoodCutPrint() *myWoodCutPrint {
	return new(myWoodCutPrint)
}

func (*myWoodCutPrint) draw(cuttable cuttable) {
	fmt.Println("my drawing on " + cuttable.getName())
}

func (*myWoodCutPrint) cut(cuttable cuttable) {
	fmt.Println("my cutting " + cuttable.getName())
}

func (*myWoodCutPrint) print(cuttable cuttable) {
	fmt.Println("my printing on " + cuttable.getName())
}

func (*myWoodCutPrint) createCuttable() cuttable {
	return new(wood)
}

type yourWoodCutPrint struct{}

func newYourWoodCutPrint() *yourWoodCutPrint {
	return new(yourWoodCutPrint)
}

func (*yourWoodCutPrint) draw(cuttable cuttable) {
	fmt.Println("your drawing on " + cuttable.getName())
}

func (*yourWoodCutPrint) cut(cuttable cuttable) {
	fmt.Println("your cutting " + cuttable.getName())
}

func (*yourWoodCutPrint) print(cuttable cuttable) {
	fmt.Println("your printing on " + cuttable.getName())
}

func (*yourWoodCutPrint) createCuttable() cuttable {
	return new(plastic)
}
