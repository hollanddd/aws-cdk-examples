package main

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestStaticSiteStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := NewStaticSiteStack(app, "MyStack", StaticSiteProps{
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
	resources := template.Get("Resources.SiteBucket397A1860.Properties.WebsiteConfiguration.IndexDocument").String()
	assert.Equal(t, "index.html", resources)
}
