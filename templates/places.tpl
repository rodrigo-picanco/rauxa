<html>
<link rel="stylesheet" href="/assets/styles.css">
<script src="https://unpkg.com/htmx.org@2.0.3"></script>
<body>
  <main>
  <h1 class="logo">Rauxa</h1>
  <table>
  <thead>
    <tr>
        <th hx-get="/places?sort=name" hx-swap="innerHTML" hx-target="body">name</th>
        <th hx-get="/places?sort=neighbourhood" hx-swap="innerHTML" hx-target="body">neighbourhood</th>
        <th hx-get="/places?sort=cuisine" hx-swap="innerHTML" hx-target="body">cuisine</th>
    </tr>
  </thead>
  <tbody>
  {{range $index, $place := .places}}
  <tr>
    <td>
       <a href="/places/{{ $index }}">
         {{.Name}}
       </a>
    </td>
    <td>{{.Neighbourhood}}</td>
    <td>{{.Cuisine}}</td>
    </tr>
  {{end}}
  </tbody>
  </table>
  </main>
</html>
