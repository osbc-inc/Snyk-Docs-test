# 그룹 일반 설정

그룹의 설정을 보고 수정하려면 **Settings** 및 **General**으로 이동하세요:

<figure><img src="../../../.gitbook/assets/image (1) (3) (1) (1).png" alt="그룹 일반 설정"><figcaption><p>그룹의 일반 설정</p></figcaption></figure>

그룹의 일반 설정에서 다음을 보고 수정할 수 있습니다:

- **그룹 이름**: Snyk 웹 UI에 표시되는 그룹의 이름을 보거나 수정합니다.
- **그룹 ID**: 이 그룹에 대한 고유 ID를 확인하고 복사합니다. [Snyk API](../../../snyk-api/)를 사용할 때 이 ID가 필요합니다.
- **세션 만료**: 기본적으로 사용자는 Snyk에서 24시간 동안 활동이 없으면 로그아웃됩니다. 그룹의 세션 만료 시간을 변경할 수 있습니다. 자세한 내용은 사용자 관리 문서의 [Snyk 그룹의 세션 길이 구성](configure-session-length-for-a-snyk-group.md)를 참조하십시오.
- **액세스 요청**: Snyk 조직에 액세스 권한이 없는 사용자가 액세스를 요청할 수 있도록 활성화합니다.
  - 이는 Snyk에 등록한 사용자이지만 조직의 특정 프로젝트에 대한 URL에 액세스하려는 경우에만 가능합니다. (예: 풀 리퀘스트로부터)
  - 이 제한은 누군가가 조직의 URL을 추측하는 리스크를 줄입니다.
  - 설정된 값은 모든 새로운 조직의 기본값으로 사용되지만 기존 Snyk 조직에 대한 **액세스 요청** 설정을 덮어쓰지 않습니다.
  - 자세한 내용은 조직에 액세스 요청의 설명서인 [요청 액세스 설정 활성화](../organizations/requests-for-access-to-an-organization.md#enable-the-request-access-setting)를 참조하십시오.
- **프로젝트 테스트 빈도**: 이 그룹에서 생성된 모든 새로운 프로젝트의 기본 테스트 빈도를 설정합니다. **프로젝트 테스트 빈도** 설정 변경은 [Snyk Infrastructure as Code](../../../scan-with-snyk/snyk-iac/scan-your-iac-source-code/) 또는 [{{Snyk Code}}](../../../scan-with-snyk/snyk-code/) 프로젝트의 기본 테스트 빈도에 영향을 주지 않음에 유의하십시오. 이 둘의 기본값은 모두 주간입니다.