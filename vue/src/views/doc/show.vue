<template>
  <div class="chart-container">
    <div class="doc_content">
      <h2 class="title" v-text="title" />
      <div v-html="data" />
    </div>
    <div class="doc_directory" />

  </div>
</template>

<script>
import { addProject } from '@/api/doc'
import { markdown } from 'markdown'
export default {
  name: 'Show',
  data() {
    return {
      projects: ['aaa', 'bbb', 'ccc'],
      data: '# hello',
      title: ''
    }
  },
  mounted() {
    this.data = markdown.toHTML(this.data)
  },
  methods: {
    getcontent() {

    },
    one() {
      console.log(12222)
      this.projects.push(new Date())
      console.log(this.projects)
    },
    addproject() {
      addProject().then(resp => {
        console.log(resp.data)
      })
    }
  }
}
</script>

<style>
.el-card__body {
  padding: 0px !important;
}
</style>

<style scoped>
.char-container{
  position: relative;
  padding:20px;
  width: 100%;
  height:85vh;
}
.box-card {
  padding: 0px;
  width: 180px;
  font-weight: 700;
  color: #777;
  height: 100px;
   margin: 20px;
}

.add-card {
  cursor: pointer;
  width: 180px;
  padding: 20px;
   height: 100px;
   font-weight: 700;
  margin: 20px;
}
.project-container {
  margin: 0 auto;
  display: flex;
  /* flex-direction:column; */
  flex-wrap: wrap ;
  max-width: 700px;
  margin-top: 30px;
}
.setting {
  float: right;
}
.div_box_card {
  padding: 20px
}
.doc_content {
    margin: 0 auto;
    width: 700px;
}
</style>

