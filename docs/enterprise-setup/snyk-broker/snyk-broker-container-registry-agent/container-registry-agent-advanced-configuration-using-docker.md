# Docker를 사용한 컨테이너 레지스트리 에이전트 고급 구성

Docker를 사용하여 브로커 컨테이너 레지스트리 에이전트를 설치하는 방법에 대한 지침은 [Snyk 브로커 - 컨테이너 레지스트리 에이전트](./)를 참조하십시오. Helm 방법을 사용한 설치 지침은 [Helm을 사용한 컨테이너 레지스트리 에이전트용 브로커 설치](install-broker-for-container-registry-agent-using-helm.md)를 참조하십시오.

## 컨테이너 레지스트리 에이전트 서버 **HTTPS 구성**

컨테이너 레지스트리 에이전트(CRA)는 기본적으로 HTTP 서버에서 실행됩니다. CRA를 로컬 연결용 HTTPS 서버로 구성할 수 있습니다. 이를 위해 실행 시 Docker 컨테이너에 SSL 인증서와 개인 키를 제공해야 합니다.

예를 들어, 만약 인증서 파일이 `./private/container-registry-agent.crt` 및 `./private/container-registry-agent.key` 경로에 있다면, 해당 파일들을 폴더를 마운트하고 `HTTPS_CERT` 및 `HTTPS_KEY` 환경 변수를 사용하여 Docker 컨테이너에 제공하세요:

```
docker run --restart=always \
       -p 8081:8081 \
       -e SNYK_PORT=8081 \
       -e HTTPS_CERT=/private/container-registry-agent.crt \
       -e HTTPS_KEY=/private/container-registry-agent.key \
       snyk/container-registry-agent:latest
```

## **내부 인증서가 있는 컨테이너 레지스트리 및 브로커 클라이언트**

기본적으로 컨테이너 레지스트리 에이전트는 컨테이너 레지스트리 및 브로커 클라이언트에 HTTPS 연결을 설정합니다. 만약 컨테이너 레지스트리 또는 브로커 클라이언트가 내부 인증서(자체 CA에 의해 서명된)를 사용하고 있다면, 해당 CA 인증서를 컨테이너 레지스트리 에이전트에 제공할 수 있습니다. 예를 들어, 만약 CA 인증서가 `./private/ca.cert.pem` 경로에 있다면, 해당 파일을 폴더를 마운트하고 `NODE_EXTRA_CA_CERTS` 환경 변수를 사용하여 Docker 컨테이너에 제공하세요:

<pre><code>docker run --restart=always \
       -p 8081:8081 \
<strong>       -e SNYK_PORT=8081 \
</strong>       -e NODE_EXTRA_CA_CERTS=/private/ca.cert.pem \
       snyk/container-registry-agent:latest
</code></pre>
