# CVWO Assignment 

This sample CVWO project is written in golang + react. It can be accessed via the Render link or ran locally.

## Getting Started

### Running the backend
Before running, create a new DB with the SQL in scheme.sql.

Next, run the follow command at the root and configure the JWT_SECRET and the DATABASE_URL in the new .env file.

$ cp .env.example .env 

To start the backend, run 
$ go run cmd/server/main.go.

### Running the frontend
Next, cd into the frontend folder and run to make your copy of the env file.

$ cp .env.example .env

Next install dependencies with:
$ npm install

Run the app with:
npm run dev

To check on the app, visit http://localhost:5173
