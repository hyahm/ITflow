<template>
  <div class="app-container">
    <p class="warn-content">
      如果没看到需要的bug, 请勾选需要显示的状态 ,永久保存，多页面生效
    </p>

    <bug-search :pageType="2" />
    <!-- <div class="filter-container">
      <el-input
        v-model="listQuery.title"
        placeholder="标题"
        style="width: 200px;"
        class="filter-item"
        @keyup.enter.native="handleFilter"
      />
      <el-select
        v-model="listQuery.level"
        placeholder="级别"
        clearable
        style="width: 90px"
        class="filter-item"
      >
        <el-option
          v-for="(item, index) in levels"
          :key="index"
          :label="item"
          :value="item"
        />
      </el-select>
      <el-select
        v-model="listQuery.project"
        placeholder="项目"
        clearable
        class="filter-item"
        style="width: 130px"
      >
        <el-option
          v-for="(item, index) in projectnames"
          :key="index"
          :label="item"
          :value="item"
        />
      </el-select>

      <el-button
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="handleFilter"
        >搜索</el-button
      >
      <el-dropdown
        :hide-on-click="false"
        :show-timeout="100"
        trigger="click"
        style="vertical-align: top;"
      >
        <el-button plain>
          状态({{ statuslength }})
          <i class="el-icon-caret-bottom el-icon--right" />
        </el-button>
        <el-dropdown-menu slot="dropdown" class="no-border">
          <el-checkbox-group
            v-model="listQuery.showstatus"
            style="padding-left: 15px;"
            @change="HandleChange"
          >
            <el-checkbox
              v-for="(item, index) in platformsOptions"
              :key="index"
              :label="item"
            >
              {{ item }}
            </el-checkbox>
          </el-checkbox-group>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
    <el-table
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%"
    >
      <el-table-column label="id" align="center" width="50">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column label="时间" width="150px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.date | parseTime("{y}-{m}-{d} {h}:{i}") }}</span>
        </template>
      </el-table-column>

      <el-table-column label="项目" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.projectname }}</span>
        </template>
      </el-table-column>

      <el-table-column label="级别" width="80px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.level }}</span>
        </template>
      </el-table-column>

      <el-table-column label="重要性" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.important }}</span>
        </template>
      </el-table-column>

      <el-table-column label="状态" align="center" width="110">
        <template slot-scope="scope">
          <el-select
            v-model="scope.row.status"
            class="filter-item"
            @change="changestatus(scope.row)"
          >
            <el-option
              v-for="(item, index) in statuslist"
              :key="index"
              :label="item"
              :value="item"
            />
          </el-select>
        </template>
      </el-table-column>

      <el-table-column label="标题" min-width="300px" align="center">
        <template slot-scope="scope">
          <router-link
            :to="'/showbug/' + scope.row.id"
            class="link-type"
            align="center"
          >
            <span>{{ scope.row.title }}</span>
          </router-link>
        </template>
      </el-table-column>
      <el-table-column label="任务者" align="center" width="300">
        <template slot-scope="scope">
          <span>{{ scope.row.handle }}</span>
          <span v-if="scope.row.handle" class="link-type" @click='handleFetchPv(scope.row.pageviews)'>{{scope.row.pageviews}}</span>-->
    <!--<span v-else>0</span>-->
    <!-- </template>
      </el-table-column> -->

    <!-- <el-table-column
        label="操作"
        align="center"
        width="230"
        class-name="small-padding fixed-width"
      >
        <template slot-scope="scope">
          <el-button type="primary" size="mini"
            ><router-link :to="'/bug/edit/' + scope.row.id"
              >编辑</router-link
            ></el-button
          >
          <el-button type="success" size="mini" @click="handleClose(scope.row)"
            >关闭</el-button
          >
        </template>
      </el-table-column>
    </el-table> -->

    <!-- <div class="pagination-container">
      <el-pagination
        :current-page="listQuery.page"
        :page-sizes="[10]"
        :page-size="listQuery.limit"
        :total="total"
        background
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div> -->
  </div>
</template>

