package funfacts

// Sett inn alle Funfucts i en struktur
type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

// Implementer funfacts-funksjon: GetFunFacts
func (ff *FunFacts) GetFunFacts(about string) []string {
	var facts []string
	switch about {
	case "sun":
		facts = ff.Sun
	case "luna":
		facts = ff.Luna
	case "terra":
		facts = ff.Terra
	default:
		facts = []string{"Sorry, I don't have any fun facts about that."}
	}
	return facts
}
