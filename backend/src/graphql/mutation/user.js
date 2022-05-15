import { UserTC } from "../../models/user"

export const createUser = UserTC.getResolver('createOne')

export const updateUser = UserTC.getResolver('updateById')

export const deleteUser = UserTC.getResolver('removeById')