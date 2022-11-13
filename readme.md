* 更改 swagger 後，要 swag init
* 如果安裝 swgger 卻無法 swag init
* follow the step below
```
$ cd ~/go/bin
$ ls
you should see "swag" then run
$ export PATH="/Users/XXX/go/bin:$PATH"
$ export PATH=$(go env GOPATH)/bin:$PATH
$ source ~/.zshrc
$ source .bash_profile
$ swag -v
```

* ref
  * https://matthung0807.blogspot.com/2021/08/go-gin-swagger-generate-rest-api-docs.html
  * https://github.com/swaggo/swag/issues/197