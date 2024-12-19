# 코드 스니펫 표시 없이 Broker 클라이언트 실행

브로커 클라이언트 이미지가 사용자의 기기에 저장된 후, `docker run` 명령어를 사용하여 이미지를 실행하고 브로커 클라이언트를 시작합니다.

## 코드 스니펫 표시 없이 Broker 클라이언트 설정을 위한 Docker run 명령어

이 섹션에서 설명된 브로커 클라이언트 실행을 위한 설정 명령어는 모든 SCM에서 사용되는 일반 명령어를 포함하고 있습니다. 일부 SCM은 Broker 클라이언트 설정을 위해 추가 매개변수를 필요로 할 수 있으며, 이러한 매개변수는 이 섹션에서 표시되어 있지만 특정 SCM을 위한 Broker 클라이언트를 설정 중일 때 해당 SCM의 설명도 확인하세요 [Install and configure Snyk Broker](../../../install-and-configure-snyk-broker/) 섹션에서.

다음은 웹 UI의 {{Snyk Code}} 결과의 코드 스니펫을 표시하지 않도록 브로커 클라이언트를 설정하는 방법을 설명합니다:

<figure><img src="../../../../../.gitbook/assets/Broker - Results - without code snippets (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (4) (1).png" alt="코드 스니펫 표시 없이 Broker 클라이언트 실행"><figcaption><p>코드 스니펫 표시 없이 Broker 클라이언트 실행</p></figcaption></figure>

**브로커 클라이언트 컨테이너 실행**을 위해 터미널에 다음 명령어를 입력하여 Snyk 브로커 클라이언트를 실행합니다:

```
docker run --restart=always \
   -p <호스트 머신 포트 번호._매핑된>:<브로커 클라이언트 컨테이너 포트 번호> \
   -e BROKER_TOKEN=<브로커 토큰> \
   -e <SCM>_TOKEN=<SCM 토큰> \
   -e <SCM_domain>=<my.SCM.domain.com_(http/s 미포함)> \  
   -e BROKER_CLIENT_URL=<http://my.broker.client:<호스트 머신 포트 번호> \
   -e PORT=<브로커 클라이언트 컨테이너 포트 번호> \
   -e GIT_CLIENT_URL=http://<Code_Agent_컨테이너_이름:Code_Agent_포트_번호> \
   --network mySnykBrokerNetwork \
   snyk/broker:<SCM_태그>
```

여기서:

* `-- restart=always`는 브로커 클라이언트 컨테이너가 종료 상태와 관계 없이 항상 다시 시작됨을 결정하는 Docker 명령어입니다.
* `-p <호스트 머신 포트 번호._매핑된>:<브로커 클라이언트 컨테이너 포트 번호>`는 호스트 머신의 물리적인 오픈 포트를 브로커 클라이언트 컨테이너의 포트에 매핑하는 것입니다. 호스트 머신과 컨테이너의 이 포트 번호는 같을 필요는 없으며, 예를 들어 `8001:8000` 과 같이 사용될 수 있습니다.\
  호스트 머신의 포트 번호는 고유해야 합니다.
