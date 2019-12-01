<template>
  <div style="padding: 15px">
    <h1 style="text-align: center">{{ list.name }}</h1>
    <div>
      <h3>请求url:</h3>
      <pre><code>{{ list.url }}</code></pre>
    </div>
    <div>
      <h3>请求方式:</h3>
      <pre v-for="(m, index) in list.methods" :key="index"><code>{{ m }}</code></pre>
    </div>

    <div v-if="list.header !== null">
      <h3>请求头:</h3>
      <table>
        <tr v-for="hr in list.header" :key="hr.key">
          <td>{{ hr.key }}:</td>
          <td style="width: 300px"><el-input v-model="hr.value" /></td>
        </tr>
      </table>
      <!--<div style="width: 10%;float: left;padding: 3px;">{{ hr.key }}</div> : <el-input style="width: 60%" v-model="hr.value"></el-input>-->
      <p v-if="list.remark !== ''"><span style="color: blue">说明：</span>{{ list.remark }}</p>
    </div>

    <div v-if="list.opts !== null">
      <h3>选项:</h3>
      <el-table
        :data="list.opts"
        style="width: 100%"
      >
        <el-table-column
          label="key"
          width="180"
        >
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.name }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="类型"
          width="180"
        >
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.type }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="是否必须"
          width="180"
        >
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.need }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="默认值"
          width="180"
        >
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.default }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="说明"
          width="180"
        >
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ scope.row.info }}</span>
          </template>
        </el-table-column>

      </el-table>
    </div>
    <div v-if="list.resp !== ''">
      <h3>请求参数：</h3>
      <el-input v-model="list.resp" />
    </div>
    <div v-if="list.result !== ''">
      <h3>返回结果：</h3>
      <pre v-if="list.calltype === 'json'">{{ list.result | jsonpretty }}</pre>
      <pre v-else>{{ list.result }}</pre>
    </div>

    <!--<vue-markdown :source="list.result" />-->
    <div v-if="list.information !== ''">
      <h3>说明：</h3>
      <vue-markdown :source="list.information" />
    </div>
    <div style="padding: 10px">
      <el-button @click="handleResp">请求</el-button>
    </div>
    <el-input v-model="url" placeholder="请求的url" />

    <pre>{{ result }}</pre>
  </div>
</template>

<script>
import { apiOne, apiResp } from '@/api/restful'
import VueMarkdown from 'vue-markdown'
// import JsonViewer from 'vue-json-viewer'
export default {
  name: 'ShowApi',
  components: {
    VueMarkdown
  },
  filters: {
    jsonpretty: function(value) {
      return JSON.stringify(JSON.parse(value), null, 4)
    }
  },
  data() {
    return {
      radio: 'text',
      url: '',
      list: {
        name: '',
        methods: [],
        opt: null,
        header: [],
        resp: '',
        result: '',
        url: '',
        information: ''
      },
      result: '',
      isjson: false,
      id: -1
    }
  },
  created() {
    this.GetQueryString()
  },
  methods: {
    GetQueryString() {
      this.id = parseInt(window.location.search.split('=')[1])
      this.getapi()
    },
    handleResp() {
      const data = {
        header: this.list.header,
        method: this.list.methods[0],
        resp: this.list.resp,
        url: this.url + this.list.url
      }

      apiResp(data).then(resp => {
        this.result = resp.data
        console.log(this.result)
        // if (typeof(this.result) == 'object' && Object.prototype.toString.call(this.result).toLowerCase() == '[object object]' && !this.result.length){
        //   this.isjson = true
        // }
      })
    },
    getapi() {
      apiOne(this.id).then(resp => {
        if (resp.data.code === 0) {
          this.list = resp.data
          // this.list.resp = JSON.parse(this.list.resp)
          // if (this.list.calltype === 'json') {
          //   this.list.result = JSON.parse(this.list.result)
          // }
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    }
  }
}
</script>

<style scoped type="text/css">
h3 {
  color: red;
}
.el-input__inner {
  width: 60px;
}
</style>
