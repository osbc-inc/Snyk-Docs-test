# Snyk IaC CLI 테스트 결과 (버전 1.939.0 및 이후)

{% hint style="info" %}
이 섹션의 지침은 테라폼, 쿠버네티스, 클라우드포메이션 및 ARM을 포함한 모든 파일 형식에서 지원되는 Snyk Infrastructure as Code에 적용됩니다.
{% endhint %}

Snyk CLI는 구성 파일을 분석하여 문제를 식별하고 해결 방법에 대한 정보와 조언을 제공합니다.

예를 들어, 테라폼 파일을 스캔할 때 다음 명령을 실행합니다:

```
snyk iac test aws_api_gateway_stage_logging.tf
```

이 명령을 실행한 결과는 다음과 같습니다:

```
Snyk Infrastructure as Code

✔ 테스트 완료.

문제

낮은 심각도 문제: 1

  [낮음] API Gateway 액세스 로깅 비활성화
  정보: Amazon API Gateway 액세스 로깅이 활성화되지 않았습니다. 조사 중에 감사 레코드를 사용할 수 없을 수 있습니다.
  규칙: https://security.snyk.io/rules/cloud/SNYK-CC-TF-138
  경로: resource > aws_api_gateway_stage[denied] > access_log_settings
  파일: aws_api_gateway_stage_logging.tf
  해결: `access_log_settings` 속성을 설정합니다.

-------------------------------------------------------

테스트 요약

  조직: demo-org

✔ 문제가 없는 파일: 0
✗ 문제가 있는 파일: 1
  잘못된 파일: 0
  무시된 문제: 0
  총 문제: 1 [ 0 critical, 0 high, 0 medium, 1 low ]
```

이 결과에는 심각도에 따라 정렬된 문제 목록이 포함되어 있습니다. 각 보고된 문제는 다음과 같은 세부 정보가 포함됩니다:

* **제목** - 감지된 문제와 해당 문제의 심각도 수준.
* **정보** - 감지된 문제의 간단한 설명.
* **규칙** - 규칙 설명서로의 링크.
* **경로** - 문제가 식별된 구성 파일 내의 속성 경로. 이 목록 다음에 있는 예제를 참조하세요.
* **파일** - 문제가 위치한 파일.
* **해결** - 문제를 해결하는 방법에 대한 간단한 설명.

속성 경로의 예는 다음과 같습니다.

```
resource > aws_api_gateway_stage[denied] > access_log_settingsresource > aws_api_gateway_stage[denied] > access_log_settings
```

아래 예는 `access_log_settings` 필드가 없는 "**denied**"로 불리는 `aws_api_gateway_stage` 블록의 내용을 나타냅니다:

{% code title="aws_api_gateway_stage_logging.tf" %}
```
resource "aws_api_gateway_stage" "denied" {
  xray_tracing_enabled = true
}
```
{% endcode %}