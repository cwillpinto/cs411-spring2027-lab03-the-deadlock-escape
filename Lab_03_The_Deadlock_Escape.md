# Lab 3: The Deadlock Escape

**CS 411: Computing Paradigms - Spring 2027**  
**Week 3: January 25-31, 2027**  
**Due:** January 31, 2027  
**Value:** 100 points

## Purpose

Two workers acquire the same locks in opposite orders. Build the wait-for graph, connect the cycle to the Coffman conditions, and refactor both paths to preserve safety without deadlock or livelock.

## Learning objectives

- Recognize all four Coffman conditions
- Trace a wait-for cycle
- Repair lock ordering without weakening the invariant and distinguish deadlock from thread starvation using progress evidence

## Before you begin

1. Review the Week 3 slide deck and assigned VoiceThread or reading.
2. Create your own repository from the Spring 2027 template linked in Blackboard. The starter contains `transfer.go`, public `transfer_test.go`, and a README.
3. Run `go test -run '^TestTransfersSequential$' ./...` first. Then run `go test -run '^TestTransfers$' -count=50 -timeout=10s ./...`. Save both outputs before editing. The stress test may time out; keep the timeout flag so the process cannot hang indefinitely.

## Part A Diagnose (30 points)

Treat the hang as a state problem.

- In `Lab3_Analysis.md` or PDF, record held and requested locks for two opposing calls, `Transfer(A,B,1)` and `Transfer(B,A,1)`.
- Draw the wait-for graph and identify the cycle. Label each account mutex and each worker. Submit it as `wait_for_graph.pdf` or `.png`.
- Explain how each Coffman condition appears. State the balance invariant and why a sequential test does not expose the lock-order flaw.

## Part B Repair (40 points)

Use one global acquisition order.

- Change `Transfer` to acquire the two account mutexes in ascending unique `Account.ID` order, regardless of transfer direction.
- Release the mutexes in reverse order, including the insufficient-funds path.
- Keep both balance updates inside the critical section. Preserve the `Account`, `NewAccount`, `Balance`, and `Transfer` API. Do not replace the two account locks with a package-wide lock.

## Part C Stress and explain (30 points)

Show that the repair terminates and preserves value.

- Run `go test -run '^TestTransfers$' -count=50 -timeout=10s ./...` and `go test -race ./...`. Save commands and full output in `Lab3_Terminal.txt`.
- Report the starting and ending total balance of the two accounts and explain why it is preserved, including when a transfer is rejected.
- Contrast deadlock, livelock, and thread starvation in this system. For each, state what worker progress would be observable. Explain why consistent lock ordering removes the demonstrated wait-for cycle but does not, by itself, guarantee fairness.

## Verification

- `go test -run '^TestTransfers$' -count=50 -timeout=10s ./...` passes and terminates.
- `go test -race ./...` passes with no data-race report.
- Total balance remains 2,000 in the public stress test, and the rejected-transfer path changes neither account.

## What to submit

Submit your own repository URL and:

1. Your own repository URL (created from the Spring 2027 template).
2. `transfer.go` and the unchanged public `transfer_test.go` in that repository.
3. `wait_for_graph.pdf` or `wait_for_graph.png`.
4. `Lab3_Analysis.pdf` or `Lab3_Analysis.md`.
5. `Lab3_Terminal.txt`.

## Rubric

| Criterion | Points |
| --- | ---: |
| Wait-for graph | 20 |
| Coffman-condition diagnosis | 10 |
| Consistent lock-order repair | 35 |
| Safety and termination tests | 25 |
| Progress-state comparison and explanation | 10 |
| **Total** | **100** |

## AI usage and academic integrity

Add an **AI Usage** section to the required analysis file. Name each tool and list the prompts or interactions that affected the work. If you did not use AI, write `AI Usage: None.` You must be able to explain and defend every submitted design and line of code.

## Late work

Late work may be accepted for up to three calendar days with a 25 percent deduction per day, following the course syllabus.
