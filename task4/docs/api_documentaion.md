
# Task Management API

## Installation & Usage
clone the repo
```
git clone git@github.com:abrishk26/a2sv-project-track.git
```

change directory
```
cd a2sv-project-track/task4
```

install dependency
```
go mod tidy
```

run the application
``` 
    go run .
```

## API EndPoint


### Get All Tasks

`GET /tasks`

Retrieves a list of all available tasks.

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
    },
    {
      "title": "Task 2 Title",
      "description": "Description of Task 2.",
      "due_date": "2025-11-15",
      "done": true
    }
  ]
}
```

**Error Response:**
```json
{
  "error": "Error message describing the issue."
}
```

### Get Task Details

`GET /tasks/:id`

Retrieves the details of a specific task identified by its unique ID.

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

**Error Response:**
```json
{
  "error": "Error message describing the issue."
}
```

### Update Task

`PUT /tasks/:id`

Updates an existing task identified by its unique ID.

**Headers:** `Content-Type: application/json`

**Request Body:**
```json
{
  "title": "New Task Title", //optional
  "description": "Updated description of the task.", //optional
  "due_date": "2025-12-31", //optional
  "done": true //optional
}
```

**Successful Response (same as GET /tasks/:id):**
```json
{
  "message": "Task updated successfully.",
  "task": {
    "title": "New Task Title",
    "description": "Updated description of the task.",
    "due_date": "2025-12-31",
    "done": true
  }
}
```

### Delete Task
`DELETE /tasks/:id`

Delete a specific task identified by its unique ID.

**Successful Response**
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

**Error Response**
```json
{
  "error": "Error message describing the issue."
}
```

### Create New Task

`POST /tasks`

Creates a new task.

**Headers:** `Content-Type: application/json`

**Request Body**
```json
{
  "title": "New Task Title",
  "description": "Description of the new task.",
  "due_date": "2025-11-30",
  "done": false
}
```

**Successful Response**
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

**Error Response**
```json
{
  "error": "Error message describing the issue."
}
```



