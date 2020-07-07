# alumni-dashboard

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Run your tests
```
npm run test
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

https://coolors.co/2b2d42-8d99ae-edf2f4-ef233c-d90429
https://coolors.co/503280-6e6d70-c0b283-dcd0c0-000000
https://www.canva.com/colors/color-palette-generator/
fab fa-connectdevelop

Storage Layer

composite primary key between email and orgCode?
avoid scan because it looks at the entire table
create secondary index on orgCode with email being sort key now
give your primary key and secondary key generic names since not all of your items will have the same data you want as a primary key
   - you can prefix the keys with the type of entity you are working with
one to many relationship
   - for limited number of many items to the one you can denormalize(just stick json in a field thats attached to the user entity)
   - for limitess many items you can make the sort key the "id" of the item
       PK = User#user1 and begins_with(sk, "order#")
filtering
   - build it into your keys, pick a partition first like a user then filter from there
   - create an attribute thats the combo of two different attributes. indexes must be unqiue so this is a good way to create that (SHIPPED#32432)


data fields:

users (emails)
   - resume
   - various attributes
orgs

access patterns

get users profile
get members of org

