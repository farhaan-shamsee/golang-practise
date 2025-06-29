# Golang

## Memory Management

- Memory management in Go happens automatically.
- ![image](https://github.com/user-attachments/assets/c6c97999-46a5-42d4-b90d-81a9f24afea7)
- Key functions:
  - `new()`: Allocates memory but does not initialize it.
  - `make()`: Allocates and initializes slices, maps, and channels.
- Differences between `new` and `make`:
  - `new` is used for value types, while `make` is for built-in reference types.
- Garbage collection is automatic in Go.
- ![image](https://github.com/user-attachments/assets/bb56daec-9be2-4ad1-86ea-6d7c860c8fe1)
- ![image](https://github.com/user-attachments/assets/1d810a23-ba92-44ba-8096-0d38e1c12bd6)
- The `runtime` package provides tools for memory management, such as `runtime.NumCPU`.

---

## GO MOD

- ![image](https://github.com/user-attachments/assets/f31dfd3a-16f4-4937-bba2-9cc6ab063409)
- **Indirect Dependencies**: The term "indirect" indicates that the package is not directly used in the code but is required by other dependencies.
- **go.sum**: Contains the hash of the modules for verification.
- **Modules Storage**: Modules are stored in `go-home/bin/package/mod/module_name/cache`.
- Useful commands:
  - `go list -m all`: Lists all dependent packages.
  - `go list -m -versions github.com/gorilla/mux`: Lists available versions of a module.
  - `go mod tidy`: Cleans up unused dependencies.
  - `go mod why github.com/gorilla/mux`: Explains why a module is needed.
  - `go mod graph`: Displays a dependency graph.
  - `go mod vendor`: Copies dependencies to the `vendor` directory.

---

## Concurrency and goroutine

- ![image](https://github.com/user-attachments/assets/7f652cc6-0681-4d17-ad15-36e701c446cf)
- Eating, using Instagram and switching on AC analogy
- Goroutines is the way to achieve parallelism
- ![image](https://github.com/user-attachments/assets/05b02c1b-adfd-4e8f-9483-e1efa6e8b539)
- ![image](https://github.com/user-attachments/assets/a0f63816-e5cf-4086-96bc-b986f6273b7b)
- ![image](https://github.com/user-attachments/assets/25969b08-4f3e-4df8-aec4-a5b3c3569292)
- Unbuffered channel is usually done for synchronous communication, because here the sender will wait for the response from the receiver then only close.
- ![image](https://github.com/user-attachments/assets/26128744-0ed2-41ca-80e9-21c2c528803f)
- In buffered channel the sending goroutine can just send the data and continue with its work and not be blocked. Although if the channel is full then the sending GR will be blocked.
- Hence the communication between sending GR and receiving GR is asynchronous.

---

## MUTEX

- Mutex is a mutual exclusion lock.
- It locks a memory used by a go routine and does not allow other go routine to interfere.

---

## CHANNELS

- Channels are used to communicate between goroutines.
- They are like pipes that connect concurrent goroutines.
- In Go, an "infinite channel" pattern is often used to continuously process work or events until some external condition stops it (such as a timeout, signal, or program exit). The channel itself may not be infinite, but the goroutine runs indefinitely, often waiting for work or performing repeated actions.
- Real Use Cases:
  - Background Workers
    - Continuously process jobs from a channel (e.g., handling web requests, processing tasks from a queue).

  - Event Listeners
    - Listen for events (like file changes, network messages) and react as long as the program runs.

  - Heartbeat/Health Checks
    - Periodically send heartbeat signals to monitor system health.

  - Polling Services
    - Regularly poll an external service or database for updates.
