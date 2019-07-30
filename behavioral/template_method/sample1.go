package template_method

import "fmt"

type woodCutPrintProcessor interface {
	draw(cuttable cuttable)
	cut(cuttable cuttable)
	print(cuttable cuttable)
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
	wood := new(wood)
	wcp.processor.draw(wood)
	wcp.processor.cut(wood)
	wcp.processor.print(wood)
}

type cuttable interface {
	getName() string
}

type wood struct{}

func (*wood) getName() string {
	return "wood"
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
