<script setup>
import PlatformToolList from "@/components/PlatformToolList.vue";
import GitRepositoryList from "@/components/GitRepositoryList.vue";
import GitReleaseList from "@/components/GitReleaseList.vue";
</script>

<template>
  <h1>Le Garage</h1>
  <p>Internal Developer Portal</p>
  <article>
    <PlatformToolList/>
  </article>
  <article>
    <GitRepositoryList/>
  </article>
  <article>
    <h2>Métriques</h2>
  </article>
  <article>
    <GitReleaseList/>
  </article>
</template>

