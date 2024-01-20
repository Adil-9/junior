<h2>Junior</h2>
<h3>Requests</h3>

<strong>"POST"</strong>  - takes json with name, surname and patronymic(not necessary) variables </br>
</br>
example:</br><h5>
{</br>
"name": "Dmitriy",</br>
"surname": "Ushakov",</br>
"patronymic": "Vasilevich" // необязательно</br>
}</br></h5>
</br>
sends request and returns and inserts into database enriched version of data if no problems encoutered</br>
</br>
<strong>"GET"</strong> - returns all instances of data if no variables were provided</br>
</br>
With variables provided return data filtered depending on the variables.</br>
</br>
Variables are: <em>"name:, "surname", "patronymic", "agef", "aget", "gender", "country", "limit", "pagination" </em></br>
</br>
<i>`name, surname, patronymic, gender, country`</i> -filters data returning instances with provided variables</br>
</br>
<i>`agef`</i> and <i>`aget`</i> returns data instances where age => agef and/or age <= agef </br>
</br>
<i>`limit`</i> - limits instances returned, by default 50</br>
</br>
<i>`pagination`</i> - returns Nth chunk of data depending on limit</br>
</br>
<strong>"DELETE"</strong> - deletes instance with provided id </br>
</br>
<strong>"PATCH"</strong> - changes instance of data with provided variables (takes same variables as in "GET" except agef, aget, limit, pagination, instead uses "age".</br>
