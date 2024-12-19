# Snyk 브로커 클라이언트 도커 이미지 다운로드 또는 업데이트

{% hint style="info" %}
브로커 클라이언트 도커 이미지를 다운로드하기 전에, 귀하의 기기가 도커 컨테이너를 실행할 수 있는지 확인하십시오.
{% endhint %}

브로커 클라이언트 설정의 첫 번째 단계는 [Docker Hub](https://hub.docker.com/r/snyk/broker)에서 브로커 클라이언트 도커 이미지를 가져오는 것입니다. 브로커 클라이언트를 실행할 각 기기에 브로커 클라이언트 이미지를 다운로드합니다. 다운로드 후, 해당 이미지는 호스트 기기에 캐시됩니다.

{% hint style="info" %}
코드 에이전트는 브로커 클라이언트 버전 4.108.0 이상에서만 지원됩니다. 이미 실행 중인 브로커 클라이언트가 있는 경우 최신 업데이트를 가져옵니다.
{% endhint %}

**코드 에이전트 도커 이미지를 가져오려면** 다음을 실행합니다:

```
docker pull snyk/broker:<SCM_tag>
```

여기서 `SCM_tag`는 각 통합 SCM마다 특정한 태그입니다. 다음과 같습니다:

**GitHub**:

```
docker pull snyk/broker:github-com
```

**GitHub Enterprise**:

```
docker pull snyk/broker:github-enterprise
```

**GitLab**:

```
docker pull snyk/broker:gitlab
```

**Bitbucket Server/Data Center**:

```
docker pull snyk/broker:bitbucket-server
```

**Azure Repo**:

```
docker pull snyk/broker:azure-repos
```

클라이언트 브로커 이미지의 도커 이미지 다운로드 프로세스가 시작됩니다. 예를 들면:

<figure><img src="../../../../../.gitbook/assets/Client Broker - Pull image - example.png" alt="클라이언트 브로커 도커 이미지 다운로드"><figcaption><p>클라이언트 브로커 도커 이미지 다운로드</p></figcaption></figure>

다운로드가 완료되면, 브로커 클라이언트 이미지가 기기에 성공적으로 다운로드되었는지 확인할 수 있습니다.

**브로커 클라이언트 도커 이미지의 다운로드 성공 확인을 위해** 도커 목록 명령을 실행합니다:

```
docker image ls
```

다음과 유사한 출력이 표시됩니다:

```
REPOSITORY           TAG                 IMAGE ID       CREATED       SIZE
snyk/broker          github-com          e999988aa7b7   7 days ago    252MB
snyk/broker          github-enterprise   0a8b4e6f518d   7 days ago   252MB
```