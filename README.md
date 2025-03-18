# Microservices POC Project

This project is a **proof-of-concept (POC)** microservices architecture built with **Go (Golang)**, using multiple services to manage courses, students, and enrollments. It includes a **GraphQL gateway** for API aggregation and service communication via **REST APIs**.

## 📌 Project Structure

```
├── course/          # Course service (manages courses)
├── student/         # Student service (manages students)
├── enrollment/      # Enrollment service (handles student enrollments in courses)
├── gateway/         # GraphQL API Gateway (aggregates data from other services)
├── docker-compose.yaml  # Docker Compose file for running services
├── .gitignore       # Git ignore file
```

## 🏗️ Microservices Overview

### 1️⃣ Course Service (`course/`)

- Provides CRUD operations for managing **courses**.
- Exposes REST API endpoints.
- Runs on **port 8082**.

### 2️⃣ Student Service (`student/`)

- Manages **student registrations**.
- Exposes REST API endpoints.
- Runs on **port 8081**.

### 3️⃣ Enrollment Service (`enrollment/`)

- Manages **student enrollments** in courses.
- Verifies student and course existence before enrollment.
- Runs on **port 8083**.

### 4️⃣ API Gateway (`gateway/`)

- Provides a **GraphQL API** to unify data access.
- Fetches data from `course`, `student`, and `enrollment` services.
- Runs on **port 8080**.

## 🚀 Getting Started

### 1️⃣ Clone the Repository

```sh
 git clone https://github.com/your-repo/microservices-poc.git
 cd microservices-poc
```

### 2️⃣ Run with Docker Compose

```sh
 docker-compose up --build
```

This will start all services and expose the **GraphQL API at `http://localhost:8080`**.

### 3️⃣ Testing the APIs

- **GraphQL Playground:** [http://localhost:8080/](http://localhost:8080/)
- **Course API:** `GET http://localhost:8082/api/courses`
- **Student API:** `GET http://localhost:8081/api/students`
- **Enrollment API:** `GET http://localhost:8083/api/enrollments`

## 🛠️ Technologies Used

- **Go (Golang)** - Language for all microservices
- **Echo** - Web framework for REST APIs
- **GraphQL (GQLGen)** - API Gateway for data aggregation
- **Docker & Docker Compose** - Containerization and orchestration
- **Distroless Base Image** - Minimal container image for security

## 📜 API Endpoints

### Course Service (`/course`)

| Method | Endpoint           | Description         |
| ------ | ------------------ | ------------------- |
| GET    | `/api/courses`     | Get all courses     |
| GET    | `/api/courses/:id` | Get course by ID    |
| POST   | `/api/courses`     | Create a new course |

### Student Service (`/student`)

| Method | Endpoint            | Description          |
| ------ | ------------------- | -------------------- |
| GET    | `/api/students`     | Get all students     |
| GET    | `/api/students/:id` | Get student by ID    |
| POST   | `/api/students`     | Create a new student |

### Enrollment Service (`/enrollment`)

| Method | Endpoint               | Description                  |
| ------ | ---------------------- | ---------------------------- |
| GET    | `/api/enrollments`     | Get all enrollments          |
| GET    | `/api/enrollments/:id` | Get enrollment by ID         |
| POST   | `/api/enrollments`     | Enroll a student in a course |

## ⚡ GraphQL Queries & Mutations

### Get all Students and Courses

```graphql
query {
	students {
		id
		name
		courses {
			title
		}
	}
	courses {
		id
		title
	}
}
```

### Enroll a Student in a Course

```graphql
mutation {
	enrollStudentInCourse(studentId: "1", courseId: "2") {
		title
		students {
			name
		}
	}
}
```

## 🛠️ Development & Contribution

1. Clone the repo and create a feature branch.
2. Modify the required service (`course`, `student`, `enrollment`, `gateway`).
3. Run `go mod tidy` to ensure dependencies are correct.
4. Test locally with `docker-compose up --build`.
5. Create a Pull Request for review.

## 🏁 Conclusion

This POC showcases a **scalable microservices architecture** using **Go, GraphQL, and Docker**. It can be extended with a **database layer, authentication, and caching** for production usage.
