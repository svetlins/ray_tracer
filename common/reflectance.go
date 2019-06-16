package common

import (
	"math"
)

func Reflectance(precomputed *Precomputed) float64 {
	cos := precomputed.EyeVector.Dot(precomputed.NormalVector)

	if precomputed.RefractiveIndexEnter > precomputed.RefractiveIndexExit {
		nRatio := precomputed.RefractiveIndexEnter / precomputed.RefractiveIndexExit
		sin2T := nRatio*nRatio * (1 - cos*cos)

		if sin2T > 1 {
			return 1.0
		}

		cos = math.Sqrt(1 - sin2T)
	}

	r0 := math.Pow((precomputed.RefractiveIndexEnter - precomputed.RefractiveIndexExit) / (precomputed.RefractiveIndexEnter + precomputed.RefractiveIndexExit), 2)

	return r0 + (1 - r0) * math.Pow(1 - cos, 5)
}
