import { TodoListServiceClient } from "./protobuf/core/proto/todolist.client"
import { createClient } from "./clientFactory"

const client = createClient(TodoListServiceClient)

export async function getAllTodos() {
    try {
        const res = await client.getTodos({})
        console.log("GRPC RESPONSE: ", res)
        return res.response
    } catch (error) {
        console.error("ERROR: ", error)
        throw error
    }
}

export default client