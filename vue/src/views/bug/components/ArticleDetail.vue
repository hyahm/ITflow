<template>
  <div class="createPost-container">
    <el-form ref="postForm" :model="postForm" class="form-container">
      <sticky :class-name="'sub-navbar ' + postForm.status">
        <!--<CommentDropdown v-model="postForm.comment_disabled" />-->
        <!--<PlatformDropdown v-model="postForm.platforms" />-->
        <!--<SourceUrlDropdown v-model="postForm.source_uri" />-->
        <el-button
          :disabled="ispub"
          style="margin-left: 10px"
          type="success"
          @click="submitForm"
          >发布
        </el-button>
      </sticky>

      <div class="" style="padding: 20px">
        <el-row>
          <el-form-item label="任务类型: ">
            <el-radio-group v-model="typ">
              <el-radio v-for="(t, key) in ts" :key="key" :label="key">{{ t }}</el-radio>
            </el-radio-group>
          </el-form-item>
          <!--<el-col :span="24" >-->
          <el-form-item prop="title" label="文章标题：">
            <el-input
              v-model="postForm.title"
              :maxlength="100"
              placeholder="请输入标题"
              clearable
              style="width: 80%"
            />
          </el-form-item>
        </el-row>

        <el-form-item style="display: inline-block; width: 400px" label="项目名称：">
          <el-select v-model="postForm.pid" @change="changeProject" placeholder="请选择">
            <el-option
              v-for="item in projects"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item
          v-if="typ == 1"
          style="display: inline-block; width: 400px"
          label="运行环境："
        >
          <el-select v-model="postForm.eid" placeholder="请选择">
            <el-option
              v-for="item in envnames"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item
          v-if="typ == 1"
          style="display: inline-block; width: 400px"
          label="应用版本："
        >
          <el-select v-model="postForm.vid" placeholder="请选择">
            <el-option
              v-for="item in versions"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item
          v-if="typ == 1"
          style="display: inline-block; width: 400px"
          label="优先级别："
        >
          <el-select v-model="postForm.lid" placeholder="请选择">
            <el-option
              v-for="item in levels"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>

        <!--<el-form-item style="margin-bottom: 40px;" prop="title">-->
        <!--<PlatformDropdown v-model="postForm.platforms" />-->
        <!--</el-form-item>-->
        <el-form-item
          v-if="typ == 1"
          style="display: inline-block; width: 400px"
          label="重要性："
        >
          <el-select v-model="postForm.iid" placeholder="请选择">
            <el-option
              v-for="important in importants"
              :key="important.id"
              :label="important.name"
              :value="important.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item style="display: inline-block; width: 400px" label="分配任务：">
          <el-select v-model="postForm.spusers" multiple placeholder="分配任务">
            <el-option
              v-for="item in users"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>

        <!--</template>-->

        <el-form-item>
          <div id="main">
            <mavon-editor ref="md" v-model="postForm.content" @imgAdd="imgAdd" />
          </div>
          <!-- <div class="editor-container">
            <Tinymce ref="editor" v-model="postForm.content" />
          </div> -->
        </el-form-item>
      </div>
    </el-form>
  </div>
</template>

<script>
import Sticky from "@/components/Sticky"; // 粘性header组件
import { fetchBug, createBug, updateBug } from "@/api/bugs";
import { uploadImg } from "@/api/uploadimg";
import {
  getEnvKeyName,
  getProjectKeyName,
  getLevels,
  getImportants,
  getTyp,
  getUserKeyNameByProject,
  getVersionKeyNameByProject,
} from "@/api/get";

const defaultForm = {
  // status: 'draft',
  title: "", // 文章题目
  content: "", // 文章内容
  id: 0,
  spusers: [],
  pid: undefined,
  lid: undefined,
  eid: undefined,
  iid: undefined,
  vid: undefined,
  tid: 1,
};

