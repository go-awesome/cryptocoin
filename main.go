package main

/**
** Author: Ravindra Sharma
**/

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/adshao/go-binance/v2"
	"github.com/fabioberger/coinbase-go"
)

var (
	binanceApiKey     = os.Getenv("BINANCE_KEY")
	binanceSecretKey  = os.Getenv("BINANCE_SECRET")
	coinbaseApiKey    = os.Getenv("COINBASE_KEY")
	coinbaseSecretKey = os.Getenv("COINBASE_SECRET")
	logPath           = os.Getenv("LOG_PATH")
)

func main() {
	// Create a client connection for binance
	client := binance.NewClient(binanceApiKey, binanceSecretKey)

	// fetching all the price and trading pairs
	prices, err := client.NewListPricesService().Do(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// create a buffer so that we don't open and write or append for each price.
		var bbuffer bytes.Buffer
		for _, p := range prices {
			bbuffer.WriteString("Symbol: " + p.Symbol + " Price: " + p.Price + "\r\n")
		}
		// write all binance data to a file
		writeLog(bbuffer.String(), "binance-pair-price.txt")
		bbuffer.Reset() // buffer reset as no longer required
	}

	// Below is for the Coinbase
	// Creating API client
	cb := coinbase.ApiKeyClient(coinbaseApiKey, coinbaseSecretKey)
	// getting Exchange Rates
	exchanges, err := cb.GetExchangeRates()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// Creating buffer same as above for binance
		var cbuffer bytes.Buffer
		for i, p := range exchanges {
			splits := strings.Split(i, "_to_")
			if len(splits) == 2 {
				// writing in buffer as string
				cbuffer.WriteString("Symbol: " + splits[0] + "/" + splits[1] + " Price: " + p + "\r\n")
			}
		}
		// writing all coinbase exchange price to a file
		writeLog(cbuffer.String(), "coinbase-pair-price.txt")
		cbuffer.Reset() // buffer reset as no longer required
	}

}

// a function to save contnet inside a log file
func writeLog(message string, filename string) {
	if _, err := os.Stat(logPath + "/" + filename); err == nil {
		f, _ := os.OpenFile(logPath+"/"+filename, os.O_APPEND|os.O_WRONLY, 0644)
		defer f.Close()
		_, _ = f.WriteString(message + "\r\n")
	} else if os.IsNotExist(err) {
		_ = ioutil.WriteFile(logPath+"/"+filename, []byte(message+"\r\n"), 0644)
	} else {
		f, _ := os.OpenFile(logPath+"/log.txt", os.O_APPEND|os.O_WRONLY, 0644)
		defer f.Close()
		_, _ = f.WriteString(filename + " - " + message + "\r\n")
	}
}
