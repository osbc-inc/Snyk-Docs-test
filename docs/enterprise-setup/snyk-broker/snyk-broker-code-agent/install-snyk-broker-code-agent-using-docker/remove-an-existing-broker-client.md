# 기존 브로커 클라이언트 제거

{% hint style="info" %}
이 섹션은 이미 실행 중인 브로커 클라이언트가 있는 경우에만 적용됩니다. 아직 브로커 클라이언트가 없는 경우에는 [브로커 클라이언트 및 코드 에이전트 간 통신을 위한 네트워크 생성](create-network-for-broker-client-and-code-agent-communication.md)로 이동하십시오.
{% endhint %}

같은 조직과 통합을 위해 이미 실행 중인 브로커 클라이언트가 있는 경우, 새로운 브로커 클라이언트를 코드 에이전트를 위해 설정하기 전에 중지하고 제거해야 합니다. 브로커 클라이언트는 코드 에이전트와 통신해야 하며, 이 통신은 브로커 클라이언트 컨테이너의 설정 명령으로 구성되므로 코드 에이전트 운영을 위해 기존 브로커 클라이언트 컨테이너를 사용할 수 없습니다.

{% hint style="info" %}
코드 에이전트를 위해 설정한 새 브로커 클라이언트는 또한 동일한 Snyk 조직과 동일한 SCM에 대한 {{Snyk 오픈 소스}} 및 {{Snyk IaC}}의 브로커 배포를 처리할 수 있습니다. [{{Snyk Container}}](../../snyk-broker-container-registry-agent/)의 경우 추가적인 브로커 클라이언트를 설정해야 합니다.
{% endhint %}

**실행 중인 브로커 클라이언트가 있는지 확인하려면** Docker Process Status 명령을 사용하십시오. 이 명령은 기계에 실행 중인 모든 컨테이너 목록을 제공합니다.

{% hint style="info" %}
실행 중인 브로커 클라이언트가 있는 경우, 이 명령은 실행 중인 브로커 클라이언트를 제거하기 위해 필요한 세부 정보도 제공합니다. 실행 중인 브로커 클라이언트를 찾으면, 그것이 코드 에이전트를 위해 동일한 Snyk 조직과 동일한 SCM으로 구성되어 있는 경우에만 제거해야 합니다.
{% endhint %}

다음을 실행하십시오:

```
docker ps
```

만약 실행 중인 브로커 클라이언트가 있다면, 다음과 유사한 출력을 얻게 됩니다:

```
CONTAINER ID   IMAGE                    COMMAND                  CREATED        STATUS        PORTS                    NAMES
6eab097879cc   snyk/broker:github-com   "broker --verbose"       18 hours ago   Up 18 hours   0.0.0.0:8001->8000/tcp   suspicious_banzai
```

**브로커 클라이언트 컨테이너를 제거하려면**, 다음을 실행하십시오:

```
docker rm -f <브로커_클라이언트_컨테이너_ID_또는_이름)
```

여기서 `container_ID_or_name`은 실행 중인 브로커 클라이언트의 컨테이너 ID 또는 이름입니다. 이러한 세부 정보는 `docker ps` 명령을 사용하여 얻을 수 있으며, 예를 들어:

<figure><img src="../../../../.gitbook/assets/Broker Client - removing.png" alt="docker ps 명령 예시"><figcaption><p>docker ps 명령 예시</p></figcaption></figure>