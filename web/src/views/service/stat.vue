<template>
  <div class="chart-container">
    <chart height="100%" width="100%" :data="chartData" />
  </div>
</template>

<script>
import Chart from './components/LineStat'
import { serviceStat } from '@/api/service'

export default {
  name: 'ServiceStat',
  components: { Chart },
  data() {
    return {
      chartData: {
        title: '',
        today: [],
        yesterday: []
      }
    }
  },
  created() {
    const id = this.$route.params && this.$route.params.id
    this.fetchStat(id)
  },
  methods: {
    fetchStat(id) {
      const query = { 'id': id }
      serviceStat(query).then(response => {
        this.chartData = {
          title: response.data.service_name + ' 服务流量统计',
          today: response.data.today,
          yesterday: response.data.yesterday
        }
      }).catch(() => {

      })
    }
  }
}
</script>

<style scoped>
.chart-container {
  position: relative;
  width: 100%;
  height: calc(100vh - 84px);
}
</style>

