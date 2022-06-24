<template>
  <el-container>
    <el-main style="padding-top:0">
      <div id="dashboard">
        <el-card>
          <template #header>
            <div class="header">
              <div style="display:flex;justify-content:space-between">
                <el-button icon="plus" type="success" style="margin-right:10px" @click="create()">添加登录用户</el-button>
                <div>
                  <el-input
                      style="width: 400px;margin-right:10px"
                      clearable
                      @keyup.enter="fetch({page: 1})"
                      @clear="fetch({page: 1})"
                      v-model="username"
                      placeholder="输入用户名称搜索,Enter"
                  />
                  <el-button type="primary" @click="fetch({page: 1})">查询</el-button>
                </div>
              </div>
            </div>

          </template>
          <el-table :stripe="true" v-loading="loading" :data="data.items" style="width: 100%">
            <el-table-column prop="username" label="用户名称"/>
            <el-table-column prop="status" label="状态" width="80" align="center">
              <template #default="scope">
                <el-tag v-if="scope.row.status === 'Y'" type="success">启用</el-tag>
                <el-tag v-if="scope.row.status === 'N'" type="error">停用</el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="last_login_time" label="登录时间" align="center"/>
            <el-table-column prop="last_login_ip" label="登录IP" align="center"/>
            <el-table-column prop="created_at" label="创建时间" align="center"/>
            <el-table-column prop="updated_at" label="修改时间" align="center"/>
            <el-table-column fixed="right" prop="operation" label="操作" width="100">
              <template #default="scope">
                <el-link @click="update(scope.row)" v-if="scope.row.status === 'Y'" type="primary"
                         style="margin-right:10px;">修改
                </el-link>
                <el-link @click="updateStatus(scope.row, {status: 'N'})" v-if="scope.row.status === 'Y'" type="danger">
                  停用
                </el-link>
                <el-link @click="updateStatus(scope.row, {status: 'Y'})" v-if="scope.row.status === 'N'" type="primary">
                  启用
                </el-link>
              </template>
            </el-table-column>
          </el-table>
          <div v-if="data && data.total > 0" style="margin-top:10px;">
            <el-pagination
                @size-change="(page_size) => fetch({page_size})"
                @current-change="(page) => fetch({page})"
                layout="total, sizes, prev, pager, next, jumper"
                small="small"
                :page-sizes="[10, 20, 30, 50, 100]"
                background
                :page-size="data.page_size"
                :current-page="data.page"
                :total="data.total"
            />
          </div>
        </el-card>
      </div>
    </el-main>
    <Form ref="formRef" @submit="submit"/>
  </el-container>
</template>

<script setup lang="ts">
import {onMounted, ref} from 'vue'
import {PageQuery, Pagination, User, userCreateApi, userListApi, userUpdateApi} from '@/api'
import {ElMessage} from 'element-plus'
import Form from './Form.vue'
import {useSync} from '@/hooks'

const formRef = ref()
const username = ref('')
const data = ref<Pagination<User>>({items: [], page: 1, page_size: 10, total: 0})

const {loading, fetch} = useSync<Partial<PageQuery>>(async (request?: Partial<PageQuery>) => {
  const {page, page_size} = data.value
  const query: { username?: string } = {}
  if (username.value) {
    query['username'] = username.value
  }

  data.value = await userListApi({page, query, page_size, ...request})
})

const create = () => {
  formRef.value?.open({
    visible: true,
    title: '添加用户信息',
    action: 'create',
  }, {username: '', password: ''})
}

const update = (user: User) => {
  formRef.value?.open({
    visible: true,
    title: '编辑用户信息',
    action: 'update',
  }, {username: user.username, password: ''})
}

const updateStatus = (user: User, status: Pick<User, 'status'>) => {
  userUpdateApi({username: user.username, ...status}).then(() => {
    ElMessage.success('处理成功')
    fetch()
  })
}

const submit = async (value: any, {action}: {action: string}) => {
  if (action === 'create') {
    await userCreateApi(value)
    ElMessage.success('添加成功')
  } else {
    await userUpdateApi(value)
    ElMessage.success('修改成功')
  }

  await fetch()
}

onMounted(() => {
  fetch()
})
</script>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
}
</style>
