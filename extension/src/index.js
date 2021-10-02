import { extend, render, useExtensionInput, BlockStack, Button, Heading, Image } from '@shopify/post-purchase-ui-extensions-react';

async function fetchPostPurchaseData() {
  // This is where you would make a request to your app server to fetch the data
  return {
    productTitle: 'Fertilizer',
    productImageURL: 'https://cdn.shopify.com/s/files/1/0551/4084/3576/products/fertilizer_95d61198-8f34-4d2e-85f7-97f3a8d994c5_320x320@2x.jpg',
  };
}

extend('Checkout::PostPurchase::ShouldRender', async ({ storage }) => {
  const postPurchaseData = await fetchPostPurchaseData();
  await storage.update(postPurchaseData);

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

