# 단계 2: Entra ID 앱 등록 생성

{% hint style="info" %}
**요약**\
[Azure Active Directory (AD) 애플리케이션 등록](https://learn.microsoft.com/en-us/azure/active-directory/develop/app-objects-and-service-principals#application-registration), [페더레이션 신원 자격 증명](https://learn.microsoft.com/en-us/azure/active-directory/develop/workload-identity-federation) 및 Snyk를 위한 [서비스 주체](https://learn.microsoft.com/en-us/azure/active-directory/develop/app-objects-and-service-principals#service-principal-object)를 선언하는 Terraform 템플릿을 다운로드했습니다. 이제 인프라를 프로비저닝해야 합니다.
{% endhint %}

Azure 구독을 스캔하기 위해 Snyk는 리더([Reader](https://learn.microsoft.com/en-us/azure/role-based-access-control/built-in-roles#reader)) 역할을 가진 서비스 주체의 권한을 사용합니다. 이렇게 함으로써 Snyk가 구독 리소스의 구성을 스캔할 수 있게 됩니다.

또한, Snyk는 구독 및 테넌트에 대한 페더레이션 자격 증명을 조직에 잠그는 보안 기능을 보유하고 있습니다. 이렇게 함으로써 누군가 자격 증명의 이름을 추측하여 해당 리소스를 보기 위해 별도의 조직에 가입하는 것을 방지합니다.

## Terraform 또는 Azure CLI를 사용하여 인프라 생성

Snyk에서 다운로드한 파일 유형에 따라 다음 도구 중 하나를 사용하여 앱 등록, 페더레이션 신원 자격 증명 및 서비스 주체를 생성할 수 있습니다:

* [Terraform](step-2-create-the-azure-ad-app-registration.md#create-azure-app-registration-infrastructure-using-terraform): Terraform CLI
* [Bash](step-2-create-the-azure-ad-app-registration.md#create-azure-app-registration-infrastructure-using-the-azure-cli): 로컬에 설치된 Azure CLI 또는 Cloud Shell을 통해 액세스

### Terraform을 사용하여 Azure 앱 등록 인프라 생성

{% hint style="info" %}
[Terraform CLI](https://www.terraform.io/downloads)를 사용하기 전에 반드시 [Azure 자격 증명을 사용하도록 구성](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs#authenticating-to-azure-active-directory)하십시오. 사용자는 응용 프로그램 관리자 또는 글로벌 관리자 디렉토리 역할 중 하나를 가지고 있어야 합니다. 이것은 Terraform을 통해 [페더레이션 신원 자격 증명](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/application\_federated\_identity\_credential#api-permissions) 및 [서비스 주체](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/service\_principal)를 생성하는 데 필요합니다.
{% endhint %}

1. 터미널에서 다운로드한 Terraform 파일이 있는 디렉터리로 이동합니다(`Snyk Web UI`에서 다운로드한 경우 `snyk-permissions-azure.tf`로 명명됨).
2. Terraform CLI를 사용하여 Terraform 프로젝트를 초기화합니다:

```bash
terraform init
```

3. Terraform 계획을 검토하고 적용합니다:

```bash
terraform apply
```

4. Terraform이 동작을 수행할지 묻는 메시지가 나타나면 `yes`를 입력합니다.

그러면 Terraform이 인프라를 프로비저닝합니다. 완료되면 다음 출력이 표시됩니다:

```bash
Apply complete! Resources: 4 added, 0 changed, 0 destroyed.

Outputs:

application_id = 12345678-9012-3456-7890-12345678abcd
```

다음 단계에서 사용할 애플리케이션 ID를 복사합니다.

### Azure CLI를 사용하여 Azure 앱 등록 인프라 생성

{% hint style="info" %}
로컬에서 [Azure CLI](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli)를 사용하는 경우에는 먼저 [`az login`](https://learn.microsoft.com/en-us/cli/azure/authenticate-azure-cli)로 로그인했는지 확인하십시오. [Azure Cloud Shell](https://portal.azure.com/#cloudshell/)에서 CLI를 사용하는 경우 이 단계는 필요하지 않습니다.
{% endhint %}

1. 터미널 또는 [Cloud Shell](https://portal.azure.com/#cloudshell/)에서 `.sh` 파일을 포함하는 디렉터리로 이동합니다.
2. `chmod` 명령을 사용하여 Bash 스크립트를 실행 가능하게 만듭니다:

```bash
chmod 755 snyk-permissions-azure.sh
```

3. 스크립트를 실행합니다:

```bash
./snyk-permissions-azure.sh
```

그러면 Azure CLI가 AD 앱 등록, 페더레이션 신원 자격 증명 및 서비스 주체를 생성합니다. 완료되면 만들어진 인프라에 대한 정보가 포함된 JSON 출력이 표시됩니다.

출력 중 마지막 항목인 **애플리케이션 ID를** 복사합니다:

```json
{
  "application_id": "12345678-9012-3456-7890-12345678abcd"
}
```

이 애플리케이션 ID를 다음 단계에서 사용할 수 있도록 복사합니다.

## 다음 단계는 무엇인가요?

다음 단계는 클라우드 환경을 생성하고 스캔하는 것입니다. [Step 3: Create and scan a Cloud Environment for Azure (Web UI)](step-3-create-and-scan-a-snyk-cloud-environment-for-azure-web-ui.md)을 참조하세요.