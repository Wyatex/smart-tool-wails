export namespace main {
	
	export class LoadReturn {
	    err: string;
	    data: string;
	
	    static createFrom(source: any = {}) {
	        return new LoadReturn(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.err = source["err"];
	        this.data = source["data"];
	    }
	}

}

