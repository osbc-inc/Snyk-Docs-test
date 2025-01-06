# Visual Studio extension configuration

## 환경 변수

프로젝트를 분석하려면 플러그인에서 Snyk CLI를 사용하며, 필요한 환경 변수가 필요합니다:

* `PATH`: 필요한 이진 파일의 경로를 지정합니다. 예를 들어 Maven의 경우
* `JAVA_HOME`: Java 종속성 분석에 사용할 JDK의 경로를 지정합니다.
* `http_proxy` 및 `https_proxy`: 프록시 서버 뒤에 있는 경우 값은 `http://username:password@proxyhost:proxyport` 형식으로 설정합니다.\
  주의: 값의 처음 `http://`는 `https://`로 바뀌지 않습니다.

이러한 변수는 웹 UI를 사용하거나 `setx` 도구를 사용하여 명령줄에서 설정할 수 있습니다.

## Visual Studio 확장 구성

플러그인이 설치된 후, 확장에 대해 다음 구성을 할 수 있습니다:

* **토큰**: 확장이 Snyk에 연결하는 데 사용하는 토큰을 입력합니다. 다른 계정으로 전환해야 할 경우 토큰을 수동으로 교체할 수 있습니다.
* **사용자 지정 엔드포인트**: 사용자 지정 멀티 테넌트 또는 단일 테넌트 설정을 위한 Snyk API 엔드포인트를 지정합니다. `https://api.snyk.io`를 사용하는 경우 별도의 구성이 필요하지 않습니다. 자세한 정보는 [IDE URLs 목록](../../../working-with-snyk/regional-hosting-and-data-residency.md#ides-urls)을 참조하십시오.
* **알 수 없는 CA 무시**: 알 수 없는 인증 기관을 무시합니다.
* **조직**: 특정 조직에 관련된 Snyk 명령을 실행하려면 ORG\_ID를 지정합니다. Snyk는 ORG\_ID를 사용할 것을 권장합니다. ORG\_NAME(조직 슬러그 이름)을 지정하는 경우, 값은 Snyk UI에서 표시되는 URL 슬러그 이름과 일치해야 합니다: `https://app.snyk.io/org/[orgslugname` 으로 표시됩니다. 지정되지 않으면 테스트를 실행하는 데 사용되는 기본 조직은 [계정 설정](https://app.snyk.io/account)에서 정의된 기본 조직입니다.
* **사용 현황 분석 보내기**: Snyk 가 확장 기능이 작동하는 방식에 대한 정보를 보내도록 Visual Studio에 허용합니다.
* **프로젝트 설정**: 추가적인 Snyk CLI 매개변수를 지정합니다.\
  .NET 프로젝트의 경우 Snyk는 `--all-projects` 추가 매개변수를 권장합니다.
* **모든 프로젝트 스캔**: 작업 디렉토리의 모든 프로젝트를 자동으로 감지하는 기능은 기본적으로 활성화되어 있습니다.
*   **실행 파일 설정**: 플러그인을 통해 CLI를 다운로드하지 않고 자체 CLI 설치를 사용할 수 있습니다.

    * **필요한 바이너리 자동 관리** 확인란이 선택된 경우, 플러그인은 CLI를 자동으로 다운로드하고 업데이트합니다.
    * **필요한 바이너리 자동 관리** 확인란 선택이 해제된 경우, CLI의 유효한 경로를 제공해야 합니다. 네트워크 구성 또는 방화벽 규칙으로 인해 CLI 다운로드가 불가능한 경우 등 다른 방법으로 CLI를 얻어야 하는 경우 이 옵션을 사용하십시오. Snyk는 항상 CLI의 최신 버전을 사용하는 것을 권장합니다.

    <figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252Fz3kifJ2sVpr9s21o4hcL%252Fimage%2520%2810%29.png%3Falt%3Dmedia%26token%3Dffae042e-1a88-4681-82d7-e8d9e5ac4e25&#x26;width=768&#x26;dpr=1&#x26;quality=100&#x26;sign=dd9edd98&#x26;sv=2" alt=""><figcaption><p>VS Code 확장 실행 파일 설정</p></figcaption></figure>
* **솔루션 설정**: 오픈 소스 스캔을 위한 `snyk test` [CLI 옵션](../../../snyk-cli/commands/test.md)을 추가로 설정합니다. **비관리** [**C/C++**](../../../supported-languages-package-managers-and-frameworks/c-c++/) **스캔** 에서는 오픈 소스 패키지의 취약점을 찾기 위해 `--unmanaged` CLI 옵션을 사용합니다. 이를 사용하려면 **모든 프로젝트 스캔**을 비활성화해야 합니다. `--unmanaged` 옵션은 비관리 C/C++ 스캔에 대해만 작동하며 다른 언어에는 사용하지 마십시오. 추가 매개변수는 또는 IaC에는 적용되지 않습니다.

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252FfwkQuYWbgFO96tfQ6x5x%252FVS_Oprions_Unmagaed.jpg%3Falt%3Dmedia%26token%3De283e0aa-dbf7-4054-9fdb-15b5e62be692&#x26;width=768&#x26;dpr=1&#x26;quality=100&#x26;sign=6d938860&#x26;sv=2" alt="VS 확장 솔루션 설정 --unmanaged"><figcaption><p>VS 확장 솔루션 설정 --unmanaged</p></figcaption></figure>

## Visual Studio 확장 결과를 위한 제품 선택

설정에서 다음 결과를 받고 싶은지 선택할 수 있습니다:

* 오픈 소스 취약점
* 보안 취약점
* 품질 문제

{% hint style="info" %}
2025년 6월 24일부터 품질 문제는 더 이상 제공되지 않을 것입니다.
{% endhint %}
