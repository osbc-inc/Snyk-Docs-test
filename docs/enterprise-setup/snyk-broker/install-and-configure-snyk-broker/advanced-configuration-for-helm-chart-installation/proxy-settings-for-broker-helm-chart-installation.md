# Broker Helm 차트 설치를 위한 프록시 설정

프록시 뒤에 Helm 차트를 사용하려면 `httpProxy` 및 `httpsProxy` 값을 설정하세요.

```
--set httpsProxy=<PROXY_URL>
```