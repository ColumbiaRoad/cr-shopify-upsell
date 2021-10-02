/* eslint-disable prettier/prettier */
/* eslint-disable no-undef */
import { useEffect, useState } from 'react';
import {
  extend,
  render,
  useExtensionInput,
  BlockStack,
  Button,
  CalloutBanner,
  Heading,
  Image,
  Text,
  TextContainer,
  Separator,
  Tiles,
  TextBlock,
  Layout,
} from '@shopify/post-purchase-ui-extensions-react';


// TODO: mock example, take this out when the offer endpoint is working
// async function fetchPostPurchaseData() {
//   // This is where you would make a request to your app server to fetch the data
//   return {
//     productTitle: 'Fertilizer',
//     productImageURL: 'https://cdn.shopify.com/s/files/1/0551/4084/3576/products/fertilizer_95d61198-8f34-4d2e-85f7-97f3a8d994c5_320x320@2x.jpg',
//   };
// }

const BACKEND_URL = 'http://localhost:4000/v1';

extend('Checkout::PostPurchase::ShouldRender', async ({ storage }) => {
  // eslint-disable-next-line no-undef
  // eslint-disable-next-line prettier/prettier
  const postPurchaseOffer = await fetch(
    `${BACKEND_URL}/offer`,
  ).then((res) => res.json());
  console.log("should be fetching")
  await storage.update(postPurchaseOffer);

  return { render: true };
});

render('Checkout::PostPurchase::Render', () => <App />);

export function App() {
  const { done, storage } = useExtensionInput();
  const { productTitle, productImageURL } = storage.initialData;
  return (
    <BlockStack spacing="loose" alignment="center">
      <Heading>{productTitle}</Heading>
      <Image source={productImageURL} />
      <Button submit onPress={done}>Click me</Button>
    </BlockStack>
  )
}

