import { model, Schema } from "mongoose";
import { composeWithMongoose } from 'graphql-compose-mongoose'

const UserSchema = new Schema({
    name: {
        type: String,
        required: true,
        unique: true,
        trim: true,

    },
    lastname: {
        type: String,
        required: true,
        trim: true
    },
    username: {
        type: String,
        lowercase:true,
        required: true,
        unique: true,
        trim: true,
    },
    password: {
        type: String,
        required: true,
        trim: true,
    }

}, {
    timestamps: true
})

export const UserModel = model('User', UserSchema)

export const UserTC = composeWithMongoose(UserModel)