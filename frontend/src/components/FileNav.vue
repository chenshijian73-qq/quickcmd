<template>
  <a-space direction="vertical" :style="{ width: '200px', borderRadius: '4px', backgroundColor: 'transparent'}">
    <a-input-search :style="{width:'99%'}" placeholder="Search ..." @search="handleSearch" @input="handleSearch" search-button/>
    <a-menu mode="pop">
      <a-menu-item v-for="(file, index) in searchFiles" :key="index" :style="{ backgroundColor: props.selectedFile.Name === file.Name ? '#d9e5f6' : 'transparent' }" @click="props.selectFile(file)">
        <div class="file-info">
          <a-switch v-model="file.Enabled" size="small" @change="changeStatus(file)"/>
          <span class="name">{{ file.Name.slice(0, file.Name.lastIndexOf('.')) }}</span>
          <span class="menu-trigger">
              <a-trigger
                  :trigger="['hover']"
                  clickToClose
                  position="right"
                  @show="file.fileMenuVisible=true"
                  @hide="file.fileMenuVisible=false"
              >
                <div :class="`button-trigger ${file.fileMenuVisible ? 'button-trigger-active' : ''}`">
                  <IconDoubleRight class="icon-double-right " />
                </div>
                <template #content>
                  <a-menu
                      :style="{ marginBottom: '-4px' }"
                      mode="popButton"
                      :tooltipProps="{ position: 'right' }"
                      showCollapseButton
                  >
                    <a-menu-item key="index" class="red-icon" @click="removeFile(file)">
                      <template #icon><IconDelete class="red-icon"/></template>
                      Remove
                    </a-menu-item>
                  </a-menu>
                </template>
              </a-trigger>
            </span>
        </div>
      </a-menu-item>
    </a-menu>
  </a-space>
</template>

<script setup>
import '../assets/nav.css'
import {ref, watch} from 'vue';
import { EditFile } from '../../wailsjs';
import {message} from "ant-design-vue";
import {
  IconDoubleRight,
  IconDelete,
} from '@arco-design/web-vue/es/icon';

const props = defineProps({
  qcFiles: {
    type: Array,
    required: true,
  },
  selectedFile: {
    type: Object,
    required: true,
  },
  selectFile: {
    type: Function,
    required: true,
  },
  removeFile: {
    type: Function,
    required: true,
  }
});

const files = ref([])

function changeStatus(file) {
  EditFile(file).then(err => {
    if (err !== ""){
      message.error(`修改文件 ${file.Name} 状态失败: ${err}`)
    }
  })
}

function removeFile(file) {
  props.removeFile(file)
}

const searchFiles = ref([])
function handleSearch(value) {
  searchFiles.value = props.qcFiles.filter(file => file.Name?.includes(value))
}

watch(() => props.qcFiles, (value) => {
  if (value) {
    searchFiles.value = value;
  } else {
    searchFiles.value = [];
  }
});

</script>

<style scoped>
</style>