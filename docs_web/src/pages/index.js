import React from 'react';
import clsx from 'clsx';
import Layout from '@theme/Layout';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import styles from './index.module.css';
import './m.css';
import HomepageFeatures from '../components/HomepageFeatures';

function HomepageHeader() {
  const {siteConfig} = useDocusaurusContext();
  const sayAccessControl = true
  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <link rel="preconnect" href="https://fonts.googleapis.com" />
      <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
      <link href="https://fonts.googleapis.com/css2?family=Signika:wght@300;400;500;600;700&display=swap" rel="stylesheet" />
      <div className="container">
        {/* <h1 className="hero__title">{siteConfig.title}</h1>
        <p className="hero__subtitle">{siteConfig.tagline}</p> */}
        <img style={{
          maxHeight: '9rem',
          marginBottom: '1rem'
        }} src='https://github.com/TheSaaSZoo/PermissionPanther/raw/main/docs_web/static/img/g1.png' />
        <h3 style={{
          marginTop: '0.5rem',
          marginBottom: '2rem'
        }}>
          The {sayAccessControl ? "access control": "permissions"} platform for developers who want to spend less time on {sayAccessControl ? "access control": "permissions"}.
        </h3>
        <div className={styles.buttons}>
          <Link
            className="button button--secondary button--lg"
            to="/docs/getting-started/quick-start">
            Master {sayAccessControl ? "Access Control" : "Permissions"} In 3 min ⏱️
          </Link>
        </div>
      </div>
    </header>
  );
}

export default function Home() {
  const {siteConfig} = useDocusaurusContext();
  return (
    <Layout
      title={`${siteConfig.title}`}
      description="Permissions for Killer Apps">
      <HomepageHeader />
      <main>
        <HomepageFeatures />
      </main>
    </Layout>
  );
}
