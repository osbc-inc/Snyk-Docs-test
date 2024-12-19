# 서드파티 도구를 위한 인증

서드파티 도구에서 Snyk를 사용할 때, Snyk는 프로세스를 시작하기 위해 인증이 필요합니다.

Snyk는 API 토큰을 제공하여 서드파티 개발자 도구와의 통합을 가능하게 합니다. 개인 계정을 통해 개인 토큰을 사용하거나 [서비스 계정](service-accounts/)을 통해 해당 계정에 연결된 토큰을 사용하여 인증할 수 있습니다. 서비스 계정을 통해 인증할 경우 개인 토큰을 사용하지 않습니다.

{% hint style="info" %}
인증 목적으로, 제3자 신원 제공업체는 리포지토리에 대한 액세스 권한이 아닌 이메일 주소만 필요합니다.
{% endhint %}

## 지원되는 신원 제공자

다음 신원 제공자 중 하나를 사용하여 Snyk와의 인증을 수행할 수 있습니다::

- GitHub
- Bitbucket
- Google
- Entra ID (이전 Azure AD)
- Docker ID
- 단일 로그인 (SSO): 엔터프라이즈 플랜과 함께 제공됩니다.\
  [인증을 위한 단일 로그인 (SSO) 설정](single-sign-on-sso-for-authentication-to-snyk/) 참조.

추가 지침은 [Git 리포지토리 (SCM)](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/)에 대한 통합 페이지에서 확인할 수 있습니다.

{% hint style="info" %}
처음 Snyk 계정을 만들 때 등록한 제공자와 다른 제공자로 로그인하면 별도의 새로운 Snyk 계정이 생성됩니다.
{% endhint %}

## 개인 토큰을 사용하여 서드파티 도구를 위해 인증하는 방법

1. [Snyk 계정](https://app.snyk.io/account)에 방문합니다.
2. **일반 계정 설정**으로 이동하여 토큰을 복사합니다.
3. 토큰 필드에서 **클릭하여 표시**한 다음 API 토큰을 선택하고 복사합니다.
4. 서드파티 인터페이스에서 Snyk 토큰을 붙여넣어 통합을 구성합니다.

<figure><img src="../.gitbook/assets/uuid-8d94edf8-b42b-e5b3-ada1-e157d18ff884-en (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (3) (16).png" alt="API token 화면; 취소; 재생성; 표시를 클릭"><figcaption><p>API 토큰 화면; 취소; 재생성; 표시를 클릭</p></figcaption></figure>