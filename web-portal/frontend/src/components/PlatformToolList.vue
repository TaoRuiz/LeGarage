<script setup>
import {onMounted, onUnmounted} from "vue";
import {usePlatformToolsStore} from "@/stores/platform-tools-store.js";

const store = usePlatformToolsStore()
onMounted(() => store.startPolling())
onUnmounted(() => store.stopPolling())
</script>

<template>
  <h2>État des services</h2>
  <ul id="services">
    <li v-for="tool in store.tools">
      <span :class="tool.healthy? 'service-up' : 'service-down'">●</span> {{ tool.name }}
    </li>
  </ul>
</template>

<style scoped>

</style>
