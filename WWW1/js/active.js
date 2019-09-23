/*********************************************************************************

	Template Name: Fixco - Multipurpose Bootstrap4 Template  
	Template URI: https://themeforest.net/user/devitems
	Description: Fixco is aesthetically well organized multipurpose. Comes with 6+ homepages available with multi style that easily modifiable one to another.
	Author: Devitems
	Author URI: https://devitems.com/
	Version: 1.0.1

	Note: This is active js. Plugins activation code here.

**********************************************************************************/


/*===============================================================================
			[ INDEX ]
=================================================================================

	Scroll Up Activation
	Banner Slider Active
	Fake Loader

=================================================================================
			[ END INDEX ]
================================================================================*/

(function ($) {
	'use strict';


	/* Scroll Up Activation */
	$.scrollUp({
		scrollText: '<i class="fa fa-angle-up"></i>',
		easingType: 'linear',
		scrollSpeed: 900,
		animation: 'slide'
	});




	/* Fake Loader */
	$('.fakeloader').fakeLoader({
		timeToHide: 1200,
		bgColor: '#2a68fc',
		spinner: 'spinner2',
	});




	/* Banner Slider Active */
	$('.banner-slider-active').slick({
		autoplay: true,
		autoplaySpeed: 6000,
		fade: true,
		adaptiveHeight: true,
		dots: false,
	});


	/* Banner Slider Active With Navigation */
	$('.banner-slider-active-with-navigation').slick({
		autoplay: false,
		arrows: true,
		prevArrow: '<span class="cr-navigation cr-navigation-prev"><i class="fa fa-angle-left"></i></span>',
		nextArrow: '<span class="cr-navigation cr-navigation-next"><i class="fa fa-angle-right"></i></span>',
		fade: true,
		adaptiveHeight: true,
		dots: false,
	});




	/* Testimonial Slider Active */
	$('.testimonial-slider-active').slick({
		autoplay: true,
		autoplaySpeed: 4000,
		adaptiveHeight: true,
		fade: true,
	});



	/* Blog Slider Active */
	$('.blog-slider-active').slick({
		autoplay: true,
		adaptiveHeight: true,
		autoplaySpeed: 6000,
		slidesToShow: 3,
		centerMode: true,
		arrows: true,
		dots: false,
		prevArrow: '<span class="cr-navigation cr-navigation-prev"><i class="fa fa-angle-left"></i></span>',
		nextArrow: '<span class="cr-navigation cr-navigation-next"><i class="fa fa-angle-right"></i></span>',
		focusOnSelect: true,
		responsive: [{
			breakpoint: 992,
			settings: {
				slidesToShow: 1,
				slidesToScroll: 1
			}
		}]
	});

	/* Testimonial 2 thumbs slider active */
	$('.testimonial2-thumb-wrap').slick({
		slidesToShow: 1,
		slidesToScroll: 1,
		asNavFor: '.testimonial2-content-wrap',
		speed: 2000,
		arrows: true,
		prevArrow: '<span class="cr-navigation cr-navigation-prev"><i class="fa fa-angle-left"></i></span>',
		nextArrow: '<span class="cr-navigation cr-navigation-next"><i class="fa fa-angle-right"></i></span>',
	});


	$('.testimonial2-content-wrap').slick({
		slidesToShow: 1,
		slidesToScroll: 1,
		asNavFor: '.testimonial2-thumb-wrap',
		arrows: false,
		dots: false,
		fade: true,
		speed: 2000,
	});



	/* Testimonial 4 thumbs slider active */
	$('.testimonial4-content-slider-active').slick({
		slidesToShow: 1,
		slidesToScroll: 1,
		asNavFor: '.testimonial4-thumb-slider-active',
		speed: 2000,
		arrows: false,
		fade: true,
		prevArrow: '<span class="cr-navigation cr-navigation-prev"><i class="fa fa-angle-left"></i></span>',
		nextArrow: '<span class="cr-navigation cr-navigation-next"><i class="fa fa-angle-right"></i></span>',
		autoplay: true,
		autoplaySpeed: 6000,
	});


	$('.testimonial4-thumb-slider-active').slick({
		slidesToShow: 1,
		slidesToScroll: 1,
		asNavFor: '.testimonial4-content-slider-active',
		arrows: false,
		dots: false,
		speed: 2000,
		autoplay: true,
		autoplaySpeed: 6000,
	});




	$('.models-thumb-slider-active').slick({
		slidesToShow: 1,
		slidesToScroll: 1,
		asNavFor: '.models-details-slider-active',
		speed: 2000,
		arrows: true,
		fade: true,
		prevArrow: '<span class="cr-navigation cr-navigation-prev"><i class="fa fa-angle-left"></i></span>',
		nextArrow: '<span class="cr-navigation cr-navigation-next"><i class="fa fa-angle-right"></i></span>',
	});

	/* Models Details Slider Active */
	$('.models-details-slider-active').slick({
		slidesToShow: 1,
		slidesToScroll: 1,
		asNavFor: '.models-thumb-slider-active',
		arrows: false,
		dots: false,
		fade: true,
		speed: 2000,
	});





	/* Brand Logo Slider Active */
	$('.brand-logos').slick({
		autoplay: true,
		adaptiveHeight: true,
		autoplaySpeed: 10000,
		slidesToShow: 5,
		arrows: true,
		centerMode: true,
		centerPadding: 0,
		prevArrow: '<span class="cr-navigation cr-navigation-prev"><i class="fa fa-angle-left"></i></span>',
		nextArrow: '<span class="cr-navigation cr-navigation-next"><i class="fa fa-angle-right"></i></span>',
		focusOnSelect: true,
		responsive: [{
				breakpoint: 992,
				settings: {
					slidesToShow: 3,
					slidesToScroll: 1
				}
			},
			{
				breakpoint: 768,
				settings: {
					slidesToShow: 2,
					slidesToScroll: 1
				}
			},
			{
				breakpoint: 576,
				settings: {
					slidesToShow: 1,
					slidesToScroll: 1
				}
			}
		]
	});


	/* Portfolio Details Slider Active */
	$('.pg-portfolio-thumbs').slick({
		slidesToShow: 1,
		slidesToScroll: 1,
		speed: 2000,
		arrows: true,
		prevArrow: '<span class="cr-navigation cr-navigation-prev"><i class="fa fa-long-arrow-left"></i>Prev</span>',
		nextArrow: '<span class="cr-navigation cr-navigation-next">Next<i class="fa fa-long-arrow-right"></i></span>',
	});


	/* Portfolio Details Slider Active */
	$('.op-testimonial-wrap').slick({
		slidesToShow: 1,
		slidesToScroll: 1,
		speed: 2000,
		arrows: false,
		autoplay: true,
		autoplaySpeed: 5000
	});



	/* Portfolio Active */
	var isotopFilter = $('.portfolio-filters');
	var isotopGrid = $('.portfolios:not(.portfolios-slider-active)');
	var isotopGridItemSelector = $('.portfolio-single');
	var isotopGridItem = '.portfolio-single';

	isotopFilter.find('button:first-child').addClass('active');

	//Images Loaded
	isotopGrid.imagesLoaded(function () {
		/*-- init Isotope --*/
		var initial_items = isotopGrid.data('show');
		var next_items = isotopGrid.data('load');
		var loadMoreBtn = $('.load-more-toggle');

		var $grid = isotopGrid.isotope({
			itemSelector: isotopGridItem,
			layoutMode: 'masonry',
		});

		/*-- Isotop Filter Menu --*/
		isotopFilter.on('click', 'button', function () {
			var filterValue = $(this).attr('data-filter');

			isotopFilter.find('button').removeClass('is-checked');
			$(this).addClass('is-checked');

			// use filterFn if matches value
			$grid.isotope({
				filter: filterValue
			});
		});

		/*-- Update Filter Counts --*/
		function updateFilterCounts() {
			// get filtered item elements
			var itemElems = $grid.isotope('getFilteredItemElements');

			if (isotopGridItemSelector.hasClass('hidden')) {
				isotopGridItemSelector.removeClass('hidden');
			}

			var index = 0;

			$(itemElems).each(function () {
				if (index >= initial_items) {
					$(this).addClass('hidden');
				}
				index++;
			});

			$grid.isotope('layout');
		}

		/*-- Function that Show items when page is loaded --*/
		function showNextItems(pagination) {
			var itemsMax = $('.hidden').length;
			var itemsCount = 0;

			$('.hidden').each(function () {
				if (itemsCount < pagination) {
					$(this).removeClass('hidden');
					itemsCount++;
				}
			});

			if (itemsCount >= itemsMax) {
				loadMoreBtn.hide();
			}

			$grid.isotope('layout');
		}

		/*-- Function that hides items when page is loaded --*/
		function hideItems(pagination) {
			var itemsMax = $(isotopGridItem).length;
			var itemsCount = 0;

			$(isotopGridItem).each(function () {
				if (itemsCount >= pagination) {
					$(this).addClass('hidden');
				}
				itemsCount++;
			});

			if (itemsCount < itemsMax || initial_items >= itemsMax) {
				loadMoreBtn.hide();
			}

			$grid.isotope('layout');
		}

		/* Load More Area */
		function loadMoreBtnChk() {
			if (loadMoreBtn.css('display') == 'none') {
				$('.load-more-toggle').parent('.loadmore-wrap').css('display', 'none');
			}
		}

		/*-- Function that Load items when Button is Click --*/
		loadMoreBtn.on('click', function (e) {
			e.preventDefault();
			showNextItems(next_items);
			loadMoreBtnChk();
		});

		hideItems(initial_items);

	});


	$('.portfolios2').imagesLoaded(function () {
		$('.portfolios2').isotope({
			itemSelector: '.portfolio-single',
			layoutMode: 'masonry',
			masonry: {
				columnWidth: 1
			}
		});
	});

	$('.portfolios4').imagesLoaded(function () {
		$('.portfolios4').isotope({
			itemSelector: '.portfolio-single',
			layoutMode: 'masonry',
			masonry: {
				columnWidth: 1,
				gutter: 0,
			}
		});
	});


	// /* Counter Up */
	$('.counter').counterUp({
		delay: 20,
		time: 3000
	});





	/* Popup Activation */
	$('.video-popup-trigger').magnificPopup({
		type: 'iframe',
		mainClass: 'mfp-fade',
		removalDelay: 160,
		preloader: true,
		fixedContentPos: false,
	});



	$('.portfolio-popup-gallery').each(function () {
		$(this).magnificPopup({
			delegate: 'a',
			type: 'image',
			mainClass: 'mfp-fade',
			preloader: true,
			fixedContentPos: false,
			removalDelay: 160,
			gallery: {
				enabled: true
			}
		});
	});




	/* Mobile Menu */
	$('nav.mainmenu').meanmenu({
		meanMenuClose: '<img src="images/icons/icon-close.png" alt="close icon">',
		meanMenuCloseSize: '18px',
		meanScreenWidth: '991',
		meanExpandableChildren: true,
		meanMenuContainer: '.mobile-menu',
		onePage: true
	});



	/* Parallax */
	$('.jarallax').jarallax();



	/* Tilter */
	if ($('.cr-tilter').length) {
		$('.cr-tilter').tilt({
			perspective: 1200,
		});
	}

	if ($('.cr-tilter-glare').length) {
		$('.cr-tilter-glare').tilt({
			glare: true,
			perspective: 1200,
		});
	}

	if ($('.cr-tilter-horaizontal').length) {
		$('.cr-tilter-horaizontal').tilt({
			axis: 'x',
		});
	}




	if ($('.youtube-bg').length) {
		$('.youtube-bg').YTPlayer();
	}





	/* Instafeed active */
	if ($('#sidebar-instagram-feed').length) {

		var userFeed = new Instafeed({
			get: 'user',
			userId: 6665768655,
			accessToken: '6665768655.1677ed0.313e6c96807c45d8900b4f680650dee5',
			target: 'sidebar-instagram-feed',
			resolution: 'thumbnail',
			limit: 6,
			template: '<li><a href="{{link}}" target="_new"> <img src="{{image}}" /><ul class="likes-comments"><li><i class="fa fa-heart-o"></i><span>{{likes}}</span></li><li><i class="fa fa-comments-o"></i><span>{{comments}}</span></li></ul></a></li>',
		});
		userFeed.run();

	}



	/* Niceselect Active */
	$('select').niceSelect();




	/* Sticky Sidebar Active */
	$('.sticky-sidebar-content, .sticky-sidebar-sidebar').theiaStickySidebar({
		additionalMarginTop: 30,
		additionalMarginBottom: 30,
	});


	if ($('#particles-js').length) {
		particlesJS.load('particles-js', 'particles.json', function () {
			console.log('callback - particles.js config loaded');
		});
	}



	if ( $('.onepage-menu').length ) {
		$('.onepage-menu').onePageNav({
			filter: ':not(.external)'
		});
	}



})(jQuery);