<script setup>
import {onMounted, onUnmounted} from "vue";
import {useGitReleasesStore} from "@/stores/git-releases-store.js";

const store = useGitReleasesStore()
onMounted(() => store.startPolling())
onUnmounted(() => store.stopPolling())
</script>
<template>
  <h2>Changelog/historique de déploiements</h2>
  <ul v-if="store.releases">
    <li v-for="release in store.releases">
      <a :href="release.url">{{ release.repositoryName }} - {{ release.tag }}</a> |
      <em v-if="release.tagName">{{ release.tagName }} | </em>
      <small>{{ release.author }} - {{ release.publishedAt }}</small>
    </li>
  </ul>
  <p v-else>
    No releases recorded yet.
  </p>
</template>

<style scoped>

</style>
