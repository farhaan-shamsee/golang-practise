# Go Revision Notes — One-Liners

## Context & Timeouts
- Default to `context.WithTimeout`, not bare `context.Background()` — always ask "what if this never responds?"
- `context.WithTimeout` = application-level deadline for the whole operation (caller's intent).
- `http.Client{Timeout: ...}` = transport-level safety net for the HTTP call itself.
- Skip a timeout only when: it's a long-running/human-approval wait, it's already bounded upstream, or it's a local/no-network operation — and say why out loud.
- Always `defer cancel()` right after `context.WithTimeout(...)` to avoid leaking resources.
- If you already set context timeout on the request, you often don't need a separate `http.Client.Timeout` — pick one as the primary control, understand the other as a fallback.

## fmt.Println vs fmt.Errorf vs %w vs %v
- Ask: "Am I *displaying* this, or *returning* this?"
- `fmt.Println` / `fmt.Printf` → for humans, right now, on screen (debugging/logging/output).
- `fmt.Errorf` → building an error value to return/propagate up the call stack.
- `%w` inside `fmt.Errorf` → wraps the original error, preserves identity for `errors.Is` / `errors.As`.
- `%v` inside `fmt.Errorf` → flattens error to text only, identity is lost.
- Rule: returning + caller may need to unwrap the cause → use `%w`. Returning + just need context text → use `%v`.

## General Function-Writing Pattern (main + helper function)
- Typical caller flow: **set up inputs → call function → check error → use result.**
- `ctx := context.Background()` (or `WithTimeout`) → `url := "..."` → `data, err := fn(ctx, url)` → `if err != nil { ... return }` → use `data`.
- Always check `err` immediately after the call that produces it.

## Pointers & Structs
- Use `&http.Client{...}` (pointer) to avoid copying the whole struct and to allow modifying it in place.

## Format Verb Bugs (%d vs %w vs %v)
- Before writing a verb, ask: "What type is the variable I'm inserting? Does the verb match that type?"
- `%d` → integers only (e.g. `resp.StatusCode`). Passing an `error` to `%d` is a bug (e.g. `fmt.Errorf("unexpected status: %d", err)` is wrong — should be `resp.StatusCode`).
- `%s` → strings. `%v` → generic/any value. `%w` → wraps an `error` specifically (only valid with `Errorf`).

## json.NewDecoder(reader).Decode(&target)
- Two chained calls: `NewDecoder(resp.Body)` builds a streaming JSON reader over the response body (no need to load the whole body into memory first); `.Decode(&target)` reads the stream and writes the result into `target`.
- Needs `&target` (pointer) because `Decode` must write into your variable — passing `target` without `&` would only fill a copy, leaving your original untouched.
- Alternative non-streaming form: `io.ReadAll` + `json.Unmarshal(body, &target)` — needs the whole body upfront; `NewDecoder(...).Decode(...)` is more memory-efficient and the common idiom for HTTP bodies.

## if-with-initializer: `if err := ...; err != nil { }`
- Packs two statements into one line: run the call, capture its result, then immediately test it — separated by `;`.
- Equivalent to writing `err := call()` then `if err != nil {}` as two separate lines.
- Key benefit: the variable (`err`) is scoped only to the `if`/`else` block, so it won't leak into or collide with an outer `err` variable already used earlier in the function.
- Trigger phrase: "call something that returns an error right where I check it, and don't let that variable outlive the check."
