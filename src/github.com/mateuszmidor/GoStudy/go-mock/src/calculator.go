//go:generate mockgen -source=$GOFILE -destination=$PWD/mocks/${GOFILE} -package=mocks
package calculator

type Calculator interface {
	Add(a, b int) int
	Mul(a, b int) int
}
