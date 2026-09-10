# Trial Booking System

A Go/PostgreSQL service for booking trial classes: a parent registers a student, the student books a seat in a trial class, and the seat is only confirmed once payment settles. This project structure based on [Scafolding Hexagonal Architecture in Go](https://soerja.medium.com/scaffolding-hexagonal-architecture-in-go-using-giter8-15e9ff3466ed).

## Data Model

- parents
    - id
    - name
- students
    - id
    - name
    - parent_id
- trial_classes
    - id
    - name
    - quota
    - available_slots
- trial_class_members -- inserted only after a booking's payment settles successfully
    - id
    - trial_classes_id
    - student_id
- bookings
    - id
    - trial_classes_id
    - student_id
    - status -- pending, confirmed, cancelled
    - idempotency_key
    - hold_expires_at -- pending booking auto-expires 5 minutes after creation
    - payment_code
    - created_at
    - updated_at
- payment_attempts
    - id
    - booking_id
    - status -- pending, success, failed, refunded
    - note
    - created_at
    - updated_at

## API Endpoints / Server Actions

| endpoint | payload | handler |
|---|---|---|
| `POST /v1/parents` | `{ "name": "..." }` | [parent.Handler.Create](internal/adapters/primary/rest/parent/handler.go) |
| `GET /v1/parents/:id` | | [parent.Handler.GetByID](internal/adapters/primary/rest/parent/handler.go) |
| `POST /v1/students` | `{ "name": "...", "parent_id": 1 }` | [student.Handler.Create](internal/adapters/primary/rest/student/handler.go) |
| `GET /v1/students/:id` | | [student.Handler.GetByID](internal/adapters/primary/rest/student/handler.go) |
| `GET /v1/students/:id/bookings` | | [booking.Handler.GetByStudentID](internal/adapters/primary/rest/booking/handler.go) |
| `POST /v1/trial_class` | `{ "name": "..." }` | [trial_class.Handler.Create](internal/adapters/primary/rest/trial_class/handler.go) |
| `POST /v1/booking` | Header `X-Idempotency-Key`, body `{ "trial_classes_id": 1, "student_id": 1 }` | [booking.Handler.BookClass](internal/adapters/primary/rest/booking/handler.go) → [booking.Service.BookClass](internal/core/services/booking/book_class.go) |
| `POST /v1/payment/settlement` | `{ "payment_code": "..." }` | [payment.Handler.Settlement](internal/adapters/primary/rest/payment/handler.go) → [payment_attempt.Service.Settlement](internal/core/services/payment_attempt/settlement.go) |


## Booking Statuses

- `bookings.status`: `pending` → `confirmed` (payment settled) or `cancelled` (hold expired / settlement failed)
- `payment_attempts.status`: `pending` → `success` (settled) or `refunded` (rollback of a settled booking). `failed` is defined in the schema for a declined-payment callback, but no code path sets it yet — `Settlement` currently only recognizes success or hold-expiry.

## Preventing Duplicate Bookings

Two mechanisms in [book_class.go](internal/core/services/booking/book_class.go):

1. **Already-a-member check** — before creating a booking, the service loads `trial_class_members` for the class and rejects the request (`"student already register in this class"`) if the student is already a confirmed member.
2. **Idempotency key** — `POST /v1/booking` requires an `X-Idempotency-Key` header, which is stored on the booking row so a client-side retry can be recognized.
<!-- 
Note: the idempotency key is persisted but not yet enforced — there's no unique constraint on `bookings.idempotency_key` and no lookup-by-key before insert, so a retried request with the same key currently creates a second pending booking rather than returning the original. That's the main gap to close before relying on it for duplicate-submit protection. -->

## Handling Payment Failure

[`Settlement`](internal/core/services/payment_attempt/settlement.go) runs these checks, in order, when a payment code is presented:

1. If `hold_expires_at` has already passed, the booking is cancelled and the class slot is released back (`available_slots + 1`) — the caller gets `"booking hold expired, slot released"` and `"payment_attempt.status" -> refunded`.
2. If the booking isn't `pending` (e.g. already confirmed or cancelled), it's rejected with `"booking is not pending"`, so a payment code can't be settled twice.
3. Otherwise, in one DB transaction: booking → `confirmed`, payment attempt → `success`, and a `trial_class_members` row is inserted.
4. If that transaction fails for any reason, the slot is rolled back and the booking is cancelled, so a failed settlement never leaves a class under-counted.

There's currently no explicit "payment declined" callback path — only success and hold-expiry are handled. Wiring a real payment gateway would add a `failed` transition here (mark the payment attempt `failed`, release the slot, leave the booking `cancelled`).

## Two Users Competing for the Last Seat

Intended design ([architecture.md](architecture.md)) is an atomic, conditional decrement of `available_slots` (`UPDATE ... WHERE available_slots > 0`, checked via affected-rows) so only one of two concurrent requests can claim the last seat — the loser is told immediately that the class is full, and only the winner ever gets a pending booking + payment code.

**Implemented.** [`BookClass`](internal/core/services/booking/book_class.go) still does a cheap `available_slots <= 0` pre-check in Go to fail fast on the common case, but the authoritative guard is [`UpdateAvailableSlots`](internal/adapters/secondary/repository/trial_class/update_available_slots.go): `UPDATE trial_classes SET available_slots = $1 WHERE id = $2 AND available_slots > 0`. Postgres row-locks the target row for that statement, so two concurrent requests targeting the last seat serialize — the first commits `available_slots` to `0`, and the second's `WHERE` clause is re-evaluated against that committed value and matches zero rows. Zero rows affected is returned to the service as an error before any booking or payment attempt is created, so only the winner ever gets a pending booking + payment code.

Every repository write and read now goes through [`db.QuerierFromContext(ctx, r.db)`](internal/infrastructure/db/tx.go) instead of its injected `db` directly, so a call made inside a `db.Do(...)` block picks up the transaction `Do` opened. That means the slot decrement, the booking insert, and the payment-attempt insert inside `BookClass`'s `db.Do` block now all run in the same transaction and roll back together if any step fails — not just the slot decrement, which was already safe on its own as a single atomic statement.

## Checks by Layer

| Layer | Responsibility |
|---|---|
| **UI** | Basic form validation (required fields) before submit — not a source of truth, purely UX. |
| **Backend (service layer)** | Business rules: seat availability, duplicate-membership check, idempotency key handling, hold-expiry logic, transaction boundaries, slot rollback on failure. |
| **Database** | Structural integrity: foreign keys (`ON DELETE CASCADE`), enum types for `booking_status` / `payment_attempt_status`, `NOT NULL` constraints. Also owns the atomic seat-decrement via a conditional `UPDATE ... WHERE available_slots > 0` — see above. |
| **Background job** | Not implemented. Expired holds are only reclaimed lazily, the next time someone calls `Settlement` on that booking — a stale pending booking with no settlement attempt keeps its slot reserved indefinitely. A periodic sweep to cancel expired pending bookings and release slots would close this. |

### Flow Diagrams

See [architecture.md](architecture.md) for the sequence diagrams (happy path and last-seat race) this design is based on.
