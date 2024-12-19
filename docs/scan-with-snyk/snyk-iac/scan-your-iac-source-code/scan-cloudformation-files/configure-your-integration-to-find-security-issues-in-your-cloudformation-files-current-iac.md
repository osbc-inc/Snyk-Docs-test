# 현재 IaC에서 보안 문제 찾기를 구성하십시오

{% hint style="info" %}
이 페이지는 현재 IaC에만 적용됩니다.
{% endhint %}

Snyk는 소스 코드 저장소로부터 CloudFormation 파일을 테스트하고 모니터링합니다. 프로덕션으로 푸시되기 전에 구성 오류를 잡아 클라우드 환경을 더 잘 보호하는 방법에 대한 조언을 제공하며, 이를 해결하는 가장 좋은 방법에 대한 지원을 제공합니다.

## 지원되는 Git 저장소 및 CloudFormation 파일 형식

Snyk는 통합된 Git 저장소에서 가져온 `JSON` 및 `YAML` 포맷의 CloudFormation 파일을 스캔합니다. `snyk iac test` CLI 명령을 사용하여 모듈을 보유한 리포지토리를 가져오거나 디렉토리 자체를 지정하여 CloudFormation 모듈 리포지토리를 스캔합니다.

CloudFormation을 스캔하면 모듈에 정적으로 구성된 모든 것에 대한 보안 피드백이 제공됩니다. 반복 및 예약된 테스트를 받으려면 사용자 지정 모듈을 SCM에서 직접 가져오는 것이 최선의 방법입니다.

## Snyk를 구성하여 CloudFormation 구성 파일을 스캔하십시오

### **CloudFormation 파일 스캔을 위한 전제 조건**

- Snyk에서 구성 중인 조직의 관리자여야 합니다.
- 이미 Git 저장소를 통합했는지 확인하십시오. 자세한 내용은 [Git 저장소 (SCM) 통합](../../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/)을 참조하십시오.

### **Snyk를 구성하여 CloudFormation 파일 스캔**

- 계정에 로그인하고 관리하려는 해당 그룹 및 조직으로 이동합니다.\
  통합은 조직 당으로 관리됩니다.
- 아래와 같이 설정을 전환하여 Infrastructure as code 파일을 감지하도록 Snyk를 활성화합니다:

<figure><img src="../../../../.gitbook/assets/snyk-iac-enable.png" alt="Infrastructure as code 구성 파일 감지 활성화"><figcaption><p>Infrastructure as code 구성 파일 감지 활성화</p></figcaption></figure>

- 필요한 경우 예시에서 AWS 탭의 **Infrastructure as code Severity settings**를 검토하고 조정하십시오.\
  검사할 파일 유형(CloudFormation, Terraform 또는 둘 다)을 선택하고, 각 API Gateway에 대해 심각도 수준을 선택하기 위해 드롭다운 메뉴에서 선택을 합니다.

<figure><img src="../../../../.gitbook/assets/image (107) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (3).png" alt="IaC 스캔을 위한 심각도 설정 선택"><figcaption><p>IaC 스캔을 위한 심각도 설정 선택</p></figcaption></figure>