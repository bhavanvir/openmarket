# OpenMarket

## Description

OpenMarket is a ready-to-use, AWS-deployable wrapper for the GraphQL Facebook Marketplace API. It serves as a drop-in replacement for complex, scraping-based solutions, streamlining access to Facebook Marketplace data. This solution is designed to be easy to deploy at the cost of increased latency.

This project is designed for the hobbyist or small business owner who wants to access Facebook Marketplace data without the need for complex scraping solutions. It is not intended for large-scale commercial use, as doing so may violate Facebook's terms of service, and may result in your IP address being blocked.

The data is also not cleaned or processed in any way, so you may need to perform additional processing to get the data in the format you need.

## Features

- **Easy Deployment:** Quickly deployable on AWS with minimal setup.
- **Proxy Support:** Integrated support for forward proxy services to bypass Facebook's rate limiting.
- **API Gateway Integration:** Uses AWS API Gateway for managing API requests.

## Requirements

To bypass Facebook's rate limiting in a production environment, use a forward proxy service like [Bright Data](https://brightdata.com). For development purposes, you can use your own IP address.

To deploy OpenMarket, ensure you have the following:

- An AWS account
- AWS CLI installed and configured
- SAM CLI installed and configured

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/bhavanvir/openmarket.git
   ```

2. **Navigate to the project directory:**

   ```bash
   cd openmarket
   ```

3. **Create a `.env.json` file:**

   Use the structure outlined in the `.env.example.json` file.

4. **Deploy the project to AWS:**

   ```bash
   make deploy
   ```

5. **Configure environment variables:**

   - Go to AWS Lambda in the AWS Console.
   - Locate the `OpenMarket` function.
   - Navigate to the `Configuration` tab.
   - Under `Environment variables`, add the following variables from the `.env.json` file:
     - `PROXY_HOST`: URL of the proxy service.
     - `PROXY_USERNAME`: Username for the proxy service.
     - `PROXY_PASSWORD`: Password for the proxy service.
   - Save the environment variables.

Once deployed, an API Gateway URL will be provided for interacting with the GraphQL API. Without a proxy service, requests will be rate limited by Facebook.

## Usage

To interact with the API, use tools like `curl` or `Postman`. Provide the `id` of the listing you want to retrieve, which can be found in the URL of the listing on Facebook Marketplace.

https://github.com/bhavanvir/openmarket/assets/20825496/8cfa4e82-a23b-4a7d-aaae-3387eee9b128

### Making Requests

Send a `POST` request to the corresponding URL with the following `JSON` body:

```json
{
  "id": "<listing_id>"
}
```

#### Local API URL

To test OpenMarket locally, use the following command:

```bash
make local
```

This will start a local server at `http://127.0.0.1:3000/l/id`.

#### AWS API URL

To test OpenMarket on AWS, use the API Gateway URL provided after deployment. The URL will be in the following format:

```bash
https://${ServerlessRestApi}.execute-api.${AWS::Region}.amazonaws.com/Prod/l/{id}
```

### Example `curl` Command

```bash
curl -X POST https://your-api-gateway-url/Prod/l/{id} -H "Content-Type: application/json" -d '{"id": "<listing_id>"}'
```

### Example Postman Request

1. Open Postman and create a new `POST` request.
2. Enter the API URL: `https://your-api-gateway-url/Prod/l/{id}/Prod/l/{id}`
3. In the `Body` tab, select `raw` and set the type to `JSON`.
4. Paste the JSON body:

   ```json
   {
     "id": "<listing_id>"
   }
   ```

5. Send the request.

## Troubleshooting

- **Rate Limiting:** Ensure you are using a proxy service to avoid Facebook's rate limiting, especially in a production environment. If you are still facing issues, try using a different proxy service, and or logging your output by means of print statements and monitoring the logs on AWS CloudWatch.
- **Environment Variables:** Double-check that all environment variables are correctly set in the AWS Lambda configuration.
- **Deployment Issues:** Verify your AWS CLI and SAM CLI configurations and ensure they are correctly installed.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
