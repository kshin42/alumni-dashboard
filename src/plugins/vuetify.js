import Vue from 'vue';
import Vuetify from 'vuetify/lib';

Vue.use(Vuetify);

export default new Vuetify({
    theme: {
        themes: {
          light: {
            primary: "#21295C",
            secondary: "#4281A4",
            tertiary: "#84BCDA",
            accent: "#48A9A6",
            error: "#EA526F",
          },
        },
      },
});
