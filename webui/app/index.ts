import Vue from 'vue';
import Component from 'vue-class-component';

const template = require('./index.html');

@Component({
    name: "app",
    template: template
})
export default class App extends Vue {
    // Initial data can be declared as instance properties
    message: string = 'Hello!';

    // Component methods can be declared as instance methods
    onClick (): void {
        window.alert(this.message)
    }
}


// export default Vue.extend({
//     name: 'app',
//     template: templateHtml,
//     data: () => {
//         return {msg: 'This is page 1'}
//     },
//     // components: {
//     //     PageHeader,
//     //     PageFooter
//     // }
// });
