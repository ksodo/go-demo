package main

//import "fmt"
import "strconv"
import "testing"

// demo test failure
func TestGetTickersFail(t *testing.T) {
  got := GetTickers("invalid")

  var expCode int32 = 0
  if expCode != got.Code {
    t.Errorf("0 != %d", got.Code)
  }

  expMessage := "Operation successful"
  if expMessage != got.Message {
    t.Errorf("%s != %s", expMessage, got.Message)
  }
}

// demo test failure
func TestGetTickersFailData(t *testing.T) {
  got := GetTickers("invalid")

  var ticker *TickersDataTicker = &got.Data.Ticker
  ticker.Buy = "0"
  gotBuy, err := strconv.ParseFloat(ticker.Buy, 64)
  if err != nil {
    t.Errorf("failed to convert %s to float", ticker.Buy)
  }
  if gotBuy <= 0.0 {
    t.Errorf("expected: %f > 0", gotBuy)
  }
}

// demo test pass
func TestGetTickersPass(t *testing.T) {
  got := GetTickers("btcusdt")

  var ticker *TickersDataTicker = &got.Data.Ticker
  gotBuy, err := strconv.ParseFloat(ticker.Buy, 64)
  if err != nil {
    t.Errorf("failed to convert %s to float", ticker.Buy)
  }
  if gotBuy <= 0.0 {
    t.Errorf("expected: %f > 0", gotBuy)
  }
}

// store result in package level variable to avoid compiler optimization
// compiler may remove function call if result is unused
var resultBenchmarkTickersAlloc *Tickers

func BenchmarkTickersAlloc(b *testing.B) {
  var t *Tickers
  for i := 0; i < b.N; i++ {
    t = new(Tickers)
  }
  resultBenchmarkGetTickers = t
}

var resultBenchmarkGetTickers *Tickers

func BenchmarkGetTickers(b *testing.B) {
  var got *Tickers
  for i := 0; i < b.N; i++ {
    got = GetTickers("btcusdt")
  }
  resultBenchmarkGetTickers = got
}
