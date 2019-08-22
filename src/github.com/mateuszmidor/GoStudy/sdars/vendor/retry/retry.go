package retry

import (
	"fmt"
	"reflect"
	"time"
)

// UntilSuccessOr5Failures tries to call fn(params) until fn returns no error OR until limit of 5 tries is reached
// The function returns same results as fn, but they need casting eg. err := result[0].Interface().(error)
// Note: last result of fn should carry error value
func UntilSuccessOr5Failures(failMessage string, fn interface{}, params ...interface{}) []reflect.Value {
	const retryLimit = 5
	const retryDelay = 1 * time.Second

	// prepare reflection
	f := reflect.ValueOf(fn)
	if f.Type().NumIn() != len(params) {
		panic(fmt.Sprintf("RetryUntilSuccessOr5Failures: %s: incorrect number of parameters provided vs required!", failMessage))
	}

	inputs := make([]reflect.Value, len(params))
	for k, in := range params {
		inputs[k] = reflect.ValueOf(in)
	}

	// repeat until success OR retry limit reached
	var result []reflect.Value
	for nRetry := 1; nRetry <= retryLimit; nRetry++ {
		result = f.Call(inputs)
		if len(result) == 0 {
			panic("RetryUntilSuccessOr5Failures: called function returned no error value")
		}
		errorValue := result[len(result)-1] // watch out for nil when casting// last result of fn call should carry error value

		if errorValue.IsNil() {
			if nRetry > 1 {
				fmt.Printf("RetryUntilSuccessOr5Failures: %s succeeded in attempt no. %d!\n", failMessage, nRetry)
			}
			return result
		}

		if nRetry < retryLimit {
			fmt.Printf("RetryUntilSuccessOr5Failures: %s failed (%d/%d), retrying in 1 second...\n", failMessage, nRetry, retryLimit)
			time.Sleep(retryDelay)
		} else {
			err := errorValue.Interface().(error)
			fmt.Printf("RetryUntilSuccessOr5Failures: %s failed (%d/%d), giving up: %s\n", failMessage, nRetry, retryLimit, err.Error())
		}
	}

	return result
}
