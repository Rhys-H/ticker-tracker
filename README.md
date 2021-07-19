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

## Example

![TSLA_example](https://user-images.githubusercontent.com/5068032/126120842-4a3026f1-62ce-4329-8bf6-4a8f8646ae89.png)

- Found two Reddit post titles containing the stock symbol that we are interested in, TSLA, for the given timeframe
- Displays the price movement as a percentage for the same timeframe
