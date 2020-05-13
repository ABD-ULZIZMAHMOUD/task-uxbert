# Chat Client-Server
```
  this project support chat between admin and normal user 
```
# How it is work ?
```
1-make database in mysql "write you database name "
2-add .env file connection to database username and password 
3- run project "go build " 
4- excute binary file "task-uxbert"
```

# Api Documentation

https://documenter.getpostman.com/view/4006411/Szmh3xLS


# packages 

for requests https://github.com/gin-gonic/gin

for mysql database https://gorm.io/docs/

for socket https://github.com/gorilla/websocket

# database structure 
![aImage of Yaktocat](https://i.ibb.co/Gk3L9vF/chat.png)


##### user table 
```
have email , password , fullname 
token => this token to Authorize user this is support user login in one devise 
type => to detrmine is admin or normal user 
```

##### Room table 
 ```
user1 and user2 this is reference to two users want to chat with ather 
note normal user can chat with admin only and admin can chat with normal user only 
because this is customer serviec chat 

fullname1 and fullname2 reference to user1 and user2 fullnames 
that for increase performance because probability of user change his name is low 
so we will not make join to get users data in rooms 

last message to show rooms whit last message  
```

##### message table 
```
content message text 

room_id referance this message connect to any room

sender and reseiver to make front-end when open room put messages in correct order 
by check user_id he login in with sender and put his message in correct order 

is_read to known if receiver read messsage or not 

```