# 용어집

## A

### Advisor

[Snyk Advisor](https://snyk.io/advisor/) 참조.

### **자산 (Snyk AppRisk)**

Snyk AppRisk 자산은 애플리케이션의 일부인 식별 가능한 개체로, 보안 및 개발자에 관련이 있습니다. Snyk는 보통 소프트웨어 개발 단계에 중점을 두며, 소프트웨어 패키지 자산을 포함하는 보안 저장소 자산을 보호하고, 컨테이너 이미지 자산과 같은 아티팩트를 구축합니다.

### 애플리케이션 (Snyk AppRisk)

애플리케이션은 비즈니스 목적을 제공하는 소프트웨어이며, 앱을 형성하는 자산으로 구성됩니다. 조직은 종종 애플리케이션의 범위를 다르게 정의합니다.

### 애플리케이션 그래프

보안 문제, 애플리케이션 자산, 자산 간 관계, 및 모든 관련 정보를 매핑합니다.&#x20;

## B

### 베이스 이미지

컨테이너 이미지를 구성하는 부모 이미지로, 일반적으로 Dockerfile의 `FROM` 지시문에서 정의됩니다. 베이스 이미지 그 자체는 다른 베이스 이미지에서 구성될 수 있습니다.

### 브로커

[Snyk 브로커](../enterprise-setup/snyk-broker/) 참조.

### 빌드 시스템

소스 코드를 가져와 배포 가능한 애플리케이션(예: 컨테이너)을 빌드하는 시스템입니다.

### 비즈니스 컨텍스트

조직의 목표, 우선순위 및 규제 요구 사항과 관련된 정보로, 앱이 비즈니스에 중요한 정도, 규정 준수 기준, 데이터 민감도 및 수익 또는 평판에 미칠 수 있는 잠재적 영향 등이 포함됩니다.

## C

### CI/CD

지속적 통합(CI), 지속적 전달(CD) 및 지속적 배포(CD)는 소프트웨어 개발 라이프사이클 모델을 이루는 요소로, 개발자들에게 작은, 빈번한 변경 사항의 개발 및 전달을 자동화하도록 안내합니다. 이는 모든 팀원이 최신 코드베이스에 접근할 수 있도록 하고 개발 중에 제출된 코드의 호환성을 보장합니다. 자세한 내용은 [Snyk CI/CD](../scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/)를 참조하세요.

### **클래스 (Snyk AppRisk)**

자산에 비즈니스 컨텍스트를 할당하고 비즈니스 중요성에 따라 자산을 분류하는 방법입니다. 자산은 A, B, C 또는 D 클래스로 할당될 수 있으며, Class A(비즈니스 중요도가 높고 민감한 데이터를 처리하며 규정 준수 대상인 자산 등)가 가장 중요하며, Class D(테스트 앱, 샌드박스 환경 등)가 가장 중요하지 않습니다. 자산은 기본적으로 Class C가 할당됩니다. 정책에서 사용되거나 정책에서 정의할 수 있도록 클래스를 사용할 수 있습니다.

### CLI

Command Line Interface. [Snyk CLI](glossary.md#snyk-cli)를 참조하세요.

### 클라우드 네이티브 애플리케이션 보안

CI/CD 파이프라인 전체에 걸쳐 보안을 구현하고, 보안을 자동화하여 마이크로서비스에 내장하고 취약점 도입을 최소화하는 반복을 극대화하는 것입니다. Snyk는 포괄적인 [CNAS 플랫폼](https://snyk.io/product/cloud-native-application-security/)을 제공합니다.\
[클라우드 네이티브 응용 프로그램 보안을 위한 클라우드 네이티브 보안 가이드](https://snyk.io/learn/cloud-native-security-for-cloud-native-applications/)를 참조하세요.

### 코드 자산 (Snyk AppRisk)

스캔된 리포지토리에서 검색된 모든 자산의 계층적 목록입니다.

### 컨테이너

컨테이너는 애플리케이션과 해당 종속성을 함께 패키징하여 단일 실행 가능한 단위로 배포할 수 있게 하는 기술입니다. 컨테이너는 운영 체제 커널에서 제공하는 추상화로, 프로세스를 시스템에서 실행 중인 다른 프로세스로부터 격리하도록 허용합니다. [Snyk 컨테이너](glossary.md#snyk-container)도 참조하세요.

### 컨테이너 엔진

사용자에게 컨테이너 이미지를 가져와 실행 가능한 컨테이너로 변환하는 애플리케이션입니다. 컨테이너 엔진은 일반적으로 컨테이너 레지스트리와 상호 작용하고 컨테이너를 실행합니다. 도커, CRI-O, LXC가 있는 컨테이너 엔진의 예시입니다.

### 컨테이너 이미지

실행 가능한 컨테이너를 제공하기 위해 컨테이너 엔진이나 실행 환경에서 구성되는 하나 이상의 파일입니다. 이미지는 컨테이너용 패키징 및 배포 형식입니다.

### 컨테이너 레지스트리

컨테이너 이미지를 저장하고 검색하는 메커니즘을 제공하는 서버입니다.

### **제어 (Snyk AppRisk)**

자산과 관련된 보안 제어입니다. Snyk AppRisk Controls 섹션으로 이동하여 보안 제어의 모든 가능한 상태를 확인하세요.

### **커버리지 (Snyk AppRisk)**

적용 가능한 자산이 애플리케이션 보안 프로그램과 관련된 보안 도구(예: Snyk Open Source)에 의해 스캔되고 테스트되었는지에 대한 평가입니다. 특정 정책 유형으로 적용해야 하는 제어사항을 지정하고 옵션으로 실행해야 하는 주기를 선택할 수 있는 정책 유형입니다.

### 커버리지 갭 (Snyk AppRisk)

정책에서 지정한 커버리지 기준을 충족하지 못하거나 스캔이 불규칙하게 이뤄지거나 스캔 전혀 이뤄지지 않는 모든 자산을 평가한 것입니다.

### CVE

일반적으로 사용되는 취약점의 식별자입니다.

### CVSS

공통 취약점 점수 시스템입니다. 취약점의 심각도를 평가하는 산업 표준으로, 0(최저)에서 10(최고)까지 점수를 사용합니다. Snyk는 CVSS를 사용합니다.

### CWE

공통 약점 분류입니다. 소프트웨어 및 하드웨어 약점을 다른 유형으로 분류하는 온라인 용어집입니다. CWE-20: 입력 유효성 검사 등의 예시가 있습니다.

## D

### DAST

동적 애플리케이션 보안 테스팅(Dynamic Application Security Testing)입니다. 실행 중인 애플리케이션을 외부에서 테스트하여 보안 문제를 찾는 보안 분석 기술입니다. [IAST](glossary.md#iast) 및 [SAST](glossary.md#sast)도 참조하세요.  

### 의존성

애플리케이션이 다른 패키지를 사용할 때, 해당 다른 패키지는 자신의 소프트웨어에서 의존성이 됩니다.

- 직접 의존성은 자신의 프로젝트에 포함하는 패키지입니다.
- 간접 의존성(또는 깊거나 체인 또는 추이적 의존성이라고도 함)은 직접 의존성 중 하나가 사용하는 패키지입니다.

### 의존성 트리

의존성 경로라고도 불리며, 소프트웨어 애플리케이션의 의존성을 보여주는 계층적 그래프입니다. 이는 직접 및 간접 의존성을 모두 포함하며 많은 수준으로 깊어질 수 있습니다.

### 개발 컨텍스트

조직 내에서 애플리케이션 개발을 둘러싼 정보 및 요구 사항으로, 소유권, 개발 도구, 환경, 팀, 워크플로우 및 프로세스와 같은 항목이 포함됩니다.

### DevOps

소프트웨어 개발 및 IT 운영을 결합하여 체계 개발 라이프사이클을 단축하는 문화 철학, 실천 방법 및 도구의 세트입니다.

### DevSecOps

보안을 신속하고 투명하게 가장 올바르게 통합하는 신흥하는 애자일 IT 및 DevOps 개발입니다.

### Dockerfile

Docker를 사용하여 컨테이너 이미지를 빌드하는 데 사용되는 텍스트 파일 형식입니다. Dockerfile에는 최종 이미지를 만들기 위해 필요한 모든 명령이 포함되어 있으며, 부모 베이스 이미지를 지정하는 것을 포함합니다.

## E

### 환경

클라우드 환경, [Project 속성](../snyk-admin/snyk-projects/project-attributes.md) 또는 Snyk 작업 인터페이스, 예를 들어 Snyk [CLI](glossary.md#cli), [웹 UI](glossary.md#snyk-web-ui), 또는 [IDE](glossary.md#ide)에 대한 참조일 수 있습니다.

### Exploit

취약점을 악용하는 방법을 시연한 것입니다. 일반적으로 Exploit이 널리 발표되면 "실제로 현장에서 활용되는 exploit"로 일반적으로 언급됩니다. [Exploits 보기](../manage-risk/prioritize-issues-for-fixing/view-exploits.md)를 참조하세요.

### Exploit Maturity

취약점의 exploit이 얼마나 실용적인지를 측정한 지표로, exploit이 널리 사용 가능하고 공격자에게 얼마나 "도움이 되는지"에 따라 결정됩니다.

## F

### 수정 가능 / 부분적으로 수정 가능

취약점이 수정 패치, 업그레이드 또는 핀을 적용하여 수정될 수 있는지를 나타내는 측정 항목입니다. [Vulnerability fix types](../scan-with-snyk/snyk-open-source/manage-vulnerabilities/vulnerability-fix-types.md)를 참조하세요.

### 수정 PR

Snyk가 사용자에게 제공할 수 있는 발견된 취약점에 대한 자동 수정이 포함된 풀 리퀘스트입니다. [Automated fix PRs](../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/create-automatic-prs-for-backlog-issues-and-known-vulnerabilities-backlog-prs.md)를 참조하세요.

## G

### Git

소프트웨어 개발 중 소스 코드의 변경 사항을 추적하는 분산 형상 관리 시스템입니다.

## I

### IaC

Infrastructure as Code. [Snyk Infrastructure as Code](glossary.md#snyk-infrastructure-as-code)를 참조하세요.

### IAST

Interactive Application Security Testing. 실행 중인 응용 프로그램의 코드 동작에 중점을 두어 운영 중 요킹 동작을 결정하는 런타임 분석 도구입니다. 패키지가 로드되는 곳, 실행 중인 응용 프로그램에서 데이터가 어떻게 흐르는지 또는 최종 사용자가 응용 프로그램과 상호 작용하는 방식 등에 대해 알아봅니다. 현재 공급 업체 및 제품에 따라 기능이 다릅니다. IAST는 응용 프로그램을 내부적으로 분석하여 코드의 취약점 소스를 추적함으로써 [DAST](glossary.md#dast)와 같은 것보다 더 자세한 통찰력을 제공합니다. [SAST](glossary.md#sast)도 참조하세요.

### IDE

Integrated Development Environment. 보통 소스 코드 편집기, 빌드 자동화 도구 및 디버거가 있는 소프트웨어 개발을 위한 시설이 있는 응용 프로그램입니다.

### 이미지

애플리케이션을 실행하는 데 필요한 소프트웨어 세트를 포함하는 컨테이너의 저장된 인스턴스입니다.  

### 이미지 레이어

컨테이너 이미지는 일반적으로 여러 다른 파일 시스템 레이어로 구성되어 있다가 실행 시 단일 파일 시스템으로 병합됩니다.

### 통합

Snyk가 작동하는 타사 제품, 애플리케이션 및 플랫폼입니다(예: GitHub와 같은 SCM 시스템). [Snyk와 통합](../integrate-with-snyk/)을 참조하세요.

### 이슈

Snyk에 의해 식별되고 목록화된 라이선스 문제,### 런타임 컨텍스트 (Snyk AppRisk)

애플리케이션이 실행되는 위치 및 방식에 대한 정보.

## S

### SARIF

정적 분석 도구의 출력에 대한 JSON 기반의 표준 형식인 정적 분석 결과 교환 형식입니다.

### SAST

정적 응용프로그램 보안 테스트. 응용프로그램을 실행하지 않고 정적 소스 코드를 검토하여 잠재적인 취약점을 식별하는 보안 분석 기술입니다. [IAST](glossary.md#iast), [DAST](glossary.md#dast), [](glossary.md#snyk-code), 및 [Snyk ](glossary.md#snyk-infrastructure-as-code)도 참조하십시오.

### SBOM

소프트웨어 구성 요소 목록. 소프트웨어의 구성 요소 목록입니다.

### SCA

소프트웨어 구성 요소 분석. 응용프로그램에서 사용 중인 오픈 소스 및 제삼자 구성 요소, 알려진 보안 취약점 및 종종 상대적인 라이센스 제한을 식별하는 보안 분석 기술입니다. [정적 코드 분석](glossary.md#static-code-analysis)과 혼동되어서는 안 됩니다. [](glossary.md#snyk-open-source)도 참조하십시오.

### 스캔된 artifact(들) (Snyk AppRisk)

Snyk AppRisk의 스캔된 artifact는 식별 정보(예: Git 원격 URL과 같은)가 포함되지 않아 리포지토리 자산으로 식별할 수 없는 Snyk에 의해 감지된 엔터티입니다.

### SCM

소스 코드 관리. 코드 저장소(repo) 또는 버전 관리 시스템이라고도 알려집니다. 개발자가 소스 코드를 저장하고 코드 변경 사항을 추적하는 데 사용하는 방법입니다. SCM은 여러 기여자로부터 업데이트를 병합할 때 충돌을 해결하는 데 도움을 줍니다. GitHub는 일반적인 SCM 시스템의 예입니다. [Git 리포지토리 (SCM)](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/)도 참조하십시오.

### SCM 리포지토리 최신 정보 (Snyk AppRisk)

SCM 리포지토리 최신 정보는 마지막 커밋 날짜를 포함하여 리포지토리의 현재 상태에 대한 즉각적인 이해를 제공합니다. 이를 통해 활성 및 휴면 프로젝트를 신속하게 식별하고 유지 보수, 보안 패치, 자원 할당에 관한 의사 결정에 도움이 됩니다. 리포지토리의 상태 및 마지막 커밋 날짜를 반영합니다.

### SDLC

소프트웨어 개발 생명주기. 개발 팀이 따르는 프로세스로, 소프트웨어를 개발하고 유지 관리하는 방법을 설명합니다.

### 보안 정책

오픈 소스 취약점을 평가하기 위한 기준 세트입니다. 보안 정책을 사용하면 특정 취약점에 대해 자동으로 우선순위를 지정하거나 우선순위를 줄일 수 있습니다. [보안 정책](../manage-risk/policies/security-policies/)을 참조하십시오.

### 심각도

취약점이나 라이센스 문제에 적용되는 심각도 수준으로, 응용프로그램의 해당 항목에 대한 위험을 나타냅니다. [심각도 수준](../manage-risk/prioritize-issues-for-fixing/severity-levels.md)을 참조하십시오.

### 스냅샷

프로젝트의 테스트 기록 내에서 개별 보고서입니다. 실행된 시점에서 정확한 종속성 트리 및 취약점 목록이 포함됩니다.

### `.snyk` 정책

Snyk가 CLI 및 CI/CD 플러그인에 대한 패치를 지정하고 특정 분석 동작을 정의하는 데 사용하는 정책 파일입니다. [.snyk 파일](../manage-risk/policies/the-.snyk-file.md)을 참조하십시오.

### Snyk

클라우드 네이티브 애플리케이션 보안(CNAS) 솔루션을 제공하는 플랫폼으로, 개발자가 코드 및 오픈 소스에서 컨테이너 및 클라우드 인프라까지 전체 응용프로그램의 보안을 소유하고 구축할 수 있게 합니다. Snyk는 또한 Snyk 플랫폼을 제공하는 회사입니다. [시작하기](./)를 참조하십시오.

### Snyk Advisor

소프트웨어 패키지를 오픈 소스 생태계 전반에서 비교할 수 있도록 하는 무료 웹 애플리케이션입니다. 커뮤니티 및 보안 데이터를 결합하여 특정 패키지의 전반적인 상태에 대한 통찰을 제공합니다. [Snyk Advisor](https://snyk.io/advisor/)를 참조하십시오.

### Snyk API

개발자가 프로그래밍 방식으로 Snyk와 통합할 수 있게 하는 Snyk 도구입니다. [Snyk API](../snyk-api/)를 참조하십시오.

### Snyk Apps

Snyk Apps는 Snyk와 통합을 구축하는 현대적이고 선호되는 방법으로, OAuth 2.0 기반의 Snyk API를 통해 리소스에 액세스하기 위한 fine-grained 스코프를 노출하고, 개발자 친화적인 경험을 제공합니다. [Snyk Apps](../snyk-api/how-to-use-snyk-apps-apis/)를 참조하십시오.

### Snyk Broker

Snyk가 개인 고객 환경 (Jira, 코드 저장소 또는 컨테이너 레지스트리)를 스캔할 수 있게 하는 클라이언트/서버 시스템으로, 에이전트나 프록시로 작동합니다. Snyk Broker는 메시지를 중계하고 사용자가 허용하는 메시지를 필터링하여, 예를 들어 일부 GitHub API만 Snyk에 노출되도록 하는 등의 기능을 제공합니다. [Snyk Broker](../enterprise-setup/snyk-broker/)를 참조하십시오.

### Snyk CLI

Snyk 플랫폼 도구로, 명령줄 인터페이스를 사용하여 의존성에서 알려진 취약점을 찾아 수정할 수 있게 합니다. [Snyk CLI](../snyk-cli/)를 참조하십시오.

### Snyk Code

Snyk 제품입니다. 소유한 응용프로그램 코드의 취약점을 찾고 수정할 수 있는 정적 응용프로그램 보안 테스트(SAST) 제품입니다. [](../scan-with-snyk/snyk-code/)를 참조하십시오.

### Snyk Container

Snyk 제품으로, 컨테이너 이미지와 Kubernetes 응용프로그램의 취약점을 찾고 수정할 수 있도록 합니다. [](../scan-with-snyk/snyk-container/)를 참조하십시오.

### Snyk IaC

Snyk 제품으로, Kubernetes, Helm 및 Terraform 구성 파일의 취약점을 찾아 수정할 수 있도록 합니다. [](../scan-with-snyk/snyk-iac/)를 참조하십시오.

### Snyk Open Source

Snyk 제품으로, 오픈 소스 취약점을 찾고 수정할 수 있도록 합니다. [](../scan-with-snyk/snyk-open-source/)를 참조하십시오.

### Snyk 플러그인

특정 언어나 빌드 시스템을 스캔하는 데 사용되는 라이브러리입니다.

### Snyk 보안 인텔리전스

Snk 클라우드 네이티브 애플리케이션 보안 플랫폼을 구동하는 구성 요소입니다.\
Snyk Intel 취약점 DB를 포함하여, 알려진 취약점에 대한 자세한 정보와 수정 지침을 제공하는 취약점 데이터베이스입니다. [취약점 DB](https://snyk.io/vuln)를 참조하십시오.

### Snyk 웹 UI

Snyk 기능에 대한 사용자 액세스를 제공하는 브라우저 기반 환경입니다. [Snyk 웹 UI](snyk-web-ui.md)를 참조하십시오.

### 소셜 트렌드

Snyk는 활발하게 논의되는 문제에 대한 트렌드 배너를 표시합니다 (이전에는 Twitter로 알려졌음). [소셜 트렌드를 갖는 취약점](../manage-risk/prioritize-issues-for-fixing/vulnerabilities-with-social-trends.md)을 참조하십시오.

### 소스

[원본](glossary.md#origin-or-source)을 참조하십시오.

### SPDX

소프트웨어 패키지 데이터 교환. 컴퓨터 소프트웨어를 배포하는 데 사용되는 소프트웨어 라이선스에 대한 정보를 문서화하는 데 사용되는 파일 형식입니다. [SPDX](https://spdx.dev/)를 참조하십시오.

### 정적 코드 분석

소스 코드를 검토하여 코드 품질, 구조 또는 성능과 관련된 문제를 식별하는 기술로, 코드 도달 가능성 결정이나 잠재적인 비효율성을 감지하는 것과 같은 작업을 포함합니다. 이 기술은 보안에 대한 고려 사항을 다룰 수 있지만, 주된 초점은 종종 보안 취약점 식별에 특히 관련되며, 코딩 결함에 따른 보안 위험을 야기할 수 있는 것들에 집중합니다. 반면, 정적 응용프로그램 보안 테스트 (SAST)는 주로 코드 내의 보안 취약점을 특정하여 이를 대상으로 합니다.

## T

### 타겟

Snyk가 스캔한 외부 리소스의 표현입니다. 모든 [Snyk 프로젝트](glossary.md#project)는 부모 타겟과 관련이 있으며, 한 타겟은 여러 프로젝트에 관련됩니다. 타겟의 구조는 [원본](glossary.md#origin-or-source)에 따라 달라집니다.

### **태그 (Snyk AppRisk)**

자산을 분류하는 방법으로, 상호 속성에 따라 자산을 인식하거나 처리하는 데 도움이 됩니다. 자산은 인벤토리에서 태그에 의해 필터링되거나 정책 규칙을 만들 때 태그별로 다르게 처리할 수 있습니다. 자산에 태그를 자동으로 할당할 수도 있고, 작성한 정책에 의해 자산에 태그를 할당할 수도 있습니다. GitHub 및 GitLab 주제는 자산 태그로 처리되며, 정책 생성에 사용할 수 있습니다.

## U

### 업그레이드 가능/패치 가능

픽스 유형: 문제가 패키지 버전을 업그레이드하거나 패치를 적용함으로써 해결될 수 있습니다.

## V

### 취약점

Snyk에서 식별된 보안 취약점입니다. [취약점 관리](../scan-with-snyk/snyk-open-source/manage-vulnerabilities/)을 참조하십시오.

## W

### 웹훅

애플리케이션이 실시간 정보를 다른 응용프로그램에 제공하는 방법입니다. Snyk는 코드 변경 사항을 확인하기 위해 웹훅을 사용합니다. [Snyk 웹훅](../snyk-api/how-to-use-snyk-webhooks-apis/)을 참조하십시오.

### 웹 UI

[Snyk 웹 UI](glossary.md#snyk-web-ui)를 참조하십시오.

### Workspace(SCM 통합)

Snyk의 기능입니다. 이를 통해 Snyk는 정확하고 신뢰할 수 있는 취약성 검사 결과를 얻기 위해 Git 리포지토리의 얕은 복사본을 처리할 수 있습니다.

[SCM 통합을 위한 Workspace](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/introduction-to-git-repository-integrations/workspaces-for-scm-integrations.md)를 참조하십시오.