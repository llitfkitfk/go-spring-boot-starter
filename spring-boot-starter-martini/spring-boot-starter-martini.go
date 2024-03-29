/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package SpringBootStarterMartini

import (
	"github.com/didi/go-spring/spring-core"
	"github.com/didi/go-spring/spring-martini"
	"github.com/go-spring/go-spring-boot/spring-boot"
	"github.com/go-spring/go-spring-boot-starter/spring-boot-web"
)

func init() {
	SpringBoot.RegisterModule(func(ctx SpringCore.SpringContext) {
		container := SpringMartini.NewMartiniContainer()
		ctx.RegisterBean(container)
		ctx.RegisterBean(SpringBootWeb.Wrapper(container))
	})
}
