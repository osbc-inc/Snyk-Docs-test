# CI/CD 문제 해결 및 자원

## CI/CD 문제 해결

이 섹션은 CI/CD 통합을 문제 해결하거나 확장하는 데 도움이 되는 몇 가지 팁을 제공합니다.

### 단계 1: Snyk CLI로 복제해 보기

CLI와 파이프라인이 동일한 엔진을 실행하는 경우 프로젝트를 복제하고 CLI로 스캔해 보십시오.

다양한 CLI 옵션을 시도해 보세요. 파이프라인에서 실행하는 동안 알려진 취약점을 찾고 수정하기 위해 Snyk CLI 도구를 사용하세요. 자세한 내용은 [CLI 문서](../../../snyk-cli/)를 참조하세요.

### 단계 2: 로그 가져오기

CLI로 복제할 수 있고 문제가 여전히 존재하는 경우, 다음 명령을 사용하여 CLI에 디버그 로깅을 출력하도록 요청하거나 로그를 캡처하기 위해 `-d` 옵션을 사용하세요:

```
snyk test -d
```

또는

```
DEBUG=* snyk test
```

### 단계 3: 플러그인 대신 CLI 사용

CLI로 네이티브 플러그인을 대체하기 위해 CLI를 설치하여 CLI를 사용해 보세요. 지침은 [Snyk CLI 설치](../../../snyk-cli/install-or-update-the-snyk-cli/)를 참조하세요.

## CI/CD에 유용한 자원

다음 저장소는 다양한 CI/CD 도구에 대한 이진 및 npm 통합 예제를 제공합니다: [GitHub CI/CD 예제](https://github.com/snyk-labs/snyk-cicd-integration-examples).

더 자세히 알아보려면 [CI/CD란 무엇인가? CI/CD 파이프라인 및 도구 설명](https://snyk.io/learn/what-is-ci-cd-pipeline-and-tools-explained/)를 참조하세요.