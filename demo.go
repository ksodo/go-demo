package main

import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"

/*
ticker response json

{
  "code": 0,
  "message": "Operation successful",
  "data": {
    "at": 1548123813,
    "ticker": {
      "buy": "0.0000353",
      "sell": "0.00003532",
      "low": "0.00003426",
      "high": "0.00003572",
      "last": "0.00003531",
      "vol": "25430185.0"
    }
  }
}
*/

type TickersDataTicker struct {
  Buy string
  Sell string
  Low string
  High string
  Last string
  Vol string
}

type TickersData struct {
  At int32
  Ticker TickersDataTicker
}

type Tickers struct {
  Code int32
  Message string
  Data TickersData
}

//var http_api_base_url string = "https://api.oceanex.pro/v1"
//var http_api_base_url string = "https://api.oceanex.cc/v1"
var http_api_base_url string = "https://api-minikube.thorex.me/v1"

func GetTickers(symbol string) *Tickers {
  url := http_api_base_url + "/tickers/" + symbol
  resp, err := http.Get(url)
  if err != nil {
    fmt.Println("error: get failed")
    return new(Tickers)
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println("error: get failed")
    return new(Tickers)
  }

  tickers := ToTickers(body)
  return tickers
}

func ToTickers(raw []byte) *Tickers {
  var tickers = new(Tickers)
  err := json.Unmarshal(raw, &tickers)
  if err != nil {
    fmt.Println("error: json decode failed")
    return new(Tickers)
  }
  return tickers
}

func main() {
  tickers := GetTickers("btcusdt")
  fmt.Println(tickers)
}
