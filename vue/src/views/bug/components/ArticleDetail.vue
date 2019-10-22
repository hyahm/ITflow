<template>
  <div class="createPost-container">
    <el-form ref="postForm" :model="postForm" :rules="rules" class="form-container" >

      <sticky :class-name="'sub-navbar ' + postForm.status">
        <!--<CommentDropdown v-model="postForm.comment_disabled" />-->
        <!--<PlatformDropdown v-model="postForm.platforms" />-->
        <!--<SourceUrlDropdown v-model="postForm.source_uri" />-->
        <el-button v-loading="loading" :disabled="ispub" style="margin-left: 10px;" type="success" @click="submitForm">发布
        </el-button>
        <!--<el-button v-loading="loading" type="warning" @click="draftForm">草稿</el-button>-->
      </sticky>

      <div class="" style="padding: 20px">
        <el-row>
          <!--<el-col :span="24" >-->
          <el-form-item prop="title" label="文章标题：">
            <el-input
              v-model="postForm.title"
              :maxlength="100"
              placeholder="请输入标题"
              clearable
              style="width: 80%;"/>
          </el-form-item>
        </el-row>

        <el-form-item style="display: inline-block;width: 300px" label="项目名称：">
          <el-select v-model="postForm.projectname" placeholder="请选择">
            <el-option
              v-for="(item, index) in projectnames"
              :key="index"
              :label="item"
              :value="item"/>
          </el-select>
        </el-form-item>

        <el-form-item style="display: inline-block;width: 300px" label="运行环境：">
          <el-select v-model="postForm.envname" placeholder="请选择">
            <el-option
              v-for="(item, index) in envnames"
              :key="index"
              :label="item"
              :value="item"/>
          </el-select>
        </el-form-item>

        <el-form-item style="display: inline-block;width: 300px" label="应用版本：">
          <el-select v-model="postForm.version" placeholder="请选择">
            <el-option
              v-for="(item, index) in versions"
              :key="index"
              :label="item"
              :value="item"/>
          </el-select>
        </el-form-item>

        <el-form-item style="display: inline-block;width: 300px" label="优先级别：">
          <el-select v-model="postForm.level" placeholder="请选择">
            <el-option
              v-for="(item, index) in levels"
              :key="index"
              :label="item"
              :value="item"/>
          </el-select>
        </el-form-item>

        <!--<el-form-item style="margin-bottom: 40px;" prop="title">-->
        <!--<PlatformDropdown v-model="postForm.platforms" />-->
        <!--</el-form-item>-->
        <el-form-item style="display: inline-block;width: 300px" label="严重级别：">
          <el-select v-model="postForm.important" placeholder="请选择">
            <el-option
              v-for="(important, index) in importants"
              :key="index"
              :label="important"
              :value="important"/>
          </el-select>
        </el-form-item>

        <el-form-item style="display: inline-block;width: 300px" label="分配任务：">
          <el-select v-model="postForm.selectuser" multiple placeholder="分配任务">
            <el-option
              v-for="(item, index) in users"
              :key="index"
              :label="item"
              :value="item"/>
          </el-select>
        </el-form-item>

        <!--</template>-->
        <el-form-item >
          <div class="editor-container">
            <Tinymce ref="editor" v-model="postForm.content"/>
          </div>
        </el-form-item>

      </div>
      <!--<el-form-item style="margin-bottom: 40px;" label-width="45px" label="摘要:">-->
      <!--<el-input type="textarea" class="article-textarea" :rows="1" autosize placeholder="请输入内容" v-model="postForm.content_short">-->
      <!--</el-input>-->
      <!--<span class="word-counter" v-show="contentShortLength">{{contentShortLength}}字</span>-->
      <!--</el-form-item>-->

      <!--<div style="margin-bottom: 20px;">-->
      <!--<Upload v-model="postForm.image_uri" />-->
      <!--</div>-->
      <!--</div>-->
    </el-form>
  </div>
</template>

<script>
import Tinymce from '@/components/Tinymce'
// import Upload from '@/components/Upload/singleImage3'
// import MDinput from '@/components/MDinput'
// import Multiselect from 'vue-multiselect'// 使用的一个多选框组件，element-ui的select不能满足所有需求
// import 'vue-multiselect/dist/vue-multiselect.min.css'// 多选框组件css
import Sticky from '@/components/Sticky' // 粘性header组件
import { validateURL } from '@/utils/validate'
import { fetchBug, createBug } from '@/api/bugs'
import { level, important } from '@/api/defaultvalue'
import { getEnv, getProject, getUsers, getVersion, getLevels, getImportants } from '@/api/get'
// import Warning from './Warning'
// import { removeToken } from '@/utils/auth'
// import { CommentDropdown, PlatformDropdown, SourceUrlDropdown } from './Dropdown'

const defaultForm = {
  // status: 'draft',
  title: '', // 文章题目
  content: '', // 文章内容
  id: -1,
  selectuser: [],
  projectname: '',
  level: '',
  envname: '',
  important: '一般',
  version: ''
}

