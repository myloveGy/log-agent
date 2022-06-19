<template>
  <el-container>
    <el-main style="padding-top:0">
      <div id="dashboard">
        <el-card>
          <template #header>
            <div class="header">
              <el-input @keyup.enter="query({page: 1})" clearable v-model="search"
                        placeholder="输入查询条件,Enter"/>
              <el-date-picker
                  v-model="datetime"
                  type="datetimerange"
                  start-placeholder="开始时间"
                  end-placeholder="结束时间"
                  style="margin-left: 8px; width: 460px"
                  value-format="YYYY-MM-DD HH:mm:ss"
                  :default-time="[new Date(2000, 2, 1, 0, 0, 0), new Date(2000, 2, 1, 23, 59, 59)]"
              />
              <el-button style="margin-left: 8px;width: 100px" type="primary" @click="query({page: 1})">查询</el-button>
            </div>

          </template>
          <div v-if="data" class="header">
            <div>
              <el-button type="primary" @click="saveJson('日志文件', data.items)" size="small" :icon="Download" circle/>
            </div>
            <div v-if="data && data.total > 1">
              <el-pagination
                  @size-change="(size) => query({page_size:size})"
                  @current-change="(page) => query({page})"
                  layout="total, sizes, prev, pager, next, jumper"
                  small="small" :page-sizes="[10,20,30,50,100, 200, 300, 500]"
                  background
                  :page-size="data.page_size"
                  :current-page="data.page"
                  :total="data.total"
              />
            </div>
          </div>

          <el-table
              :stripe="true"
              v-loading="loading"
              :data="data.items"
              style="width: 100%"
              @sort-change="sortChange"
              default-expand-all
          >
            <el-table-column fixed type="expand" prop="document" label="Document" width="100">
              <template #default="scope">
                <json-viewer
                    :value="scope.row"
                    :expand-depth=10
                    copyable
                    boxed
                    sort
                />
              </template>
            </el-table-column>
            <el-table-column prop="datetime" label="datetime" width="180" sortable/>
            <el-table-column prop="server_name" label="server_name"/>
            <el-table-column prop="request_method" label="request_method"/>
            <el-table-column prop="request_uri" label="request_uri"/>
            <el-table-column prop="request_duration" label="request_duration"
                             :formatter="(row, column, value) => value.toFixed(4)"/>
            <el-table-column prop="ip" label="ip" width="180"/>
            <el-table-column prop="city" label="city" width="180" show-overflow-tooltip/>
          </el-table>
          <div v-if="data && data.total > 1" style="margin-top:10px;">
            <el-pagination
                @size-change="(size) => query({page_size:size})"
                @current-change="(page) => query({page})"
                layout="total, sizes, prev, pager, next, jumper"
                small="small"
                :page-sizes="[10, 20, 30, 50, 100, 200, 300, 500]"
                background
                :page-size="data.page_size"
                :current-page="data.page"
                :total="data.total"
            />
          </div>
        </el-card>
      </div>
    </el-main>
  </el-container>
</template>

<script setup lang="ts">
import JsonViewer from 'vue-json-viewer'
import {Download} from '@element-plus/icons-vue'
import {databaseQueryApi} from '../api'
import {onMounted, ref} from 'vue'
import {ElMessage} from 'element-plus'
import {saveJson} from '../utils'

const formatQuery = (query: string) => {
  if (query.substring(0, 1) === '{' && query.substring(-1, 1) === '}') {
    return JSON.parse(query)
  }

  const search: Record<string, string> = {}
  query.split('and').forEach(v => {
    const condition = v.trim().split(':')
    if (condition.length === 2) {
      const name = condition[0].trim()
      search[name] = condition[1].trim()
    }
  })

  return search
}

const loading = ref<boolean>(false)
const datetime = ref<string[]>([])
const search = ref<string>('')
const params = ref<{ collection: string, order: 'desc' | 'asc' }>({
  collection: 'laravel',
  order: 'desc',
})

const data = ref<any>({items: [], total: 0, page_size: 10, page: 1})

const query = (requestQuery: any = {}) => {
  loading.value = true
  const {page, page_size} = data.value
  const request: any = {
    page,
    page_size,
    query: {},
    ...params.value,
    ...requestQuery,
  }

  // 时间搜索
  if (datetime.value && datetime.value.length === 2) {
    request.start_time = datetime.value[0]
    request.end_time = datetime.value[1]
  }

  try {
    if (search.value) {
      request.query = formatQuery(search.value)
    }
  } catch (e) {
    ElMessage.error('查询参数解析失败: ' + e)
  }

  databaseQueryApi(request).then(response => {
    data.value = response
  }).finally(() => {
    loading.value = false
  })
}

const sortChange = ({order}: any) => {
  params.value.order = order === 'ascending' ? 'asc' : 'desc'
  query()
}

onMounted(() => {
  query()
})
</script>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
}
</style>
