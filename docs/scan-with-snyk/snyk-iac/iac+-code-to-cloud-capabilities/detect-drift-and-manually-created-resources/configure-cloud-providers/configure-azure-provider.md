# Azure 공급자 구성

## Azure 공급자의 인증

`iac describe`를 사용하려면 Azure 계정에 인증된 요청을 보낼 수 있도록 자격 증명을 설정해야 합니다. Snyk는 [환경 변수](https://docs.microsoft.com/en-us/azure/developer/go/azure-sdk-authorization#use-environment-based-authentication)에서 구성 정보를 가져옵니다.

Azure 인증 구성에 대한 안내는 [Terraform 문서](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs#authenticating-to-azure)를 참조하세요.

```
# 여기서는 클라이언트 시크릿을 사용하는 서비스 주체 계정을 사용합니다
$ AZURE_SUBSCRIPTION_ID=00000000-0000-0000-0000-000000000000\
  AZURE_TENANT_ID=00000000-0000-0000-0000-000000000000\
  AZURE_CLIENT_ID=00000000-0000-0000-0000-000000000000\
  AZURE_CLIENT_SECRET=password\
  snyk iac describe --to="azure+tf"
```

또한 **az** CLI를 사용하여 인증할 수도 있습니다. 그런 다음 `AZURE_SUBSCRIPTION_ID`만 지정해야 합니다.

```
$ AZURE_SUBSCRIPTION_ID=00000000-0000-0000-0000-000000000000\
  snyk iac describe --to=azure+tf
```

## 최소 권한 정책 <a href="#least-privileged-policy" id="least-privileged-policy"></a>

`iac describe` 명령은 계정에 대해 읽기 전용 권한이 필요합니다. 전체 Azure 계정을 스캔하려면 구독에 **Reader** 역할을 설정하십시오. 다음 스크린샷에 나와 있는 대로 설치할 수 있습니다.

<figure><img src="https://docs.driftctl.com/assets/images/auth-d38df6fe7a4318ec9ebf82d0e5f9edae.png" alt="Azure 공급자를 위해 Reader 역할 설정"><figcaption><p>Azure 공급자를 위해 Reader 역할 설정</p></figcaption></figure>

특정 리소스 그룹만 스캔하려는 경우 특정 제한된 리소스 그룹에 **Reader** 역할을 할당할 수 있습니다.