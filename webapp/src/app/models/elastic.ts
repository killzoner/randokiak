export class EsRootDoc {
    constructor(
        public hits: EsRootHits) { }
}

export class EsRootHits {
    constructor(
        public total: {value: number},
        public hits: EsHits[]) { }
}

export class EsHits {
    constructor(
        // tslint:disable-next-line: variable-name
        public _source: any) { }
}
