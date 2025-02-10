# CI/CD 파이프라인에 Snyk 추가 및 구성하기

빌드 파이프라인에서 Snyk을 사용하여 새 취약점이 도입되는 것을 방지하는 개선 필터로 사용합니다. 설정한 실패 기준에 기반하여 작동합니다.

팀이 애플리케이션의 취약점을 이해하고 개발 주기 초기에 이를 수정하는 프로세스를 개발한 후, Snyk을 구성하여 취약점이 감지되면 빌드를 실패시켜 애플리케이션에 취약점을 도입하는 것을 방지할 수 있습니다.

## 입수 요구사항 없음

파이프라인에 테스트를 추가하는 장점은 Snyk PR Checks에 필요한 소스 제어 통합을 가져오지 않아도 되며, 심지어 PR을 테스트하는 경우라도 파이프라인에 테스트를 추가할 수 있어서 프로덕션 빌드에 새로운 취약점이 들어갈 가능성을 더 줄일 수 있습니다.

## 파이프라인 옵션

Snyk을 빌드 파이프라인에 추가할 때 일반적으로 다음 옵션을 사용할 수 있습니다:

* 사용자 도구에 대한 특정 [파이프라인 통합](../../../scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/)을 사용합니다.
* [Snyk CLI](../../../snyk-cli/)를 사용하고 특정 명령어를 직접 실행합니다.

각 옵션에는 장점이 있습니다. 기존 파이프라인 통합을 사용하면 더 신속하고 구성하기 쉬울 수 있지만, Snyk CLI를 사용하면 실패 기준에서 더 많은 옵션과 유연성을 제공받을 수 있습니다.

## 파이프라인 테스트 필터

파이프라인에서 테스트를 실행할 때 필터를 사용하여 테스트가 통과하거나 실패할 조건을 결정할 수 있습니다. 이 중에서도 가장 일반적인 것은 심각도 임계값으로, High 또는 Critical 심각도 취약점만 있을 때 빌드를 실패하도록 지정할 수 있습니다.

## CLI 지원 도구

파이프라인에서 Snyk CLI를 사용할 때 추가 기능을 제공하는 여러 [CLI 도구](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/)를 사용할 수 있습니다. 이 도구들은 다음과 같습니다:

* [snyk-delta](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-delta.md): 두 세트의 결과를 비교하고 새로운 취약점을 식별하는 데 사용됩니다. 이는 PR Checks 기능이 새로운 취약점만을 테스트하는 방식과 유사합니다.
* [snyk-filter](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md): `High 심각도 취약점이 세 개 이상 발견되면 실패`와 같은 복잡한 실패 기준에 사용할 수 있습니다.

## 추가 정보

[CI/CD Best Practices](https://www.youtube.com/watch?v=6QS9gRQ0WVU)는 CI/CD 체크 사항을 자세히 다루는 웨비나로, 이 기능을 점진적으로 도입하는 예시를 포함하고 있습니다.
