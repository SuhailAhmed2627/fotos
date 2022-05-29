import { GET_IMAGES } from "../actions/types";

export default function postReducer(state: any = null, action: any) {
	switch (action.type) {
		case GET_IMAGES:
			return action.payload;
		default:
			return state;
	}
}
