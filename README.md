# Software Engineering(CSE 4123) Project API

* go
* gin
* gorm
* postgresql

### Installation & Run:
 
* Before running, you should set the database informations on  ***internal/orm/main.go***. 

Then on terminal;
```
scripts/run.sh
```
For build;
```
scripts/build.sh
```
## API
#### /files
* `GET` : Static folder.
### Without Token
#### /login
* `POST` : `{email:"a@b.com",password:"pass"}`
#### /register
* `POST` : `{name:"username",email:"a@b.com",password:"pass"}`
#### /ringtones
* `GET` : Returns all ringtones.
#### /ringtones-featured
* `GET` : Returns 3 featured ringtones.
#### /ringtones/:id
* `GET` : Returns ringtone detail.
#### /categories
* `GET` : Returns all categories.
#### /categories/:id
* `GET` : Returns categories ringtones.
#### /recent-activities
* `GET` : Returns recent activities.

### With User Token
#### /user/settings/:id
* `GET` : Returns user profile.
#### /user/ringtone/buy
* `POST` : Buy ringtone. `{"ringtone_id": 1, "user_id": 2}`
#### /user/ringtone/like
* `POST` : Like ringtone. `{"ringtone_id": 1, "user_id": 2}`
#### /user/ringtone/dislike
* `POST` : Dislike ringtone. `{"ringtone_id": 1, "user_id": 2}`
#### /user/ringtone/comment
* `POST` : Comment on ringtone. `{"ringtone_id": 1, "user_id": 2, "comment":"comment text"}`
 