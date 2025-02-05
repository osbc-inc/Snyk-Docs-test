# JavaScript 및 Node.js에 대한 지침

| 제품                                                                           | 설명                                                                                                                                                                                                                                                |
| ---------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Snyk 코드**                                                                  | 소스 코드 분석을 사용하여 코드를 보안 취약점에 대해 검사합니다.                                                                                                                                                                                                              |
| <p><strong>Snyk 오픈 소스</strong><br><br>어떤 언어나 패키지 관리자에 대해 기능이 제한될 수 있습니다.</p> | <ul><li>모든 계획에서 오픈 소스 취약점 테스트 및 모니터링 가능</li><li>모든 계획에서 오픈 소스 의존성 업그레이드 버전 업 가능</li><li>라이선스 준수(유료 계획)</li></ul>                                                                                                                                  |
| **Snyk 인프라스트럭처 as 코드**                                                       | <p>쿠버네티스 배포 파일, Terraform 또는 Cloudformation 템플릿을 사용하여 새 애플리케이션을 배포할 때 구성 문제를 검사합니다.<br>자세한 내용은 <a href="../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/">Snyk CLI for </a>를 참조하세요.</p>                                  |
| **Snyk 컨테이너**                                                                | 컨테이너를 빌드하는 경우 컨테이너 이미지에 대한 문제를 검사합니다.                                                                                                                                                                                                             |
| **Snyk 통합 IaC with cloud context**                                           | <ul><li>코드에서 클라우드까지 보안을 제공합니다.</li><li>클라우드 및 컨테이너에서 런타임 구성 문제를 검색하여 인프라 획일을 감지하고 문제를 원본에서 해결합니다.</li></ul><p>자세한 내용은 <a href="../../scan-with-snyk/snyk-iac/iac+-code-to-cloud-capabilities/">Snyk 통합 IaC with cloud context</a> 페이지를 참조하세요.</p> |

## 유효성 검사, 모니터링, 경보 및 게이팅

### **SCM 통합으로**

* **Snyk 엔터프라이즈 계획에서만**, Snyk은 제품의 프로덕션에서 사용되는 컨테이너 이미지와 그들의 오픈 소스 또는 Linux 기반 패키지를 모니터링하기 위해 쿠버네티스 통합을 사용하여 알려진 취약점을 고객에게 알릴 수 있습니다.
* **모든Snyk 계획에서**, 프로덕션 통합이 없는 경우, [`snyk monitor`](../../snyk-cli/commands/monitor.md) CLI 명령을 사용하여 스냅샷을 취하고 프로덕션으로 푸시되는 항목을 모니터링하세요.

## 패키지 레지스트리 통합 (Artifactory/Nexus)

Artifactory, Nexus, npm Teams 및 npm Enterprise Package Registry 통합은 Snyk 엔터프라이즈 계획 사용자에게 제공됩니다.

Snyk Open Source Gatekeeper 플러그인은 Artifactory 및 Nexus와 통합하여 취약점 및 라이선스 문제가 있는 패키지 다운로드를 차단합니다.

Snyk Open Source는 또한 Artifactory, Nexus, npm Teams 및 npm Enterprise와 통합하여 응용 프로그램의 보안 테스트를 지원합니다. Snyk는 이 통합을 사용하여 의존성 해결, 고침 계산 및 재잠금 잠금 파일을 수행합니다.

만약 프로젝트가 이러한 저장소에서 프라이빗 의존성을 참조하지만 Snyk 엔터프라이즈 사용자가 아니라면, 이러한 의존성을 해결하고 테스트에 포함시키기 위해 Snyk CLI를 정확하게 구성된 지역 환경(빌드 파이프라인과 같은)에서 사용할 수 있습니다.

자세한 내용은 다음을 참조하세요:

* 패키지 레지스트리 통합: [npm Teams 및 npm Enterprise](../../scan-with-snyk/snyk-open-source/package-repository-integrations/npm-teams-and-npm-enterprise-integration.md), [Artifactory 레지스트리 설정](../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/) 및 [Nexus Repository Manager 설정](../../scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/).
* Gatekeeper 플러그인: [Artifactory Gatekeeper 플러그인](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/gatekeeper-plugins/artifactory-gatekeeper-plugin.md) 및 [Nexus Repository Manager Gatekeeper 플러그인](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/gatekeeper-plugins/nexus-repository-manager-gatekeeper-plugin.md)

## 언어 및 패키지 관리자 고려 사항

### devDependencies 분석

