# Random Crockford generator
This library generates pseudo random strings based on [Crockford's Base32](https://en.wikipedia.org/wiki/Base32#Crockford.27s_Base32) <br />


## Usage
```go
r := random_crockford.NewRandom(10)
fmt.Println(r.Next())
```

## License
MIT