export default {
  name: "ArticleDetail",
  components: {
    // Tinymce,
    Sticky,
  },
  props: {
    isEdit: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      postForm: Object.assign({}, defaultForm),
      ispub: false,
      versions: [], // 随项目变化的版本号
      importants: [],
      levels: [],
      oses: [],
      users: [], // 随项目变化的用户
      projects: [],
      envnames: [],
      ts: [],
      typ: "1",
    };
  },
  activated() {
    this.getimportants();
    this.getproject();
    this.getenv();
    this.getlevels();
  },
  created() {
    this.getimportants();
    this.getproject();
    this.getlevels();
    this.getenv();
    this.gettyp();
    if (this.isEdit) {
      const id = this.$route.params && this.$route.params.id;
      this.postForm.id = parseInt(id);
      this.fetchData(id);
    } else {
      this.postForm = Object.assign({}, defaultForm);
    }
  },
  methods: {
    getversionMap(projectid) {
      getVersionKeyNameByProject({ project_id: projectid }).then((resp) => {
        this.versions = resp.data.data;
      });
    },
    getuserMap(projectid) {
      getUserKeyNameByProject({ project_id: projectid }).then((resp) => {
        this.users = resp.data.data;
      });
    },
    gettyp() {
      getTyp().then((resp) => {
        this.ts = resp.data.ts;
      });
    },
    changeProject() {
      // 选择不用项目会显示不同的用户
      this.getuserMap(this.postForm.pid);
      this.getversionMap(this.postForm.pid);
    },

    getimportants() {
      getImportants().then((resp) => {
        this.importants = resp.data.data;
      });
    },
    getlevels() {
      getLevels().then((resp) => {
        this.levels = resp.data.data;
      });
    },
    getenv() {
      getEnvKeyName().then((resp) => {
        this.envnames = resp.data.data;
      });
    },
    getproject() {
      getProjectKeyName().then((resp) => {
        this.projects = resp.data.data;
      });
    },

    fetchData(id) {
      fetchBug(id).then((resp) => {
        const dd = resp.data.data;
        this.postForm = dd;
        this.getuserMap(this.postForm.pid);
        this.getversionMap(this.postForm.pid);
      });
    },
    submitForm() {
      this.ispub = true;
      // this.postForm.display_time = parseInt(this.display_time / 1000)
      if (this.postForm.title.length > 40) {
        this.$message({
          message: "标题长度必须小于40位",
          type: "error",
        });
        this.ispub = false;
        return;
      }
      if (this.postForm.spusers.length < 1) {
        this.$message({
          message: "请选择指定给谁",
          type: "error",
        });
        this.ispub = false;
        return;
      }
      if (this.postForm.pid == undefined) {
        this.$message({
          message: "请选择项目名称",
          type: "error",
        });
        this.ispub = false;
        return;
      }
      if (this.typ === 1 && this.postForm.lid == undefined) {
        this.$message({
          message: "请选择项目级别",
          type: "error",
        });
        this.ispub = false;
        return;
      }
      if (this.typ === 1 && this.postForm.iid == undefined) {
        this.$message({
          message: "请选择项目严重程度",
          type: "error",
        });
        this.ispub = false;
        return;
      }
      if (this.postForm.content.length < 1) {
        this.$message({
          message: "请填写内容",
          type: "error",
        });
        this.ispub = false;
        return;
      }
      if (this.typ === 1 && this.postForm.eid == undefined) {
        this.$message({
          message: "请选择运行环境",
          type: "error",
        });
        this.ispub = false;
        return;
      }
      if (this.typ === 1 && this.postForm.vid == undefined) {
        this.$message({
          message: "请选择版本",
          type: "error",
        });
        this.ispub = false;
        return;
      }
      this.$refs.postForm.validate((valid) => {
        if (valid) {
          if (this.isEdit) {
            this.postForm.tid = parseInt(this.tid);
            updateBug(this.postForm).then(() => {
              this.$notify({
                title: "成功",
                message: "修改成功",
                type: "success",
              });
            });
          } else {
            createBug(this.postForm).then((resp) => {
              if (this.postForm.id === 0) {
                this.$notify({
                  title: "成功",
                  message: "发布成功",
                  type: "success",
                });
                this.$router.push({ path: "/bug/edit/" + resp.data.id });
              } else {
                this.$notify({
                  title: "成功",
                  message: "修改成功",
                  type: "success",
                });
              }
            });
          }
        }
        this.ispub = false;
      });
    },
    imgAdd(pos, $file) {
      // 第一步.将图片上传到服务器.
      var formdata = new FormData();
      formdata.append("image", $file);
      uploadImg(formdata).then((resp) => {
        this.$refs.md.$img2Url(pos, resp.data.url);
      });
    },
    draftForm() {
      if (this.postForm.content.length === 0 || this.postForm.title.length === 0) {
        this.$message({
          message: "请填写必要的标题和内容",
          type: "warning",
        });
        return;
      }
      this.$message({
        message: "保存成功",
        type: "success",
        duration: 1000,
      });
    },
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
