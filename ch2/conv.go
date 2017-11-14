package ch2

import (
	"flag"
	"fmt"
	"strconv"
	"log"
	"io"
	"os"
)

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius {
	return Celsius(k) + AbsoluteZeroC
}

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}

// FtToM ...
func FtToM(f Foot) Metre {
	return Metre(0.3048 * f)
}

// MToFt ...
func MToFt(m Metre) Foot {
	return Foot(3.28083 * m)
}

func PToKg(p Pound) Kilogram {
	return Kilogram(0.453592 * p)
}

func KgToP(kg Kilogram) Pound {
	return Pound(2.2046226 * kg)
}

func unitConverter(t, l, w bool, v float64) {
	if t {
		c := Celsius(v)
		f := Fahrenheit(v)
		k := Kelvin(v)
		fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s, %s = %s, %s = %s\n",
			c, CToF(c), c, CToK(c), f, FToC(f), f, FToK(f), k, KToC(k), k, KToF(k))
	}
	if l {
		ft := Foot(v)
		m := Metre(v)
		fmt.Printf("%s = %s, %s = %s\n",
			ft, FtToM(ft), m, MToFt(m))
	}
	if w {
		p := Pound(v)
		kg := Kilogram(v)
		fmt.Printf("%s = %s, %s = %s\n",
			p, PToKg(p), kg, KgToP(kg))
	}
}

// UnitConverterMain ...
func UnitConverterMain() {
	t := flag.Bool("t", false, "temperature convert")
	l := flag.Bool("l", false, "length convert")
	w := flag.Bool("w", false, "weight convert")

	flag.Parse()

	input := flag.Args()
	if len(input) != 0 {
		for _, s := range input {
			v, err := strconv.ParseFloat(s, 64)
			if err != nil {
				log.Println(s, ":", err)
			} else {
				unitConverter(*t, *l, *w, v)
			}
			fmt.Println()
		}
	} else {
		var v float64
		for {
			fmt.Print(">> ")
			_, err := fmt.Scanf("%f", &v)
			if err == io.EOF {
				os.Exit(0)
			} else if err != nil {
				panic(err)
			}
			unitConverter(*t, *l, *w, v)
			fmt.Println()
		}
	}
}
