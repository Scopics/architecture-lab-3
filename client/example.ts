import { Client, ClientConfig, Order } from './restaurant/client';

(async () => {
  const clientConfig: ClientConfig = { baseUrl: 'http://localhost:8080' };
  const client = Client(clientConfig);

  console.log('<-- Scenario 1 Get list of menu items -->');
  try {
    const menu = await client.listMenu();
    console.log(menu);
  } catch (err: any) {
    console.log(`Unable to get list menu: ${err.message}`);
  }

  console.log();

  console.log('<-- Scenario 2 Add new Order -->');
  try {
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
  } catch (err: any) {
    console.log(`Unable to post a new order: ${err.message}`);
  }
})()