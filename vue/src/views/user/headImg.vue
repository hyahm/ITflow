<template>
  <div id="head" style="padding-left: 20px">
    <p class="warn-content">
      修改个人头像
    </p>
    <el-button type="primary" @click="toggleShow">设置头像</el-button>
    <my-upload
      v-model="show"
      :params="params"
      :headers="headers"
      :url="url"
      field="upload"
      method="POST"
      @crop-success="cropSuccess"
      @crop-upload-success="cropUploadSuccess"
      @crop-upload-fail="cropUploadFail"
    />
    <img :src="avatar">
  </div>
</template>

<script>
import myUpload from 'vue-image-crop-upload'
import { getToken } from '@/utils/auth'
// import { mapState, mapGetters } from 'vuex'
export default {
  name: 'HeadImg',
  components: {
    myUpload
  },
  data() {
    return {
      show: false,
      imgname: 'head',
      url: process.env.VUE_APP_BASE_API + '/upload/headimg',
      params: {
        token: getToken(),
        name: 'avatar'
      },
      headers: {
        smail: '*_~',
        'X-Token': getToken()
      }
      // avatar: ''
    }
  },
  computed: {
    avatar: {
      get() {
        return this.$store.state.user.avatar
      },
      set(val) {
        this.$store.commit('user/SET_AVATAR', val)
      }
    }
  },
  created() {
    this.checktoken()
  },
  methods: {
    checktoken() {
      if (getToken() === '') {
        this.$router.push('/login')
      }
    },
    toggleShow() {
      this.show = !this.show
      this.imgname = (new Date()).valueOf().toString()
    },
    cropSuccess(imgDataUrl, field) {

      // console.log("imgDataUrl")
      // console.log(imgDataUrl)
      // this.imgDataUrl = imgDataUrl
      // console.log('cropSuccess')
    },
    /**
     * upload success
     *
     * [param] jsonData   服务器返回数据，已进行json转码
     * [param] field
     */
    cropUploadSuccess(jsonData, field) {
      this.avatar = jsonData.url
      this.$store.dispatch('user/setHeadImage', jsonData.url).then(() => {
        this.loading = false
        // this.$router.push({ path: this.redirect || '/' })
      }).catch(() => {
        this.loading = false
      })
    },
    /**
     * upload fail
     *
     * [param] status    server api return error status, like 500
     * [param] field
     */
    cropUploadFail(status, field) {
      console.log('field: ' + field)
    }
  }
}
</script>
