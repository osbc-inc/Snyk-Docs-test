# Docker와 함께 프록시 지원

프록시 구성에 대한 자세한 내용은 [도커를 프록시 서버 사용하도록 구성](https://docs.docker.com/network/proxy/)을 참조하십시오.

```
 -e HTTP_PROXY=http://my.proxy.address:8080
 -e HTTPS_PROXY=http://my.proxy.address:8080
 -e NO_PROXY=*.test.example.com,.example2.com,127.0.0.0/8
```

프록시가 사용자 이름과 암호 인증을 필요로 하는 경우, 다음 추가 환경 변수를 제공하십시오:

```
-e PROXY_AUTH=userID:userPass
```