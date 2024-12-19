# Docker를 사용한 브로커 클라이언트용 HTTPS

브로커 클라이언트는 기본적으로 HTTP 서버를 실행합니다. 로컬 연결을 위해 HTTPS 서버를 실행하도록 구성할 수 있습니다. 이를 위해서는 실행 시 Docker 컨테이너에 SSL 인증서와 개인 키를 제공해야 합니다.

예를 들어, 인증서 파일들이 로컬에서 `./private/broker.crt` 및 `./private/broker.key` 경로에 있다면, 다음과 같이 폴더를 마운트하고 `HTTPS_CERT` 및 `HTTPS_KEY` 환경 변수를 사용하여 Docker 컨테이너에 이 파일들을 제공합니다:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e GITHUB_TOKEN=secret-github-token \
           -e PORT=8000 \
           -e HTTPS_CERT=/private/broker.crt \
           -e HTTPS_KEY=/private/broker.key \
           -e BROKER_CLIENT_URL=https://my.broker.client:8000 \
           -v /local/path/to/private:/private \
       snyk/broker:github-com
```

`BROKER_CLIENT_URL`이 HTTPS 프로토콜로 변경되었음에 주목하세요.