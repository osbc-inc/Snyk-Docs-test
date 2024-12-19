# 이클립스 플러그인

## **초기 스캔, 개발 중 문제 해결: 보안 포지션 강화**

개발 라이프사이클 초기에 보안 검사를 통합하면 보안 검토를 원활하게 통과하고 나중에 비용이 많이 소요되는 수정을 피할 수 있습니다.

Snyk Eclipse 확장 프로그램을 사용하면 코드, 오픈 소스 의존성 및 {{인프라스트럭처 코드}} (IaC) 구성 요소를 분석할 수 있습니다. IDE에서 바로 확인할 수 있는 실질적인 인사이트로 문제를 해결할 수 있습니다.

**주요 기능:**

- **라인 내 문제 강조:** 보안 문제가 코드 내에서 직접 표시되며, 유형 및 심각성에 따라 분류되어 빠른 식별 및 해결이 가능합니다.
- **폭넓은 스캐닝:** 이 확장 프로그램은 다음과 같은 다양한 보안 문제를 스캔합니다:
  - [**오픈 소스 보안**](https://snyk.io/product/open-source-security-management/)**:** 직접 및 간접적인 오픈 소스 의존성에서 취약점 및 라이선스 문제를 감지합니다. 자동화된 수정 제안으로 해결이 용이합니다. [{{Snyk 오픈 소스}} 문서](https://docs.snyk.io/scan-using-snyk/snyk-open-source)에서 자세한 내용을 살펴보세요.
  - [**코드 보안**](https://snyk.io/product/snyk-code/)**:** 사용자 지정 코드에서 보안 취약점을 식별합니다. [{{Snyk 코드}} 문서](https://docs.snyk.io/scan-using-snyk/snyk-code)에서 자세한 내용을 살펴보세요.
  - [**IaC 보안**](https://snyk.io/product/infrastructure-as-code-security/)**:** {{인프라스트럭처 코드}} 템플릿 (Terraform, Kubernetes, CloudFormation, Azure Resource Manager)에서 구성 문제를 발견합니다. [IaC 문서](https://docs.snyk.io/scan-using-snyk/snyk-iac)에서 자세한 내용을 살펴보세요.
- **다양한 언어 및 프레임워크 지원:** {{Snyk 오픈 소스}} 및 {{Snyk 코드}}는 다양한 패키지 관리자, 프로그래밍 언어 및 프레임워크를 지원하며, 최신 기술을 지원하기 위해 지속적으로 업데이트됩니다. 지원되는 언어, 패키지 관리자 및 프레임워크에 대한 최신 정보는 [지원되는 언어 기술 페이지](https://docs.snyk.io/supported-languages-package-managers-and-frameworks)를 참조하십시오.

## 확장 프로그램 설치 및 설정 방법

{% hint style="info" %}
최신 Snyk 이클립스 플러그인은 Eclipse 2024-03 또는 그 이상에서 지원됩니다.&#x20;

이전 플러그인 버전은 Eclipse 2023-03 또는 그 이상에서 지원됩니다.
{% endhint %}

다음 환경에서 Eclipse 플러그인을 사용할 수 있습니다:

- 리눅스: AMD64 및 ARM64
- 윈도우: 386, AMD64 및 ARM64
- 맥OS: AMD64 및 ARM64

[Eclipse 마켓플레이스](https://marketplace.eclipse.org/content/snyk-security)에서 언제든지 무료로 플러그인을 설치하고, Free 계획을 포함한 모든 Snyk 계정에서 사용할 수 있습니다.&#x20;

플러그인을 설치할 때 라이선스 동의를 받고 **Snyk 보안** 인증서를 추가하여 설치를 완료하십시오 (이는 한 번만 발생합니다).

다음 Eclipse 확장 프로그램 문서의 지침을 따라 진행하십시오:

- [Eclipse 플러그인과 함께 CLI 및 랭귀지 서버 다운로드](https://docs.snyk.io/ide-tools/eclipse-plugin/download-the-cli-and-language-server-with-the-eclipse-plugin)
- [Eclipse 플러그인을 위한 인증](https://docs.snyk.io/ide-tools/eclipse-plugin/authentication-for-the-eclipse-plugin)
- [Eclipse 플러그인 폴더 신뢰](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/folder-trust)
- [Eclipse 플러그인 구성](https://docs.snyk.io/ide-tools/eclipse-plugin/configuration-of-the-eclipse-plugin)
- [Eclipse 플러그인을 위한 환경 변수](https://docs.snyk.io/ide-tools/eclipse-plugin/environment-variables-for-the-eclipse-plugin)
- [Eclipse 프로젝트를 보호하기 위해 Snyk 플러그인 사용](https://docs.snyk.io/ide-tools/eclipse-plugin/use-the-snyk-plugin-to-secure-your-eclipse-projects)

## 지원

문제 해결 및 알려진 문제에 대해서는 [이클립스 플러그인을 위한 문제 해결](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/troubleshooting-for-the-eclipse-plugin)을 참조하십시오.

도움이 필요한 경우 [Snyk 지원](https://support.snyk.io)에 요청을 제출하십시오.