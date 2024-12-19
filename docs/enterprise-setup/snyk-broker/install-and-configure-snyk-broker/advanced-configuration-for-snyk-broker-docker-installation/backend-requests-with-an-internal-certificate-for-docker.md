# 도커를 이용한 내부 인증서를 사용하는 백엔드 요청

기본적으로 브로커 클라이언트는 백엔드 시스템(GitHub, BitBucket, Jira 등)으로의 HTTPS 연결을 설정합니다. 만약 백엔드 시스템이 내부 인증서(CA인증서로 서명된)를 제공한다면, 해당 CA인증서를 브로커 클라이언트에 제공할 수 있습니다.

예를 들어, CA 인증서가 `./private/ca.cert.pem`에 있는 경우, 다음과 같이 폴더를 마운트하고 `NODE_EXTRA_CA_CERT` 환경 변수를 사용하여 Docker 컨테이너에 제공합니다. 다음은 Bitbucket의 예시입니다:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e BITBUCKET_USERNAME=username \
           -e BITBUCKET_PASSWORD=password \
           -e BITBUCKET=your.bitbucket-server.domain.com \
           -e BITBUCKET_API=your.bitbucket-server.domain.com/rest/api/1.0 \
           -e PORT=8000 \
           -e NODE_EXTRA_CA_CERTS=/private/ca.cert.pem \
           -v /local/path/to/private:/private \
       snyk/broker:bitbucket-server
```

[브로커 버전 4.166.0 (2023-10-10)부터](https://github.com/snyk/broker/releases/tag/v4.166.0), 사용자 지정 CA인증서 지시사항은 `NODE_EXTRA_CA_CERTS`로 변경되었으며, 사용자 정의 CA를 사용하려면 이를 설정해야 합니다. 이를 위해 `CA_CERT` 환경 변수는 더 이상 사용되지 않습니다.

백엔드 시스템으로의 모든 요청에 대해 기본 CA 인증서 목록을 완전히 대체하므로 백엔드 시스템에서 사용하는 인증서에 필요한 완전한 체인이어야 합니다.

이 인증서는 `PEM` 형식이어야하며 `DER`는 지원되지 않습니다. 지원되는 인증서 유형은 다음과 같습니다:

* `TRUSTED CERTIFICATE`
* `X509 CERTIFICATE`
* `CERTIFICATE`

다음은 예시입니다.

```
-----BEGIN CERTIFICATE-----
<base64-encoded certificate>
-----END CERTIFICATE----
-----BEGIN CERTIFICATE-----
<base64-encoded certificate>
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
<base64-encoded certificate>
-----END CERTIFICATE-----
```  