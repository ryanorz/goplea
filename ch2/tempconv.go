package ch2

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }

type Foot float64
type Metre float64

func (f Foot) String() string { return fmt.Sprintf("%gft", f) }
func (m Metre) String() string { return fmt.Sprintf("%gm", m) }

type Pound float64
type Kilogram float64

func (p Pound) String() string { return fmt.Sprintf("%g£", p) }
func (kg Kilogram) String() string { return fmt.Sprintf("%gKg", kg) }
