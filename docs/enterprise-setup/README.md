# 기업 설정

기업 구성은 기업 Snyk 솔루션을 이해하고 계획하며 전파하는 데 도움을 주는 데 초점을 맞춥니다. 이에는 다음이 포함됩니다.

* [서드파티 도구용 인증](authentication-for-third-party-tools.md)
* [SSO 설정](single-sign-on-sso-for-authentication-to-snyk/)
* [Snyk 브로커](snyk-broker/)
* [서비스 계정](service-accounts/)

이 페이지는 Snyk 엔터프라이즈 플랜을 사용하는 동안 인증 및 연결하는 방법에 대한 지침을 제공합니다.

[기업 구현 가이드](../implement-snyk/enterprise-implementation-guide/)는 오픈 소스와 코드 분석에 초점을 맞추고 있습니다. 컨테이너 및 인프라 코드에 대한 업데이트가 계획되어 있습니다.

이 페이지에서는 여러 팀 구성원이 Snyk를 사용하여 피드백을 받고 참여시키는 방법을 살펴봅니다.

{% hint style="info" %}
Snyk를 무료 또는 팀 플랜으로 사용하고 있고 엔터프라이즈 플랜으로 업그레이드하려는 경우 [기업 구현 가이드](../implement-snyk/enterprise-implementation-guide/)를 참조하십시오.
{% endhint %}

Snyk에는 전체 소프트웨어 개발 라이프사이클을 보호하는 데 도움이 되는 여러 도구와 프로세스가 있습니다. Snyk를 사용하여 코딩하는 동안 스캔하거나 작업하지 않을 때 코드를 모니터링할 수 있습니다. Snyk는 깃 리포지토리 통합을 통해 프로젝트 전체의 문제에 대한 가시성을 제공하거나 플러그인, CLI 또는 선별된 컨테이너를 통해 CI/CD에 통합할 수도 있습니다.

Snyk를 평가하거나 기업 배포를 계획 중인 사용자, 대부분의 프로그래밍 언어에 대해 Snyk는 깃 리포지토리와 통합하는 것을 권장합니다.

{% hint style="info" %}
개인 상황에 따라 최선의 기술 스택, 환경, 작업 흐름에 맞는 도구를 선택하는 것이 가장 중요합니다. 자세한 내용은 [지원되는 언어 및 프레임워크](../supported-languages-package-managers-and-frameworks/)를 참조하십시오.
{% endhint %}

