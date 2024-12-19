# 로컬 IaC 사용자 정의 규칙 번들 사용

{% hint style="info" %}
예제에서 `bundle.tar.gz`를 보면 이 부분을 귀하의 번들 이름으로 대체할 수 있습니다. 예를 들어 `bundle-v1.0.0.tar.gz` 또는 `./bundles/team-bundle.tar.gz`와 같이 사용할 수 있습니다.
{% endhint %}

프로젝트 폴더에서 다음 명령을 실행하십시오:

```
snyk iac test --rules=bundle.tar.gz
```

이제 구성 스캔 결과에는 기본 Snyk 규칙과 사용자 정의 규칙에서 발견된 문제가 모두 포함됩니다. [IaC CLI 테스트 결과 이해하기](../../../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/understand-the-iac-cli-test-results/)를 참조하세요.

로컬 사용자 정의 규칙 번들을 문제 해결하려면 다음 명령을 `--d` 옵션으로 실행하여 디버그 로그를 활성화하십시오:

```
snyk iac test --rules=bundle.tar.gz -d
```

가능한 문제점에는 다음이 포함됩니다:

- 번들 경로가 잘못 제공되었거나 존재하지 않는 번들 경로가 제공되었습니다. `--rules` 옵션에 전달된 경로가 현재 위치에서 액세스할 수 있는지 확인하십시오. 오류는 다음과 같습니다.

```
We were unable to extract the rules provided at: ./invalid/location/bundle.tar.gz
```

- 손상된 또는 잘못된 번들이 제공되었습니다. 번들을 생성하려면 [SDK 시작 안내](../writing-rules-using-the-sdk/)의 지침을 따랐는지 확인하십시오. 오류는 다음과 같습니다.

```
We were unable run the test. Please run the command again with the `-d` flag and contact support@snyk.io with the contents of the output.
```

설명할 수 없는 불일치를 발견한 경우, [Snyk 지원팀에 문의](https://support.snyk.io)하세요.