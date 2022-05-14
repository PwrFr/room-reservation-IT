import { createServer } from 'http'
import express from 'express'
import ApolloServer from 'apollo-server-express'
import cors from 'cors'

const app = express()

app.use(express.json())
app.use(express.urlencoded({ extended: true }))
app.use(cors())

app.get('/', (req, res) => {
    res.json({ message: 'Server is running' })
})

const apolloServer = new ApolloServer({
    schema: '',

})

const httpServer = createServer(app)

httpServer.listen({ port: 3001 })