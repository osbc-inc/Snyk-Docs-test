# Broker 클라이언트 설정

Snyk 기관에서 설정하려는 SCM에 대해 현재 실행 중인 Broker 클라이언트가 없는지 확인합니다. 이미 실행 중인 Broker 클라이언트가 있는 경우, 해당 Broker 클라이언트를 제거합니다.

그런 다음, Snyk Broker 클라이언트 도커 이미지를 다운로드하거나 업데이트합니다. Code Agent는 Broker 클라이언트 버전 4.108.0 및 이후 버전에서만 지원됩니다. 최신 업데이트를 가져옵니다.

그 다음, Docker 컨테이너를 실행합니다. 브로커 클라이언트를 [코드 스니펫 표시와 함께 실행](run-the-broker-client-with-the-code-snippets-display.md)하거나 [코드 스니펫 표시 없이 실행](run-the-broker-client-without-the-code-snippet-display.md)할 수 있습니다.

프록시 서버와 함께 Broker 클라이언트를 설정할 수 있습니다. 프록시를 사용하는 인프라에서 코드 에이전트를 사용하려면, 코드 에이전트로의 요청이 프록시를 통하지 않고 우회되었는지 확인합니다.

다음과 같은 환경 변수를 Broker 클라이언트에 대한 `docker run` 명령에 추가하고 프록시를 우회하는 명령을 추가하십시오.

```
-e HTTP_PROXY=http://my.proxy.address:<port_no.>
-e HTTPS_PROXY=http://my.proxy.address:<port_no.>
-e NO_PROXY=<code_agent_container_name>
```

프록시가 사용자 이름 및 암호 인증을 요구하는 경우, 다음 환경 변수를 추가하십시오.

```
-e PROXY_AUTH=userID:userPass
```

또한, 코드 에이전트 구성 요소에 이러한 환경 변수를 추가하십시오. 자세한 내용은 [프록시 서버와 함께 코드 에이전트 설정](../set-up-the-code-agent.md#set-up-the-code-agent-to-work-with-a-proxy-server)을 참조하십시오.

프록시를 사용하여 Docker 컨테이너를 사용하는 자세한 정보는 [Docker를 프록시 서버로 설정](https://docs.docker.com/network/proxy/)를 참조하십시오.