import { HelloRequest } from '../gen/helloworld_pb';
import  { GreeterPromiseClient }  from  '../gen/helloworld_grpc_web_pb';

const client = new GreeterPromiseClient('http://localhost:8080', null, null);

export const sayHello = async (name) => {
    const request = new HelloRequest();
    request.setName(name);
    const response = await client.sayHello(request);
    return response.toObject();
}