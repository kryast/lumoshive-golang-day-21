<!-- method create user -->
curl -X POST http://localhost:8080/create -d '{"name":"Ahmad 1", "email":"test1@gmail.com", "password":"pass"}'

<!-- method login user -->
curl -X POST  http://localhost:8080/login -d '{"email":"test1@gmail.com", "password":"pass"}'

<!-- method create task default status incomplete -->
curl -X POST http://localhost:8080/create-task -d '{"description":"description 1"}'