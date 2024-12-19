# Visual Studio Code 확장 프로그램

## **일찍 검사하고 개발 중에 수정: 보안 자세를 높이세요**

개발 라이프사이클 초기에 보안 체크를 통합하면 보안 검토를 원활하게 통과하고 나중에 비용이 많이 드는 수정을 피할 수 있습니다.

Snyk Visual Studio Code 확장 프로그램을 사용하면 코드, 오픈 소스 종속성 및 {{IaC(Infrastructure as Code)}} 구성을 분석할 수 있습니다. IDE에서 직접적인 인사이트를 통해 문제를 발생하는 즉시 해결할 수 있습니다.

**주요 기능:**

- **라인 내 문제 강조:** 보안 문제는 코드 내에 직접 표시되며 종류와 심각성별로 분류되어 신속한 식별과 해결이 가능합니다.
- **포괄적인 스캔:** 확장 프로그램은 다음과 같은 다양한 보안 문제를 스캔합니다:
  - [**오픈 소스 보안**](https://snyk.io/product/open-source-security-management/)**:** 직접 및 간접적인 오픈 소스 종속성에서 취약점 및 라이선스 문제를 감지합니다. 자동 수정 제안으로 간편화된 복구가 가능합니다. [Snyk 오픈 소스 문서](https://docs.snyk.io/scan-using-snyk/snyk-open-source)에서 자세히 알아보세요.
  - [**코드 보안**](https://snyk.io/product/snyk-code/)**:** 사용자 지정 코드에서 보안 취약점을 식별합니다. [Snyk 코드 문서](https://docs.snyk.io/scan-using-snyk/snyk-code)에서 자세히 알아보세요.
  - [**IaC 보안**](https://snyk.io/product/infrastructure-as-code-security/)**:** {{IaC(Infrastructure as Code)}} 템플릿(Terraform, Kubernetes, CloudFormation, Azure Resource Manager)에서 구성 문제를 발견합니다. [IaC 문서](https://docs.snyk.io/scan-using-snyk/snyk-iac)에서 자세히 알아보세요.
- **다양한 언어 및 프레임워크 지원:** Snyk 오픈 소스 및 Snyk 코드는 다양한 패키지 관리자, 프로그래밍 언어 및 프레임워크를 지원하며 최신 기술을 지원하기 위해 계속 업데이트됩니다. 지원되는 언어, 패키지 관리자 및 프레임워크에 대한 최신 정보는 [지원되는 언어 기술 페이지](https://docs.snyk.io/supported-languages-package-managers-and-frameworks)에서 확인하세요.

## 확장 프로그램 설치 및 설정 방법

다음 환경에서 Snyk Visual Studio Code 확장 프로그램을 사용할 수 있습니다:

- Linux: AMD64 및 ARM64
- Windows: 386, AMD64 및 ARM64
- macOS: AMD64 및 ARM64

Snyk Visual Studio Code 확장 프로그램은 원격 및 컨테이너 환경을 지원하지 않습니다:

- [Cloud VS Code IDE](https://code.visualstudio.com/docs/editor/vscode-web)
- [VS Code 원격 개발](https://code.visualstudio.com/docs/remote/remote-overview)
- [컨테이너 내부](https://code.visualstudio.com/docs/devcontainers/containers)

언제든지 무료로 [Visual Studio Code marketplace](https://marketplace.visualstudio.com/items?itemName=snyk-security.snyk-vulnerability-scanner)에서 플러그인을 설치하고, 무료 계정을 포함한 모든 Snyk 계정으로 사용할 수 있습니다. 자세한 내용은 [VS Code 확장 프로그램 설치 가이드](https://code.visualstudio.com/docs/editor/extension-marketplace#_install-an-extension)를 참조하세요.

확장 프로그램이 설치되면 자동으로 [Snyk CLI를 다운로드](https://docs.snyk.io/snyk-cli)하여 [Language Server](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/snyk-language-server)를 포함합니다.

다른 Visual Studio Code 확장 프로그램 문서를 따라 진행하세요:

- [Visual Studio Code 확장 프로그램 구성](visual-studio-code-extension-configuration.md)
- [Visual Studio Code 확장 프로그램 인증](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-code-extension/visual-studio-code-extension-authentication)
- [Visual Studio Code Workspace 신뢰](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-code-extension/workspace-trust)
- [Visual Studio Code 확장 프로그램으로 분석 실행](https://docs.snyk.io/integrate-with-snyk/use-snyk-in-your-ide/visual-studio-code-extension/run-an-analysis-with-visual-studio-code-extension)
- [Visual Studio Code 확장 프로그램에서 분석 결과 보기](https://docs.snyk.io/integrate-with-snyk/use-snyk-in-your-ide/visual-studio-code-extension/view-analysis-results-from-visual-studio-code-extension)

## 지원

문제 해결 및 알려진 문제에 대해서는 [Visual Studio Code 확장 프로그램 문제 해결](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-code-extension/troubleshooting-for-visual-studio-code-extension)를 참조하세요.

도움이 필요하면 [Snyk 지원](https://support.snyk.io)으로 요청을 제출하세요.