<script>
import { closeBug, changeStatus } from "@/api/bugs";

import BugSearch from "@/views/bug/components/Search";
import {
  getMyProject,
  getStatus,
  getShowStatus,
  getPermStatus,
  getLevels
} from "@/api/get";

export default {
  name: "ArticleList",
  components: {
    BugSearch
  },
  // filters: {
  //   statusFilter(status) {
  //     const statusMap = {
  //       published: "success",
  //       draft: "info",
  //       deleted: "danger"
  //     };
  //     return statusMap[status];
  //   },
  //   typeFilter(type) {
  //     return calendarTypeKeyValue[type];
  //   }
  // },
  data() {
    return {
      // list: null,
      // total: 0,
      // listLoading: true,
      // listQuery: {
      //   page: 1,
      //   limit: 10,
      //   level: "",
      //   project: "",
      //   title: "",
      //   showstatus: []
      // },
      // projectnames: [],
      // platformsOptions: [],
      // statuslist: [],
      // levels: [],
      // statuslength: 0
    };
  }
  // mounted() {
  //   this.getstatus();
  //   this.getpname();
  //   this.getlevels();
  //   this.handleFilter();
  //   this.getmystatus();
  // },
  // activated() {
  //   this.getmystatus();
  //   this.getpname();
  //   this.getstatus();
  //   this.getlevels();
  // },

  // methods: {
  //   getlevels() {
  //     getLevels().then(resp => {
  //       this.levels = resp.data.levels;
  //     });
  //   },
  //   HandleChange() {
  //     const data = {
  //       checkstatus: this.listQuery.showstatus
  //     };

  //     statusFilter(data).then(_ => {
  //       this.statuslength = this.listQuery.showstatus.length;
  //       this.handleFilter();
  //     });
  //   },
  //   getstatus() {
  //     getStatus().then(resp => {
  //       this.platformsOptions = resp.data.statuslist;
  //     });
  //     // 可以修改的权限
  //     getPermStatus().then(resp => {
  //       this.statuslist = resp.data.statuslist;
  //     });
  //   },
  //   //
  //   getmystatus() {
  //     // 需要显示的状态
  //     getShowStatus().then(resp => {
  //       this.listQuery.showstatus = resp.data.checkstatus;
  //       this.statuslength = this.listQuery.showstatus.length;
  //     });
  //   },
  //   getpname() {
  //     getMyProject().then(resp => {
  //       this.projectnames = resp.data.name;
  //     });
  //   },
  //   handleClose(row) {
  //     this.$confirm("此操作将关闭bug, 是否继续?", "提示", {
  //       confirmButtonText: "确定",
  //       cancelButtonText: "取消",
  //       type: "warning"
  //     })
  //       .then(() => {
  //         closeBug(row.id).then(_ => {
  //           this.list = this.list.filter(items => {
  //             return items.id !== row.id;
  //           });
  //           this.$message({
  //             message: "已关闭",
  //             type: "success"
  //           });
  //         });
  //       })
  //       .catch(() => {
  //         this.$message({
  //           type: "info",
  //           message: "已取消删除"
  //         });
  //       });
  //   },
  //   changestatus(row) {
  //     const param = {
  //       id: row.id,
  //       status: row.status
  //     };
  //     changeStatus(param).then(response => {
  //       this.$notify({
  //         title: "成功",
  //         message: "修改成功",
  //         type: "success"
  //       });
  //     });
  //   },
  //   handleFilter() {
  //     // 获取过滤后的bug
  //     this.listLoading = true;
  //     searchMyBugs(this.listQuery).then(resp => {
  //       this.list = resp.data.articlelist;
  //       this.total = resp.data.total;
  //       this.listQuery.page = resp.data.page;
  //     });
  //     this.listLoading = false;
  //   },
  //   handleSizeChange(val) {
  //     this.listQuery.limit = val;
  //     this.handleFilter();
  //   },
  //   handleCurrentChange(val) {
  //     this.listQuery.page = val;
  //     this.handleFilter();
  //   }
  // }
};
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
