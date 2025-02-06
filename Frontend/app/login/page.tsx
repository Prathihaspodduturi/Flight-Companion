"use client";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { FaEnvelope, FaLock } from "react-icons/fa";
import Link from "next/link";
import styles from "../../styles/Login.module.css";

export default function Login() {
  const router = useRouter();

  const [formData, setFormData] = useState({ email: "", password: "" });
  const [errorMessage, setErrorMessage] = useState("");
  const [successMessage, setSuccessMessage] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  useEffect(() => {
    const userSession = localStorage.getItem("user");
    if (userSession) {
      router.push("/dashboard"); // Redirect if already logged in
    }
  }, [router]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    setErrorMessage("");
    setSuccessMessage("");
    setIsSubmitting(true);

    setTimeout(() => {
      const storedUser = localStorage.getItem("registeredUser");
      if (!storedUser) {
        setErrorMessage("❌ No registered user found. Please sign up first.");
        setIsSubmitting(false);
        return;
      }

      const registeredUser = JSON.parse(storedUser);
      if (
        formData.email === registeredUser.email &&
        formData.password === registeredUser.password
      ) {
        setSuccessMessage("✅ Login Successful! Redirecting...");
        localStorage.setItem("user", JSON.stringify(registeredUser)); // Store user session
        setTimeout(() => router.push("/dashboard"), 2000);
      } else {
        setErrorMessage("❌ Invalid email or password. Please try again.");
        setIsSubmitting(false);
      }
    }, 1500);
  };

  return (
    <div className={styles.loginContainer}>
      <div className={styles.loginBox}>
        <h2 className={styles.loginTitle}>Login</h2>

        {errorMessage && <p className={styles.errorMessage}>{errorMessage}</p>}
        {successMessage && (
          <p className={styles.successMessage}>{successMessage}</p>
        )}

        <form onSubmit={handleSubmit}>
          <div className={styles.inputGroup}>
            <FaEnvelope className={styles.inputIcon} />
            <input
              type="email"
              name="email"
              placeholder="Email"
              className={styles.inputField}
              onChange={handleChange}
              required
            />
          </div>

          <div className={styles.inputGroup}>
            <FaLock className={styles.inputIcon} />
            <input
              type="password"
              name="password"
              placeholder="Password"
              className={styles.inputField}
              onChange={handleChange}
              required
            />
          </div>

          <button
            type="submit"
            className={styles.submitButton}
            disabled={isSubmitting}
          >
            {isSubmitting ? "Logging in..." : "Login"}
          </button>
        </form>

        {/* 🔹 Signup Link Below */}
        <p className={styles.registerRedirect}>
          Don&apos;t have an account?{" "}
          <Link href="/register" className={styles.registerLink}>
            Sign up
          </Link>
        </p>
      </div>
    </div>
  );
}
