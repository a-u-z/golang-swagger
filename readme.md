* 更改 swagger 後，要 swag init
* 如果安裝 swgger 卻無法 swag init
* follow the step below
```
$ cd ~/go/bin
$ ls
you should see "swag" then run
$ export PATH="/Users/XXX/go/bin:$PATH"
$ source ~/.zshrc
$ swag -v
```

* ref
  * https://matthung0807.blogspot.com/2021/08/go-gin-swagger-generate-rest-api-docs.html