# room-reservation-IT

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
- Tailwind
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

