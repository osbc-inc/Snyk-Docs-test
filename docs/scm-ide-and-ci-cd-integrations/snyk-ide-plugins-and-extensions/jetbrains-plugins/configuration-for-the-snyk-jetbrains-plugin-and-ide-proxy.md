# Snyk JetBrains 플러그인 및 IDE 프록시 설정

플러그인을 위해 다음 구성을 설정하려면 **Preferences**, **Tools**로 이동한 다음 **Snyk**를 선택하세요.

## **Snyk 계정**

* **인증 방법:** OAuth2 또는 API 토큰으로 인증할지를 지정합니다. 기본값은 `OAuth2`입니다.
* **토큰:** Snyk와의 인증에 사용할 토큰을 설정합니다. 자세한 내용은 [JetBrains 플러그인을 위한 인증](authentication-for-the-jetbrains-plugins.md)을 참조하세요.
* **사용자 정의 엔드포인트:** 사용자 정의 다중 테넌트 또는 단일 테넌트 설정을 위한 Snyk API 엔드포인트를 지정합니다. `https://api.snyk.io`를 사용하는 경우 별도의 구성이 필요하지 않습니다. 자세한 내용은 [IDE URLs 목록](../../../working-with-snyk/regional-hosting-and-data-residency.md#ides-urls)을 참조하세요.
* **알 수 없는 CA 무시:** 필요한 경우 SSL 인증서를 무시합니다.
*   **조직:** `snyk test`를 실행할 조직을 설정합니다 (CLI의 `--org=` 옵션과 유사). Snyk는 `ORG_ID`를 사용할 것을 권장합니다. 조직 슬러그 이름을 지정하는 경우 값은 Snyk UI에서 조직 URL 슬러그와 일치해야 합니다: `https://app.snyk.io/org/[orgslugname]`.

    이를 지정하지 않으면 테스트를 실행하기 위해 웹 계정 설정에서 정의된 기본 조직이 사용됩니다.

## 스캔 설정

* **Snyk Open Source**: 오픈 소스 종속성을 위한 스캐너를 활성화합니다. 기본적으로 활성화되어 있습니다.  
* **Snyk Infrastructure as Code**: Terraform 및 Kubernetes 코드의 보안 취약 구성요소를 스캔하는 스캐너를 활성화합니다. 기본적으로 활성화되어 있습니다.  
* **Snyk Container vulnerabilities**: 컨테이너 이미지 및 Kubernetes 애플리케이션의 취약점을 스캔하는 스캐너를 활성화합니다. 기본적으로 활성화되어 있습니다.  
* **Snyk Code Security issues**: 애플리케이션 코드의 보안 취약점을 스캔하는 스캐너를 활성화합니다. 기본적으로 활성화되어 있습니다.  
* **Snyk Code Quality issues**: 애플리케이션 코드의 품질 문제를 스캔하는 스캐너를 활성화합니다. 기본적으로 비활성화되어 있습니다.  
* **Severity selection**: 문제의 심각도(낮음~치명적)에 따라 필터링합니다.  
* **All Issues vs Net New Issues**: 모든 문제를 표시할지 또는 새로운 문제만 표시할지 지정합니다. 새로운 문제만 표시하려면 Git 저장소가 필요하며, 기본 브랜치와의 비교를 통해 새로 발견된 문제를 확인합니다.  
* **Additional parameters**: 오픈 소스 스캔을 위한 추가 `snyk test` [CLI 옵션](../../../snyk-cli/cli-commands-and-options-summary.md)을 설정합니다.  

  **비관리형** [**C/C++**](../../../supported-languages-package-managers-and-frameworks/c-c++/) **스캔**의 경우, `--unmanaged` CLI 옵션을 사용하여 오픈 소스 패키지의 취약점을 찾을 수 있습니다. 이 옵션은 비관리형 C/C++ 스캔에만 적용되며, 다른 언어에서는 사용하지 마십시오. 추가 매개변수는 Snyk Code 또는 IaC에 적용되지 않습니다.

---

## **실행 파일 설정**

이 옵션을 사용하면 Snyk CLI와 관련된 플러그인 동작을 사용자 정의할 수 있습니다.

* **CLI 다운로드 기본 URL:** CLI의 대체 다운로드 위치를 지정할 수 있습니다. 이 위치는 https://downloads.snyk.io와 동일한 파일 및 폴더 구조를 따라야 합니다. 예를 들어, FIPS를 지원하는 CLI는 https://static.snyk.io/fips를 기본 URL로 사용할 수 있습니다.  
* **Snyk CLI 경로:** Snyk CLI의 파일 경로를 변경할 수 있습니다.  
* **자동으로 필요한 바이너리 관리**가 체크되어 있으면 플러그인은 CLI를 다운로드하고, CLI 경로에 정의된 대로 정기적으로 업데이트합니다. 네트워크 구성(예: 방화벽 규칙)으로 인해 CLI를 다운로드할 수 없는 경우 이 옵션을 체크 해제하고 다른 방법으로 CLI를 가져와야 합니다.  
* **CLI 릴리스 채널:** CLI의 릴리스 채널(**preview**, **rc**, **stable**)을 지정할 수 있습니다. 여기에서 버전을 고정할 수도 있으며, 예를 들어 `v1.1293.1`과 같이 버전을 지정할 수 있습니다.  

---

## **사용자 경험**

**시작 및 저장 시 자동 스캔:** 파일을 저장하거나 프로젝트를 열 때 자동 스캔을 활성화합니다.

---

## JetBrains 플러그인용 프록시

인터넷에 연결하기 위해 프록시 서버가 필요한 경우 [JetBrains IDE 설정](https://www.jetbrains.com/help/idea/settings-http-proxy.html)을 통해 설정하세요. Snyk 플러그인은 IDE 설정을 자동으로 사용합니다.