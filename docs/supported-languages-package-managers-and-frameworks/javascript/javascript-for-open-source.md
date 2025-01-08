# 오픈 소스용 JavaScript

지원되는 패키지 매니저 및 기능에 대해서는 [JavaScript 세부 정보](./)를 참조하십시오.

{% hint style="info" %}
공식 릴리스만이 추적됩니다. 기본 브랜치에 대한 커밋은 공식 릴리스 또는 태그에 포함되지 않은 한 식별되지 않습니다.

JavaScript 패키지의 경우 npmjs 패키지 레지스트리에 릴리스되어야 합니다.
{% endhint %}

## 오픈 소스 및 라이선스

다음은 npm, pnpm 및 Yarn에 대한 Snyk Open Source 지원과 Lerna에 대한 부분적인 지원 요약입니다.

### npm

다음 표는 npm 잠금 파일 버전 및 Snyk 기능 가용성을 제공합니다.

<table><thead><tr><th>잠금 파일 버전</th><th width="144">CLI 지원</th><th>SCM 지원</th><th>라이선스 스캔</th><th>수정 PR</th></tr></thead><tbody><tr><td>잠금 파일 v1</td><td>✔︎</td><td>✔︎</td><td>✔︎</td><td>✔︎</td></tr><tr><td>잠금 파일 v2</td><td>✔︎</td><td>✔︎</td><td>✔︎</td><td>✔︎</td></tr><tr><td>잠금 파일 v3</td><td>✔︎</td><td>✔︎</td><td>✔︎</td><td>✔︎</td></tr></tbody></table>

#### **피어 종속성**

npm v7 이상에서 **피어 종속성**의 동작이 기본적으로 설치되는 경우 달라집니다. npm v7+ 프로젝트에 맞추기 위해 Snyk는 피어 종속성이 설치되었으며 기본적으로 스캔된 것으로 간주합니다.

npm v7+ 프로젝트에서 피어 종속성은 `package.json`의 `peerDependenciesMeta` 객체에 `optional`로 명시된 경우에만 무시됩니다. 아래와 같이 `cache-manager`에 대해 표시된 것처럼:

```json
{
    ...
    "peerDependenciesMeta": {
        "cache-manager": {
            "optional": true
        }
    },
    ...
}
```

npm v6 이하에서는 피어 종속성은 기본적으로 스캔되지 않으며, 패키지 매니저가 기본적으로 설치하지 않기 때문에입니다. 피어 종속성을 스캔하려면 이를 설치한 후 CLI를 `--peer-dependencies` 옵션과 함께 실행하십시오.

#### **잠금 파일 버전**

프로젝트의 종속성 트리를 생성하기 위해 Snyk는 `package-lock.json` 잠금 파일을 사용합니다. 이러한 잠금 파일은 다른 버전으로 제공됩니다.

