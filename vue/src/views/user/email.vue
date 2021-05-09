<template>
  <div>
    <el-form class="form-container">
      <div style="height: 20px" />

      <el-form-item label="邮箱:">
        <el-input
          v-model="email"
          :maxlength="100"
          placeholder="请输入标题"
          style="width: 60%;"
        />
      </el-form-item>

      <div>
        <el-button
          type="success"
          style="margin-left: 20px"
          plain
          @click="changeEmail"
          >修改</el-button
        >
      </div>
    </el-form>
  </div>
</template>

<script>
import { myEmail, setEmail } from "@/api/user";
export default {
  name: "Changepwd",
  data() {
    return {
      email: ""
    };
  },
  mounted() {
    this.getEmailHandle();
  },
  methods: {
    getEmailHandle() {
      myEmail().then(response => {
        this.email = response.data.email;
      });
    },
    changeEmail() {
      if (this.newpassword !== this.repassword) {
        this.$message({
          message: "新密码不一致",
          type: "error"
        });
        return;
      }
      setEmail(this.email).then(_ => {
        this.$message({
          message: "修改密码成功",
          type: "success"
        });
        this.clean();
      });
    }
  }
};
</script>

<style type="text/css">
.el-form-item {
  margin-left: 20px;
}
</style>
