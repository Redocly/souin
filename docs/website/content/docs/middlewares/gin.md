+++
weight = 507
title = "Gin"
icon = "extension"
description = "Use Souin directly in the Gin web server"
tags = ["Beginners", "Advanced"]
+++

## Usage
Here is the example about the Souin initialization.
```go
import (
	"net/http"

	souin_gin "github.com/Redocly/souin/plugins/gin"
	"github.com/gin-gonic/gin"
)

func main(){

    // ...
	s := souin_gin.New(souin_gin.DevDefaultConfiguration)
	e.Use(s.Process)
    // ...

}
```
With that your application will be able to cache the responses if possible and returns at least the `Cache-Status` HTTP header with the different directives mentionned in the RFC specification.  
You have to pass a Souin `BaseConfiguration` structure into the `NewHTTPCache` method (you can use the `DefaultConfiguration` variable to have a built-in production ready configuration).  

Look at the configuration section to discover [all configurable keys here]({{% relref "/docs/configuration" %}}).

Other resources
---------------
You can find an example for a docker-compose stack inside the [`examples` folder on the Github repository](https://github.com/Redocly/souin/tree/master/plugins/gin/examples).
Look at the [`BaseConfiguration` structure on pkg.go.dev documentation](https://pkg.go.dev/github.com/Redocly/souin/pkg/middleware#BaseConfiguration).
