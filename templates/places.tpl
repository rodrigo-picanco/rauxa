<html>
  <head>
    <link rel="stylesheet" href="/assets/styles.css">
    <script src="https://unpkg.com/htmx.org@2.0.3"></script>
  </head>
  <body>
    <main>
      <h1 class="logo">Rauxa</h1>
      <table class="content">
        <thead>
          <tr>
            <th hx-get="/?sort=name" hx-swap="innerHTML" hx-target="body" class="no-select">name&ShortDownArrow;</th>
            <th hx-get="/?sort=neighbourhood" hx-swap="innerHTML" hx-target="body" class="no-select">neighbourhood&ShortDownArrow;</th>
            <th hx-get="/?sort=cuisine" hx-swap="innerHTML" hx-target="body" class="no-select">cuisine&ShortDownArrow;</th>
          </tr>
        </thead>
        <tbody>
          {{range $index, $place := .places}}
            <tr>
              <td>
                <a href="/{{ $index }}">
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
  </body>
</html>
