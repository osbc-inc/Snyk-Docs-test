# Snyk JetBrains 플러그인 및 IDE 프록시 설정

플러그인을 위해 다음 구성을 설정하려면 **Preferences**, **Tools**로 이동한 다음 **Snyk**를 선택하세요.

## **Snyk 계정**

* **인증 방법:** OAuth2 또는 API 토큰으로 인증할지를 지정합니다. 기본값은 `OAuth2`입니다.
* **토큰:** Snyk와의 인증에 사용할 토큰을 설정합니다. 자세한 내용은 [JetBrains 플러그인을 위한 인증](authentication-for-the-jetbrains-plugins.md)을 참조하세요.
* **사용자 정의 엔드포인트:** 사용자 정의 다중 테넌트 또는 단일 테넌트 설정을 위한 Snyk API 엔드포인트를 지정합니다. `https://api.snyk.io`를 사용하는 경우 별도의 구성이 필요하지 않습니다. 자세한 내용은 [IDE URLs 목록](../../../working-with-snyk/regional-hosting-and-data-residency.md#ides-urls)을 참조하세요.
* **알 수 없는 CA 무시:** 필요한 경우 SSL 인증서를 무시합니다.
*   **조직:** `snyk test`를 실행할 조직을 설정합니다 (CLI의 `--org=` 옵션과 유사). Snyk는 `ORG_ID`를 사용할 것을 권장합니다. 조직 슬러그 이름을 지정하는 경우 값은 Snyk UI에서 조직 URL 슬러그와 일치해야 합니다: `https://app.snyk.io/org/[orgslugname]`.

    이를 지정하지 않으면 테스트를 실행하기 위해 웹 계정 설정에서 정의된 기본 조직이 사용됩니다.

## 스캔 구성

* **{{Snyk 오픈 소스}}**: 오픈 소스 종속성을 위한 스캐너를 활성화합니다. 기본적으로 활성화되어 있습니다.
* **Snyk Infrastructure as Code**: Terraform 및 Kubernetes 코드의 취약한 구성을 위한 스캐너를 활성화합니다. 기본적으로 활성화되어 있습니다.
* **{{Snyk Container}} 취약점**: 컨테이너 이미지 및 Kubernetes 애플리케이션의 취약점을 위한 스캐너를 활성화합니다. 기본적으로 활성화되어 있습니다.
* **{{Snyk Code}} 보안 문제**: 애플리케이션 코드의 보안 취약점을 위한 스캐너를 활성화합니다. 기본적으로 활성화되어 있습니다.
* **{{Snyk Code}} 품질 문제**: 애플리케이션 코드의 코드 품질 문제를 위한 스캐너를 활성화합니다. 기본적으로 비활성화되어 있습니다.
* **심각도 선택**: 낮음부터 심각한까지의 문제를 심각도에 따라 필터링합니다.
* **모든 문제 vs 새로운 문제만 보기**: 모든 문제 또는 새로운 문제만을 볼지를 지정합니다. 후자는 Git 리포지토리를 필요로 하며, 발견된 문제를 기본 브랜치의 결과와 비교합니다.
*   **추가 매개변수**: 추가 `snyk test` [CLI 옵션](../../../snyk-cli/cli-commands-and-options-summary.md)을 설정합니다. 

    **관리되지 않는** [**C/C++**](../../../supported-languages-package-managers-and-frameworks/c-c++/) **스캔**을 위해, 옵션 `--unmanaged`를 사용하여 오픈 소스 패키지의 취약점을 찾을 수 있습니다. 이 옵션은 관리되지 않는 C/C++ 스캔에만 적용되며, 다른 언어에는 사용하지 마십시오. 추가 매개변수는 {{Snyk Code}}나 IaC에는 적용되지 않습니다.

## **실행 파일 설정**

이러한 옵션은 Snyk CLI에 대한 플러그인 의존성 처리를 사용자화하는 데 도움이 됩니다.&#x20;

* **CLI 다운로드를 위한 기본 URL:** CLI의 대체 다운로드 위치를 지정합니다. 이 위치는 https://downloads.snyk.io와 동일한 파일 및 폴더 레이아웃을 준수해야 합니다. 예를 들어, FIPS 지원 CLI는 기본 URL을 https://static.snyk.io/fips로 사용할 것입니다.
* **Snyk CLI 경로:** Snyk CLI의 파일 경로를 변경합니다.
* **필요한 이진 파일 자동 관리** 옵션을 선택하면 플러그인이 CLI를 다운로드하고 정의된 CLI 경로에 정기적으로 업데이트합니다. 네트워크 구성으로 인해 CLI 다운로드가 불가능한 경우, 예를 들어 방화벽 규칙 때문에 CLI를 다른 방법으로 얻어야 하는 경우 이 옵션을 선택 해제하세요.
* **CLI 릴리스 채널:** CLI의 릴리스 채널(**preview**, **rc**, **stable**)을 지정합니다. 여기서 CLI를 특정 버전에 고정하여 지정하는 것도 가능하며, 예를 들어 `v1.1293.1`과 같이 버전을 지정할 수 있습니다.

## **사용자 경험**

**시작 및 저장 시 자동 스캔:** 파일을 저장하거나 프로젝트를 열 때 자동 스캔을 활성화합니다.

## JetBrains 플러그인용 프록시

인터넷에 연결하기 위해 프록시 서버를 사용해야 하는 경우, [JetBrains IDE 설정](https://www.jetbrains.com/help/idea/settings-http-proxy.html)을 사용하여 구성하세요. Snyk 플러그인은 IDE 설정을 자동으로 사용합니다.