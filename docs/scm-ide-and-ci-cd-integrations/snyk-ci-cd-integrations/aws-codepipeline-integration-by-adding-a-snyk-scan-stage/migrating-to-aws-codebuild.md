# AWS CodeBuild로 마이그레이션

{% hint style="warning" %}
**AWS CodePipeline을 위한 Snyk 통합이 종료될 예정임**

\
**조치 필요**

서비스 및 고객의 보안을 보호하기 위해, Snyk은 **AWS CodePipeline**과의 통합을 폐기하기 시작했습니다. 중단을 최소화하기 위해, **AWS CodeBuild** 및 **Snyk CLI**를 사용하여 동일한 사용 사례와 기능을 지원하는 대체 방법으로 전환하는 것을 권장합니다.

**이주 일정**

**2024년 10월 30일** 이후, 새로운 또는 기존 파이프라인에 Snyk 플러그인을 추가하거나 수정할 수 없게 됩니다. 기존 파이프라인은 6개월 동안 그대로 작동하지만, 가능한 빨리 새 프로세스로 이주하는 것이 좋습니다. **2025년 4월 30일** 이전에 Snyk CLI로 이주하여 CI/CD 워크플로우에 지장이 없도록 해야 합니다. 아래 이주 안내서의 단계를 따라 Snyk CLI를 AWS CodeBuild와 함께 사용하는 방법을 참조하십시오.
{% endhint %}

