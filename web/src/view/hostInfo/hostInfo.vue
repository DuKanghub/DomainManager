<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="IP">
          <el-input v-model="searchInfo.ip" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="主机名">
          <el-input v-model="searchInfo.hostname" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="端口">
          <el-input v-model="searchInfo.port" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="启用" prop="active">
          <el-select v-model="searchInfo.active" clearable placeholder="请选择">
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
      <el-table-column label="IP" prop="ip" width="150" />
      <el-table-column label="主机名" prop="hostname" width="120" />
      <el-table-column label="端口" prop="port" width="120" />
      <el-table-column label="启用" prop="active" width="120">
        <template slot-scope="scope">{{ scope.row.active|formatBoolean }}</template>
      </el-table-column>
      <el-table-column label="用户" prop="user_id" width="120">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="分组" prop="group_id" width="120">
        <template slot-scope="scope">
          {{ filterDict(scope.row.group_id,"HostGroup") }}
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateHostInfo(scope.row)">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
          <el-button type="warning" icon="el-icon-s-platform" size="mini" @click="sshRow(scope.row)">远程</el-button>
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
        <el-form-item label="IP:">
          <el-input v-model="formData.ip" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="主机名:">
          <el-input v-model="formData.hostname" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="端口:">
          <el-input v-model.number="formData.port" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="启用:">
          <el-switch v-model="formData.active" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="用户:">
          <el-select v-model="formData.user_id" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in SSHUsers" :key="key" :label="item.name" :value="item.ID" />
          </el-select>
        </el-form-item>
        <el-form-item label="分组:">
          <el-select v-model="formData.group_id" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in HostGroupOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  createHostInfo,
  deleteHostInfo,
  deleteHostInfoByIds,
  updateHostInfo,
  findHostInfo,
  getHostInfoList
} from '@/api/hostInfo' //  此处请自行替换地址
import { getSSHUsers } from '@/api/sshUser'
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'
import Term from '@/view/term/term'

export default {
  name: 'HostInfo',
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
      listApi: getHostInfoList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      // SSHUserOptions: [],
      SSHUsers: [],
      HostGroupOptions: [],
      formData: {
        ip: '',
        hostname: '',
        port: 22,
        active: true,
        user_id: 1,
        group_id: 0
      }
    }
  },
  async created() {
    await this.getTableData()
    // await this.getDict('SSHUser')
    await this.getDict('HostGroup')
    const res = await getSSHUsers()
    if (res.code === 0) {
      this.SSHUsers = res.data
    }
    // console.log(this.SSHUsers)
  },
  methods: {
  // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.active === '') {
        this.searchInfo.active = null
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
        this.deleteHostInfo(row)
      })
    },
    sshRow(row) {
      console.log(row)
      // 新标签页打开, TODO: 跳转后会报404
      // const newUrl = this.$router.resolve(
      //   {
      //     path: 'term',
      //     params: { id: row.ID }
      //   }
      // )
      // console.log(newUrl)
      // window.open(newUrl.href, '_blank')
      // 直接写死跳转地址，可以成功跳转
      window.open('/#/layout/hostManager/term/' + row.ID, '_blank')
      // 当前页面打开
      // this.$router.push({
      //   name: 'term',
      //   params: {
      //     id: row.ID
      //   }
      // })
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
      const res = await deleteHostInfoByIds({ ids })
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
    async updateHostInfo(row) {
      const res = await findHostInfo({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.rehost
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        ip: '',
        hostname: '',
        port: 22,
        active: true,
        user_id: 1,
        group_id: 0
      }
    },
    async deleteHostInfo(row) {
      const res = await deleteHostInfo({ ID: row.ID })
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
          res = await createHostInfo(this.formData)
          break
        case 'update':
          console.log(this.formData)
          res = await updateHostInfo(this.formData)
          break
        default:
          res = await createHostInfo(this.formData)
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
