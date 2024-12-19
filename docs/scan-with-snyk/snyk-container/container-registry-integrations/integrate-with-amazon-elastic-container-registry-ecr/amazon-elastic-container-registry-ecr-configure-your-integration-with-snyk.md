# Amazon Elastic Container Registry (ECR) - Snyk와의 통합 구성

IAM 역할을 생성하거나 업데이트한 후에는 계속하기 전에 AWS가 서버에서 역할을 업데이트하기 위해 몇 분 정도를 기다려야 합니다.

1. AWS에서 **Summary** 섹션의 **Role** 영역 상단에 표시되는 **Role ARN** 키를 복사합니다.
2. Snyk 계정에 로그인합니다.
3. **Integrations**으로 이동하고 Amazon ECR 옵션을 클릭합니다.\
   설정 영역에서 Amazon ECR 구성 페이지가 로드됩니다.
4. 다음과 같이 자격 증명을 입력합니다:
   1. **AWS 지역**—`region-part-#` 형식을 사용하십시오. 예를 들어, `eu-west-3`.\
      가져오려는 리포지토리와 이미지가 사용 가능하게 하려면 AWS 계정에 구성된 기본 지역을 입력해야 합니다.
   2. **Role ARN**—AWS 계정에서 다음과 같은 형식으로 복사합니다: `arn:aws:iam:::role/`.
5. **Save**를 클릭합니다.

다음은 예시입니다:

```
   arn:aws:iam::881001789406:role/TestSnykIntegration_role
```

Snyk는 연결값을 테스트하고, 입력한대로 Amazon ECR 통합 세부 정보가 표시되는 페이지가 다시 로드됩니다. 또한 세부 정보가 저장되었음을 나타내는 확인 메시지도 화면 상단에 녹색으로 표시됩니다.

AWS에 대한 연결이 실패하면 **Connected to Amazon ECR** 섹션 아래에 알림이 표시됩니다.  