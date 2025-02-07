# Snyk Broker - 코드 에이전트

{% hint style="warning" %}
**사용 중단됨**

코드 에이전트는 사용 중단되었으며 더 이상 유지되지 않습니다.

Snyk 브로커를 사용하여 Snyk 코드 분석을 실행하는 기본 방법은 \[브로커를 통한 Git 복제]\((../git-clone-through-broker.md) (브로커 코드)를 통해 실행하는 것입니다. 코드 에이전트는 이점이 없는 대안입니다. 자세한 내용은 Snyk 통합 컨설턴트 또는 기술 지원 담당자에게 문의하거나 [Snyk 지원팀](https://support.snyk.io)에 문의하십시오.

코드 에이전트에 대한 자동 \[PR 확인(../../../scan-with-snyk/pull-requests/pull-request-checks/) 기능은 지원되지 않습니다.

코드 에이전트를 비활성화하려면 페이지 [브로커를 통한 Git 복제](../git-clone-through-broker.md)에 대한 지침에 따르십시오.
{% endhint %}

Snyk 코드를 자체 호스팅된 Git 서버에 연결하려면 Snyk 브로커를 설치한 후에 코드 에이전트를 추가할 수 있습니다.

## 코드 에이전트 작동 방식

Snyk 코드를 브로커에 설치한 후 코드 에이전트를 추가하면 Snyk 브로커를 통해 Snyk 코드가 자체 호스팅된 Git 서버와 연결됩니다.

코드 에이전트는 [도커 이미지](https://hub.docker.com/r/snyk/code-agent/)로 제공됩니다. 코드 에이전트는 Snyk 브로커 버전 4.108.0 및 이후 버전에서만 지원됩니다. 이미 실행 중인 브로커 클라이언트가 있는 경우 최신 도커 이미지를 가져와 업데이트해야 합니다.

브로커 클라이언트 및 코드 에이전트를 배포하면 두 개의 별도 서비스가 생성됩니다. 이러한 서비스는 브로커 서버, Snyk 코드 AI 엔진, 및 Snyk 웹 UI와 함께 코드 분석 워크플로우를 가능하게 합니다:

1. 사용자가 자체 호스팅된 Git 서버에서 저장소를 Snyk로 가져와 코드 분석을 시작하는 요청을 Snyk 웹 UI에서 직접 하는 경우 또는 엔드포인트 [대상 가져오기](../../../snyk-api/reference/import-projects-v1.md#org-orgid-integrations-integrationid-import)를 사용하는 경우
2. 요청은 브로커 서버에 도착하고, 브로커 클라이언트로 보내진 후 코드 에이전트로 전송됩니다. 브로커 클라이언트는 통합된 SCM에 대한 연결 세부 정보와 요구되는 저장소를 제공하는 방식으로 코드 에이전트에 자동으로 연결합니다.
3. 코드 에이전트는 통합된 SCM에 연결되어 로컬 저장소를 안전하게 클론합니다. 클론 된 저장소는 코드 에이전트 컨테이너에 임시로 저장됩니다. 클론은 HTTPS 연결을 통해 수행됩니다. SCM이 HTTPS를 지원하지 않는 경우 반대 프록시를 사용합니다. 자세한 내용은 Snyk의 기술 담당자에게 문의하십시오.
4. 코드 에이전트는 지원되는 파일을 필터링하고 클론 저장소를 AI 엔진으로 보냅니다.
5. AI 엔진은 취약점 문제를 검색하기 위해 코드 파일을 분석합니다. 분석 결과는 Snyk 웹 UI로 다시 전송됩니다. 클론된 파일은 클라우드 제공 업체의 저장 최소 정책에 따라 캐시됩니다.

<figure><img src="../../../.gitbook/assets/Code Agent - diagram - new - 4.png" alt="분석 브로커의 워크플로우"><figcaption><p>분석 브로커의 워크플로우</p></figcaption></figure>

## 코드 에이전트 사전 요구 사항

브로커 클라이언트를 설치했음

코드 에이전트 배포 방법을 사용하여 가져올 수있는 최대 파일 크기는 1MB입니다. 자세한 내용은 [Snyk 코드 분석용 파일 크기 제한](../../../supported-languages-package-managers-and-frameworks/technical-specifications-and-guidance.md#file-size-limit-for-snyk-code-analysis)

Docker 컨테이너를 실행할 수있는 능력이 있어야합니다. 예를 들어 Docker Desktop 또는 Kubernetes를 사용합니다.

{% hint style="info" %}
**EU 및 AU용 멀티 테넌트 설정**\
EU 또는 AU 멀티 테넌트 환경에서 브로커, 코드 에이전트 또는 둘 다 설정할 때 추가 환경 변수가 필요합니다.\
예: `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`\
URL에 대한 자세한 내용은 [지역 호스팅 및 데이터 소재](../../../working-with-snyk/regional-hosting-and-data-residency.md)를 참조하십시오.
{% endhint %}

코드 에이전트를 브로커 중복 해결책의 일부로 배포할 수 없습니다.

**코드 에이전트** 구성요소를 실행하는 데 필요한 최소 요구 사항은:

* **CPU** - 1 vCPU
* **메모리** - 2Gb
* **디스크 공간** - 2Gb\
  사용 가능한 디스크 공간은 동시에 가져올 수있는 저장소의 최대 크기를 결정합니다. 이 크기를 초과하는 저장소를 가져오려면 사용 가능한 디스크 공간을 늘려야합니다. 그러나 2Gb보다 큰 저장소를 가져 오기 전에 귀하의 Snyk 팀과 상의하는 것이 강력히 권장됩니다.
* **네트워크:**
  * SCM 연결 - 분석할 저장소를 저장하는 SCM에 대한 HTTPS 통신. HTTP 전용 SCM 배포 지원은 코드 에이전트와 SCM 사이에 반대 프록시를 배치하여 해결할 수 있습니다.
  * 동일한 브로커 클라이언트를 다른 Snyk 제품에 사용하고 해당 클라이언트에서 자동 PR 체크를 활성화하려는 경우 다음을 구성해야합니다:\
    브로커 클라이언트에서 브로커 클라이언트에 대한 통신을 허용하는 내부 연결 (SCM)을 구성해야합니다. 이는 인터넷에서 들어오는 것이 아닙니다. [브로커 클라이언트 실행](install-snyk-broker-code-agent-using-docker/set-up-the-broker-client/run-the-broker-client-without-the-code-snippet-display.md) 참조.
  * 분석 엔진 연결 - [https://deeproxy.snyk.io/](https://deeproxy.snyk.io/)에 위치한 코드 분석 엔진으로의 외부 통신.
* 인터넷 대역폭 및 연결 - 소스 코드를 브로커 서버에 업로드하는 속도는 대역폭이 낮고 인터넷 연결이 느린 경우 영향을 받습니다.
* **Snyk API 토큰** - 코드 에이전트 구성 요소를 Snyk 계정과 인증하는 데 Snyk API 토큰이 필요합니다. 자세한 내용은 [Snyk API 토큰 얻기](../../../getting-started/#obtain-and-use-your-snyk-api-token)을 참조하십시오.
