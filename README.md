# google-auth-sample
Google OAuth2認証を用いたログインサービス

## todo
### application
- [x] oauth2
- [x] session
    - [ ] session save redis
- [ ] datastore
    - [ ] sqlite3
    - [ ] mysql
        - http://doc.gorm.io/database.html
- [ ] migration
    - use GORM
- [ ] package manager

#### Fix
- [ ] oauth2callback時の`invalid state parameter`

### development environment setup
- [x] docker image
- [x] docker-compose setup
    - [x] mysql
    - [x] redis
    - [x] web

## Getting started

```
$ docker-compose build
$ docker-compose up -d
```

## 参考
### Docker
- https://blog.golang.org/docker
- https://github.com/hiromaily/go-gin-wrapper/blob/master/docker-compose.yml
- https://shamaton.orz.hm/blog/archives/310
- https://slideship.com/users/@deadcheat/presentations/2018/01/KWpsfgXdfU2uNGP7tDNVYX/
- http://unknownplace.org/archives/golang-development-enviroment-with-docker.html

### Google OAuth2
- https://github.com/GoogleCloudPlatform/golang-samples/tree/master/getting-started/bookshelf
- https://murashun.jp/blog/20150920-01.html
- https://godoc.org/golang.org/x/oauth2#Config.AuthCodeURL
- https://medium.com/eureka-engineering/go%E8%A8%80%E8%AA%9E%E3%81%A7%E3%81%AE%E6%B1%BA%E6%B8%88%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0%E3%81%A8%E3%83%9E%E3%82%A4%E3%82%AF%E3%83%AD%E3%82%B5%E3%83%BC%E3%83%93%E3%82%B9%E5%8C%96%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6-d80dd4e6f684