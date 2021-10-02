post purchase extension for Carbon Offset Post Purchase Upsell Shopify app
===================================================================

In development, the post purchase extensions runs locally, so you'll need the Shopify CLI [CLI instructions](https://shopify.dev/apps/tools/cli/installation) and you'll need a browser extension [Chrome version](https://cdn.shopify.com/static/checkout-post-purchase/dev-browser-extension/chrome-0.1.0-latest.zip?shpxid=40c7b06c-B636-4357-CDE3-D31443B5FF2E). (When it is deployed it will be hosted on Shopify.)

You'll also need a .env file. This should work for now if you fill in your API Key, Secret, and your shop. You can give it whatever you want to the title.

`SHOPIFY_API_KEY=theAPIKeyForYourTestApp
SHOPIFY_API_SECRET=theAPISecretForYourTestApp
SHOP=your-shop.myshopify.com
SCOPES=write_products,write_customers,write_draft_orders
EXTENSION_TITLE=carbon-offset-upsell
EXTENSION_UUID=testing`

(When you actually register and push the extension, you'll need a real UUID. The CLI creates one for you, but I think they'll just take what you give them, so I'm not sure of any reason why you can't generate your own.)

In the extension directory, run `shopify extension serve`. Then make a test purchase on your test store, and you should see the upsell afterwards.
## Functionality
- Not much yet :)
- Show an offer after a customer makes a purchase