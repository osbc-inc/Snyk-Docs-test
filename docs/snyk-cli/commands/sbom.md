---
description: 로컬 파일 시스템에서 SBOM 문서를 생성합니다.
---

# SBOM

## 전제 조건

**기능 가용성:** 이 기능은 Snyk 엔터프라이즈 플랜을 이용하는 고객에게 제공됩니다.

**참고:** SBOM 생성 기능을 사용하기 위해서는 최소 CLI 버전 1.1071.0을 사용해야 합니다.

`snyk sbom` 기능은 인터넷 연결이 필요합니다.

## 사용법

`$ snyk sbom --format=<cyclonedx1.4+json|cyclonedx1.4+xml|cyclonedx1.5+json|cyclonedx1.5+xml|cyclonedx1.6+json|cyclonedx1.6+xml|spdx2.3+json> [--org=<ORG_ID>] [--file=<FILE>] [--unmanaged] [--dev] [--all-projects] [--name=<NAME>] [--version=<VERSION>] [--exclude=<NAME>[,<NAME>...]] [--detection-depth=<DEPTH>] [--prune-repeated-subdependencies|-p] [--maven-aggregate-project] [--scan-unmanaged] [--scan-all-unmanaged] [--sub-project=<NAME>] [--gradle-sub-project=<NAME>] [--all-sub-projects] [--configuration-matching=<CONFIGURATION_REGEX>] [--configuration-attributes=<ATTRIBUTE>[,<ATTRIBUTE>]] [--init-script=<FILE>] [--json-file-output=<OUTPUT_FILE_PATH>] [<TARGET_DIRECTORY>]`

## 설명

`snyk sbom` 명령은 Snyk에서 지원하는 생태계 내의 로컬 소프트웨어 프로젝트에 대한 SBOM을 생성합니다.

지원되는 형식은 CycloneDX v1.4 (JSON 또는 XML), CycloneDX v1.5 (JSON 또는 XML), CycloneDX v1.6 (JSON 또는 XML), 그리고 SPDX v2.3 (JSON)이 있습니다.

SBOM은 지원되는 모든 오픈 소스 패키지 관리자 및 관리되지 않는 소프트웨어 프로젝트에 대해 생성할 수 있습니다.

## 종료 코드

가능한 종료 코드와 그 의미:

**0**: 성공 (프로세스 완료), SBOM이 성공적으로 생성됨\
**2**: 실패, 명령 다시 실행해 보세요. 디버그 로그를 출력하려면 `-d`를 사용하세요.

## 디버그

디버그 로그를 출력하려면 `-d` 또는 `--debug` 옵션을 사용하세요.

## 옵션

### `--format=<cyclonedx1.4+json|cyclonedx1.4+xml|cyclonedx1.5+json|cyclonedx1.5+xml|cyclonedx1.6+json|cyclonedx1.6+xml|spdx2.3+json>`

필수 항목입니다. 생성될 SBOM의 출력 형식을 지정합니다.

원하는 SBOM 출력 형식을 설정합니다. 사용 가능한 옵션은 `cyclonedx1.4+json`, `cyclonedx1.4+xml`, `cyclonedx1.5+json`, `cyclonedx1.5+xml`, `cyclonedx1.6+json`, `cyclonedx1.6+xml`, `spdx2.3+json`입니다.

### `[--org=<ORG_ID>]`

특정 Snyk 조직과 관련된 Snyk 명령을 실행하려면 `<ORG_ID>` (이름 또는 UUID)를 지정합니다. `<ORG_ID>`는 일부 기능 가용성과 개인 테스트 한도에 영향을 줍니다.

이 옵션은 기본 조직에서 API 권한이 없는 경우 사용하세요.

생략된 경우 계정의 기본 조직이 사용됩니다.

큰 회사용 [계정 설정](https://app.snyk.io/account)에서 현재 기본 조직인 `<ORG_ID>`를 설정해주세요.

모든 새로 테스트하는 프로젝트가 기본 조직 하에 테스트되도록 하려면 기본값을 설정하세요. 기본값을 재정의해야 하는 경우 `--org=<ORG_ID>` 옵션을 사용하세요.

여러 조직을 가지고 있는 경우 CLI를 사용하여 기본값을 설정할 수 있습니다:

`$ snyk config set org=<ORG_ID>`

**참고:** `--org=<orgslugname>`도 사용할 수 있습니다. `ORG_ID`는 CLI 및 API에서 작동하며, 조직 슬러그 이름은 CLI에서만 작동합니다. API에서는 작동하지 않습니다.

자세한 정보는 [CLI에서 사용할 조직 선택하는 방법](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/how-to-select-the-organization-to-use-in-the-cli) 문서를 참조하세요.

### `[--file=<file>]` or `[--f=<file>]`

SBOM의 기반이 될 원하는 매니페스트 파일을 지정합니다.

기본적으로 `sbom` 명령은 현재 작업 디렉토리에 있는 지원되는 매니페스트 파일을 감지합니다.

### `[--unmanaged]`

관리되지 않는 소프트웨어 프로젝트용 SBOM 생성합니다.

### `[--dev]`

SBOM 출력에 개발용 종속성을 포함합니다.

일부 패키지 관리자에만 해당되며, 예를 들어 npm의 `devDependencies`나 Gemfile의 `:development` 종속성과 같은 것입니다.

SPDX 형식과 함께 `--dev`를 사용하면 개발 전용 종속성이 `DEV_DEPENDENCY_OF` 관계에 포함됩니다.

CycloneDX 형식과 함께 `--dev`를 사용하면 개발 전용 종속성이 비개발 종속성과 구별되지 않습니다.

**참고:** 이 옵션은 Maven, npm, Yarn 프로젝트에서 사용할 수 있습니다.