이 안내서는보안 스캔 워크플로우를 [Snyk 및 AWS CodePipeline 통합](./)에서 [AWS CodeBuild](https://aws.amazon.com/codebuild/)로 이주하는 단계를 설명합니다. Snyk CLI 및 CodeBuild의 내장 기능을 사용하여 CI/CD 파이프라인에서 Snyk 소프트웨어 구성 분석 (SCA) 스캔을 실행하고 결과를 통합할 수 있는 보다 간소화되고 설정 가능한 솔루션을 달성할 수 있습니다.

## 이주 목표

* **현재 설정**: 워크플로우는 Snyk CodePipeline 플러그인을 CodePipeline의 전용 단계에서 사용합니다.
* **대상 설정**: Snyk 스캔이 사용자 정의 CodeBuild 빌드 단계에서 수행됩니다. 이 빌드 단계는 Snyk CLI를 직접 활용하여 스캔을 실행하고 결과를 파이프라인에 통합합니다.

## 전제 조건

* CodeBuild 및 CodePipeline 서비스가 활성화된 AWS 계정
* Snyk 계정 및 Snyk CLI가 구성되어 있는 Snyk
* CodeBuild 프로젝트 구성 및 환경 변수에 대한 이해
* 기존 CodePipeline 단계 및 Snyk과의 상호 작용에 대한 이해

## 이주 단계

보안 스캔 워크플로우를 [Snyk 및 AWS CodePipeline 통합](./)에서 [AWS CodeBuild](https://aws.amazon.com/codebuild)로 이주하는 단계에 따라 이러한 섹션의 단계를 따르십시오.

### CodeBuild 설정

* AWS 계정에서 새로운 CodeBuild 프로젝트를 생성합니다.
* 프로그래밍 언어 및 종속성에 따라 프로젝트에 호환되는 [베이스 이미지](https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html)를 선택합니다.
* [Snyk CLI를 계정에 인증하는 방법](../../../snyk-cli/authenticate-to-use-the-cli.md)을 검토하고, 환경 변수를 사용하여 Snyk CLI 토큰과 같은 중요 정보를 저장하는 고려합니다.

{% hint style="info" %}
AWS CodeBuild의 기본 **서비스 역할**에는 `/CodeBuild/`로 시작하는 모든 비밀을 AWS Secrets Manager에서 가져올 수 있는 IAM 권한이 포함되어 있습니다. 자세한 내용은 이 가이드의 맨 아래있는 [문제 해결](migrating-to-aws-codebuild.md#troubleshooting) 섹션을 참조하십시오.
{% endhint %}

* 빌드 명령 구성:
  * 운영 체제에 적합한 명령을 사용하여 [Snyk CLI 설치](../../../snyk-cli/install-or-update-the-snyk-cli/)합니다.
  * CLI를 사용하여 Snyk 스캔을 실행하는 빌드 명령을 정의합니다.
  * 프로젝트 스냅샷을 Snyk에 업로드하여 연속 모니터링을 위한 빌드 명령을 정의합니다 (선택 사항).
* 자세한 내용은 아래 예제인 `buildspec.yaml`을 검토하십시오:

```yaml
# buildspec.yaml
version: 0.2

phases:
  build:
    commands:
      # 최신 Snyk CLI를 GitHub 릴리스에서 설치합니다
      - latest_version=$(curl -Is "https://github.com/snyk/cli/releases/latest" | grep "^location" | sed 's#.*tag/##g' | tr -d "\r")
      - snyk_cli_dl_linux="https://github.com/snyk/cli/releases/download/${latest_version}/snyk-linux"
      - curl -Lo /usr/local/bin/snyk $snyk_cli_dl_linux
      - chmod +x /usr/local/bin/snyk
      
      # Snyk CLI를 인증합니다
      - snyk auth $SNYK_TOKEN
      
      # Snyk SCA 스캔을 수행합니다; 취약점이 발견된 경우 계속 진행합니다
      - snyk test || true
      
      # 프로젝트 스냅샷을 Snyk에 업로드하여 연속 모니터링합니다
      - snyk monitor
```

### CodePipeline 설정

일부 [오픈 소스](https://snyk.io/product/open-source-security-management/) 프로젝트는 Snyk CLI로 테스트하기 전에 프로젝트를 빌드해야 할 수도 있습니다. Snyk가 오픈 소스 스캔을 실행하기 전에 프로젝트를 빌드해야 하는지 확인하려면 Snyk [문서](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md)를 검토한 후 해당 섹션에서 지시사항을 따르십시오.

#### 빌드된 프로젝트가 필요한 경우

* 기존 CodePipeline을 편집하거나 새로 생성합니다.
* 프로젝트를 빌드하기 위한 새로운 단계를 생성하거나 기존 빌드 단계를 편집합니다.
* 예제 `buildspec.yaml`에서 명령을 빌드 단계에 추가하여 프로젝트 빌드 직후 Snyk 스캔이 수행되도록합니다.

{% hint style="info" %}
Snyk 오픈 소스 스캔은 Snyk이 전체 빌드 워크스페이스에 액세스할 수 있도록 빌드 프로세스와 동일한 CodeBuild 액션 내에서 이루어져야 합니다.
{% endhint %}

#### 빌드된 프로젝트가 필요하지 않은 경우

* 기존 CodePipeline을 편집하거나 새로 생성합니다.
* 원본 코드 획득 단계 이후에 새로운 빌드 단계를 추가합니다.
* 이 단계에 새로 생성한 CodeBuild 프로젝트를 선택합니다.
* Snyk이 소스 코드를 직접 스캔할 수 있도록 입력 아티팩트 아래의 SourceArtifact를 선택합니다.

### 결과 처리

CodePipeline에서 사용하는 Snyk 통합은 일부 [Snyk CLI](../../../snyk-cli/commands/) 기능 및 옵션만 지원합니다. CodeBuild에서 Snyk CLI를 사용하면 모든 Snyk CLI 기능을 사용할 수 있습니다. 그러나 Pipeline 통합과 가장 유사하게 동작을 복제하려면 다음 팁을 따를 수 있습니다:

* `snyk test` 명령은 취약점을 발견하면 비정상 종료 코드를 반환합니다. 이 동작을 우회하기 위해 명령 끝에 `|| true`를 추가하는 것을 고려해보세요.
* [snyk-to-html](https://github.com/snyk/snyk-to-html) 도구를 사용하여 `snyk test --json | snyk-to-html -o snyk-results.html`와 유사한 명령을 실행하여 스캔 결과의 HTML 보고서를 생성할 수 있습니다.
* AWS CodePipeline 통합에서 구성한 동작을 복제하기 위해 다음 [CLI 옵션](https://docs.snyk.io/snyk-cli/commands)을 사용하십시오:
  * [--org=\<ORG\_ID>](https://docs.snyk.io/snyk-cli/commands/test#org-less-than-org_id-greater-than) - 특정 Snyk 조직에 연결된 Snyk 명령을 실행할 조직을 지정합니다.
  * [--severity-threshold=\<low|medium|high|critical>](https://docs.snyk.io/snyk-cli/commands/test#severity-threshold-less-than-low-or-medium-or-high-or-critical-greater-than) - 지정된 수준 또는 그 이상의 취약점만 보고합니다.
  * [--all-projects](https://docs.snyk.io/snyk-cli/commands/test#all-projects) - 작업 디렉토리의 모든 프로젝트를 자동 감지합니다.
  * [--project-name=\<PROJECT\_NAME>](https://docs.snyk.io/snyk-cli/commands/monitor#project-name-less-than-project_name-greater-than) - `snyk monitor` 명령에 사용할 사용자 정의 Snyk 프로젝트 이름을 지정합니다.

### 테스트 및 확인

* CodePipeline에서 수동 빌드를 트리거하여 새로운 CodeBuild 통합 테스트합니다.
* Snyk 스캔이 예상대로 실행되고 결과를 출력하는지 확인합니다.
* 이후 파이프라인 단계가 스캔 출력을 적절하게 처리하는지 확인합니다.

### 배포

* 테스트가 완료되면 업데이트된 CodePipeline을 배포하는 것을 고려합니다.
* 파이프라인이 성공적으로 Snyk 스캔을 실행하고 통합 문제를 해결하는지 모니터링합니다.

### 다음 단계

CI/CD 파이프라인에 추가적인 보안 스캔을 통합하는 데 [Snyk CLI](https://docs.snyk.io/snyk-cli) 문서를 참조하십시오.

## 결론

이러한 단계와 고려 사항을 따르면 Snyk와 AWS CodePipeline 통합에서 AWS CodeBuild를 사용하여 더 간소하고 설정 가능한 솔루션으로 보안 스캔 워크플로우를 성공적으로 이주할 수 있습니다.

## 문제 해결

**AWS Secrets Manager에 Snyk 토큰을 저장하고 AWS CodeBuild에서 사용하는 방법은 무엇인가요?**

AWS Secrets Manager 환경 변수를 사용하는 경우, 토큰을 AWS Secrets Manager에 **일반 텍스트**로 저장하고 CodeBuild 서비스 역할이 IAM에서 `secretsmanager:GetSecretValue` 권한을 갖도록 한 후 AWS CodeBuild에서의 환경 변수의 **값**을 AWS Secrets Manager에서의 **비밀 이름**으로 설정해야 합니다.
