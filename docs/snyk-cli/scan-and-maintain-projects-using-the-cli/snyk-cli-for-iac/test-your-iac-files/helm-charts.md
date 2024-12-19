# Helm 차트

Helm 차트를 스캔하기 위해서는 Helm 템플릿을 Kubernetes 매니페스트 파일로 렌더링한 다음, 이를 Snyk CLI의 `snyk iac` 명령어를 사용하여 스캔합니다.

예를 들어, `./helm` 디렉토리에 위치한 Helm 프로젝트가 있다면, 다음 명령어를 실행하여 템플릿 파일을 `./output` 디렉토리에 출력할 수 있습니다:

{% tabs %}
{% tab title="macOS/Linux/Unix" %}
```
helm template ./helm --output-dir ./output
snyk iac test ./output
```
{% endtab %}

{% tab title="Windows PowerShell" %}
```
helm template .\helm\ --output-dir .\output\
snyk iac test .\output\
```
{% endtab %}
{% endtabs %}

Unix 기반의 터미널에서는 `helm template`의 출력을 직접 한 파일로 보내는 것도 가능합니다:

```
helm template ./helm > output.yaml
snyk iac test output.yaml
```

Snyk CLI는 현재 표준 입력에서 읽을 수 없습니다. Snyk는 이 기능에 대해 작업 중입니다.