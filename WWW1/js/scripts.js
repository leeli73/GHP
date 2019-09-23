/*********************************************************************************

    Template Name: Fixco - Multipurpose Bootstrap4 Template  
    Template URI: https://themeforest.net/user/devitems
    Description: Fixco is aesthetically well organized multipurpose. Comes with 6+ homepages available with multi style that easily modifiable one to another.
    Author: Devitems
    Author URI: https://devitems.com/
    Version: 1.0.1

    Note: This is scripts js. All custom scripts here.

**********************************************************************************/

/*===============================================================================
			[ INDEX ]
=================================================================================

	Write here

=================================================================================
			[ END INDEX ]
================================================================================*/

(function ($) {
    'use strict';

    /* Scroll Scripts */
    $(window).on('scroll', function () {

        function winScrollPosition() {
            var scrollPos = $(window).scrollTop(),
                winHeight = $(window).height();
            var scrollPosition = Math.round(scrollPos + (winHeight / 1.2));
            return scrollPosition;
        }


        function progressAddClassOnScroll() {
            var trigger = $('.progress-bar');

            if (trigger.length) {
                var triggerPos = Math.round(trigger.offset().top);

                if (triggerPos < winScrollPosition()) {
                    trigger.each(function () {
                        $(this).addClass('fill');
                    });
                }
            }

        }
        progressAddClassOnScroll();


        var scrollPos = $(this).scrollTop();
        if (scrollPos > 300) {
            $('.sticky-header').addClass('is-sticky');
        } else {
            $('.sticky-header').removeClass('is-sticky');
        }

    });




    /* Expandable Searchbox */
    function expandableSearchbox() {
        var trigger = $('.search-trigger'),
            container = $('.expandable-searchbox');

        trigger.on('click', function () {
            $(this).find('i.fa').toggleClass('fa-close');
            container.toggleClass('is-visible');
        });

        container.on('focus', 'input', function () {
            $(this).parents('form').addClass('focused');
        });
        container.on('focusout', 'input', function () {
            $(this).parents('form').removeClass('focused');
            var $this = $(this);
            if ($this.val().trim().length !== 0) {
                $(this).parents('form').addClass('focused');
            }
        });

        $('<button class="close">close</button>').appendTo(container)
            .on('click', function () {
                container.removeClass('is-visible');
                trigger.find('i.fa').removeClass('fa-close');
            });

    }
    expandableSearchbox();



    /* Expandable Menu */
    function expandableMenu() {
        var trigger = $('.slidemenu-trigger');
        var container = $('.mainmenu');
        trigger.on('click', function () {
            container.toggleClass('menu-expanded');
            $(this).find('i').toggleClass('fa-bars');
            $(this).find('i').toggleClass('fa-close');
        });
    }
    expandableMenu();


    /* Sticky List Blog */
    function stickyListBlog() {
        $('.list-blog').each(function () {
            if ($(this).hasClass('sticky')) {
                $('<span class="sticky-icon"><i class="bi bi-pin"></i></span>').prependTo($(this).find('.grid-blog-header'));
            }
        });
    }
    stickyListBlog();


    /* Sticky Grid Blog */
    function stickyGridBlog() {
        $('.grid-blog:not(.list-blog)').each(function () {
            if ($(this).hasClass('sticky')) {
                $('<span class="sticky-icon"><i class="bi bi-pin"></i></span>').prependTo($(this).find('.grid-blog-thumb'));
            }
        });
    }
    stickyGridBlog();


    function menuLastElement() {
        $('.mainmenu > ul > li').slice(3).addClass('last-element');
    }
    menuLastElement();



    /* Sticky Header */




})(jQuery);