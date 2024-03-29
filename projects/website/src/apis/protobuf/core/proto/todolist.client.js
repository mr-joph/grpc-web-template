// @generated by protobuf-ts 2.9.3 with parameter output_javascript_es2018
// @generated from protobuf file "proto/todolist.proto" (package "core", syntax proto3)
// tslint:disable
import { TodoListService } from "./todolist";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service core.TodoListService
 */
export class TodoListServiceClient {
    constructor(_transport) {
        this._transport = _transport;
        this.typeName = TodoListService.typeName;
        this.methods = TodoListService.methods;
        this.options = TodoListService.options;
    }
    /**
     * @generated from protobuf rpc: GetTodos(google.protobuf.Empty) returns (core.GetTodosOut);
     */
    getTodos(input, options) {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: NewTodo(core.NewTodoIn) returns (core.NewTodoOut);
     */
    newTodo(input, options) {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: EditTodo(core.Todo) returns (core.EditTodoOut);
     */
    editTodo(input, options) {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: CheckTodo(core.CheckTodoIn) returns (core.CheckTodoOut);
     */
    checkTodo(input, options) {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: DeleteTodo(core.DeleteTodoIn) returns (google.protobuf.Empty);
     */
    deleteTodo(input, options) {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: CompleteAll(google.protobuf.Empty) returns (google.protobuf.Empty);
     */
    completeAll(input, options) {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: ClearCompleted(google.protobuf.Empty) returns (google.protobuf.Empty);
     */
    clearCompleted(input, options) {
        const method = this.methods[6], opt = this._transport.mergeOptions(options);
        return stackIntercept("unary", this._transport, method, opt, input);
    }
}
