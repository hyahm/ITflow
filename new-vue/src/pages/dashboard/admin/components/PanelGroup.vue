<template>
  <div></div>
</template>

<script setup lang="ts">
// import { getprojectlist } from 'src/api/dashboard';
// import { reactive, watch } from 'vue';

// const group = reactive({
//   countusers: 0,
//   countgroups: 0,
//   countbugs: 0,
//   countcomplete: 0,
//   projectnames: [],
// });
// watch(
//   group.projectnames,
//   async () => {
//     await getprojectcount();
//   },
//   { deep: true },
// );

// async function getlist() {
//   await getlist().then((response) => {
//     if (response.data.code === 0) {
//       group.countusers = response.data.data.countusers;
//       group.countgroups = response.data.data.countgroups;
//     }
//   });
// }
// async function getprojectcount() {
//   await getprojectlist().then((response) => {
//     if (response.data.code === 0) {
//       group.countbugs = response.data.data.countbugs;
//       group.countcomplete = response.data.data.countcomplete;
//     }
//   });
// }
// function handleSetLineChartData(type: string) {
//   const emit = defineEmits(['handleSetLineChartData']);
//   emit('handleSetLineChartData', type);
// }
// getpname() {
//   getproject().then(response => {
//     const arr = response.data
//     for (let i = 0; i < arr.length; i++) {
//       const aa = {}
//       aa.value = arr[i]
//       aa.label = arr[i]
//       this.projectnames.push(aa)
//     }
//   })
// }
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.panel-group {
  margin-top: 18px;
  .card-panel-col {
    margin-bottom: 32px;
  }
  .card-panel {
    height: 108px;
    cursor: pointer;
    font-size: 12px;
    position: relative;
    overflow: hidden;
    color: #666;
    background: #fff;
    box-shadow: 4px 4px 40px rgba(0, 0, 0, 0.05);
    border-color: rgba(0, 0, 0, 0.05);
    &:hover {
      .card-panel-icon-wrapper {
        color: #fff;
      }
      .icon-people {
        background: #40c9c6;
      }
      .icon-message {
        background: #36a3f7;
      }
      .icon-money {
        background: #f4516c;
      }
      .icon-shoppingCard {
        background: #34bfa3;
      }
    }
    .icon-people {
      color: #40c9c6;
    }
    .icon-message {
      color: #36a3f7;
    }
    .icon-money {
      color: #f4516c;
    }
    .icon-shoppingCard {
      color: #34bfa3;
    }
    .card-panel-icon-wrapper {
      float: left;
      margin: 14px 0 0 14px;
      padding: 16px;
      transition: all 0.38s ease-out;
      border-radius: 6px;
    }
    .card-panel-icon {
      float: left;
      font-size: 48px;
    }
    .card-panel-description {
      float: right;
      font-weight: bold;
      margin: 26px;
      margin-left: 0px;
      .card-panel-text {
        line-height: 18px;
        color: rgba(0, 0, 0, 0.45);
        font-size: 16px;
        margin-bottom: 12px;
      }
      .card-panel-num {
        font-size: 20px;
      }
    }
  }
}
</style>
