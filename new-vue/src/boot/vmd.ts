// src/boot/vmd.ts
import { boot } from 'quasar/wrappers';
import VueMarkdownEditor  from '@kangc/v-md-editor';
import '@kangc/v-md-editor/lib/style/base-editor.css';
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js';

VueMarkdownEditor.use(vuepressTheme)
export default boot(({ app }) => {
   app.use(VueMarkdownEditor );
});
