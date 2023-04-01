# Goroutine vs Thread

## Schedule management
Goroutines are executed by the go-runtime which is a cooperative scheduler, not the OS. Threads are executed by the kernel which is a preemptive scheduler so application has no control over the execution.

## Communication
Goroutines communicate using channels. Threads have no established way of communication.

## Speed
Goroutines are faster than Thread.

## Infrastructure dependency
Goroutines are just functions executed by the runtime, But threads are dependent upon the underlying infrastructure such as OS.

## Growable stack
Goroutines has growable stack memory model. It uses 2kb of memory at first then it can grow gradually. Contrarily threads need at least 1MB to run.