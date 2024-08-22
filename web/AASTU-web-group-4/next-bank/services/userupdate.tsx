import Cookie from "js-cookie";

const API_BASE_URL = "https://bank-dashboard-o9tl.onrender.com";
const token = Cookie.get("accessToken");

// Update User Details - PUT Request
export const updateUserDetails = async (userData: any) => {
  try {
    const response = await fetch(`${API_BASE_URL}/user/update`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(userData),
    });

    if (!response.ok) {
      throw new Error("Failed to update user details");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// Update User Preferences - PUT Request
export const updatePreference = async (userData: any) => {
  try {
    const response = await fetch(`${API_BASE_URL}/user/update-preference`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(userData),
    });

    if (!response.ok) {
      throw new Error("Failed to update user preferences");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// Fetch User Details - GET Request
export const fetchUserDetails = async (username: string) => {
  try {
    const response = await fetch(`${API_BASE_URL}/user/${username}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      throw new Error("Failed to fetch user details");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// Fetch Random Investment Data - GET Request
export const randomInvestmentData = async () => {
  try {
    const response = await fetch(`${API_BASE_URL}/user/random-investment-data`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      throw new Error("Failed to fetch investment data");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// Fetch Current User - GET Request
export const currentUser = async () => {
  try {
    const response = await fetch(`${API_BASE_URL}/user/current`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      throw new Error("Failed to fetch current user details");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};