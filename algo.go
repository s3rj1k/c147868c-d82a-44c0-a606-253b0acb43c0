package main

import "fmt"

// List of algoritm types.
const (
	BaseTypeAlgo    = "base"
	Custom1TypeAlgo = "custom1"
	Custom2TypeAlgo = "custom2"
)

// In desribes input data.
type In struct {
	A bool    `json:"A"`
	B bool    `json:"B"`
	C bool    `json:"C"`
	D float64 `json:"D"`
	E int     `json:"E"`
	F int     `json:"F"`
}

// Out describes output data.
type Out struct {
	H string  `json:"H"`
	K float64 `json:"K"`
}

// IsValid checks if Out object is valid.
func (out Out) IsValid() bool {
	return out.H != ""
}

// GetH returns H dependent on algoritm type.
func (in In) GetH(algo string) (H string) {
	if in.A && in.B && !in.C {
		H = "M"
	}

	if in.A && in.B && in.C {
		H = "P"
	}

	if !in.A && in.B && in.C {
		H = "T"
	}

	/*
		if algo == Custom1TypeAlgo {
			// no action
		}
	*/

	if algo == Custom2TypeAlgo {
		if in.A && in.B && !in.C {
			H = "T"
		}

		if in.A && !in.B && in.C {
			H = "M"
		}
	}

	return H
}

// GetK returns K dependent on algoritm type and H.
func (in In) GetK(algo, H string) (K float64) {
	if H == "M" {
		K = in.D + (in.D * float64(in.E) / 10)
	}

	if H == "P" {
		K = in.D + (in.D * float64(in.E-in.F) / 25.5)
	}

	if H == "T" {
		K = in.D - (in.D * float64(in.F) / 30)
	}

	if algo == Custom1TypeAlgo {
		if H == "P" {
			K = 2.0*in.D + (in.D * float64(in.E) / 100)
		}
	}

	if algo == Custom2TypeAlgo {
		if H == "M" {
			K = float64(in.F) + in.D + (in.D * float64(in.E) / 100)
		}
	}

	return K
}

// Process returns Out object based on algoritm type.
func (in In) Process(algo string) (out Out, err error) {
	out.H = in.GetH(algo)
	out.K = in.GetK(algo, out.H)

	if !out.IsValid() {
		return Out{}, fmt.Errorf("invalid input data for %s algoritm", algo)
	}

	return out, nil
}
