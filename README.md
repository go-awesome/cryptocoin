# Get the price list for all pairs from Binance

Please Set 5 environment variables as mentioned for following

2 environment variable For Binance:

BINANCE_KEY="xxxx"

BINANCE_SECRET="xxxx"

2 environment variable For Coinbase:

COINBASE_KEY="xxxx"

COINBASE_SECRET="xxxx"

1 environment variable for Log path without ending trailing slash like below

LOG_PATH="/home/centos/go/src/cryptocoin"

# Build the package by following command

user@centos$ go build -ldflags "-s -w" main.go
user@centos$

# Once done, you can run the file as:

user@centos: ./main

# after the file has completed the process, it will store all the exchange pair and data on 2 file names:

1. binance-pair-price-with-time.txt
2. coinbase-pair-price-with-time.txt

Thanks

Happy Trading :)