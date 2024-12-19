# Snyk IaC with Broker for self-hosted Git

Snyk Broker는 내부 Git 서버를 Snyk에 연결할 수 있게 해줍니다. 이를 통해 Git 서버가 인터넷에 접속할 수 없더라도 Snyk와 연동이 가능합니다.

기본적으로 Snyk Broker는 Snyk Open Source 및 Docker 파일에 대한 정보만 호스트합니다. `.tf` 또는 `.yaml`과 같은 Infrastructure as Code 파일을 분석하려면 Broker를 설정해야 합니다.

자세한 내용은 [Snyk Broker - Infrastructure as Code detection](../../../enterprise-setup/snyk-broker/snyk-broker-infrastructure-as-code-detection/)을 참조하세요. Snyk Broker에 대한 자세한 정보는 [Broker](../../../enterprise-setup/snyk-broker/) 문서를 확인하세요.