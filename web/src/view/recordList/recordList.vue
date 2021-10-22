<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" ref="searchInfoRef">
        <el-form-item label="所属账号:">
          <el-select v-model="searchInfo.account" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in DnsRamUserOptions" :key="key" :label="item.label" :value="item.value"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="域名:">
          <el-input placeholder="dukanghub.com" v-model="searchInfo.domain" @keyup.enter.native="onSubmit" />
        </el-form-item>
        <el-form-item label="类型:">
          <el-select v-model="searchInfo.type" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in RecordTypeOptions" :key="key" :label="item.label" :value="item.value"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="线路:">
          <el-select v-model="searchInfo.line" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in RecordLineOptions" :key="key" :label="item.label" :value="item.value"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="主机记录:">
            <el-input placeholder="www" v-model="searchInfo.rr" @keyup.enter.native="onSubmit" />
        </el-form-item>
        <el-form-item label="记录值:">
            <el-input placeholder="IP或域名" v-model="searchInfo.value" @keyup.enter.native="onSubmit" />
        </el-form-item>
        <el-form-item label="状态：">
          <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in RecordStatusOptions" :key="key" :label="item.label" :value="item.value"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
          <el-button icon="el-icon-edit" size="mini" type="primary" @click="updateDialog">批量更新</el-button>
          <el-button icon="el-icon-edit" size="mini" type="primary" @click="flushrecords">刷新解析缓存</el-button>
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
      <el-table-column label="更新日期" width="180">
        <template slot-scope="scope">{{ scope.row.UpdatedAt|formatDate }}</template>
      </el-table-column>
      <el-table-column label="域名" prop="domain" width="120" />
      <el-table-column label="主机记录" prop="rr" width="80" />
    
      <el-table-column label="记录类型" prop="type" width="120">
        <template slot-scope="scope">
          {{filterDict(scope.row.type,"RecordType")}}
        </template>  
      </el-table-column>     
      <el-table-column label="线路" prop="line" width="80">
        <template slot-scope="scope">
          {{filterDict(scope.row.line,"RecordLine")}}
        </template>   
      </el-table-column>    
      <el-table-column label="记录值" prop="value" width="180" />
      <el-table-column label="状态" prop="status" width="80">
          <template slot-scope="scope">
          {{filterDict(scope.row.status,"RecordStatus")}}
        </template>
      </el-table-column>   
      <el-table-column label="所属账号" prop="account" width="120">
        <template slot-scope="scope">
          {{filterDict(scope.row.account,"DnsRamUser")}}
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateRecord(scope.row)">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 20, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
    <!-- 变更修改的弹窗 -->
    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="域名:" v-if='isShow'>
          <el-input v-model="formData.domain" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="主机记录:" v-if='isShow'>
          <el-input v-model="formData.rr" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型:" v-if='isShow'>
          <el-select v-model="formData.type" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in RecordTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="线路:" v-if='isShow'>
          <el-select v-model="formData.line" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in RecordLineOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="记录值:">
          <el-input v-model="formData.value" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:">
          <el-select v-model="formData.status" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in RecordStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="所属账号:">
          <el-select v-model="formData.account" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in DnsRamUserOptions" :key="key" :label="item.label" :value="item.value" />
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
  createRecord,
  deleteRecord,
  deleteRecordByIds,
  updateRecord,
  updateRecords,
  findRecord,
  flushRecords,
  getRecordList
} from '@/api/recordList' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'
export default {
  name: 'RecordList',
  mixins: [infoList],
  data() {
    return {
      listApi: getRecordList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      RecordLineOptions: [],
      RecordTypeOptions: [],
      RecordStatusOptions: [],
      DnsRamUserOptions: [],
      searchInfo: { //解决mounted后域名搜索不能输入的问题
        domain: ""
      },
      formData: {
        domain: '',
        rr: '',
        type: 1,
        line: 1,
        value: '',
        record_id: '',
        status: 1,
        account: 0
      },
      isShow: true
    }
  },
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
  async created() {
    if (this.$route.params.domain) {//域名详情跳转过来使用
      this.searchInfo.domain = this.$route.params.domain
      this.onSubmit()
    } else {
      await this.getTableData()
    }
    
    await this.getDict('DnsRamUser')
    await this.getDict('RecordType')
    await this.getDict('RecordLine')
    await this.getDict('RecordStatus')
  },
  // mounted() { //域名详情跳转过来使用
  //   console.log(this.$route.params)
  //   this.searchInfo.domain = this.$route.params.domain
  //   // this.page = this.$route.params.page
  //   // this.pageSize = this.$route.params.pageSize
  //   this.onSubmit()
  // },
  methods: {
  // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
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
        this.deleteRecord(row)
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
      const res = await deleteRecordByIds({ ids })
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
    async flushrecords() {
      const res = await flushRecords()
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '更新成功'
        })
        this.getTableData()
      }
    },
    // 更新单条数据
    async updateRecord(row) {
      const res = await findRecord({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        // rerecord来自后端findRecord定义
        this.formData = res.data.rerecord
        this.dialogFormVisible = true
      }
    },
    // 批量更新
    async onUpdateMulti(formData) {
      const records = this.$refs.multipleTable.selection
      for (let i = 0; i < records.length; i++) {
        records[i].value = formData.value || records[i].value
        records[i].status = formData.status || records[i].status
        records[i].account = formData.account || records[i].account
      }
      const res = await updateRecords({ records })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '更新成功'
        })
        if (this.tableData.length === records.length && this.page > 1) {
          this.page--
        }
        //this.updateVisible = false
        this.getTableData()
        this.closeDialog()
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        domain: '',
        rr: '',
        type: 0,
        line: 0,
        value: '',
        record_id: '',
        status: 0,
        account: 0
      }
    },
    async deleteRecord(row) {
    
      const res = await deleteRecord({ ID: row.ID })
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
          
          res = await createRecord(this.formData)
          break
        case 'update':
          res = await updateRecord(this.formData)
          break
        case 'updateMulti':
          this.onUpdateMulti(this.formData)
          break
        default:
          res = await createRecord(this.formData)
          break
      }
      if (this.type !== 'updateMulti') {  
        if (res.code === 0) {
            this.$message({
            type: 'success',
            message: '创建/更改成功'
            })
            this.closeDialog()
            this.getTableData()
        }
      }
    },
    openDialog() {
      this.type = 'create'
      this.formData = { //新增记录时设置一些默认值
            domain: this.searchInfo.domain,
            rr: '',
            type: 1,
            line: 1,
            value: '',
            status: 1,
            account: 1
          }
      this.dialogFormVisible = true
      this.isShow = true
    },
    updateDialog() {
      if (this.$refs.multipleTable.selection.length <= 0) {
        this.$message.error('请选择需要批量更新的内容')
      } else {
        this.formData = {}
        this.type = 'updateMulti'
        this.dialogFormVisible = true
        this.isShow = false
      }
    }
  }
}
</script>

<style>
</style>
