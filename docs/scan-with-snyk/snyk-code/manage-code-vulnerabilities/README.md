# 코드 취약성 관리

## Snyk 웹 UI에서 코드 취약점 관리를 위한 전제 조건

Snyk Code를 사용하여 취약점을 관리하기 전에 다음 사항을 확인하십시오:

* [시작하기](../../../getting-started/)에서 단계를 완료했습니다.
* 귀하의 저장소에는 [지원되는 언어 및 플랫폼](../../../supported-languages-package-managers-and-frameworks/)에서 코드가 포함되어 있습니다.
* [Snyk Code를 설정했습니다](../configure-snyk-code.md).

## Snyk 코드의 프로젝트 테스트 방법

프로젝트를 테스트할 때 매번 코드를 스냅숏으로 찍어 현재의 상태를 분석하여 취약점을 찾습니다. Snyk Code가 분석 가능한 소스 코드를 포함하는 모든 파일은 코드 분석에서 종합됩니다.

저장소를 가져올 때 Snyk은 저장소에 있는 파일 유형에 기반하여 다른 Snyk 프로젝트가 포함된 Target 폴더를 생성합니다. Target 폴더의 이름에는 저장소 이름, 통합된 Git 저장소 계정 이름 및 해당 아이콘, 그리고 저장소에 대한 생성된 Snyk 프로젝트 수가 포함됩니다. [첫 Snyk 프로젝트 보기](../../../implement-snyk/walkthrough-code-repository-projects/view-your-first-snyk-projects.md)을 참조하십시오.

Snyk Code는 한 저장소에서 가져온 모든 파일에 대해 단일 프로젝트를 생성합니다. 이는 저장소 코드에서 감지된 취약점을 하나의 프로젝트로 집계하여 취약성 문제의 데이터 흐름을 여러 파일에 걸쳐 제시합니다.

