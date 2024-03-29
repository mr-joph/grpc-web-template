// @generated by protobuf-ts 2.9.3 with parameter output_javascript_es2018
// @generated from protobuf file "proto/todolist.proto" (package "core", syntax proto3)
// tslint:disable
// @generated by protobuf-ts 2.9.3 with parameter output_javascript_es2018
// @generated from protobuf file "proto/todolist.proto" (package "core", syntax proto3)
// tslint:disable
import { Empty } from "../google/protobuf/empty";
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { WireType } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { Todo } from "./todo";
// @generated message type with reflection information, may provide speed optimized methods
class GetTodosOut$Type extends MessageType {
    constructor() {
        super("core.GetTodosOut", [
            { no: 1, name: "todos", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Todo }
        ]);
    }
    create(value) {
        const message = globalThis.Object.create((this.messagePrototype));
        message.todos = [];
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated core.Todo todos */ 1:
                    message.todos.push(Todo.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* repeated core.Todo todos = 1; */
        for (let i = 0; i < message.todos.length; i++)
            Todo.internalBinaryWrite(message.todos[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message core.GetTodosOut
 */
export const GetTodosOut = new GetTodosOut$Type();
// @generated message type with reflection information, may provide speed optimized methods
class NewTodoIn$Type extends MessageType {
    constructor() {
        super("core.NewTodoIn", [
            { no: 1, name: "content", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value) {
        const message = globalThis.Object.create((this.messagePrototype));
        message.content = "";
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string content */ 1:
                    message.content = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* string content = 1; */
        if (message.content !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.content);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message core.NewTodoIn
 */
export const NewTodoIn = new NewTodoIn$Type();
// @generated message type with reflection information, may provide speed optimized methods
class NewTodoOut$Type extends MessageType {
    constructor() {
        super("core.NewTodoOut", [
            { no: 1, name: "todo", kind: "message", T: () => Todo }
        ]);
    }
    create(value) {
        const message = globalThis.Object.create((this.messagePrototype));
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* core.Todo todo */ 1:
                    message.todo = Todo.internalBinaryRead(reader, reader.uint32(), options, message.todo);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* core.Todo todo = 1; */
        if (message.todo)
            Todo.internalBinaryWrite(message.todo, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message core.NewTodoOut
 */
export const NewTodoOut = new NewTodoOut$Type();
// @generated message type with reflection information, may provide speed optimized methods
class EditTodoOut$Type extends MessageType {
    constructor() {
        super("core.EditTodoOut", [
            { no: 1, name: "todo", kind: "message", T: () => Todo }
        ]);
    }
    create(value) {
        const message = globalThis.Object.create((this.messagePrototype));
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* core.Todo todo */ 1:
                    message.todo = Todo.internalBinaryRead(reader, reader.uint32(), options, message.todo);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* core.Todo todo = 1; */
        if (message.todo)
            Todo.internalBinaryWrite(message.todo, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message core.EditTodoOut
 */
export const EditTodoOut = new EditTodoOut$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CheckTodoIn$Type extends MessageType {
    constructor() {
        super("core.CheckTodoIn", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "completed", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = globalThis.Object.create((this.messagePrototype));
        message.id = "";
        message.completed = false;
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                case /* bool completed */ 2:
                    message.completed = reader.bool();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        /* bool completed = 2; */
        if (message.completed !== false)
            writer.tag(2, WireType.Varint).bool(message.completed);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message core.CheckTodoIn
 */
export const CheckTodoIn = new CheckTodoIn$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CheckTodoOut$Type extends MessageType {
    constructor() {
        super("core.CheckTodoOut", [
            { no: 1, name: "todo", kind: "message", T: () => Todo }
        ]);
    }
    create(value) {
        const message = globalThis.Object.create((this.messagePrototype));
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* core.Todo todo */ 1:
                    message.todo = Todo.internalBinaryRead(reader, reader.uint32(), options, message.todo);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* core.Todo todo = 1; */
        if (message.todo)
            Todo.internalBinaryWrite(message.todo, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message core.CheckTodoOut
 */
export const CheckTodoOut = new CheckTodoOut$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteTodoIn$Type extends MessageType {
    constructor() {
        super("core.DeleteTodoIn", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value) {
        const message = globalThis.Object.create((this.messagePrototype));
        message.id = "";
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message core.DeleteTodoIn
 */
export const DeleteTodoIn = new DeleteTodoIn$Type();
/**
 * @generated ServiceType for protobuf service core.TodoListService
 */
export const TodoListService = new ServiceType("core.TodoListService", [
    { name: "GetTodos", options: {}, I: Empty, O: GetTodosOut },
    { name: "NewTodo", options: {}, I: NewTodoIn, O: NewTodoOut },
    { name: "EditTodo", options: {}, I: Todo, O: EditTodoOut },
    { name: "CheckTodo", options: {}, I: CheckTodoIn, O: CheckTodoOut },
    { name: "DeleteTodo", options: {}, I: DeleteTodoIn, O: Empty },
    { name: "CompleteAll", options: {}, I: Empty, O: Empty },
    { name: "ClearCompleted", options: {}, I: Empty, O: Empty }
]);
