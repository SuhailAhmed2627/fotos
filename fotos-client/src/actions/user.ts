import { User } from "../types";
import { LOGIN, CREATE_EVENT } from "./types";

export const loginSuccess = (data: User | null) => {
	return {
		type: LOGIN,
		payload: data,
	};
};

export const createEvent = (data: any) => {
	return {
		type: CREATE_EVENT,
		payload: data,
	};
};
