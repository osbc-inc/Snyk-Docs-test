# Snyk 브로커 클라이언트 업그레이드

Snyk는 정기적으로 브로커 클라이언트를 업데이트하여 새로운 기능, 버그 수정 및 기타 기능을 제공합니다. 릴리스 노트와 함께 모든 버전의 전체 목록은 [Snyk Broker GitHub](https://github.com/snyk/broker/releases)에서 확인할 수 있습니다. Snyk는 해당 페이지의 [RSS 피드를 구독](https://github.com/snyk/broker/releases.atom)하여 새로운 버전이 릴리스될 때마다 정보를 받도록 권장합니다.

브로커를 업그레이드하면 Snyk가 올바르게 기능하도록 필요한 새로운 규칙이 추가될 수 있습니다. 따라서, API 허용 목록을 다시 초기화해야 합니다. [사용자 정의 허용 목록을 설정](https://docs.snyk.io/snyk-admin/snyk-broker/how-to-install-and-configure-your-snyk-broker-client/advanced-configuration-for-snyk-broker-docker-installation#custom-approved-listing-filter)하여 1MB 이상의 파일을 지원하도록 한 경우와 같이 허용 목록에 규칙을 추가하거나 제거한 경우, 새로운 허용 목록에 이러한 변경 사항을 다시 적용해야 합니다. 이 내용은 Helm이나 Docker로 설치하는 경우에도 해당됩니다.