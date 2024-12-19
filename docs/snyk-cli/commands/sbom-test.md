# SBOM 테스트

**기능 이용 가능 여부:** 이 기능은 Snyk 엔터프라이즈 요금제의 고객에게 제공됩니다.

## 사용법

`snyk sbom test --experimental --file=<파일_경로> [<옵션>]`

## 설명

`snyk sbom test` 명령어는 오픈 소스 패키지에서의 취약점을 검사하기 위해 SBOM 파일을 확인합니다.

## 종료 코드

가능한 종료 코드 및 의미:

**0**: 성공 (스캔 완료), 취약점 없음\
**1**: 조치 필요 (스캔 완료), 취약점 발견\
**2**: 실패, 명령어 다시 실행 시도

## Snyk CLI 구성

환경 변수를 사용하여 Snyk CLI를 구성하고 Snyk API와 연결하기 위한 변수를 설정할 수 있습니다. [Snyk CLI 구성](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli) 참조

## 디버그

디버그 로그를 출력하려면 `-d` 또는 `--debug` 옵션을 사용하세요.

## 옵션

### `--experimental`

필수. 실험적인 명령어 기능을 사용합니다. 현재 명령어가 실험 단계에 있기 때문에 이 옵션이 필요합니다.

### `--file=<파일_경로>`

필수. SBOM 문서의 파일 경로를 지정합니다.

`snyk sbom test` 명령어는 다음 파일 형식을 허용합니다:

* **CycloneDX:** JSON 버전 1.4, 1.5 및 1.6
* **SPDX:** JSON 버전 2.3

제공된 SBOM 파일 내의 패키지 및 구성 요소는 PackageURL (purl)을 통해 식별되어야 합니다.

지원되는 purl 유형은: `apk`, `cargo`, `cocoapods`, `composer`, `deb`, `gem`, `generic`, `golang`, `hex`, `maven`, `npm`, `nuget`, `pub`, `pypi`, `rpm`, `swift`입니다.

예: `$ snyk sbom test --experimental --file=bom.cdx.json`

### `--json`

콘솔에 결과를 JSON 데이터 구조로 출력합니다.

예: `$ snyk sbom test --experimental --file=bom.cdx.json --json`