소프트웨어 개발 라이프사이클 내에서 최적의 통합 지점을 선택하는 방법 및 현재 보안 수준에 따라 팀의 경우에 필요한 [회사에서 Snyk 통합](https://learn.snyk.io/lesson/integrate-snyk-at-your-company/)을 참조하십시오.

Snyk가 사용자에게 무엇을 제공할 수 있는지 알아보려면 프로젝트를 시도하십시오.

이 페이지의 나머지 부분에서는 깃 리포지토리를 Snyk에 연결하고 해당 리포지토리에 있는 Snyk 프로젝트의 스캔 결과를 표시하는 방법에 대해 설명합니다. 이 단계는 Snyk를 팀 또는 팀과 함께 구현하기 전에 Snyk가 작동하는 방식을 이해하는 데 중점을 둡니다.

{% hint style="info" %}
Snyk는 각 스캔 유형(오픈 소스, 코드, 컨테이너 또는 IaC)별로 제한된 무료 테스트를 제공합니다. 무제한 테스트를 위해서는 필요한 유형의 테스트에 대한 유료 플랜 및 해당 유형의 설정이 활성화되어 있는지 확인하십시오.
{% endhint %}

## Snyk 계정 생성 또는 로그인

Snyk 기능을 사용하려면 로컬 환경에서도 Snyk 계정이 있어야 합니다. 프로젝트에 Snyk를 시도해 보려면 무료 계정을 생성하십시오. 기업에서 이미 Snyk를 사용하고 있다면 SSO를 통해 Snyk 계정을 자동으로 설정할 수 있습니다. 자세한 내용은 [기존 계정 생성 또는 로그인](../getting-started/#create-or-log-in-to-a-snyk-account)을 참조하십시오.

## **{Snyk Code 활성화**

Snyk에서 새로운 조직을 만들 때, Snyk 코드(SAST) 스캔이 기본적으로 비활성화됩니다. Snyk는 프로젝트를 Snyk에 가져오기 전에 Snyk 코드 제품을 활성화하는 것을 권장합니다.

1. **설정 > {{Snyk Code}}** 옵션을 선택하십시오.
2. {{Snyk Code}}를 활성화하기 위해 토글을 클릭한 후 **변경 사항 저장**을 클릭하십시오.

## **Snyk 프로젝트 추가**

깃 리포지토리 통합을 연결하려면 Snyk 프로젝트를 추가하십시오. 이 비디오에서 프로세스를 보여줍니다.

{% embed url="https://thoughtindustries-1.wistia.com/medias/9hwr0bnvko" %}
프로젝트 추가 비디오
{% endembed %}

## **초기 Snyk 통합 설정 구성**

깃 리포지토리가 연결되면([깃 리포지토리 통합(소스 코드 관리자)](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/) 참조) 자동화 프로세스를 사용할 수 있습니다. 새로운 취약점을 자동으로 확인하고 병합 요청을 자동으로 생성하거나 의존성 업데이트_pull 요청을 자동으로 생성할 수 있습니다. 초기에는 이러한 옵션을 비활성화하는 것이 좋습니다.

각 Snyk 프로젝트의 설정은 Snyk 조직 통합 설정에서 상속됩니다. 이러한 설정이 비활성화되어 있는지 다음 단계를 따라 확인하십시오: 병합 요청을 위한 기본 Snyk 테스트, 자동 수정 병합 요청, 자동 의존성 업그레이드 병합 요청, Dockerfile 기본 이미지 자동 업데이트.

1. 왼쪽 탐색 메뉴에서 **통합** 페이지를 선택하십시오.
2. 깃 리포지토리 통합을 위한 **설정 톱니바퀴 아이콘**을 선택하십시오.
3. **병합 요청을 위한 기본 Snyk 테스트** 섹션에서 다음 사항이 비활성화되어 있는지 확인하십시오:
   1. **오픈 소스 보안 및 라이선스** (PR이 열릴 때 기본 확인 사항)
   2. **자동 수정 병합 요청**: **신규 취약점** 및 **알려진 취약점(백로그)**
   3. **Dockerfile 기본 이미지 자동 업데이트**
   4. **의존성 업그레이드 된 자동 병합 요청**

{% hint style="info" %}
Snyk는 몇 개 이상의 프로젝트를 추가하기 전에 이러한 옵션 및 알림 기본값을 정의하는 것을 권장합니다. 팀이 이러한 옵션을 구현할 준비가 되면 Snyk는 보안 성숙도에 따라 이러한 옵션을 정의하는 것을 권장합니다. 자세한 내용은 [통합 구성](../implement-snyk/enterprise-implementation-guide/phase-2-configure-account/set-visibility-and-configure-an-organization-template/configure-integrations.md)을 참조하십시오.
{% endhint %}

## **Snyk 스캔 결과 확인**

오픈 소스 프로젝트의 경우, Snyk는 의존성 및 라이선스 구성 요소 문제에 대한 가시성을 제공합니다. 일부 패키지 관리자는 특정 문제를 해결하기 위한 풀 요청을 열도록 안내하는 링크를 제공하기도 합니다. 다음 비디오에서 이 과정이 보여집니다.

{% embed url="https://thoughtindustries-1.wistia.com/medias/z0lcfht0na" %}
문제 해결 비디오
{% endembed %}

Snyk는 프로프리에터리 코드에서 잠재적 보안 및 품질 문제에 대한 정보를 제공하며 권장 사항과 일부 수정 예제를 제공합니다. 다음 비디오에서 보여집니다.

{% embed url="https://thoughtindustries-1.wistia.com/medias/p32o09thet" %}
코드 문제 검토 비디오
{% endembed %}

만약 리포지토리에서 Dockerfile을 식별한 경우, Snyk는 감지된 취약점을 포함한 기본 이미지에 대한 정보를 제공합니다. Snyk는 더 안전한 이미지로 교체하기 위한 풀 요청을 열 수 있습니다. 다음 비디오에서 보여집니다.

{% embed url="https://thoughtindustries-1.wistia.com/medias/3s7h3r9o8h" %}
기본 이미지 교체 비디오
{% endembed %}

리포지토리에서 Kubernetes 또는 Terraform 파일을 식별하는 경우, Snyk는 설정에 대한 정보를 제공합니다. 다음 비디오에서 이에 대한 설명이 나와 있습니다.

{% embed url="https://thoughtindustries-1.wistia.com/medias/l29m2dwrk8" %}
인프라 문제 검토 비디오
{% endembed %}

## Snyk CLI로 스캔

일부 패키지 관리자들은 **로컬 환경에서의 컨텍스트에 의존**합니다. 이러한 패키지 관리자들은 **로컬 환경에서 스캔하거나 CI/CD 파이프라인의 일부로 스캔하는 것이 가장 정확한 결과를 제공**합니다.

Snyk CLI 또는 CI/CD 플러그인을 사용하려면 [Snyk CLI 설치](../snyk-cli/install-or-update-the-snyk-cli/)를 시작하십시오. 설치한 후 CLI를 사용하려면 [기계를 인증하여 CLI와 Snyk 계정을 연결](../snyk-cli/authenticate-to-use-the-cli.md)해야 합니다. 다음 비디오에서 나와 있는 내용대로 인증해야 합니다.

{% embed url="https://thoughtindustries-1.wistia.com/medias/ava7rrg7al" %}
CLI 인증 비디오
{% endembed %}

[**Snyk test**](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/)를 사용한 스캔은 해당 프로젝트의 오픈 소스 패키지 문제 및 수정 지침을 제공합니다. 다음 비디오에서 나와 있는 내용을 보여줍니다.

{% embed url="https://thoughtindustries-1.wistia.com/medias/b8vrvtmnbu" %}
Snyk test 비디오
{% endembed %}

[`snyk code test`](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/)를 실행하면 해당 프로젝트의 코드에 대한 정적 코드 분석 테스트가 수행되며 감지된 취약점 문제 목록, 테스트에 대한 일반 정보 및 테스트 결과 요약을 반환합니다.

[`snyk container test`](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/)를 실행하면 컨테이너 이미지에서 발견된 취약점 목록과 더 안전한 기본 이미지로 업그레이드할 권고 사항이 반환됩니다.

[`snyk iac test`](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/)를 실행하면 발견된 인프라 코드 파일의 문제를 해결하는 방법에 대한 조언이 제공됩니다.

## Snyk Enterprise 플랜을 구현하는 다음 단계

* 개발자들이 IDE 또는 CLI를 사용하여 로컬 환경에서 Snyk를 시도하도록 하려면 [지역에서 스캔 시작하기](../implement-snyk/walkthrough-initiate-a-scan-locally.md)을 검토하십시오.
* 귀하의 기술 스택에 대한 구체적인 권장 사항을 얻으려면 해당 언어에 대한 가이드를 참조하십시오.
* 더 많은 팀으로 Snyk를 전파할 계획이 준비되면 자세한 내용은 [기업 구현 가이드](../implement-snyk/enterprise-implementation-guide/)를 확인하십시오.
* 넒은 대상을 대상으로 Snyk를 전파하기 위한 [개발자 출시 패키지](https://assets.ctfassets.net/4un77bcsnjzw/2YfaqJNMsogGNJM6BBQz4p/8f5ca77b9c40a1bbe14cc9fb0aa05462/Snyk-developer-launch-package.pdf)를 확인하십시오. 추가 전략, 커뮤니케이션 템플릿 및 확장 대상을 위한 체크리스트가 포함되어 있습니다.