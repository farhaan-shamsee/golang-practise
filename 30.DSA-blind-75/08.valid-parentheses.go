/*
Problem: Valid Parentheses
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

Real-world use case for Platform Engineers:
- Validating JSON, YAML, or nested configuration files
- Checking syntax in Kubernetes manifests
- Validating Terraform HCL configurations
- Parsing nested API responses or DSL expressions
- Validating bracket matching in log queries

Approach: Stack-based Solution

Algorithm:
1. Create an empty stack
2. Iterate through each character:
   - If opening bracket: push to stack
   - If closing bracket:
     a. If stack is empty → return false (no matching open)
     b. Pop from stack
     c. Check if popped bracket matches current closing bracket
     d. If not matching → return false
3. After loop: return true if stack is empty, false otherwise

Time Complexity: O(n) - single pass through string
Space Complexity: O(n) - stack can hold up to n/2 opening brackets

Example 1: s = "()[]{}"
  i=0: '(' → push '('  → stack = ['(']
  i=1: ')' → pop '(', matches ✓ → stack = []
  i=2: '[' → push '['  → stack = ['[']
  i=3: ']' → pop '[', matches ✓ → stack = []
  i=4: '{' → push '{'  → stack = ['{']
  i=5: '}' → pop '{', matches ✓ → stack = []
  Stack empty → return true

Example 2: s = "([)]"
  i=0: '(' → push '('  → stack = ['(']
  i=1: '[' → push '['  → stack = ['(', '[']
  i=2: ')' → pop '[', expected ']' but got ')' ✗ → return false

Example 3: s = "((("
  i=0: '(' → push '('  → stack = ['(']
  i=1: '(' → push '('  → stack = ['(', '(']
  i=2: '(' → push '('  → stack = ['(', '(', '(']
  Stack not empty → return false (unmatched opening brackets)

Key Insight:
The stack ensures that the most recent opening bracket is matched first (LIFO).
This naturally handles nested structures like {[()]} where innermost must close first.
*/
package main

import "fmt"

func isValid(s string) bool {
	// Map of closing to opening brackets
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []rune{}

	for _, char := range s {
		// Check if it's a closing bracket
		if opening, isClosing := pairs[char]; isClosing {
			// Stack empty or top doesn't match
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			// Pop from stack
			stack = stack[:len(stack)-1]
		} else {
			// It's an opening bracket, push to stack
			stack = append(stack, char)
		}
	}

	// Valid only if all brackets are matched (stack empty)
	return len(stack) == 0
}

func mainIsValid() {
	test1 := "()"
	fmt.Printf("Input: \"%s\" → Output: %v\n", test1, isValid(test1)) // true

	test2 := "()[]{}"
	fmt.Printf("Input: \"%s\" → Output: %v\n", test2, isValid(test2)) // true

	test3 := "(]"
	fmt.Printf("Input: \"%s\" → Output: %v\n", test3, isValid(test3)) // false

	test4 := "([)]"
	fmt.Printf("Input: \"%s\" → Output: %v\n", test4, isValid(test4)) // false

	test5 := "{[]}"
	fmt.Printf("Input: \"%s\" → Output: %v\n", test5, isValid(test5)) // true

	// Real-world example: Validate JSON-like structure
	jsonLike := `{"key": [1, 2, {"nested": true}]}`
	fmt.Printf("\nJSON-like: \"%s\"\n", jsonLike)
	fmt.Printf("Brackets valid: %v\n", isValid(jsonLike)) // true
}
