package design_search_autocomplete

type AutocompleteSystem struct{}

func Constructor(sentences []string, times []int) AutocompleteSystem { return AutocompleteSystem{} }
func (ac *AutocompleteSystem) Input(c byte) []string                 { return nil }
