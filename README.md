# trade
this is simple golang implementation using calculate max difference from remote file

## 1 installing dependencies
```
make install
```

## 2 testing the code
```
make test
```

## 3 build
make build
./trade


## How to run solver
### 1 calculate max profit
```
go run main.go max-trade
```

and the response console will be like this:
```
{"max_profit": 299994700,"buy_hour": 11789,"sell_hour": 19090,}
```

### 2 unique string

```
go run main.go unique-string
```

and the response console will be like this:
```
{"first_occurence": "sfgxclryzidpuvejaqbtwmhkno","smallest_lexicographical_order": "abcdefghijklmnopqrstuvwxyz"}
```

### 3 find H(n) of six

```
go run main.go factor-six
```

type number to find factor six

```
Enter text: 128
```

and the response console will be like this :
```
{"mathched_factor_count":19}
```

*note : need to improve on step 3
