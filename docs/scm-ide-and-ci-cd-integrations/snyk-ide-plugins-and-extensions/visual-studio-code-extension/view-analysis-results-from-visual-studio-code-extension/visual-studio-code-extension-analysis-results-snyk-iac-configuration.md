# 분석 결과: Snyk IaC 구성

Snyk IaC 구성 분석은 각 스캔마다 Terraform, Kubernetes, AWS CloudFormation 및 Azure Resource Manager (ARM) 코드에서 발견된 문제를 보여줍니다. Snyk CLI를 기반으로 한 이 스캔은 로컬 개발에 빠르고 친숙하며 백그라운드에서 실행되며 기본적으로 활성화되어 있습니다.

Visual Studio Code 결과 화면의 **문제** 탭에서 프로젝트에서 발견된 모든 구성 문제를 볼 수 있습니다.

<figure><img src="../../../../.gitbook/assets/Screenshot 2023-03-16 at 14.32.48.png" alt="문제 탭"><figcaption><p>문제 탭</p></figcaption></figure>

## Snyk IaC 구성 편집기 창

편집기 내에서 이 문제들이 보이며, 자세한 정보는 호버시 사용 가능합니다.

<figure><img src="../../../../.gitbook/assets/Screenshot 2023-03-16 at 15.21.10.png" alt="Snyk IaC 구성 이슈"><figcaption><p>Snyk IaC 구성 이슈</p></figcaption></figure>

**빠른 수정**을 선택하여 코드 액션을 통해 문제에 대한 세부 정보 패널을 열 수 있습니다.

<figure><img src="../../../../.gitbook/assets/Screenshot 2023-03-16 at 15.17.50.png" alt="빠른 수정"><figcaption><p>빠른 수정</p></figcaption></figure>

세부 정보 패널에는 **설명**, **영향** 문장, 문제가 도입된 **경로**, 그리고 제안된 **복구**가 표시됩니다.

<figure><img src="../../../../.gitbook/assets/Screenshot 2023-03-16 at 14.32.23.png" alt="Snyk IaC 구성 문제에 대한 세부 정보 패널"><figcaption><p>Snyk IaC 구성 문제에 대한 세부 정보 패널</p></figcaption></figure>

## Snyk IaC 구성 문제 창

구성 문제 창에서 문제에 관한 정보가 표시됩니다. 문제를 클릭하여 자세히 알아볼 수 있습니다:

<figure><img src="../../../../.gitbook/assets/Screenshot 2023-03-16 at 15.14.16.png" alt="Snyk IaC 구성 문제 창"><figcaption><p>Snyk IaC 구성 문제 창</p></figcaption></figure>

다음 정보가 표시됩니다: 

- 문제 설명
- 문제 영향
- 문제 경로
- 복구 세부 사항
- 참조 링크들