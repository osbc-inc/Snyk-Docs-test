# Artifactory Repository - Docker를 사용한 설치 및 구성

{% hint style="info" %}
**기능 가용성**

Artifactory Repository와의 통합은 엔터프라이즈 요금제에서만 사용할 수 있습니다. 자세한 정보는 [요금제 및 가격](https://snyk.io/plans/)을 참조하세요.
{% endhint %}

설치하기 전에 [전제 조건](./)과 [Docker를 사용한 설치에 대한 일반 지침](../install-and-configure-broker-using-docker.md)를 검토하세요.

이 통합은 온프레미스 Artifactory Repository 배포와의 안전한 연결을 보장하는 데 유용합니다.

Artifactory Repository와의 비브로커 통합에 대한 정보는 [Artifactory Repository 설정](../../../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/)을 참조하십시오. Artifactory 컨테이너 레지스트리와의 브로커 통합에 대한 정보는 [Snyk 브로커 - 컨테이너 레지스트리 에이전트](https://docs.snyk.io/snyk-admin/snyk-broker/snyk-broker-container-registry-agent)를 참조하십시오.

## Artifactory 레지스트리에 사용할 브로커 구성

Artifactory 레지스트리 배포에 브로커 클라이언트를 사용하려면 `docker pull snyk/broker:artifactory`를 **실행**하십시오. 환경 변수의 정의는 [Artifactory ](artifactory-repository-environment-variables-for-snyk-broker.md)[Repository - Snyk 브로커를 위한 환경 변수](artifactory-repository-environment-variables-for-snyk-broker.md)를 참조하십시오.

## Artifactory Repository를 위한 브로커 클라이언트 설정을 위한 Docker 실행 명령어

{% hint style="info" %}
**기본값 이외의 지역에서의 멀티 테넌트 설정**\
기본값 이외의 지역에서 Snyk 브로커를 설정할 때, 특정 URL이 필요한 추가 환경 변수가 필요합니다. URL 및 예제에 대한 자세한 내용은 [지역 호스팅 및 데이터 주거지, 브로커 URL](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)를 참조하십시오.
{% endhint %}

다음 명령어를 **복사**하여 Artifactory 레지스트리와 함께 사용할 완전히 구성된 브로커 클라이언트를 설정하세요. 관련 구성을 제공하여 Docker 컨테이너를 실행할 수 있습니다:

```console
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e ARTIFACTORY_URL=<yourdomain>.artifactory.com/artifactory \
       snyk/broker:artifactory
```

**npm 또는 Yarn 통합**을 위해 다음 **명령어**를 사용하세요.

```
docker run --restart=always \
  -p 8000:8000 \
  -e BROKER_TOKEN=secret-broker-token \
  -e ARTIFACTORY_URL=acme.com/artifactory \
  -e RES_BODY_URL_SUB=http://acme.com/artifactory \ 
  snyk/broker:artifactory
```

**Docker run 명령어 사용 대안**으로, Artifactory 통합을 위해 환경 변수를 재정의할 파생 Docker 이미지를 사용할 수 있습니다. 환경 변수에 대한 자세한 내용은 [파생 Docker 이미지](../derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md)를 참조하십시오.

## 브로커 클라이언트 컨테이너 시작 및 Artifactory Repository와의 연결 확인

브로커 클라이언트 컨테이너를 시작하기 위해 브로커 클라이언트 구성을 붙여넣으십시오.

Artifactory 통합 설정 페이지를 새로 고침하여 연결 상태를 확인할 수 있습니다. 연결이 올바르게 설정되면 연결 오류가 없습니다.
