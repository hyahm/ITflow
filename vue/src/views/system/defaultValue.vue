<template>
  <div style="padding-left: 20px">
    <p class="warn-content">某些选项的默认值</p>
    <div style="margin-top: 20px">
      bug创建时的状态:
      <el-select v-model="form.created" placeholder="Select">
        <el-option
          v-for="status in statuslist"
          :key="status.id"
          :label="status.name"
          :value="status.id"
        />
      </el-select>
    </div>
    <div style="margin-top: 20px">
      bug完成时的状态:
      <el-select v-model="form.completed" placeholder="Select">
        <el-option
          v-for="status in statuslist"
          :key="status.id"
          :label="status.name"
          :value="status.id"
        />
      </el-select>
    </div>
    <div style="margin-top: 20px">
      bug转交时的状态:
      <el-select v-model="form.pass" placeholder="Select">
        <el-option
          v-for="status in statuslist"
          :key="status.id"
          :label="status.name"
          :value="status.id"
        />
      </el-select>
    </div>
    <div style="margin-top: 20px">
      bug领取后的状态:
      <el-select v-model="form.receive" placeholder="Select">
        <el-option
          v-for="status in statuslist"
          :key="status.id"
          :label="status.name"
          :value="status.id"
        />
      </el-select>
    </div>
    <el-button style="margin-top: 20px" type="primary" plain @click="handleSave"
      >保存</el-button
    >
  </div>
</template>

<script>
import { defaultValue, save, important, level } from "@/api/defaultvalue";
import { getStatus, getImportants, getLevels } from "@/api/get";
export default {
  name: "DefaultValue",
  data() {
    return {
      form: {
        created: undefined,
        completed: undefined,
        pass: undefined,
        receive: undefined,
      },

      statuslist: [],
      importants: [],
      levels: [],
    };
  },
  created() {
    this.getdefaultstatus();
    this.getstatuslist();
  },
  methods: {
    // getlevels() {
    //   getLevels().then((resp) => {
    //     this.levels = resp.data.levels;
    //   });
    // },
    // getdefaultlevel() {
    //   level().then((resp) => {
    //     this.form.defaultlevel = resp.data.defaultlevel;
    //   });
    // },
    // getimportantlist() {
    //   getImportants().then((resp) => {
    //     this.importants = resp.data.importants;
    //   });
    // },
    // getdefaultimportant() {
    //   important().then((resp) => {
    //     this.form.defaultimportant = resp.data.defaultimportant;
    //   });
    // },
    getdefaultValue() {
      defaultValue().then((resp) => {
        this.form.created = resp.data.data.created;
        this.form.completed = resp.data.data.completed;
        this.form.pass = resp.data.data.pass;
        this.form.receive = resp.data.data.receive;
      });
    },
    getstatuslist() {
      getStatus().then((resp) => {
        this.statuslist = resp.data.data;
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
      console.log(this.form);
      save(this.form).then((_) => {
        this.$message.success("保存成功");
      });
    },
  },
};
</script>

<style scoped></style>
