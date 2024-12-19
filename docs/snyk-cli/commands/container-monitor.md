# 컨테이너 모니터

## 사용법

`snyk container monitor [<OPTIONS>] [<IMAGE>]`

## 설명

`snyk container monitor` 명령은 프로젝트의 컨테이너 이미지 레이어와 종속성을 캡처하고 해당 스냅샷을 취약점을 모니터링하며 결과를 [snyk.io](https://snyk.io)로 전송합니다.

코드를 프로덕션에 통합하기 전에 `container monitor` 명령을 사용하여 코드의 모니터링할 스냅샷을 가져가서 프로덕션으로 취약점을 전파하는 것을 피할 수 있습니다. 기본으로는 매일인 빈도를 변경하려면 설정에서 테스트 빈도를 선택하세요.

코드를 변경한 경우 `container monitor` 명령을 다시 실행해야 합니다.

자세한 정보는 [컨테이너 보안을 위한 Snyk CLI](https://docs.snyk.io/products/snyk-container/snyk-cli-for-container-security)을 참조하세요.

## 종료 코드

가능한 종료 코드와 그 의미:

**0**: 성공, 이미지 레이어와 종속성이 캡처됨\
**2**: 실패, 명령을 다시 실행해보세요. 디버그 로그를 출력하려면 `-d`를 사용하세요.\
**3**: 실패, 지원되는 프로젝트를 감지하지 못함

## Snyk CLI 구성

환경 변수를 사용하여 Snyk CLI를 구성하고 Snyk API와 연결하는 변수를 설정할 수 있습니다.

컨테이너 명령에 적용되는 환경 변수에 대해 자세히 알아보려면 [Snyk CLI 구성](https://docs.snyk.io/features/snyk-cli/configure-the-snyk-cli)을 확인하세요.

## 디버그

디버그 로그를 출력하려면 `-d` 옵션을 사용하세요.

## 옵션

### `--org=<ORG_ID>`

특정 Snyk 조직에 묶인 Snyk 명령을 실행하려면 `<ORG_ID>`를 지정하세요. `<ORG_ID>`는 일부 기능의 가용성과 개인 테스트 한도에 영향을 줍니다.

여러 조직이 있는 경우 CLI에서 다음을 사용하여 기본값을 설정할 수 있습니다.

`$ snyk config set org=<ORG_ID>`

기본값: 현재 선호하는 조직의 `<ORG_ID>`는 [계정 설정](https://app.snyk.io/account)에서 설정된 것입니다.

**참고:** `--org=<orgslugname>`도 사용할 수 있습니다. `ORG_ID`는 CLI 및 API에서 모두 작동합니다. 조직 슬러그 이름은 CLI에서 작동하지만 API에서는 작동하지 않습니다.

`orgslugname`은 Snyk UI의 조직 URL에 표시된 슬러그 이름과 일치해야 합니다: `https://app.snyk.io/org/[orgslugname]`. orgname은 작동하지 않습니다.

자세한 내용은 [CLI에서 사용할 조직 선택하는 방법](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/how-to-select-the-organization-to-use-in-the-cli)을 참조하세요.

### `--file=<FILE_PATH>`

상세한 조언을 위해 이미지의 Dockerfile 경로를 포함시킵니다.

### `--project-name=<PROJECT_NAME>`

사용자 지정 Snyk 프로젝트 이름을 지정하세요.

### `--policy-path=<PATH_TO_POLICY_FILE>`

수동으로 `.snyk` 정책 파일의 경로를 전달하세요.

### `--json`

콘솔에 결과를 JSON 데이터 구조로 인쇄합니다.

예: `$ snyk container test --json`

**참고:** 역할이 프로젝트 속성을 편집할 권한이 없는 경우 `monitor` 명령이 실패합니다. 계속 진행하는 방법은 [Snyk CLI에서 프로젝트 속성을 편집하려는 데 필요한 권한(역할)](https://docs.snyk.io/snyk-admin/manage-permissions-and-roles/manage-member-roles#permissions-role-required-to-edit-project-attributes-from-the-snyk-cli)을 참조하세요.

### `--target-reference=<TARGET_REFERENCE>`

이 프로젝트를 구별하는 참조를 지정하세요. 예: 브랜치 이름이나 버전. 동일한 참조를 가진 프로젝트를 해당 참조를 기준으로 그룹화할 수 있습니다.

더 많은 정보는 [모니터링을 위해 브랜치 또는 버전별 프로젝트 그룹화](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/group-projects-by-branch-or-version-for-monitoring)를 참조하세요.

### `--project-environment=<ENVIRONMENT>[,<ENVIRONMENT>]...>`

프로젝트 환경을 하나 이상의 값으로 설정하세요(쉼표로 구분). 프로젝트 환경을 지우려면 `--project-environment=`으로 설정하세요.

허용되는 값: `frontend`, `backend`, `internal`, `external`, `mobile`, `saas`, `onprem`, `hosted`, `distributed`

더 많은 정보는 [프로젝트 속성](https://docs.snyk.io/snyk-admin/snyk-projects/project-attributes)를 참조하세요.

### `--project-lifecycle=<LIFECYCLE>[,<LIFECYCLE]...>`

프로젝트 수명주기를 하나 이상의 값으로 설정하세요(쉼표로 구분). 프로젝트 수명주기를 지우려면 `--project-lifecycle=`으로 설정하세요.

허용되는 값: `production, development, sandbox`

더 많은 정보는 [프로젝트 속성](https://docs.snyk.io/snyk-admin/snyk-projects/project-attributes)를 참조하세요.

### `--project-business-criticality=<BUSINESS_CRITICALITY>[,<BUSINESS_CRITICALITY>]...>`

프로젝트 비즈니스 중요도를 하나 이상의 값으로 설정하세요(쉼표로 구분). 프로젝트 비즈니스 중요도를 지우려면 `--project-business-criticality=`으로 설정하세요.

허용되는 값: `critical`, `high`, `medium`, `low`

더 많은 정보는 [프로젝트 속성](https://docs.snyk.io/snyk-admin/snyk-projects/project-attributes)를 참조하세요.

### `--project-tags=<TAG>[,<TAG>]...>`

프로젝트 태그를 하나 이상의 값으로 설정하세요(등호 구분된 쉼표로 구분된 키-값 쌍).

예: `--project-tags=department=finance,team=alpha`

프로젝트 태그를 지우려면 `--project-tags=`으로 설정하세요.

허용되는 문자를 포함한 자세한 정보는 [프로젝트 태그](https://docs.snyk.io/snyk-admin/snyk-projects/project-tags)를 참조하세요.

### `--tags=<TAG>[,<TAG>]...>`

`--project-tags`의 별칭입니다.

### `--app-vulns`

컨테이너 이미지로부터 어플리케이션 종속성뿐만 아니라 운영 체제에서도 취약점을 감지할 수 있도록 허용합니다. CLI 버전 1.1090.0 (2023년 01월 24일) 이상에서는 Snyk가 기본적으로 이미지의 어플리케이션 종속성을 스캔하므로 `--app-vulns` 옵션을 명시할 필요가 없습니다.

CLI 버전 1.962.0부터 v1.1089.0까지는 결과에서 운영 체제 및 어플리케이션 취약점을 JSON 형식으로 보려면 `--app-vulns` 옵션을 `--json` 옵션과 함께 사용하세요.

자세한 정보는 [컨테이너 이미지에서 어플리케이션 취약점 감지](https://docs.snyk.io/scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/detect-application-vulnerabilities-in-container-images)를 참조하세요.

### `--exclude-app-vulns`

어플리케이션 취약점을 스캔하지 않도록 허용합니다. CLI 버전 1.1090.0 (2023년 01월 24일) 이상에서는 `app-vulns`가 기본적으로 활성화됩니다.

이전 릴리스에서는 `--exclude-app-vulns`와 함께 사용할 수 없습니다.

자세한 정보는 [컨테이너 이미지에서 어플리케이션 취약점 감지](https://docs.snyk.io/scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/detect-application-vulnerabilities-in-container-images)를 참조하세요.

### `--nested-jars-depth`

`app-vulns`가 활성화된 경우, `--nested-jars-depth=n` 옵션을 사용하여 Snyk가 언팩해야 하는 중첩된 jar의 수준을 설정하세요. 깊이는 숫자여야 합니다.

### `--exclude-base-image-vulns`

기본 이미지에서만 도입된 취약점을 표시하지 않습니다. 운영 체제 패키지에 대해서만 작동합니다. `snyk container test`만을 사용하는 경우에만 사용할 수 있습니다. `monitor`와 함께 사용하면 효과가 없습니다.

### `--platform=<PLATFORM>`

다중 아키텍처 이미지의 경우 테스트할 플랫폼을 지정하세요.

지원되는 플랫폼: `linux/amd64`, `linux/arm64`, `linux/riscv64`, `linux/ppc64le`, `linux/s390x`, `linux/386`, `linux/arm/v7`, 또는 `linux/arm/v6`

### `--username=<CONTAINER_REGISTRY_USERNAME>`

컨테이너 레지스트리에 연결할 때 사용할 사용자 이름을 지정하세요. 도커가 있을 때는 로컬 도커 바이너리 자격 증명을 우선합니다.

### `--password=<CONTAINER_REGISTRY_PASSWORD>`

컨테이너 레지스트리에 연결할 때 사용할 비밀번호를 지정하세요. 도커가 있을 때는 로컬 도커 바이너리 자격 증명을 우선합니다.

## 컨테이너 모니터 명령 예시

**도커 이미지 스캔 및 모니터**

`$ snyk container monitor <image>`