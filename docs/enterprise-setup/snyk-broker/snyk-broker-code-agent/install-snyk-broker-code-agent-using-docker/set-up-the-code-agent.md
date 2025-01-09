# 코드 에이전트 설정

Docker를 사용하여 코드 에이전트를 설정하는 방법:

* [Code Agent 도커 이미지 다운로드 또는 업데이트](set-up-the-code-agent.md#download-or-update-the-code-agent-docker-image).
* [Code Agent 컨테이너 실행](set-up-the-code-agent.md#run-the-code-agent-container).
* 프록시 서버를 사용하는 경우, [코드 에이전트를 프록시 서버와 함께 작동하도록 설정](set-up-the-code-agent.md#set-up-the-code-agent-to-work-with-a-proxy-server).

## 코드 에이전트 Docker 이미지 다운로드 또는 업데이트

### **Code Agent Docker 이미지 다운로드**

[도커 허브](https://hub.docker.com/r/snyk/code-agent/)에서 Code Agent Docker 이미지를 가져옵니다. 최신 Docker 이미지 버전을 가져와 사용하는 것이 권장됩니다. 각 Code Agent를 실행할 기계에 Code Agent Docker 이미지를 다운로드합니다. 일반적으로 Docker 이미지는 호스트 기계에 캐싱됩니다.

Code Agent Docker 이미지를 가져오려면 터미널에 다음을 입력합니다:

```
docker pull snyk/code-agent
```

다음과 같이 Code Agent Docker 이미지에 대한 다운로드 프로세스가 시작됩니다:

<figure><img src="../../../../.gitbook/assets/Code Agent - Pull docker image - New.png" alt="Code Agent Docker 이미지 다운로드 프로세스"><figcaption><p>Code Agent Docker 이미지 다운로드 프로세스</p></figcaption></figure>

### **Code Agent Docker 이미지 업데이트**

Code Agent Docker 이미지를 다시 가져옵니다. `latest` 태그를 사용하는 경우 이미지는 자동으로 업데이트됩니다. 그렇지 않은 경우 새 이미지 태그를 제공합니다:

```
docker pull snyk/code-agent:<image_tag>
```

이전 Code Agent 컨테이너를 제거하거나 중지합니다.

다음 섹션의 단계에 따라 새 Code Agent 컨테이너를 시작합니다.

## 코드 에이전트 컨테이너 실행

Code Agent 이미지를 컴퓨터에 저장하면, 터미널에 다음 명령을 입력하여 Snyk Broker Code Agent 이미지를 기반으로 하는 컨테이너를 실행합니다:

```
docker run --name <container_name> \
-p <호스트_기계_포트_no._mapped to>:<Code_Agent_container_port_no.> \
-e PORT=<Code_Agent_container_port_no.> -e SNYK_TOKEN=<Snyk_API_token> --network <network_name> \
snyk/code-agent:<image_tag>
```

여기서:

* `--name <container_name>`은 새로운 Code Agent 컨테이너의 이름입니다. 이 이름은 다음에 실행할 Broker 클라이언트의 `GIT_CLIENT_URL` 매개 변수를 정의하는 데 사용됩니다. 예: `code-agent`.
* `-p <호스트_기계_포트_no._mapped to>:<Code_Agent_container_port_no.>`은 호스트 기계의 물리적 오픈 포트를 Code Agent 컨테이너의 포트로 매핑하는 것입니다. 호스트 기계 및 컨테이너의 포트 번호는 동일할 필요는 없습니다.\
  예: `3001:3000`.\
  호스트 기계의 포트 번호는 고유해야 합니다.
* `-e PORT`은 Code Agent 컨테이너의 외부 연결을 받는 포트입니다. 기본값은 `3000`입니다. 이 포트 번호는 이전 매개 변수인 `-p`에서 `<Code_Agent_container_port_no.>`와 동일해야 합니다.
* `-e SNYK_TOKEN`은 Snyk 웹 UI의 **계정 설정** 페이지에 표시된 대로 Snyk API 토큰입니다.
* `--network`은 이전에 생성된 [도커 브리지 네트워크](create-network-for-broker-client-and-code-agent-communication.md)의 이름입니다. 예: `mySnykBrokerNetwork`.
* `snyk/code-agent:<image_tag>`은 Code Agent 컨테이너의 Docker 이미지입니다. `latest`를 사용하지 않는 경우 태그를 지정합니다.

코드 에이전트 설정이 성공적으로 완료되면, 터미널에 다음 메시지가 나타납니다:

`{ ..., "msg":"Application started", ... }`

<figure><img src="../../../../.gitbook/assets/Code Agent - Exmaple - success.png" alt="Code Agent 설정 완료 시 성공 메시지"><figcaption><p>Code Agent 설정이 완료되었을 때 나타나는 성공 메시지</p></figcaption></figure>

### 코드 에이전트 컨테이너 설정 및 세부 사항 확인

다음을 실행합니다:

```
docker ps
```

다음과 비슷한 출력이 표시됩니다:

```
CONTAINER ID   IMAGE            COMMAND                 CREATED      STATUS      PORTS                    NAMES
eebd7d4f0568   snyk/code-agent "docker-entrypoint.s…"   9 days ago   Up 9 days   0.0.0.0:3000->3000/tcp   code-agent
```

### 코드 에이전트 실행 예시

다음 예시에서는 터미널에 다음 명령을 입력하여 Code Agent 컨테이너를 시작했습니다:

```
docker run --name code-agent \
-p 3000:3000 \
-e PORT=3000 -e SNYK_TOKEN=fa7fxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx --network mySnykBrokerNetwork \
snyk/code-agent
```

여기서:

* `--name`은 새 Code Agent 컨테이너의 이름인 `code-agent`입니다.
* `-p` - 호스트 기계의 포트 `3000`이 Code Agent 컨테이너의 포트 `3000`으로 매핑되어 있습니다.
* `-e PORT`는 외부 연결을 받는 Code Agent 컨테이너의 포트인 `3000`입니다.
* `-e SNYK_TOKEN`은 Snyk API 토큰인 `fa7f….`입니다.
* `--network`은 클라이언트 브로커와의 통신에 사용되는 도커 브리지 네트워크의 이름인 `mySnykBrokerNetwork`입니다.
* `snyk/code-agent`는 Code Agent 컨테이너의 Docker 이미지입니다.

### **내부 인증서가 있는 Git 인스턴스에 연결**

기본적으로 Code Agent는 Git 인스턴스에 HTTPS 연결을 설정합니다. Git 인스턴스가 내부 인증서를 제공하는 경우 (사용자의 자체 CA로 서명된), CA 인증서를 Code Agent에 제공할 수 있습니다.

예를 들어, CA 인증서가 `./private/ca.cert.pem`에 있는 경우, 폴더를 마운트하고 `CA_CERT` 환경 변수를 사용하여 Docker 컨테이너에 제공합니다:

```
docker run --name code-agent \
-p 3000:3000 \
-e PORT=3000 -e SNYK_TOKEN=fa7fxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx --network mySnykBrokerNetwork \
-e CA_CERT=/private/ca.cert.pem \
-v /local/path/to/private:/private \
snyk/code-agent
```

## 프록시 서버와 함께 작동하도록 코드 에이전트 설정

프록시를 사용하는 인프라에서 코드 에이전트를 사용하려면 코드 에이전트의 `docker run` 명령에 다음 환경 변수를 추가합니다:

```
-e HTTP_PROXY=http://my.proxy.address:<port_no.> \
-e HTTPS_PROXY=http://my.proxy.address:<port_no.>
```

프록시가 사용자 이름 및 암호 인증을 요구하는 경우 다음 추가 환경 변수를 추가합니다:

```
-e PROXY_AUTH=userID:userPass
```

또한, Broker Client 구성 요소에 이러한 환경 변수를 추가하고 Code Agent를 우회하는 명령을 추가합니다.

프록시를 사용하여 Docker 컨테이너를 사용하는 자세한 정보는 [도커가 프록시 서버를 사용하도록 구성](https://docs.docker.com/network/proxy/)를 참조하십시오.

### **사용자 지정 인증서**

프록시로 보호된 Code Agent를 사용하려면 (HTTPS), Code Agent의 `docker run` 명령에 다음 환경 변수를 추가합니다:

```
-e HTTP_PROXY=http://my.proxy.address:<port_no.> \
-e HTTPS_PROXY=https://my.proxy.address:<port_no.>
```

다음 단계는 실행 중인 Code Agent 버전에 따라 다릅니다. `latest` 태그를 사용하는 경우, 가장 가까운 버전화된 이미지를 찾으려면:

* 로컬 이미지의 `digest`를 [Docker Hub Code Agent Tags](https://hub.docker.com/r/snyk/code-agent/tags)와 비교합니다: `docker images snyk/code-agent --digest`
* 로컬 이미지가 빌드된 시점보다 이전에 릴리스된 `x.y.z` 형태의 다음 이미지 태그를 찾습니다.

### **Version `1.18.0` 이후**

사용자 정의 Certificate Authority를 신뢰하려면 다음 중 하나가 있어야 합니다:

* 단일 Certificate Authority (`PEM`으로 인코딩)
* 여러 Certificate Authority를 포함하는 디렉토리 (`PEM`으로 인코딩)

단일 인증서를 신뢰하려면 Code Agent의 `docker run` 명령에 다음 인수를 추가합니다:

```
-v local/path/to/ca.pem:/etc/certs/ca.pem \
-e GIT_SSL_CAINFO='/etc/certs/ca.pem'
```

인증서 디렉토리를 신뢰하려면 Code Agent의 `docker run` 명령에 다음 인수를 추가합니다:

```
-v local/path/to/certdirectory:/etc/certs
-e GIT_SSL_CAPATH='/etc/certs'
```

### **Version `1.16.0` - `1.17.0`**

앞의 단계를 따르고 다음 인수를 Code Agent의 `docker run` 명령에 추가합니다:

```
-e CODE_AGENT_GIT_CLI=true
```

### **Version `1.15.2` 및 이전**

Code Agent `1.15.2` 이하에서는 사용자 정의 Certificate Authorities의 신뢰를 지원하지 않으며 모든 인증서를 신뢰해야 합니다.

Code Agent의 `docker run` 명령에 다음 환경 변수를 추가합니다:

```
-e NODE_TLS_REJECT_UNAUTHORIZED=0
```
