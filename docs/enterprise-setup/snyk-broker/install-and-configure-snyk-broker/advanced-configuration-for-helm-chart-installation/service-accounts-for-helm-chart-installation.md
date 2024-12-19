# Helm 차트 설치를 위한 서비스 계정

기존 서비스 계정을 사용하려면 다음 매개변수를 설치 명령에 추가하십시오:

```
--set serviceAccount.create=false \  
--set serviceAccount.name=<기존_서비스_계정_입력> \
```  