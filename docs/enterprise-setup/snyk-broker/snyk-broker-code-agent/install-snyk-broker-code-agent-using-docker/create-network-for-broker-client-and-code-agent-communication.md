# Broker 클라이언트와 Code Agent 간 통신을 위한 네트워크 생성

Snyk Broker를 Code Agent와 함께 실행하려면 그들 사이에 네트워크 연결을 설정해야 합니다. 이를 위해 여기에 설명된 대로 Docker 브리지 네트워크를 만들 수 있습니다.

{% hint style="info" %}
Docker 브리지 네트워크에 대한 자세한 정보는 [Docker 문서](https://docs.docker.com/network/bridge/)를 참조하세요.

Broker 클라이언트와 Code Agent 간의 이 네트워크 연결을 설정하려면 Ngrok와 같은 다른 방법과 도구도 사용할 수 있습니다.
{% endhint %}

**Docker 브리지 네트워크를 생성하려면** 다음을 실행합니다:

```bash
docker network create <network_name>
```

여기서 `network_name`은 Broker 클라이언트와 Code Agent 사이의 새 네트워크에 대한 이름입니다. 예를 들어:

```bash
docker network create mySnykBrokerNetwork
```

여기서 `mySnykBrokerNetwork`가 새 네트워크의 이름입니다.

**네트워크 생성을 확인하려면** 다음을 실행합니다:

```bash
docker network ls
```

아래와 유사한 출력이 표시됩니다:

```bash
NETWORK ID     NAME                     DRIVER    SCOPE
cb352a803eb0   mySnykBrokerNetwork   bridge    local
```