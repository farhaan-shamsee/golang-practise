# Coding Fundamentals --- Daily Revision Cheat Sheet

> **Purpose:** 5--10 minute daily revision for coding interview
> fundamentals.\
> **Current focus:** Functions, loops, Big-O, traversal, state, early
> exit, two pointers, and building results.

------------------------------------------------------------------------

## 1. Core Mental Model

For most beginner/interview problems:

**INPUT → TRAVERSE → CHECK → UPDATE → RETURN**

Typical structure:

``` text
Function
   ↓
Loop
   ↓
Condition
   ↓
Update state
   ↓
Return result
```

------------------------------------------------------------------------

## 2. Function vs Loop

### Function

A function organizes and reuses logic.

``` go
func findMax(nums []int) int {
    // ...
}
```

Think:

> **Function = WHAT the program does**

### Loop

A loop repeats work.

``` go
for _, num := range nums {
    // process num
}
```

Think:

> **Loop = HOW you process multiple elements**

**Important:** Function and loop are not alternatives. They are commonly
used together.

------------------------------------------------------------------------

## 3. Time Complexity --- Big-O

Ask:

> **How does the amount of work grow when the input gets bigger?**

Let `n = input size`.

  Complexity       Meaning             Typical example
  ---------------- ------------------- -------------------------
  **O(1)**         Constant            `nums[0]`
  **O(log n)**     Problem shrinks     Binary search
  **O(n)**         One traversal       Find max
  **O(n log n)**   Efficient sorting   Merge sort
  **O(n²)**        Nested traversal    Compare every pair
  **O(2ⁿ)**        Exponential         Some recursive problems
  **O(n!)**        Very expensive      Permutations

### For now, master:

**O(1), O(n), O(n²)**

------------------------------------------------------------------------

## 4. Big-O Rules

### Rule 1 --- Drop constants

``` text
O(2n)   → O(n)
O(10n)  → O(n)
O(100n) → O(n)
```

### Rule 2 --- Drop lower-order terms

``` text
O(n² + n)       → O(n²)
O(n² + 10n + 5) → O(n²)
```

The fastest-growing term dominates.

### Rule 3 --- Sequential loops usually add

``` go
for ... { } // O(n)

for ... { } // O(n)
```

Together:

``` text
O(n) + O(n)
= O(2n)
= O(n)
```

### Rule 4 --- Nested loops usually multiply

``` go
for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
    }
}
```

``` text
n × n = n²
Time = O(n²)
```

**But:** nested loops do not automatically mean O(n²).

If the inner loop always runs 10 times:

``` text
n × 10 = 10n = O(n)
```

------------------------------------------------------------------------

## 5. Space Complexity

Ask:

> **How much EXTRA memory does my algorithm need as input grows?**

### O(1) space

A fixed number of variables:

``` go
max := nums[0]
count := 0
sum := 0
```

Even if the input contains 1 million elements:

``` text
Extra memory ≈ constant
Space = O(1)
```

### O(n) space

A result/data structure that can grow with the input:

``` go
result := []int{}

for _, num := range nums {
    if num > 5 {
        result = append(result, num)
    }
}
```

Potentially `n` elements:

``` text
Space = O(n)
```

------------------------------------------------------------------------

## 6. Traversal

**Traversal = going through elements one by one.**

### Array / slice

``` go
for _, num := range nums {
    // process num
}
```

### String

``` go
for _, ch := range s {
    // process ch
}
```

Think:

> **"I need to look at the elements one by one."**

Typical simple traversal:

``` text
Time  = O(n)
Space = O(1)
```

when only a few variables are maintained.

------------------------------------------------------------------------

## 7. State Pattern

A huge number of easy interview problems are:

> **TRAVERSE + MAINTAIN STATE**

Template:

``` text
INITIALIZE STATE

FOR EACH ELEMENT:
    CHECK SOMETHING
    UPDATE STATE

RETURN STATE
```

Common state:

``` text
max
min
sum
count
found
result
```

### Find maximum

``` text
max = first element

for each number:
    if number > max:
        max = number

return max
```

Complexity:

``` text
Time  = O(n)
Space = O(1)
```

------------------------------------------------------------------------

## 8. Accumulator Pattern

An **accumulator** continuously builds a value.

### Sum

``` text
sum = 0

for each number:
    sum += number

return sum
```

### Count

``` text
count = 0

for each number:
    if condition:
        count++
```

Think:

> **Accumulator = keep updating an answer while traversing.**

------------------------------------------------------------------------

## 9. Search + Early Exit

Sometimes you don't need to process everything.

Example:

> Does the array contain `7`?

``` text
for each number:
    if number == 7:
        return true

return false
```

Complexity:

``` text
Best case  = O(1)
Worst case = O(n)
Space      = O(1)
```

