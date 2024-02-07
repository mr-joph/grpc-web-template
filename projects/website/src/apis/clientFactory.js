import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport"

const coreTransport = new GrpcWebFetchTransport({
    format: "binary",
    baseUrl: "http://localhost:8080",
})

export function createClient(serviceClass) {
    return new serviceClass(coreTransport)
}