* `-e BROKER_TOKEN`은 특정 조직과 특정 통합 SCM과 연결된 브로커 토큰입니다.
* `-e <SCM_TOKEN>`은 특정 통합 SCM의 SCM 토큰입니다.
* `-e <SCM_domain>=`은 `http/https`를 제외한 SCM 도메인명입니다. 예를 들어 `snyk.git.com`입니다. 각 SCM에 대해 해당 SCM을 위한 매개변수를 사용합니다:
    * **GitHub** - `-e <SCM_domain>` 매개변수는 필요하지 않습니다.
    * **GitHub Enterprise**: `-e GITHUB`\
      [GitHub Enterprise](../../../install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-docker.md)를 위해 다음 매개변수를 추가하세요:\
      `-e GITHUB_API=<your.ghe.domain.com/api/v3_(http/s 미포함)> \`\
      `-e GITHUB_GRAPHQL=<your.ghe.domain.com/api_(http/s 미포함)> \`
    * **Azure Repos**: `-e AZURE_REPOS_HOST`\
      [Azure Repos](../../../install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-azure-repos.md)를 위해 다음 매개변수를 추가하세요:\
      `-e AZURE_REPOS_ORG=<azure_repo_org_name> \`
    * **Bitbucket Server/Data Center**: `-e BITBUCKET`\
      [Bitbucket Server/Data Center](../../../install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/data-center.md)를 위해 다음 매개변수를 추가하세요:\
      `-e BITBUCKET_API=<your.bitbucket-server.domain.com/rest/api/1.0_(http/s 미포함)> \`
    * **GitLab**: `-e GITLAB`
* \[선택 사항] `-e BROKER_CLIENT_URL=`은 브로커 클라이언트의 호스트 머신 URL입니다. URL에는 호스트 머신의 IP 주소 또는 포트 번호를 포함할 수 있습니다. 예를 들어, `http://localhost:8000`와 같이 사용될 수 있습니다.\
  동일한 브로커 클라이언트를 다른 Snyk 제품에 사용하고 Automatic PR Checks 기능을 활성화하려는 경우에만 이 매개변수를 추가하세요. Automatic PR Checks 기능은 Code Agent에 지원되지 않으므로 Code Agent에 대한 이 매개변수는 사용할 필요가 없습니다.
* `-e PORT`은 외부 연결을 수락하는 브로커 클라이언트 컨테이너의 포트 번호입니다. 기본값은 `8000`입니다. 이 포트 번호는 `-p` 이전 매개변수의 `<브로커 클라이언트 컨테이너 포트 번호>`와 동일해야 합니다.
* `-e GIT_CLIENT_URL`은 실행 중인 Code Agent 컨테이너의 포트에 대한 URL입니다. URL에는 Code Agent 컨테이너의 이름과 포트 번호를 포함해야 합니다. 예를 들어, `http://code-agent:3000`과 같이 사용될 수 있습니다.
* `--network`은 [Docker 브릿지 네트워크](../create-network-for-broker-client-and-code-agent-communication.md)의 이름이며, Code Agent와의 통신에 사용됩니다.
* `snyk/broker:<SCM_tag>`은 특정 통합 SCM을 위한 [브로커 클라이언트의 Docker 이미지](download-or-update-the-snyk-broker-client-docker-image.md)입니다.

브로커 클라이언트 설정이 성공적으로 완료되면, 다음 메시지가 터미널에 표시됩니다:

`{ ..., "msg":"successfully established a websocket connection to the broker server", ... }`

<figure><img src="../../../../../.gitbook/assets/Broker Client - Setup success message.png" alt="브로커 클라이언트 설정 확인 메시지"><figcaption><p>브로커 클라이언트 설정 확인 메시지</p></figcaption></figure>

브로커 클라이언트 컨테이너의 설정과 세부 정보를 확인하려면 다음을 실행하세요:

```
docker ps
```

다음과 유사한 출력이 표시됩니다:

```
CONTAINER ID   IMAGE                     COMMAND                  CREATED             STATUS             PORTS                    NAMES
7d9a410e7eaa   snyk/broker:azure-repos   "broker --verbose"       About an hour ago   Up About an hour   0.0.0.0:8000->8000/tcp   sweet_williams
6216e27b8d28   snyk/code-agent           "docker-entrypoint.s…"   2 hours ago         Up 2 hours         0.0.0.0:3000->3000/tcp   code-agent
```

## 코드 스니펫 표시 없이 브로커 클라이언트 실행 예시

### **통합된 GitHub 서버용 Broker 클라이언트 실행**

다음 명령어가 사용되어 통합된 GitHub 서버용 Broker 클라이언트를 실행했습니다:

```
docker run --restart=always \
   -p 8000:8000 \
   -e BROKER_TOKEN=b1dcxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx \
   -e GITHUB_TOKEN=ghp_xxxxxxxxxxxxxxxxxxxxxx \
   -e BROKER_CLIENT_URL=http://localhost:8000 \
   -e PORT=8000 \
   -e GIT_CLIENT_URL=http://code-agent:3000 \
   --network mySnykBrokerNetwork \
snyk/broker:github-com
```

