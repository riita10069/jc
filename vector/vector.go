package vector

type Vector struct {
	Value []float64
}

func NewVector(value []float64) Vector {
	return Vector{Value: value}
}

func (v *Vector) HammingDistance(other Vector) float64 {
	var humming_distance float64 = 0
	for i, value := range v.Value {
		if value != other.Value[i] {
			humming_distance += 1
		}
	}
	return humming_distance
}

func (v *Vector) GetLength() int {
	return len(v.Value)
}
