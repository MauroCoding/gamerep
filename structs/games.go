package structs

type Game struct {
	name, studio string
	year, price  float64
}

func (g *Game) getInfo() string {
	g.name += " original"
	return g.name + " " + g.studio
}
