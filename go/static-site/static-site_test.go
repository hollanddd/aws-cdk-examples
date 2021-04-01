package main

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/jsii-runtime-go"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestStaticSiteStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := NewStaticSiteStack(app, "MyStack", StaticSiteProps{
		DomainName: jsii.String("test.com"),
		StackProps: awscdk.StackProps{
			Env: Env(),
		},
	})

	// THEN
	bytes, err := json.Marshal(app.Synth(nil).GetStackArtifact(stack.ArtifactId()).Template())
	if err != nil {
		t.Error(err)
	}

	template := gjson.ParseBytes(bytes)
	resources := template.Get("Resources").String()
	assert.Equal(t, "never gets here", resources)
}
