# Azure 리포지토리(TFS)

Snyk 계정에 새로운 통합을 추가하려면 먼저 통합을 설치하려는 수준 유형을 결정해야 합니다.

* [그룹 수준](azure-repositories-tfs.md#group-level-snyk-apprisk-integrations) - Snyk 응용 프로그램에 통합을 추가하여 Snyk AppRisk Essentials 또는 Snyk AppRisk Pro에서 사용할 수 있게 합니다.
* [조직 수준](azure-repositories-tfs.md#organization-level-snyk-integrations) - Snyk 애플리케이션에 통합을 추가하여 Snyk AppRisk 이외의 모든 Snyk 제품에서 사용할 수 있게 합니다.

{% hint style="info" %}
Snyk AppRisk에 대한 통합을 설정하려면 그룹 수준의 통합 메뉴를 사용하십시오.
{% endhint %}

## 조직 수준 - Snyk 통합

{% hint style="info" %}
**기능 이용 가능성**\
Azure DevOps Server 2020 이상, 즉 TFS와의 통합은 엔터프라이즈 요금제에서만 사용할 수 있습니다. 자세한 정보는 [요금제 및 가격 정보](https://snyk.io/plans/)를 참조하세요.

Snyk은 Git만 지원합니다. Snyk은 현재 Team Foundation 버전 제어 (TFVC)와의 통합을 지원하지 않습니다.
{% endhint %}

### Azure 저장소 (TFS) 통합 사전 준비 조건

* [Snyk 조직 관리자](../../snyk-admin/user-roles/pre-defined-roles.md) 사용자 역할.
* Azure 프로젝트. 아직 프로젝트가 없는 경우 [Azure DevOps](https://docs.microsoft.com/en-us/azure/devops/user-guide/sign-up-invite-teammates?view=azure-devops)에서 생성하거나 [온프레미스 Azure DevOps](https://docs.microsoft.com/en-us/azure/devops/organizations/projects/create-project?view=azure-devops) 인스턴스에서 설정하십시오.
* 필요한 개인 액세스 토큰 (PAT) 액세스 범위. 필요한 권한에 대한 자세한 내용은 [Azure 저장소 (TFS) 권한 요구 사항](./#azure-repositories-tfs-permission-requirements)을 참조하세요.

### Azure 저장소 (TFS) 통합 기능

Snyk은 Microsoft Azure 저장소와 통합하여 프로젝트를 가져와 저장소의 소스 코드를 모니터링할 수 있습니다. Snyk 가져온 프로젝트를 검사하여 의존성의 알려진 보안 취약점을 확인하며 필요에 따라 테스트 빈도를 제어할 수 있습니다.

Azure 저장소 통합을 사용하여 다음을 수행할 수 있습니다:

* 통합된 모든 저장소에서 보안 스캔을 지속적으로 수행
* 오픈 소스 구성 요소의 취약점 식별
* 자동 수정 및 업그레이드 제공

통합을 구성한 후 Snyk은 다음을 수행합니다:

1. 선택한 항목을 평가하고 루트 폴더와 모든 하위 폴더의 해당 매니페스트 파일을 가져옵니다.
2. 가져온 각 의존성은 Snyk 취약점 데이터벤이스와 테스트되어 알려진 취약점이 있는지 확인하기 위해 Snyk 애플리케이션에 의해 현재 푸시된 코드와 사용 중인 의존성이 정확히 어떤 것인지 결정하기 위해 PAT와 연결된 권한을 사용하여 리포지토리와 직접 통신합니다.
3. 설정한 환경에 따라 발견된 취약점이 있는 경우 이를 즉시 수정할 수 있도록 이슈를 고칠 수 있도록 이메일이나 전용 Slack 채널을 통해 알려줍니다.

### Azure 저장소 (TFS) 통합 설정 방법

Snyk을 Azure 저장소와 연결하는 프로세스는 다음 단계를 포함합니다:

1. Snyk을 위한 고유한 Azure DevOps 개인 액세스 토큰(PAT)을 생성합니다. 이는 사용자 이름과 비밀번호 조합에 따라 생성되며, Snyk이 귀하의 Azure 저장소에 액세스해야 하는 특정 권한으로 구성됩니다. 자세한 정보는 [개인 액세스 토큰(PAT) 구성하기](azure-repositories-tfs.md#configure-a-personal-access-token-pat)를 참조하세요.
2. Snyk 웹 UI를 통해 [통합 활성화](azure-repositories-tfs.md#integrate-using-the-snyk-web-ui)합니다.
3. Snyk을 통해 테스트 및 모니터링할 [프로젝트 및 저장소를 선택](azure-repositories-tfs.md#add-projects-to-snyk-for-azure-repos)합니다.\
   저장소 루트 폴더에 위치하지 않은 매니페스트 파일은 사용자 지정 파일 위치에 입력할 수도 있습니다.

### **개인 액세스 토큰(PAT) 구성하기**

Snyk에서 사용할 고유한 PAT을 생성하고 복사합니다. Azure에서 활성화해야 하는 PAT 액세스 범위에 대한 자세한 내용은 [Azure 저장소 (TFS) 권한 요구 사항](./#azure-repositories-tfs-permission-requirements)을 참조하세요.

### Snyk 웹 UI를 사용한 통합

1. [Snyk 계정](https://app.snyk.io)에 로그인하고 **통합**으로 이동합니다.
2. **Azure 저장소** 타일에서 설정 아이콘을 클릭하여 **조직 설정** > **통합** > **Azure 저장소** > **계정 자격 증명**을 엽니다.
3. **계정 자격 증명** 페이지에서 주어진 지침에 주의를 기울입니다. 요금제에 따라 Azure DevOps Organization만 입력하거나 전체 URL을 입력해야 할 수도 있습니다.
   * **조직 설정**: 귀하의 조직을 위한 슬러그를 입력합니다.\
     예: `your-azure-devops-org`
   * **호스트 설정**: 전체 URL을 입력합니다.\
     예: `https://dev.azure.com/your-azure-devops-org`\
     또는 공개적으로 접근 가능한 사용자 지정 URL을 입력할 수도 있습니다.
4. **저장**을 클릭한 다음 생성한 PAT을 입력합니다.
5. **저장**을 클릭합니다.\
   입력한 연결 값에 대한 테스트를 Snyk이 수행하고 페이지가 다시로드되어 Azure 저장소 통합 정보가 표시됩니다. 정보가 업데이트되었다는 확인 메시지가 화면 상단에 나타납니다.

{% hint style="info" %}
Azure와의 연결이 실패하면 **Azure 저장소** 카드 제목 아래에 알림이 표시됩니다.
{% endhint %}

### Snyk에 Azure 저장소 프로젝트 추가

Snyk은 [Snyk 지원하는](../../supported-languages-package-managers-and-frameworks/) 언어를 위한 루트 폴더 및 사용자 지정 파일 위치를 평가하여 Azure 저장소를 테스트하고 모니터링합니다.

기본 프로젝트를 추가하려면 다음을 수행합니다:

1. Snyk에서 **프로젝트** > **프로젝트 추가**로 이동합니다.
2. 가져올 프로젝트의 관련 저장소나 도구를 선택합니다.\
   선택한 통합을 위해 사용 가능한 저장소가 새 창에 표시됩니다.
3. Snyk이 보안 및 라이선스 문제를 모니터링하도록 원하는 저장소를 선택합니다.
4. 특정 조직의 모든 리포지토리를 가져오려면 **조직**을 선택하십시오.
5. **선택한 저장소 추가**를 클릭합니다.\
   Snyk은 종속성 파일 트리 전체를 스캔하고 프로젝트로 가져옵니다.

### 사용자 지정 파일 위치 추가 및 폴더 제외

#### 사용자 지정 파일 위치 추가

기본 경로가 아닌 Azure 저장소 종속성을 추가하려면 다음 방법을 사용합니다.

1. Snyk에서 **프로젝트** > **프로젝트 추가 > Azure 저장소 > 설정**으로 이동합니다.
2. **사용자 지정 파일 위치 추가(선택 사항)** 목록을 열고 \*\*저장소 선택...\*\*을 클릭하여 사용자 정의 경로를 구성합니다.
3. 텍스트 필드에 `매니페스트 파일 위치에 상대적인 경로`를 입력합니다.

{% hint style="warning" %}
상대적인 경로 필드는 대소문자를 구분합니다.
{% endhint %}

#### 가져오기에서 폴더 제외

선택 사항인 **폴더 제외** 필드는 모든 Azure 저장소에 적용되는 대소문자가 구분되는 패턴을 입력합니다.

### 저장소 가져오기 확인

저장소가 가져오면 화면 상단에 초록색으로 확인 메시지가 나타납니다. 선택한 파일은 고유 아이콘으로 표시되며 조직 및 저장소별로 명명됩니다. Azure Repos 필터 옵션을 선택하여 그 프로젝트만 볼 수 있도록 필터링할 수 있습니다.

{% hint style="info" %}
Azure 저장소 통합은 다른 [Snyk SCM 통합](./)과 유사하게 작동합니다. 프로젝트를 계속 모니터링, 수정 및 관리하려면 [프로젝트](../../snyk-admin/snyk-projects/) 설명서를 참조하세요.
{% endhint %}

## 그룹 수준 - Snyk AppRisk 통합

통합 페이지에는 자동으로 동기화되는 기존 Snyk 조직의 데이터를 포함하여 모든 활성 통합이 표시되어 통합 허브에 액세스할 수 있습니다.

### Azure DevOps 설정 안내

#### 검색된 엔티티 <a href="#azure-devops-pulled-entities" id="azure-devops-pulled-entities"></a>

* 저장소 - Snyk AppRisk에서 검색된 엔티티입니다.

#### Snyk AppRisk를 사용하여 통합 <a href="#azure-devops-integrate-using-snyk-apprisk" id="azure-devops-integrate-using-snyk-apprisk"></a>

1. 프로필 이름 (`필수`): 통합 프로필 이름을 입력합니다.
2. 조직 (`필수`): 모든 관련 Azure DevOps 조직의 이름을 입력합니다.
3. 액세스 토큰 (`필수`): Azure DevOps 설정에서 Azure DevOps PAT를 생성합니다.
   * 액세스 토큰 (`필수`): [Azure DevOps 설정에서 개인 액세스 토큰 생성](azure-repositories-tfs.md#generate-a-personal-access-token-from-your-azure-devops-settings) 지침을 따릅니다.
   * API URL (`필수`): 예를 들어 API URL은 [`https://dev.azure.com/`](https://dev.azure.com/)입니다. 공개적으로 접근 가능한 사용자 지정 URL을 사용할 수 있습니다.
4. 브로커 토큰 (`필수`): AppRisk용 Snyk 브로커를 사용하는 경우 브로커 토큰을 생성하여 추가합니다.
   * [Snyk 브로커용 브로커 토큰 생성](../../enterprise-setup/snyk-broker/snyk-broker-code-agent/install-snyk-broker-code-agent-using-docker/obtain-the-required-tokens-for-setup.md#obtain-your-broker-token-for-snyk-broker-code-agent) 페이지를 참조하여 브로커 토큰을 생성합니다.
5. Backstage 카탈로그 추가 (`선택 사항`): Backstage 카탈로그를 추가하려면 [SCM 통합을 위한 Backstage 파일](application-context-for-scm-integrations/) 페이지의 지침을 따릅니다.

{% hint style="warning" %}
다음 PAT 토큰 권한 요구 사항은 Snyk AppRisk 통합을 위한 것입니다. SCM 통합에 대해 자세히 알아보려면 Snyk SCM 통합 페이지의 [Azure 저장소 (TFS) 권한 요구 사항](./#azure-repositories-tfs-permission-requirements)을 참조하십시오.
{% endhint %}

#### Azure DevOps 설정에서 개인 액세스 토큰 생성

1. Azure DevOps를 열어 **프로필** 메뉴로 이동합니다.
2. **개인 액세스 토큰**을 클릭한 다음 **새 토큰**을 선택합니다.
3. 다음 범위를 선택합니다:
   * 권한
     * **코드** - 읽기
     * **프로젝트 및 팀** - 읽기
     * **분석** - 읽기
     * **멤버 권리 관리** - 읽기
   * 조직 - **모든 접근 가능 조직** 또는 특정 조직을 선택합니다.
4. 만료 기간을 12개월로 설정합니다.
5. 생성된 개인 액세스 토큰을 복사하고 안전한 보관소를 통해 공유합니다.

#### API 버전 <a href="#azure-devops-api-version" id="azure-devops-api-version"></a>

API 버전 정보는 [Azure DevOps REST API v6](https://learn.microsoft.com/en-us/rest/api/azure/devops/core/?view=azure-devops-rest-6.0)를 사용하여 API에 대한 정보에 액세스할 수 있습니다.
