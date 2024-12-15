package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
)

var (
	Armory = raidengine.Armory{
		Tactics: map[string][]raidengine.Strike{
			"maturity_1": {
				AC_01,
				AC_02,
				AC_03,
				AC_04,
				BR_01,
				BR_02,
				BR_03,
				DO_01,
				DO_02,
				LE_02,
				LE_03,
				LE_04,
				QA_01,
				QA_02,
			},
			"maturity_2": {
				AC_05,
				BR_04,
				BR_05,
				BR_06,
				BR_08,
				DO_03,
				DO_04,
				DO_05,
				DO_06,
				DO_07,
				DO_11,
				DO_12,
				LE_01,
				QA_03,
				QA_04,
				QA_06,
				QA_07,
			},
			"maturity_3": {
				AC_07,
				DO_08,
				DO_09,
				DO_10,
				QA_05,
			},
		},
	}
)
