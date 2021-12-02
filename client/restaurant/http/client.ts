import { http } from "./http"

export const client = (baseUrl: string) => {
  return {
    post: (path: string, data: any = {}) => http(baseUrl + path, 'POST', data),
    get: (path: string) => http(baseUrl + path, 'GET'),
  }
}
