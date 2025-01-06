# Run the Broker client with the code snippets display

브로커 클라이언트 이미지가 당신의 컴퓨터에 저장되면, `docker run` 명령어를 사용하여 이미지를 실행하고 브로커 클라이언트를 시작합니다.

## 코드 스니펫을 표시할 브로커 클라이언트를 설정하기 위한 Docker 실행 명령어

이 섹션에서 설명하는 브로커 클라이언트 실행을 위한 설정 명령어는 모든 SCMs에 사용되는 일반적인 명령어를 포함합니다. 일부 SCMs는 브로커 클라이언트 설정에 추가 매개변수를 필요로하며, 해당 매개변수는 이 섹션에서 표시되지만, 특정 SCM을 위한 브로커 클라이언트를 설정하는 경우 해당 SCM의 지침 또한 참조하십시오 [Snyk 브로커 설치 및 구성](../../../install-and-configure-snyk-broker/) 섹션.

다음은 결과를 웹 UI에서 표시하는 방식으로 브로커 클라이언트를 설정하는 방법에 대해 설명합니다:

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252FarerukRHyQnx9xKiziLs%252FBroker%2520-%2520Results%2520-%2520with%2520code%2520snippets.png%3Falt%3Dmedia%26token%3Dfd7841b2-87f6-48ed-ac56-f7b4323d166e&#x26;width=768&#x26;dpr=1&#x26;quality=100&#x26;sign=a520526b&#x26;sv=2" alt="코드 스니펫을 표시하는 브로커 클라이언트 실행"><figcaption><p>코드 스니펫을 표시하는 브로커 클라이언트 실행</p></figcaption></figure>

코드 스니펫을 표시하도록 브로커 클라이언트를 설정할 때, 기본 브로커 클라이언트 설정에 사용한 매개변수와 함께 다음 파일과 매개변수를 추가합니다:

사전 정의된 `accept.json` 파일을 다운로드하여 컴퓨터에 저장합니다. 이 `accept.json` 파일은 각 SCM에 맞게 사용자 정의되어 있으며, 코드 스니펫을 표시하는 데 필요한 규칙을 포함합니다.

다음 매개변수를 추가하여 `accept.json` 파일을 브로커 클라이언트에 마운트하는 설정 명령어에 넣어줍니다:

`-e ACCEPT=/<폴더_이름>/accept.json`

`-v / local/path/to/<폴더_이름>:/<폴더_이름>`

## **accept.json 파일 다운로드**

웹 UI에서 결과의 코드 스니펫을 표시하려면, 먼저 사전 정의된 `accept.json` 파일을 다운로드합니다. `accept.json` 파일을 다운로드할 때, 통합된 SCM에 맞게 사용자 정의된 파일을 선택하고 접근 가능한 위치에 저장합니다.

1. 통합 설정 페이지에서 통합 용 `accept.json` 파일을 찾아 다운로드합니다:
   * [GitHub](../../../install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/github-install-and-configure-using-docker.md)
   * [GitHub Enterprise](../../../install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-docker.md)
   * [Gitlab](../../../install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-gitlab.md)
   * [Bitbucket Server/Data Center](../../../install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/data-center.md)
   * [Azure Repos](../../../install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-azure-repos.md)
2. 다운로드한 파일이 `accept.json`으로 이름이 지어졌는지 확인합니다. 다운로드 과정에서 파일 이름이 변경되었다면 원래 이름으로 파일 이름을 변경합니다.
3. 파일을 `/private/accept.json`과 같이 안전하고 분리된 폴더에 저장합니다.

## **accept.json 파일을 사용하여 브로커 클라이언트 실행**

터미널에 다음 명령어를 입력하여 브로커 클라이언트를 시작합니다:

```
docker run --restart=always \
   -p <host_machine_port_no._mapped to>:<Broker_Client_container_port_ no.> \
   -e BROKER_TOKEN=<Broker_token> \
   -e <SCM>_TOKEN=<SCM_token> \
   -e <SCM_domain>=<my.SCM.domain.com_(without_http/s)> \  
   -e BROKER_CLIENT_URL=<http://my.broker.client:<host_machine_port_no.> \
   -e PORT=<Broker_Client_container_port_no.> \
   -e GIT_CLIENT_URL=http://<Code_Agent_container_name:Code_Agent_port_no.> \
   -e ACCEPT=/<folder_name>/accept.json \
   -v /local/path/to/<folder_name>:/<folder_name> \
   --network mySnykBrokerNetwork \
   snyk/broker:<SCM_tag>
```

여기서:

* `-- restart=always`는 Broker client 컨테이너가 종료 상태에 관계없이 항상 다시 시작되도록하는 Docker 명령어입니다.
* `-p <host_machine_port_no._mapped to>:<Broker_Client_container_port_ no.>`는 호스트 머신의 물리적인 오픈 포트를 브로커 클라이언트 컨테이너의 포트에 매핑하는 것입니다. 이 호스트 머신 및 컨테이너의 포트 번호는 동일할 필요가 없습니다. 예를 들어, `8001:8000`와 같이 호스트 머신의 포트 번호는 고유해야 합니다.
* `-e BROKER_TOKEN`은 특정 조직과 특정 통합 SCM과 관련된 브로커 토큰입니다.
* `-e <SCM_TOKEN>`은 해당 통합 SCM의 SCM 토큰입니다.
* `-e <SCM_domain>=`은 `http/https`가 없는 SCM 도메인 이름입니다. 예를 들어, `snyk.git.com`. 각 SCM마다 다음 매개변수를 사용하십시오:
  * **GitHub** - `-e <SCM_domain>` 매개변수는 필요하지 않습니다.
  * **GitHub Enterprise**: `-e GITHUB`\
    [GitHub Enterprise](../../../install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-docker.md)의 경우 다음 매개변수도 추가하십시오:\
    `-e GITHUB_API=<your.ghe.domain.com/api/v3_(without_http/s)> \`\
    `-e GITHUB_GRAPHQL=<your.ghe.domain.com/api_(without_http/s)> \`
  * **Azure Repos**: `-e AZURE_REPOS_HOST`\
    [Azure Repos](../../../install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-azure-repos.md)의 경우 다음 매개변수도 추가하십시오:\
    `-e AZURE_REPOS_ORG=<azure_repo_org_name> \`
  * **Bitbucket Server/Data Center**: `-e BITBUCKET`\
    [Bitbucket Server/Data Center](../../../install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/data-center.md)의 경우 다음 매개변수도 추가하십시오:\
    `-e BITBUCKET_API=<your.bitbucket-server.domain.com/rest/api/1.0_(without http/s)> \`
  * **GitLab**: `-e GITLAB`
* \[선택 사항] `-e BROKER_CLIENT_URL=`은 브로커 클라이언트의 호스트 머신 URL입니다. URL에는 호스트 머신의 IP 주소나 포트 번호가 포함됩니다. 예를 들어, `http://localhost:8001`와 같이 입력합니다.\
  이 매개변수는 동일한 브로커 클라이언트를 다른 Snyk 제품에서 사용 중이고, 자동 PR 확인 기능을 활성화하려는 경우에만 추가합니다. 자동 PR 확인 기능은 코드 에이전트에 대해 지원되지 않기 때문에 코드 에이전트에 대해 이 매개변수를 사용할 필요가 없습니다.
* `-e PORT`은 브로커 클라이언트 컨테이너가 외부 연결을 받는 포트 번호입니다. 기본값은 `8000`입니다. 이 포트 번호는 `-p` 매개변수의 `<Broker_Client_container_port_ no.>`와 동일해야 합니다.
* `-e GIT_CLIENT_URL=`은 실행 중인 코드 에이전트 컨테이너의 포트로의 URL입니다. URL에는 코드 에이전트 컨테이너의 이름과 포트 번호가 포함됩니다. 예를 들어, `http://code-agent:3000`와 같이 입력합니다.
* `-e ACCEPT=`은 다운로드한 `accept.json` 파일을 저장하는 폴더의 이름입니다. 예를 들어, `/private/accept.json`과 같이 입력합니다.
* `-v /local/path/to/<folder_name>:/<folder_name>`은 `accept.json` 파일을 저장하는 폴더의 경로입니다. 예를 들어, `/private:/private`와 같이 입력합니다.
* `--network`은 [Docker 브릿지 네트워크](../create-network-for-broker-client-and-code-agent-communication.md)의 이름입니다. 코드 에이전트와의 통신에 사용됩니다.
* `snyk/broker:<SCM_tag>`은 특정 통합 SCM을 위한 [브로커 클라이언트의 Docker 이미지](download-or-update-the-snyk-broker-client-docker-image.md)입니다.

브로커 클라이언트 설정이 성공적으로 완료되면, 터미널에 다음 메시지가 나타납니다:

`{ ..., "msg":"successfully established a websocket connection to the broker server", ... }`

<figure><img src="../../../../../.gitbook/assets/Broker Client - Setup success message (1).png" alt="브로커 클라이언트 설정 확인 메시지"><figcaption><p>브로커 클라이언트 설정 확인 메시지</p></figcaption></figure>

**브로커 클라이언트 컨테이너의 설정과 세부사항을 확인**하기 위해 다음을 실행합니다:

```
docker ps
```

다음과 유사한 출력이 표시됩니다:

```
CONTAINER ID   IMAGE                     COMMAND                  CREATED             STATUS             PORTS                    NAMES
1cef6e3e3290   snyk/broker:github-com    "broker --verbose"       About an hour ago   Up About an hour   0.0.0.0:8000->8000/tcp   nifty_cori  
6216e27b8d28   snyk/code-agent           "docker-entrypoint.s…"   2 hours ago         Up 2 hours         0.0.0.0:3000->3000/tcp   code-agent
```

## 코드 스니펫을 표시하는 브로커 클라이언트 실행 예시

### **통합 GitHub 서버용 브로커 클라이언트 실행**

다음 명령어는 통합된 GitHub 서버를 위해 브로커 클라이언트를 실행하는 데 사용되었습니다:

