# 산행 안내: 로컬 스캔 시작

이 페이지들은 결과를 확인하기 위해 몇 가지 스캔을 시도하는 방법에 대해 설명합니다.

Snyk은 전체 소프트웨어 개발 라이프사이클을 안전하게 하는 데 도움이 되는 다양한 도구와 프로세스를 보유하고 있습니다. Snyk을 사용하면 코딩하는 동안 코드를 유효성 검사할 수 있습니다. 또한 작업 중이 아닌 경우에도 코드를 모니터링할 수 있습니다. Snyk은 또한 Git 저장소 통합을 통해 프로젝트 전반의 문제에 대한 가시성을 제공하고, 통합, CLI 또는 커버트 컨테이너를 통해 CI/CD와 작업할 수 있습니다. Snyk을 개발 프로세스의 여러 지점에 통합하여 개발자를 지원하고 가시성을 확보하며 응용 프로그램을 보호하는 것은 일반적인 실천 방법입니다.

스캔을 처음 수행하는 경우이거나 작업 중인 단일 응용 프로그램에 대한 결과에 관심이 있는 경우, 로컬 환경에서 스캔을 시작하는 것이 좋습니다. 이 가이드에서 다루고 있습니다.

개인 또는 팀으로 책임을 지고있는 일련의 응용 프로그램이 있다면, Snyk은 몇 번의 클릭만으로 저장소의 문제에 대한 가시성을 시작할 수 있도록 Git 저장소 통합을 구성하는 것을 권장합니다.

{% hint style="info" %}
개인 상황에 가장 적합한 기술 스택, 환경 및 워크플로우에 가장 적합한 도구 또는 도구를 선택하는 것이 중요합니다. 자세한 내용은 언어 페이지를 참조하십시오.
{% endhint %}

보안 성숙도 수준 및 현재 보안 조치 가능 단계에 대해 알아보기 위해 소프트웨어 개발 라이프사이클의 통합 지점을 선택하는 방법에 대해 알아보려면 [회사에서 Snyk 통합하는 방법](https://learn.snyk.io/lesson/integrate-snyk-at-your-company/) 수업을 참조하십시오.

{% hint style="info" %}
코드 스캔을 수행하려면 Snyk Code를 활성화해야 합니다. 자세한 내용은 [Snyk Code 배포](../scan-with-snyk/snyk-code/#deployment)를 참조하십시오.
{% endhint %}

## 프로젝트 테스트

이 안내서에서는 로컬 개발 환경에서 샘플 또는 단일 프로젝트를 테스트하는 방법을 설명합니다. Snyk CLI를 사용하여도 가능합니다.

{% hint style="info" %}
Snyk 무료 계획은 월별 제한된 무료 테스트를 제공합니다. 무제한 테스트를 원하는 경우 유료 계획을 고려하십시오.
{% endhint %}

### 계정 생성 또는 로그인

로컬 환경에서도 Snyk 기능을 사용하려면 Snyk 계정이 필요합니다. 프로젝트를 시도해 보려면 무료 계정을 생성하십시오. 이미 Snyk을 사용 중인 경우, 단일 로그인을 사용하여 Snyk 계정을 가져올 수 있습니다. 자세한 내용은 [시작하기](../getting-started/)를 참조하십시오.

### 로컬 개발 환경에서 프로젝트 테스트

로컬 개발 환경에서 단일 프로젝트를 스캔하려면 IDE와 함께 Snyk 플러그인 또는 익스텐션이 필요합니다. 또한 이 플러그인이나 익스텐션을 Snyk 계정과 연동해야 합니다. 이러한 과정은 다음 비디오에서 시연됩니다.

{% embed url="https://thoughtindustries-1.wistia.com/medias/bva0d7k46b" %}
IDE 설치 및 Snyk 인증하기
{% endembed %}

{% hint style="warning" %}
IDE를 인증할 때 신뢰할 수 있는 폴더를 스캔한다는 경고가 표시될 수 있습니다. Snyk은 코드를 실행하므로, 의존성 정보를 가져 오기 위해 패키지 관리자를 호출하는 것과 같은 작업이 진행되어야 합니다. 계속하려면 스캔하려는 폴더를 신뢰해야 합니다.
{% endhint %}

로컬 프로젝트에서 Snyk IDE 플러그인 또는 익스텐션을 사용하여 오픈 소스 패키지 문제를 포함한 정보를 제공합니다.

{% embed url="https://thoughtindustries-1.wistia.com/medias/bny3j0ywui" %}
오픈 소스 의존성 문제 검토 비디오
{% endembed %}

로컬 프로젝트에서 Snyk IDE 플러그인 또는 익스텐션을 사용하여 코드 문제와 관련 정보를 제공합니다.

{% embed url="https://thoughtindustries-1.wistia.com/medias/fazwzk6oi3" %}
코드 문제 검토 비디오
{% endembed %}

### Snyk CLI로 프로젝트 테스트

일부 패키지 관리자는 로컬 환경에서의 컨텍스트에 의존하기 때문에 로컬 환경에서 또는 CI/CD 파이프라인의 일부로 테스트하는 것이 가장 정확한 결과를 제공합니다.

먼저 [Snyk CLI 설치](../snyk-cli/install-or-update-the-snyk-cli/)를 수행해야 합니다. 설치 후에는 이를 Snyk 계정에 인증해야 합니다. 이러한 과정은 다음 비디오에서 시연됩니다.

{% embed url="https://thoughtindustries-1.wistia.com/medias/ava7rrg7al" %}
CLI 인증 비디오
{% endembed %}

[**Snyk test**](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/)로 수행하는 스캔은 오픈 소스 패키지 문제와 관련하여 수정 조언을 제공합니다. 이를 다음 비디오에서 시연합니다.

{% embed url="https://thoughtindustries-1.wistia.com/medias/b8vrvtmnbu" %}
Snyk test 비디오
{% endembed %}

[**Snyk code test**](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/)로 수행하는 스캔은 해당 프로젝트의 코드에 대한 정적 코드 분석을 수행하고 발견된 취약점 문제 목록, 테스트에 대한 일반 정보 및 테스트 결과 요약을 반환합니다.

[**Snyk container test**](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/)로 수행하는 스캔은 컨테이너 이미지에서 발견된 취약점 목록과 더 안전한 기본 이미지로 업그레이드하는 권장 사항을 반환합니다.

[**Snyk iac test**](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/)로 수행하는 스캔은 발견된 인프라 코드 파일의 문제를 해결하는 방법에 대한 조언을 제공합니다.

## 다음 단계

* Snyk를 프로세스 전체에 롤아웃하는 가이던스가 필요한 작은 팀이 있다면 [팀 구현 가이드](team-implementation-guide/)를 참조하십시오.
* 기술 스택에 맞는 구체적인 권장 사항을 얻으려면 해당 언어 페이지를 참조하십시오.
* Snyk 사용을 확장하고 기업 전반에 Snyk을 도입하고 다양한 팀을 관련시키고 싶다면 [기업 구현 가이드](enterprise-implementation-guide/)를 읽어보십시오.