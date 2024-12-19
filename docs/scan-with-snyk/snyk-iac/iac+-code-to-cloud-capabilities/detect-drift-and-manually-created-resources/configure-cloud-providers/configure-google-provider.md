# 구글 공급업체 구성

## Google 공급업체의 인증

`iac describe`를 사용하려면 GCP 프로젝트에 인증된 요청을 만들기 위한 자격 증명을 설정해야 합니다.

`iac describe` 명령어는 Cloud Asset API를 사용하므로 **서비스 계정을 사용해야 합니다**.

서비스 계정 설정에 대한 자세한 내용은 [GoogleCloud 문서](https://cloud.google.com/docs/authentication/production)를 참조하십시오.

```
GOOGLE_APPLICATION_CREDENTIALS=your-creds.json \
  CLOUDSDK_CORE_PROJECT=my-project \
  snyk iac describe --to="gcp+tf"
```

[GoogleCloud SDK 환경 변수](https://cloud.google.com/sdk/docs/properties#setting_properties_via_environment_variables)에서 어떤 `env var`도 사용할 수 있습니다.

## 최소 특권 정책 <a href="#least-privileged-policy" id="least-privileged-policy"></a>

`iac describe` 명령어는 귀하의 계정에서 리소스를 나열하기 위해 [Google Asset API](https://console.cloud.google.com/apis/api/cloudasset.googleapis.com/overview)와 [Cloud Resource Manager API](https://console.cloud.google.com/marketplace/product/google/cloudresourcemanager.googleapis.com)를 사용합니다. 아래 스크린샷에 표시된 대로 사용 중인 GCP 프로젝트에 이러한 API를 활성화해야 합니다.

<figure><img src="https://docs.driftctl.com/assets/images/enable_api-dffb8e57a0ce1c667527ede14b2728df.png" alt="Enable Cloud Asset API"><figcaption><p>Cloud Asset API 활성화</p></figcaption></figure>

리소스를 나열하려면 적어도 [Cloud Asset Viewer](https://cloud.google.com/iam/docs/understanding-roles#cloud-asset-roles) 역할이 필요합니다.

## **필수 역할**

`iac describe`를 딥 모드로 사용하려면 리소스의 세부 정보를 검색할 수 있는 액세스 및 **Cloud Asset Viewer 역할만으로는 충분하지 않습니다**. 세부 정보를 얻으려면 프로젝트에 [**Viewer**](https://cloud.google.com/iam/docs/understanding-roles#basic-definitions) 기본 역할을 설정해야 합니다. IAM 정책을 읽으려면 프로젝트에 [iam.securityReviewer](https://cloud.google.com/iam/docs/understanding-roles#iam-roles) 역할도 필요합니다.

```
# 리소스 열거를 허용하기 위한 필수 역할
roles/cloudasset.viewer

# 딥 모드 전용 필수 역할
roles/viewer
```