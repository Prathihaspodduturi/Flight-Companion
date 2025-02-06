"use client";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import {
  validateEmail,
  validatePassword,
  validateName,
} from "../utils/validators";
import {
  FaUser,
  FaEnvelope,
  FaLock,
  FaCheckCircle,
  FaTimesCircle,
} from "react-icons/fa";
import styles from "../../styles/Register.module.css";
import Link from "next/link";

export default function Register() {
  const router = useRouter();

  const [formData, setFormData] = useState({
    firstName: "",
    lastName: "",
    email: "",
    password: "",
  });

  const [errors, setErrors] = useState({
    firstName: "First name is required",
    lastName: "Last name is required",
    email: "Email is required",
    password: "Password must meet the criteria.",
  });

  const [successMessage, setSuccessMessage] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  useEffect(() => {
    const userSession = localStorage.getItem("user");
    if (userSession) {
      router.push("/dashboard"); // Redirect if already logged in
    }
  }, []);

  // Password validation checks
  const passwordChecks = {
    length: formData.password.length >= 12,
    uppercase: /[A-Z]/.test(formData.password),
    lowercase: /[a-z]/.test(formData.password),
    number: /\d/.test(formData.password),
    specialChar: /[@$!%*?&]/.test(formData.password),
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });

    const newErrors = { ...errors };

    if (name === "firstName" || name === "lastName") {
      newErrors[name] = validateName(value)
        ? ""
        : "No numbers or special characters allowed";
    }

    if (name === "email") {
      newErrors.email = validateEmail(value) ? "" : "Invalid email format";
    }

    if (name === "password") {
      newErrors.password = validatePassword(value)
        ? ""
        : "Password must meet the criteria below.";
    }

    setErrors(newErrors);
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (Object.values(errors).some((error) => error !== "")) {
      alert("Please fix the errors before submitting.");
      return;
    }

    setIsSubmitting(true);

    setTimeout(() => {
      localStorage.setItem("registeredUser", JSON.stringify(formData)); // Store user details
      setSuccessMessage("🎉 Registration successful! Redirecting to login...");
      setTimeout(() => router.push("/login"), 2000);
    }, 1000);
  };

  return (
    <div className={styles.registerContainer}>
      <div className={styles.registerBox}>
        <h2 className={styles.registerTitle}>Create an Account</h2>

        {successMessage && (
          <p className={styles.successMessage}>{successMessage}</p>
        )}

        <form onSubmit={handleSubmit}>
          <div className={styles.inputGroup}>
            <FaUser className={styles.inputIcon} />
            <input
              type="text"
              name="firstName"
              placeholder="First Name"
              className={styles.inputField}
              onChange={handleChange}
              required
            />
          </div>
          {errors.firstName && (
            <p className={styles.errorText}>{errors.firstName}</p>
          )}

          <div className={styles.inputGroup}>
            <FaUser className={styles.inputIcon} />
            <input
              type="text"
              name="lastName"
              placeholder="Last Name"
              className={styles.inputField}
              onChange={handleChange}
              required
            />
          </div>
          {errors.lastName && (
            <p className={styles.errorText}>{errors.lastName}</p>
          )}

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
          {errors.email && <p className={styles.errorText}>{errors.email}</p>}

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
          {errors.password && (
            <p className={styles.errorText}>{errors.password}</p>
          )}

          <ul className={styles.passwordChecklist}>
            {Object.entries(passwordChecks).map(([key, value]) => (
              <li key={key} className={value ? styles.valid : styles.invalid}>
                {value ? (
                  <FaCheckCircle className={styles.checkIcon} />
                ) : (
                  <FaTimesCircle className={styles.crossIcon} />
                )}
                {key === "length" && "At least 12 characters"}
                {key === "uppercase" && "At least one uppercase letter"}
                {key === "lowercase" && "At least one lowercase letter"}
                {key === "number" && "At least one number"}
                {key === "specialChar" &&
                  "At least one special character (@, $, !, %, *, ?, &)"}
              </li>
            ))}
          </ul>

          <button
            type="submit"
            className={styles.submitButton}
            disabled={isSubmitting}
          >
            {isSubmitting ? "Processing..." : "Sign Up"}
          </button>
        </form>

        <p className={styles.loginRedirect}>
          Already have an account? <Link href="/login">Login</Link>
        </p>
      </div>
    </div>
  );
}
