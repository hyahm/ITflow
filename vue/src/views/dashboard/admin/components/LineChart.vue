<template>
  <div :class="className" :style="{height:height,width:width}" />
</template>

<script>
import echarts from 'echarts'
require('echarts/theme/macarons') // echarts theme
import { debounce } from '@/utils'
import { getBugCount } from '@/api/dashboard'

export default {
  props: {
    className: {
      type: String,
      default: 'chart'
    },
    width: {
      type: String,
      default: '100%'
    },
    height: {
      type: String,
      default: '350px'
    },
    autoResize: {
      type: Boolean,
      default: true
    }
    // chartData: {
    //   type: Object,
    //   default: null
    // }
  },
  data() {
    return {
      chart: null,
      chartData: {
        createdData: [],
        completedData: [],
        xData: []
      }
    }
  },

  mounted() {
    this.x()
    this.getdata()
    if (this.autoResize) {
      this.__resizeHanlder = debounce(() => {
        if (this.chart) {
          this.chart.resize()
        }
      }, 100)
      window.addEventListener('resize', this.__resizeHanlder)
    }

    // 监听侧边栏的变化
    const sidebarElm = document.getElementsByClassName('sidebar-container')[0]
    sidebarElm.addEventListener('transitionend', this.__resizeHanlder)
  },
  beforeDestroy() {
    if (!this.chart) {
      return
    }
    if (this.autoResize) {
      window.removeEventListener('resize', this.__resizeHanlder)
    }

    const sidebarElm = document.getElementsByClassName('sidebar-container')[0]
    sidebarElm.removeEventListener('transitionend', this.__resizeHanlder)

    this.chart.dispose()
    this.chart = null
  },
  methods: {
    x() {
      var data = new Date()
      this.chartData.xData[0] = data.getDate()
      for (var i = 0; i < 6; i++) {
        var time = data.getTime() - 24 * 60 * 60 * 1000
        data = new Date(time)
        this.chartData.xData.unshift(data.getDate())
      }
    },
    getdata() {
      getBugCount().then(resp => {
        this.chartData.createdData = resp.data.created
        this.chartData.completedData = resp.data.completed
        this.chart = echarts.init(this.$el, 'macarons')
        this.setOptions(this.chartData)
      })
    },
    setOptions(data) {
      var { createdData, completedData, xData } = data
      this.chart.setOption({
        xAxis: {
          data: xData,
          boundaryGap: false,
          axisTick: {
            show: false
          }
        },
        grid: {
          left: 10,
          right: 10,
          bottom: 20,
          top: 30,
          containLabel: true
        },
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'cross'
          },
          padding: [5, 10]
        },
        yAxis: {
          axisTick: {
            show: true
          }
        },
        legend: {
          data: ['created', 'completed']
        },
        series: [{
          name: 'create', itemStyle: {
            normal: {
              color: '#FF005A',
              lineStyle: {
                color: '#FF005A',
                width: 2
              }
            }
          },
          smooth: true,
          type: 'line',
          data: createdData,
          animationDuration: 2800,
          animationEasing: 'cubicInOut'
        },
        {
          name: 'complete',
          smooth: true,
          type: 'line',
          itemStyle: {
            normal: {
              color: '#3888fa',
              lineStyle: {
                color: '#3888fa',
                width: 2
              },
              areaStyle: {
                color: '#f3f8ff'
              }
            }
          },
          data: completedData,
          animationDuration: 2800,
          animationEasing: 'quadraticOut'
        }]
      })
    }

  }
}
</script>
