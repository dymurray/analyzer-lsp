package java

import (
	"reflect"
	"strings"
	"testing"
)

func Test_parseUnresolvedSources(t *testing.T) {
	tests := []struct {
		name      string
		mvnOutput string
		wantErr   bool
		wantList  []javaArtifact
	}{
		{
			name: "valid sources output",
			mvnOutput: `
[INFO] --- dependency:3.5.0:sources (default-cli) @ spring-petclinic ---
[INFO] The following files have been resolved:
[INFO]    org.apache.tomcat:tomcat-servlet-api:jar:sources:9.0.46 -- module tomcat.servlet.api (auto)
[INFO]    com.fasterxml.jackson.core:jackson-core:jar:sources:2.12.3
[INFO]    com.fasterxml.jackson.core:jackson-databind:jar:sources:2.12.3
[INFO] The following files have NOT been resolved:
[INFO]    org.apache.tomcat:tomcat-servlet-api:jar:9.0.46:provided -- module java.servlet
[INFO]    com.fasterxml.jackson.core:jackson-core:jar:2.12.3:compile -- module com.fasterxml.jackson.core
[INFO]    com.fasterxml.jackson.core:jackson-databind:jar:2.12.3:compile -- module com.fasterxml.jackson.databind
[INFO]    io.konveyor.demo:config-utils:jar:1.0.0:compile -- module config.utils (auto)
[INFO] --- maven-dependency-plugin:3.5.0:sources (default-cli) @ spring-petclinic ---
[INFO] -----------------------------------------------------------------------------
[INFO] The following files have NOT been resolved:
[INFO]    org.springframework.boot:spring-boot-actuator:jar:sources:3.1.0:compile
`,
			wantErr: false,
			wantList: []javaArtifact{
				{
					packaging:  JavaArchive,
					GroupId:    "io.konveyor.demo",
					ArtifactId: "config-utils",
					Version:    "1.0.0",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputReader := strings.NewReader(tt.mvnOutput)
			gotList, gotErr := parseUnresolvedSources(outputReader)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("parseUnresolvedSources() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
			if !reflect.DeepEqual(gotList, tt.wantList) {
				t.Errorf("parseUnresolvedSources() gotList = %v, wantList %v", gotList, tt.wantList)
			}
		})
	}
}
