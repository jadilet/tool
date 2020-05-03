## About
  * This tool computes the md5 hash of http request response.
  * Concurrently computes the md5 hash.
  * The default value of concurrent runnig goroutines is ten(10).

## Build
  ```
  go build main.go
  ```

## Run
```
./main google.com yandex.com

Output:

http://google.com 2e907bf6a12b97bb43c44bbe360d473c
http://yandex.com c1c40a19f7e20912eae6fdc05230de81

# Set the flag for to limit the concurrent running goroutines

./main -parallel 2 google.com yandex.com

Output:

http://google.com bd1a5b3f83bf2e1c14287cf6136a3688
http://yandex.com 9c4ef66cb22885f5cac515785ab30d48
```
