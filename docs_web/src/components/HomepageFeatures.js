import React from 'react';
import clsx from 'clsx';
import styles from './HomepageFeatures.module.css';

const FeatureList = [
  {
    title: 'Permission Groups and Inheritance',
    Svg: require('../../static/img/undraw_docusaurus_mountain.svg').default,
    description: (
      <>
        Repeating yourself sucks. Use <a href="/docs/getting-started/concepts#permission-groups">Permission Groups</a> and <a href="/docs/getting-started/concepts#inheritance-and-recursion">Inheritance</a> to make integration a breeze.
      </>
    ),
  },
  {
    title: 'Focus on What Matters',
    Svg: require('../../static/img/undraw_docusaurus_tree.svg').default,
    description: (
      <>
        Spend less time on permissions, and more time building your awesome app.
      </>
    ),
  },
  {
    title: 'For The Love Of Open Source',
    Svg: require('../../static/img/undraw_docusaurus_react.svg').default,
    description: (
      <>
        Permission Panther is Open Source, because everyone should know how their apps work.
      </>
    ),
  },
];

function Feature({Svg, title, description}) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} alt={title} />
      </div>
      <div className="text--center padding-horiz--md">
        <h3 className={clsx(styles.feature__title)}>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
