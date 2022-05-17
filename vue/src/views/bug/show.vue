<template>
  <div>
    <div></div>
    <div class="components-container">
      <div>
        <h1 style="text-align: center">{{ bug.title }}</h1>
      </div>

      <el-card class="box-card" style="background-color: #8cbda4">
        <el-row>
          <el-col :span="5">
            <span>项目名: {{ bug.projectname }}</span>
          </el-col>
          <el-col v-if="bug.typ == 1" :span="5">
            <span>版本：{{ bug.version }}</span>
          </el-col>
          <el-col v-if="bug.typ == 1" :span="4">
            <span>级别：{{ bug.level }}</span>
          </el-col>
          <el-col v-if="bug.typ == 1" :span="4">
            <span>重要性：{{ bug.important }}</span>
          </el-col>
          <el-col v-if="bug.typ == 1" :span="4">
            <span>环境：{{ bug.envname }}</span>
          </el-col>
        </el-row>
      </el-card>
      <el-card class="box-card" v-if="bug.typ == 1" style="background-color: #8cbda4">
        <span style="padding-top: 20px">URL：{{ bug.url }}</span>
      </el-card>
      <div id="main">
        <mavon-editor
          style="width: 100%; min-height: 10px"
          :value="bug.content"
          :box-shadow="false"
          :subfield="false"
          :toolbars-flag="false"
          default-open="preview"
          :editable="false"
          :scroll-style="true"
        />
      </div>
      <div v-for="(cc, index) in bug.comments" :key="index" style="margin-bottom: 5px">
        <el-card class="box-card">
          <p>
            {{ cc.date | parseTime("{y}-{m}-{d} {h}:{i}") }}, 处理人:{{ cc.user }}，
            事件：{{ cc.info }}
          </p>
          <!-- <span>转交原因：{{ cc.info }}</span> -->
        </el-card>
      </div>
    </div>

    <!-- <el-dialog
      :close-on-click-modal="false"
      :title="textMap[dialogStatus]"
      :visible.sync="dialogFormVisible"
    >
      <el-form
        ref="dataForm"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="70px"
        style="width: 400px; margin-left: 50px"
      >
        <el-form-item label="状态">
          <el-select v-model="bug.status" class="filter-item" placeholder="Please select">
            <el-option
              v-for="(item, index) in statusOptions"
              :key="index"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item style="margin-bottom: 40px" label="任务给：">
          <el-select v-model="temp.selectusers" multiple placeholder="请选择指定的用户">
            <el-option
              v-for="(item, index) in users"
              :key="index"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="理由">
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
    </el-dialog> -->
  </div>
</template>

<script>
import { showBug, passBug } from "@/api/bugs";

export default {
  name: "ShowBug",
  // components: { Sticky },
  data() {
    return {
      users: [],
      loading: false,
      bug: {
        title: "",
        appversion: "",
        status: "",
        comments: [],
      },
      statusOptions: [],
      dialogFormVisible: false,
      textMap: {
        update: "Edit",
        create: "Create",
      },
      dialogStatus: "",
      temp: {
        id: undefined,
        remark: "",
        status: "",
        selectusers: "",
      },
      myBackToTopStyle: {
        right: "50px",
        bottom: "50px",
        width: "40px",
        height: "40px",
        "border-radius": "4px",
        "line-height": "45px", // 请保持与高度一致以垂直居中 Please keep consistent with height to center vertically
        background: "#e7eaf1", // 按钮的背景颜色 The background color of the button
      },
      rules: {
        type: [{ required: true, message: "type is required", trigger: "change" }],
        timestamp: [
          {
            type: "date",
            required: true,
            message: "timestamp is required",
            trigger: "change",
          },
        ],
        title: [{ required: true, message: "title is required", trigger: "blur" }],
      },
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      const url = window.location.href;
      const ul = url.split("/");
      const id = ul[ul.length - 1];
      if (id % 1 === 0) {
        this.temp.id = parseInt(id);
        showBug(id).then((resp) => {
          this.bug = resp.data.data;
          document.title = this.bug.title;
        });
      }
    },
    // updateData() {
    //   this.dialogFormVisible = true;
    //   this.temp.status = this.bug.status;
    //   passBug(this.temp).then(resp => {
    //     const data = resp.data;
    //     this.bug.comment.push({
    //       date: data.date,
    //       info: data.remark,
    //       user: data.user,
    //       passuser: data.selectusers
    //     });
    //     this.temp.remark = "";
    //     this.temp.status = data.status;
    //     this.temp.selectusers = "";
    //     this.$message({
    //       message: "操作成功",
    //       type: "success"
    //     });
    //   });
    //   this.dialogFormVisible = false;
    // },
    // handleDelete(row) {
    //   this.$notify({
    //     title: "成功",
    //     message: "删除成功",
    //     type: "success",
    //     duration: 2000
    //   });
    //   const index = this.list.indexOf(row);
    //   this.list.splice(index, 1);
    // }
  },
};
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
@import "src/styles/mixin.scss";

.createPost-container {
  position: relative;
  .createPost-main-container {
    padding: 40px 45px 20px 50px;
    .postInfo-container {
      position: relative;
      @include clearfix;
      margin-bottom: 10px;
      .postInfo-container-item {
        float: left;
      }
    }
    .editor-container {
      min-height: 500px;
      margin: 0 0 30px;
      .editor-upload-btn-container {
        text-align: right;
        margin-right: 10px;
        .editor-upload-btn {
          display: inline-block;
        }
      }
    }
  }
  .word-counter {
    width: 40px;
    position: absolute;
    right: -10px;
    top: 0px;
  }
}
</style>
<style>
#tinymcecontent img {
  max-width: 800px;
  text-align: center;
}
</style>
