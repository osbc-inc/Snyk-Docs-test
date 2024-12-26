# 컨테이너

## 설명

`snyk container` 명령은 컨테이너 이미지에 대한 취약점을 테스트하고 지속적으로 모니터링하며 컨테이너 이미지용 SBOM을 생성합니다.

더 많은 정보는 [Snyk CLI for ](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container)를 참조하십시오.

## `snyk container` 명령 및 도움말 문서

다음은 `snyk container` 명령과 관련된 도움말 옵션이 나열되어 있습니다:

- [`container test`](container-test.md), `container test --help`: 알려진 취약점을 테스트합니다.
- [`container monitor`](container-monitor.md), `container monitor --help`: 컨테이너 이미지 레이어 및 종속성을 캡처하고 [snyk.io](https://snyk.io)에서 취약점을 모니터링합니다.
- [`container sbom`](container-sbom.md), `container sbom --help`: 컨테이너 이미지용 SBOM을 생성합니다.