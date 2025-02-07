# 동일한 네임스페이스에 여러 Brokers 배포하기

## Helm 차트 버전 2.x.x

Helm 차트 버전 2.x.x는 릴리스 이름을 기반으로 모든 오브젝트 이름에 접미사를 자동으로 추가합니다. 동일한 네임스페이스에 추가할 각 브로커에 대해 다른 릴리스 이름을 사용하면 됩니다.

{% hint style="info" %}
호환성을 유지하기 위해 2.1.0 버전은 단순히 --set disableSuffixes=true 또는 disableSuffixes=true를 values 파일에 추가하여 아래 나열된 1.x.x 동작으로 되돌릴 수 있는 disableSuffixes 플래그를 도입했습니다.
{% endhint %}

## Helm 차트 버전 1.x.x

기존 브로커와 동일한 네임스페이스에 추가적인 브로커를 배포하려면 다음 예제를 따르세요.

### 기존 서비스 계정과 함께 배포

```
helm install <고유_차트_이름_입력> snyk-broker/snyk-broker \
             --set scmType=github-com \
             --set brokerToken=<브로커_토큰_입력> \
             --set scmToken=<저장소_토큰_입력> \
             --set brokerClientUrl=<브로커_클라이언트_URL>:<브로커_클라이언트_포트_입력> \
             --set serviceAccount.create=false \
             --set serviceAccount.name=<기존_서비스_계정> \
             -n <기존_네임스페이스>
```

### 새로운 서비스 계정과 함께 배포

```
helm install <고유_차트_이름_입력> snyk-broker/snyk-broker \
             --set scmType=github-com \
             --set brokerToken=<브로커_토큰_입력> \
             --set scmToken=<저장소_토큰_입력> \
             --set brokerClientUrl=<브로커_클라이언트_URL>:<브로커_클라이언트_포트_입력> \
             --set serviceAccount.name=<생성될_새로운_서비스_계정> \
             -n <기존_네임스페이스>
```
