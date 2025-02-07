# Eclipse 플러그인 구성

Snyk 환경설정에서 다음 구성 옵션을 설정할 수 있습니다.

<figure><img src="../../../.gitbook/assets/image (2) (12).png" alt=""><figcaption><p>Snyk 환경설정</p></figcaption></figure>

* `Snyk API 토큰`: Snyk에서 인증 토큰을 설정합니다.
* `경로`: Gradle 또는 Maven과 같은 필요한 서드파티 도구를 찾기 위한 경로 추가를 지정합니다.
* `사용자 정의 엔드포인트`: 사용자 지정 멀티테넌시 또는 싱글테넌시 설정을 위한 Snyk API 엔드포인트를 지정합니다. `https://api.snyk.io`를 사용하는 경우 별도의 구성이 필요하지 않습니다. 자세한 내용은 [IDE URL 목록](../../../working-with-snyk/regional-hosting-and-data-residency.md#ides-urls)을 참조하세요.
* `알려지지 않은 인증 권한 허용`: SSL 연결을 위한 인증 확인을 비활성화합니다.
* `Snyk 오픈 소스 활성화됨`: 랭귀지 서버를 통해 Snyk 오픈 소스 종속성 스캔을 활성화 또는 비활성화합니다. 기본값: 베타 동안 `활성화` 상태
* `Snyk 코드 보안활성화됨`: 랭귀지 서버를 통해 Snyk 코드 스캔을 활성화 또는 비활성화합니다. 기본값: 베타 동안 `비활성화` 상태
* `Snyk Infrastructure-as-Code 활성화됨`: 랭귀지 서버를 통해 Snyk IaC 스캔을 활성화 또는 비활성화합니다. 기본값: 베타 동안 `활성화` 상태
* `시작시 자동으로 스캔 및 저장`: 자동으로 스캔할지 여부
* `조직`: 스캔에 사용할 Snyk 조직을 지정합니다. Snyk는 `ORG_ID`를 사용하도록 권장합니다. 조직 슬러그 이름을 지정하는 경우, 값은 Snyk UI의 조직 URL 슬러그 이름 (`[orgslugname]`)과 일치해야 합니다: `https://app.snyk.io/org/[orgslugname]`.
* `추가 매개변수`: CLI로 전달할 추가 매개변수를 지정합니다. 예를 들어, `.NET Projects`의 경우 Snyk는 `--all-projects` 추가 매개변수를 추가하도록 권장합니다.
* `추가 환경 변수`: 랭귀지 서버에 환경 변수를 추가합니다. 여러 변수는 `;`로 구분할 수 있습니다. 예: `JAVA_HOME=/usr/local/bin;GOPATH=/usr/local/bin`
* `Snyk 이진 파일 자동 업데이트 및 설치`: 비활성화된 경우 업데이트가 다운로드되지 않으며, 업데이트를 수동으로 수행해야 합니다. Snyk는 항상 최신 버전의 CLI를 사용하도록 권장합니다. CLI의 위치가 현재 이진 파일을 가리키도록 해야 합니다.
* `CLI 다운로드를 위한 기본 URL:` CLI를 위한 대체 다운로드 호스트를 지정합니다. 예: `https://downloads.snyk.io/fips`. 이는 다음 파일을 제공해야 합니다. GitHub [릴리스](https://github.com/snyk/cli/releases)도 참조하세요.
  * %Base URL%/cli/v%VERSION%/%CLI-BINARY-NAME%
  * %Base URL%/cli/v%VERSION%/%CLI-BINARY-NAME%.sha256
  * %Base URL%/cli/v%VERSION%/sha256sums.txt.asc
  * %Base URL%/cli/v%VERSION%/release.json
  * %Base URL%/cli/stable/version
  * %Base URL%/cli/stable/%CLI-BINARY-NAME%
  * %Base URL%/cli/stable/%CLI-BINARY-NAME%.sha256
  * %Base URL%/cli/stable/ls-protocol-version-%PROTOCOL\_VERSION%
  * %Base URL%/cli/stable/release.json
  * %Base URL%/cli/stable/sha256sums.txt.asc
* `Snyk CLI`: Snyk CLI의 위치를 지정합니다.
* `에러 보고서를 Snyk로 보내기`: 랭귀지 서버에서 발생한 오류를 빠른 버그 수정을 위해 Snyk로 보냅니다. 기본값: `활성화`.
* `사용 통계를 Snyk로 보내기`: 워크플로를 개선하기 위해 사용 데이터를 Snyk에 제공합니다. 기본값: `활성화`.
* `신뢰할 수 있는 폴더`: 안전한 고려해야 하는 디렉터리를 지정합니다. 예: 모든 프로젝트의 상위 디렉토리.
