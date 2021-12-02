import { Client, ClientConfig, Order } from './restaurant/client';

(async () => {
  const clientConfig: ClientConfig = { baseUrl: 'http://localhost:8080' };
  const client = Client(clientConfig);

  // Get list of menu items 
  const menu = await client.listMenu();
  console.log(menu);

  // Add new Order
  const order: Order = {
    table: 10,
    items: [
      { itemId: 3, quantity: 2 },
      { itemId: 5, quantity: 1 },
      { itemId: 6, quantity: 3 },
    ],
  };
  const newOrder = await client.addNewOrder(order);
  console.log(newOrder);

})()