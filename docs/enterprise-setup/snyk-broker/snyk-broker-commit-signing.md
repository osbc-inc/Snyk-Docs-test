# Snyk Broker - 커밋 서명

{% hint 스타일 = "info" %}
**릴리스 상태**

Snyk 브로커 커밋 서명은 초기 액세스 상태이며 엔터프라이즈 요금제에서만 제공됩니다. 이 기능을 사용하려면 Snyk 계정 팀에 문의하십시오.
{% endhint %}

버전 v4.151.0부터 Snyk 브로커 클라이언트는 GitHub 통합을 위한 커밋 서명을 지원합니다. 브로커 설정을 통해 GitHub 커밋에 대한 수정 PR을 GPG 키와 구성한 전용 사용자로 서명할 수 있습니다.

## 커밋 서명 요구 사항

* 브로커 클라이언트 버전 v4.151.0 이상
* GitHub 계정이 **Access->SSH and GPG keys** 섹션 아래에서 올바르게 구성된 GPG 키로 커밋 서명 구성

## 커밋 서명 구성

1. 커밋 서명을 사용하려면 브로커 클라이언트에 다음 환경 변수를 제공하십시오:
   * **`GPG_PRIVATE_KEY`**: ASCII로 보호된 GPG 개인 키를 내보낸 값. 값은 `-----BEGIN PGP PRIVATE KEY BLOCK-----`로 시작하고 `-----END PGP PRIVATE KEY BLOCK-----`로 끝나야 합니다.
   * **`GPG_PASSPHRASE`**: GPG 개인 키의 암호.
   * **`GIT_COMMITTER_NAME`**: 커미터 이름을 설정하는 데 사용됩니다.
   * **`GIT_COMMITTER_EMAIL`**: 커미터 이메일을 설정하는 데 사용됩니다.
2. Snyk 미리보기 설정에서 **브로커 클라이언트 커밋 서명**을 활성화하십시오.

## 커밋 서명 문제 해결

GitHub에서 커밋이 `Unverified`로 표시되는 경우 다음을 수행해야 합니다:

* GPG 공개 키가 올바른 GitHub 사용자에 가져온 것이고 이메일 주소가 GitHub에서 환경 변수와 동일한지 확인합니다.
* Snyk 조직의 브로커 클라이언트에 커밋 서명 기능이 활성화되었는지 로그를 확인하고 브로커 클라이언트가 시작될 때 다음 메시지가 나타나는지 확인하십시오: `loading commit signing rules (enabled=true, rulesCount=5)`