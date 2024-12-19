# CLI 도움말

Snyk CLI는 프로젝트의 보안 취약점 및 라이선스 문제를 스캔하고 모니터링합니다.

자세한 정보는 [Snyk 웹 사이트](https://snyk.io)를 방문하십시오.

자세한 내용은 [CLI 문서](https://docs.snyk.io/features/snyk-cli)를 참조하십시오.

## 시작하는 방법

1. `snyk auth`를 실행하여 인증합니다.
2. `snyk test`로 로컬 프로젝트를 테스트합니다.
3. `snyk monitor`로 새로운 취약점에 대한 알림을 받습니다.

## 사용 가능한 명령어

각 Snyk CLI 명령어에 대해 자세히 알아보려면 `--help` 옵션을 사용하십시오. 예를 들어, `snyk auth --help`.

**참고:** 문서 사이트의 도움말은 CLI의 `--help`와 동일합니다.

### [`snyk auth`](auth.md)

Snyk 계정을 사용하여 Snyk CLI를 인증합니다.

### [`snyk test`](test.md)

오픈 소스 취약점 및 라이선스 문제를 테스트합니다.

**참고**: 알려진 오픈 소스 종속성을 스캔하려면 `snyk test --unmanaged`를 사용하십시오 (C/C++ 전용).

### [`snyk monitor`](monitor.md)

오픈 소스 취약점 및 라이선스 문제를 스냅샷 및 계속해서 모니터링합니다.

### [`snyk container`](container.md)

이 명령은 컨테이너 이미지를 취약점에 대해 테스트하고 지속적으로 모니터링하며 컨테이너 이미지용 SBOM을 생성합니다.

### [`snyk iac`](iac.md)

이 명령은 Infrastructure as Code 파일에서 보안 문제를 찾아 보고하며, 인프라 드리프트 및 미관리 리소스를 감지하고 추적 및 경보를 작성하며 .driftigore 파일을 생성합니다.

### [`snyk code`](code.md)

`Static Code Analysis`를 사용하여 보안 문제를 찾는 `snyk code test` 명령어입니다.

### [`snyk sbom`](sbom.md)

Snyk에서 지원하는 생태계에서 SBOM 문서를 생성하거나 테스트합니다.

### [`snyk log4shell`](log4shell.md)

Log4Shell 취약점을 찾습니다.

### [`snyk config`](config.md)

Snyk CLI 구성을 관리합니다.

### [`snyk policy`](policy.md)

패키지에 대한`.snyk` 정책을 표시합니다.

### [`snyk ignore`](ignore.md)

명시된 문제를 무시하도록`.snyk` 정책을 수정합니다.

## 디버그

디버그 로그를 출력하려면 `-d` 옵션을 사용하십시오.

## Snyk CLI 구성

환경 변수를 사용하여 Snyk CLI를 구성하고 Snyk API와의 연결을 설정하는 변수를 설정할 수 있습니다. [Snyk CLI 구성](https://docs.snyk.io/features/snyk-cli/configure-the-snyk-cli)을 참조하십시오.