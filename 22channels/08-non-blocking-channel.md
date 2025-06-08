### Explanation

This program uses channels to demonstrate non-blocking operations and communication between goroutines, although it doesn't actually launch any goroutines. It shows how channels can be checked for data availability without blocking execution.

1. **Channel Creation**:
   ```go
   messages := make(chan string)
   signals := make(chan bool)
   ```

   - `messages` and `signals` are two channels. `messages` is for string data, and `signals` is for boolean values.
   - Note that these channels are unbuffered, meaning they can only hold data when there's something to send and something to receive simultaneously.

2. **First Select Statement**:
   ```go
   select {
   case msg := <-messages:
       fmt.Println("received message", msg)
   default:
       fmt.Println("no message received")
   }
   ```

   - **Purpose**: Check if there's a message available to receive from the `messages` channel.
   - **Case**: `msg := <-messages` tries to receive a message from `messages`. If there were a message, it would be printed.
   - **Default**: If the case cannot proceed (no message is available), it prints "no message received".
   - **Outcome**: Since nothing is sent to `messages`, it goes to the default and prints "no message received".

3. **Second Select Statement - Sending a Message**:
   ```go
   msg := "hi"
   select {
   case messages <- msg:
       fmt.Println("sent message", msg)
   default:
       fmt.Println("no message sent")
   }
   ```

   - **Purpose**: Attempt to send a message `"hi"` to the `messages` channel.
   - **Case**: `messages <- msg`: Tries to send the message `"hi"` to the `messages` channel.
   - **Default**: If the send cannot proceed (because `messages` is unbuffered and empty), it prints "no message sent".
   - **Outcome**: The channel is unbuffered and there's no receiving operation, so it defaults with "no message sent".

4. **Third Select Statement - Receive from Either Channel**:
   ```go
   select {
   case msg := <-messages:
       fmt.Println("received message", msg)
   case sig := <-signals:
       fmt.Println("received signal", sig)
   default:
       fmt.Println("no activity")
   }
   ```

   - **Purpose**: Attempt to receive from either channel.
   - **Cases**:
     - `msg := <-messages`: Tries to receive a message from `messages`.
     - `sig := <-signals`: Tries to receive a signal from `signals`.
   - **Default**: If neither case is successful, it prints "no activity".
   - **Outcome**: Since neither channel has data, it defaults to "no activity".

### Summary

- **Non-blocking Operations**: This example highlights how to avoid blocking your program when attempting to send or receive on channels.
- **Default**: Each `select` statement has a `default` case to ensure that the program doesn't hang waiting for channel operations where data is absent.
- **Data Flow**: Illustrates conditional checks on channels before proceeding with operations, useful in situations where channel state is unpredictable.

### Real-World Use Case

This pattern is particularly useful when you're dealing with operations that might pause or delay due to unavailable data or synchronization events. Using `select` with a default case can prevent hanging, making for responsive applications.