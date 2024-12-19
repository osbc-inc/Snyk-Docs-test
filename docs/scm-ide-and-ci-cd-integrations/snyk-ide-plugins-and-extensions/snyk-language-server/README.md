# Snyk Language Server

Snyk는 통합 개발 환경 또는 편집기에서 Snyk의 기능을 사용할 수 있도록 허용하는 IDE 통합을 제공합니다. 이 페이지에서는 [Language Server Protocol](https://microsoft.github.io/language-server-protocol/)을 지원하는 IDE나 편집기에 대한 진단을 제공할 수 있는 Snyk Language Server에 대해 설명합니다. 모든 IDE 플러그인 및 사용 방법에 대한 정보는 문서의 [IDE를 위한 Snyk](https://docs.snyk.io/ide-tools)를 참조하십시오.

Snyk Language Server는 취약점, 오픈 소스 라이선스 문제 및 인프라 구성 오류를 검사하고 문제 유형 및 심각도에 따라 보안 문제를 분류하여 결과를 반환합니다.

오픈 소스의 경우 직접 및 간접 의존성에 대해 자동 알고리즘 기반 수정 제안을 받습니다.

Snyk Language Server는 다음 유형의 문제를 스캔합니다:

- [**오픈 소스 보안**](https://snyk.io/product/open-source-security-management/) - Snyk 프로젝트에 편입된 직접 및 간접(순환) 오픈 소스 의존성의 보안 취약성 및 라이선스 문제. [오픈 소스 문서](https://docs.snyk.io/products/snyk-open-source)도 참조하십시오.
- [**코드 보안**](https://snyk.io/product/snyk-code/) - 코드에 대한 보안 취약성. [Snyk Code 문서](https://docs.snyk.io/products/snyk-code)도 참조하십시오.
- [**Infrastructure as Code (IaC) 보안**](https://snyk.io/product/infrastructure-as-code-security/) - IaC 템플릿(Terraform, Kubernetes, CloudFormation, Azure Resource Manager)의 구성 문제. [Snyk Infrastructure as Code 문서](https://docs.snyk.io/products/snyk-infrastructure-as-code)도 참조하십시오.

Language Server를 설치하고 구성한 후 매번 실행하거나 파일을 열거나 저장할 때마다, Snyk가 프로젝트의 매니페스트 파일, 소유 코드 및 구성 파일을 스캔합니다. Snyk는 행동을 취할 수 있는 취약점, 라이선스 또는 구성 문제 세부 정보를 제공하고 결과를 LSP를 지원하는 편집기 또는 IDE에서 네이티브로 표시합니다.

이 페이지에서는 지원되는 환경, 지원 및 피드백 제공 및 설치 지침에 대해 설명합니다.

## 지원되는 운영 체제 및 아키텍처

{% hint style="경고" %}
Snyk 플러그인은 배급 업체에서 지원 종료(End Of Life, EOL)된 모든 운영 체제에서 지원되지 않습니다.
{% endhint %}

다음 환경에서 Language Server를 사용할 수 있습니다:

- Linux: AMD64 및 ARM64
- Linux Alpine: 386 및 AMD64
- Windows: 386, AMD64, 386 호환성을 통한 ARM
- MacOS: AMD64 및 ARM64

## Language Server 다운로드 위치

Snyk Language Server는 Visual Studio Code(VS Code) 및 Eclipse 플러그인을 사용할 때에만 자동으로 다운로드됩니다. Language Server를 수동으로 다운로드할 수도 있습니다. 다음 쉘 스크립트는 그 방법을 보여줍니다.

{% code title="getLanguageServer.sh" lineNumbers="true" %}
```bash
#!/bin/bash
# 이 파일은 최신 Language Server를 다운로드할 수 있게 해 주며, 이는 미관리 에디터와 IDE에 통합하는 데 유용합니다.
# 이는 Snyk에서 다운로드를 지원하지 않은 모든 에디터를 위해 정기적으로 다운로드 및 업데이트해야 하는데, 이러한 작업을 시스템 관리자와 사용자가 할 수 있도록 하는 스크립트입니다.
# Snyk는 항상 최신 버전의 Language Server를 사용하는 것을 권장합니다.

set -e
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m | tr '[:upper:]' '[:lower:]')
if [[ $ARCH == "x86_64" ]]; then
  ARCH="amd64"
fi
if [[ $ARCH == "aarch64" ]]; then
  ARCH="arm64"
fi
PROTOCOL_VERSION=3
VERSION=$(curl https://static.snyk.io/snyk-ls/$PROTOCOL_VERSION/metadata.json | jq .version | sed -e s/\"//g)
wget -O /usr/local/bin/snyk-ls "https://static.snyk.io/snyk-ls/$PROTOCOL_VERSION/snyk-ls_${VERSION}_${OS}_${ARCH}"
```
{% endcode %}

PROTOCOL\_VERSION은 3이지만, 계속된 개발에 따라 증가할 수 있습니다.

## Snyk Language Server 구성

### Snyk LSP 명령행 플래그

`-c <FILE>`은 모든 다른 파일들 앞에 로드할 구성 파일을 지정하는 데 사용됩니다.

`-l <LOGLEVEL>`은 로그 수준(`trace`, `debug`, `info`, `warn`, `error`, `fatal`)을 지정하는 데 사용됩니다. 기본 로그 수준은 `info`입니다.

`-o <FORMAT>`은 문제에 대한 출력 형식(`md` 또는 `html`)을 지정하는 데 사용됩니다.

`-f <FILE>`은 콘솔에 로깅 대신 로그 파일을 지정하는 데 사용됩니다.

`-licenses`는 Language Server에 의해 사용되는 [라이선스](https://github.com/snyk/snyk-ls/tree/main/licenses)를 표시합니다.

### **LSP 초기화 옵션**

`initializationOptions?: LSPAny;` 내에서 [초기화 메시지](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialize)의 일부로, Snyk는 다음 설정을 지원합니다:

```json
{
  "activateSnykOpenSource": "true", // {{Snyk Open Source}}를 활성화함 - 기본값은 true
  "activateSnykCode": "false", // 조직에서 활성화된 경우 {{Snyk Code}}를 활성화함 - 기본값은 false
  "activateSnykIac":  "true", // Infrastructure as Code를 활성화함 - 기본값은 true
  "insecure": "false", // 사용자 지정 CA(Certification Authorities) 허용
  "endpoint":  "https://example.com", // 다중 테넌트 및 단일 테넌트 설정에 필요한 Snyk API 엔드포인트
  "additionalParams": "--all-projects", // 빈칸으로 구분된 Snyk CLI의 추가 매개변수
  "additionalEnv":  "MAVEN_OPTS=-Djava.awt.headless=true;FOO=BAR", // 세미콜론으로 구분된 추가 환경 변수
  "path": "/usr/local/bin", // CLI에서 사용하는 시스템 경로에 추가
  "sendErrorReports":  "true", // 오류를 Snyk에 보낼지 여부 - 기본값은 true
  "organization": "a string", // curl -H "Authorization: token $(snyk config get api)"  https://api.snyk.io/v1/cli-config/settings/sast | jq .org 결과의 조직 이름
  "enableTelemetry":  "true", // 사용자 분석을 추적할 수 있는지 여부
  "manageBinariesAutomatically": "true", // CLI/LS 바이너리를 자동으로 다운로드 및 업데이트할지 여부
  "cliPath":  "/a/patch/snyk-cli", // CLI의 위치 또는 다운로드할 위치
  "token":  "secret-token", // Snyk 토큰, 예: snyk config get api
  "automaticAuthentication": "true", // 스캔 시작 시 LS가 자동으로 인증할지 여부 (기본값: true)
  "enableTrustedFoldersFeature": "true", // 폴더를 신뢰할 것인지 여부
  "trustedFolders": ["/a/trusted/path", "/another/trusted/path"], // 신뢰할 폴더 목록
}
```

모든 .NET 프로젝트에 대해 Snyk는 `--all-projects` 추가 매개변수로 추가하도록 권장합니다.

## Snyk Language Server의 인증

Snyk Language Server가 시작되면 초기화 옵션 `token`에 토큰이 있는지 확인합니다. 토큰이 없는 경우, Snyk Language Server는 Snyk CLI를 사용하여 인증을 시도합니다. CLI도 인증되지 않은 경우, Snyk Language Server가 브라우저 창을 열어 인증을 수행합니다. 웹 브라우저에서 성공적으로 인증을 완료하면 Snyk Language Server가 CLI로부터 자동으로 Snyk 인증 토큰을 검색합니다.

## Snyk Language Server의 환경 변수

Snyk Language Server와 Snyk CLI는 기능을 위해 특정 환경 변수를 지원하고 필요로 합니다:

1. 사용할 HTTP 프록시를 정의하는 `HTTP_PROXY`, `HTTPS_PROXY` 및 `NO_PROXY`
2. Java JVM 기반 프로젝트를 분석하기 위해 `JAVA_HOME`
3. Maven 프로젝트를 분석하기 위해 `PATH` (maven, python 등을 찾을 수 있도록)

## Snyk Language Server의 환경 변수 자동 구성

환경에 이러한 변수를 자동으로 추가하려면, Snyk Language Server는 우선 순위를 결정하기 위해 다음 파일을 검색합니다. 실행 파일이 이미 구성된 환경에서 호출되지 않은 경우 (예: `zsh -i -c 'snyk-ls'`를 통해), 필요한 변수를 설정하기 위해 `-c` 명령행 플래그로 구성 파일을 지정할 수도 있습니다. Snyk Language Server는 주어진 우선도 및 순서에 따라 다음 파일을 읽습니다.

```
-c 플래그를 통해 지정된 구성 파일
<작업 디렉토리>/.snyk.env
$HOME/.snyk.env
```

환경 변수 형식의 `VARIABLENAME=VARIABLEVALUE`을 포함하는 줄은 이미 로드된 변수들을 덮어쓰지 않는 한 환경에 자동으로 추가됩니다. 이는 `dotenv` 형식을 준수합니다. `.profile`, `.zshrc` 등의 경우, 변수가 직접 내보내지는 경우 (예: `export VARIABLENAME=VARIABLEVALUE`와 같이), 로드되지 않습니다. 예를 들어 다음과 같이 분리해야 합니다.

```bash
VARIABLENAME=VARIABLEVALUE
export VARIABLENAME
```

PATH 변수는 다른 모든 변수와 다르게 처리됩니다. PATH 변수는 파일 및 환경에서 찾은 모든 PATH 변수의 집합이므로, 현재 작업 디렉토리 `.`이 자동으로 경로에 추가되어, LSP 클라이언트가 현재 작업 디렉토리에 Snyk CLI를 다운로드하면 Language Server가 해당 CLI를 찾을 수 있습니다.

변수를 구성 파일을 통해 추가하는 것 외에, Snyk Language Server는 Linux 및 macOS에서 다음 디렉토리를 경로에 추가합니다:

- /bin
- $HOME/bin
- /usr/local/bin
- $JAVA\_HOME/bin

JAVA\_HOME이 설정되어 있지 않은 경우, Snyk Language Server는 먼저 `path`에서, 그리고 다음 디렉토리에서 Java 실행 파일을 찾아 부모 디렉토리를 JAVA\_HOME에 추가합니다. 다음 디렉토리가 재귀적으로 검색됩니다:

- /usr/lib
- /usr/java
- /opt
- /Library
- $HOME/.sdkman
- C:\Program Files
- C:\Program Files (x86)

Maven 실행 파일을 찾아 경로에 부모 디렉토리를 추가하기 위해 동일한 디렉토리가 검색됩니다.

## **Snyk CLI**

자동으로 관리되는 Snyk CLI를 찾기 위해 [XDG Data Home](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html#variables) 및 `PATH` 경로가 자동으로 스캔되어 해당 운영 체제에 따른 파일 (예: macOS의 `snyk-macos`, Linux의 `snyk-linux`, Windows의 `snyk-win.exe`)이 발견되는 첫 번째 경로가 환경에 추가됩니다. 나중에 CLI의 모든 기능에 사용됩니다. CLI의 경로는 또한 초기화 옵션 `cliPath`를 사용하여 수동으로 설정할 수 있습니다.

## 폴더 신뢰

Snyk는 취약점을 분석하기 위해 코드베이스에서 자동으로 코드를 실행할 수 있습니다. 이는 Snyk Open Source에 대한 의존성 정보를 얻기 위해 패키지 관리자(예: pip, Gradle, Maven, Yarn, npm 등)를 호출하는 것을 포함합니다. 악의적 구성을 가진 신뢰할 수 없는 코드에서 이러한 프로그램을 실행하는 것은 시스템을 악의적 코드 실행 및 공격에 노출시킬 수 있습니다.

신뢰할 수 없는 폴더에서 Language Server를 사용하는 것에 대비하여, Snyk Language Server는 이러한 폴더에서 스캔을 실행하기 전에 해당 폴더에 대한 신뢰를 요청합니다. 의심스러울 경우, 신뢰를 부여하지 마십시오.

신뢰 기능은 기본적으로 활성화됩니다. 폴더가 신뢰되면 모든 하위 폴더도 신뢰됩니다. 폴더가 신뢰되면, Snyk Language Server는 Language Server 클라이언트에 대한 사용자 지정 `$ / snyk.addTrustedFolders` 통지를 제공하며, 이 통지에는 신뢰할 폴더 경로 목록이 포함됩니다. 클라이언트는 이를 기반으로 해당 통지를 가로채서 결정을 저장하고 IDE 또는 편집기 저장 메커니즘에 신뢰를 보존할 수 있습니다.

신뢰 대화 상자는 초기 옵션에