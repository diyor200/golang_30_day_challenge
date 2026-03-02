Golang Pro Dev — Day 1 Challenge
 Title: Concurrent Log Analyzer
 Problem Statement
You are building a backend utility that analyzes application logs.
You are given a slice of log entries (around 1,000 lines).
Each log entry is a string in the following format:

[DATE] [LEVEL] message

Your Task
Build a concurrent log processor that:
Processes logs using goroutines
Extracts the log level (INFO, WARN, ERROR)
Counts occurrences of each level
Returns a summary map

Expected output:
```map[ERROR:1 INFO:2 WARN:1]```
Functional Requirements
• Implement the function:
Use concurrency (goroutines required)
• Avoid race conditions
• Do not use global variables
• Program must not panic
 Edge Cases to Handle
• Malformed log lines
Example:
• Missing log level
• Empty slice input
Your program should skip invalid logs gracefully.
 What We’re Evaluating
• Proper goroutine usage
• Understanding of synchronization
• Safe shared memory handling
• Clean and readable code
• Structured thinking
 Bonus (Optional)
• Make worker count configurable
• Compare sequential vs concurrent performance
• Add context cancellation support
 Goal of Day 1:
Not just “use goroutines” —
but use them safely and intentionally.