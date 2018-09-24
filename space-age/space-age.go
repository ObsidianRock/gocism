package space

//Planet in the Solar System
type Planet string

const (
	secToYears = 365.25 * 24 * 60 * 60

	Earth   = 1
	Mercury = 0.2408467
	Venus   = 0.61519726
	Mars    = 1.8808158
	Jupiter = 11.862615
	Saturn  = 29.447498
	Uranus  = 84.016846
	Neptune = 164.79132
)

//Age calculates how old someone would be on different Planets
func Age(seconds float64, p Planet) float64 {

	var years float64

	years = seconds / secToYears

	switch p {
	case Planet("Earth"):
		years /= Earth
	case Planet("Mercury"):
		years /= Mercury
	case Planet("Venus"):
		years /= Venus
	case Planet("Mars"):
		years /= Mars
	case Planet("Jupiter"):
		years /= Jupiter
	case Planet("Saturn"):
		years /= Saturn
	case Planet("Uranus"):
		years /= Uranus
	case Planet("Neptune"):
		years /= Neptune
	default:
		years *= 0
	}

	return years
}
