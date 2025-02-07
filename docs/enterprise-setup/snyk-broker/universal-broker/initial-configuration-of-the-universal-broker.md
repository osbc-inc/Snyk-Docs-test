# 유니버셜 브로커의 초기 구성

유니버설 브로커를 구현하는 과정은 구성 방법에 상관없이 전체적인 단계는 동일합니다. `snyk-broker-config` 명령어를 사용하면 온보딩을 쉽게 할 수 있습니다. 반면에 직접 API 호출은 전반적인 Universal Broker 모델을 이해해야 합니다. 다음은 이러한 단계입니다.

{% hint style="info" %}
사전 조건: 배포, 자격 증명 참조 및 연결을 만들기 위해 테넌트 관리자여야 합니다.
{% endhint %}

* **한 번만:** Snyk 브로커 앱을 조직에 설치합니다. 이렇게 하면 설치 ID, 클라이언트 ID 및 클라이언트 비밀번호가 반환되는데, 이는 모두 Snyk 플랫폼과 상호 작용하는 데 필요합니다. 배포를 생성하려면 조직 ID가 필요합니다.
* **한 번만:** 테넌트 ID와 설치 ID에 대한 Universal Broker 배포를 생성합니다.
* **한 번만:** 연결에 필요한 자격 증명 참조를 만듭니다.
* **한 번만:** 원하는 연결 또는 연결을 생성합니다.
* **각 조직 통합을 위해 한 번씩:** 연결을 사용해야 하는 조직과 통합을 구성합니다.

<figure><img src="../../../.gitbook/assets/image 7-121224.png" alt=""><figcaption><p>GitHub를 위한 Universal Broker 구현 단계 그림</p></figcaption></figure>

## `snyk-broker-config` 명령을 사용한 구성 (권장) <a href="#using-snyk-broker-config-cli" id="using-snyk-broker-config-cli"></a>

`snyk-broker-config` 명령을 사용하여 Universal Broker를 구성하는 단계는 다음과 같습니다.

{% hint style="info" %}
이 명령을 사용하려면 Node 18 이상이 설치되어 있어야 합니다.
{% endhint %}

1. `npm i -g snyk-broker-config`을 실행합니다.
2. 필요한 환경 변수를 설정합니다.
   * 이미 설정되어 있지 않은 경우 `SNYK_TOKEN`을 설정합니다. 이는 개인 API 키여야 합니다.
   * https://api.snyk.io를 타겟팅하고 있지 않은 경우 `SNYK_API_HOSTNAME`을 설정합니다. 예: `export SNYK_API_HOSTNAME=https://api.eu.snyk.io`. [브로커 URL](../../../working-with-snyk/regional-hosting-and-data-residency.md#broker-urls)을 참조하세요.
3. 나머지 단계를 따르는 동안 경험이 원할하게 진행될 수 있도록 필요한 환경 변수를 설정하는 등 필요한 추가 환경 변수를 설정합니다. 예:
   * 모든 명령에서 `TENANT_ID`를 입력할 필요가 없도록합니다.
   * 이미 보유하고 있는 경우 `INSTALL_ID`를 설정하고, 그렇지 않으면 도구가 설치 프로세스를 안내합니다.
4. 사용 가능한 명령을 나열하려면 `snyk-broker-config commands`를 실행합니다.
5. 사용 가능한 대화식 워크플로우를 나열하려면 `snyk-broker-config workflows`를 실행합니다.
6. 배포를 생성합니다.
   * `snyk-broker-config workflows deployments create`를 실행합니다.
   * 필요한 메타데이터 키-값 쌍을 추가합니다.
   * **참고**: 이 워크플로는 이미 설치되어 있지 않은 경우 Snyk 앱을 설치하도록 요청합니다. 설치 ID(환경 변수로 내보내기) 및 자격 증명 출력을 기록하여 이후 단계에서 사용하도록 합니다. 자격 증명을 분실하면 앱을 다시 설치하고 배포를 재생성하거나 업데이트해야 합니다.
   * **선택 사항**: 배포가 생성된 후에 `snyk-broker-config workflows deployments get`을 사용하여 배포를 보십시오.
