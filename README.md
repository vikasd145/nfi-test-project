# nfi-test-project
- Start the service by `docker-compose up -d` but first make sure docker daemon is running
- If due to some reason docker-compose not working create binary `go build ./...` and run binary `./nfi-test-project` 
- Above command will start service on port 8080 and postgres server on port 5432
- Also create databases and minmal table structure

- Register endpoint curl command ```curl --location --request POST 'localhost:8080/register'```
- Deposit endpoint curl command ```curl --location 'localhost:8080/deposit' \
  --form 'user_id="1"' \
  --form 'amount="100"'```
- Withdraw endpoint curl command ```curl --location 'localhost:8080/withdraw' \
  --form 'user_id="1"' \
  --form 'amount="100"' \```