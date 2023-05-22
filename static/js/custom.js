$(document).ready(function(){

	AOS.init({
		duration: 1200,
	 	once: true,
	 	offset: 100
	});

	$(window).resize(function(){
		AOS.init({
			duration: 1200,
		 	once: true,
		 	offset: 100
		});
	});

	alert('teste');

});

// $('.menu-responsive-btn').click(function(event){
// 	event.preventDefault();
// 	$('.menu-responsive').toggleClass('menu-responsive-open');
// });

// $(document).ready(function(){
// 	if( $(window).width() < 1000 )
// 	{
// 		$('#revolutionSlider li').attr('data-transition', 'fade');
// 		$('#revolutionSlider li').attr('data-masterspeed', 0);
// 	}
// });

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
		}, 3800);
		setTimeout(function(){
			$('.switch.switch-chat-off').addClass('switch-featured');
			$('.switch .switch-container .switch-circle').css('transition', 'left .3s ease-out')
		}, 4800);
	}

});
