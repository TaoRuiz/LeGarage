<script setup>

</script>

<template>
  <h1>Le Garage</h1>
  <p>Internal Developer Portal</p>

  <article>
    <h2>État des services</h2>
    <ul id="services">
<!--      {{range .Services}}-->
<!--      <li>{{if .Healthy}}<span class="service-up">●</span>{{else}}<span class="service-down">●</span>{{end}} {{.Name}}-->
<!--      </li>-->
<!--      {{end}}-->
    </ul>
  </article>
  <article>
    <h2>Activité Git</h2>
<!--    {{range .Repos}}-->
<!--    <div>-->
<!--      <a href="{{.Repo.HTMLURL}}">{{.Repo.FullName}}</a>-->
<!--      {{if .Repo.Description}}<p>{{.Repo.Description}}</p>{{end}}-->
<!--      <ul>-->
<!--        {{range .Commits}}-->
<!--        <li>-->
<!--          <a href="{{.HTMLURL}}">{{slice .SHA 0 7}}</a>-->
<!--          — {{.Commit.Message}}-->
<!--          <small>{{.Commit.Author.Name}} · {{.Commit.Author.Date.Format "2006-01-02"}}</small>-->
<!--        </li>-->
<!--        {{end}}-->
<!--      </ul>-->
<!--    </div>-->
<!--    {{end}}-->

  </article>
  <article>
    <h2>Métriques</h2>
  </article>
  <article>
    <h2>Changelog/historique de déploiements</h2>
<!--    <ul>-->
<!--      {{range .Releases}}-->
<!--      <li>-->
<!--        <a href="{{.HTMLURL}}">{{.Repo}} — {{.TagName}}</a>-->
<!--        {{if .Name}}<em>{{.Name}}</em>{{end}}-->
<!--        <small>{{.Author}} · {{.PublishedAt.Format "2006-01-02"}}</small>-->
<!--      </li>-->
<!--      {{else}}-->
<!--      <li>No releases recorded yet.</li>-->
<!--      {{end}}-->
<!--    </ul>-->
  </article>
</template>

<style scoped>

</style>
