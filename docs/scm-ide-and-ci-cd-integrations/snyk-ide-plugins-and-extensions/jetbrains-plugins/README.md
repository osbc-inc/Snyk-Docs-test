# JetBrains 플러그인

## **초반 스캔, 개발 중 문제 해결: 보안 수준 제고**

개발 수명주기 초기에 보안 검사를 통합하면 보안 리뷰를 매끄럽게 통과하고 이른 시일 내 비싼 수정을 피할 수 있습니다.

Snyk JetBrains 플러그인을 사용하면 코드, 오픈 소스 종속성, 도커 이미지 및  구성을 분석할 수 있습니다. IDE에서 바로 사용 가능한 인사이트를 통해 문제가 발생할 때 해결할 수 있습니다.

**주요 기능:**

- **라인 내 문제 강조:** 보안 문제가 코드 내에서 직접 표시되며, 타입 및 심각도에 따라 분류되므로 빠르게 식별하고 해결할 수 있습니다.
- **종합적인 스캐닝:** 해당 확장 도구는 다음과 같은 풍부한 범위의 보안 문제를 스캔합니다.
  - [**오픈 소스 보안**](https://snyk.io/product/open-source-security-management/)**:** 직간접적인 오픈 소스 종속성에서 취약점 및 라이선스 문제를 식별합니다. 자동으로 수정 제안을 하여 해결을 단순화합니다. [Snyk 오픈 소스 문서](https://docs.snyk.io/scan-using-snyk/snyk-open-source)에서 자세히 살펴보세요.
  - [**코드 보안**](https://snyk.io/product/snyk-code/)**:** 사용자 지정 코드에서 보안 취약점을 확인합니다. [Snyk 코드 문서](https://docs.snyk.io/scan-using-snyk/snyk-code)에서 자세히 살펴보세요.
  - [**IaC 보안**](https://snyk.io/product/infrastructure-as-code-security/)**:** IaC 템플릿(Terraform, Kubernetes, CloudFormation, Azure Resource Manager)에서 구성 문제를 드러냅니다. [IaC 문서](https://docs.snyk.io/scan-using-snyk/snyk-iac)에서 자세히 살펴보세요.
  - [**컨테이너 보안**](https://snyk.io/product/container-vulnerability-management/): 베이스 이미지에서 보안 취약점을 찾습니다; 에서 지원하는 모든 [운영 체제 배포판](https://docs.snyk.io/scan-using-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container)을 지원합니다. 또한 [Snyk 컨테이너 문서](https://docs.snyk.io/scan-using-snyk/snyk-container)를 참조하세요.
- **다양한 언어 및 프레임워크 지원:**  및 는 다양한 패키지 관리자, 프로그래밍 언어 및 프레임워크를 지원하며, 최신 기술을 지원하기 위한 업데이트가 지속적으로 이루어집니다. 지원되는 언어, 패키지 관리자 및 프레임워크에 대한 최신 정보를 보려면 [지원되는 언어 기술 페이지](https://docs.snyk.io/supported-languages-package-managers-and-frameworks)를 참조하세요.

## 확장 프로그램 설치 및 설정 방법

{% hint style="info" %}
최신 Snyk JetBrains 플러그인은 모든 JetBrains IDE 2023.3 이상에서 지원됩니다.

이전 플러그인 버전은 JetBrains IDE 2020.3 이상에서 지원됩니다.
{% endhint %}

다음 환경에서 Snyk JetBrains 플러그인을 사용할 수 있습니다.

- Linux: 386, AMD64 및 ARM64
- Linux Alpine: 386 및 AMD64
- Windows: 386, AMD64 및 ARM64
- MacOS: AMD64 및 ARM64

플러그인을 언제든지 무료로 [JetBrains marketplace](https://plugins.jetbrains.com/plugin/10972-snyk-vulnerability-scanner)에서 설치하고 Snyk 계정(무료 플랜 포함)을 통해 사용할 수 있습니다. 자세한 내용은 [IDEA 플러그인 설치 가이드](https://www.jetbrains.com/help/idea/managing-plugins.html)를 참조하세요.

확장 프로그램을 설치하면 자동으로 [Snyk CLI](https://docs.snyk.io/snyk-cli)가 다운로드되며, [Language Server](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/snyk-language-server)를 포함합니다.

다른 JetBrains 플러그인 문서의 지침을 따릅니다.

- [JetBrains 플러그인용 구성, 환경 변수 및 프록시](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/jetbrains-plugins/configuration-environment-variables-and-proxy-for-the-jetbrains-plugins)
- [JetBrains 플러그인 인증](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/jetbrains-plugins/authentication-for-the-jetbrains-plugins)
- [JetBrains 플러그인 폴더 신뢰](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/jetbrains-plugins/jetbrains-plugin-folder-trust)
- [JetBrains 플러그인을 사용하여 분석 실행](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/jetbrains-plugins/run-an-analysis-with-the-jetbrains-plugins)

## 지원

문제 해결 및 알려진 문제에 대해서는 [JetBrains 플러그인 문제 해결](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/jetbrains-plugins/troubleshooting-for-the-jetbrains-plugin)을 참조하세요.

도움이 필요한 경우, [요청](https://support.snyk.io)을 Snyk 지원에 제출하세요.