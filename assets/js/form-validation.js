(function ($, undefined) {
    $(document).ready(function () {
        var requiredText = "This field is required";
        var emailText = "Please enter valid email";
        var $form = $('#createBlog');
        $form.validate({
                    rules: {
                        'title': {
                            required: true
                        },
                        'description': {
                            required: true
                        },
                    },
                    messages: {
                        'title': {
                            required: requiredText
                        },
                        'description': {
                            required: requiredText
                        },
                    }
                }
        );

        var $form = $('#createComment');
        $form.validate({
                    rules: {
                        'comment': {
                            required: true
                        },
                    },
                    messages: {
                        'comment': {
                            required: requiredText
                        },
                    }
                }
        );

        var $form = $('#registration');
        $form.validate({
                    rules: {
                        'username': {
                            required: true,
                        },
                        'password': {
                            required: true,
                        },
                        'email': {
                            required: true,
                            email: true
                        },

                    },
                    messages: {
                        'username': {
                            required: requiredText
                        },
                        'password': {
                            required: requiredText
                        },
                        'email': {
                            required: requiredText,
                            email: emailText
                        }
                    }
                }
        );

        var $form = $('#contact');
        $form.validate({
                    rules: {
                        'name': {
                            required: true,
                        },
                        'email': {
                            required: true,
                            email: true
                        },
                        'message': {
                            required: true,
                        },
                    },
                    messages: {
                        'name': {
                            required: requiredText
                        },
                        'email': {
                            required: requiredText,
                            email: emailText
                        },
                        'message': {
                            required: requiredText
                        },
                    }
                });
    })
})(jQuery);