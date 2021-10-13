/* eslint-disable prettier/prettier */
/* eslint-disable no-undef */
import {useEffect, useState} from 'react';
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

const BACKEND_URL = 'http://localhost:4000/v1';  // TODO: change this to the real host!

extend('Checkout::PostPurchase::ShouldRender', async ({ storage, inputData }) => {
  // eslint-disable-next-line no-undef
  // eslint-disable-next-line prettier/prettier

  const postPurchaseOffer = await fetch(
    `${BACKEND_URL}/offer?shop=` + inputData.shop.domain,
  ).then((res) => res.json());
  await storage.update(postPurchaseOffer);
  return { render: true };
});


render('Checkout::PostPurchase::Render', () => <App />);

export function App() {
  const {
    storage,
    inputData,
    calculateChangeset,
    applyChangeset,
    done,
  } = useExtensionInput();
  const [loading, setLoading] = useState(true);
  const [calculatedPurchase, setCalculatedPurchase] = useState();

  useEffect(() => {
    async function calculatePurchase() {
      // Request Shopify to calculate shipping costs and taxes for the upsell
      const result = await calculateChangeset({changes});
      setCalculatedPurchase(result.calculatedPurchase);
      setLoading(false);
    }

    calculatePurchase();
  }, []);

  const {
    variantId,
    productTitle,
    productImageURL,
    productDescription,
  } = storage.initialData;
  const changes = [{type: 'add_variant', variantId, quantity: 1}];
  // Extract values from the calculated purchase
  const shipping =
      calculatedPurchase?.addedShippingLines[0]?.priceSet?.presentmentMoney
          ?.amount;
  const taxes =
      calculatedPurchase?.addedTaxLines[0]?.priceSet?.presentmentMoney?.amount;
  const total = calculatedPurchase?.totalOutstandingSet.presentmentMoney.amount;
  const discountedPrice =
      calculatedPurchase?.updatedLineItems[0].totalPriceSet.presentmentMoney
          .amount;
  const originalPrice =
      calculatedPurchase?.updatedLineItems[0].priceSet.presentmentMoney.amount;

  async function acceptOffer() {
    setLoading(true);

    // Make a request to your app server to sign the changeset
    const token = await fetch(`${BACKEND_URL}/sign-changeset`, {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify({
        referenceId: inputData.initialPurchase.referenceId,
        changes: changes,
        token: inputData.token,
      }),
    })
        .then((response) => response.json())
        .then((response) => response.token);

    // Make a request to Shopify servers to apply the changeset
    await applyChangeset(token);


    // Redirect to the thank-you page
    done();
  }

  function declineOffer() {
    setLoading(true);
    done();
  }

  return (
      <BlockStack spacing="loose">
        <CalloutBanner>
          <BlockStack spacing="tight">
            <TextContainer>
              <Text size="medium" emphasized>
                Take Climate action today!
              </Text>
            </TextContainer>
            <TextContainer>
              <Text size="medium">Add the offer below to your order </Text>
              <Text size="large" emphasized>
                and help fight climate change!
              </Text>
            </TextContainer>
          </BlockStack>
        </CalloutBanner>
        <Layout
            media={[
              {viewportSize: 'small', sizes: [1, 0, 1], maxInlineSize: 0.9},
              {viewportSize: 'medium', sizes: [532, 0, 1], maxInlineSize: 420},
              {viewportSize: 'large', sizes: [560, 38, 340]},
            ]}
        >
          <Image description="product photo" source={productImageURL} />
          <BlockStack />
          <BlockStack>
            <Heading>{productTitle}</Heading>
            <PriceHeader
                discountedPrice={discountedPrice}
                originalPrice={originalPrice}
                loading={!calculatedPurchase}
            />
            <ProductDescription textLines={productDescription} />
            <BlockStack spacing="tight">
              <Separator />
              <MoneyLine
                  label="Subtotal"
                  amount={discountedPrice}
                  loading={!calculatedPurchase}
              />
              <MoneyLine
                  label="Shipping"
                  amount={shipping}
                  loading={!calculatedPurchase}
              />
              <MoneyLine
                  label="Taxes"
                  amount={taxes}
                  loading={!calculatedPurchase}
              />
              <Separator />
              <MoneySummary
                  label="Total"
                  amount={total}
                  loading={!calculatedPurchase}
              />
            </BlockStack>
            <BlockStack>
              <Button onPress={acceptOffer} submit loading={loading}>
                Pay now · {formatCurrency(total)}
              </Button>
              <Button onPress={declineOffer} subdued loading={loading}>
                Decline this offer
              </Button>
            </BlockStack>
          </BlockStack>
        </Layout>
      </BlockStack>
  );
}

function PriceHeader({discountedPrice, originalPrice, loading}) {
  return (
      <TextContainer alignment="leading" spacing="loose">
        <Text size="large">
          {!loading && formatCurrency(originalPrice)}
        </Text>
      </TextContainer>
  );
}

function ProductDescription({textLines}) {
  return (
      <BlockStack spacing="xtight">
            <TextBlock subdued>
              {textLines}
            </TextBlock>
      </BlockStack>
  );
}

function MoneyLine({label, amount, loading = false}) {
  return (
      <Tiles>
        <TextBlock size="small">{label}</TextBlock>
        <TextContainer alignment="trailing">
          <TextBlock emphasized size="small">
            {loading ? '-' : formatCurrency(amount)}
          </TextBlock>
        </TextContainer>
      </Tiles>
  );
}

function MoneySummary({label, amount}) {
  return (
      <Tiles>
        <TextBlock size="medium" emphasized>
          {label}
        </TextBlock>
        <TextContainer alignment="trailing">
          <TextBlock emphasized size="medium">
            {formatCurrency(amount)}
          </TextBlock>
        </TextContainer>
      </Tiles>
  );
}

function formatCurrency(amount) {
  if (!amount || parseInt(amount, 10) === 0) {
    return 'Free';
  }
  return `$${amount}`;
}


