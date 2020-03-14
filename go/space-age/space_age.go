package space

type Planet string

var planets = map[Planet]float64{
	Planet("Mercury"): 0.2408467,
	Planet("Venus"): 0.61519726,
	Planet("Earth"): 1.0,
	Planet("Mars"): 1.8808158,
	Planet("Jupiter"): 11.862615,
	Planet("Saturn"): 29.447498,
	Planet("Uranus"): 84.016846,
	Planet("Neptune"): 164.79132,
}

var secondsInEarthYear float64 = 31557600

// Age takes a float representing seconds and a planet as arguments and returns a float representing age converted
// to that planets time
func Age (s float64, p Planet) float64 {
	// pAge = ageInSeconds / (converted planetary years)
	return s / (secondsInEarthYear * planets[p])
}