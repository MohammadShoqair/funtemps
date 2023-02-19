package conv

func FarhenheitToCelsius(fahrenheit float64) float64 {

	celsius := (fahrenheit - 32) * 5 / 9

	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	return celsius
}

func FarhenheitToKevlin(Farhenheit float64) float64 {

	Kelvin := (Farhenheit-32)*5/9 + 273.15

	return Kelvin
}

// //cccccccccccccccc
func CelsiusToFarhenheit(Celsius float64) float64 {

	farhrenheit := Celsius*9/5 + 32

	return farhrenheit
}

func CelsiusToKelvin(Celsius float64) float64 {
	Kelvin := Celsius + 273.15

	return Kelvin
}

// kkkkkkkkkkkkkkkkkkkkkkkkkkkkk
func KelvinToCelsius(Kelvin float64) float64 {
	Celsius := Kelvin - 273.15

	return Celsius
}

func KelvinToFarhenheit(Kelvin float64) float64 {
	Farhrenheit := (Kelvin-273.15)*9/5 + 32

	return Farhrenheit
}

func GetSunTemp() string {
	return "Temperatur på ytre lag av Solen 5506.85°C.\nTemperatur i Solens kjerne er 15 000 000°C."
}
