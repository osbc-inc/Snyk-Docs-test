# IaC 테스트에 클라우드 컨텍스트 추가

{% hint style="info" %}
**릴리스 상태**\
클라우드 컨텍스트 기능은 [IaC+](./)에서만 사용 가능하며 AWS를 지원합니다.

{{Snyk IaC}}+는 지금 클로즈드 베타로 진행 중이며 새 고객을 더 이상 받지 않습니다.\
사용 가능한 기능에 대한 자세한 내용은 [현재 IaC로 시작하기](https://docs.snyk.io/scan-using-snyk/snyk-iac/getting-started-with-current-iac)를 참조하십시오.
{% endhint %}

## 클라우드 컨텍스트란?

{{Snyk IaC}}의 클라우드 컨텍스트 기능은 [IaC+](./)를 통해 배포된 클라우드 인프라에서 정보를 활용하여 IaC 테스트에서 특정 이슈를 억제합니다.

예를 들어, Terraform 구성이 퍼블릭 액세스 블록이 없는 Amazon S3 버킷을 선언하지만 계정 수준의 퍼블릭 액세스 블록이 있을 경우 Snyk는 AWS 계정의 클라우드 컨텍스트를 적용하여 버킷이 공개 액세스 블록에 의해 안전하지 않다는 오보 형태의 이슈를 억제합니다.

클라우드 컨텍스트 없이 나타나는 결과의 예시:

```
테스트 요약

  조직: demo-production
  프로젝트 이름: terraform

✔ 이슈 없는 파일: 0
✗ 이슈 있는 파일: 1
  무시된 이슈: 0
  총 이슈: 15 [ 0 심각, 7 높음, 3 중간, 5 낮음 ]
```

클라우드 컨텍스트가 있는 결과의 예시:

```
테스트 요약

  조직: demo-production
  프로젝트 이름: terraform

✔ 이슈 없는 파일: 0
✗ 이슈 있는 파일: 1
  무시된 이슈: 0
  클라우드 컨텍스트 - 억제된 이슈: 5
  총 이슈: 10 [ 0 심각, 2 높음, 3 중간, 5 낮음 ]
```

결과 요약은 억제된 이슈의 수를 나열합니다. 예를 들어, `클라우드 컨텍스트 - 억제된 이슈: 5`와 같습니다. 이 억제된 이슈는 총 이슈 수에 포함되지 않습니다. 예를 들어, `총 이슈: 10 [ 0 심각, 2 높음, 3 중간, 5 낮음 ]`과 같습니다.

현재로서는 Amazon Web Services (AWS)의 Terraform이 지원됩니다.

{{Snyk IaC}}는 [Snyk에서 컨텍스트 가져오기](add-cloud-context-to-your-iac-tests.md#bringing-context-from-a-snyk-cloud-scan)를 통해 IaC 테스트 결과에서 클라우드 컨텍스트를 적용하고 이슈를 억제할 수 있습니다.

## Snyk에서 컨텍스트 가져오기 <a href="#bringing-context-from-a-snyk-cloud-scan" id="bringing-context-from-a-snyk-cloud-scan"></a>

[Snyk 클라우드 환경](../getting-started-with-iac+-and-cloud-scans/key-concepts-for-iac+-and-cloud.md#environments)이 있다면, 클라우드 공급업체 계정에 대한 Snyk의 이미 알고 있는 정보를 활용하여 IaC 테스트에서 클라우드 컨텍스트를 적용하고 오보 형태의 결과를 줄일 수 있습니다.

[`snyk iac test`](../../../snyk-cli/commands/iac-test.md)와 함께 `--snyk-cloud-environment=<ENVIRONMENT_ID>` 옵션을 사용하여 Snyk에게 IaC 테스트에 대한 컨텍스트로 사용해야 하는 클라우드 환경을 알려줍니다.

예를 들어, 다음 명령은 현재 작업 디렉토리의 IaC를 테스트하고, Snyk 클라우드 환경 `93786877-c9f8-0000-1234-abcd1234efgh`의 최신 스캔 결과로부터 클라우드 컨텍스트를 적용합니다:

```
snyk iac test --snyk-cloud-environment=93786877-c9f8-0000-1234-abcd1234efgh
```

환경 ID를 찾으려면 [환경 ID 찾기](../getting-started-with-iac+-and-cloud-scans/snyk-environments/find-an-environment-id.md)를 참조하십시오.

Snyk 클라우드 환경을 만드는 방법에 대한 정보는 [IaC+ 문서](./)를 참조하십시오.