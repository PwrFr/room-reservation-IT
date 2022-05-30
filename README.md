# room-reservation-IT

## Demo
### Student
- Login login with google OAuth

![image](https://user-images.githubusercontent.com/54875724/170971445-8c175b74-402f-4970-96ba-26ca1959e6a5.png)
- Home Page show all rooms

![image](https://user-images.githubusercontent.com/54875724/170971491-efa52da2-ec70-4691-8f78-a1424d87b7cc.png)
- Reserve room Feature Fill in Purpose, Data and Attendance for using this room 

![image](https://user-images.githubusercontent.com/54875724/170971582-b0435a7c-39a0-4220-a2c2-733d4a0cd1bc.png)
![image](https://user-images.githubusercontent.com/54875724/170971622-9b0135d6-a8e2-47d6-8203-0d2d91427efc.png)
- Reservation Page 

![image](https://user-images.githubusercontent.com/54875724/170971710-2a3a5ea3-2340-4144-8600-5b09228d790a.png)
- Canceling Reservation
- 
![image](https://user-images.githubusercontent.com/54875724/170972011-c54366b7-de5f-4fb4-81ce-fdf2b6e83976.png)
![image](https://user-images.githubusercontent.com/54875724/170972146-4cba3886-4e1e-4497-b1e1-8c485c5eaf45.png)
### Staff
- Approving Page

![image](https://user-images.githubusercontent.com/54875724/170972228-31f87cc1-198a-4888-8adf-fd776a5f3f32.png)
![image](https://user-images.githubusercontent.com/54875724/170972247-07121134-46c7-49e3-a9cb-70a97fad3c5e.png)
- Manage Room ( Cant show this Page cuz DB just running out ;-; ) This Page can Change status of room to Availible, Not availible and Maintainance IF 
1. room isnt using
2. request using room is out of date

![image](https://user-images.githubusercontent.com/54875724/170973868-be5109ee-3168-4d0f-98e8-179a5da5336b.png)


Use this email to Demo project on : https://statuesque-queijadas-fa6f75.netlify.app/
```
Student : 
gmail: student.it00@gmail.com
password: student@itkmitl1234
```
```
Staff :
gmail: staff01.it@gmail.com
password: staff@itkmitl1234
```
## How to run Frontend
```
cd web
npm i
npm run dev
```

## How to Create Schema
```
1. go to schema.graphqls then add your data Type, Query and Mutation
2. Delete schema.resolvers.go (Only you have new Query or Mutation)
3. go generate ./graph/
Model_gen.go and schema.resolver.go will automatic generate 
4. implement Resolver on schema.resolvers.go

Doc:
- https://gqlgen.com/
- https://github.com/99designs/gqlgen
```

## How to run Backend
```
cd api
go run server.go
```

## Technology
- NuxtJS
- Vite
- Vuetify

## Color Palette
![image](https://user-images.githubusercontent.com/54875724/167244726-1aa8ca6f-33b1-4484-a7fc-0e83aa023c57.png)

``` 
#DCDDE0
```
``` 
#747377
```
``` 
#4D4C4B
```
``` 
#D0754A
```

## UI Styling
- Css Framework : Tailwind
- UI Marterials : Vuetify 

## Vite-Nuxt With Nuxt/Axios
There is a problem that Vite-Nuxt version 0.3.5 Can't be use with Nuxt/Axios,
So i Downgraded Vite-Nuxt from 0.3.5 to 0.2.4

