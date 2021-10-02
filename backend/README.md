Backend service for Carbon Offset Post Purchase Upsell Shopify app
===================================================================

To run the service install docker, docker compose and `make run`


## Functionality
- Shopify App install flow including Billing API setup
- Subscribe to Shopify Webhooks
- Process Shopify Order Paid webhooks


## Database migrations using goose
- `make db/up env=local service=offset`
- `make db/down env=local service=offset`

## Misc
If you want to view the API docs and you are running it with Docker then point your browser to http://localhost:4000/v1/docs/index.html
