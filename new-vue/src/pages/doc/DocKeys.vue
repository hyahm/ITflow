<template>
  <div class="app-container"></div>
</template>

<script setup lang="ts">
// import { fetchList, addKey, delKey, checkName } from 'api/key';
// import { getToken } from 'utils/auth';
// export default {
//   name: 'KeysList',
//   data() {
//     return {
//       list: [],
//       user: 1,
//       key: 2,
//       data: {},
//       headers: {},
//       giturl: '',
//       total: 0,
//       dialogFormVisible: false,
//       auths: ['', '用户认证', 'key认证'],
//       listLoading: true,
//       formLabelWidth: '120px',
//       newname: '',
//       listQuery: {
//         page: 1,
//         limit: 20,
//       },
//       form: {
//         id: 0,
//         name: '',
//         pri: '',
//         pub: '',
//         typ: 1,
//         user: '',
//         created: 0,
//         uptime: 0,
//         password: '',
//       },
//     };
//   },
//   created() {
//     this.getList();
//   },
//   mounted() {
//     this.headers = {
//       'X-Token': getToken(),
//     };
//   },
//   methods: {
//     handleAdd() {
//       this.form = {
//         id: 0,
//         name: '',
//         pri: '',
//         pub: '',
//         typ: 1,
//         user: '',
//         created: 0,
//         uptime: 0,
//         password: '',
//       };
//       this.dialogFormVisible = true;
//     },
//     handleUpdate(row) {
//       this.form = row;
//       this.dialogFormVisible = true;
//     },
//     handleDomain(val) {
//       // 判断子域名是否被使用
//       checkName(val).then((resp) => {
//         if (resp.data.code !== 0) {
//           this.$message.error('name 重复');
//         }
//       });
//     },
//     handleAuthMethod(val) {
//       this.form.auth = val;
//     },

//     handleAddDoc() {
//       this.form.typ = parseInt(this.form.typ);
//       addKey(this.form).then((resp) => {
//         if (this.form.id > 0) {
//           for (let i = 0; i < this.list.length; i++) {
//             if (this.list[i].id == this.form.id) {
//               this.list[i] = this.form;
//               this.list[i].uptime = resp.data.time;
//               break;
//             }
//           }
//         } else {
//           this.form.id = resp.data.id;
//           this.form.created = resp.data.time;
//           this.list.push(this.form);
//         }
//       });

//       this.dialogFormVisible = false;
//       this.$message.success('更新成功');
//     },
//     getList() {
//       this.listLoading = true;
//       fetchList(this.listQuery).then((response) => {
//         const data = response.data;
//         this.list = data.keys;
//         this.total = data.total;
//         this.listLoading = false;
//       });
//     },
//     handleSizeChange(val) {
//       this.listQuery.limit = val;
//       this.getList();
//     },
//     handleCurrentChange(val) {
//       this.listQuery.page = val;
//       this.getList();
//     },
//     handleDeleteKey(id) {
//       this.$confirm('此操作将永久删除该文档, 是否继续?', '提示', {
//         confirmButtonText: '确定',
//         cancelButtonText: '取消',
//         type: 'warning',
//       })
//         .then(() => {
//           delKey(id).then((resp) => {
//             if (resp.data.code === 0) {
//               const l = this.list.length;
//               for (let i = 0; i < l; i++) {
//                 if (this.list[i].id === id) {
//                   this.list.splice(i, 1);
//                   break;
//                 }
//               }
//               this.$message({
//                 type: 'success',
//                 message: '删除成功!',
//               });
//             }
//           });
//         })
//         .catch(() => {
//           this.$message({
//             type: 'info',
//             message: '已取消删除',
//           });
//         });
//     },
//   },
// };
</script>

<style scoped>
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
