<script setup>
import {onMounted, onUnmounted} from "vue";
import {useGitRepositoriesStore} from "@/stores/git-repositories-store.js";

const store = useGitRepositoriesStore()
onMounted(() => store.startPolling())
onUnmounted(() => store.stopPolling())
</script>

<template>
  <h2>Activité Git</h2>
  <div v-for="repository in store.repositories">
    <a href="{{repository.url}}">{{ repository.name }}</a>
    <p v-if="repository.description">{{ repository.description }}</p>
    <ul>
      <li v-for="commit in repository.commits">
        <a :href="commit.url">{{ commit.name }}</a> {{ commit.message }}
        <small>{{ commit.author }} · {{ commit.date }}</small>
      </li>
    </ul>
  </div>
</template>

<style scoped>

</style>
