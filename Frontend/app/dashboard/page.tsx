"use client";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import styles from "../../styles/Dashboard.module.css";
import Link from "next/link";

export default function Dashboard() {
  const router = useRouter();
  const [user, setUser] = useState<{ firstName: string; email: string } | null>(
    null
  );

  useEffect(() => {
    const storedUser = localStorage.getItem("user");
    if (!storedUser) {
      router.push("/login");
    } else {
      setUser(JSON.parse(storedUser));
    }
  }, [router]);

  return (
    <div className={styles.dashboardContainer}>
      <div className={styles.dashboardBox}>
        <h2 className={styles.welcomeTitle}>
          Welcome, {user?.firstName || "Traveler"}! ✈️
        </h2>
        <p className={styles.welcomeMessage}>
          Manage your flights and connect with travel buddies.
        </p>

        <p className={styles.dashboardDescription}>
          Connect with fellow travelers, manage your flight details, and check
          real-time flight status.
        </p>

        <div className={styles.actions}>
          <Link href="/flightform" className={styles.actionButton}>
            👥 Find Travelers
          </Link>
          <Link href="/flightlist" className={styles.actionButton}>
            ✈ Add Your Flight
          </Link>
          <Link href="/flightstatus" className={styles.actionButton}>
            📡 Track Flight Status
          </Link>
        </div>
      </div>
    </div>
  );
}
