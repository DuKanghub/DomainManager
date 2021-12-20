<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="任务名称">
          <el-input v-model="searchInfo.job_name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="命令">
          <el-input v-model="searchInfo.command" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="创建者">
          <el-input v-model="searchInfo.creator" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="执行主机">
          <el-input v-model="searchInfo.exec_host" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="启用" prop="enabled">
          <el-select v-model="searchInfo.enabled" clearable placeholder="请选择">
            <el-option
              key="true"
              label="是"
              value="true"
            />
            <el-option
              key="false"
              label="否"
              value="false"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="searchInfo.comment" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
          <el-popover v-model="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
              <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <el-button slot="reference" icon="el-icon-delete" size="mini" type="danger" style="margin-left: 10px;">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      ref="multipleTable"
      border
      stripe
      style="width: 100%"
      tooltip-effect="dark"
      :data="tableData"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{ scope.row.CreatedAt|formatDate }}</template>
      </el-table-column>
      <el-table-column label="任务名称" prop="job_name" width="120" />
      <el-table-column label="执行频率" prop="spec" width="180" />
      <el-table-column label="命令" prop="command" width="320" />
      <el-table-column label="创建者" prop="creator" width="80" />
      <el-table-column label="执行主机" prop="exec_host" width="120" />
      <el-table-column label="启用" prop="enabled" width="50">
        <template slot-scope="scope">
          <span :style="{color: scope.row.enabled ? 'green' : 'red'}">{{ scope.row.enabled|formatBoolean }}</span>
        </template>
      </el-table-column>
      <el-table-column label="备注" prop="comment" width="120" /> <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateCronJob(scope.row)">变更</el-button>
          <el-button size="small" type="warning" icon="el-icon-edit" class="table-button" @click="deployCronJob(scope.row)">部署</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="任务名称:">

          <el-input v-model="formData.job_name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="执行频率:">

          <el-input v-model="formData.spec" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="命令:">
          <el-input v-model="formData.command" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="执行主机:">
          <el-input v-model="formData.exec_host" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="启用:">

          <el-switch v-model="formData.enabled" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="备注:">

          <el-input v-model="formData.comment" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createCronJob,
  deleteCronJob,
  deleteCronJobByIds,
  updateCronJob,
  findCronJob,
  deployCronJob,
  getCronJobList
} from '@/api/cronJob' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'
export default {
  name: 'CronJob',
  filters: {
    formatDate: function(time) {
      if (time !== null && time !== '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? '是' : '否'
      } else {
        return ''
      }
    }
  },
  mixins: [infoList],
  data() {
    return {
      listApi: getCronJobList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],

      formData: {
        job_name: '',
        spec: '',
        command: '',
        creator: 0,
        exec_host: '',
        enabled: true,
        comment: ''

      }
    }
  },
  async created() {
    await this.getTableData()
  },
  methods: {
  // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.enabled === '') {
        this.searchInfo.enabled = null
      }
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteCronJob(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteCronJobByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateCronJob(row) {
      const res = await findCronJob({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.recronJob
        this.dialogFormVisible = true
      }
    },
    async deployCronJob(row) {
      if (row.exec_host === '') {
        this.$message({
          type: 'error',
          message: '没有指定执行的主机IP'
        })
        return
      }
      const res = await deployCronJob(row)
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '部署成功'
        })
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        job_name: '',
        spec: '',
        command: '',
        creator: 0,
        exec_host: '',
        enabled: true,
        comment: ''

      }
    },
    async deleteCronJob(row) {
      const res = await deleteCronJob({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case 'create':
          res = await createCronJob(this.formData)
          break
        case 'update':
          res = await updateCronJob(this.formData)
          break
        default:
          res = await createCronJob(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  }
}
</script>

<style>
</style>
