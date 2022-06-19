<template>
  <el-dialog width="500px" draggable v-model="props.visible" :title="props.title" @close="emit('close')">
    <el-form ref="formRef" :model="form" :rules="rules" label-width="120px" label-position="top">
      <el-form-item label="用户名称" prop="username">
        <el-input v-model="form.username" placeholder="请输入用户名称"/>
      </el-form-item>
      <el-form-item label="登录密码" prop="password">
        <el-input type="password" v-model="form.password" placeholder="请输入登录密码"/>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button type="primary" @click="submit()">确定</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import {defineEmits, defineExpose, ref} from 'vue'
import {UserRequest} from '@/api'
import type {FormInstance} from 'element-plus'

const formRef = ref<FormInstance>()
const rules = {
  username: [
    {required: true, message: '请输入索引文本', trigger: 'blur'},
  ],
  password: [
    {required: true, message: '请输入JSON文档', trigger: 'blur'},
  ],
}

interface Props {
  visible: boolean,
  title: string,
  action: string,
}

const props = ref<Props>({visible: false, title: '', action: 'create'})

const emit = defineEmits<{
  (e: 'change'): void,
  (e: 'close'): void,
  (e: 'submit', data: any, props: Props): void,
}>()

const form = ref<UserRequest>({username: '', password: ''})

const submit = async () => {
  await formRef.value?.validate()
  await emit('submit', form.value, props.value)
  props.value = {...props.value, visible: false}
}

const open = (params: Props, values: UserRequest) => {
  props.value = params
  form.value = values
}

defineExpose({open})
</script>
