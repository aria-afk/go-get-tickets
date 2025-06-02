# create dummy data

This should achieve a few things:

- One is just be a psuedo unit test of inserting data in a structured manner into the database
(ideally similar data flow to the actual UX)

- Two is for it to provide a base set of data that we can use during staging or development to test and build out
the api.

## TODO:
- Write a sql function to do the above in a structured order, make sure the data is somewhat meaningful for testing later on 
- Use embed to take that function and run it in its own main.go
- make it easy to reset/take down :) 
