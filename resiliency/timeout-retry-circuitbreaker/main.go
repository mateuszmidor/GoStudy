package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/failsafe-go/failsafe-go/circuitbreaker"
	"github.com/failsafe-go/failsafe-go/failsafehttp"
	"github.com/failsafe-go/failsafe-go/retrypolicy"
	"github.com/failsafe-go/failsafe-go/timeout"
)

func main() {
	const maxAttempts = 3                   // make 3 attempts for a single request before failing the request
	const circuitBreakerThreshold = 1       // open the circuit after 1 request failure (failure counts when all request attempts are exhausted)
	const attemptTimeOut = 2 * time.Second  // per-attempt timeout (request is made of attempts)
	const requestTimeOut = 10 * time.Second // per-request timeout (sum time of all attempts)

	// 1. Circuit Breaker policy. Circuit Breaker policy is cross-requests - counts total sum of failed requests (request is considered failed after exhausting all attempts).
	breaker := circuitbreaker.NewBuilder[*http.Response]().
		HandleIf(func(resp *http.Response, err error) bool {
			requestFailed := err != nil || (resp != nil && resp.StatusCode >= 500)
			fmt.Println("checking circuit breaker policy: request failed:", requestFailed)
			return requestFailed
		}).
		WithFailureThreshold(circuitBreakerThreshold).
		WithDelay(10 * time.Second). // Time in OPEN state
		Build()

	// 2. Retry policy with Exponential Backoff. Retry policy is per-request
	retryPolicy := retrypolicy.NewBuilder[*http.Response]().
		HandleIf(func(resp *http.Response, err error) bool {
			needRetry := err != nil || (resp != nil && resp.StatusCode >= 500) // retry on any error. This should be narrowed down to transient errors only
			fmt.Println("checking retry policy: need retry:", needRetry)
			return needRetry
		}).
		WithBackoff(100*time.Millisecond, 5*time.Second). // Exponential Backoff
		WithMaxAttempts(maxAttempts).
		Build()

	// 3. Timeout policy (Hard Deadline for single attempt)
	attemptTimeout := timeout.New[*http.Response](attemptTimeOut)

	// 4. Combining policies in correct order
	// check breaker -> check retry -> check timeout
	roundTripper := failsafehttp.NewRoundTripper(nil, breaker, retryPolicy, attemptTimeout) // 3 attempt failures = 1 breaker failure (forgiving protection)
	client := &http.Client{Transport: roundTripper}                                         // note this is regular stdlib http client, easy to replace if ever needed

	// 5. Demonstrate circuit breaker behavior with 2 sequential requests;
	// first will fail because no such domain exists, second will fail because circuit breaker is open after first request failure.
	urls := []string{"https://non-existent.com", "https://www.google.com"}
	for _, url := range urls {
		fmt.Println("### Requesting", url)

		// prepare request
		reqCtx, _ := context.WithTimeout(context.Background(), requestTimeOut) // ignore request cancel for brevity
		req, _ := http.NewRequestWithContext(reqCtx, http.MethodGet, url, nil) // ignore error handling for brevity

		// send request
		resp, err := client.Do(req)

		// handle errors
		if err != nil {
			fmt.Println("Error", err)
			continue
		}

		// clean up (but this won't be reached as we trigger errors each time)
		resp.Body.Close()
		fmt.Printf("Response from %s status: %d (unexpected)\n", url, resp.StatusCode)
	}
}
