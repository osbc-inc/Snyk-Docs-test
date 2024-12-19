# IAC 소스 사용

## **지원되는 IaC 소스**

현재, `snyk iac describe` 명령어는 다음과 같은 Terraform 상태를 읽는 것을 지원합니다:

* 로컬: `--from="tfstate://terraform.tfstate"`
* S3: `--from="tfstate+s3://my-bucket/path/to/state.tfstate"`
* GCS: `--from="tfstate+gs://my-bucket/path/to/state.tfstate"`
* HTTPS: `--from="tfstate+https://my-url/state.tfstate"`
* Terraform Cloud / Terraform Enterprise: `--from="tfstate+tfcloud://WORKSPACE_ID"`
* Azure blob storage: `--from="tfstate+azurerm://container-name/path/to/state.tfstate"`

지원되지 않는 백엔드를 사용하려면 `terraform`을 사용하여 상태를 파일로 보내고 그 파일을 `snyk iac describe`와 함께 사용할 수 있습니다:

```
$ terraform state pull > state.tfstate
$ snyk iac describe --from="tfstate://state.tfstate"
```

## **S3 읽기 전용 액세스 IAM 정책**

`snyk iac describe` 명령어는 읽기 전용 액세스가 필요합니다. 다음 정책은 상태 파일에 최소한의 액세스 권한을 보장합니다.

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "s3:ListBucket",
      "Resource": "arn:aws:s3:::mybucket"
    },
    {
      "Effect": "Allow",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::mybucket/path/to/my/key"
    }
  ]
}
```

## **HTTP + GitLab**

HTTP 백엔드는 GitLab API를 사용하여 GitLab에서 관리하는 Terraform 상태를 지원합니다.

`read_api` 스코프를 가진 액세스 토큰이 있는 GitLab 리포지토리가 필요합니다.

다음 명령어를 사용하세요:

```
$ GITLAB_TOKEN=<access_token> \
snyk iac describe \
--from="tfstate+https://gitlab.com/api/v4/projects/<project_id>/terraform/state/<path_to_state>" \
--headers "Authorization=Bearer ${GITLAB_TOKEN}"
```

GitLab에서 Terraform State를 관리하는 방법에 대한 자세한 정보는 GitLab 문서 웹사이트의 [GitLab 관리 Terraform 상태](https://docs.gitlab.com/ee/user/infrastructure/terraform\_state.html)를 참조하세요.

## **Azure Blob Storage**

Azure Blob Storage에서 상태에 액세스하려면 다음 환경 변수를 정의하세요:

```
$ export AZURE_STORAGE_ACCOUNT=...
$ export AZURE_STORAGE_KEY=...
$ snyk iac describe --from="tfstate+azurerm://my-container/terraform.tfstate"
```

이 값을 Azure 콘솔에서 다음 스크린샷과 같이 찾을 수 있습니다:

<figure><img src="https://docs.driftctl.com/assets/images/azure_storage_account_keys-ccb38d8792616d4376050fc6b715a6ef.png" alt="Azure 저장소 환경 변수"><figcaption><p>Azure 저장소 환경 변수</p></figcaption></figure>