import express=require('express');
import * as bodyParser from 'body-parser';

const ExpressPort=3000;
class Server {
    private server=express();
    constructor() {
        
    }

    registerRoutes(){
        this.server.get('/showdown/:id',(request:express.Request,resposne:express.Response)=>{
            //console.log(request.params.id);
            resposne.status(200).send(request.params.id);
        });
        this.server.post('/showdown',(request:express.Request, response:express.Response)=>{
            //can not return only the nubmer of the request due to
            // express behaviour when sending only number as response to post
            // this becomes the status of the response
            response.status(200).send(request.body);
        });
    }

    startServer(){
        this.server.use(bodyParser.json());
        this.server.use(bodyParser.urlencoded({extended:true}));
        this.registerRoutes();
        this.server.listen(ExpressPort,()=>{
            console.log(`server is listening on port - ${ExpressPort}`);
        })
    }
}
let server=new Server();
server.startServer();