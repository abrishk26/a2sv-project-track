# Task Management API

## Installation & Usage

Clone the repository:
```bash
git clone git@github.com:abrishk26/a2sv-project-track.git
```

Change directory:
```bash
cd a2sv-project-track/task6
```

Create a `.env` file:
```bash
touch .env
```

Set your MongoDB Atlas connection string and secret key in the `.env` file:
```env
mongo=<your MongoDB Atlas connection string>
SECRET_KEY=<your secret key>
```
> Make sure your IP address is allowed in your MongoDB Atlas network access settings.

Install dependencies:
```bash
go mod tidy
```

Run the application:
```bash
go run .
```

---

## Authentication

All endpoints **require a Bearer token** in the `Authorization` header **except** the following:

- `POST /users/register` – register a new user  
- `POST /users/login/:id` – login an existing user (requires `id` from register response)

Both require a JSON body with `username` and `password` fields.

---

## API Endpoints

### Register User
`POST /users/register`

**Request Body:**
```json
{
  "username": "your_username",
  "password": "your_password"
}
```

**Successful Response:**
```json
{
  "message": "User registered successfully.",
  "user": {
    "id": "USER_ID",
    "username": "your_username",
    "role": "user"
  }
}

```

---

### Login User
`POST /users/login/:id`

**Request Body:**
```json
{
  "username": "your_username",
  "password": "your_password"
}
```

**Successful Response:**
```json
{
  "message": "Login successful.",
  "token": "<JWT Token>"
}
```

---

### Get All Tasks
`GET /tasks`

**Requires:** Authorization Bearer Token

**Successful Response:**
```json
{
  "message": "Tasks retrieved successfully.",
  "tasks": [
    {
      "title": "Task 1 Title",
      "description": "Description of Task 1.",
      "due_date": "2025-10-26",
      "done": false
    }
  ]
}
```

---

### Get Task Details
`GET /tasks/:id`

**Requires:** Authorization Bearer Token

**Successful Response:**
```json
{
  "message": "Task retrieved successfully.",
  "task": {
    "title": "Task Title",
    "description": "Description of the task.",
    "due_date": "2025-12-31",
    "done": false
  }
}
```

---

### Create New Task
`POST /tasks`

**Requires:** Authorization Bearer Token (Admin)

**Request Body:**
```json
{
  "title": "New Task Title",
  "description": "Description of the new task.",
  "due_date": "2025-11-30",
  "done": false
}
```

**Successful Response:**
```json
{
  "message": "Task created successfully.",
  "task": {
    "title": "New Task Title",
    "description": "Description of the new task.",
    "due_date": "2025-11-30",
    "done": false
  }
}
```

---

### Update Task
`PUT /tasks/:id`

**Requires:** Authorization Bearer Token (Admin)

**Request Body:**
```json
{
  "title": "Updated Task Title",
  "description": "Updated description.",
  "due_date": "2025-12-31",
  "done": true
}
```

**Successful Response:**
```json
{
  "message": "Task updated successfully.",
  "task": {
    "title": "Updated Task Title",
    "description": "Updated description.",
    "due_date": "2025-12-31",
    "done": true
  }
}
```

---

### Delete Task
`DELETE /tasks/:id`

**Requires:** Authorization Bearer Token (Admin)

**Successful Response:**
```json
{
  "message": "Task deleted successfully.",
  "task": {
    "title": "Deleted Task Title",
    "description": "Description of the deleted task.",
    "due_date": "2025-09-01",
    "done": true
  }
}
```

---

### Get All Users
`GET /users`

**Requires:** Authorization Bearer Token (Admin)

---

### Get Single User
`GET /users/:id`

**Requires:** Authorization Bearer Token (Admin)

---

### Update User
`PUT /users/:id`

**Requires:** Authorization Bearer Token (Admin)

---

### Delete User
`DELETE /users/:id`

**Requires:** Authorization Bearer Token (Admin)

