# 명령 줄을 사용한 IaC 제외

Snyk CLI를 사용하여 `iac test` 명령으로 디렉터리나 대량의 IaC 파일을 스캔할 때 실수로 스캔에 원치 않는 파일이나 디렉터리가 포함될 수 있습니다. 이런 경우 명령 줄 도구를 사용하여 스캔에서 특정 파일이나 디렉터리를 제외할 수 있습니다. 이 문서에서는 가장 일반적인 사용 사례에 대한 해결책을 설명합니다.

{% hint style="info" %}
다음 예제에서는 `find`와 `xargs`와 같은 명령 줄 도구를 사용합니다. 이 도구들이 UNIX와 유사한 환경에서 보편적으로 사용됩니다. 사용하는 플랫폼에서 도구들이 제대로 작동하는지 확인해야 합니다.
{% endhint %}

## 잘못된 유형의 파일 제외

프로젝트에는 다양한 유형의 파일이 포함될 수 있는데, 특정 확장자를 가진 파일만을 스캔하고 나머지는 제외하고 싶을 수 있습니다. 다음 명령은 현재 작업 디렉터리와 그 하위 디렉터리에 포함된 `.tf` 확장자를 가진 파일만을 스캔합니다. `.tf` 확장자가 없는 파일은 스캔되지 않습니다.

```
find . -type file -name '*.tf' | xargs snyk iac test
```

## 이름으로 디렉터리 제외

매우 큰 프로젝트에서는 서로 다른 목적을 가진 파일들을 다른 디렉터리에 저장하는 것이 일반적입니다. 예를 들어 개발, 스테이징, 및 프로덕션 환경용 IaC 파일을 서로 다른 디렉터리에 저장할 수 있습니다. 이러한 프로젝트를 스캔할 때 일부 디렉터리를 제외하고 싶을 수 있습니다. 이 명령은 현재 디렉터리와 그 하위 디렉터리에 있는 `.tf` 확장자를 가진 모든 파일을 스캔하지만 `prod` 하위 디렉터리에 있는 파일은 제외합니다.

```
find . -name '*.tf' -not -path './prod/*' | xargs snyk iac test
```  