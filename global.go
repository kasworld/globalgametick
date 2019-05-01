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

// gametickmaker의 global instance.
// 서버가 종료 되면 흐름이 멎고 재시작하면 계속 됨.
package globalgametick

import (
	"time"

	"github.com/kasworld/gametick"
	"github.com/kasworld/gametickmaker"
)

var G_GameTick gametickmaker.GameTickMaker = gametickmaker.New(0, 1)

func SetGameTickMaker(lasttick gametick.GameTick, ntimes float64) {
	G_GameTick = gametickmaker.New(lasttick, ntimes)
}

func AccelerateBy(n float64) {
	G_GameTick = G_GameTick.GetAcceleratedBy(n)
}

func GetGameTick() gametick.GameTick {
	return G_GameTick.GetGameTick()
}

func ToUTCTime(tk gametick.GameTick) time.Time {
	return G_GameTick.ToUTCTime(tk)
}

func FromTimeToTickType(s time.Time) gametick.GameTick {
	return G_GameTick.FromTimeToTickType(s)
}