**`devDependencies`** 분석은 기본적으로 비활성화되어 있습니다. 이러한 종속성은 일반적으로 프로덕션으로 전환되지 않으며 보안 및 개발자에게 “노이즈”로 간주됩니다. dev-dependencies에 대한 테스트를 활성화하려면:

* CLI 및 CI/CD 통합에 **--dev** 매개변수를 사용합니다.
* SCM 통합의 경우, 관련 구성 항목의 **설정 > 언어**를 설정합니다.

### optionalDependencies 분석

optionalDependencies는 CLI, CI/CD 및 SCM 통합을위한 기본 포함 항목입니다.

### npm

Snyk은 lockfile과 함께 또는 없이 종속성 트리를 작성할 수 있습니다. lockfile이 있으면 사용됩니다.

* **로컬 및 CI/CD**: lockfile이 없고 스캔이 CLI 또는 IDE로 이루어지는 경우, Snyk은 `node_modules`을 살펴 무엇이 설치되었는지를 판단합니다.
* **SCM 통합**: lockfile이 없는 경우, Snyk는 빌드 시점에서 트리가 어떻게 보일지 대략적으로 계산합니다. 이는 lockfile이 없을 때 개발 중인 프로젝트나 다음 빌드가 어떻게 보일지에 대한 통찰을 제공하는 데 매우 유용합니다

npm 사용자로서, 종속성을 처리할 때 언제든지 npm-audit이 사용 가능한 상황에서 '**왜** **Snyk**?'라고할 수 있습니다. 다음 능력을 제공받습니다:

* Snyk는 오픈 소스 뿐만 아니라 당신의 제1 자 코드를 안전하게 보호합니다. 만약 인프라스트럭처 as 코드 및/또는 컨테이너를 사용 중이라면 Snyk는는 가시성과 해결책을 제공할 수 있습니다.
* 개인 및 회사 모두에게 설계되었습니다.
* 오픈 소스 측면에서:
  * Snyk 보안팀이 추가한 선택기준, 업데이트 및 추가 가치를 모두받습니다.
  * 자동화 된 보안 조치가 있습니다.
* 중앙 보고서 작성.
* Git 코드 저장소 통합 뿐만 아니라, Snyk는 파이프라인과 프로덕션의 가시성을 포함하여 여러 프로그래밍 언어 및 패키지 관리자 지원합니다.
* 무시 기능.