### Key idea

> **Return immediately when you already know the answer.**

This is **early exit**.

------------------------------------------------------------------------

## 10. `break` vs `return`

### `break`

Stops the **loop**.

``` go
for ... {
    if condition {
        break
    }
}

fmt.Println("continues here")
```

### `return`

Stops the **function**.

``` go
for ... {
    if condition {
        return true
    }
}

return false
```

Remember:

> **break → exits loop**

> **return → exits function**

------------------------------------------------------------------------

## 11. Two Pointers

Useful for:

-   strings
-   arrays
-   comparing opposite ends
-   reversing
-   palindrome problems

Basic setup:

``` text
left  → beginning
right → end
```

Then:

``` text
while left < right:

    process left/right

    left++
    right--
```

Mental model:

> **Two ends → process → move inward**

------------------------------------------------------------------------

## 12. Palindrome Pattern

Question:

> Is `"racecar"` a palindrome?

Mental model:

``` text
r a c e c a r
↑           ↑
L           R
```

Compare:

``` text
s[left] == s[right]
```

Then:

``` text
left++
right--
```

If mismatch:

``` text
return false
```

If pointers meet/cross without mismatch:

``` text
return true
```

Complexity:

``` text
Time  = O(n)
Space = O(1)*
```

`*` assumes the implementation does not require a growing copy of the
string.

### Important

Prefer:

``` text
while left < right
```

When the pointers meet or cross, the comparison work is finished.

------------------------------------------------------------------------

## 13. Reverse Pattern

Same two-pointer idea.

``` text
left = 0
right = n - 1
```

While:

``` text
left < right
```

Swap:

``` text
left ↔ right
```

Then:

``` text
left++
right--
```

Mental model:

> **Two ends → swap → move inward**

Conceptually:

``` text
Time = O(n)
```

------------------------------------------------------------------------

## 14. Building a Result

Example:

``` text
Input:
[1, 4, 7, 8, 10]

Output:
[4, 8, 10]
```

Mental model:

``` text
result = empty

for each number:
    if condition:
        add to result

return result
```

This is:

> **Traversal + Condition + Result**

Because the result can grow with `n`:

``` text
Time  = O(n)
Space = O(n)
```

------------------------------------------------------------------------

## 15. State vs Growing Result

  Problem           State                                  Extra Space
  ----------------- ------------------------ -------------------------
  Find max          `max`                                     **O(1)**
  Find min          `min`                                     **O(1)**
  Sum               `sum`                                     **O(1)**
  Count             `count`                                   **O(1)**
  Search            early return / `found`                    **O(1)**
  Palindrome        `left`, `right`                         **O(1)**\*
  Reverse           `left`, `right`            **O(1)** conceptually\*
  Filter elements   `result[]`                                **O(n)**

`*` Exact space for string manipulation can depend on the
language/implementation.

------------------------------------------------------------------------

# 16. Pattern Recognition Cheat Sheet

  When the problem says...                 Think...
  ---------------------------------------- -----------------------------
  Find maximum                             Traversal + State
  Find minimum                             Traversal + State
  Calculate sum                            Traversal + Accumulator
  Count occurrences                        Traversal + Counter
  Does it exist?                           Traversal + Early Exit
  Find first matching item                 Traversal + Early Exit
  Reverse                                  Two Pointers
  Palindrome                               Two Pointers + Comparison
  Return elements satisfying a condition   Traversal + Result Array
  Compare every element with every other   Nested Loops → likely O(n²)
  Input gets divided in half               Think O(log n)

------------------------------------------------------------------------

# 17. 30-Second Interview Checklist

**Do NOT immediately start typing.**

Ask:

``` text
1. What is the INPUT?

2. What is the OUTPUT?

3. Do I need to TRAVERSE?

4. Do I need a CONDITION?

5. What STATE do I need?
   max / min / count / sum / found / result

6. Can I EXIT EARLY?

7. Do I need TWO POINTERS?

8. Am I building a RESULT?

9. What is the TIME COMPLEXITY?

10. What is the SPACE COMPLEXITY?
```

Then say:

> **"My approach is..."**

Explain the logic briefly **before coding**.

------------------------------------------------------------------------

# 18. Interview Explanation Template

When the interviewer gives you a problem:

### Step 1 --- Clarify

``` text
"What is the expected input?"
"What should I return?"
"Are there any constraints?"
```

### Step 2 --- Explain approach

``` text
"My approach is to traverse the input,
check the condition, maintain the required state,
and return the result."
```

### Step 3 --- Code

Only after the logic is clear.

### Step 4 --- Test

Use:

``` text
Normal case
Edge case
Empty input
Single element
```

### Step 5 --- Complexity

Always finish with:

``` text
Time:  O(...)
Space: O(...)
```

------------------------------------------------------------------------

# 19. The Three Patterns to Remember First

If you forget everything else, remember these:

### Pattern A --- Traverse + State

``` text
for each element:
    check
    update state

return state
```

Examples:

**max, min, sum, count**

------------------------------------------------------------------------

### Pattern B --- Traverse + Early Exit

``` text
for each element:
    if found:
        return answer

return default
```

Examples:

**search, first match, existence check**

------------------------------------------------------------------------

### Pattern C --- Two Pointers

``` text
left  = start
right = end

while left < right:
    compare/swap
    move inward
```

Examples:

**reverse, palindrome**

------------------------------------------------------------------------

# 20. Golden Rules ⭐

``` text
1. Don't code immediately.
   Understand the problem first.

2. Think in patterns, not individual solutions.

3. Function organizes logic.
   Loop repeats logic.

4. One full traversal → usually O(n).

5. Nested n × n traversal → O(n²).

6. Constants disappear in Big-O.
   O(10n) → O(n).

7. Fixed extra variables → O(1) space.

8. Growing result/map/list → usually O(n) space.

9. If you know the answer, return early.

10. Two ends?
    Think two pointers.

11. Always know:
    Time complexity + Space complexity.

12. Explain the algorithm before writing code.
```

------------------------------------------------------------------------

# 🧠 Daily 5-Minute Revision

Before coding practice, mentally answer:

``` text
What is traversal?
→ Go through elements one by one.

What is state?
→ Information maintained while traversing.

What is an accumulator?
→ A value continuously updated while processing.

What is early exit?
→ Return as soon as the answer is known.

What is O(n)?
→ Work grows linearly with input.

What is O(n²)?
→ Work grows roughly as n × n.

What is O(1) space?
→ Fixed extra memory.

When do I use two pointers?
→ Often when processing from both ends.

What is the basic interview flow?
→ Understand → Explain → Code → Test → Complexity.
```

------------------------------------------------------------------------

## Current Learning Roadmap --- Updated

Completed:

-   [x] Functions vs loops
-   [x] Time complexity
-   [x] Space complexity
-   [x] Big-O basics
-   [x] Traversal
-   [x] State / accumulator
-   [x] Early exit
-   [x] `break` vs `return`
-   [x] Two pointers
-   [x] Palindrome
-   [x] Reverse
-   [x] Building a result
-   [x] Hash Maps / frequency counting
-   [x] Two Sum
-   [x] Sliding window
-   [x] Stack
-   [x] Parentheses validation
-   [x] Monotonic stack
-   [x] Duplicate handling with index-based stack
-   [x] Amortized complexity
-   [x] Binary search

Next only if needed for the interview:

-   [ ] Queue
-   [ ] Sorting/common array patterns
-   [ ] Recursion
-   [ ] Mixed interview problems
-   [ ] Go-specific coding patterns
-   [ ] Concurrency patterns

------------------------------------------------------------------------

# 21. Hash Maps / Frequency Counting

A **map is memory that lets you answer a question quickly**.

Common patterns:

``` text
number → frequency
number → seen
number → index
key    → useful information
```

### Frequency map

``` go
freq := make(map[int]int)

for _, val := range nums {
    freq[val]++
}
```

Think:

> "I need to know how many times each value appeared."

Typical complexity:

``` text
Time  = O(n)
Space = O(n)
```

### Seen map

``` go
seen := make(map[int]bool)

for _, val := range nums {
    if seen[val] {
        return val
    }
    seen[val] = true
}
```

### Two Sum pattern

For the classic index-returning version:

``` go
seen := make(map[int]int)

for i, val := range nums {
    needed := target - val

    if j, ok := seen[needed]; ok {
        return []int{j, i}
    }

    seen[val] = i
}
```

Important:

> **Check before storing** so the same element is not used twice.

Complexity:

``` text
Time  = O(n)
Space = O(n)
```

------------------------------------------------------------------------

# 22. Sliding Window

Use sliding window when a problem asks about a **contiguous
window/subarray/substring**.

For a fixed-size window:

``` text
First window → calculate it

Move window:
newSum = oldSum - leaving + entering
```

Typical complexity:

``` text
Time  = O(n)
Space = O(1)
```

Edge cases:

``` text
window size <= 0
window size > array length
empty input
```

------------------------------------------------------------------------

# 23. Stack

Stack = **LIFO** --- Last In, First Out.

Go slice implementation:

``` go
stack := []int{}

stack = append(stack, 10)       // push
top := stack[len(stack)-1]      // peek
stack = stack[:len(stack)-1]    // pop
```

### Valid parentheses

``` text
Opening bracket → PUSH

Closing bracket:
    if stack empty → invalid
    compare with top
    pop if matching

End:
    stack empty → valid
```

