```
$ bundle install
$ go build -buildmode=c-shared -o lib.dylib lib.go
$ bundle exec ruby main.rb
```