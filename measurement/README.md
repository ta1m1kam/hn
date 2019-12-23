# Measurement between goroutine and without goroutine. 

This directory is to measure time that getting HackerNews between goroutine and without goroutine.

## without goroutine

```shell
❯ go run measurement/hn_api.go 5
5
3.221009秒
```

## with goroutine

```shell
❯ go run measurement/hn_api_goroutine.go 5
5
1.729498秒
```
