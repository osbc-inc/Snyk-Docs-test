# 단계 1: 서비스 계정 IaC 템플릿 다운로드 (웹 사용자 인터페이스)

클라우드 환경을 생성하기 전에, Snyk에게 귀하의 Google 프로젝트의 자원 구성을 스캔할 권한을 부여하는 Google 서비스 계정을 선언하는 인프라 구성 (IaC) 템플릿을 다운로드해야 합니다.

이 템플릿은 또한 귀하의 Google Cloud 프로젝트를 위해 [Google 서비스 API](https://cloud.google.com/service-usage/docs/enabled-service) 세트를 활성화합니다. 이를 통해 Snyk가 프로젝트의 자원을 스캔하기 위해 필요한 API를 사용할 수 있게 됩니다.

이 IaC 템플릿을 사용하여 [단계 2: Google 서비스 계정 생성](step-2-create-the-google-service-account-web-ui.md)에서 롤을 프로비저닝할 것입니다.

## IaC 템플릿 다운로드

1. [Snyk 웹 UI](https://app.snyk.io/)에서 **Integrations > Cloud platforms**로 이동합니다.
2. **Google Cloud**를 선택합니다.
3. **Google Cloud Environment 추가** 모달에서 `snyk-permissions-google.tf` 파일을 다운로드하기 위해 **Terraform** 버튼을 선택합니다:

<figure><img src="../../../../../.gitbook/assets/Bildschirmfoto 2023-07-18 um 12.16.54 (1) (1).png" alt="Snyk 클라우드 추가 Google Cloud 환경 모달"><figcaption><p>Google Cloud 환경 추가 모달</p></figcaption></figure>

이제 [단계 2: Google 서비스 계정 생성](step-2-create-the-google-service-account-web-ui.md)로 진행할 수 있습니다.

{% hint style="info" %}
**Organization Settings (톱니바퀴 아이콘) > Cloud environments**에서도 클라우드 환경을 추가할 수 있습니다. [환경 보기, 클라우드 환경 추가](../../../../../scan-with-snyk/snyk-iac/getting-started-with-iac+-and-cloud-scans/snyk-environments/view-add-and-remove-environments.md#add-an-environment)를 참조하세요.
{% endhint %}