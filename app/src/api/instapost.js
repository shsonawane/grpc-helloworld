import { InstaPostRequest } from '../gen/instaposts_pb';
import  { InstaPostPromiseClient }  from  '../gen/instaposts_grpc_web_pb';

const client = new InstaPostPromiseClient('http://localhost:8080', null, null);

export const getPostStats = async () => {
    const request = new InstaPostRequest();
    request.setId("TestPost");
    const response = await client.getPosts(request);
    return response.getStats();
}