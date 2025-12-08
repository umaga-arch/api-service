// types.ts

export type HttpMethod = 'GET' | 'POST' | 'PUT' | 'DELETE' | 'HEAD' | 'OPTIONS' | 'PATCH';

export type RoutePath = string;

export interface Route {
  path: RoutePath;
  method: HttpMethod;
  handler: (request: any, response: any) => void;
}

export interface Request {
  method: HttpMethod;
  path: RoutePath;
  headers: { [key: string]: string };
  body: any;
}

export interface Response {
  statusCode: number;
  headers: { [key: string]: string };
  body: any;
}