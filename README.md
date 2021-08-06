# Student-api

First create a DB with name 'crud'

## Custom Port Facility

 User can pass the custom port facility by passing its value in the flag during the runtime call as " **go run \.main.go -port 8070**" .(Here I used 8070 as custom port)
 Otherwise it will open the application on **Port 8080 by default**.


# Available Endpoints

POST   /student                  
GET    /student                 
GET    /student/:id              
PUT    /student/:id              
DELETE /student/:id



### Validation

Email is checked for its correctness

Phone number is checked for its 10-digit length

Name is checked for its empty field


### Sorting 

Sorting is defined in the URL by using string "sort_field=`field name`" & "dir = `direction of sorting`"

For example "/student?sort_field=phone&sort_dir=asc" (Here we have choosen Column field as phone and Direction for Sorting as Ascending)


 ### Pagination
 
 Pagination is added on GET request of the API, it defines the custom limit and pages on request.
 
 ### Check For Existing Data (Phone Number or Email)
 
 In order to avoid duplicacy of Data, the API will check for Existing Data with similar Phone number and Email Id and will prevent from creating duplicate entries.
 
 
