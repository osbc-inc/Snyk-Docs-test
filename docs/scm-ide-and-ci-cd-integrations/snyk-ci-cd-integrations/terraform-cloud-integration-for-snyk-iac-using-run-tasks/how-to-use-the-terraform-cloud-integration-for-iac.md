# IaC를 위한 Terraform Cloud 통합 사용 방법

통합이 설정되면 Snyk은 워크스페이스에서 트리거된 각 실행에 대해 Terraform 계획을 스캔합니다.

## Terraform 계획 스캔 결과 보기

1. Terraform Cloud 워크스페이스에서 트리거된 각 실행에 대해, Snyk Terraform 계획 스캔 결과는 `Plan` 단계가 완료된 후 트리거되는 `run tasks` 단계 아래에 나타납니다.
2. 스캔 결과는 `passed` 또는 `failed` 상태로 나타납니다. Snyk이 Terraform 계획 파일에서 문제를 발견하면 스캔이 실패합니다.
3. Terraform Cloud에서 실행 작업 결과의 **Details** 링크를 클릭하여 Snyk에서 자세한 내용을 볼 수 있습니다.
   1. `{YOUR_TFC_ORG_NAME}/{YOUR_TFC_WORKSPACE_NAME}`로 명명된 Target 하위에 있는 terraform-plan.json을 검색하여 Snyk **Projects** 탭에서 결과를 찾을 수도 있습니다.
   2. 좌측 패널의 필터를 사용하여 Terraform Cloud 프로젝트만 표시할 수도 있습니다.
4. Snyk에는 ({YOUR_TFC_ORG_NAME}/{YOUR_TFC_WORKSPACE_NAME})를 사용하는 각 워크스페이스에 대해 단일 프로젝트(`terraform-plan.json`)가 생성됩니다. 모든 프로젝트 페이지에는 최신 스캔 결과가 표시됩니다.
5. 기존 스캔 결과를 보려면 해당 프로젝트의 **History** 탭으로 이동하여 보고 싶은 과거 스냅샷을 선택하십시오.

## Terraform 계획 스캔 사용자 정의

Snyk Terraform Cloud 통합은 다음과 같은 사용자 정의 수준을 제공합니다:

* 심각도 임계값: 실패의 최소 심각도 수준을 설정합니다. 이는 Snyk의 통합 페이지에서 설정할 수 있습니다.
* 사용자 정의 심각도: 기본값을 덮어쓰는 문제에 대해 사용자 정의 심각도를 설정합니다(예: [SNYK-CC-00172](https://security.snyk.io/rules/cloud/SNYK-CC-00172)).
* 강제 수준: 실패가 apply를 차단하는지 여부를 결정합니다. 이 설정은 Terraform Cloud를 통해 제어됩니다. 예를 들어, `Advisory` 수준은 최소 심각도 임계값 내에서 문제가 발견되더라도 apply를 차단하지 않습니다.

## 참고 사항과 제한 사항

* Snyk은 Terraform Cloud에서 최신 실행 내에서 완료된 각 `plan` 단계에 대한 이벤트를 수신합니다.
* 스캔을 트리거하는 유일한 방법은 Terraform Cloud를 통해 새 실행을 트리거하는 것입니다.
* Snyk 웹 UI를 통해 Terraform 계획 파일의 재스캔을 트리거할 수 없습니다.
* Snyk 통합을 사용자 정의하면(예: 심각도 임계값 변경 또는 정책 심각도 사용자 정의), Snyk에서 변경 사항이 적용되려면 Terraform Cloud에서 새 실행을 트리거해야 합니다.