import mongoose from 'mongoose'

// 
const uri = 'mongodb+srv://graphql-test.919sq.mongodb.net'
const options = {
    dbName: 'GraphQL-Test',
    user: 'pwrfr',
    pass: 'Knight1234'
}

export default mongoose.connect(uri, options)