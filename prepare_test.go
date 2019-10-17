/*
 * Copyright 2014-2019 Li Kexian
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Go module for domain whois info parse
 * https://www.likexian.com/
 */

package whoisparser

import (
	"strings"
	"testing"

	"github.com/likexian/gokit/assert"
	"github.com/likexian/gokit/xfile"
)

func TestPrepare(t *testing.T) {
	tests := []string{
		"it_git.it",
		"it_google.it",
		"ch_google.ch",
		"ch_switch.ch",
		"fr_git.fr",
		"pm_git.pm",
		"re_git.re",
		"tf_git.tf",
		"wf_git.wf",
		"yt_git.yt",
		"fr_google.fr",
		"pm_google.pm",
		"re_google.re",
		"tf_google.tf",
		"wf_google.wf",
		"yt_google.yt",
		"fr_ovh.fr",
		"ru_git.ru",
		"ru_google.ru",
		"jp_git.jp",
		"jp_google.jp",
		"edu_git.edu",
		"edu_snai.edu",
		"edu_us.edu",
		"int_esa.int",
		"int_wto.int",
		"mo_yp.mo",
		"mo_moo.mo",
		"hk_git.hk",
		"hk_google.hk",
		"hk_ibm.hk",
		"tw_git.tw",
		"tw_google.tw",
		"tw_msn.tw",
		"su_git.su",
		"su_google.su",
	}

	for _, v := range tests {
		whoisRaw, err := xfile.ReadText("./examples/" + v)
		assert.Nil(t, err)
		whoisPre, err := xfile.ReadText("./examples/" + v + ".pre")
		assert.Nil(t, err)
		result := Prepare(whoisRaw)
		assert.Equal(t, strings.TrimSpace(result), strings.TrimSpace(whoisPre))
	}
}
