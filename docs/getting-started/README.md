# 시작하기



{% hint style="info" %}
Snyk에서 지원하는 언어, 패키지 관리자 및 프레임워크를 사용해야 합니다. [지원되는 언어, 패키지 관리자 및 프레임워크](../supported-languages-package-managers-and-frameworks/)를 참조하십시오.
{% endhint %}

## 지원되는 브라우저

{% hint style="info" %}
Snyk는 Microsoft Internet Explorer를 지원하지 않습니다.
{% endhint %}

Snyk는 다음 웹 브라우저의 최신 버전을 지원합니다:

* [Chrome](https://www.google.com/chrome/)
* [Edge](https://www.microsoft.com/en-us/edge?form=MA13FJ)
* [Firefox](https://www.mozilla.org/en-US/firefox/new/)
* [Safari](https://www.apple.com/safari/) (단, [Opening Fix PR](../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/) 제외)

{% hint style="info" %}
Snyk는 브라우저에서 Javascript를 활성화해야 합니다.
{% endhint %}

Snyk 어플리케이션에서 기본 작업을 수행하기 위해:

## Snyk 계정 생성 또는 로그인

무료 계정을 생성하거나 요금제를 가입하려면 [snyk.io](https://snyk.io/)로 이동하십시오. 자세한 내용은 [Snyk 가격제도](https://docs.snyk.io/implement-snyk/enterprise-implementation-guide/trial-limitations)를 참조하십시오.

회사가 기존 Snyk 계정을 보유하고 있고 단일 로그인 (SSO)을 사용하는 경우, 회사 관리자가 제공한 SSO 링크를 사용하십시오.

회사가 Snyk를 사용하려면 초대가 필요한 경우, 처음으로 로그인할 때 조직 목록이 표시될 수 있으며, 이는 Snyk에서 프로젝트에 대한 액세스를 제어하는 조직입니다. 조직에 액세스를 요청하려면 조직 관리자 이름을 선택하십시오.

{% hint style="info" %}
회사에서 사용하는 인증 공급자와 다른 인증 공급자로 로그인하는 경우, 새로운 계정이 생성됩니다. 이 경우 회사의 올바른 조직에 로그인되지 않습니다.
{% endhint %}

Snyk Web UI에 로그인하면 Snyk는 기본 설정된 조직을 표시합니다. Snyk는 또한 CLI를 사용하여 로컬에서 프로젝트를 테스트할 때 기본 조직의 설정을 사용합니다. 기본 조직을 변경하려면 [계정 환경 설정 관리](snyk-web-ui.md#manage-account-preferences-and-settings)를 참조하십시오.

## Snyk 통합 설정

Snyk가 스캔할 위치를 알기 위해 환경에 액세스 권한을 제공해야 합니다. 필요한 통합 유형은 사용하는 시스템, 스캔하려는 항목 및 통합을 추가할 위치에 따라 다릅니다. [조직](https://docs.snyk.io/integrate-with-snyk#integrations-for-snyk) 또는 [그룹](https://docs.snyk.io/integrate-with-snyk#integrations-for-snyk-apprisk)에 통합을 추가해야 합니다. 사용 가능한 통합 프로그램에 대한 정보는 [Snyk SCM 통합](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-scm-integrations) 및 [Snyk와 통합](https://docs.snyk.io/integrate-with-snyk)을 참조하십시오.

코드를 스캔하려면 먼저 Snyk를 그 코드를 보유한 저장소와 통합해야 합니다.

### 가이드된 과정

Snyk 계정을 만든 후 선택적인 시작 워크스루 프롬프트를 따라 정보를 제공하여 Snyk가 경험을 안내하도록 할 수 있습니다. 이 과정에는 통합 방법 선택, 액세스 권한 설정, 자동화 설정 구성 및 해당 통합 인증이 포함됩니다.

원격 코드 저장소에 인증하지 않고 코드를 스캔하려는 경우 CLI 통합을 선택할 수 있습니다. 이를 통해 로컬 머신에서 스캔을 실행하고 결과를 Snyk의 조직에 업로드할 수 있습니다.

### 수동 과정

Snyk에 통합을 언제든지 수동으로 추가할 수 있습니다. 이를 위해 **통합 > 소스 제어**로 이동하십시오. 자세한 내용은 [Snyk와 통합](../integrate-with-snyk/)을 참조하십시오.

{% hint style="info" %}
조직에 대해 이미 구성된 통합은 **구성됨**으로 표시됩니다.
{% endhint %}

## Snyk API 토큰 얻기 및 사용하기

{% hint style="warning" %}
인증하기 전에 지역을 올바르게 설정해야 합니다. 자세한 내용은 [지역 호스팅 및 데이터 저장 위치](../working-with-snyk/regional-hosting-and-data-residency.md)를 참조하고 [지역별 URL 목록](../working-with-snyk/regional-hosting-and-data-residency.md#regional-urls)을 확인하십시오.
{% endhint %}

Snyk API 토큰은 사용자 프로필에서 이용 가능한 개인 토큰입니다. Snyk API 토큰은 Snyk 계정과 특정 조직과 관련이 없습니다.

무료 및 팀 플랜 및 평가판 사용자는 사용자 프로필에서만 이 개인 토큰에 액세스할 수 있습니다. 이 개인 토큰은 로컬 또는 빌드 머신에서 실행되는 Snyk CLI 및 IDE를 인증하는 데 사용될 수 있습니다. API 또는 CI/CD에 인증할 때는 주의하여 개인 토큰을 사용하십시오.

개인 Snyk API 토큰을 얻으려면:

1. Snyk에 로그인하고 개인 **계정 설정**로 이동합니다.
2. **일반** 설정에서 API 토큰 아래에서 **클릭하여 표시**를 선택합니다.
3. API 키를 강조 표시하고 복사합니다.

새 API 토큰을 원하는 경우 **철회 및 재생성**을 선택할 수 있지만, 이전 API 토큰이 무효화됩니다.

API 토큰을 사용할 때와 서비스 계정 토큰을 사용할 때에 대한 자세한 내용은 엔터프라이즈 플랜 사용자에게만 제공되는 [API 인증](../snyk-api/rest-api/authentication-for-api/)을 참조하십시오.

## 스캔 및 문제 식별을 위한 프로젝트 가져오기

Snyk 프로젝트는 예를 들어 오픈 소스 종속성을 나열한 매니페스트 파일과 같이 이슈를 스캔하는 Snyk에서 스캔하는 항목입니다.

프로젝트를 가져오면 Snyk는 해당 가져온 프로젝트를 스캔하고 결과를 확인할 수 있도록 표시합니다.

프로젝트를 가져 오는 것은 다음과 같은 작업을 수행합니다:

* Snyk를 해당 프로젝트에 대한 정기 스캔 실행으로 설정합니다. [사용 설정](../snyk-admin/groups-and-organizations/usage-settings.md)을 참조하십시오.
* 특히 기본 Snyk 테스트가 풀 및 머지 요청에서 실행되도록 하는 일부 자동화, 이는 프로젝트에 취약점이 추가되는 것을 방지해 줍니다. 이러한 자동화는 조건에 따라 빌드를 실패시키며 [통합 설정](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/)에서 비활성화하거나 사용자 정의할 수 있습니다.

## Snyk AppRisk 설정

{% hint style="info" %}
Snyk AppRisk는 엔터프라이즈 플랜 사용자에게만 제공됩니다. 자세한 내용은 [플랜 및 가격](https://snyk.io/plans/)을 참조하십시오.
{% endhint %}

Snyk AppRisk를 통해 애플리케이션 보안 팀은 현대화되고 높은 성능의 개발자 보안 프로그램을 구현하고 관리하며 확장할 수 있습니다. 이는 애플리케이션 보안 포지처 관리 (ASPM) 하에 있는 사용 사례를 다룹니다.

자세한 내용은 [Snyk AppRisk](../scan-with-snyk/snyk-apprisk/)를 참조하십시오.

## 결과 검토 및 문제 수정

프로젝트를 가져온 후 Snyk가 해당 프로젝트를 스캔하고 문제를 식별하면 스캔 결과를 확인하고 문제 해결을 위한 조치를 취할 수 있습니다. 심각도 수준별로 (Critical, High, Medium 또는 Low) 그룹화된 발견된 문제 수를 볼 수 있습니다. 자세한 내용은 [심각도 수준](../manage-risk/prioritize-issues-for-fixing/severity-levels.md)을 참조하십시오.

스캔 결과 및 사용 가능한 작업은 스캔하는 프로젝트 유형에 따라 달라집니다:

* 오픈 소스 라이브러리: [Snyk Open Source](../scan-with-snyk/snyk-open-source/) 참조
* 어플리케이션 코드: [Snyk Code](../scan-with-snyk/snyk-code/) 참조
* 컨테이너 이미지: [Snyk Container](../scan-with-snyk/snyk-container/scan-container-images.md/) 참조
* 인프라스트럭처 코드 (IaC), 쿠버네티스, 헬름 및 테라폼 구성 파일 및 클라우드 잘못된 구성: [Snyk IaC](../scan-with-snyk/snyk-iac/)를 참조하십시오.