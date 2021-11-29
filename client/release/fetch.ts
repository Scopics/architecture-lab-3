import { request, IncomingMessage, RequestOptions } from 'http';

export type fetchMethod = 'GET' | 'PUT' | 'POST';

export const fetch = async (reqUrl: string, method: fetchMethod = 'GET', data?: any) => {
  const url = new URL(reqUrl);

  const options: RequestOptions = {
    hostname: url.hostname,
    port: url.port,
    path: url.pathname,
    method: method,
  };

  return new Promise((resolve, reject) => {
    const req = request(options, (res: IncomingMessage) => {
      let data: string = '';
  
      res.on("data", (chunk) => {
        data += chunk;
      });
      
      res.on("end", () => {
        const parsedData = JSON.parse(data);
        resolve(parsedData);
      });
    });

    req.on('error', (err) => {
      reject(err);
    });

    if (data) {
      req.write(JSON.stringify(data));
    }

    req.end();
  })

}