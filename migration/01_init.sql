CREATE TYPE booking_status AS ENUM ('pending', 'confirmed', 'cancelled');
CREATE TYPE payment_attempt_status AS ENUM ('pending', 'success', 'failed', 'refunded');

CREATE TABLE parents (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE students (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    parent_id INT NOT NULL REFERENCES parents(id) ON DELETE CASCADE
);

CREATE TABLE trial_classes (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    quota INT NOT NULL DEFAULT 0,
    available_slots INT NOT NULL DEFAULT 0
);

CREATE TABLE trial_class_members (
    id BIGSERIAL PRIMARY KEY,
    trial_classes_id INT NOT NULL REFERENCES trial_classes(id) ON DELETE CASCADE,
    student_id INT NOT NULL REFERENCES students(id) ON DELETE CASCADE
);

CREATE TABLE bookings (
    id BIGSERIAL PRIMARY KEY,
    trial_classes_id INT NOT NULL REFERENCES trial_classes(id) ON DELETE CASCADE,
    student_id INT NOT NULL REFERENCES students(id) ON DELETE CASCADE,
    status booking_status NOT NULL,
    idempotency_key VARCHAR(255),
    payment_code VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE payment_attempts (
    id BIGSERIAL PRIMARY KEY,
    booking_id INT NOT NULL REFERENCES bookings(id) ON DELETE CASCADE,
    status payment_attempt_status NOT NULL,
    note TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
