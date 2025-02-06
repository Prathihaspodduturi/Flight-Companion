"use client";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import {
  FaUser,
  FaEnvelope,
  FaGlobe,
  FaEdit,
  FaSignOutAlt,
} from "react-icons/fa";
import styles from "../../styles/Profile.module.css";

// Define the User Type
interface User {
  firstName: string;
  lastName: string;
  email: string;
  country?: string;
}

export default function Profile() {
  const router = useRouter();
  const [user, setUser] = useState<User | null>(null);
  const [isEditing, setIsEditing] = useState(false);
  const [updatedUser, setUpdatedUser] = useState<User | null>(null);

  useEffect(() => {
    const storedUser = localStorage.getItem("user");
    if (storedUser) {
      const parsedUser: User = JSON.parse(storedUser);
      setUser(parsedUser);
      setUpdatedUser(parsedUser);
    } else {
      router.push("/login"); // Redirect if no user is found
    }
  }, []);

  const handleEdit = () => {
    setIsEditing(true);
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (updatedUser) {
      setUpdatedUser({ ...updatedUser, [e.target.name]: e.target.value });
    }
  };

  const handleSave = () => {
    if (updatedUser) {
      localStorage.setItem("user", JSON.stringify(updatedUser));
      setUser(updatedUser);
      setIsEditing(false);
    }
  };

  const handleLogout = () => {
    localStorage.removeItem("user");
    router.push("/login");
  };

  return (
    <div className={styles.profileContainer}>
      <div className={styles.profileCard}>
        <h2 className={styles.profileTitle}>Profile</h2>

        {user && (
          <>
            <div className={styles.userInfo}>
              <FaUser className={styles.icon} />
              {isEditing ? (
                <input
                  type="text"
                  name="firstName"
                  value={updatedUser?.firstName || ""}
                  onChange={handleChange}
                  className={styles.inputField}
                />
              ) : (
                <p>
                  {user.firstName} {user.lastName}
                </p>
              )}
            </div>

            <div className={styles.userInfo}>
              <FaEnvelope className={styles.icon} />
              <p>{user.email}</p>
            </div>

            <div className={styles.userInfo}>
              <FaGlobe className={styles.icon} />
              {isEditing ? (
                <input
                  type="text"
                  name="country"
                  value={updatedUser?.country || ""}
                  onChange={handleChange}
                  className={styles.inputField}
                />
              ) : (
                <p>{user.country || "Not Provided"}</p>
              )}
            </div>

            <div className={styles.buttonGroup}>
              {isEditing ? (
                <button onClick={handleSave} className={styles.saveButton}>
                  Save
                </button>
              ) : (
                <button onClick={handleEdit} className={styles.editButton}>
                  <FaEdit /> Edit Profile
                </button>
              )}
              <button onClick={handleLogout} className={styles.logoutButton}>
                <FaSignOutAlt /> Logout
              </button>
            </div>
          </>
        )}
      </div>
    </div>
  );
}
