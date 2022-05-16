import { model, Schema } from "mongoose";
import { composeWithMongoose } from 'graphql-compose-mongoose'

const UserSchema = new Schema({
    name: {
        type: String,
        required: true
    },
    lastname: {
        type: String,
        required: true
    }

})

export const UserModel = model('User', UserSchema)

export const UserTC = composeWithMongoose(UserModel)