<template>
  <div>
    <el-table
      ref="multipleTable"
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

      <el-table-column label="日期" width="150px" align="center">
        <template slot-scope="scope">
          <span>{{
            scope.row.createtime | parseTime("{y}-{m}-{d} {h}:{i}")
          }}</span>
        </template>
      </el-table-column>

      <el-table-column label="项目" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.pid | toProjectName(projectMap) }}</span>
        </template>
      </el-table-column>

      <el-table-column label="优先级" width="80px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.lid | toLevelName(levelMap) }}</span>
        </template>
      </el-table-column>

      <el-table-column label="重要性" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.iid | toImportantName(importantMap) }}</span>
        </template>
      </el-table-column>

      <el-table-column label="处理者" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.spusers | toUserName(userMap) }}</span>
        </template>
      </el-table-column>

      <el-table-column label="状态" align="center" width="110">
        <template slot-scope="scope">
          <span>{{ scope.row.sid | toStatusName(statusMap) }}</span>
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

      <el-table-column
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
          <el-button type="success" size="mini" @click="resume(scope.row.id)"
            >恢复</el-button
          >
          <!--<el-button type="danger" size="mini" @click="handleRemove(scope.row)">{{ $t('list.remove') }}</el-button>-->
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getImportants, getUserKeyName } from "@/api/get";

export default {
  name: "Show",
  filters: {
    toProjectName(id, projectMap) {
      return projectMap.get(id);
    },
    toLevelName(id, levelMap) {
      return levelMap.get(id);
    },
    toStatusName(id, statusMap) {
      return statusMap.get(id);
    },
    toImportantName(id, importantMap) {
      return importantMap.get(id);
    },
    toUserName(ids, userMap) {
      let names = [];
      for (let id of ids) {
        names.push(userMap.get(id));
      }
      return names.join(", ");
    }
  },
  props: {
    // 数据
    list: {
      type: Array,
      default: []
    },
    // project 映射关系
    projectMap: {
      type: Map,
      require: true
    },
    levelMap: {
      type: Map,
      require: true
    },
    statusMap: {
      type: Map,
      require: true
    }
  },
  data() {
    return {
      listLoading: false,
      importantMap: new Map(),
      userMap: new Map()
    };
  },
  created() {
    this.getImportant();
    this.getUsers();
  },
  methods: {
    getImportant() {
      getImportants().then(resp => {
        for (let v of resp.data.data) {
          this.importantMap.set(v.id, v.name);
        }
      });
    },
    getUsers() {
      getUserKeyName().then(resp => {
        for (let v of resp.data.data) {
          this.userMap.set(v.id, v.name);
        }
      });
    }
  }
};
</script>
