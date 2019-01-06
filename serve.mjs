// -*- javascript -*-

import process from "process";
import http from "http";
import url from "url";

const server = http.createServer();

server.on('request', async (req, res)=> {
    if (req.method === 'GET' && req.url === '/') {
        res.statusCode = 200;
        res.setHeader('content-type', 'text/plain; charset=utf-8');
        res.end('k');
    } else {
        res.statusCode = 404;
        res.end();
    }
});

server.listen(process.env['PORT'] || 8080);
