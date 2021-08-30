<template>
  <div>
    <div class="filter-container">
      <el-input
        v-model="listQuery.title"
        placeholder="标题"
        style="width: 200px"
        class="filter-item"
      />
      <el-select
        v-model="listQuery.level_id"
        placeholder="级别"
        clearable
        style="width: 90px"
        class="filter-item"
      >
        <el-option
          v-for="item in levels"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        />
      </el-select>
      <el-select
        v-model="listQuery.project_id"
        placeholder="项目名"
        clearable
        class="filter-item"
        style="width: 130px"
      >
        <el-option
          v-for="item in projectnames"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        />
      </el-select>
      <el-button
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="searchHandle"
        >搜索</el-button
      >
      <el-dropdown
        :hide-on-click="false"
        :show-timeout="100"
        trigger="click"
        style="vertical-align: top"
      >
        <el-button plain>
          状态({{ statuslength }})
          <i class="el-icon-caret-bottom el-icon--right" />
        </el-button>
        <el-dropdown-menu slot="dropdown" class="no-border">
          <el-checkbox-group
            v-model="showstatus"
            style="padding-left: 15px"
            @change="HandleChange"
          >
            <el-checkbox
              v-for="item in allStatus"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            >
            </el-checkbox>
          </el-checkbox-group>
        </el-dropdown-menu>
      </el-dropdown>
    </div>

    <show
      :list="list"
      :pageType="pageType"
      :statusMap="statusidMap"
      :projectMap="projectMap"
      :levelMap="levelMap"
    />

    <div class="pagination-container">
      <el-pagination
        :current-page="listQuery.page"
        :page-sizes="[10, 15, 20, 30]"
        :page-size="listQuery.limit"
        :total="total"
        background
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script>
import {
  bugFilter,
  searchMyBugs,
  searchAllBugs,
  searchMyTasks,
} from "@/api/search"; // 水波纹指令
import { resumeBug } from "@/api/bugs";
import { statusFilter } from "@/api/status";
import Show from "./Bug/Show";
import {
  getProjectKeyName,
  getLevels,
  getShowStatus,
  getStatus,
} from "@/api/get";
export default {
  name: "BugSearch",
  components: {
    Show,
  },
  props: {
    // 1, 2, 3, 4  对应5种类型页面
    pageType: {
      type: Number,
      default: 1,
    },
  },
  data() {
    return {
      list: [],
      listQuery: {
        limit: 15,
        page: 1,
        level_id: undefined,
        project_id: undefined,
        showstatus: [],
        title: "",
      },
      showstatus: [],
      total: 0,
      projectnames: [],
      projectMap: new Map(),
      levels: [],
      levelMap: new Map(),
      allStatus: [],
      showstatus: [],
      statusMap: new Map(),
      statusidMap: new Map(),
      statuslength: 0,
    };
  },
  activated() {
    this.getpname();
    this.getstatus();
  },
  created() {
    this.getpname();
    this.getstatus();
  },
  methods: {
    getstatus() {
      const status = getStatus();
      const level = getLevels();
      const show = getShowStatus();

      Promise.all([status, level, show]).then((values) => {
        this.allStatus = values[0].data.data;
        this.statusMap.clear();
        for (let v of this.allStatus) {
          this.statusMap.set(v.name, v.id);
          this.statusidMap.set(v.id, v.name);
        }

        this.levels = values[1].data.data;
        this.levelMap.clear();
        for (let v of this.levels) {
          this.levelMap.set(v.id, v.name);
        }
        this.showstatus = [];
        for (let v of values[2].data.data) {
          this.showstatus.push(this.statusidMap.get(v));
        }
        this.statuslength = this.showstatus.length;
        this.handleFilter();
      });
      
    },

    getpname() {
      getProjectKeyName().then((resp) => {
        this.projectnames = resp.data.data;
        this.projectMap.clear();
        for (let v of this.projectnames) {
          this.projectMap.set(v.id, v.name);
        }
      });
    },
    getSidByShowStatus() {
      let sids = [];
      for (let v of this.showstatus) {
        sids.push(this.statusMap.get(v));
      }
      return sids;
    },
    HandleChange() {
      let sids = this.getSidByShowStatus();
      this.statuslength = this.showstatus.length;
      const data = {
        showstatus: sids,
      };
      statusFilter(data).then(() => {});
    },
    handleFilter() {
      this.listQuery.showstatus = this.getSidByShowStatus();
      switch (this.pageType) {
        case 1:
          // 垃圾箱
          bugFilter(this.listQuery).then((resp) => {
            this.list = resp.data.data;
            this.total = resp.data.total;
            this.listQuery.page = resp.data.page;
          });
          break;
        case 2:
          // 我创建的bug
          searchMyBugs(this.listQuery).then((resp) => {
            this.list = resp.data.data;
            this.total = resp.data.total;
            this.listQuery.page = resp.data.page;
          });
          break;
        case 3:
          // 所有的bug
          searchAllBugs(this.listQuery).then((resp) => {
            this.list = resp.data.data;
            this.total = resp.data.total;
            this.listQuery.page = resp.data.page;
          });
          break;
        case 4:
          // 我的任务
          searchMyTasks(this.listQuery).then((resp) => {
            this.list = resp.data.data;
            this.total = resp.data.total;
            this.listQuery.page = resp.data.page;
          });
          break;
      }
    },
    searchHandle() {
      this.handleFilter();
    },
    handleSizeChange(val) {
      this.listQuery.limit = val;
      this.handleFilter();
    },
    handleCurrentChange(val) {
      this.listQuery.page = val;
      this.handleFilter();
    },
    resume(id) {
      resumeBug(id).then((_) => {
        const l = this.list.length;
        for (var i = 0; i < l; i++) {
          if (this.list[i].id === id) {
            this.list.splice(i, 1);
            this.$message.success("恢复成功");
            return;
          }
        }
      });
    },
  },
};
</script>
