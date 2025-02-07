# .snyk 파일

`.snyk` 파일은 사용자가 로컬로 사용하거나 작업 흐름의 일부로 사용하여 Snyk이 방지하는 문제를 제어하고, 스캔에서 파일을 제외하고, 프로젝트 수준에서 Python 버전을 설정하며, CLI 및 CI/CD 플러그인에 대한 패치를 지정하는 능력입니다.

`.snyk` 파일은 Snyk 제품들 사이에서 작동 방식이 다릅니다. `.snyk` 파일을 배포할 때는 파일 생성 방법, 사용 가능한 위치 및 사용 목적을 검토하며 시작합니다. 자세한 내용은 [Snyk 오픈 소스와 함께 `.snyk` 파일 사용](the-.snyk-file.md#use-the-.snyk-file-with-snyk-open-source), [Snyk 코드에서 `.snyk` 파일 사용](the-.snyk-file.md#use-the-.snyk-file-with-snyk-code) 및 [Snyk IaC에서 `.snyk` 파일 사용](the-.snyk-file.md#use-the-.snyk-file-with-snyk-iac)을 참조하십시오.

`.snyk` 파일은 `snyk ignore` 명령을 사용하여 생성할 수 있습니다. 이렇게 하면 파일과 무시 규칙이 생성됩니다. 또는 텍스트 또는 코드 편집기를 사용하여 파일을 생성할 수도 있습니다. 형식은 YAML입니다. 자세한 내용은 [`.snyk` 파일 생성 방법](the-.snyk-file.md#how-to-create-the-.snyk-file)을 참조하십시오.

`.snyk` 파일 사용 시 고려해야 할 중요 사항:

* CLI는 자동으로 `.snyk` 파일을 사용하며, 제품이 데이터베이스에 무시를 지원하거나 Snyk 규칙을 사용하는 경우 웹 UI에서 생성된 무시 규칙과 함께 사용합니다.
* 빌드 시스템의 일부로 사용되는 CI/CD 플러그인이나 CLI는 파일이 존재하는 경우 스캔 도중에 `.snyk` 파일을 사용합니다.
* `.snyk` 파일을 코드와 병합하면 SCM을 Snyk에 가져올 때, `.snyk` 파일의 규칙이 웹 UI에서 생성된 데이터베이스 규칙 위에 적용됩니다.
* 무시를 지정하기 위해 `.snyk` 파일을 사용하면 문제가 감지되고 모니터링된 후에만 가능한 Snyk 웹 UI에서 무시를 지정할 필요가 없습니다. Snyk 데이터베이스의 무시 규칙을 무시하려면 `.snyk` 파일을 사용할 수 있습니다. 자세한 내용은 [데이터베이스의 무시 규칙 재정의하는 방법](the-.snyk-file.md#how-to-override-the-ignore-rules-in-the-database)을 참조하십시오.

## `.snyk` 파일 생성 방법

`snyk ignore` 명령을 사용하여 `.snyk` 파일을 생성할 수 있습니다. 자세한 내용은 [무시](../../snyk-cli/commands/ignore.md) 명령 CLI 도움말을 참조하십시오.

기존 `.snyk` 파일이 없는 경우 다음 코드로 파일을 생성하고 채울 수 있습니다:\
`# Snyk (https://snyk.io) policy file, patches or ignores known vulnerabilities` `version: v1.25.0`

* 현재 정책 스키마 버전이므로 `version`을 `v1.25.0`으로 설정해야 합니다.
* 무시 블록(들)은 [무시 명령 설명서](../../snyk-cli/commands/ignore.md#description)나 이 페이지의 예제에 표시된 관련 구문을 따라야 합니다.

자세한 내용은 [`.snyk` 파일의 구문](the-.snyk-file.md#syntax-of-the-.snyk-file)을 참조하십시오.

일반적으로 `.snyk` 파일은 코드 저장소의 루트에 생성되어야 합니다. 그러나 SCM 가져오기의 경우 `.snyk` 파일은 스캔에 필요한 모든 파일이 있는 디렉토리와 동일한 디렉토리에 있어야 합니다. 예를 들어 매니페스트 파일과 관련이 있는 경우입니다. [단일 저장소 및 복잡한 프로젝트에서 `.snyk` 파일 사용](the-.snyk-file.md#use-the-.snyk-file-with-monorepos-and-complex-projects)을 참조하십시오.

Git 저장소 오픈 소스 스캔에서 **취약성 수정** 버튼을 클릭하고 Snyk 패치가 가능하며 업그레이드가 불가능한 경우 `.snyk` 파일이 추가되어 패치를 지정합니다. Snyk 패치는 npm 및 Yarn에만 지원됩니다.

다음 예시는 취약성 수정 PR을 사용하여 패치 규칙을 생성하는 `.snyk` 파일을 작성하는 방법을 보여줍니다:

```
# Snyk (https://snyk.io) 정책 파일, 패치 또는 이미 알려진 취약점 무시
version: v1.25.0
ignore: {}
# 패치는 취약점을 수정하는 데 필요한 최소한의 변경을 적용합니다
patch:
  'npm:hawk:20160119':
    - tap > codecov.io > request > hawk:
        patched: '2020-01-20T14:26:34.404Z'
```

Snyk은 또한 [snyk-policy 패키지](https://www.npmjs.com/package/snyk-policy)를 제공하여 일반적으로 `.snyk`로 명명된 정책 파일을 생성할 수 있습니다.

{% hint style="info" %}
패키지의 버전은 `.snyk` 파일에 입력해야 하는 정책 스키마 버전과 동일하지 않습니다.
{% endhint %}

## Snyk 코드와 함께 `.snyk` 파일 사용

Snyk 코드 스캔에서 스캔된 파일과 코드 분석 프로젝트를 생성하는 `exclude from import` 옵션을 사용할 파일 또는 디렉터리를 지정하기 위해 `.snyk` 파일을 사용할 수 있습니다. `exclude from import` 옵션은 단지 Snyk 코드에서 지원되며, Snyk 웹 UI 및 CLI를 통해 수행되는 가져오기에만 사용할 수 있습니다.

`snyk monitor` 명령을 사용하는 대신 코드 저장소 통합을 사용하여 가져오는 프로젝트의 경우, `--policy-path` 옵션이 사용할 수 없습니다. `.snyk` 파일은 `.snyk` 파일이 있는 경로와 동일한 경로에서 발견된 프로젝트에만 적용됩니다.

자세한 내용은 [import 프로세스에서 디렉터리 및 파일 제외](../../scan-with-snyk/import-project-repository/exclude-directories-and-files-from-project-import.md)을 참조하십시오.

## Snyk IaC와 함께 `.snyk` 파일 사용

IaC 무시 규칙에 대해서는 [.snyk 정책 파일을 사용한 IaC 무시](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/iac-ignores-using-the-.snyk-policy-file.md)를 참조하십시오. 더 많은 정보는 [드리프트에 대한 자원 무시](../../scan-with-snyk/snyk-iac/iac+-code-to-cloud-capabilities/detect-drift-and-manually-created-resources/ignore-resources-for-drift.md)를 참조하십시오.

## Snyk Open Source와 함께 `.snyk` 파일 사용

프로젝트의 `.snyk` 파일은 `snyk test`와 `snyk monitor` 명령 및 API 또는 Snyk 웹 UI를 통해 수행되는 테스트에 대한 무시 및 기타 설정을 적용하는 데 사용됩니다. `.snyk` 파일은 빌드 시간에 적용할 Snyk 패치를 정의하여 업그레이드로 수정할 수 없는 취약점을 해결하고, `snyk protect` 명령을 대체하는 `@snyk/protect` [패키지](https://github.com/snyk/snyk/tree/master/packages/snyk-protect)를 적용하며, Python (Pip) 버전을 위한 `language settings:`를 적용합니다.

### 오픈 소스 프로젝트와 함께 `.snyk` 파일 작동 방식

Snyk은 SCM 통합, Snyk CLI 및 CI/CD 통합을 통한 스캔 시 Snyk 데이터베이스 및 `.snyk` 파일의 무시 규칙을 확인합니다.

프로젝트에 `.snyk` 파일이 있는 경우, `snyk test` 명령은 해당 파일을 무시 메커니즘으로 사용하며, 웹 UI에서 설정된 무시 대신에 사용합니다.

`.snyk` 파일이 SCM 프로젝트에 포함되어 있으면 Snyk는 데이터베이스 무시 및 `.snyk` 무시를 모두 고려합니다.

`.snyk` 파일을 코드 저장소에 포함하고 `language-settings:` 값을 설정하면 코드 저장소 스캔 시 프로젝트 수준 Python 설정을 생성할 수 있습니다.

* 예를 들어 GitHub 스캔에서는, Snyk 웹 UI에서 **조직 > 설정 > Snyk 오픈 소스 > Python > Pip Python 버전** 옵션으로 조직 수준에서 Python 버전을 제어합니다.
* `language settings:` 값을 사용 가능한 UI 언어 설정 중 하나로 설정하여 코드 저장소 스캔에서 해당 저장소의 SCM 스캔에 사용할 Python 버전을 조직 수준 설정을 무시하고 UI 옵션 중 사용할 수 있게 설정할 수 있습니다.

{% hint style="info" %}
만일 Snyk에 프로젝트를 최초 가져오 때 `.snyk` 파일이 존재하지 않았다면 프로젝트를 다시 가져와야 합니다.
{% endhint %}

Python 버전 지원에 대한 자세한 내용은 [Python 버전 지원](../../supported-languages-package-managers-and-frameworks/python/#python-version-support)을 참조하십시오.

오픈 소스 프로젝트와 `.snyk` 파일을 사용하는 방법에 대해 더 많은 정보를 원하시면 다음을 참조하십시오:

* [CLI를 사용하여 취약점 무시](https://docs.snyk.io/snyk-cli/fix-vulnerabilities-from-the-cli/ignore-vulnerabilities-using-snyk-cli)
* [오류 메시지: CLI를 통한 무시가 이 조직에서는 활성화되지 않았습니다. 문제를 웹사이트를 통해 불평하시길 바랍니다](https://support.snyk.io/s/article/Error-message-Ignoring-via-the-CLI-is-not-enabled-for-this-organization-Please-ignore-issues-via-our-website)

### 오픈 소스를 위한 `.snyk` 파일 예시

#### Python을 위한 언어 버전 설정

프로젝트에 Python 3.7을 지정하기 위해 `.snyk` 파일을 수동으로 수정합니다:

```
# Snyk (https://snyk.io) 정책 파일, 패치 또는 이미 알려진 취약점 무시.
version: v1.25.0
language-settings: 
  python: "3.7"
```

자세한 내용은 [Git 프로젝트에서 Python 버전 설정](../../supported-languages-package-managers-and-frameworks/python/#setting-python-version-in-git-projects)을 참조하십시오.

#### 취약점 무시 규칙 설정

{% hint style="warning" %}
`expires`필드는 옵션입니다. 영구적으로 무시해야 하는 경우 이 필드를 생략하십시오.
{% endhint %}

특정 경로의 특정 취약점을 무시합니다:

<pre><code><strong>ignore:
</strong>  SNYK-JS-BSON-561052:
    - mongodb > mongodb-core > bson:
        reason: None given
        expires: '2020-06-19T20:36:54.553Z'
</code></pre>

모든 경로에 대해 취약점을 무시합니다:

```
ignore:
  SNYK-JS-BSON-561052:
    - '*':
        reason: None Given
        expires: 2020-04-04T17:33:45.004Z
```

여러 경로에서 특정 취약점을 무시합니다:

<pre><code><strong>ignore:
</strong><strong>  SNYK-JS-DOTPROP-543489:
</strong>    - configstore > dot-prop:
        reason: None given
        expires: '2020-06-19T20:36:54.553Z'
    - snyk > configstore > dot-prop:
        reason: None given
        expires: '2020-06-19T20:36:54.553Z'
</code></pre>

#### 라이선스 무시 규칙 설정

패키지에 대한 라이선스 문제를 무시하려면 `snyk test` 명령의 출력에서 라이선스의 ID를 찾습니다.

라이선스 ID는 라이선스 문제 URL의 일부이며, 이러한 URL 예시에서 라이선스 ID는 `snyk:lic:npm:symbol:MPL-2.0`입니다: [https://snyk.io/vuln/snyk:lic:npm:symbol:MPL-2.0](https://snyk.io/vuln/snyk:lic:npm:symbol:MPL-2.0).

### **Snyk CLI 및 `.snyk` 파일을 사용한 Snyk Open Source 관리**  

## **Snyk CLI와 `.snyk` 파일 사용**  

### **사용할 CLI 명령어**  

Snyk CLI는 `.snyk` 파일을 생성하고 조회하는 명령어를 제공합니다.  

- `snyk policy` 명령어는 패키지의 `.snyk` 정책을 표시합니다.  
- `snyk ignore` 명령어는 특정 문제를 무시하도록 `.snyk` 파일을 수정합니다.  

```
snyk ignore --id='vulnerabilityID' --expiry='date-string' --reason='text string'
```  

다음 예제는 `SNYK-JS-BSON-561052` 취약점을 무시하는 규칙을 생성하는 명령어입니다. 이 규칙은 디스크상의 모든 경로에서 해당 라이브러리를 무시하도록 설정합니다.  

```
snyk ignore --id='SNYK-JS-BSON-561052' --expiry='2018-04-01' --reason='testing'
```  

### **Monorepo 및 복잡한 프로젝트에서 `.snyk` 파일 사용**  

Snyk CLI는 `.snyk` 파일이 분석 중인 매니페스트(manifest)에 적용된다고 가정합니다. 그러나 복잡한 프로젝트나 monorepo의 경우 여러 개의 매니페스트가 하위 폴더에 있을 수 있으며, 중앙 집중식 무시 정책을 사용하고 싶을 수도 있습니다. `.snyk` 파일은 프로젝트 루트에 위치해야 하며, 매니페스트 파일과 함께 있어야 합니다.  

만약 `.snyk` 파일이 프로젝트 루트에 존재하지 않는 경우(예: 중앙 집중식 정책 사용), `--policy-path` 옵션을 사용하여 경로를 명시적으로 지정해야 합니다.  

CLI를 사용하여 `.snyk` 무시 정책을 생성했는데 Snyk가 취약점을 성공적으로 무시하지 않는다면, 다음 옵션을 추가하십시오.  

```
--policy-path=/path/path/file
```  

전체 명령어는 다음과 같습니다.  

```
snyk ignore --id=IssueID [--expiry=expiry] [--reason='reason for ignoring'] [--policy-path=/path/path/file]
```  

---

## **데이터베이스의 무시 규칙을 재정의하는 방법**  

프로젝트에 `.snyk` 파일이 있으면 `snyk test` CLI 명령어는 해당 파일을 무시 규칙(ignore mechanism)으로 사용하며, 웹 UI에서 설정한 무시 규칙을 덮어씁니다. 즉, `.snyk` 파일이 있는 프로젝트에서 `snyk test` 명령어를 실행하면 Snyk Web UI에서 설정한 모든 무시 설정이 무시됩니다.  

그러나 `.snyk` 파일이 포함된 프로젝트가 SCM(Source Control Management)에 있는 경우, Snyk는 데이터베이스의 무시 규칙과 `.snyk` 파일의 무시 규칙을 모두 고려합니다.  

### **관리자 사용자만 무시 규칙을 변경하도록 설정하는 방법**  

웹 UI에서 설정된 무시 규칙을 `.snyk` 파일을 통해 재정의하려면, **관리자 사용자만** 무시 규칙을 변경할 수 있도록 설정해야 합니다. 이를 위해 다음 단계를 수행하십시오.  

1. Snyk 계정에 로그인합니다.  
2. **Settings** 메뉴에서 **General**을 선택합니다.  
3. 다음 옵션 중 하나를 선택합니다.  
   - **Admin users only** - 관리자만 무시 설정을 변경할 수 있습니다.  
   - **All users in any environment** - 모든 사용자가 무시 설정을 변경할 수 있습니다.  

---

## **`.snyk` 파일의 문법(Syntax)**  

`.snyk` 파일의 최상위 키는 다음과 같습니다.  

- `language-settings:`  
- `ignore:`  
- `patch:`  

### **`language-settings:`**  
이 값은 현재 사용 중인 Python 버전을 나타냅니다. Python 버전 설정 예제는 [Python 언어 버전 설정](the-.snyk-file.md#set-the-language-version-for-python) 섹션을 참조하십시오.  

### **`ignore:` (무시 규칙)**  
무시 규칙은 다음 형식으로 작성됩니다.  

```
ignore:
  snyk-vulnid:
    - path to library using > separator :
        reason: 'text string'
        expires: 'YYYY-MM-DDThh:mm:ss.fffZ'
```  

- `expires` 필드는 선택 사항입니다.  
- 특정 취약점을 영구적으로 무시하려면 `expires` 필드를 생략하면 됩니다.  

예제:  

```
ignore:
  snyk-vulnid:
    - path to library using > separator :
        reason: 'text string'
```  

### **`patch:` (패치 규칙)**  
패치 규칙은 다음 형식으로 작성됩니다.  

```
'npm:library:yyyymmdd’ :
  - path to library using > separator:
    patched: 'datetime string'
  - path to library using > separator > to > another > path:
    patched: 'YYYY-MM-DDThh:mm:ss.fffZ'
```