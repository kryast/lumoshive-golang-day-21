<!-- method create user -->
curl -X POST http://localhost:8080/create -d '{"name":"Ahmad 1", "email":"test1@gmail.com", "password":"pass"}'

<!-- method login user -->
curl -X POST  http://localhost:8080/login -d '{"email":"test1@gmail.com", "password":"pass"}'

<!-- method create task default status incomplete without middleware -->
curl -X POST http://localhost:8080/create-task -d '{"description":"description 1"}'

<!-- method create task default status incomplete with middleware -->
curl -X POST http://localhost:8080/create-task -H "token: 12345" -H "role: user" -d '{"description": "description 6"}'

<!-- method read all task without middleware-->
curl -X GET http://localhost:8080/read-task

<!-- method read all task with middleware -->
curl -X GET http://localhost:8080/read-task -H "token: 12345" -H "role: user"

<!-- method update task by ID without middleware -->
curl -X POST http://localhost:8080/update-task -d '{"id": 1}'

<!-- method update task by ID with middleware -->
curl -X POST http://localhost:8080/update-task -H "token: 12345" -H "role: user" -d '{"id": 4}'