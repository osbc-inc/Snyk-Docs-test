# GCR 액세스 권한 활성화

## **GCR 권한을 활성화하는 데 필요한 사전 조건**

[Cloud Resources Manager API](https://console.cloud.google.com/apis/library/cloudresourcemanager.googleapis.com?q=cloud%20resource%20manager\&id=16f5d23e-c895-4b9d-88e4-864c1766636f\&project=next-for-integration-testing)를 활성화하십시오. 이 API는 Snyk와 통합할 예정인 Google 계정에서 사용할 수 있어야 합니다.

Google의 해당 프로젝트에서 이 Snyk 통합을 위해 서비스 계정을 생성했는지 확인하십시오.

## **GCR 권한 활성화 단계**

1. Google Cloud Platform 콘솔 [자격 증명](https://console.cloud.google.com/apis/credentials) 페이지로 이동하여 통합할 프로젝트를 선택한 다음 **새 서비스 계정 키 설정**을 선택합니다.
2. 열리는 화면에서 **서비스 계정**을 선택하고 이 통합을 위해 만든 드롭다운 목록에서 **서비스 계정**을 선택하고 해당 계정에 대해 다음 값으로 **새 키**를 구성합니다:
   * **서비스 계정 이름** - 나중에 해당 사용처를 기억하기 위해 계정에 고유한 이름을 할당합니다.
   * **역할** - 스토리지 오브젝트 뷰어(roles/storage.objectViewer)
   * **서비스 계정 ID** - 비워 둡니다.
   * **키 유형** - JSON
3. 프로젝트에 대한 키를 생성하려면 **만들기**를 클릭하십시오.
4. JSON 파일의 전체 내용을 복사합니다.

복사한 데이터를 GCR에 대한 통합을 구성할 때 붙여넣기 유의하십시오.