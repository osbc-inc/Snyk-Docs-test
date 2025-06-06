# Visual Studio Code 확장 프로그램 구성, 환경 변수 및 프록시

## Visual Studio Code 확장 프로그램의 구성

플러그인을 설치한 후에는 다음 확장 프로그램을 위한 구성을 설정할 수 있습니다.

### Snyk 계정

* **인증 방법:** OAuth2 또는 API 토큰으로 인증할지를 지정합니다. `OAuth2`가 기본값입니다.
* **사용자 지정 엔드포인트**: 사용자 지정 다중 테넌트 또는 단일 테넌트 설정의 Snyk API 엔드포인트를 지정합니다. `https://api.snyk.io`를 사용하는 경우, 구성이 필요하지 않습니다. 자세한 내용은 [IDE URL 목록](../../../working-with-snyk/regional-hosting-and-data-residency.md#ides-urls)을 참조하세요.
*   **조직**: `snyk test`를 실행할 조직을 설정하는 것으로, CLI의 `--org=` 옵션과 유사합니다. Snyk`ORG_ID`를 사용할 것을 권장합니다. 조직 슬러그 이름을 지정하는 경우, 값은 URL에서 조직의 슬러그 이름과 일치해야 합니다. URL에서 조직의 슬러그 이름은 다음과 같습니다: `https://app.snyk.io/org/[orgslugname]`.

    이를 지정하지 않으면 테스트를 실행할 때 사용되는 기본적인 조직, 즉 [웹 계정 설정](https://app.snyk.io/account)에서 정의된 기본 조직이 사용됩니다.
* **Snyk에 오류 보고 전송**: 해당 보고서를 분석하여 플러그인의 안정성을 향상시키기 위해 Snyk에 도움을 줍니다.

### 스캔 구성

* **오픈 소스**: 오픈 소스 종속성에 대한 스캐너를 활성화합니다. 기본적으로 활성화됩니다.
* **Snyk Code 보안 문제**: 애플리케이션 코드의 보안 취약점에 대한 스캐너를 활성화합니다. 기본적으로 활성화됩니다.
* **Snyk Code 품질 문제**: 애플리케이션 코드의 코드 품질 문제에 대한 스캐너를 활성화합니다. 기본적으로 비활성화됩니다.

{% hint style="info" %}
2025년 6월 24일부터, 품질 문제는 더 이상 제공되지 않습니다.
{% endhint %}

* **인프라로 코드**: Terraform 및 Kubernetes 코드의 보안 구성에 대한 스캐너를 활성화합니다. 기본적으로 활성화됩니다.
* **심각도 선택**: 심각도에 따라 문제를 필터링합니다. Low부터 Critical까지.
* **모든 문제 vs 새로운 문제**: 모든 문제를 볼지 또는 새로운 문제만 볼지 지정합니다. 후자는 Git 저장소를 필요로 하며, 확장 프로그램은 결과를 기본 브랜치의 결과와 비교합니다.
*   **추가 매개변수**: Open Source 스캔을 위해 추가적인 `snyk test` [CLI 옵션](https://docs.snyk.io/snyk-cli/cli-reference#options-for-multiple-commands)을 설정합니다.

    **관리되지 않는** [**C/C++**](../../../supported-languages-package-managers-and-frameworks/c-c++/) **스캔**을 위해, 옵션인 `--unmanaged`를 사용하여 오픈 소스 패키지의 취약점을 찾습니다. 이 옵션은 관리되지 않는 C/C++ 스캔에만 작동합니다. 다른 언어에는 이 옵션을 사용하지 마십시오. 추가 매개변수는 나 IaC에는 적용되지 않습니다.

## **사용자 경험**

* **스캔 모드:** **자동** 옵션은 파일을 저장할 때와 프로젝트를 여는 경우 자동으로 스캔을 활성화합니다. 코드 및 IaC와 작동합니다.
* **오픈 소스 보안 자동 스캔:** 오픈 소스 보안 분석을 자동 모드로 실행할 수 있도록 설정합니다.

## 실험적

이 섹션에는 갑작스럽게 변경될 수 있는 실험적인 기능이 포함되어 있습니다.

이러한 설정은 안정적인 기능의 일부가 아니며 공식적으로 지원되지 않습니다.

## 초기화

**신뢰하는 폴더**: 특정 폴더를 신뢰할 것으로 표시하는 목록이 포함된 `settings.json` 파일에 링크합니다. 이 설정은 고급 경우나 특정 폴더를 신뢰하지 않도록 지정해야 할 때에만 사용합니다.

## CLI 및 Language Server

* **자동 종속성 관리**가 선택되어 있는 경우, 플러그인은 정의된 CLI 경로 및 Language Server 경로를 정기적으로 다운로드하여 업데이트합니다. 네트워크 구성으로 인해 CLI를 다운로드하는 것이 불가능한 경우, 예를 들어 방화벽 규칙 등으로 인해 다른 방법으로 이러한 종속성을 획득해야 하는 경우, 이 옵션의 선택을 해제하십시오.
* **CLI 경로:** CLI 파일 경로를 변경할 수 있습니다 (옵션 필드).
* **Language Server 경로:** Language Server 파일 경로를 변경할 수 있습니다 (옵션 필드).

## 환경 변수

프로젝트를 분석하기 위해 플러그인이 활용하는 Snyk CLI는 환경 변수가 필요합니다:

* `PATH`: Maven과 같은 필요한 이진 파일의 경로를 지정합니다.
* `JAVA_HOME`: Java 종속성을 분석할 때 사용하려는 JDK의 경로를 지정합니다.

이러한 변수를 설정하는 것은 단순히 `~/.bashrc`을 사용하여 쉘 환경에만 설정하는 것으로는 충분하지 않습니다. IDE를 명령 행에서 시작하지 않거나 쉘 환경을 사용하여 IDE를 시작하는 스크립트 파일을 만들지 않는 한.

* `Windows`에서는 GUI를 사용하거나 `setx` 도구를 사용하여 명령줄에서 환경 변수를 설정할 수 있습니다.
* `macOS`에서는 `launchd` 프로세스가 Finder에서 직접 IDE를 시작하도록 환경 변수를 알아야 합니다. `launchctl setenv` 명령을 사용하여, 예를 들어, 시작할 때 또는 사용자 로그인 시에 사용하는 스크립트를 통해 Finder를 통해 시작된 응용 프로그램을 위한 환경 변수를 설정할 수 있습니다.\
  `macOS` UI에 환경 변수를 제공하는 방법은 운영 체제 버전 간에 변경될 수 있으므로, `~/.bashrc`을 사용하여 정의된 쉘 환경을 활용할 수 있는 작은 쉘 스크립트를 생성하는 것이 더 쉬울 수 있습니다.
* `Linux`에서는 `file /etc/environment`를 업데이트하여 환경 변수를 창 관리자 및 UI로 전파할 수 있습니다.

## 프록시

프록시 뒤에 있을 경우, VS Code 프록시 설정 `Application > Proxy` 를 사용하여 프록시 설정을 구성하거나 `http_proxy`와 `https_proxy` 환경 변수를 사용하여 프록시 설정을 설정할 수 있습니다.

예를 들어, 일반적으로 사용되는 **Proxy Strict SSL** 설정은 프록시 서버 인증서를 에 특화된 CA 목록에 대해 확인해야 함을 지정합니다.
