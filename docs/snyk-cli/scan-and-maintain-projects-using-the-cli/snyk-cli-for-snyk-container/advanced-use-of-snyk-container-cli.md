# {{Snyk Container}} CLI의 고급 사용

## 아카이브 스캔

로컬 Docker 데몬이나 원격 레지스트리에서 이미지를 스캔하는 것 외에도 Snyk는 `snyk container test docker-archive:<filename>.tar` 또는 `snyk container test oci-archive:<filename>.tar`을 실행하여 Docker 또는 OCI 아카이브를 직접 스캔하거나 모니터링할 수 있습니다.

예를 들어:

```
snyk container test docker-archive:archive.tar
snyk container test oci-archive:archive.tar
```

{% hint style = "info" %}
`crane`의 경우, Snyk는 `--format=oci`와 `--format=legacy` 형식만 지원합니다.&#x20;
{% endhint %}

## 다중 플랫폼 이미지 테스트

일부 리포지토리는 여러 다른 이미지를 가리키는 다중 매니페스트를 나타냅니다. 특정 플랫폼을 위해 이미지를 명시적으로 스캔하려면 Snyk CLI `container test` 명령을 사용할 수 있습니다.

예를 들어:

```
snyk container test --platform=linux/arm64 debian
```

`--platform` 옵션에는 다음 중 하나가 포함되어야 합니다:

* linux/amd64
* linux/arm64
* linux/riscv64
* linux/ppc64le
* linux/s390x
* linux/386
* linux/arm/v7
* linux/arm/v

## 원격 컨테이너 레지스트리에 인증

Docker가 설치된 경우, Snyk CLI `container` 명령은 미리 구성된 레지스트리 인증을 사용합니다. Docker를 사용하지 않는 경우에는 다음 중 하나의 방법으로 자격 증명을 명령줄에 전달할 수 있습니다:

* 다음 환경 변수를 사용하십시오: `SNYK_REGISTRY_USERNAME` 및 `SNYK_REGISTRY_PASSWORD`
* 사용자 이름 및 암호를 전달하십시오:

```
snyk container test <repository>:<tag> --username= --password=
```

{% hint style = "info" %}
둘 다 전달되면 옵션이 환경 변수보다 우선합니다.
{% endhint %}

## 다른 일반적으로 사용되는 CLI 옵션

자주 사용되는 CLI 옵션은 다음과 같습니다:

* `--json` - 다른 도구와 통합하는 데 유용합니다.
* `--sarif` - 다른 도구와 통합하는 데 유용합니다. 이 옵션은 `container test`에서만 사용할 수 있습니다. 또한 [OASIS 정적 분석 결과 교환 형식 (SARIF)](https://www.oasis-open.org/committees/tc_home.php?wg_abbrev=sarif)을 참조하십시오.
* `--exclude-base-image-vulns` - `container test`에서만 사용할 수 있습니다.
* `--severity-threshold` - `container test`에서만 사용할 수 있습니다.
* `--exclude-app-vulns`
* `--nested-jars-depth`
* `--fail-on` - `container test`에서만 사용할 수 있습니다.

자세한 내용과 CLI 옵션에 대해서는 [Snyk CLI Container](../../commands/container.md) 도움링크를 참조하거나 다음을 실행하여 도움말을 표시하십시오:

```
snyk container --help
```