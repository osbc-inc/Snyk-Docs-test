# 오픈 소스용 Go

{% hint style="warning" %}
2023년 1월 1일부터, Snyk는 govendor 프로젝트를 지원하지 않았습니다. 일반적인 보안 최선의 실천으로, Snyk은 지속적으로 유지 및 최신 상태를 유지하는 도구를 사용하는 것을 권장합니다.

Snyk이 govendor 프로젝트를 더 이상 지원하지 않기 때문에 경고가 발생하고 결과가 제공되지 않습니다.
{% endhint %}

## Snyk 오픈 소스 지원하기

{% hint style="info" %}
**지원되는 Go 버전**

Snyk은 Go의 모든 버전을 지원하며, Go [모든 릴리스](https://go.dev/dl/) 페이지에서 나열된 최신 안정 버전을 포함합니다.
{% endhint %}

지원되는 패키지 관리자 및 피처에 대한 자세한 내용은 [Go 세부 정보](./)를 참조하십시오.

{% hint style="info" %}
공식 릴리스만이 추적됩니다. 커밋(기본 브랜치를 포함)은 공식 릴리스 또는 태그에 포함되지 않는 한 식별되지 않습니다.

패키지 관리자가 있는 프로젝트의 경우, 패키지 관리자에 대한 릴리스를 의미합니다.

Go 및 관리되지 않는 스캔(C/C++)의 경우, GitHub 리포지토리에 공식 릴리스 또는 태그가 있어야 합니다.
{% endhint %}

## Go Modules 및 dep 지원

{% hint style="info" %}
**기능 사용 가능**\
일부 기능은 귀하의 요금제에 따라 사용 가능하지 않을 수 있습니다. 자세한 정보는 [요금제 및 가격](https://snyk.io/plans/)을 참조하십시오.
{% endhint %}

Snyk은 [Go Modules](https://golang.org/ref/mod) 및 [dep](https://github.com/golang/dep)를 통해 관리되는 종속성을 가진 Go 프로젝트의 테스트 및 모니터링을 지원합니다.

<table><thead><tr><th width="167">패키지 관리자/기능</th><th width="126">CLI 지원</th><th width="147">SCM 지원</th><th width="160">라이선스 스캔</th><th>고치기 PRs</th></tr></thead><tbody><tr><td><a href="https://golang.org/ref/mod">Go Modules</a></td><td>✔︎</td><td>✔︎</td><td>✔︎</td><td></td></tr><tr><td><a href="https://github.com/golang/dep">dep</a></td><td>✔︎</td><td>✔︎</td><td>✔︎</td><td></td></tr></tbody></table>

### **Go Modules 및 CLI**

Snyk은 Go Modules 프로젝트를 CLI에서 패키지 수준에서 스캔하며, Snyk은 로컬 소스 코드에 완전한 액세스 권한을 갖기 때문에 모듈 수준이 아닙니다.

[Go 표준 라이브러리](https://pkg.go.dev/std)의 패키지는 지원되지 않으며 의존성 트리에 포함되지 않습니다.

주요 Go 트리 외부에 있는 [Go 프로젝트의 일부인](https://pkg.go.dev/golang.org/x) `golang.org/x/` 하위의 패키지는 지원됩니다.

의존성 트리를 빌드하기 위해 Snyk는 `go list -json -deps ./...` 명령 및 `Imports`에서 찾은 종속성을 사용합니다.

`TestImports` 및 `XTestImports`는 지원되지 않습니다.

CLI를 사용하여 Go Modules 프로젝트를 테스트할 때, Snyk은은 의존성이 설치되었는지 여부는 확인하지 않지만 프로젝트 루트에 `go.mod` 파일이 있어야 합니다. `go list`는 이를 사용하여 프로젝트 소스 코드를 통해 완전한 종속성 트리를 빌드합니다.

Go의 각 버전은 `go list -json -deps` 명령에 대해 다른 결과를 생성합니다. 이는 종속성 트리와 Snyk CLI가 발견하는 취약점에 영향을 줄 수 있습니다.

### **Dep 및 CLI**

의존성 트리를 빌드하기 위해 Snyk은 `Gopkg.lock` 파일을 분석합니다.

CLI를 사용하여 dep 프로젝트를 테스트할 때, Snyk는 의존성을 설치해야 합니다. 이를 위해 `dep ensure`를 실행하십시오.

### **Dep 및 Git**

의존성 트리를 빌드하기 위해 Snyk은Git 리포지토리의 `Gopkg.lock` 파일을 분석합니다.

### **Go Modules 및 Git**

기본적으로, Git을 사용하여 가져온 Go Modules 프로젝트의 경우 CLI에서 테스트된 프로젝트와 마찬가지로 **패키지** 수준이 아닌 **모듈** 수준에서 의존성이 해결됩니다. 따라서 Git을 사용하여 가져올 때 CLI와 달리 이를 포함한 더 많은 종속성과 문제가 보고될 수 있습니다.

최적의 해결책을 얻으려면 [전체 소스 코드 분석을 활성화](go-for-open-source.md#enable-full-source-code-analysis)하십시오.

전체 소스 코드 분석을 활성화하면, Snyk은 `go list -json -deps ./...` 명령을 사용하여 완전한 의존성 트리를 빌드합니다. 그렇지 않으면 `go mod graph`를 사용합니다.

#### 전체 소스 코드 분석 활성화

Git에서 가져온 Go Modules 프로젝트에 대해 가능한 가장 정확한 의존성 트리를 빌드하려면, Snyk은귀하의 리포지토리의 모든 파일에 액세스해야 합니다.

이를 통해 Snyk은 귀하의 `.go` 소스 파일에서 `import` 문을 볼 수 있으며, 응용 프로그램에서 사용하는 특정 패키지를 결정할 수 있습니다. 이에 액세스하지 않으면, Snyk은귀하의 `go.mod` 파일에 나열된 모듈의 모든 패키지를 포함할 것입니다.

전체 소스 코드 분석을 활성화하려면 다음과 같이 설정을 조정하십시오:

1. 계정에 로그인하고 그룹 및 조직을 선택하십시오.
2. **설정**으로 이동한 다음 **언어**를 선택하십시오.
3. **Go**에 대한 **편집 설정**을 선택하십시오.
4. 전체 소스 코드 분석을 켜거나 끄려면 토글을 사용합니다.

<figure><img src="../../.gitbook/assets/image (149) (1).png" alt=""><figcaption><p>전체 소스 코드 분석 활성화</p></figcaption></figure>

다양한 Snyk 기능이 필요로 하는 귀하의 리포지토리의 액세스 수준에 대해 자세한 내용은 [Snyk 가귀하의 데이터를 처리하는 방법](../../working-with-snyk/how-snyk-handles-your-data.md)을 참조하십시오.

#### **비공개 모듈**

메인 프로젝트 리포지토리와 동일한 Git 조직 내의 비공개 Git 리포지토리에서 모듈을 의존하는 Go 모듈 프로젝트가 지원됩니다.

다른 Git 조직의 리포지토리에 비공개 모듈이 있는 경우, 프로젝트 가져오기가 제대로 작동하지 않을 수 있습니다. 코드가 다른 조직의 Git 서브모듈을 사용하는 경우도 마찬가지입니다.

비공개 모듈에 다른 Git 조직의 다른 비공개 모듈이 있는 경우 프로젝트 가져오기가 작동하지 않습니다. 모두 동일한 Git 조직의 일부여아 한다는 것은 주요 프로젝트 리포지토리와 동일한 조직 내의 다른 모듈을 포함한 모든 비공개 모듈이여아 한다는 것을 의미합니다.

비공개 모듈 지원은 [전체 소스 코드 분석](go-for-open-source.md#enable-full-source-code-analysis)이 활성화된 상태에 따라 다른 SCM에서 다릅니다.

| 전체 소스 코드 분석이 활성화되어 있는 경우                                                                                                               | 전체 소스 코드 분석이 비활성화되어 있는 경우                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| <ul><li>Azure Repos</li><li>Bitbucket Cloud</li><li>Bitbucket Server</li><li>GitHub</li><li>GitLab</li><li>GitHub Enterprise</li></ul> | <ul><li>Bitbucket Cloud</li><li>GitHub</li><li>GitHub Enterprise</li></ul> |

### **GO에 대한 Snyk 브로커 지원**

{% hint style="warning" %}
Snyk 브로커는 [전체 소스 코드 분석](go-for-open-source.md#enable-full-source-code-analysis)이 비활성화된 경우에만 지원됩니다.
{% endhint %}

새로운 [Snyk 브로커](../../enterprise-setup/snyk-broker/) 클라이언트를 사용하여 가져온 Go Modules 프로젝트는 예상대로 작동해야 합니다.

2020년 12월 30일 이전에 생성된 클라이언트에 지원을 추가하려면, 변경 사항에 따라 `accept.json` 파일에 `go.mod` 및 `go.sum`를 추가하십시오. ([풀 리퀘스트](https://github.com/snyk/broker/pull/299/files) 참조).

브로커를 통합한 개인 Go Modules를 사용하는 경우, 각 개인 모듈에 정의된 `go.mod` 파일이 있어야 합니다.

####