```
docker run --restart=always \
   -p 8000:8000 \
   -e BROKER_TOKEN=b1dcxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx \
   -e GITHUB_TOKEN=ghp_xxxxxxxxxxxxxxxxxxxxxx \
   -e BROKER_CLIENT_URL=http://localhost:8000 \
   -e PORT=8000 \
   -e GIT_CLIENT_URL=http://code-agent:3000 \
   -e ACCEPT=/private/accept.json \
   -v /private:/private \
   --network mySnykBrokerNetwork \
snyk/broker:github-com
```

여기서:

* `-p 8000:8000`는 호스트 컴퓨터의 포트 번호 `8000`을 브로커 클라이언트 컨테이너의 포트 번호 `8000`에 매핑합니다. 이것은 브로커 클라이언트 컨테이너와 브로커 서버 및 코드 에이전트 간 통신에 사용됩니다.
* `-e BROKER_TOKEN`은 특정 조직과 GitHub과 관련된 브로커 토큰입니다.
* `-e GITHUB_TOKEN`은 GitHub 저장소에 액세스하기 위한 GitHub 토큰입니다.
* `-e BROKER_CLIENT_URL`은 브로커 클라이언트의 호스트 컴퓨터 URL인 `http://localhost:8000`입니다.
* `-e PORT`는 브로커 클라이언트 컨테이너가 연결을 수락하는 로컬 포트인 `8000`입니다.
* `-e GIT_CLIENT_URL=http://code-agent:3000`는 실행 중인 코드 에이전트 컨테이너의 포트로의 URL입니다. URL에는 코드 에이전트 컨테이너의 이름과 포트 번호인 `code-agent:3000`이 포함됩니다.
* `-e ACCEPT=/private/accept-github.json`는 다운로드한 `accept.json` 파일을 저장하는 폴더 이름과 GitHub용 파일인 `accept.json`을 나타냅니다.
* `-v /private:/private`는 `accept-github.json` 파일을 저장하는 폴더인 `private`와 해당 폴더 이름인 `private`의 경로입니다.
* `--network`는 브로커 클라이언트와 통신에 사용되는 Docker 브리지 네트워크의 이름입니다.
* `snyk/broker:github-com`은 GitHub용 브로커 클라이언트의 Docker 이미지입니다.

### **통합 Azure Repos 서버용 브로커 클라이언트 실행**

다음 명령어는 통합된 Azure Repos 서버를 위해 브로커 클라이언트를 실행하는 데 사용되었습니다:

```
docker run --restart=always \
   -p 8001:8001 \
   -e BROKER_TOKEN=fe5bxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx \
   -e AZURE_REPOS_TOKEN=brmyxxxxxxxxxxxxxxxx \
   -e AZURE_REPOS_ORG=snyktest \
   -e AZURE_REPOS_HOST=dev.azure.com \
   -e BROKER_CLIENT_URL=http://localhost:8001 \
   -e PORT=8001 \
   -e GIT_CLIENT_URL=http://code-agent:3000 \
   -e ACCEPT=/private/accept.json \
   -v ./private:/private \
   --network mySnykBrokerNetwork \
snyk/broker:azure-repos
```

여기서:

* `-p 8001:8001`은 호스트 머신의 포트 번호 `8001`을 Broker Client 컨테이너의 포트 번호 `8001`로 매핑하며, Broker Client 컨테이너와 Broker Server 및 Code Agent 간의 통신에 사용됩니다.
* `-e BROKER_TOKEN`은 특정 Organization 및 Azure Repos와 연결된 Broker 토큰입니다.
* `-e AZURE_REPOS_TOKEN`은 Azure Repos 저장소에 접근하기 위한 Azure Repo 토큰입니다.
* `-e AZURE_REPOS_ORG`은 Azure Repos 조직 이름으로, 예: `snyktest`입니다.
* `-e AZURE_REPOS_HOST`은 Azure Repos 서버의 도메인 이름으로, 예: `dev.azure.com`입니다.
* `-e BROKER_CLIENT_URL`은 Broker 클라이언트가 실행 중인 호스트 머신의 URL로, 예: `http://localhost:8001`입니다.
* `-e PORT`는 Broker Client 컨테이너가 연결을 수락하는 로컬 포트로, 예: `8001`입니다.
* `-e GIT_CLIENT_URL=http://code-agent:3000`은 실행 중인 Code Agent 컨테이너의 포트 URL입니다. URL에는 Code Agent 컨테이너 이름 `code-agent`와 해당 포트 번호 `3000`이 포함됩니다.
* `-e ACCEPT=/private/accept.json`은 다운로드된 `accept.json` 파일을 저장하는 폴더의 이름(`private`)과 GitHub을 위한 파일 이름(`accept.json`)입니다.
* `-v ./private:/private`는 `accept.json` 파일을 저장하는 폴더(`private`)의 경로와 폴더 이름(`private`)을 지정합니다.
* `--network`는 Broker 클라이언트와 통신하기 위해 사용되는 Docker 브릿지 네트워크 이름으로, 예: `mySnykBrokerNetwork`입니다.
* `snyk/broker:azure-repos`는 Azure Repos를 위한 Broker 클라이언트의 Docker 이미지입니다.