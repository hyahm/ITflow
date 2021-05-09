<template>
  <div style="padding-left: 20px">
    <p class="warn-content">
      某些选项的默认值
    </p>
    <div style="margin-top: 20px">
      bug创建时的状态:
      <el-select
        v-model="form.created"
        placeholder="Select"
        @change="handleCreated"
      >
        <el-option
          v-for="(status, index) in statuslist"
          :key="index"
          :label="status"
          :value="status"
        />
      </el-select>
    </div>
    <div style="margin-top: 20px">
      bug完成时的状态:
      <el-select
        v-model="form.completed"
        placeholder="Select"
        @change="handleCompleted"
      >
        <el-option
          v-for="(status, index) in statuslist"
          :key="index"
          :label="status"
          :value="status"
        />
      </el-select>
    </div>
    <!-- <div style="margin: 10px 0 10px 10px"> 严重级别:
      <el-select v-model="form.defaultlevel" placeholder="Select" @change="handleChangeLevel">
        <el-option
          v-for="(important, index) in levels"
          :key="index"
          :label="important"
          :value="important"
        />
      </el-select>
    </div> -->
    <el-button style="margin-top: 20px" type="primary" plain @click="handleSave"
      >保存</el-button
    >
  </div>
</template>

<script>
import { status, save, important, level } from "@/api/defaultvalue";
import { getStatus, getImportants, getLevels } from "@/api/get";
export default {
  name: "DefaultValue",
  data() {
    return {
      form: {
        created: "",
        completed: ""
      },

      statuslist: [],
      importants: [],
      levels: []
    };
  },
  created() {
    this.getdefaultstatus();
    this.getstatuslist();
  },
  methods: {
    getlevels() {
      getLevels().then(resp => {
        this.levels = resp.data.levels;
      });
    },
    getdefaultlevel() {
      level().then(resp => {
        this.form.defaultlevel = resp.data.defaultlevel;
      });
    },
    getimportantlist() {
      getImportants().then(resp => {
        this.importants = resp.data.importants;
      });
    },
    getdefaultimportant() {
      important().then(resp => {
        this.form.defaultimportant = resp.data.defaultimportant;
      });
    },
    getdefaultstatus() {
      status().then(resp => {
        this.form.created = resp.data.created;
        this.form.completed = resp.data.completed;
      });
    },
    getstatuslist() {
      getStatus().then(resp => {
        this.statuslist = resp.data.statuslist;
      });
    },
    handleCompleted(e) {
      this.form.completed = e;
    },
    handleCreated(e) {
      this.form.created = e;
    },
    handleChangeLevel(e) {
      this.form.defaultlevel = e;
    },
    handleSave() {
      save(this.form).then(_ => {
        this.$message.success("保存成功");
      });
    }
  }
};
</script>

<style scoped></style>
