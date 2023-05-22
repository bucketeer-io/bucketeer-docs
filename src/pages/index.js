import React from 'react';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import useBaseUrl from '@docusaurus/useBaseUrl';
import Layout from '@theme/Layout';
import Intro from '../theme/sections/intro/intro';
import Features from '../theme/sections/features/features';



import jquery from 'jquery';
if(typeof window !== 'undefined'){
	window.$ = window.jQuery=jquery;
}
import AOS from 'aos';
// import "aos/dist/aos.css";

import 'font-awesome/css/font-awesome.min.css'
require('bootstrap/dist/css/bootstrap.css');
require('/static/css/aos.css');
require('/static/css/custom.css');

export default function Home() {

	if(typeof window !== 'undefined')
	{

		function checkhome()
		{
			var checkhome_timer = setInterval(() => {
				if( window.location.pathname != "/" )
				{
					clearInterval(checkhome_timer);
					$(".navbar").removeClass("home-navbar");
					$(".main-wrapper").removeClass("main-wrapper-home");
					checkhomeback();
				}
			}, 100);
		}

		function checkhomeback() 
		{
			var checkhomeback_timer = setInterval(() => {
				if( window.location.pathname == "/" )
				{
					clearInterval(checkhomeback_timer);
					$(".navbar").addClass("home-navbar");
					$(".main-wrapper").addClass("main-wrapper-home");
					checkhome();
				}
			}, 100);
		}

		// $(".navbar").addClass("home-navbar");
		setTimeout(function(){
			$(".main-wrapper").addClass("main-wrapper-home");
		}, 0);

		checkhome();
		$(".navbar").addClass("internal-navbar");

		AOS.init({
			duration: 1200,
		 	once: true,
		 	offset: 50
		});
		
		$(window).resize(function(){
			AOS.init({
				duration: 1200,
			 	once: true,
			 	offset: 50
			});
		});
		
		$(document).ready(function(){
			$('.navbar').addClass('home-navbar');
		});

		$(window).scroll(function(){
			if( !$(".navbar").hasClass("home-navbar-active") )
			{
				if( $(window).scrollTop() > 55 )
				{
					$(".navbar").addClass("home-navbar-active");
				}
				else
				{
					$(".navbar").removeClass("home-navbar-active");
				}
			}
			else
			{
				if( $(window).scrollTop() <= 55 )
				{
					$(".navbar").removeClass("home-navbar-active");
				}
			}
		});

		$(document).ready(function(){

			$('.switch').click(function(){
				if( !$(this).hasClass('switch-enable') )
				{
					$(this).addClass('switch-enable');

					if( $(this).hasClass('switch-featured') )	
					{
						$(this).removeClass('switch-featured');
					}

					if( $(this).hasClass('switch-sidebar-off') )
					{
						$('.tp-caption.service .service-top').addClass('service-top-disable');
						$('.tp-caption.service .service-sidebar').removeClass('service-sidebar-disable');
						$('.tp-caption.service .service-content').removeClass('service-content-adjust');
						$(this).removeClass('switch-sidebar-off');
						$(this).addClass('switch-sidebar-on');
					}

					if( $(this).hasClass('switch-sidebar-off') )
					{
						$('.tp-caption.service .service-top').addClass('service-top-disable');
						$('.tp-caption.service .service-sidebar').removeClass('service-sidebar-disable');
						$('.tp-caption.service .service-content').removeClass('service-content-adjust');
						$(this).removeClass('switch-sidebar-off');
						$(this).addClass('switch-sidebar-on');
					}

					if( $(this).hasClass('switch-chat-off') )
					{
						$('.tp-caption.service .service-chat').removeClass('service-chat-disable');
						$(this).removeClass('switch-chat-off');
						$(this).addClass('switch-chat-on');
					}
				}
				else
				{
					$(this).removeClass('switch-enable');

					if( $(this).hasClass('switch-sidebar-on') )
					{
						$('.tp-caption.service .service-top').removeClass('service-top-disable');
						$('.tp-caption.service .service-sidebar').addClass('service-sidebar-disable');
						$('.tp-caption.service .service-content').addClass('service-content-adjust');
						$(this).removeClass('switch-sidebar-on');
						$(this).addClass('switch-sidebar-off');
					}

					if( $(this).hasClass('switch-chat-on') )
					{
						$('.tp-caption.service .service-chat').addClass('service-chat-disable');
						$(this).removeClass('switch-chat-on');
						$(this).addClass('switch-chat-off');
					}
				}
			});

		});

		$(document).ready(function(){

			if( $('section').hasClass('slider-home') )
			{
				setTimeout(function(){
					$('.switch .switch-container .switch-circle').css('transition', 'left .6s ease-out')
					$('.switch.switch-sidebar-off').click();
				}, 2600);
				setTimeout(function(){
					$('.switch.switch-chat-off').addClass('switch-featured');
					$('.switch .switch-container .switch-circle').css('transition', 'left .3s ease-out')
				}, 3600);
			}

		});

	}

	const context = useDocusaurusContext();
	const {siteConfig = {}} = context;

	return (
		<Layout
			description={siteConfig.tagline}
			keywords={siteConfig.customFields.keywords}
			metaImage={useBaseUrl(`img/${siteConfig.customFields.image}`)}
			wrapperClassName={'page-home'}
		>
			<Intro />
			<Features />
		</Layout>
	);
}