Typical complexity:

``` text
Time  = O(n)
Space = O(n)
```

------------------------------------------------------------------------

# 24. Monotonic Stack

Useful for:

> "Find the next greater element to the right."

The stack stores elements that are **waiting for an answer**.

For:

``` text
[2, 1, 5, 3, 4]
```

Result:

``` text
[5, 5, -1, 4, -1]
```

Prefer storing **indices** when duplicates are possible.

Key idea:

``` text
Each index is pushed once.
Each index is popped at most once.
```

Therefore the nested-looking loops are still:

``` text
Time  = O(n) amortized
Space = O(n)
```

Interview lesson:

> **A nested loop does not automatically mean O(n²). Count the total
> work across all iterations.**

------------------------------------------------------------------------

# 25. Binary Search

Use binary search when the search space is sorted and you can eliminate
half the possibilities.

Core variables:

``` text
left
right
mid
```

Standard implementation:

``` go
func binarySearch(arr []int, target int) int {
    left := 0
    right := len(arr) - 1

    for left <= right {
        mid := left + (right-left)/2

        if arr[mid] == target {
            return mid
        }

        if target > arr[mid] {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return -1
}
```

Mental model:

``` text
target > arr[mid]
→ discard LEFT side including mid
→ left = mid + 1

target < arr[mid]
→ discard RIGHT side including mid
→ right = mid - 1
```

Important:

``` text
while left <= right
```

because when `left == right`, one element still needs checking.

Complexity:

``` text
Time  = O(log n)
Space = O(1)
```

------------------------------------------------------------------------

# 26. Pattern Recognition Cheat Sheet

  Problem clue                    Think
  ------------------------------- ------------------------
  Find max/min/sum/count          Traversal + State
  Does it exist?                  Traversal + Early Exit
  Reverse / palindrome            Two Pointers
  Need frequency                  Hash Map
  Need fast "have I seen this?"   Seen Map
  Need complement/lookup          Hash Map
  Contiguous window/subarray      Sliding Window
  Matching nested brackets        Stack
  Next greater/smaller element    Monotonic Stack
  Input/search space is sorted    Binary Search
  Compare every pair              Usually O(n²)
  Repeatedly eliminate half       O(log n)

------------------------------------------------------------------------

# 27. Current Interview Mental Model

Before coding:

``` text
1. Understand INPUT
2. Understand OUTPUT
3. Identify the PATTERN
4. Explain the ALGORITHM
5. Consider EDGE CASES
6. Code
7. Test with an example
8. State TIME + SPACE complexity
```

High-value question:

> **"What information do I need to remember while I move through the
> input?"**

Possible answers:

``` text
Nothing extra        → simple traversal
A running value      → state / accumulator
Previously seen data → hash map
Both ends            → two pointers
A contiguous region  → sliding window
Unresolved items     → stack
Monotonic candidates → monotonic stack
A shrinking search   → binary search
```

------------------------------------------------------------------------

# 28. Golden Rules ⭐

``` text
1. Don't code immediately.
   Understand the problem first.

2. Think in patterns, not individual solutions.

3. Function organizes logic.
   Loop repeats logic.

4. One full traversal → usually O(n).

5. Nested n × n traversal → O(n²).

6. Constants disappear in Big-O.
   O(10n) → O(n).

7. Fixed extra variables → O(1) space.

8. Growing result/map/list → usually O(n) space.

9. If you know the answer, return early.

10. Two ends?
    Think two pointers.

11. Contiguous region?
    Think sliding window.

12. Need counts or fast lookup?
    Think hash map.

13. Unresolved "next greater/smaller"?
    Think monotonic stack.

14. Sorted search space?
    Think binary search.

15. Always know:
    Time complexity + Space complexity.

16. Explain the algorithm before writing code.
```

------------------------------------------------------------------------

# 🧠 Daily 5-Minute Revision

Before coding practice, mentally answer:

\`\`\`text What is traversal? → Go through elements one by one.

What is state? → Information maintained while traversing.

What is an accumulator? → A value continuously updated while processing.

What is early exit? → Return as soon as the answer is known.

What is O(n)? → Work grows linearly with input.

What is O(n²)? → Work grows roughly as n × n.

What is O(1) space? → Fixed extra memory.

When do I use two pointers? → Often when processing from both ends.

When do I use a hash map? → When I need counts, seen values, indices, or
fast lookup.

When do I use sliding window? → For a contiguous region with a moving
boundary.

When do I use a stack? → When the most recent unresolved item matters.

When do I use a monotonic stack? → For next greater/smaller style
problems.

When do I use binary search? → When the search space is sorted and can
be halved.

What is the interview flow? → Understand → Explain → Code → Test →
Complexity.
