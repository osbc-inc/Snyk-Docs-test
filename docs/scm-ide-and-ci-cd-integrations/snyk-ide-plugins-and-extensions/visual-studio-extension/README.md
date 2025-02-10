# Visual Studio 확장 프로그램

## **초보를 검사하고 개발 중에 수정하라: 보안 수준을 높이세요**

개발 생명주기 초기에 보안 검사를 통합하면 보안 검토를 원활하게 통과하고 나중에 비싼 수정을 피할 수 있습니다.

Snyk 비주얼 스튜디오 확장 프로그램을 사용하면 코드 및 오픈 소스 의존성을 분석할 수 있습니다. IDE에서 직접적으로 취할 수 있는 통찰력으로 문제를 해결할 수 있습니다.

**주요 특징:**

* **포괄적인 스캔:** 이 확장 프로그램은 다음과 같은 다양한 보안 문제를 스캔합니다:
  * [**오픈 소스 보안**](https://snyk.io/product/open-source-security-management/)**:** 직접 및 간접적인 오픈 소스 종속성에서 취약점과 라이선스 문제를 감지합니다. 자동 수정 제안으로 문제 해결이 간소화됩니다. [문서](https://docs.snyk.io/scan-using-snyk/snyk-open-source)에서 자세히 알아보세요.
  * [**코드 보안**](https://snyk.io/product/snyk-code/)**:** 사용자 지정 코드에서 보안 취약점을 식별합니다. [문서](https://docs.snyk.io/scan-using-snyk/snyk-code)에서 자세히 알아보세요.
* **넓은 언어 및 프레임워크 지원:** 와 는 다양한 패키지 관리자, 프로그래밍 언어 및 프레임워크를 지원하며 최신 기술을 지원하기 위해 지속적으로 업데이트됩니다. 지원되는 언어, 패키지 관리자 및 프레임워크에 대한 가장 최신 정보는 [지원되는 언어 기술 페이지](https://docs.snyk.io/supported-languages-package-managers-and-frameworks)를 참조하세요.

## 확장 프로그램 설치 및 설정 방법

{% hint style="info" %}
최신 버전의 Snyk 비주얼 스튜디오 확장 프로그램은 비주얼 스튜디오 2022 (버전 17.5 이상)을 지원합니다.

이전 플러그인 버전은 Visual Studio 2015, 2017 및 2019을 지원합니다.
{% endhint %}

다음 환경에서 Snyk 비주얼 스튜디오 확장 프로그램을 사용할 수 있습니다:

* Windows: 386, AMD64 및 ARM64
* MacOS: ARM64 프로세서를 사용하는 맥 내부 Windows 가상 머신에서 비주얼 스튜디오 Windows 플러그인

[Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner-vs-2022)에서 언제든지 무료로 플러그인을 설치하고 사용할 수 있으며 무료 요금제를 포함한 모든 Snyk 계정에서 사용할 수 있습니다. 자세한 정보는 [VS 확장 프로그램 설치 가이드](https://learn.microsoft.com/en-us/visualstudio/ide/finding-and-using-visual-studio-extensions?view=vs-2022#find-and-install-extensions)를 참조하세요.

확장 프로그램이 설치된 후에는 **확장 > Snyk** 메뉴를 통해 Snyk을 사용할 수 있습니다.

<figure><img src="../../../.gitbook/assets/image (351) (1) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Snyk extensions menu"><figcaption><p>Snyk 확장 메뉴</p></figcaption></figure>

**뷰 > 기타 창 > Snyk**을  통해 Snyk 도구 창을 열 수도 있습니다.

도구 창이 열리면 Snyk 확장 프로그램이 [Snyk CLI,](https://docs.snyk.io/snyk-cli) 즉 [언어 서버](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/snyk-language-server)를 다운로드하고 있습니다.

다음은 다른 비주얼 스튜디오 확장 프로그램 설명서의 지침에 따라 계속 진행하세요:

* [비주얼 스튜디오 확장 프로그램 구성](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/visual-studio-extension-configuration)
* [비주얼 스튜디오 확장 프로그램 인증](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/visual-studio-extension-authentication)
* [비주얼 스튜디오 워크스페이스 신뢰](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/workspace-trust)
* [비주얼 스튜디오 확장 프로그램으로 분석 실행](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/run-an-analysis-with-visual-studio-extension)
* [비주얼 스튜디오 확장 프로그램에서 분석 결과 보기](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/view-analysis-results-from-visual-studio-extension)

## 지원

문제 해결 및 알려진 문제에 대해서는 [비주얼 스튜디오 확장 프로그램 문제 해결 및 알려진 문제](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/troubleshooting-and-known-issues-with-visual-studio-extension)을 참조하세요.

도움이 필요한 경우 [Snyk 지원](https://support.snyk.io)에 요청을 제출하세요.
