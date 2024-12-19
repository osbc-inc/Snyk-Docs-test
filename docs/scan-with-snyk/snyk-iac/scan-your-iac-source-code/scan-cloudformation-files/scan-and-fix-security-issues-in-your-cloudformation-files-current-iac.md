# CloudFormation 파일에서 보안 문제를 스캔하고 수정하십시오 (현재 IaC)

{% hint style="info" %}
현재 IaC에만 해당하는 페이지입니다.
{% endhint %}

Snyk은 CloudFormation 코드를 구성 오류 및 보안 문제에 대해 스캔합니다. 구성 파일이 스캔된 후, Snyk은 관리자가 구현한 설정에 따라 구성 오류를 보고하고 필요에 따라 수정을 권장합니다.

## **CloudFormation 파일에서 문제를 스캔하고 수정하기 위한 사전 조건**

* 관리자가 조직을 선호하는 Git 저장소와 통합하고 [구성 파일에서 보안 문제를 찾기 위한 통합 설정](configure-your-integration-to-find-security-issues-in-your-cloudformation-files-current-iac.md)에 설명된대로 구성 파일을 감지하는 것을 활성화합니다.
* Snyk 계정 및 CloudFormation 파일은 `JSON` 및 `YAML` 형식이어야 합니다.

## 구성 파일 스캔 및 수정

* 계정에 로그인하고 해당 그룹 및 조직으로 이동합니다.
* 관리자가 인프라스트럭처 코드 기능을 활성화하기 전에 테스트를 위해 저장소를 가져왔다면, 추가 프로젝트 화면에서 해당 저장소를 다시 가져와서 CloudFormation 코드를 감지해야합니다:

<figure><img src="../../../../.gitbook/assets/screenshot_2020-07-09_at_12.44.03 (1) (1) (3) (3) (2) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1)  (13).png" alt="추가 프로젝트 화면"><figcaption><p>추가 프로젝트 화면</p></figcaption></figure>

각 저장소가 스캔될 때마다 모든 CloudFormation 파일이 별도의 프로젝트로 가져와지며, 예시와 유사하게 저장소별로 그룹화됩니다.

CloudFormation 파일을 가져오기 위해 저장소를 다시 가져왔다면, Snyk은 이미 가져온 응용 프로그램 매니페스트 파일을 가져와 다시 테스트하며, 테스트 시간을 "지금"으로 표시합니다.

<figure><img src="../../../../.gitbook/assets/image (231) (1).png" alt="CloudFormation 프로젝트 목록"><figcaption><p>CloudFormation 프로젝트 목록</p></figcaption></figure>

* CloudFormation 코드의 스캔 결과 및 세부 정보를 보려면 프로젝트 링크를 클릭하십시오:

<figure><img src="../../../../.gitbook/assets/image (139) (1) (1) (1) (2) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (3).png" alt="CloudFormation 프로젝트 세부 정보"><figcaption><p>CloudFormation 프로젝트 세부 정보</p></figcaption></figure>