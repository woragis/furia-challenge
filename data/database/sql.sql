-- PostgreSQL

-- 1. Users
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role TEXT CHECK (role IN ('student', 'admin', 'teacher')) NOT NULL
);

-- 2. Students
CREATE TABLE students (
    student_id SERIAL PRIMARY KEY,
    user_id INTEGER UNIQUE REFERENCES users(user_id) ON DELETE CASCADE,
    date_of_birth DATE,
    grade TEXT,
    location TEXT,
    enrollment_date DATE DEFAULT CURRENT_DATE
);

-- 3. Courses
CREATE TABLE courses (
    course_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    subject TEXT NOT NULL,
    teacher_id INTEGER REFERENCES users(user_id) ON DELETE SET NULL,
    start_date DATE,
    end_date DATE
);

-- 4. Lessons
CREATE TABLE lessons (
    lesson_id SERIAL PRIMARY KEY,
    course_id INTEGER REFERENCES courses(course_id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    content TEXT,
    lesson_date DATE
);

-- 5. Quizzes
CREATE TABLE quizzes (
    quiz_id SERIAL PRIMARY KEY,
    course_id INTEGER REFERENCES courses(course_id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    instructions TEXT,
    date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    date_due DATE
);

-- 6. Questions
CREATE TABLE questions (
    question_id SERIAL PRIMARY KEY,
    quiz_id INTEGER REFERENCES quizzes(quiz_id) ON DELETE CASCADE,
    question_text TEXT NOT NULL,
    question_type TEXT CHECK (question_type IN ('multiple_choice', 'true_false', 'short_answer')) NOT NULL
);

-- 7. Answers
CREATE TABLE answers (
    answer_id SERIAL PRIMARY KEY,
    question_id INTEGER REFERENCES questions(question_id) ON DELETE CASCADE,
    answer_text TEXT NOT NULL,
    is_correct BOOLEAN NOT NULL
);

-- 8. Student_Quiz_Progress
CREATE TABLE student_quiz_progress (
    progress_id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES students(student_id) ON DELETE CASCADE,
    quiz_id INTEGER REFERENCES quizzes(quiz_id) ON DELETE CASCADE,
    score DECIMAL(5,2),
    completed_at TIMESTAMP
);

-- 9. Student_Lesson_Progress
CREATE TABLE student_lesson_progress (
    progress_id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES students(student_id) ON DELETE CASCADE,
    lesson_id INTEGER REFERENCES lessons(lesson_id) ON DELETE CASCADE,
    status TEXT CHECK (status IN ('completed', 'in_progress')) NOT NULL
);

-- 10. Assignments
CREATE TABLE assignments (
    assignment_id SERIAL PRIMARY KEY,
    course_id INTEGER REFERENCES courses(course_id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT,
    due_date DATE
);

-- 11. Student_Assignment_Submissions
CREATE TABLE student_assignment_submissions (
    submission_id SERIAL PRIMARY KEY,
    assignment_id INTEGER REFERENCES assignments(assignment_id) ON DELETE CASCADE,
    student_id INTEGER REFERENCES students(student_id) ON DELETE CASCADE,
    submission_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    file_url TEXT,
    grade DECIMAL(5,2)
);

-- 12. Notifications
CREATE TABLE notifications (
    notification_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE,
    message TEXT NOT NULL,
    status TEXT CHECK (status IN ('read', 'unread')) DEFAULT 'unread',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 13. Certificates
CREATE TABLE certificates (
    certificate_id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES students(student_id) ON DELETE CASCADE,
    course_id INTEGER REFERENCES courses(course_id) ON DELETE CASCADE,
    issued_date DATE DEFAULT CURRENT_DATE,
    certificate_url TEXT
);

-- 14. Payments
CREATE TABLE payments (
    payment_id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES students(student_id) ON DELETE CASCADE,
    amount DECIMAL(10,2) NOT NULL,
    payment_method TEXT CHECK (payment_method IN ('credit_card', 'bank_transfer')),
    payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status TEXT CHECK (status IN ('completed', 'pending')) NOT NULL
);
