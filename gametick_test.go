// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package global_gametick

import (
	"testing"
	"time"

	"github.com/kasworld/gametick"
	"github.com/kasworld/gametickmaker"
)

func TestFrame(t *testing.T) {
	t.Logf("%v", GetGameTick())
	t.Logf("%v", GetGameTick().ToIntString())
	v := GetGameTick().ToIntString()
	gtm, err := gametick.FromIntString(v)
	t.Logf("%v, %v %v", v, gtm, err)
	t.Logf("%x", GetGameTick().ToInt64NanoSec())
}

func TestFrame2(t *testing.T) {
	t.Logf("%v", G_GameTick)
	SetGameTickMaker(0, 1)
	t.Logf("%v", G_GameTick)

	gtm := gametickmaker.New(3600*gametick.Nano, 2)
	t.Logf("%v", gtm)
	tt := gtm.GetGameTick()
	t.Logf("%v", tt)
	tt2 := gtm.ToUTCTime(tt)
	t.Logf("%v", tt2)
	t.Logf("%v", time.Now().UTC())
}
