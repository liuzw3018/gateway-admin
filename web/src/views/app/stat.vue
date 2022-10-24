<template>
  <div class="chart-container">
    <chart height="100%" width="100%" :data="chartData" />
  </div>
</template>

<script>
import Chart from './components/LineStat'
import { appStat } from '@/api/app'

export default {
  name: 'AppStat',
  components: { Chart },
  data() {
    return {
      chartData: {
        app_name: '',
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
      appStat(query).then(response => {
        this.chartData = {
          app_name: response.data.app_name + ' 租户流量统计',
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

