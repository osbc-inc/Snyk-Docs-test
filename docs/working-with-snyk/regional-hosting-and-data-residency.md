# 지역 호스팅 및 데이터 주거지

{% hint style="info" %}
**기능 가용성**

지역 호스팅과 데이터 주거지는 엔터프라이즈 플랜에서만 사용할 수 있습니다. 자세한 정보는 [플랜 및 요금](https://snyk.io/plans)을 참조하세요.
{% endhint %}

데이터 주거지는 Snyk가 선택한 데이터 하위 집합을 호스팅하는 지역을 제어할 수 있도록 합니다. GDPR 관련 정보는 [Snyk가 GDPR 규정을 준수하는 방법](how-snyk-handles-your-data.md#how-snyk-maintains-gdpr-compliance)을 참조하세요.

데이터 주거지는 [{{Snyk 오픈 소스}}](../scan-with-snyk/snyk-open-source/), [{{Snyk 코드}}](../scan-with-snyk/snyk-code/), [{{Snyk 컨테이너}}](../scan-with-snyk/snyk-container/), 그리고 [{{Snyk IaC}}](../scan-with-snyk/snyk-iac/)을 위해 사용할 수 있습니다. Snyk는 여러 지역에서 데이터를 호스팅할 수 있습니다.

이 페이지에서 제공하는 정보:

- [지역별 저장된 데이터 유형](regional-hosting-and-data-residency.md#regionally-stored-data-types)
- [글로벌로 저장된 데이터 유형](regional-hosting-and-data-residency.md#globally-stored-data-types)
- [사용 가능한 Snyk 지역](regional-hosting-and-data-residency.md#available-snyk-regions)
- [지역적 멀티-테넌트 및 싱글-테넌트 호스팅](regional-hosting-and-data-residency.md#regional-multi-and-single-tenant-hosting)

[지역 URL 목록](regional-hosting-and-data-residency.md#regional-urls)도 제공됩니다.

## 지역 및 글로벌 데이터

Snyk는 고품질 서비스를 제공하기 위해 서브프로세서를 사용합니다. 따라서 모든 데이터 유형이 원하는 지역에 저장될 수 있는 것은 아닙니다. 서브프로세서 목록은 [Snyk 웹사이트](https://snyk.io/policies/subprocessors/)에서 이용할 수 있습니다.

데이터 처리 방법에 대한 제품별 예제는 [Snyk가 데이터를 처리하는 방법](how-snyk-handles-your-data.md)을 참조하세요.

### 지역적으로 저장된 데이터 유형

- 취약점 데이터
- 취약점 소스
- 감사 로그
- 통합 관련 데이터
- 고객 소스 코드

### 글로벌로 저장된 데이터 유형

- 청구 데이터
- 고객 관계 관리 데이터
- 운영 로그 및 메트릭
- 제품 분석
- 지원 티켓
- 사용자 인증 데이터

## 사용 가능한 Snyk 지역

{% hint style="warning" %}
한 번 지역을 선택하면 해당 지역의 데이터를 다른 지역으로 이전할 수 없습니다. 새 지역으로 이동하려면 완전한 재온보딩이 필요합니다.
{% endhint %}

시스템을 초기 도입할 때 멀티테넌트 호스팅 지역을 선택하기 위해 회계팀과 협력할 수 있습니다. 싱글테넌트 사용 가능성(Snyk 사설 클라우드)의 경우, 도입 전에 회계팀에 문의하시기 바랍니다. Snyk 기능을 사용할 때에는 SNYK-US-01 URL과 다른 특정 URL을 사용합니다. URL 목록은 [지역 URL](regional-hosting-and-data-residency.md#regional-urls)에서 확인할 수 있습니다.

인증하기 전에 환경을 설정하여 지역을 설정해야 합니다. SNYK-US-01 URL을 사용할 때에는 해당되지 않습니다. 자세한 내용은 [snyk config environment CLI 도움말](../snyk-cli/commands/config-environment.md)을 참조하세요.

Snyk는 다음 지역을 위해 데이터 주거지를 제공합니다:

|                 지역                 |           URL          |
| :------------------------------------: | :--------------------: |
|             SNYK-US-01 (미국)            |   https://app.snyk.io  |
|             SNYK-US-02 (미국)            | https://app.us.snyk.io |
|     SNYK-EU-01 (독일, 프랑크푸르트)    | https://app.eu.snyk.io |
|         SNYK-AU-01 (호주)         | https://app.au.snyk.io |
| SNYK-GOV-01 (정부용 Snyk (미국)) | https://app.snykgov.io |

싱글테넌트 배포는 Snyk 엔지니어링의 아키텍처 서비스 지원 확인에 따라 여기에 나열되지 않은 더 많은 지역을 지원할 수 있습니다.

{% hint style="info" %}
Snyk AppRisk는 모든 지역에서 작동합니다. 특정 지역을 설정하고 특정 URL을 사용할 필요가 없습니다.
{% endhint %}

## 지역 멀티-테넌트 및 싱글-테넌트 호스팅

{% hint style="warning" %}
SNYK-US-01 지역은 2024년 9월 이전에 가입한 고객 및 무료 플랜 사용자에게 제공됩니다.
{% endhint %}

SNYK-US-01 지역의 무료 또는 팀 플랜 사용자는 URL을 설정할 필요가 없습니다. 다른 지역에서 호스팅할 경우 해당 지역을 위해 환경을 설정해야 합니다.

SNYK-US-02, EU, AU 데이터 센터 Snyk 계정은 [엔터프라이즈 플랜](https://snyk.io/plans/)을 구입한 경우에만 사용할 수 있습니다.

Snyk는 지역 멀티테넌트 및 싱글테넌트 지역에서 거의 동일한 기능, 지원 및 성능을 제공합니다. 지역 간 기능 동등성에 대한 최신 소개를 확인하려면 회계팀에 문의하십시오.

모든 Snyk AppRisk 기능은 모든 멀티테넌트 환경에서 지원되지만, 싱글테넌트 환경의 사용 가능성은 회계팀 확인이 필요합니다.

## 통합 고려 사항

Snyk 에코시스템에서 특정 통합을 설정할 때 특별한 고려 사항이 있습니다.

- Snyk 런타임 센서를 Helm 차트를 사용하여 설치하는 경우, Snyk API 기본 URL을 제공해야 합니다. Snyk 런타임 센서 설명 페이지의 [Helm 차트 사용](../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/snyk-runtime-sensor.md#using-a-helm-chart) 섹션의 5단계에서 필요한 모든 정보를 찾을 수 있습니다.
- 서드 파티 통합을 설정하는 경우, 해당 서드파티의 기본 API URL을 지정해야 하는지 확인하십시오.

## 지역 URL

### 로그인 및 웹 UI URL

**SNYK-US-01:** [https://app.snyk.io/](https://app.snyk.io/)

**SNYK-US-02:** [https://app.us.snyk.io/](https://app.us.snyk.io/)

**SNYK-EU-01:** [https://app.eu.snyk.io/](https://app.eu.snyk.io/)

**SNYK-AU-01:** [https://app.au.snyk.io/](https://app.au.snyk.io/)

### 지원 포털 링크

지역별로 티켓을 확인하고 제출하려면 해당 지역의 링크를 사용하십시오.

**SNYK-US-01**: [https://support.snyk.io/](https://support.snyk.io/)

**SNYK-US-02**: [https://support.snyk.io/services/auth/sso/MT_US](https://support.snyk.io/services/auth/sso/MT_US)

**SNYK-EU -01:**[https://snyk-mt-eu-prod-1.eu.auth0.com/samlp/xU6rUSC7zvEco2ndKemfJNT6oKtc13Qw](https://snyk-mt-eu-prod-1.eu.auth0.com/samlp/xU6rUSC7zvEco2ndKemfJNT6oKtc13Qw)

**SNYK-AU-01:**[https://snyk-mt-au-prod-1.au.auth0.com/samlp/HnGYPWKxM9JegYL2aq0OAdBAwGJDz0vQ](https://snyk-mt-au-prod-1.au.auth0.com/samlp/HnGYPWKxM9JegYL2aq0OAdBAwGJDz0vQ)

### API URL

다음의 기본 URL을 사용하여 참조 설명서를 사용합니다.

**SNYK-US-01:** **API v1** : https://api.snyk.io/v1/ 및 **REST** API: https://api.snyk.io/rest/

**SNYK-US-02**: **API v1:** https://api.us.snyk.io/v1/ 및 **REST** API: https://api.us.snyk.io/rest/

**SNYK-EU-01: API v1**: https://api.eu.snyk.io/v1/ 및 **REST** API: https://api.eu.snyk.io/rest/

**SNYK-AU-01: API v1**: https://api.au.snyk.io/v1/ 및 **REST** API: https://api.au.snyk.io/rest/

### CLI 및 CI 파이프라인 URL

CLI와 CI 실행 CLI는 인스턴스에 대해 실행하도록 설정해야 합니다.

이를 위해 [CLI v1.1293.0](https://updates.snyk.io/announcing-snyk-cli-v1-1293-0-299452) 및 최신 버전 이 후에 [snyk config environment 명령](../snyk-cli/commands/config-environment.md)을 사용하세요. 예를 들어:

`snyk config environment SNYK-US-02`

[snyk config environment](../snyk-cli/commands/config-environment.md#supported-environment-urls-mappings) 도움말에 나열된 지원 환경 URL 매핑을 참조하세요.

### IDE URL

{% hint style="warning" %}
최신 IDE 플러그인 버전을 사용해야 합니다. 다음은 필요한 최소 버전을 지정한 것입니다:\
VSCode - 1.2.18\
Visual Studio - 1.1.21\
IntelliJ - 2.4.32
{% endhint %}

Snyk IDE 확장은 CLI와 유사하게 변경 가능한 옵션을 갖고 있으며, 적절한 엔드포인트를 사용하도록 설정해야 합니다. IDE에서 Snyk에 대한 설정에서 **Custom Endpoint**를 SNYK-US-02, SNYK-EU-01, 그리고 SNYK-AU-01을 위한 적절한 값으로 설정하세요.

**SNYK-US-01:** `https://api.snyk.io` (구성 필요 없음)

**SNYK-US-02:** `https://api.us.snyk.io`

**SNYK-EU-01:** `https://api.eu.snyk.io`

**SNYK-AU-01 :** `https://api.au.snyk.io`

### 브로커 URL

[github.com/snyk/broker](https://github.com/snyk/broker)를 사용하고 컨테이너에서 추가 환경 변수를 추가하세요:

**SNYK-US-01**: `https://broker.snyk.io` (구성 필요 없음)

**SNYK-US-02**: `-e BROKER_SERVER_URL=https://broker.us.snyk.io`

**SNYK-EU-01:** `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`

**SNYK-AU-01:** `-e BROKER_SERVER_URL=https://broker.au.snyk.io`

Helm 차트로 배포하는 경우 [https://github.com/snyk/snyk-broker-helm](https://github.com/snyk/snyk-broker-helm)를 사용하고 다음 변수를 추가하세요:

**SNYK-US-02:** `--set brokerServerUrl=https://broker.us.snyk.io`

**SNYK-EU-01:** `--set brokerServerUrl=https://broker.eu.snyk.io`

**SNYK-AU-01:** `--set brokerServerUrl=https://broker.au.snyk.io`

### 브로커 고가용성 (HA) 모드 URL

[고가용성 모드](../enterprise-setup/snyk-broker/high-availability-mode.md) 설명을 따르되 BROKER_DISPATCHER_BASE_URL에 다음 세부 정보를 사용하세요:

**SNYK-US-02**: `-e BROKER_DISPATCHER_BASE_URL=https://api.us.snyk.io`

**SNYK-EU-01:** `-e BROKER_DISPATCHER_BASE_URL=https://api.eu.snyk.io`

**SNYK-AU-01 :** `-e BROKER_DISPATCHER_BASE_URL=https://api.au.snyk.io`

Helm 차트로 배포하는 경우, `values.yaml` 파일을 편집하여 brokerDispatcherUrl에 관련 세부 정보를 포함하세요.

### 브로커와 코드 에이전트 URL

{% hint style="warning" %}
**폐기됨**

코드 에이전트는 폐기되었으며 더 이상 지원되지 않습니다.

Snyk 브로커를 통해 Snyk 코드 분석을 실행하는 권장 방법은 [브로커드 코드](../enterprise-setup/snyk-broker/git-clone-through-broker.md)를 통한 것입니다. 코드 에이전트는 이점이 없는 대안 방법입니다. 자세한 정보는 Snyk 통합 컨설턴트 또는 기술 지원 매니저에 문의하시거나 [Snyk 지원](https://support.snyk.io)팀에 문의하세요.

Snyk 브로커 - 코드 에이전트에 대한 자동 [PR 확인](../scan-with-snyk/pull-requests/pull-request-checks/) 기능은 Snyk 브로커 - 코드 에이전트에서 지원되지 않습니다.
{% endhint %}

[Snyk Broker - Code Agent](../enterprise-setup/snyk-broker/snyk-broker-code-agent/) 설명을 따르되 Code Agent 컨테이너에 추가 환경 변수를 추가하세요:

**SNYK-US-02** :`-e UPSTREAM_URL=https://deeproxy.us.snyk.io`

**SNYK-EU-01 :**`-e UPSTREAM_URL=https://deeproxy.eu.snyk.io`

**SNYK-AU-01 :**`-e UPSTREAM_URL=https://deeproxy.au.snyk.io`

Helm 차트로 배포된 Broker와 Code Agent를 사용하는 경우 [https://github.com/snyk/snyk-broker-helm](https://github.com/snyk/snyk-broker-helm)