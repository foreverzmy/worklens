export interface StreamMessage {
	type: string;
	data: any;
}

export interface RepoReference {
	name: string;
	shortName: string;
	hash: string;
	category: 'branch' | 'tag' | 'remote' | 'stash' | 'other';
	type: string;
}