잠금 파일 v1은 npm v5와 v6에서 사용되었습니다. npm v7에서 두 가지 새로운 형식이 도입되었습니다 - 잠금 파일 v2 및 잠금 파일 v3. 자세한 내용은 [lockfileVersion](https://docs.npmjs.com/cli/v9/configuring-npm/package-lock-json#lockfileversion)를 참조하십시오.

사용 중인 잠금 파일 형식은 다음과 같이 `package-lock.json`에서 확인할 수 있습니다:

```json
{
    ...
    "lockfileVersion": 3,
    ...
}
```

npm에 특정 잠금 파일 버전을 생성하도록 강제하려면 npm `--lockfile-version` 매개변수를 사용하십시오.

```bash
npm install --lockfile-version=2
```

### **pnpm**

{% hint style="info" %}
**릴리스 상태**

Snyk CLI pnpm 지원은 이른 액세스 상태입니다.

[Snyk Priview](../../snyk-admin/snyk-preview.md)를 사용하여 이를 활성화하고 CLI v1.1293.0 이상을 설치하십시오.
{% endhint %}

다음 표는 pnpm 버전 및 Snyk 기능 가용성 매트릭스를 보여줍니다.

<table><thead><tr><th>pnpm 버전</th><th>CLI 지원</th><th width="151">SCM 지원</th><th>라이선스 스캔</th><th>수정 PR</th></tr></thead><tbody><tr><td>pnpm 7</td><td>✔︎ (이른 액세스)</td><td></td><td>✔︎ (이른 액세스)</td><td></td></tr><tr><td>pnpm 8</td><td>✔︎ (이른 액세스)</td><td></td><td>✔︎ (이른 액세스)︎</td><td></td></tr><tr><td>pnpm 9</td><td>✔︎ (이른 액세스)︎</td><td></td><td>✔︎ (이른 액세스)︎</td><td></td></tr></tbody></table>

**잠금 파일 버전**

Snyk은 프로젝트의 종속성 트리를 생성하기 위해 `pnpm-lock.yaml` 잠금 파일을 사용합니다.

지원되는 잠금 파일 버전은 pnpm 7, 8 및 9에 해당하는 5.4, 6.x 및 9.x입니다.

pnpm 잠금 파일에는 [bundledDependencies](https://docs.npmjs.com/cli/v10/configuring-npm/package-json#bundledependencies)이 포함되어 있지 않으므로 Snyk은 이를 스캔하지 않습니다.

### **Yarn**

Snyk은은 프로젝트 종속성의 표현을 생성하기 위해 Yarn 잠금 파일(`yarn.lock`)을 사용합니다.

프로젝트를 스캔하는 데 사용하는 파일은 패키지 매니저의 버전 업그레이드에 따라 변경될 수 있습니다. Snyk은 내부적으로 지원되는 버전만 나열합니다.

이 페이지에 나열된 버전보다 더 새로운 Yarn 버전을 사용하는 경우 Yarn이 이미 지원되는 잠금 파일 버전을 사용하기 때문에 Snyk이 예상대로 작동하는 것을 발견할 수 있습니다. 그 Yarn 버전은 아마도 평가되지 않았으며, 따라서 이 페이지에 추가되지 않았을 것입니다.

다음 표는 Yarn 버전 및 Snyk 기능 가용성 매트릭스를 보여줍니다.

<table><thead><tr><th>Yarn 버전</th><th width="148">CLI 지원</th><th>SCM 지원</th><th>라이선스 스캔</th><th>수정 PR</th></tr></thead><tbody><tr><td>Yarn 1</td><td>✔︎</td><td>✔︎</td><td>✔︎</td><td>✔︎</td></tr><tr><td>Yarn 2</td><td>✔︎</td><td>✔︎</td><td>✔︎</td><td>✔︎</td></tr><tr><td>Yarn 3</td><td>✔︎</td><td>✔︎</td><td>✔︎</td><td>✔︎</td></tr></tbody></table>

{% hint style="info" %}
Yarn의 다른 버전에는 다른 기능 집합이 있기 때문에 패키지 관리자의 작동 방식과 일치하도록 Snyk의 지원에 차이가 있습니다.

**해결**은 Yarn v2 이상에서 지원됩니다. Yarn v1 해결은 지원되지 않습니다.
{% endhint %}

### Lerna

Snyk은 **Lerna**를 완전히 지원하지는 않습니다. 프로젝트가 Yarn Workspaces를 사용하여 설정된 경우, 프로젝트를 스캔하는 방법은 모든 Yarn Workspaces 프로젝트와 동일하게 스캔할 수 있습니다.

Yarn Workspaces를 사용하여 설정된 Lerna 프로젝트의 경우 다음처럼 `snyk test`와 `snyk monitor`를 실행할 수 있습니다.

예제 패키지마다 다음을 사용할 수 있습니다:

<pre class="language-javascript"><code class="lang-javascript"><strong>snyk monitor --file=packages/example-package/package.json
</strong></code></pre>

또는 중첩된 `package.json` 파일을 자동으로 스캔하는 스크립트를 지정할 수 있습니다:

```javascript
ls packages | xargs -I PKG_NAME snyk monitor --file=packages/PKG_NAME/package.json
```

## npm, pnpm 및 Yarn을 사용한 스캔 시작 단계

다음 표에는 종속성을 스캔하는 단계가 나열되어 있습니다. `snyk test` 및 `snyk monitor`와 같은 기본 명령에 대해 다루고 있습니다. 전체 CLI 명령 목록은 [CLI 명령 및 옵션 요약](../../snyk-cli/cli-commands-and-options-summary.md)을 참조하십시오.

| 패키지 매니저 | 시작 방법                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | 설명                                                                                                                                                                                                                                         |
| ------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| npm     | <ol><li>npm 설치</li><li><code>package.json</code> 및 <code>package-lock.json</code> 파일이 있는 디렉토리, 즉 npm 프로젝트 파일이 있는 디렉토리에 있는지 확인</li><li>(선택 사항) <code>npm install</code> 실행</li><li><a href="../../snyk-cli/cli-commands-and-options-summary.md">Snyk 명령</a> 실행</li><li>(선택 사항) <a href="../../snyk-cli/commands/test.md#options-for-npm-projects">snyk test</a> 및 <a href="../../snyk-cli/commands/monitor.md#options-for-npm-projects">snyk monitor</a>을 위한 명령 옵션 실행</li></ol>                      | <p>Snyk는 종속성 트리를 빌드하기 위해 당신의 <code>package-lock.json</code> 파일을 분석합니다.</p><p><code>package-lock.json</code>이 없는 경우, Snyk는 당신의 <code>node_modules</code> 폴더를 분석합니다.</p><p>대안으로, 먼저 잠금 파일을 생성하기 위해 <code>npm install</code>을 실행할 수 있습니다.</p> |
| pnpm    | <ol><li>pnpm 설치</li><li><code>package.json</code> 또는 <code>pnpm</code> 및 <code>pnpm-lock.yaml</code> 파일이 있는 디렉토리, 즉 pnpm 프로젝트 파일이 있는 디렉토리에 있는지 확인</li><li>(선택 사항) <code>pnpm install</code> 실행</li><li><a href="../../snyk-cli/cli-commands-and-options-summary.md">Snyk 명령</a> 실행</li><li>(선택 사항) <a href="../../snyk-cli/commands/test.md#options-for-npm-projects">snyk test</a> 및 <a href="../../snyk-cli/commands/monitor.md#options-for-npm-projects">snyk monitor</a>을 위한 명령 옵션 실행</li></ol> | <p>Snyk는 당신의 <code>pnpm-lock.yaml</code> 파일을 분석하여 종속성 트리를 빌드합니다.<br><br><code>pnpm-lock.yaml</code>이 없는 경우, Snyk는 당신의 <code>node_modules</code> 폴더를 분석합니다.<br><br>대안으로, 먼저 잠금 파일을 생성하기 위해 <code>pnpm install</code>을 실행할 수 있습니다.</p>       |
| Yarn    | <ol><li>Yarn 설치</li><li><code>package.json</code> 및 <code>yarn.lock</code> 파일이 있는 디렉토리, 즉 Yarn 프로젝트 파일이 있는 디렉토리에 있는지 확인</li><li>(선택 사항) <code>yarn install</code> 실행</li><li><a href="../../snyk-cli/cli-commands-and-options-summary.md">Snyk 명령</a> 실행</li><li>(선택 사항) <a href="../../snyk-cli/commands/test.md#options-for-yarn-projects">snyk test</a> 및 <a href="../../snyk-cli/commands/monitor.md#options-for-yarn-projects">snyk monitor</a>을 위한 명령 옵션 실행</li></ol>                         | <p>Snyk는 당신의 <code>yarn.lock</code> 파일을 분석하여 종속성 트리를 빌드합니다.</p><p><code>yarn.lock</code>이 없는 경우, Snyk는 당신의 <code>node_modules</code> 폴더를 분석합니다.</p><p>대안으로, 먼저 잠금 파일을 생성하기 위해 <code>yarn install</code>을 실행할 수 있습니다.</p>                   |

## 모노레포 및 워크스페이스 지원

Yarn, npm 및 pnpm은 다중 하위 프로젝트를 포함하는 모노레포를 관리하는 데 도움을 주는 워크스페이스를 지원합니다.

다음 CLI 옵션에 대해 워크스페이스가 Snyk CLI에서 지원됩니다:

* `--all-projects` : Yarn, npm 및 pnpm 워크스페이스 프로젝트 및 기타 지원되는 생태계 프로젝트를 발견 및 스캔합니다. 워크스페이스 프로젝트를 스캔할 때 루트 잠금 파일이 참조됩니다.
* `--detection-depth` : 검색할 하위 디렉토리 수를 지정합니다.
* `--strict-out-of-sync=false` : 워크스페이스의 패키지에 대해 동기화되지 않은 잠금 파일을 테스트할 수 있도록 허용합니다. 이 옵션이 `false`로 설정되면, 동기화되지 않은 매니페스트와 잠금 파일을 사용하여 Snyk 테스트를 실행할 수 있습니다.
* `--policy-path` : 테스트 중 Snyk이이 사용할 정책 경로를 지정합니다.

### 워크스페이스 스캔 예제

현재 디렉토리 및 5단계 하위 디렉토리 내의 모든 워크스페이스 프로젝트를 스캔하고, 감지된 기타 프로젝트 유형도 포함합니다.

```bash
snyk test --all-projects --strict-out-of-sync=false --detection-depth=6 
```

모든 감지된 워크스페이스에 대해 적용할 무시 및 패치를 한 곳에서 관리하려면 공통 `.snyk` 정책 파일을 사용할 수 있습니다. 이 파일에 대한 자세한 내용은 [The .snyk file](https://docs.snyk.io/manage-risk/policies/the-.snyk-file)을 참조하십시오.

```bash
snyk test --all-projects --strict-out-of-sync=false --policy-path=src/.snyk
```

### **npm 워크스페이스 예제**

npm v7에서는 워크스페이스를 지원합니다. 자세한 내용은 [lockfile version and Snyk feature availability matrix](javascript-for-open-source.md#npm)을 참조하십시오.

npm 프로젝트의 모든 워크스페이스를 감지하고 스캔하려면 위에서 설명한 CLI 옵션을 사용하십시오.

### **pnpm 워크스페이스 예제**

pnpm 워크스페이스는 루트 디렉토리에 `package.json`, `pnpm-lock.yaml` 및 `pnpm-workspace.yaml` 파일이 있어야 합니다.

pnpm 프로젝트의 모든 워크스페이스를 감지하고 스캔하려면 위에서 설명한 CLI 옵션을 사용하십시오.

### **Yarn 워크스페이스 예제**

{% hint style="info" %}
`nohoist`는 Yarn 워크스페이스에서 지원되지 않습니다.
{% endhint %}

Yarn 프로젝트의 모든 워크스페이스를 감지하고 스캔하려면 모노레포 및 워크스페이스에 대해 정의된 CLI 옵션과 다음 Yarn 전용 옵션을 사용하십시오:

`--yarn-workspaces` : 루트에 락파일이 있는 경우 `--all-projects` 대신 사용하여 Yarn 워크스페이스 프로젝트만 감지하고 스캔합니다. 다른 생태계는 무시됩니다.