자세한 내용은 [npm용 Snyk](./#npm)를 참조하세요.

### Yarn

yarn.lock 및 package.json이 필요합니다.

* Yarn 1, 2 및 3은 GIT 및 CLI에서 완벽하게 지원됩니다.
* CLI에서 lock 파일이 없는 경우, `node_modules` 폴더가 종속성 트리를 작성하는 데 사용됩니다.
* Yarn Workspaces에서 `nohoist`는 지원되지 않습니다.

자세한 내용은 [Yarn용 Snyk](./#yarn) 참조하세요.

### Lerna/PNPM

공식적으로 지원되지는 않지만 Yarn Workspaces로 구성된 경우 CLI 또는 IDE에서 Snyk 결과를 얻을 수 있습니다.

### 비관리 JavaScript

엔터프라이즈 계획으로, Snyk API에 액세스할 수 있으므로 종속성 및 이들의 종속성의 전체 목록을 얻을 수 있습니다.

취약점을 테스트하려면 다음 API 엔드포인트를 사용할 수 있습니다:

* [이름과 버전에 따라 공개 패키지의 문제 테스트하기](../../snyk-api/reference/test-v1.md#test-npm-packagename-version)
* [테스트 Dep 그래프](../../snyk-api/reference/test-v1.md#test-dep-graph)
* [패키지에 대한 문제 목록](../../snyk-api/reference/issues.md#orgs-org_id-packages-purl-issues)

### 동기화되지 않은 lockfiles

lockfile과 패키지 파일이 동기화되어 있을 때의 동작 제어는 다음을 사용하여 수행할 수 있습니다:

* CLI 추가 값: `--strict-out-of-sync`, `--fail-on`
* Git Scans용 WebUI:
  * **설정 > 언어 > 자바스크립트**

## Snyk 통합 공통 사용 패턴

npm 및 Yarn은 잘 설계된 패키지 관리자입니다. npm 및 Yarn 패키지에 대한 주요 고려 사항은 동일한 저장소 또는 빌드에 여러 패키지 관리자 또는 프로젝트가 있는지 및 각각에 대한 크리티컬/높은 임계값과 같은 기준을 적용할지 여부입니다.

소스 코드 스캐닝을 위해 이 생태계는 특별한 옵션이 필요하지 않고 기본 기능과 명령을 갖춘 Git 및 Snyk CLI에서 테스트가 실행됩니다.

## Snyk CLI 팁과 트릭

[CLI 참고 시트](https://snyk.io/blog/snyk-cli-cheat-sheet/)를 참조하세요.

### 테스트 항목

세부 정보를 보려면 CLI에서 `--help` 옵션을 사용하여 Snyk CLI 명령에 대한 자세한 내용을 확인합니다.

모든 CLI 명령의 목록은 [CLI 명령 및 옵션 요약](../../snyk-cli/cli-commands-and-options-summary.md)을 참조하세요.

#### 오픈 소스 라이브러리

`snyk test` 명령은 찾은 첫 번째 manifest를 테스트하고 해당 단일 진입점에서 테스트를 수행합니다. 디렉토리의 모든 manifest를 분석하도록 Snyk에게 요청하려면 다음 옵션을 사용하세요:

* `--all-projects`: 이 옵션은 이 디렉터리의 모든 Yarn 및 다른 프로젝트를 감지 및 스캔합니다.
* `--yarn-workspaces`: Yarn Workspaces를 사용하는 경우 `--all-projects` 플래그를 사용하여 다른 패키지 관리자나 Yarn 워크 스페이스와 함께 패키지를 테스트 및 모니터링하세요.

{% hint style="info" %}
옵션이 필요한 패키지 관리자를 사용 중인 경우, `--file=`로 개별 목표를 지정하는 것이 좋습니다.
{% endhint %}

#### 코드베이스

* 프레임워크 지원 - [지원되는 언어, 프레임워크 및 기능 개요](../) 참조
* 소스 코드 분석을 수행하려면 프로젝트의 루트에서 `snyk code test` 명령을 사용합니다.

#### 컨테이너

* 컨테이너 스캔의 일환으로 Snyk 애플리케이션(오픈 소스) 취약점을 자동으로 찾습니다. 프로덕션 이전 단계에서 CLI를 통해 Snyk을 통합하고 제품의 내부에 무엇이 있는지의 추가 신호 및 통찰력을 활용하세요.
* Node.JS 애플리케이션을 컨테이너에서 배포하는 경우, 컨테이너 베이스 이미지가 포함하는 것 외에도 (Linux, 오픈 소스) 취약한 패키지를 번들로 제공하는 SDK SDK Container는 애플리케이션의 공격 표면을 최소화하는 안전한 베이스 이미지를 식별할 수 있습니다.
* 작업하고자 하는 레이어를 필터링하려면 즉, 안전하고 체계적인 베이스 이미지를 식별하거나 채택할 레이어, 또는 애플리케이션(운영 체제) 취약점을 식별하는 작업을 하려면 [Snyk 컨테이너 보안을위한 CLI](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/) 참조

#### **Infrastructure as code**

[Infrastructure as code 보안](https://snyk.io/product/infrastructure-as-code-security/) 참조

### **옵션 및 플러그인**

* 로컬 또는 빌드 타임에 보고서를 생성하는 데 도움이 되는 경우 [snyk-to-html 플러그인](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md) 확인하세요.
* `--json` 및 `--sarif` 옵션을 사용하여 프로그래밍 방식으로 액세스 할 수 있는 출력 생성 확인하세요.
* 고급 필터링 옵션에 대해서는 [snyk-filter](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md) 참조

## Node.js 및 JavaScript 개발자에게 영향을 미치는 추가 보안 주제

다음은 Snyk 보안 팀과 Developer Relations에서 제공하는 이 생태계에 영향을 미치는 기사 모음입니다. 산업, 보안 및 기술 관련 추가 기사에 대해서는 [Snyk 블로그](https://snyk.io/blog/)를 방문하십시오:

* [현대 소프트웨어 공급망 보안 강화](https://snyk.io/blog/software-supply-chain-security/)
* [최신 npm 패키지 생성 모범 사례](https://snyk.io/blog/best-practices-create-modern-npm-package/)
* [npm에서 종속성 혼동 공격을 감지하고 방지하여 공급망 보안 유지](https://snyk.io/blog/detect-prevent-dependency-confusion-attacks-npm-supply-chain-security/)
* [Node.js 버전 전환](https://snyk.io/blog/mastering-node-js-version-management-and-npm-registry-sources-like-a-pro/)
* [JavaScript 및 Node.js의 오픈 소스 프로젝트를 위한 DevSecOps 도구](https://snyk.io/blog/devsecops-tools-for-open-source-projects-in-javascript-and-node-js/)
