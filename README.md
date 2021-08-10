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

Email is checked for its validness.

Phone number is checked for its valid 10-digit length.

Name is checked if it is not empty.


### Sorting 

Sorting is defined in the URL by using string` "sort_field= field_name"  & "dir = direction_of_sorting" `

For example` "/student?sort_field=phone&sort_dir=asc" `(Here we have choosen Column field as phone and Direction for Sorting as Ascending)


 ### Pagination
 
 Pagination is added on GET request of the API, it defines the custom limit and pages on request.
 
 ### Check For Existing Data (Phone Number or Email)
 
 In order to avoid duplicacy of Data, the API will check for Existing Data with similar Phone number and Email Id and will prevent from creating duplicate entries.
 
 
 ### Data Filtering 
 
 User can now filter the data from the tables by Passing Search string as` "search= search_key " `(Here search_key is the value to be searched) in the URL. 
 It will filter all the data that matches the passed string. 
 Note:- It is not necessary to pass the full length of the search key, the API will search for all the data fields that matches the passed string wheather it is partial or full.
 
 ### Time Based Filtering
 
 User can now filter data on daily/weekly/monthly/yearly basis
 User has to pass the Filter String as ` "time= time_key"  `(Here time_key can be daily/weekly/monthly/yearly) in the URL.
 
 For Example = ` "\student?time=daily" ` (Here we are filtering based on data created in last 24hours or Daily Basis)
