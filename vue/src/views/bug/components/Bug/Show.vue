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
          <span>{{ scope.row.createtime | parseTime("{y}-{m}-{d} {h}:{i}") }}</span>
        </template>
      </el-table-column>

      <el-table-column label="完成时间" width="150px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.deadline | parseTime("{y}-{m}-{d} {h}:{i}") }}</span>
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
          <el-tag>{{ scope.row.sid | toStatusName(statusMap) }}</el-tag>
        </template>
      </el-table-column>

      <el-table-column label="标题" min-width="300px" align="center">
        <template slot-scope="scope">
          <router-link
            target="_blank"
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
        v-if="pageType != 3"
        class-name="small-padding fixed-width"
      >
        <template slot-scope="scope">
          <el-button type="primary" size="mini" v-if="pageType == 2"
            ><router-link :to="'/bug/edit/' + scope.row.id">编辑</router-link></el-button
          >

          <el-button
            type="success"
            size="mini"
            v-if="pageType == 3"
            @click="resume(scope.row.id)"
            >恢复</el-button
          >
          <el-button
            type="success"
            size="mini"
            v-if="pageType == 4"
            @click="Receive(scope.row)"
            >领取</el-button
          >

          <el-button
            type="danger"
            v-if="pageType == 2"
            size="mini"
            @click="handleRemove(scope.row.id)"
            >删除</el-button
          >
          <el-button
            type="primary"
            v-if="pageType == 4"
            size="mini"
            @click="handlePass(scope.row)"
            >转交</el-button
          >

           <el-button
            type="primary"
            v-if="pageType == 4 && scope.row.uid == $store.getters.uid"
            size="mini"
            @click="handleComplete(scope.row)"
            >完成</el-button
          >
        </template>
      </el-table-column>
    </el-table>

    <el-dialog
      v-if="pageType == 4"
      :close-on-click-modal="false"
      title="任务完成转交"
      :visible.sync="dialogFormVisible"
    >
      <el-form
        ref="dataForm"
        :model="temp"
        label-position="left"
        label-width="70px"
        style="width: 400px; margin-left: 50px"
      >
        <el-form-item style="margin-bottom: 40px" label="任务给：">
          <el-select
            v-model="temp.spusers"
            filterable
            multiple
            allow-create
            default-first-option
            placeholder="请选择指定的用户"
          >
            <el-option
              v-for="item in userMap"
              :key="item[0]"
              :label="item[1]"
              :value="item[0]"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="说明:">
          <el-input
            v-model="temp.remark"
            :autosize="{ minRows: 2, maxRows: 4 }"
            type="textarea"
            placeholder="Please input"
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="updateData">确认</el-button>
      </div>
    </el-dialog>

    <el-dialog
      v-if="pageType == 4"
      :close-on-click-modal="false"
      title="领取任务"
      :visible.sync="openReceive"
    >
      <el-form>
        <el-form-item label="完成时间:">
          <el-date-picker
            v-model="receive.deadline"
            type="datetime"
            value-format="timestamp"
            placeholder="选择日期时间"
          >
          </el-date-picker>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="openReceive = false">取消</el-button>
        <el-button type="primary" @click="receiveHandle">确认</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getImportants, getUserKeyName, getUserKeyNameByProject } from "@/api/get";
import { defaultValue } from "@/api/defaultvalue";
import { passBug, delBug, receiveBug,completeBug } from "@/api/bugs";
// pageType:    1： 垃圾箱   2： 我创建的bug    3： 所有bug     4: 我的任务
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
    },
  },
  props: {
    // 数据
    list: {
      type: Array,
      default: [],
    },
    // project 映射关系
    projectMap: {
      type: Map,
      require: true,
    },
    levelMap: {
      type: Map,
      require: true,
    },
    statusMap: {
      type: Map,
      require: true,
    },
    pageType: {
      type: Number,
      require: true,
    },
  },
  data() {
    return {
      dialogFormVisible: false,
      listLoading: false,
      importantMap: new Map(),
      userMap: new Map(),
      default: {
        created: 0,
        pass: 0,
      },
      temp: {
        bid: undefined, // bugid
        remark: "",
        spusers: [],
      },
      receive: {
        id: 0,
        deadline: 0,
      },
      openReceive: false,
    };
  },
  created() {
    this.getImportant();
    this.getUsers();
  },
  methods: {
    handleComplete(row) {
      completeBug({id: row.id}).then(()=>{
        this.$router.go(0)
      })
    },
    handleRemove(id) {
      delBug(id).then((resp) => {
        this.list = this.list.filter((m) => m.id != id);
        
      });
    },
    updateData() {
      if (this.temp.spusers.length === 0) {
        this.$message.error("至少选择一个处理人");
        return;
      }
      passBug(this.temp).then((resp) => {
        this.$message({
          message: "操作成功",
          type: "success",
        });
        setTimeout(()=>{
          this.$router.go(0)
        }, 2000)
      });
      this.dialogFormVisible = false;
    },
    getImportant() {
      getImportants().then((resp) => {
        for (let v of resp.data.data) {
          this.importantMap.set(v.id, v.name);
        }
      });
      defaultValue().then((resp) => {
        this.default = resp.data.data;
      });
    },
    getUsers() {
      getUserKeyName().then((resp) => {
        for (let v of resp.data.data) {
          this.userMap.set(v.id, v.name);
        }
      });
    },
    handlePass(row) {
      this.temp.bid = parseInt(row.id); // copy obj
      this.temp.spusers = [];
      this.users = [];
      getUserKeyNameByProject({project_id: row.pid}).then((resp) => {
        this.users = resp.data.data;
        for (var i = 0; i < this.users.length; i++) {
          for (var j = 0; j < row.spusers.length; j++) {
            if (this.temp.spusers[j] === this.users[i]) {
              this.temp.spusers.push(this.users[i]);
            }
          }
        }
      });
      this.dialogFormVisible = true;
    },
    Receive(row) {
      if (row.sid != this.default.pass && row.sid != this.default.created) {
        this.$message({
          message: "此状态无法领取",
          type: "error",
        });
        return;
      }
      this.openReceive = true;
      this.receive.id = row.id;
      this.receive.deadline = new Date().getTime();
    },
    receiveHandle(id) {
      // 确认领取任务
      this.receive.deadline = parseInt(this.receive.deadline / 1000);
      receiveBug(this.receive).then((resp) => {
        
 
        this.$message({
          message: "领取成功",
          type: "success",
        });
        setTimeout(()=>{
          this.$router.go(0)
        },2000)
      });
      this.openReceive = false;
    },
  },
};
</script>