API 엔드포인트 [대상 가져오기](../../../snyk-api/reference/import-projects-v1.md#org-orgid-integrations-integrationid-import)를 사용하여 여러 저장소를 자동으로 가져올 수 있습니다.

## 가져오기부터 재테스트까지의 코드 테스트

다음은 테스트 단계를 기반으로 한 Snyk Code에서의 테스팅 프로세스에 대한 개요를 제공합니다.

| 단계                                                                                 | 설명                              |
| ---------------------------------------------------------------------------------- | ------------------------------- |
| [저장소 가져오기](../import-project-with-snyk-code.md)                                    | 저장소를 가져올 때 수행됩니다.               |
| [주기적인 테스트 일정](../../../snyk-admin/snyk-projects/view-and-edit-project-settings.md) | 일정에 따라 자동으로 수행됩니다.              |
| [요청에 따른 테스트 (저장소 재테스트)](./#retesting-code-repository)                              | **지금 재테스트**를 선택하여 요청에 따라 수행됩니다. |

### 저장소 재테스트

저장소에서 가장 최근의 취약점을 확인하려면 **지금 재테스트** 옵션을 선택하여 수동 테스트를 수행할 수 있습니다. 이렇게 하면 Snyk Code가 저장소의 최신 스냅숏을 찍고 소스 코드 파일을 분석하게 됩니다. 그 결과는 코드 분석 페이지에 표시됩니다. Snyk는 수동 테스트를 새로운 테스트로 계산합니다. [테스트로 간주하는 것은 무엇입니까?](../../../working-with-snyk/what-counts-as-a-test.md)를 참조하십시오.

또한 `.snyk` 파일의 제외 규칙을 가져온 저장소에 적용하기 위해 **지금 재테스트** 옵션을 사용할 수 있습니다. [프로젝트 가져오기에서 디렉토리 및 파일 제외](../../import-project-repository/exclude-directories-and-files-from-project-import.md)를 참조하십시오.

<figure><img src="../../../.gitbook/assets/Retest Code.png" alt="저장소 재테스트"><figcaption><p>저장소 재테스트</p></figcaption></figure>

## 프로젝트 필터

Snyk 웹 UI의 프로젝트 페이지에는 Snyk 프로젝트를 분류하고 각 기준에 맞는 프로젝트 수를 표시하는 필터 창이 있습니다. [프로젝트 정보](../../../snyk-admin/snyk-projects/project-information.md)를 참조하십시오.

**파일 또는 취약점 유형으로 그룹화** 기능은 다음 추가 옵션을 제공합니다:

* **파일별 그룹화**: 이 옵션은 여러 취약점을 포함하는 특정 파일을 식별하여 문제가 있는 파일에 집중하거나 보다 엄격한 검토나 리팩터링이 필요한 파일에 집중할 수 있도록 돕습니다.
* **취약점 유형별 그룹화**: 이 옵션은 SQL Injection 또는 Cross-Site Scripting (XSS)와 같이 취약점을 유형별로 분류하여 코드베이스 내에서 가장 일반적인 취약점 유형을 해결하는 데 도움이 됩니다.

자세한 예제 및 사용 시나리오에 대해서는 Snyk 사용자 문서의 각 섹션을 참조하십시오.

## 취약점 문제

다음 옵션을 사용하여 코드 분석 페이지에서 문제의 표시를 변경할 수 있습니다:

### **파일** 또는 **취약점 유형**으로 그룹화

다수의 문제가 있는 문제 파일을 식별하거나 자주 발생하는 취약점 유형을 해결합니다. 이 필터링 옵션을 사용하여 취약성이 집중적으로 발생할 가능성이 있는 위치를 결정할 수 있습니다.

### 심각도 수준별로 정렬

가장 높은 심각도 수준의 문제부터 낮은 심각도 수준의 문제가 표시되도록 취약점 문제를 정렬합니다.

### 기준별로 취약성 필터링

다음 표에 나열된 다른 기준에 따라 발견된 취약점 문제를 필터링합니다.

| 취약점 문제 필터                                                                                                     | 설명                                                                                                       |
| ------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| [심각도 수준](../../../manage-risk/prioritize-issues-for-fixing/severity-levels.md)                                | 특정 심각도 수준의 문제를 표시합니다. Snyk Code는 **High**, **Medium**, 그리고 **Low** 심각도 수준을 사용하며 **Critical**은 사용하지 않습니다. |
| [우선 순위 점수](../../../manage-risk/prioritize-issues-for-fixing/priority-score.md#calculation-of-priority-score) | 특정 우선 순위 점수 범위에서 문제를 표시합니다.                                                                              |
| 상태                                                                                                            | **열린** 문제나 **무시된** 문제를 표시합니다.                                                                            |
| [언어](../../../supported-languages-package-managers-and-frameworks/)                                           | 분석된 저장소에서 발견된 특정 언어로 작성된 코드 파일에서 발견된 문제를 표시합니다. 필터 창에는 분석된 저장소에서 발견된 프로그래밍 언어만 표시됩니다.                    |
| [취약점 유형](../snyk-code-security-rules/)                                                                        | 특정 취약점 유형과 관련된 문제를 표시합니다. [Snyk 코드 보안 규칙](../snyk-code-security-rules/)을 참조하십시오.                         |

<figure><img src="../../../.gitbook/assets/Vulnerability issues.png" alt="취약점 문제 필터링, 정렬 및 그룹화 개요"><figcaption><p>취약점 문제 필터링, 정렬 및 그룹화</p></figcaption></figure>

## **코드 취약점 스캔**

코드 취약점을 스캔하고 관리하기 위해 다음 작업을 확인할 수 있습니다.

### 저장소에서 취약점 보기

1. Snyk 웹 UI에 로그인하고 [그룹 및 조직](../../../snyk-admin/groups-and-organizations/)을 선택합니다.
2. **프로젝트**로 이동하여 저장소 프로젝트를 포함하는 Target 폴더를 선택합니다.
3. **코드 분석** 프로젝트를 열어 Snyk Code에 의해 감지된 모든 취약점 문제를 확인합니다.

결과를 이해하려면 [코드 분석의 세분화](breakdown-of-code-analysis.md)를 참조하십시오.

### 추가 저장소 가져오기

Snyk 계정에 기존 프로젝트가 있는 경우 Snyk이 테스트할 추가 저장소를 추가할 수 있습니다. [Snyk에 저장소 가져오기](../import-project-with-snyk-code.md#import-repository-to-snyk)를 참조하십시오.

### 테스트에서 저장소 제거

더 이상 취약점을 테스트할 필요가 없는 경우 코드 분석 프로젝트를 제거하거나 가져온 저장소를 삭제할 수 있습니다. [가져온 저장소 제거](../../import-project-repository/remove-imported-repository-from-a-project.md)를 참조하십시오.

### 디렉토리 및 파일 제외

특정 파일 및 디렉토리를 Snyk Code에 의해 가져오지 않도록 제외하려면 저장소에 `.snyk` YAML 정책 파일을 생성해야 합니다. [프로젝트 가져오기에서 디렉토리 및 파일 제외](../../import-project-repository/exclude-directories-and-files-from-project-import.md)를 참조하십시오.

개방 소스 종속성 스캔(SCA)에서만 지원되는 기능으로서, Git 저장소를 통해 저장소를 가져오는 동안 제외 대화상자를 사용하여 가져오기를 제외할 디렉토리를 지정할 수 있습니다.

### 저장소 외부 링크 열기

통합된 Git 저장소 플랫폼에서 저장소에 액세스하려면 코드 분석 프로젝트로 이동하여 저장소 이름을 선택하십시오.

<figure><img src="../../../.gitbook/assets/Open repository external link.png" alt="외부 저장소 링크 개요"><figcaption><p>외부 저장소 링크</p></figcaption></figure>

### 프로젝트 기록 보기

**코드 분석** 프로젝트의 **History** 페이지에 결과 기록이 표시됩니다. 이 페이지에서는 테스트가 수행될 때 가져온 스냅숏이 표시됩니다. 모든 테스팅 단계에 대한 Snyk Code 테스트 결과를 검토할 수 있습니다. [가져오기부터 재테스트까지의 코드 테스트](./#code-testing-from-import-to-retest)를 참조하십시오.

**History** 페이지에서는 두 가지 명확하게 다른 스냅숏만 표시됩니다. 스냅숏이 고유하게 간주되는 경우는 마지막 평가 이후에 저장소나 관련 취약점 결과가 변경된 경우로, 이에 따라 변화를 보여주는 스냅숏이 존재합니다. 최근 테스트 이후에 저장소나 취약점 결과에 변경이 없었을 경우, 새로운 스냅숏은 이전과 동일한 것을 복제할 것입니다. 따라서 이는 **History** 페이지에서 추가 테스트 실행으로 나열될 것입니다. 다만 페이지가 여러 테스트 항목을 제공할 수 있으나 최대 두 항목이 다른 결과를 제공할 것입니다.

프로젝트 기록 보기:

1. Snyk 웹 UI에 로그인하고 [그룹 및 조직](../../../snyk-admin/groups-and-organizations/)을 선택합니다.
2. **프로젝트**로 이동하여 저장소 프로젝트를 포함하는 Target 폴더를 선택합니다.
3. **코드 분석** 프로젝트를 열고 **History**로 이동합니다.
4. 목록에서 테스트를 선택하여 프로젝트의 이력 스냅숏을 확인합니다.
5. (옵션) **최신 스냅숏 보기** 선택. 이 옵션은 최신 스냅숏을 열 경우 사용할 수 없습니다.

<figure><img src="../../../.gitbook/assets/Project history.png" alt="Snyk Code 프로젝트 기록"><figcaption><p>Snyk Code 프로젝트 기록</p></figcaption></figure>

### 프로젝트 설정 관리

다음과 같이 프로젝트 설정을 관리합니다:

* 주기적인 테스트 예약: [테스트 및 자동 풀 리퀘스트 빈도 설정](../../../snyk-admin/snyk-projects/#test-frequency-settings).
* 프로젝트 ID 검색: [프로젝트의 고유 식별자 검색](../../../snyk-admin/snyk-projects/#project).
* 프로젝트 비활성화: [데이터를 삭제하지 않고 프로젝트 일시적으로 비활성화](../../../snyk-admin/snyk-projects/#delete-activate-or-deactivate).
* 프로젝트 삭제: [프로젝트 및 관련 데이터 영구적으로 제거](../../../snyk-admin/snyk-projects/#delete-activate-or-deactivate).

## 다음 단계는?

* [코드 분석 세분화 보기](breakdown-of-code-analysis.md)
* [코드 취약점 자동으로 수정](fix-code-vulnerabilities-automatically.md)
* [Snyk 코드 보안 규칙 보기](../snyk-code-security-rules/)
