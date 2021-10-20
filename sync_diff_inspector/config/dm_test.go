// Copyright 2021 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/pingcap/check"
)

type getDMTaskCfgSuite struct{}

var _ = Suite(&getDMTaskCfgSuite{})

func testHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `{"result":true,"cfgs":["is-sharding = true\\nshard-mode = \\"pessimistic\\"\\nonline-ddl-scheme = \\"\\"\\ncase-sensitive = false\\nname = \\"test\\"\\nmode = \\"all\\"\\nsource-id = \\"mysql-replica-01\\"\\nserver-id = 0\\nflavor = \\"\\"\\nmeta-schema = \\"dm_meta\\"\\nheartbeat-update-interval = 1\\nheartbeat-report-interval = 10\\nenable-heartbeat = false\\ntimezone = \\"Asia/Shanghai\\"\\nrelay-dir = \\"\\"\\nuse-relay = false\\nfilter-rules = []\\nmydumper-path = \\"./bin/mydumper\\"\\nthreads = 4\\nchunk-filesize = \\"64\\"\\nstatement-size = 0\\nrows = 0\\nwhere = \\"\\"\\nskip-tz-utc = true\\nextra-args = \\"\\"\\npool-size = 16\\ndir = \\"./dumped_data.test\\"\\nmeta-file = \\"\\"\\nworker-count = 16\\nbatch = 100\\nqueue-size = 1024\\ncheckpoint-flush-interval = 30\\nmax-retry = 0\\nauto-fix-gtid = false\\nenable-gtid = false\\ndisable-detect = false\\nsafe-mode = false\\nenable-ansi-quotes = false\\nlog-level = \\"\\"\\nlog-file = \\"\\"\\nlog-format = \\"\\"\\nlog-rotate = \\"\\"\\npprof-addr = \\"\\"\\nstatus-addr = \\"\\"\\nclean-dump-file = true\\n\\n[from]\\n  host = \\"127.0.0.1\\"\\n  port = 3306\\n  user = \\"root\\"\\n  password = \\"/Q7B9DizNLLTTfiZHv9WoEAKamfpIUs=\\"\\n  max-allowed-packet = 67108864\\n\\n[to]\\n  host = \\"127.0.0.1\\"\\n  port = 4000\\n  user = \\"root\\"\\n  password = \\"\\"\\n  max-allowed-packet = 67108864\\n\\n[[route-rules]]\\n  schema-pattern = \\"sharding*\\"\\n  table-pattern = \\"t*\\"\\n  target-schema = \\"db_target\\"\\n  target-table = \\"t_target\\"\\n\\n[[route-rules]]\\n  schema-pattern = \\"sharding*\\"\\n  table-pattern = \\"\\"\\n  target-schema = \\"db_target\\"\\n  target-table = \\"\\"\\n\\n[[mapping-rule]]\\n  schema-pattern = \\"sharding*\\"\\n  table-pattern = \\"t*\\"\\n  source-column = \\"id\\"\\n  target-column = \\"id\\"\\n  expression = \\"partition id\\"\\n  arguments = [\\"1\\", \\"sharding\\", \\"t\\"]\\n  create-table-query = \\"\\"\\n\\n[block-allow-list]\\n  do-dbs = [\\"~^sharding[\\\\\\\\d]+\\"]\\n\\n  [[block-allow-list.do-tables]]\\n    db-name = \\"~^sharding[\\\\\\\\d]+\\"\\n    tbl-name = \\"~^t[\\\\\\\\d]+\\"\\n","is-sharding = true\\nshard-mode = \\"pessimistic\\"\\nonline-ddl-scheme = \\"\\"\\ncase-sensitive = false\\nname = \\"test\\"\\nmode = \\"all\\"\\nsource-id = \\"mysql-replica-02\\"\\nserver-id = 0\\nflavor = \\"\\"\\nmeta-schema = \\"dm_meta\\"\\nheartbeat-update-interval = 1\\nheartbeat-report-interval = 10\\nenable-heartbeat = false\\ntimezone = \\"Asia/Shanghai\\"\\nrelay-dir = \\"\\"\\nuse-relay = false\\nfilter-rules = []\\nmydumper-path = \\"./bin/mydumper\\"\\nthreads = 4\\nchunk-filesize = \\"64\\"\\nstatement-size = 0\\nrows = 0\\nwhere = \\"\\"\\nskip-tz-utc = true\\nextra-args = \\"\\"\\npool-size = 16\\ndir = \\"./dumped_data.test\\"\\nmeta-file = \\"\\"\\nworker-count = 16\\nbatch = 100\\nqueue-size = 1024\\ncheckpoint-flush-interval = 30\\nmax-retry = 0\\nauto-fix-gtid = false\\nenable-gtid = false\\ndisable-detect = false\\nsafe-mode = false\\nenable-ansi-quotes = false\\nlog-level = \\"\\"\\nlog-file = \\"\\"\\nlog-format = \\"\\"\\nlog-rotate = \\"\\"\\npprof-addr = \\"\\"\\nstatus-addr = \\"\\"\\nclean-dump-file = true\\n\\n[from]\\n  host = \\"127.0.0.1\\"\\n  port = 3307\\n  user = \\"root\\"\\n  password = \\"/Q7B9DizNLLTTfiZHv9WoEAKamfpIUs=\\"\\n  max-allowed-packet = 67108864\\n\\n[to]\\n  host = \\"127.0.0.1\\"\\n  port = 4000\\n  user = \\"root\\"\\n  password = \\"\\"\\n  max-allowed-packet = 67108864\\n\\n[[route-rules]]\\n  schema-pattern = \\"sharding*\\"\\n  table-pattern = \\"t*\\"\\n  target-schema = \\"db_target\\"\\n  target-table = \\"t_target\\"\\n\\n[[route-rules]]\\n  schema-pattern = \\"sharding*\\"\\n  table-pattern = \\"\\"\\n  target-schema = \\"db_target\\"\\n  target-table = \\"\\"\\n\\n[[mapping-rule]]\\n  schema-pattern = \\"sharding*\\"\\n  table-pattern = \\"t*\\"\\n  source-column = \\"id\\"\\n  target-column = \\"id\\"\\n  expression = \\"partition id\\"\\n  arguments = [\\"2\\", \\"sharding\\", \\"t\\"]\\n  create-table-query = \\"\\"\\n\\n[block-allow-list]\\n  do-dbs = [\\"~^sharding[\\\\\\\\d]+\\"]\\n\\n  [[block-allow-list.do-tables]]\\n    db-name = \\"~^sharding[\\\\\\\\d]+\\"\\n    tbl-name = \\"~^t[\\\\\\\\d]+\\"\\n"]}`)
}

