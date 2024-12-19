# Custom Rules 프로젝트 생성

## Custom Rules 프로젝트란 무엇인가요?

Snyk CLI를 사용하여 Custom Rules 프로젝트를 생성할 수 있습니다. 이는 사용자 정의 규칙 및 규칙 사양을 위한 디렉터리 구조입니다. CLI는 테스트, 빌드, 번들링 및 사용자 정의 규칙을 업로드하기 위해 해당 구조를 활용합니다.

Custom Rules 프로젝트를 생성하려면, 빈 프로젝트를 만들거나 [snyk-labs/iac-to-cloud-example-custom-rules](https://github.com/snyk-labs/iac-to-cloud-example-custom-rules) 리포지토리와 같은 미리 작성된 프로젝트를 복제할 수 있습니다.

### 옵션 1: Snyk CLI를 사용하여 빈 프로젝트 생성

다음 명령을 사용합니다:

<pre><code><strong>snyk iac rules init
</strong></code></pre>

이 명령은 대화식 프롬프트를 도입하고 프로젝트, 규칙, 규칙 사양(테스트용) 및 리소스 관련을 설정할 수 있습니다. `project`를 선택하고 이후의 프롬프트에 응답하여 디렉터리 구조를 생성할 수 있습니다.

**대화식 프롬프트에서 예시 출력:**

```
What do you want to initialize? project
Project name: project1
```

### 옵션 2: IaC to Cloud 예제 Custom Rules 리포지토리 가져오기

```
git clone https://github.com/snyk-labs/iac-to-cloud-example-custom-rules.git
```