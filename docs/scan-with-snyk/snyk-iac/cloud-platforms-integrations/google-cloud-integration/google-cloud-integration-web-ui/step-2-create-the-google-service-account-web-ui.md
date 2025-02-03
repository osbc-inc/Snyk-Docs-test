# 2단계: Google 서비스 계정 만들기(웹 UI)

{% hint style="info" %}
**요약**\
Snyk을 위한 [Google 서비스 계정](https://cloud.google.com/iam/docs/service-accounts)을 선언하는 Terraform 템플릿을 다운로드했습니다. 이제 인프라를 프로비저닝해야 합니다.
{% endhint %}

Google 서비스 계정을 생성하는 프로세스는 [Snyk 웹 UI](./) 또는 [Snyk API](../google-cloud-integration-api/)를 사용하여 Google 프로젝트를 온보딩하는 경우에 동일합니다.

Google 클라우드 프로젝트를 스캔하기 위해 Snyk은 Snyk이 프로젝트 리소스 구성을 스캔할 수 있도록 허용하는 범위가 엄격히 지정된 Google 서비스 계정의 권한을 사용합니다.

생성하는 서비스 계정은 다음과 같은 읽기 전용 Identity and Access Management (IAM) 역할을 부여받습니다:

* [보안 리뷰어](https://cloud.google.com/iam/docs/understanding-roles#iam.securityReviewer)
* [뷰어](https://cloud.google.com/iam/docs/understanding-roles)

Snyk 클라우드의 서비스 계정은 [서비스 계정 토큰 생성자](https://cloud.google.com/iam/docs/understanding-roles#iam.serviceAccountTokenCreator) IAM 역할을 부여받아 서비스 계정에 대한 짧은 수명의 자격 증명을 생성할 수 있도록 합니다.

또한 Snyk은 서비스 계정을 온보딩한 조직에 잠금 메커니즘을 구현했습니다. 이는 누구도 서비스 계정 이름을 추측하여 해당 리소스를 확인하기 위해 별도의 조직에 온보딩할 수 없도록 하는 보안 기능입니다.

## Google 클라우드 프로젝트 ID 설정

Snyk은 Terraform 템플릿에서 `project_id` [변수](https://www.terraform.io/language/values/variables)로 지정된 Google 클라우드 프로젝트를 스캔합니다. 다음 방법 중 하나를 사용하여 변수 값을 설정해야 합니다:

* **Terraform 템플릿에서 `project_id` 변수를 직접 설정합니다.** 템플릿의 4번째 줄에서 `project_id` 변수의 기본 값을 프로젝트 ID로 변경합니다:

```
default = "your-project-id"
```

* **Terraform을 적용할 때 `project_id` 변수를 설정합니다.** 다음 섹션인 [Terraform 적용](step-2-create-the-google-service-account-web-ui.md#apply-terraform)에서 Terraform을 적용하여 Google 서비스 계정을 생성합니다. 그 때 `-var` 옵션을 사용하여 `project_id` 변수를 프로젝트 ID로 설정합니다:

```
terraform apply -var="project_id=your-project-id"
```

* **`GOOGLE_PROJECT` 환경 변수를 사용합니다.** Terraform의 [문서](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#full-reference)를 참조하세요.

## Terraform 적용

{% hint style="info" %}
[Terraform CLI](https://www.terraform.io/downloads)를 사용하기 전에 [Google Cloud 자격 증명을 구성](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/getting_started)하도록 합니다.
{% endhint %}

Terraform을 사용하여 Google 서비스 계정을 프로비저닝하려면 다음을 수행합니다:

1. 터미널에서 `.tf` 파일이 포함된 디렉토리로 이동합니다(Web UI에서 다운로드한 경우 `snyk-permissions-google.tf`로 이름이 지정됨).
2. Terraform CLI를 사용하여 Terraform 프로젝트를 초기화합니다:

```
terraform init
```

3\. Terraform 계획을 검토하고 적용합니다:

```
terraform apply
```

4\. Terraform이 해당 작업을 수행할지 묻는 메시지가 나타나면 `yes`를 입력합니다.

그러면 Terraform이 Google 서비스 계정을 생성합니다. 완료되면 다음 출력이 표시됩니다:

```
Apply complete! Resources: 22 added, 0 changed, 0 destroyed.

Outputs:

service_account_email = "snyk-cloud-mt-us-abcd1234@my-project.iam.gserviceaccount.com"
identity_provider = "https://iam.googleapis.com/projects/12345567/locations/global/workloadIdentityPools/workload-identity-123456/providers/identity-provider-123456"
```

서비스 계정 이메일 및 식별 공급자를 복사하여 [다음 단계](step-3-create-and-scan-a-cloud-environment-for-google-web-ui.md)에서 사용하세요.

## 다음은 무엇인가요?

다음 단계는 클라우드 환경을 생성하고 스캔하는 것입니다. [단계 3: Google용 클라우드 환경 생성 및 스캔(웹 UI)](step-3-create-and-scan-a-cloud-environment-for-google-web-ui.md)를 참조하세요.
