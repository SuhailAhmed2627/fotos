export interface Event {
	eventId: number;
	name: string;
	url?: string;
	description?: string;
	users?: {
		username: string;
	}[];
	createdBy?: string;
}

export interface User {
	name: string;
	username: string;
	events: Event[];
	userToken: string;
	firstLogin: boolean;
}

export interface ImageType {
	name: string;
	src: string;
	uploaderName: string;
	eventId?: number;
	width: number;
	height: number;
	clickedAt?: string;
	id?: number;
	faceProcessed?: boolean;
	hasYou?: boolean;
}
