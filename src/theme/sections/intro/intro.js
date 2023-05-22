
import React from 'react';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import useBaseUrl from '@docusaurus/useBaseUrl';

export default function Intro() {

	const context = useDocusaurusContext();
	const {siteConfig = {}} = context;

	const component = 'shift-intro';

	return (
		<section className="slider-home">
			<div className="container">
				<div className="tp-caption block-left" data-aos="fade-down" data-aos-delay="0" data-aos-duration="1000">
					<img src="img/slider/img-block_left.png" />
				</div>

				<div className="tp-caption block-right" data-aos="fade-up" data-aos-delay="0">
					<img src="img/slider/img-block_right.png" />
				</div>

				<div className="tp-caption main-label" data-aos="fade-up-small" data-aos-delay="900" data-aos-duration="1000">Maximize Development with<br /> Minimum Risk</div>

				<div className="tp-caption bottom-label" data-aos="fade-up-small" data-aos-delay="1100" data-aos-duration="1000">Bucketeer is an open-source platform created to help teams make better decisions,<br />reduce deployment lead time and release risk through feature flags.</div>

				<div className="tp-caption dashboard d-@none d-lg-inline-block" data-aos="fade-up-small" data-aos-delay="1000">
					<img src="img/slider/img_title-top_bucketeer-dashboard.png" />

					<div className="tp-caption dashboard-list d-none d-lg-inline-block">
						<img src="img/slider/img_title-top_bucketeer_list_slide-off_chat-off.png" />
					</div>

					<div className="tp-caption dashboard-switch-1">
						<span className="title d-inline-block d-lg-none">Sidebar</span>
						<div className="switch switch-sidebar-off"><div className="switch-container"><div className="switch-circle"></div></div></div>
					</div>

					<div className="tp-caption dashboard-switch-2">
						<span className="title d-inline-block d-lg-none" style={{position: 'relative',left: '494px',top: '-93px',fontSize: '17px'}}>Chat</span>
						<div className="switch switch-chat-off"><div className="switch-container"><div className="switch-circle"></div></div></div>
					</div>
				</div>


				<div className="tp-caption arrow d-none d-lg-inline-block" data-aos="fade-up-small" data-aos-delay="1100">
					<img src="img/slider/img_arrow.png" />
				</div>

				<div className="tp-caption service">
					<div className="service-screen" data-aos="fade-up-small" data-aos-delay="1200">
						<img src="img/slider/service-top.png" className="service-top d-block" />
						<img src="img/slider/service-sidebar.png" className="service-sidebar service-sidebar-disable d-inline-block" />
						<img src="img/slider/service-content.png" className="service-content service-content-adjust d-inline-block" />
						<img src="img/slider/service-chat.png" className="service-chat service-chat-disable" />
					</div>
				</div>
			</div>
		</section>
	);
}