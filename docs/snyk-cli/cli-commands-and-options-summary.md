# CLI 명령 및 옵션 요약

{% hint style="info" %}
이 페이지는 **CLI 명령**과 각 명령에 대한 옵션을 요약합니다. 자세한 내용은 이 요약 안에 있는 링크를 사용하여 사용 중인 명령어의 도움링크 문서를 열어주세요. 도움링크 문서는 CLI의 도움링크와 동일합니다.
{% endhint %}

## 사용법

`snyk [COMMAND] [SUBCOMMAND] [OPTIONS] [PACKAGE] [CONTEXT-SPECIFIC-OPTIONS]`

## 설명

Snyk CLI는 프로젝트에서 알려진 취약점을 찾아 해결하는 빌드 시간 도구입니다. 더 자세한 Snyk CLI 및 Snyk 설명은 [Snyk CLI](https://docs.snyk.io/snyk-cli)를 참조하세요. Snyk CLI 사용 방법에 대한 소개는 [CLI로 시작하기](https://docs.snyk.io/snyk-cli/getting-started-with-the-cli)에서 확인할 수 있습니다.

## 사용 가능한 CLI 명령

각 Snyk CLI 명령에 대한 자세한 정보는 `--help` 옵션을 사용하여 확인하십시오. 예를 들어, `snyk auth --help` 또는 `snyk container --help`와 같이 사용합니다. 이 목록의 각 명령은 해당 문서에서 도움링크 페이지로 연결됩니다.

**참고:** Snyk CLI 명령에 대한 옵션 목록이 이 페이지에 있습니다. 각 명령에 대한 도움링크에서 옵션에 대한 자세한 내용이 설명되어 있습니다.

### [`snyk auth`](https://docs.snyk.io/snyk-cli/commands/auth)

Snyk 계정으로 Snyk CLI를 인증합니다.

### [`snyk config`](https://docs.snyk.io/snyk-cli/commands/config)

Snyk CLI 구성을 관리합니다.

### [`snyk config environment`](https://docs.snyk.io/snyk-cli/commands/config-environment)

`--snyk auth` 명령을 실행하기 전에 지역 환경을 설정하는 데 사용합니다.

### [`snyk test`](https://docs.snyk.io/snyk-cli/commands/test)

오픈 소스 취약점 및 라이선스 문제를 위해 프로젝트를 테스트합니다.

### [`snyk monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

프로젝트를 스냅샷으로 만들고 오픈 소스 취약점 및 라이선스 문제를 계속해서 모니터링합니다.

### [`snyk code`](https://docs.snyk.io/snyk-cli/commands/code)

`snyk code` 명령의 이름을 확인합니다. 도움 옵션과 함께 사용하는 방법: `snyk code test`

### [`snyk code test`](https://docs.snyk.io/snyk-cli/commands/code-test)

소스 코드를 검사하여 알려진 보안 문제를 확인합니다(정적 응용 프로그램 보안 테스트).

### [`snyk container`](https://docs.snyk.io/snyk-cli/commands/container)

`snyk container` 명령, `snyk container monitor`, `snyk container test`를 출력합니다.

### [`snyk container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

컨테이너 이미지 계층과 종속성을 캡처하고 [snyk.io](https://snyk.io)에서 취약점을 모니터링합니다.

### [`snyk container SBOM`](https://docs.snyk.io/snyk-cli/commands/container-sbom)

컨테이너 이미지에 대한 SBOM을 생성합니다.

### [`snyk container test`](https://docs.snyk.io/snyk-cli/commands/container-test)

컨테이너 이미지를 알려진 취약점으로 테스트합니다.

### [`snyk iac`](https://docs.snyk.io/snyk-cli/commands/iac)

`snyk iac` 명령 목록을 출력합니다: `snyk iac describe`, `snyk iac update-exclude-policy`, `snyk iac test`.

### [`snyk iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

알려진 보안 문제를 확인합니다.

### [`snyk iac capture`](https://docs.snyk.io/snyk-cli/commands/iac-capture)

코드에서 Terraform 상태 파일로부터 리소스 매핑을 생성하는 데 필요한 최소한의 정보를 포함한 매핑 아티팩트를 생성하고 해당 매핑 아티팩트를 Snyk에 보냅니다.

### [`snyk iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

인프라 드리프트와 관리되지 않는 리소스를 감지하고 추적 및 경고를 합니다.

### [`snyk iac rules init`](https://docs.snyk.io/snyk-cli/commands/iac-rules-init)

사용자 정의 규칙 프로젝트 구조, 관계, 규칙 또는 사양을 초기화합니다.

### [`snyk iac rules test`](https://docs.snyk.io/snyk-cli/commands/iac-rules-test)

모든 사용자 정의 규칙에 대한 테스트를 실행합니다.

### [`snyk iac rules push`](https://docs.snyk.io/snyk-cli/commands/iac-rules-push)

사용자 정의 규칙 번들을 묶어서 Snyk 클라우드 API에 업로드합니다.

### [`snyk iac update-exclude-policy`](https://docs.snyk.io/snyk-cli/commands/iac-update-exclude-policy)

`snyk iac describe`에서 사용할 제외 정책 규칙을 생성합니다.

### [`snyk ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

지정된 문제를 무시하기 위해 `.snyk` 정책을 수정합니다.

### [`snyk log4shell`](https://docs.snyk.io/snyk-cli/commands/log4shell)

Log4Shell 취약점을 찾습니다.

### [`snyk policy`](https://docs.snyk.io/snyk-cli/commands/policy)

패키지에 대한 `.snyk` 정책을 표시합니다.

### [`snyk sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

Snyk에서 지원하는 생태계 내 소프트웨어 프로젝트에 대한 SBOM을 생성합니다.

### [`snyk sbom test`](https://docs.snyk.io/snyk-cli/commands/sbom-test)

오픈 소스 패키지의 취약점을 확인하는 SBOM을 확인합니다.

### [`snyk apps`](https://docs.snyk.io/snyk-cli/create-a-snyk-app-using-the-snyk-cli)

Snyk CLI를 사용하여 Snyk 앱을 만듭니다. 자세한 내용은 [Snyk 앱](../snyk-api/how-to-use-snyk-apps-apis/)을 참조하세요.

## CLI 명령의 하위 명령

다음은 Snyk CLI 명령의 하위 명령 목록입니다. 각 하위 명령 뒤에 해당되는 명령이 나옵니다. 명령어는 도움 문서로 연결됩니다. 각 하위 명령에 대한 자세한 정보는 해당 도움 문서를 참조하세요.

`get <KEY>`: [`config`](https://docs.snyk.io/snyk-cli/commands/config)의 하위 명령

`set <KEY>=<VALUE>`: [`config`](https://docs.snyk.io/snyk-cli/commands/config)의 하위 명령

`unset <KEY>`: [`config`](https://docs.snyk.io/snyk-cli/commands/config)의 하위 명령

`clear`: [`config`](https://docs.snyk.io/snyk-cli/commands/config)의 하위 명령

`environment`: [`config`](https://docs.snyk.io/snyk-cli/commands/config)의 하위 명령

## Snyk CLI 구성

환경 변수를 사용하여 Snyk CLI를 구성하고 Snyk CLI가 Snyk API에 연결할 수 있도록 변수를 설정할 수 있습니다. [Snyk CLI 구성](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli)을 참조하세요.

## 디버깅

자세한 내용은 [Snyk CLI 디버깅](debugging-the-snyk-cli.md)을 참조하세요.

## CLI 명령의 종료 코드

`test` 명령에 대한 종료 코드는 모두 동일합니다. 다음 도움 문서에서 종료 코드를 확인하세요:

- [`snyk test` 종료 코드](https://docs.snyk.io/snyk-cli/commands/test#exit-codes)
- [`snyk container test` 종료 코드](https://docs.snyk.io/snyk-cli/commands/container-test#exit-codes)
- [`snyk iac test` 종료 코드](https://docs.snyk.io/snyk-cli/commands/iac-test#exit-codes)
- [`snyk code test` 종료 코드](https://docs.snyk.io/snyk-cli/commands/code-test#exit-codes)

다음 도움 문서에 나열된 다른 CLI 명령에 대한 종료 코드가 있습니다:

- [`snyk monitor` 종료 코드](https://docs.snyk.io/snyk-cli/commands/container-monitor#exit-codes)
- [`snyk container monitor` 종료 코드](https://docs.snyk.io/snyk-cli/commands/container-monitor#exit-codes)
- [`snyk iac describe` 종료 코드](https://docs.snyk.io/snyk-cli/commands/iac-describe#exit-codes)
- [`snyk iac update-exclude-policy` 종료 코드](https://docs.snyk.io/snyk-cli/commands/iac-update-exclude-policy#exit-codes)
- [`snyk log4shell` 종료 코드](https://docs.snyk.io/snyk-cli/commands/log4shell#exit-codes)
- [`snyk sbom` 종료 코드](https://docs.snyk.io/snyk-cli/commands/sbom#exit-codes)
- [`snyk container sbom` 종료 코드](https://docs.snyk.io/snyk-cli/commands/container-sbom#exit-codes)

## 여러 명령의 옵션

Snyk CLI 명령의 옵션 목록이 옵션에 적용되는 명령 뒤에 나옵니다. 명령은 해당 도움 문서로 연결됩니다. 각 옵션에 대한 자세한 정보는 도움 문서를 참조하세요.

`--all-projects`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--fail-fast`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--detection-depth=<DEPTH>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac/iac-test), [`sbom`](commands/sbom.md)

`--exclude=<NAME>[,<NAME>]...>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](commands/sbom.md)

`--prune-repeated-subdependencies, -p`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](commands/sbom.md)

`--print-deps`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test)

`--remote-repo-url=<URL>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [iac test](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--dev`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](commands/sbom.md)

`--org=<ORG_ID>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [container monitor](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac/iac-test), [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe), [`iac capture`](https://docs.snyk.io/snyk-cli/commands/iac-capture), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom), [`container sbom`](https://docs.snyk.io/snyk-cli/commands/container-sbom)

`--file=<FILE>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--file=<FILE_PATH>`: [container test](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`sbom test`](https://docs.snyk.io/snyk-cli/commands/sbom-test)

`--package-manager=<PACKAGE_MANAGER_NAME>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--unmanaged:` [test](commands/test.md)`,` [monitor](commands/monitor.md). [사용하지 않는 리소스에 대한 스캔 옵션](https://docs.snyk.io/snyk-cli/cli-reference#options-for-scanning-using-unmanaged) 및 [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom) 명령어 도움 링크에서 이 옵션을 또 사용하는 방법을 참조하세요.

`--ignore-policy`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test), [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--trust-policies` [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--show-vulnerable-paths=<none|some|all>` [`test`](https://docs.snyk.io/snyk-cli/commands/test)

`--project-name=<PROJECT_NAME>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`containernyk.io/snyk-cli/commands/monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)`,`[`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

`--policy-path=<PATH_TO_POLICY_FILE>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test), [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe), [`ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

`--json`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test), [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe), [`sbom test`](https://docs.snyk.io/snyk-cli/commands/sbom-test)

`--json-file-output=<OUTPUT_FILE_PATH>`: [`test`](commands/test.md), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test),[`sbom`](https://docs.snyk.io/snyk-cli/commands/test)

`--sarif`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--sarif-file-output=<OUTPUT_FILE_PATH>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--severity-threshold=<low|medium|high|critical>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test), [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--fail-on=<all|upgradable|patchable>`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [test](https://docs.snyk.io/snyk-cli/commands/test)

`--project-environment=<ENVIRONMENT>[,<ENVIRONMENT>]...`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--project-lifecycle=<LIFECYCLE>[,<LIFECYCLE>]...`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--project-business-criticality=<BUSINESS_CRITICALITY>[,<BUSINESS_CRITICALITY>]...`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--project-tags=<TAG>[,<TAG>]...`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--tags=<TAG>[,<TAG>]...`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

## `snyk auth` 커맨드 옵션

\--`auth-type=<TYPE>`\
`--client-secret=<SECRET>`\
`--client-id=<ID>` [`snyk auth`](https://docs.snyk.io/snyk-cli/commands/auth)

## `snyk code test` 커맨드 옵션

`--include-ignores`: [`code test`](https://docs.snyk.io/snyk-cli/commands/code-test)

## `snyk config environment` 커맨드 옵션

`--no-check` [snyk config environment](https://docs.snyk.io/snyk-cli/commands/config-environment)

## `snyk container` 커맨드 옵션

`--app-vulns`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

`--exclude-app-vulns`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor), [`container sbom`](https://docs.snyk.io/snyk-cli/commands/container-sbom)

`--nested-jars-depth`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

`--exclude-base-image-vulns`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

`--platform=<PLATFORM>`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

`--username=<CONTAINER_REGISTRY_USERNAME>`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

`--password=<CONTAINER_REGISTRY_PASSWORD>`: [`container test`](https://docs.snyk.io/snyk-cli/commands/container-test), [`container monitor`](https://docs.snyk.io/snyk-cli/commands/container-monitor)

## `snyk iac test` 커맨드 옵션

`--scan=<TERRAFORM_PLAN_SCAN_MODE>`: [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--target-name=<TARGET_NAME>`: [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--rules=<PATH_TO_CUSTOM_RULES_BUNDLE>`: [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--var-file=<PATH_TO_VARIABLE_FILE>`: [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

`--report`: [`iac test`](https://docs.snyk.io/snyk-cli/commands/iac-test)

## `snyk iac capture` 커맨드 옵션

`--stdin`: [`iac capture`](https://docs.snyk.io/snyk-cli/commands/iac-capture)

`PATH`: [`iac capture`](https://docs.snyk.io/snyk-cli/commands/iac-capture)

## `snyk iac describe` 커맨드 옵션

`--from=<STATE>[,<STATE>...]`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--to=<PROVIDER+TYPE>`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac/iac-describe)

`--service=<SERVICE>[,<SERVICE]...>`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--quiet`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--filter`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--html`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac/iac-describe)

`--html-file-output=<OUTPUT`_`FILE`_`PATH>`: [`iac-describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--fetch-tfstate-headers`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--tfc-token`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--tfc-endpoint`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--tf-provider-version`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--strict`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--deep`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

`--tf-lockfile`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

-`-config-dir`: [`iac describe`](https://docs.snyk.io/snyk-cli/commands/iac-describe)

## `snyk iac update-exclude-policy` 커맨드 옵션

`--exclude-changed`: [`iac update-exclude-policy`](https://docs.snyk.io/snyk-cli/commands/iac-update-exclude-policy)

`--exclude-missing`: [`iac update-exclude-policy`](https://docs.snyk.io/snyk-cli/commands/iac-update-exclude-policy)

`--exclude-unmanaged`: [`iac update-exclude-policy`](https://docs.snyk.io/snyk-cli/commands/iac-update-exclude-policy)

## `snyk iac rules push` 커맨드 옵션

`--delete`: [`iac rules push`](commands/iac-rules-push.md)

## `snyk iac rules test` 커맨드 옵션

`--update-expected`: [`iac rules test`](commands/iac-rules-test.md)

## `snyk ignore` 커맨드 옵션

`--id=<ISSUE_ID>`: [`ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

`--expiry=<EXPIRY>`: [`ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

`--reason=<REASON>`: [`ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

`--path=<PATH_TO_RESOURCE>`: [`ignore`](https://docs.snyk.io/snyk-cli/commands/ignore)

## `snyk sbom` 및 `snyk container sbom` 커맨드 옵션

`--format=<cyclonedx1.4+json|cyclonedx1.4+xml|cyclonedx1.5+json|cyclonedx1.5+xml|cyclonedx1.6+json|cyclonedx1.6+xml|spdx2.3+json>`: [`snyk sbom`](https://docs.snyk.io/snyk-cli/commands/sbom), [snyk container sbom](https://docs.snyk.io/snyk-cli/commands/container-sbom)

`[--file=] or [--f=]`: [`snyk sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`[--name=<NAME>]`: [`snyk sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`[--version=<VERSION>]`: [`snyk sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`[<TARGET_DIRECTORY>]`: [`snyk sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`<IMAGE>`: [`snyk container sbom`](https://docs.snyk.io/snyk-cli/commands/container-sbom)

## 메이븐 프로젝트를 위한 옵션

`--maven-aggregate-project`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--scan-unmanaged`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--scan-all-unmanaged`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## 그레들 프로젝트를 위한 옵션

`--sub-project=<NAME>`, `--gradle-sub-project=<NAME>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--all-sub-projects`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--configuration-matching=<CONFIGURATION_REGEX>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--configuration-attributes=<ATTRIBUTE>[,<ATTRIBUTE>]...`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--init-script=<FILE`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## .Net 및 NuGet 프로젝트를 위한 옵션

`--file=.sln`: [test](https://docs.snyk.io/snyk-cli/commands/test)

`--file=<filename>.sln`: [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

-`-file=packages.config`: [test](https://docs.snyk.io/snyk-cli/commands/test), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--assets-project-name`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--packages-folder`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--project-name-prefix=<PREFIX_STRING>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--project-name-prefix=my-group/`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--dotnet-runtime-resolution`:  `test,` [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--dotnet-target-framework`: `test,` [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## Options for npm projects

`--strict-out-of-sync=true|false`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [sbom](https://docs.snyk.io/snyk-cli/commands/sbom)

## Options for pnpm projects

`--dev:` [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--all-projects`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--fail-on`: [`test`](https://docs.snyk.io/snyk-cli/commands/test)

`--prune-repeated-subdependencies`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## Options for Yarn projects

`--strict-out-of-sync=true|false`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--yarn-workspaces`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

## Options for CocoaPods projects

`--strict-out-of-sync=true|false`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## Options for Python projects

`--command=<COMMAND>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--skip-unresolved=true|false`: [`test`](https://docs.snyk.io/snyk-cli/commands/test), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--File=<filename>:` [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--pakage-manager=<package manager>`: [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

## Options for Go projects

다음 옵션은 지원되지 않습니다:

`--fail-on=<all|upgradable|patchable>`: [`test`](https://docs.snyk.io/snyk-cli/commands/test)

## Options for scanning using `--unmanaged`

`--org=<ORG_ID>`: [`test`](commands/test.md), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--json`: [`test`](commands/test.md), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--json-file-output=<OUTPUT_FILE_PATH>`: [`test`](commands/test.md)

`--remote-repo-url=<URL>`: [`test`](commands/test.md)

`--severity-threshold=<low|medium|high|critical>:` [`test`](commands/test.md)

`--target-reference=<TARGET_REFERENCE>`: [`test`](commands/test.md), [monitor](commands/monitor.md)

`--max-depth`: [`test`](commands/test.md), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor), [`sbom`](https://docs.snyk.io/snyk-cli/commands/sbom)

`--print-dep-paths:` [`test`](commands/test.md), [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

`--project-name=c-project`: [`monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)

## `-- [<CONTEXT-SPECIFIC_OPTIONS>]`

이러한 옵션은 `snyk test` 및 `snyk monitor` 명령어와 함께 사용됩니다. 자세한 내용은 [`snyk test`](https://docs.snyk.io/snyk-cli/commands/test) 및 [`snyk monitor`](https://docs.snyk.io/snyk-cli/commands/monitor)의 도움말을 참조하십시오.