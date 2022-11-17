# Microservice for managing users balance

## Project description.
Implemented 4 handlers: GetBalance, AddBalance, CreateOrder, FinishOrder

Code covered with tests

## Build instructions.
After git clone it should be enough to run in the project directory:

`make`

Service should be accessed via port 8080 on local host.

## HTTP API specification.
### /get-balance
Getting the user's balance.
Accepts a user id.

Sample request:

`curl -X POST http://localhost:8080/get-balance -H 'Content-Type: application/json' -d '{ "user_id" : 1 }'`

### /add-balance
Adding funds to the balance.
Accepts user id and how much money to deposit.

Sample request:

`curl -X POST http://localhost:8080/add-balance -H 'Content-Type: application/json' -d '{ "id" : 1, "service_id" : 2, "user_id" : 1, "amount" : 50 }'`

### /create-order
Reservation of funds from the main balance in a separate account.
Accepts user id, service id, order id, cost.

Sample request:

`curl -X POST http://localhost:8080/create-order -H 'Content-Type: application/json' -d '{ "id" : 1, "service_id" : 2, "user_id" : 1, "amount" : 50 }'`

### /finish-order
Recognition of revenue - writes off money from the reserve, adds data to the report for accounting.
Accepts user id, service id, order id, amount.

Sample request:

`curl -X POST http://localhost:8080/finish-order -H 'Content-Type: application/json' -d '{ "id" : 1, "service_id" : 2, "user_id" : 1, "amount" : 50}'`
