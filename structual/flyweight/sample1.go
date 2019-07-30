package flyweight

type greekLetter struct {
	Upper string
	Lower string
}

type greekLetterFactory struct {
	letterMap map[string]greekLetter
}

func (factory *greekLetterFactory) GetLetter(key string) greekLetter {
	return factory.letterMap[key]
}

var singleton = greekLetterFactory{
	letterMap: map[string]greekLetter{
		"A": {"Α", "α"},
		"V": {"Β", "β"},
		"G": {"Γ", "γ"},
		// ... more letters
	},
}

func GetFactoryInstance() *greekLetterFactory {
	return &singleton
}