export default {
  name: 'ArticleDetail',
  components: {
    Tinymce,
    // MDinput,
    // Multiselect,
    Sticky
    // Warning
    // CommentDropdown,
    // PlatformDropdown,
    // SourceUrlDropdown
  },
  props: {
    isEdit: {
      type: Boolean,
      default: false
    }
  },
  data() {
    const validateSourceUri = (rule, value, callback) => {
      if (value) {
        if (validateURL(value)) {
          callback()
        } else {
          this.$message({
            message: '外链url填写不正确',
            type: 'error'
          })
          callback(null)
        }
      } else {
        callback()
      }
    }
    return {
      postForm: Object.assign({}, defaultForm),
      loading: false,
      userListOptions: [],
      ispub: false,
      rules: {
        // image_uri: [{ valiimportantsdator: validateRequire }],
        // title: [{ validator: validateRequire }],
        // content: [{ validator: validateRequire }],
        source_uri: [{ validator: validateSourceUri }]
        // source_uri: [{ validator: validateSourceUri, trigger: 'blur' }]
      },
      classname: [],
      versions: [],
      importants: [],
      levels: [],
      oses: [],
      users: [],
      projectnames: [],
      envnames: []
    }
  },
  activated() {
    this.getuser()
    this.getversion()
    this.getproject()
    this.getenv()
    this.getlevels()
    this.getimportants()
    this.defaultimportant()
    this.defaultlevel()
  },
  created() {
    this.getimportants()
    this.defaultimportant()
    this.defaultlevel()
    this.getuser()
    this.getproject()
    this.getversion()
    this.getlevels()
    this.getenv()
    if (this.isEdit) {
      const id = this.$route.params && this.$route.params.id
      this.postForm.id = parseInt(id)
      this.fetchData(id)
    } else {
      this.postForm = Object.assign({}, defaultForm)
    }
  },
  methods: {
    defaultimportant() {
      important().then(resp => {
        if (resp.data.statuscode === 0) {
          this.postForm.important = resp.data.defaultimportant
        }
      })
    },
    getimportants() {
      getImportants().then(resp => {
        if (resp.data.statuscode === 0) {
          if (resp.data.importants !== null) {
            this.importants = resp.data.importants
          }
        }
      })
    },
    defaultlevel() {
      level().then(resp => {
        if (resp.data.statuscode === 0) {
          this.postForm.level = resp.data.defaultlevel
        }
      })
    },
    getlevels() {
      getLevels().then(resp => {
        if (resp.data.statuscode === 0) {
          if (resp.data.levels !== null) {
            this.levels = resp.data.levels
          }
        }
      })
    },
    getenv() {
      getEnv().then(response => {
        if (response.data.statuscode === 0) {
          this.envnames = response.data.envlist
        }
      })
    },
    getproject() {
      getProject().then(response => {
        if (response.data.statuscode === 0) {
          this.projectnames = response.data.projectlist
        }
      })
    },
    getversion() {
      getVersion().then(response => {
        if (response.data.statuscode === 0) {
          this.versions = response.data.versionlist
        }
      }).catch(err => {
        console.log(err)
      })
    },
    getuser() {
      getUsers().then(resp => {
        if (resp.data.statuscode === 0) {
          if (resp.data.users !== null) {
            this.users = resp.data.users
          }
        }
      }).catch(err => {
        console.log(err)
      })
    },
    fetchData(id) {
      fetchBug(id).then(response => {
        if (response.data.statuscode === 0) {
          const dd = response.data
          this.postForm.title = dd.title
          this.postForm.content = dd.content
          this.postForm.importance = dd.importance
          this.postForm.version = dd.version
          this.postForm.selectuser = dd.handle
          this.postForm.envname = dd.env
          this.postForm.projectname = dd.projectname
        }
      }).catch(err => {
        console.log(err)
      })
    },
    submitForm() {
      this.ispub = true
      // this.postForm.display_time = parseInt(this.display_time / 1000)
      if (this.postForm.title.length > 40) {
        this.$message({
          message: '标题长度必须小于40位',
          type: 'error'
        })
        this.ispub = false
        return
      }
      if (this.postForm.selectuser.length < 1) {
        this.$message({
          message: '请选择指定给谁',
          type: 'error'
        })
        this.ispub = false
        return
      }
      if (this.postForm.projectname.length < 1) {
        this.$message({
          message: '请选择项目名称',
          type: 'error'
        })
        this.ispub = false
        return
      }
      this.$refs.postForm.validate(valid => {
        if (valid) {
          createBug(this.postForm).then(resp => {
            if (resp.data.statuscode === 0) {
              if (this.postForm.id === -1) {
                this.$notify({
                  title: '成功',
                  message: '发布成功',
                  type: 'success'
                })
              } else {
                this.$notify({
                  title: '成功',
                  message: '修改成功',
                  type: 'success'
                })
              }
            } else {
              this.$notify({
                title: '成功',
                message: '修改失败',
                type: 'error'
              })
            }
            // this.$router.push('/bug/allbugs')
          })
        }
        this.ispub = false
      })
    },
    draftForm() {
      if (this.postForm.content.length === 0 || this.postForm.title.length === 0) {
        this.$message({
          message: '请填写必要的标题和内容',
          type: 'warning'
        })
        return
      }
      this.$message({
        message: '保存成功',
        type: 'success',
        showClose: true,
        duration: 1000
      })
      // this.postForm.status = 'draft'
    }
  }
}

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
