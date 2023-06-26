# nfi-test-project
## Requirements
* Docker
* Docker Compose
## Setup and Running the Project
1. Clone the repository
```bash
git clone https://github.com/vikasd145/nfi-test-project.git
cd nfi-test-project
```
2. Build the Docker image and run it using Docker Compose,Below command will start service on port 8080 and postgres server on port 5432 Also create databases and minmal table structure
```bash
docker-compose up --build
```
- If due to some reason docker-compose not working run binary `./nfi-test-project` you need to have postgres server running on port 5432
- Can use postman or Curl for following command Register endpoint
```bash
 curl --location --request POST 'localhost:8080/register'
- ```
- Deposit endpoint curl command 
```bash
curl --location 'localhost:8080/deposit' \
  --form 'user_id="1"' \
  --form 'amount="100"'
  ```
- Withdraw endpoint curl command 
```bash
curl --location 'localhost:8080/withdraw' \
  --form 'user_id="1"' \
  --form 'amount="100"' \
  ```
## Possible Improvements
1. User Authentication and Authorization: The current version doesn't provide user authentication and authorization, which is essential in real-world applications. This can be implemented using JWT or OAuth protocols.

2. Tests: More comprehensive unit and integration tests can be written to ensure the application behaves as expected during development and before deployment.

3. Database Improvement: Can design databses table in better manner with better transactional abilities.

4. Rate Limiting: It might be useful to introduce rate limiting to protect the API from being overwhelmed by too many requests in a short period of time.

5. Pagination: If there are a lot of users, it would be more efficient to return results in pages rather than all at once.

6. Logging and Monitoring: Implementing robust logging and monitoring would be critical for spotting and understanding issues in production.

7. Environment Variables for Config: Right now, the configuration is hardcoded in a file. A good practice is to keep configurations that can change depending on the environment (like database credentials, hostnames, ports etc.) in environment variables.

8. API Documentation: Implementing API documentation with tools such as Swagger can be very helpful for the consumers of the API.