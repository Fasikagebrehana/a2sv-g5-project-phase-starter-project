import { call, put, takeLatest } from 'redux-saga/effects';
import axios from 'axios';
import { setUser } from './slice/userSlice';

interface FetchUserAction {
  type: string;
  payload: {
    username: string;
    token: string;
  };
}

// Function to fetch user data from the API
function fetchUserData(username: string, token: string) {
  return axios.get(`https://bank-dashboard-rsf1.onrender.com/user/${username}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
}

// Saga to handle the user fetch request
function* fetchUser(action: FetchUserAction) {
  try {
    const response = yield call(fetchUserData, action.payload.username, action.payload.token);
    const data = response.data;
    yield put(setUser(data));
  } catch (e) {
    console.error('Failed to fetch user data', e);
  }
}

function* userSaga() {
  yield takeLatest('USER_FETCH_REQUESTED', fetchUser);
}

export default userSaga;
