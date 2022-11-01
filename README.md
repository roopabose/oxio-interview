a. Explanation of how we would install and run your submission
    

Download and install Go from [here](https://go.dev/doc/install) 
    
To run the Rest API: Please go to your `terminal` and go to the project path and 
    execute below command:
    
`go mod tidy`

`go run .`

There are 2 ways you can execute the end point:
1. Use below Curl Operation:

   `curl --location --request GET 'http://localhost:8080/v1/phone-numbers?phoneNumber=%2B12125690123' \
   --header 'Content-Type: text/plain' \
   --data-raw '/v1/phone-numbers?phoneNumber=631%20311%208150'`
2. Go to Postman and type the url: http://localhost:8080/v1/phone-numbers?phoneNumber=%2B12125690123 and make sure the method is set to Get and press send button

To run the test cases.
Please go to the project path in your terminal and execute:


`go test`


b. Explanation of your choice of programming language, framework, library.
    

Have used Go as I have mentioned before, I am new to Go and I am just in love with the language as it requires minimum code 
    and it is also fast.


I have used `github.com/nyaruka/phonenumbers` library for phone number validations, as I researched and found that this library is used daily in production for parsing and validation of numbers across the world, so is well maintained.  

c. Explanation of how you would deploy to production.

Execute below command at your project path
`GOOS=linux GOARCH=amd64 go build
`
A new file called `jobs` has been created.This is the compiled binary file that is needed to execute our api.You just have to upload this binary file to your server and run it.

d. Explanation of assumptions you made


1: I have assumed, that the phone number can only miss the country code, that is, country code can be missing from the number for an errorneous case

2: I have assumed spaces will always be after the country code and area code/ Or no spaces at all (Did not validate according to spaces position)

3: CountryCode Parameter if given should be 2 Characters

4: The Url with Country code parameter will be as
`http://localhost:8080/v1/phone-numbers?phoneNumber=12125690123&countryCode=US
`
(the '1' is already there in the phone number)





e. Explanation of improvements you wish to make

1: The error handling can be better. In terms for error messages, I can have error msg enums with code, so that with different scenarios we can send code with messages

2: The handling of invalid cases, Had doubt regarding the messages for  invalid cases.

3: The Validation in terms of spaces, if we know the valid number of digits for country code,
for area code.
