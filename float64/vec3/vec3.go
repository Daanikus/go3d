// Package vec3 contains a 3D float64 vector type T and functions.
package vec3

import (
	"fmt"
	"math"

	"github.com/ungerik/go3d/float64/generic"
)

var (
	// Zero holds a zero vector.
	Zero = T{}

	// UnitX holds a vector with X to one.
	UnitX = T{1, 0, 0}
	// UnitY holds a vector with Y to one.
	UnitY = T{0, 1, 0}
	// UnitZ holds a vector with Z to one.
	UnitZ = T{0, 0, 1}

	// Red holds the color red.
	Red = T{1, 0, 0}
	// Green holds the color green.
	Green = T{0, 1, 0}
	// Blue holds the color black.
	Blue = T{0, 0, 1}
	// Black holds the color black.
	Black = T{0, 0, 0}
	// White holds the color white.
	White = T{1, 1, 1}

	// MinVal holds a vector with the smallest possible component values.
	MinVal = T{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64}
	// MaxVal holds a vector with the highest possible component values.
	MaxVal = T{+math.MaxFloat64, +math.MaxFloat64, +math.MaxFloat64}
)

// T represents a 3D vector.
type T [3]float64

// From copies a T from a generic.T implementation.
func From(other generic.T) T {
	switch other.Size() {
	case 2:
		return T{other.Get(0, 0), other.Get(0, 1), 0}
	case 3, 4:
		return T{other.Get(0, 0), other.Get(0, 1), other.Get(0, 2)}
	default:
		panic("Unsupported type")
	}
}

// Parse parses T from a string. See also String()
func Parse(s string) (r T, err error) {
	_, err = fmt.Sscanf(s, "%f %f %f", &r[0], &r[1], &r[2])
	return r, err
}

// String formats T as string. See also Parse().
func (vec *T) String() string {
	return fmt.Sprintf("%f %f %f", vec[0], vec[1], vec[2])
}

// Rows returns the number of rows of the vector.
func (vec *T) Rows() int {
	return 3
}

// Cols returns the number of columns of the vector.
func (vec *T) Cols() int {
	return 1
}

// Size returns the number elements of the vector.
func (vec *T) Size() int {
	return 3
}

// Slice returns the elements of the vector as slice.
func (vec *T) Slice() []float64 {
	return []float64{vec[0], vec[1], vec[2]}
}

// Get returns one element of the vector.
func (vec *T) Get(col, row int) float64 {
	return vec[row]
}

// IsZero checks if all elements of the vector are zero.
func (vec *T) IsZero() bool {
	return vec[0] == 0 && vec[1] == 0 && vec[2] == 0
}

// Length returns the length of the vector.
// See also LengthSqr and Normalize.
func (vec *T) Length() float64 {
	return float64(math.Sqrt(vec.LengthSqr()))
}

// LengthSqr returns the squared length of the vector.
// See also Length and Normalize.
func (vec *T) LengthSqr() float64 {
	return vec[0]*vec[0] + vec[1]*vec[1] + vec[2]*vec[2]
}

// Scale multiplies all element of the vector by f and returns vec.
func (vec *T) Scale(f float64) *T {
	vec[0] *= f
	vec[1] *= f
	vec[2] *= f
	return vec
}

// Scaled returns a copy of vec with all elements multiplies by f.
func (vec *T) Scaled(f float64) T {
	return T{vec[0] * f, vec[1] * f, vec[2] * f}
}

// Invert inverts the vector.
func (vec *T) Invert() *T {
	vec[0] = -vec[0]
	vec[1] = -vec[1]
	vec[2] = -vec[2]
	return vec
}

// Inverted returns an inverted copy of the vector.
func (vec *T) Inverted() T {
	return T{-vec[0], -vec[1], -vec[2]}
}

// Normalize normalizes the vector to unit length.
func (vec *T) Normalize() *T {
	sl := vec.LengthSqr()
	if sl == 0 || sl == 1 {
		return vec
	}
	vec.Scale(1 / math.Sqrt(sl))
	return vec
}

// Normalized returns a unit length normalized copy of the vector.
func (vec *T) Normalized() T {
	v := *vec
	v.Normalize()
	return v
}

// Normal returns an orthogonal vector.
func (vec *T) Normal() T {
	n := Cross(vec, &UnitZ)
	if n.IsZero() {
		return UnitX
	}
	return n.Normalized()
}

// Add adds another vector to vec.
func (vec *T) Add(v *T) *T {
	vec[0] += v[0]
	vec[1] += v[1]
	vec[2] += v[2]
	return vec
}

// Sub subtracts another vector from vec.
func (vec *T) Sub(v *T) *T {
	vec[0] -= v[0]
	vec[1] -= v[1]
	vec[2] -= v[2]
	return vec
}

// Mul multiplies the components of the vector with the respective components of v.
func (vec *T) Mul(v *T) *T {
	vec[0] *= v[0]
	vec[1] *= v[1]
	vec[2] *= v[2]
	return vec
}

// Add returns the sum of two vectors.
func Add(a, b *T) T {
	return T{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

// Sub returns the difference of two vectors.
func Sub(a, b *T) T {
	return T{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

// Mul returns the component wise product of two vectors.
func Mul(a, b *T) T {
	return T{a[0] * b[0], a[1] * b[1], a[2] * b[2]}
}

// Dot returns the dot product of two vectors.
func Dot(a, b *T) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

// Cross returns the cross product of two vectors.
func Cross(a, b *T) T {
	return T{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[0],
	}
}

// Angle returns the angle between two vectors.
func Angle(a, b *T) float64 {
	return math.Acos(Dot(a, b))
}

// Min returns the component wise minimum of two vectors.
func Min(a, b *T) T {
	min := *a
	if b[0] < min[0] {
		min[0] = b[0]
	}
	if b[1] < min[1] {
		min[1] = b[1]
	}
	if b[2] < min[2] {
		min[2] = b[2]
	}
	return min
}

// Max returns the component wise maximum of two vectors.
func Max(a, b *T) T {
	max := *a
	if b[0] > max[0] {
		max[0] = b[0]
	}
	if b[1] > max[1] {
		max[1] = b[1]
	}
	if b[2] > max[2] {
		max[2] = b[2]
	}
	return max
}
