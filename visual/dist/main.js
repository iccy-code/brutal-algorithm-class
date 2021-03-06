// @ts-nocheck
import { chan, } from "https://creatcodebuild.github.io/csp/dist/csp.js";
import * as csp from "https://creatcodebuild.github.io/csp/dist/csp.js";
import * as sort from "./sort.js";
import * as component from "./components.js";
async function main() {
    DefineComponent();
    // init an array
    let array = [];
    for (let i = 0; i < 50; i++) {
        array.push(Math.floor(Math.random() * 50));
    }
    // event channels
    let insertQueue = chan();
    let mergeQueue = chan();
    let stop = chan();
    let resume = chan();
    let resetChannel = chan();
    // @ts-ignore
    let onReset = csp.multi(resetChannel);
    let mergeQueue2 = (() => {
        let c = chan();
        (async () => {
            let numebrsToRender = [].concat(array);
            await c.put(numebrsToRender);
            while (1) {
                let [numbers, startIndex] = await mergeQueue.pop();
                // console.log(numbers);
                for (let i = 0; i < numbers.length; i++) {
                    numebrsToRender[i + startIndex] = numbers[i];
                }
                await c.put(numebrsToRender);
            }
        })();
        return c;
    })();
    // console.log(mergeQueue2);
    // Components
    let onLanguageChange = component.languages("languages");
    component.SortVisualizationComponent("insertion-sort", insertQueue, onLanguageChange.copy());
    component.SortVisualizationComponent("merge-sort", mergeQueue, onLanguageChange.copy());
    component.DataSourceComponent("data-source-1", array, resetChannel, onLanguageChange);
    // Kick off
    console.log("begin sort", array);
    // @ts-ignore
    sort.Sorter2(sort.InsertionSort, onReset.copy(), insertQueue);
    // @ts-ignore
    sort.Sorter2(sort.MergeSort, onReset.copy(), mergeQueue);
}
main();
function DefineComponent() {
    // Web Components
    customElements.define("sort-visualization", class extends HTMLElement {
        constructor() {
            super();
            let template = document.getElementById("sort-visualization");
            let templateContent = template.content;
            const shadowRoot = this.attachShadow({ mode: "open" })
                .appendChild(templateContent.cloneNode(true));
        }
    });
    customElements.define("data-source", class extends HTMLElement {
        constructor() {
            super();
            let template = document.getElementById("data-source");
            let templateContent = template.content;
            const shadowRoot = this.attachShadow({ mode: "open" })
                .appendChild(templateContent.cloneNode(true));
        }
    });
}
//# sourceMappingURL=main.js.map