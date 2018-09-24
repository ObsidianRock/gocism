package space

//Planet in the Solar System
type Planet string

const (
	secToYears = 365.25 * 24 * 60 * 60
)

var planets = map[Planet]float64{
	Planet("Earth"):   1,
	Planet("Mercury"): 0.2408467,
	Planet("Venus"):   0.61519726,
	Planet("Mars"):    1.8808158,
	Planet("Jupiter"): 11.862615,
	Planet("Saturn"):  29.447498,
	Planet("Uranus"):  84.016846,
	Planet("Neptune"): 164.79132,
}

//Age calculates how old someone would be on different Planets
func Age(seconds float64, p Planet) float64 {

	var years float64

	years = seconds / secToYears

	convertion, ok := planets[p]
	if !ok {
		return 0
	}

	years /= convertion

	return years

}
