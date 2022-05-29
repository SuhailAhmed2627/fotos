import { CREATE_EVENT, LOGIN } from "../actions/types";

export default function postReducer(state: any = null, action: any) {
	switch (action.type) {
		case LOGIN:
			return action.payload;
		default:
			return state;
	}
}
