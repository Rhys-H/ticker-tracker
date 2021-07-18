# ticker-tracker

Track mentions of stock symbols on social media and analyze the impact on trade volume and price fluctuation.

I initially only plan on using Reddit as a source of data but will consider other data sources, such as Twitter, in the future.

## Usage

- Uses Alpaca API to retrieve security information, so you will need to create
  an account if you want to run the application
- Create a file in the project root, `.env`, with the following variables:

  - `API_KEY=<YOUR_API_KEY>`
  - `API_SECRET=<YOUR_API_SECRET>`

- Run from CLI using:
  ```shell
  go run .
  ```
