# Google Artifact Registry(GAR)와 통합하기

Snyk은 [Google Artifact Registry (GAR)](https://cloud.google.com/artifact-registry)과 통합되어 컨테이너를 취약점으로 모니터링하고 작업 중에 이를 수정할 수 있습니다. Snyk은 정기적으로 가져온 컨테이너 이미지를 테스트합니다.

## GAR 통합 권한 활성화

### **GAR 통합 권한 활성화를 위한 전제 조건**

[Artifact Registry API 활성화](https://cloud.google.com/artifact-registry/docs/enable-service)를 통해 Snyk와 통합할 Google 계정의 권한을 활성화합니다. Google이 활성화를 전파하는 데 몇 분 정도 소요됩니다.

### **GAR 통합 권한 활성화 단계**

1. Google Cloud Console [Credentials](https://console.cloud.google.com/apis/credentials) 페이지로 이동합니다.
2. 이미 선택되어 있지 않은 경우 구성할 Google 프로젝트를 선택합니다.
3. **Credentials** 버튼을 선택하고 **Service Account**를 선택합니다.
4. 새로운 서비스 계정에 고유한 이름 및 ID를 지정하고 **Create**를 선택합니다.
5. 서비스 계정 권한 페이지에서 **Select a role** 및 **Artifact Registry Reader**를 선택합니다. 또한 **resourcemanager.projects.list** 권한을 가진 추가 역할, 예를 들어 **Browser** 또는 **Viewer**를 추가해야 합니다.
6. 계정이 생성된 후 **Service Accounts** 섹션에서 관련 계정을 선택합니다.
7. 서비스 계정 세부 정보 페이지에서 **Keys** 섹션에서 **Add Key** 및 **Create New Key**를 선택합니다.
8. **Key type**으로 JSON을 선택하고 **Create**를 선택합니다.

프로젝트용으로 JSON 키가 생성됩니다. JSON 파일의 전체 내용을 복사하여 다음 단계인 [GAR를 위한 인테그레이션 구성](integrate-with-google-artifact-registry-gar.md#configure-integration-for-gar)에 사용하세요.

## GAR를 위한 인테그레이션 구성

Snyk을 통한 Google Artifact Registry 계정의 인테그레이션을 구성하여 취약점을 스캔하고 보안 및 라이선스 문제를 해결합니다.

### GAR 인테그레이션 구성 전제 조건

GAR 통합을 위한 [권한 활성화](integrate-with-google-artifact-registry-gar.md#enable-permissions-for-gar-integration)

### GAR 인테그레이션 구성 단계

1. Snyk 웹 UI에서 조직으로 이동합니다.
2. **Integrations**을 선택합니다.
3. 컨테이너 레지스트리 섹션에서 **Google Artifact Registry**를 선택합니다.
4. 계정 자격 증명 섹션에 [Artifact Registry](https://cloud.google.com/artifact-registry/docs/repositories/repo-locations) 호스트 이름을 입력합니다. 일반적으로 **\<Your Region Name>-docker.pkg.dev**이지만 경우에 따라 특정 지역 또는 다중 지역을 사용해야 할 수도 있습니다. 예를 들어 **us-east1-docker.pkg.dev** 또는 **us-docker.pkg.dev**와 같습니다.
5. JSON 키 파일 필드에 [권한을 활성화](integrate-with-google-artifact-registry-gar.md#enable-permissions-for-gar-integration)할 때 다운로드한 JSON 키 파일 전체 내용을 붙여넣습니다.
6. **Save**를 선택합니다.

Snyk은 자격 증명을 확인하고 성공하면 페이지가 다시로드되어 연결이 성공했다는 알림이 표시됩니다.
