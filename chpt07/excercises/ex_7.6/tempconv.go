package tempconv

/*
Пример создания своего флага, путем реализации интерфейса
 Value представляет собой интерфейс значения, хранящегося в флаге,
 type Value interface {
	String() string
	Set(string) error
}
*/
import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }
func KToC(k Kelvin) Celsius     { return Celsius(k - 273.15) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + 273.15) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// *celsiusFlag соответствует интерфейсу flag.Value
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("Неверная температура %q", s)
}

// CelsiusFlag определяет флаг Celsius с указанным именем, значением по умолчанию и
// строкой-инструкцией по применению и возвращает адрес переменной-флага. Аргумент флага
// должен содержать числовое значение и единицу изменения, например "100С".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
