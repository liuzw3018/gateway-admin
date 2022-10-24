<template>
  <div class="dashboard-editor-container">
    <github-corner class="github-corner" />

    <panel-group @handleSetLineChartData="handleSetLineChartData" />

    <el-row :gutter="50">
      <el-col :xs="24" :sm="24" :lg="16">
        <div class="chart-wrapper">
          <line-chart :chart-data="chartData" />
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :lg="8">
        <div class="chart-wrapper">
          <pie-chart :pie-data="pieData" />
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import GithubCorner from '@/components/GithubCorner'
import PanelGroup from './components/PanelGroup'
import LineChart from './components/LineChart'
import PieChart from './components/PieChart'
import { flowStat, serviceStat } from '@/api/dashboard'

export default {
  name: 'DashboardAdmin',
  components: {
    PanelGroup,
    GithubCorner,
    LineChart,
    PieChart
  },
  data() {
    return {
      chartData: {
        today: [],
        yesterday: []
      },
      pieData: {
        http: 1,
        tcp: 2,
        grpc: 3
      }
    }
  },
  created() {
    this.handleSetLineChartData()
    this.handleSetPieChartData()
  },
  methods: {
    handleSetLineChartData() {
      flowStat().then(response => {
        this.chartData.today = response.data.today
        this.chartData.yesterday = response.data.yesterday
      })
    },
    handleSetPieChartData() {
      serviceStat().then(response => {
        for (const responseKey in response.data.data) {
          if (response.data.data[responseKey].load_type === 0) {
            this.pieData.http = response.data.data[responseKey].value
          } else if (response.data.data[responseKey].load_type === 1) {
            this.pieData.tcp = response.data.data[responseKey].value
          } else {
            this.pieData.grpc = response.data.data[responseKey].value
          }
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard-editor-container {
  padding: 32px;
  background-color: rgb(240, 242, 245);
  position: relative;

  .github-corner {
    position: absolute;
    top: 0px;
    border: 0;
    right: 0;
  }

  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
  }
}

@media (max-width: 1024px) {
  .chart-wrapper {
    padding: 8px;
  }
}
</style>
