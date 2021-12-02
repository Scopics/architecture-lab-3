import { client } from './http/client';

export interface ClientConfig {
  baseUrl: string,
}

export interface ClientMethods {
  listMenu: () => Promise<any>
}

export interface Order {
  table: number,
  items: Array<{
    itemId: number,
    quantity: number,
  }>
}

export const Client = ({ baseUrl }: ClientConfig) => {
  const httpClient = client(baseUrl);

  return {
    listMenu: () => httpClient.get('/restaurant'),
    addNewOrder: (items: Order) => httpClient.post('/restaurant', items),
  }
}