여기서:

* `-p 8000:8000`는 호스트 머신의 포트 번호 `8000`이 브로커 클라이언트 컨테이너의 포트 번호 `8000`으로 매핑되어 있는 것입니다. 이는 브로커 클라이언트 컨테이너와 Broker 서버 및 Code Agent 사이의 통신에 사용됩니다.
* `-e BROKER_TOKEN`은 특정 조직 및 GitHub과 연결된 브로커 토큰입니다.
* `-e GITHUB_TOKEN`은 GitHub 리포지터리에 액세스하기 위한 GitHub 토큰입니다.
* `-e BROKER_CLIENT_URL`은 브로커 클라이언트의 호스트 머신 URL인 `http://localhost:8000`입니다.
* `-e PORT`은 브로커 클라이언트가 연결을 수락하는 로컬 포트이며 `8000`입니다.
* `-e GIT_CLIENT_URL=http://code-agent:3000`는 실행 중인 Code Agent 컨테이너의 포트에 대한 URL입니다. URL에는 Code Agent 컨테이너의 이름인 `code-agent`와 포트 번호인 `3000`이 포함되어 있습니다.
* `--network`는 브로커 클라이언트와의 통신에 사용되는 Docker 브릿지 네트워크의 이름입니다 `mySnykBrokerNetwork`.
* `snyk/broker:github-com`은 GitHub을 위한 브로커 클라이언트의 Docker 이미지입니다.

### **통합된 Azure Repos 서버용 Broker 클라이언트 실행**

다음 명령어가 사용되어 통합된 Azure Repos 서버용 Broker 클라이언트를 실행했습니다:

```
docker run --restart=always \
   -p 8000:8001 \
   -e BROKER_TOKEN=fe5bxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx \
   -e AZURE_REPOS_TOKEN=brmyxxxxxxxxxxxxxxxx \
   -e AZURE_REPOS_ORG=snyktest \
   -e AZURE_REPOS_HOST=dev.azure.com \
   -e BROKER_CLIENT_URL=http://localhost:8000 \
   -e PORT=8001 \
   -e GIT_CLIENT_URL=http://code-agent:3000 \
   --network mySnykBrokerNetwork \
snyk/broker:azure-repos
```

여기서:

* `-p 8000:8001`은 호스트 머신의 포트 번호 `8000`이 브로커 클라이언트 컨테이너의 포트 번호 `8001`로 매핑되어 있는 것입니다. 이는 브로커 클라이언트 컨테이너와 Broker 서버 및 Code Agent 사이의 통신에 사용됩니다.
* `-e BROKER_TOKEN`은 특정 조직 및 Azure Repos와 연결된 브로커 토큰입니다.
* `-e AZURE_REPOS_TOKEN`은 Azure Repos 리포지터리에 액세스하기 위한 Azure Repos 토큰입니다.
* `-e AZURE_REPOS_ORG`는 Azure Repos 조직의 이름인 `snyktest`입니다.
* `-e AZURE_REPOS_HOST`는 Azure Repos 서버의 도메인 이름인 `dev.azure.com`입니다.
* `-e BROKER_CLIENT_URL`은 브로커 클라이언트의 호스트 머신 URL인 `http://localhost:8000`입니다.
* `-e PORT`은 브로커 클라이언트가 연결을 수락하는 로컬 포트이며 `8001`입니다.
* `-e GIT_CLIENT_URL=http://code-agent:3000`는 실행 중인 Code Agent 컨테이너의 포트에 대한 URL입니다. URL에는 Code Agent 컨테이너의 이름인 `code-agent`와 포트 번호인 `3000`이 포함되어 있습니다.
* `--network`는 브로커 클라이언트와의 통신에 사용되는 Docker 브릿지 네트워크의 이름입니다 `mySnykBrokerNetwork`.
* `snyk/broker:azure-repos`는 Azure Repos를 위한 브로커 클라이언트의 Docker 이미지입니다.