func equal(a *DataSource, b *DataSource) bool {
	return a.Host == b.Host && a.Port == b.Port && a.Password == b.Password && a.User == b.User
}

func (s *getDMTaskCfgSuite) TestGetDMTaskCfg(c *C) {
	mockServer := httptest.NewServer(http.HandlerFunc(testHandler))
	defer mockServer.Close()

	dmTaskCfg, err := getDMTaskCfg(mockServer.URL, "test")
	c.Assert(err, IsNil)
	c.Assert(dmTaskCfg, HasLen, 2)
	c.Assert(dmTaskCfg[0].SourceID, Equals, "mysql-replica-01")
	c.Assert(dmTaskCfg[1].SourceID, Equals, "mysql-replica-02")

	cfg := NewConfig()
	cfg.DMAddr = mockServer.URL
	cfg.DMTask = "test"
	err = cfg.adjustConfigByDMSubTasks()
	c.Assert(err, IsNil)

	// after adjust config, will generate source tables for target table
	c.Assert(cfg.DataSources, HasLen, 3)
	c.Assert(equal(cfg.DataSources["target"], &DataSource{
		Host:     dmTaskCfg[0].To.Host,
		Port:     dmTaskCfg[0].To.Port,
		Password: dmTaskCfg[0].To.Password,
		User:     dmTaskCfg[0].To.User,
	}), IsTrue)

	c.Assert(equal(cfg.DataSources["mysql-replica-01"], &DataSource{
		Host:     dmTaskCfg[0].From.Host,
		Port:     dmTaskCfg[0].From.Port,
		Password: dmTaskCfg[0].From.Password,
		User:     dmTaskCfg[0].From.User,
	}), IsTrue)

	c.Assert(equal(cfg.DataSources["mysql-replica-02"], &DataSource{
		Host:     dmTaskCfg[1].From.Host,
		Port:     dmTaskCfg[1].From.Port,
		Password: dmTaskCfg[1].From.Password,
		User:     dmTaskCfg[1].From.User,
	}), IsTrue)

}
