package calculator_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mateuszmidor/GoStudy/go-mock/src/mocks"
)

func TestAdd(t *testing.T) {
	// given
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	calc := mocks.NewMockCalculator(mockCtrl)

	// when
	calc.EXPECT().Add(1, 4).Return(5).Times(1)

	// then
	if calc.Add(1, 4) != 5 {
		t.Error("1 + 4 should be 5")
	}
}

func TestMul(t *testing.T) {
	// given
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	calc := mocks.NewMockCalculator(mockCtrl)
	const a, b, result = 4, 5, 20

	// when
	calc.EXPECT().Mul(gomock.Eq(a), gomock.Eq(b)).DoAndReturn(
		// signature of anonymous function must have the same number of input and output arguments as the mocked method.
		func(int, int) int {
			return result
		},
	).MaxTimes(1)

	// then
	if calc.Mul(a, b) != result {
		t.Errorf("%d + %d should be %d", a, b, result)
	}
}
