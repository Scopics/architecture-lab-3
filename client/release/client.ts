import { fetch } from './fetch';


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
  
  return {
    listMenu: () => fetch(baseUrl + '/restaurant'),
    addNewOrder: (items: Order) => fetch(baseUrl + '/restaurant', 'POST', items),
  }
}