7. 연결을 생성하고 구성합니다.
   * 연결을 생성하려면 `snyk-broker-config workflows connections create`를 실행합니다.
   * **어떤 배포를 사용하시겠습니까?** 프롬프트에 응답하여 표시된 목록에서 배포를 선택합니다.
   * **어떤 종류의 연결을 생성하시겠습니까?** 프롬프트에 응답하여, 표시된 목록에서 생성할 연결의 유형을 선택합니다.
     * `github` 및 변형, `bitbucket` 및 변형, `gitlab`, `azure`와 같은 SCM 연결 유형 또는 컨테이너 레지스트리 연결(다음 단계 참조), 패키지 관리자 연결, Jira 및 기타 옵션이 포함됩니다.
     * 컨테이너 레지스트리 유형의 브로커 연결의 경우 CR\_AGENT\_URL을 지정하여 컨테이너 레지스트리 에이전트를 가리켜야 합니다.
     * Universal Broker와 별도의 컨테이너 레지스트리 에이전트를 구성하고 실행해야 합니다. [컨테이너 레지스트리 에이전트를 구성 및 실행](../snyk-broker-container-registry-agent/#configuring-and-running-the-container-registry-agent)하는 지침을 따르세요.
   * 프롬프트에 응답하여 각 필수 필드에 대한 구성을 제공합니다.
     * 연결에 사용할 사용자 친화적인 이름을 입력합니다. 공백이 허용되지 않음에 유의하세요.
     * **broker\_client-url**(브로커 클라이언트의 호스트명 및 포트, 예: https://my.broker.company.com:8000)을 입력합니다.
     * 프롬프트에 응답하여 `github-token (민감): 사용할 자격 증명 참조를 선택하십시오. 또는 새로 만들기?`와 같은 프롬프트에 대한 응답을 제공합니다.
   * 메시지 **Connection created** 및 **Ready to configure integrations to use this connection**을 보게 되면 브로커 클라이언트를 실행할 수 있습니다.
8. 연결이 생성된 후에 `snyk-broker-config workflows connections integrate`을 사용하여 새로 생성된 연결을 사용하여 통합을 구성합니다. 프롬프트에 대응하여 사용할 `배포`, 사용할 **연결**, 통합하려는 조직의 `OrgID` 및 `github` 유형인 `integration ID`를 입력하세요. **integration ID**는 **Organization Integrations** 설정에서 찾을 수 있거나 [Integrations](../../../snyk-api/reference/integrations-v1.md) API를 사용하여 검색할 수 있습니다.

일부 통합은 Snyk 웹 UI에서 통합 ID가 표시되지 않을 수 있습니다.

* JIRA, Artifactory 및 Nexus의 경우 통합 ID가 필요하지 않습니다.
* SCM을 통하지 않는 AppRisk 연결 유형의 경우 베타 프로세스 중에는 통합 단계가 필요하지 않습니다. AppRisk 문서에서 연결에 나열된 연결 식별자를 어디에 복사해야 하는지 확인하세요.
* 다른 모든 연결 유형의 경우, 초기에 Snyk UI에 통합 ID가 표시되지 않을 수 있습니다. 통합 ID를 볼 수 있도록 수동으로 값들을 입력해야 할 수 있습니다.

## 예시: 새 연결의 첫 구성 <a href="#quick-examples-below" id="quick-examples-below"></a>

1. [앞 단계](initial-configuration-of-the-universal-broker.md#using-snyk-broker-config-cli)에 설명된대로 첫 번째 설정 및 연결을 생성하여 시작합니다. [첫 번째 설정 및 연결 녹화](https://asciinema.org/a/YqSmUHEWMcDPeQKm6lpeG3qhM)를 참조하세요.
2. 연결을 조직 ID와 통합 ID로 통합합니다. [통합 녹화](https://asciinema.org/a/I2QJxi9MDEeThRZTLD1aTv9cN)를 참조하세요.
3. 연결에 대한 조직 ID 및 통합 ID 목록을 확인합니다. [조직 ID 및 통합 ID 목록 녹화](https://asciinema.org/a/5RWuySWT0M2dDI9mARJjeZS5g)를 참조하세요.
4. 브로커 클라이언트를 실행합니다.

## 컨테이너 엔진 또는 Kubernetes 클러스터에서 브로커 배포 실행

`-e BROKER_SERVER_URL=https://broker`, `REGION.snyk.io`로 원하는 환경을 타겟팅합니다.&#x20;

`broker.snyk.io`를 사용하지 않는 경우 환경 변수 또는 변수를 추가합니다. 참조 자격 증명에 정의된 변수와 연결된 값을 추가합니다. 참조가 누락되면 연결이 설정되지 않으며 브로커 클라이언트 로그에 오류 항목이 기록됩니다.



```
docker run --restart=always 

-p 8000:8000 

-e ACCEPT_CODE=true 

-e DEPLOYMENT_ID=<DEPLOYMENTID> 

-e CLIENT_ID=<CLIENTID> 

-e CLIENT_SECRET=<CLIENTSECRET> 

-e UNIVERSAL_BROKER_ENABLED=true 

-e PORT=8000 

-e BROKER_HA_MODE_ENABLED=true 

-e <YOUR_CREDENTIALS_REFERENCE>=<secret value> 

snyk/broker:universal






```
