# Helm을 사용하여 문제 해결 및 고유 인증서 제공을 위한 매개변수

SSL 검사 문제를 해결하려면 `tlsRejectUnauthorized` 매개변수를 `disable`로 설정할 수 있습니다.

```
--set tlsRejectUnauthorized=disable
```

고유 인증 기관을 신뢰하려면 `caCert` 매개변수에 인증서 파일 이름을 전달할 수 있습니다. 파일은 Helm 차트 디렉토리 내에 있어야 합니다.

```
--set caCert=<CERT_NAME>
```

대신에 인증서 내용을 `caCertFile` 매개변수에 제공할 수도 있습니다.

```
--set caCertFile="-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----"
```

CA 신뢰 문제를 해결하려면 `disableCaCertTrust` 매개변수를 `true`로 설정할 수 있습니다.

```
--set disableCaCertTrust=true
```

Broker를 HTTPS 서버로 실행하려면 파일을 `httpsCert` 및 `httpsKey` 매개변수에 전달할 수 있습니다. 파일은 Helm 차트 디렉토리 내에 있어야 합니다.

```
--set httpsCert=<CERT_NAME> --set httpsKey=<CERT_KEY>
```

고유 인증서를 사용하는 자세한 정보는 [Docker를 위한 내부 인증서로 백엔드 요청](../advanced-configuration-for-snyk-broker-docker-installation/backend-requests-with-an-internal-certificate-for-docker.md) 및 [Docker용 Broker 클라이언트용 HTTPS](../advanced-configuration-for-snyk-broker-docker-installation/https-for-broker-client-with-docker.md)를 참조하십시오.