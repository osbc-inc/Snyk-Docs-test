# 지역 호스팅 및 데이터 보관

{% hint style="info" %}
**기능 이용 가능성**

지역 호스팅 및 데이터 보관은 기업용 요금제에서만 이용 가능합니다. 자세한 정보는 [요금제 및 가격](https://snyk.io/plans)을 참조하십시오.
{% endhint %}

데이터 보존은 Snyk이 선택한 데이터의 일부를 호스팅하는 지역을 제어할 수 있도록 합니다. GDPR에 대한 정보는 [Snyk이 GDPR 준수를 유지하는 방법](how-snyk-handles-your-data.md#how-snyk-maintains-gdpr-compliance)을 참조하십시오.

데이터 보존은 [Snyk Open Source](../scan-with-snyk/snyk-open-source/), [Snyk Code](../scan-with-snyk/snyk-code/), [Snyk Container](../scan-with-snyk/snyk-container/) 및 [Snyk IaC](../scan-with-snyk/snyk-iac/)용으로 사용할 수 있습니다. Snyk은 여러 지역에서 데이터를 호스팅할 수 있습니다.

이 페이지에는 다음 정보가 제공됩니다.

* [지역별 저장 데이터 유형](regional-hosting-and-data-residency.md#regionally-stored-data-types)
* [전 세계 저장 데이터 유형](regional-hosting-and-data-residency.md#globally-stored-data-types)
* [사용 가능한 Snyk 지역](regional-hosting-and-data-residency.md#available-snyk-regions)
* [지역 멀티테넌트 및 싱글테넌트 호스팅](regional-hosting-and-data-residency.md#regional-multi-and-single-tenant-hosting)

[지역별 URL 목록](regional-hosting-and-data-residency.md#regional-urls)도 제공됩니다.

## 지역별 및 글로벌 데이터

Snyk은 서비스 품질을 제공하기 위해 서브프로세서를 사용합니다. 따라서 모든 데이터 유형을 선택한 지역에 저장할 수는 없습니다. 서브프로세서 목록은 [Snyk 웹사이트](https://snyk.io/policies/subprocessors/)에서 확인할 수 있습니다.

제품 별 데이터 처리 방법에 대한 자세한 내용은 [Snyk이 데이터를 처리하는 방법](how-snyk-handles-your-data.md)을 참조하십시오.

### 지역별 저장 데이터 유형

* 취약점 데이터
* 취약점 소스
* 감사 로그
* 통합 관련 데이터
* 고객 소스 코드

### 전 세계 저장 데이터 유형

* 결제 데이터
* 고객 관계 관리 데이터
* 운영 로그 및 메트릭
* 제품 분석
* 지원 티켓
* 사용자 인증 데이터

## 사용 가능한 Snyk 지역

{% hint style="warning" %}
한 번 지역을 선택하면 해당 지역의 데이터를 다른 지역으로 이전할 수 없습니다. 새로운 지역으로 이동하려면 완전한 재온보딩이 필요합니다.
{% endhint %}

시스템의 초기 온보딩 중에 계정 팀과 협력하여 멀티테넌트 호스팅 지역을 선택할 수 있습니다. 싱글테넌트 이용 가능성 (Snyk 프라이빗 클라우드)은 온보딩 이전에 계정 팀에 문의하시기 바랍니다. Snyk 기능을 사용할 때 SNYK-US-01 URL과 다른 URL을 사용하게 됩니다. URL 목록은 [지역별 URL](regional-hosting-and-data-residency.md#regional-urls)에서 확인할 수 있습니다.

인증을 위해 지역을 설정하도록 환경을 구성해야 합니다. SNYK-US-01 URL을 사용하는 경우에는 해당 사항이 적용되지 않습니다. 자세한 사항은 [snyk config environment CLI 도움말](../snyk-cli/commands/config-environment.md)을 참조하십시오.

Snyk은 다음 지역을 위해 데이터 보존을 제공합니다:

|              지역             |           URL          |
| :-------------------------: | :--------------------: |
|       SNYK-US-01 (미국)       |   https://app.snyk.io  |
|       SNYK-US-02 (미국)       | https://app.us.snyk.io |
|   SNYK-EU-01 (독일, 프랑크푸르트)   | https://app.eu.snyk.io |
|       SNYK-AU-01 (호주)       | https://app.au.snyk.io |
| SNYK-GOV-01 (정부용 Snyk (미국)) | https://app.snykgov.io |

싱글테넌트 배포는 Snyk 엔지니어링에 의해 아키텍처 서비스 지원 가능성이 검증된 추가 지역을 지원할 수 있습니다.

{% hint style="info" %}
Snyk AppRisk는 모든 지역에서 작동합니다. 지역을 설정하고 특정 URL을 사용할 필요가 없습니다.
{% endhint %}

## 지역별 멀티테넌트 및 싱글테넌트 호스팅

{% hint style="warning" %}
SNYK-US-01 지역은 2024년 9월 이전 가입 고객 및 무료 요금제 사용자에게 사용 가능합니다.
{% endhint %}

SNYK-US-01 지역에서의 무료 요금제 또는 팀 요금제 사용자는 URL을 위한 구성이 필요하지 않습니다. 다른 지역에 호스팅하려는 경우 해당 지역용으로 환경 설정해야 합니다.

SNYK-US-02, EU 및 AU 데이터 센터 계정은 [기업 요금제](https://snyk.io/plans/)를 구입한 경우에만 이용 가능합니다.

지역 멀티테넌트 및 싱글테넌트 지역에서 SNYK-US-01과 비슷한 기능, 지원 및 성능을 제공합니다. 지역 간 기능 동등성의 최신 개요에 대한 자세한 정보는 계정 팀에 문의하십시오.

모든 멀티테넌트 환경에서 모든 Snyk AppRisk 기능이 지원되며, 싱글테넌트 환경에서는 계정 팀과 가용성을 확인해야 합니다.

## 통합 고려 사항

Snyk 생태계에서 특정 통합을 설정할 때 특별한 사항이 있습니다.

* Helm 차트를 사용하여 Snyk 런타임 센서를 설치하는 경우, Snyk API 기본 URL을 제공해야 합니다. Snyk 런타임 센서 문서 페이지의 [Helm 차트 사용](../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/snyk-runtime-sensor.md#using-a-helm-chart) 섹션의 5단계에서 필요한 모든 세부 정보를 찾을 수 있습니다.
* 서드파티 통합을 설정하는 경우, 해당 서드파티의 기본 API URL을 지정해야 하는지 확인하십시오.

## 지역별 URL

### 로그인 및 웹 UI URL

**SNYK-US-01:** [https://app.snyk.io/](https://app.snyk.io/)

**SNYK-US-02:** [https://app.us.snyk.io/](https://app.us.snyk.io/)

**SNYK-EU-01:** [https://app.eu.snyk.io/](https://app.eu.snyk.io/)

**SNYK-AU-01:** [https://app.au.snyk.io/](https://app.au.snyk.io/)

### 지원 포털 링크

지역별로 티켓을 확인하고 제출하려면 해당 지역의 링크를 사용하십시오.

**SNYK-US-01**: [https://support.snyk.io/](https://support.snyk.io/)

**SNYK-US-02**: [https://support.snyk.io/services/auth/sso/MT\_US](https://support.snyk.io/services/auth/sso/MT_US)

**SNYK-EU-01:** [https://support.snyk.io/services/auth/sso/MT\_EU](https://support.snyk.io/services/auth/sso/MT_EU)

**SNYK-AU-01:** [https://support.snyk.io/services/auth/sso/MT\_AU](https://support.snyk.io/services/auth/sso/MT_AU)

### API URL

참조 문서에서 다음 기본 URL을 사용하십시오.

**SNYK-US-01:** **API v1** : https://api.snyk.io/v1/ 및 **REST** API: https://api.snyk.io/rest/

**SNYK-US-02**: **API v1:** https://api.us.snyk.io/v1/ 및 **REST** API: https://api.us.snyk.io/rest/

**SNYK-EU-01: API v1:** https://api.eu.snyk.io/v1/ 및 **REST** API: https://api.eu.snyk.io/rest/

**SNYK-AU-01: API v1:** https://api.au.snyk.io/v1/ 및 **REST** API: https://api.au.snyk.io/rest/

### CLI 및 CI 파이프라인 URL

CLI 및 CI 실행 CLI는 인스턴스와 대조되도록 설정되어야 합니다.

이를 위해 [CLI v1.1293.0](https://updates.snyk.io/announcing-snyk-cli-v1-1293-0-299452) 및 이후 버전을 사용하려면 [snyk config environment 명령](../snyk-cli/commands/config-environment.md)을 사용하십시오. 예를 들어:

`snyk config environment SNYK-US-02`

[snyk config environment URL 매핑 지원](../snyk-cli/commands/config-environment.md#supported-environment-urls-mappings)은 `snyk config environment` 도움말에 나열되어 있습니다.

### IDE URL

{% hint style="warning" %}
IDE 플러그인의 최신 버전을 사용하고 있는지 확인하십시오. 다음이 필요한 최소 버전을 지정합니다:\
VSCode - 1.2.18\
Visual Studio - 1.1.21\
IntelliJ - 2.4.32
{% endhint %}

Snyk IDE 확장 프로그램에는 CLI와 유사하게 수정 가능한 옵션이 있으며 적절한 엔드포인트를 사용하도록 설정해야 합니다. IDE의 Snyk 확장 설정에서 **Custom Endpoint**를 해당 값으로 설정하십시오.

**SNYK-US-01:** `https://api.snyk.io` (구성이 필요 없음)

**SNYK-US-02:** `https://api.us.snyk.io`

**SNYK-EU-01:** `https://api.eu.snyk.io`

**SNYK-AU-01 :** `https://api.au.snyk.io`

### Broker URL

[github.com/snyk/broker](https://github.com/snyk/broker)를 사용하고 컨테이너에 추가 환경 변수를 추가하십시오.

**SNYK-US-01:** `https://broker.snyk.io` (구성이 필요 없음)

**SNYK-US-02:** `-e BROKER_SERVER_URL=https://broker.us.snyk.io`

**SNYK-EU-01:** `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`

**SNYK-AU-01:** `-e BROKER_SERVER_URL=https://broker.au.snyk.io`

Helm 차트로 배포된 Broker의 경우, [https://github.com/snyk/snyk-broker-helm](https://github.com/snyk/snyk-broker-helm)을 사용하고 다음 변수를 추가하십시오:

**SNYK-US-02:** `--set brokerServerUrl=https://broker.us.snyk.io`

**SNYK-EU-01:** `--set brokerServerUrl=https://broker.eu.snyk.io`

**SNYK-AU-01:** `--set brokerServerUrl=https://broker.au.snyk.io`

### 고가용성(HA) 모드 Broker URL

[고가용성 모드](../enterprise-setup/snyk-broker/high-availability-mode.md) 지침에 따르되, BROKER\_DISPATCHER\_BASE\_URL의 다음 세부 정보를 사용하십시오:

**SNYK-US-02:** `-e BROKER_DISPATCHER_BASE_URL=https://api.us.snyk.io`

**SNYK-EU-01:** `-e BROKER_DISPATCHER_BASE_URL=https://api.eu.snyk.io`

**SNYK-AU-01 :** `-e BROKER_DISPATCHER_BASE_URL=https://api.au.snyk.io`

Helm 차트로 배포된 Broker의 경우, `values.yaml` 파일을 편집하여 brokerDispatcherUrl에 해당하는 세부 정보를 추가하십시오.

### 코드 에이전트가 있는 Broker URL

{% hint style="warning" %}
**사용 중단**

코드 에이전트는 더 이상 사용되지 않으며 유지 관리되지 않습니다.

Snyk Broker를 통해 Snyk 코드 분석을 실행하는 기본 방법은 [Brokered Code](../enterprise-setup/snyk-broker/git-clone-through-broker.md)을 통해 실행하는 것입니다. 코드 에이전트는 혜택이 없는 대체 방법입니다. 자세한 내용은 Snyk 통합 컨설턴트 또는 기술 담당자에 문의하시거나 [Snyk 지원팀에](https://support.snyk.io) 문의하십시오.

자동 [PR 확인](../scan-with-snyk/pull-requests/pull-request-checks/) 기능은 Snyk Broker - 코드 에이전트에서 지원되지 않습니다.
{% endhint %}

[Snyk Broker - 코드 에이전트](../enterprise-setup/snyk-broker/snyk-broker-code-agent/) 지침을 따르되, 코드 에이전트 컨테이너에 추가 환경 변수를 추가하십시오:

**SNYK-US-02**: `-e UPSTREAM_URL=https://deeproxy.us.snyk.io`

**SNYK-EU-01**: `-e UPSTREAM_URL=https://deeproxy.eu.snyk.io`

**SNYK-AU-01**: `-e UPSTREAM_URL=https://deeproxy.au.snyk.io`

Helm 차트로 배포된 코드 에이전트와 함께 사용하는 경우, [https://github.com/snyk/snyk-broker-helm](https://github.com/snyk/snyk-broker-helm)을 사용하고 코드 에이전트 차트에 다음 변수를 추가하십시오:

**SNYK-US-02**: `--set upstreamUrlCodeAgent=https://deeproxy.us.snyk.io`

**SNYK-EU-01**: `--set upstreamUrlCodeAgent=https://deeproxy.eu.snyk.io`

**SNYK-AU-01**: `--set upstreamUrlCodeAgent=https://deeproxy.au.snyk.io`

### Snyk 코드 로컬 엔진(SCLE)이 있는 Broker

`values-customer-settings.yml`파일을 설정하여 해당 지역에 맞는 올바른 Broker 서버 URL을 [브로커 URL](regional-hosting-and-data-residency.md#broker-urls) 섹션에서 찾아 설정하십시오.

그런 다음 `values-customer-settings.yml`에 추가 변수를 추가합니다:

**SNYK-US-02**\
`deeproxy:`\
`verificationEndpoint: "https://api.us.snyk.io/v1/validate/token/snyk-to-deepcode-proxy-validation"`

**SNYK-EU-01**\
`deeproxy:`\
`verificationEndpoint: "https://api.eu.snyk.io/v1/validate/token/snyk-to-deepcode-proxy-validation"`

**SNYK-AU-01**\
`deeproxy:`\
`verificationEndpoint: "https://api.au.snyk.io/v1/validate/token/snyk-to-deepcode-proxy-validation"`
