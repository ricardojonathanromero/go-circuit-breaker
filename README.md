# go-circuit-breaker
An example about Circuit Breaker Design Patter Using Golang

## What is the Circuit Breaker?
The Circuit Breaker pattern is a design pattern used in software development to improve the
resilience and fault tolerance of distributed systems, particularly in situations where
services or components may fail or become unreliable.

When one service or component becomes slow or unreliable, it can lead to a chain reaction of
failures as other services that depend on it also experience delays or failures. The Circuit
Breaker pattern helps prevent these cascading failures by temporarily isolating the unreliable
service when it exceeds a failure threshold.

![alt image](https://miro.medium.com/v2/resize:fit:1034/format:webp/1*q1W_VWUUxpetzM-5djExhg.png)

1. **Closed State**: In the closed state, the circuit breaker allows requests to pass through
   to the underlying service or component as usual. It monitors the responses to these requests
   for failures, such as timeouts, exceptions, or other errors.

2. **Open State**: When the circuit breaker detects that the failure rate (e.g., the percentage
    of failed requests) exceeds a certain threshold, it transitions to the open state. In this
    state, the circuit breaker prevents any requests from reaching the unreliable service,
    effectively “breaking” the circuit to that service. Instead, it may return predefined fallback 
    responses or take other actions to mitigate the failure.

3. **Half-Open State**: After a predefined time or under specific conditions, the circuit breaker
    may transition to a half-open state. In this state, the circuit breaker allows a limited
    number of test requests to pass through to the service. If these test requests succeed, it 
    indicates that the service has recovered, and the circuit breaker transitions back to the
    closed state. If the test requests continue to fail, the circuit breaker remains in the open state.
