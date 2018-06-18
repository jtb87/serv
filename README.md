# PhoneBook

## Build Steps

0. run command "go get" to download all dependencies  
1. set environment variables 
2. go build -o phonebook 
3. start executable (server will create logfile "phonebook.log" in same directory)
4. server is now running on localhost:9090 

## API routes 
#DB is seeded with a small number of entries

### /api/entry
method: GET

returns: a list of all entries in the phonebook 

### /api/entry?search={search}
method: GET

returns: a list of entries where the email address startswith the searched term

### /api/entry/{email} 
method: GET

returns: object with the specified resourceId (email address) 

### /api/entry/{email}
method: DELETE 

returns: deletes the resource with the specified resourceid (email address) 

###  /api/entry
method: POST

returns: creates a resource returns the resourceId
data: 
```javascript
{
"firstname":"string",
"lastname":"string",
"email":"string",
"phonenumber":"string",
}
```    

###  /api/entry
method: PUT

returns: edits the resource with the specified resourceId (emailaddress) 
data: 
```javascript
{
"firstname":"string",
"lastname":"string",
"email":"string",
"phonenumber":"string",
}
```    


