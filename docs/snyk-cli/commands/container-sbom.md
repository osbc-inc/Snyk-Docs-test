# 컨테이너 SBOM

## 전제 조건

**기능 사용 가능 여부:** 이 기능은 현재 초기 액세스 중이며 Snyk 엔터프라이즈 요금제를 사용하는 고객에게 제공됩니다.

**참고:** SBOM 생성 기능을 사용하려면 최소 CLI 버전 1.1226.0을 사용해야 합니다.

`snyk container sbom` 기능을 사용하려면 인터넷 연결이 필요합니다.

## 사용법

`$ snyk container sbom --format=<cyclonedx1.4+json|cyclonedx1.4+xml|cyclonedx1.5+json|cyclonedx1.5+xml|cyclonedx1.6+json|cyclonedx1.6+xml|spdx2.3+json> [--org=<ORG_ID>] [--exclude-app-vulns] <IMAGE>`

## 설명

`snyk container sbom` 명령어는 컨테이너 이미지를 위한 SBOM을 생성합니다.

지원되는 형식에는 CycloneDX v1.4 (JSON 또는 XML), CycloneDX v1.5 (JSON 또는 XML), CycloneDX v1.6 (JSON 또는 XML), 그리고 SPDX v2.3 (JSON)이 포함됩니다.

SBOM은 이미지 내의 운영 체제 종속성 및 애플리케이션 종속성을 위해 생성될 수 있습니다. 관리되지 않는 종속성은 지원되지 않습니다.

## 종료 코드

가능한 종료 코드 및 의미:

**0**: 성공 (프로세스가 완료됨), SBOM이 성공적으로 생성됨\
**2**: 실패, 명령어를 재실행하십시오. 디버그 로그를 보려면 `-d`를 사용하십시오.

## 디버깅

디버그 로그를 확인하려면 `-d` 또는 `--debug` 옵션을 사용하십시오.

## 옵션

### `--format=<cyclonedx1.4+json|cyclonedx1.4+xml|cyclonedx1.5+json|cyclonedx1.5+xml|cyclonedx1.6+json|cyclonedx1.6+xml|spdx2.3+json>`

필수 항목입니다. 생성될 SBOM의 출력 형식을 지정합니다.

생성될 SBOM의 출력 형식을 설정합니다. 사용 가능한 옵션은 `cyclonedx1.4+json`, `cyclonedx1.4+xml`, `cyclonedx1.5+json`, `cyclonedx1.5+xml`, `cyclonedx1.6+json`, `cyclonedx1.6+xml`, `spdx2.3+json`입니다.

### `[--org=<ORG_ID>]`

특정 Snyk 조직에 연결된 Snyk 명령을 실행하기 위해 `<ORG_ID>`(이름 또는 UUID)를 지정합니다. `<ORG_ID>`는 일부 기능의 가용성 및 개인 테스트 한도에 영향을 줍니다.

기본 조직에서 API 권한이 없는 경우 이 옵션을 사용합니다.

이 옵션을 생략하면 계정의 기본 조직이 사용됩니다.

이것은 [계정 설정](https://app.snyk.io/account)에서 현재 기본 조직인 `<ORG_ID>`입니다.

모든 새로 테스트된 프로젝트가 기본 조직 하에 테스트되도록 하려면 기본값을 설정하십시오. 기본값을 무시하려면 `--org=<ORG_ID>` 옵션을 사용하십시오.

여러 가지 조직을 보유한 경우 CLI를 통해 기본값을 설정할 수 있습니다:

`$ snyk config set org=<ORG_ID>`

**참고:** `--org=<orgslugname>`를 사용할 수도 있습니다. `ORG_ID`는 CLI 및 API에서 모두 작동합니다. 조직 슬러그 이름은 CLI에서는 작동하지만 API에서는 작동하지 않습니다.

자세한 내용은 [CLI에서 사용할 조직 선택하는 방법](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/how-to-select-the-organization-to-use-in-the-cli) 문서를 참조하십시오.

### `[--exclude-app-vulns]`

기본적으로 Snyk는 이미지 내의 운영 체제 종속성 및 애플리케이션 종속성을 스캔하고 SBOM을 생성합니다.

`--exclude-app-vulns`를 추가하여 애플리케이션 종속성의 생성을 비활성화할 수 있습니다.

애플리케이션 스캔에 대한 자세한 정보는 [컨테이너 이미지에서 애플리케이션 취약점 감지하기](https://docs.snyk.io/scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/detect-application-vulnerabilities-in-container-images)를 참조하십시오.

### `<IMAGE>`

필수 항목입니다. SBOM 문서를 생성할 이미지입니다.

**참고:** 이미지는 `repo:tag` 형식으로 지정해야 합니다.

## snyk container sbom 명령어에 대한 예시

### 이미지를 위해 CycloneDX JSON 문서 생성하기

`$ snyk container sbom --format=cyclonedx1.6+json redis:latest`

### 이미지를 위해 CycloneDX JSON 문서 생성하고 표준 출력을 파일로 리디렉션하기

`$ snyk container sbom --format=cyclonedx1.6+json redis:latest > mySBOM.json`

### 애플리케이션 종속성을 제외하고 이미지를 위해 SPDX JSON 문서 생성하기

`$ snyk container sbom --format=spdx2.3+json redis:latest --exclude-app-vulns`

### 다이제스트를 통해 컨테이너 이미지 참조하기

`$ snyk container sbom --format=cyclonedx1.6+xml alpine@sha256:c5c5fda71656f28e49ac9c5416b3643eaa6a108a8093151d6d1afc9463be8e33`