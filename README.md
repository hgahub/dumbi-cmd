# Dumbi CMD
Dumbi CMD is a tool for building, changing, and versioning data infrastructure

## SonarQube

```shell
# Download 
docker pull sonarqube
docker pull sonarsource/sonar-scanner-cli

# Start the server by running:
docker run -d --name sonarqube -e SONAR_ES_BOOTSTRAP_CHECKS_DISABLE=true -p 9000:9000 sonarqube:latest

# Settings
export SONARQUBE_URL=http://[IP_ADDRESS]:9000
```

### Release
```shell
goreleaser release --snapshot --rm-dist
```