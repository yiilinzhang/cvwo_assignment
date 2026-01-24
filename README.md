# CVWO Assignment 

This sample CVWO project is written in golang + react. It can be accessed via the Render link or ran locally.

## Getting Started

### Deployed app
Access the live app (frontend) at: https://cvwo-assignment-frontend.onrender.com
Test account:
Username: test
password: test
 
### Running locally
### Running the backend
Before running, create a new DB with the SQL in schema.sql at /internal/db/schema.sql
$ createdb cvwo_assignment
$ psql -d cvwo_assignment -f internal/db/schema.sql

To seed values:
psql -d cvwo_assignment -f internal/db/seeds.sql

Next, run the follow command at the root and configure the JWT_SECRET and the DATABASE_URL in the new .env file.

$ cp .env.example .env 

To start the backend, run 
$ go run cmd/server/main.go

To view the backend visit: http://localhost:8000

### Running the frontend
Next, cd into the frontend folder and run to make your copy of the env file located at frontend/.env.example

$ cp .env.example .env

Install dependencies:
$ npm install

Run the app:
npm run dev

To check on the app, visit http://localhost:5173

You may create a new account or use the test user at:
Username: alice
Password: alice123

