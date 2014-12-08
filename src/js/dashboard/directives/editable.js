angular
	.module('Dashboard')
	.directive('editable', ['$compile', editable]);

function editable($compile) {
	var directive = {
        restrict: 'A',
        scope: {
            category: '=category',
            entry: '=entry',
            compare: '=compare'
        },
        link: function(scope, element, attrs) {
            var edit = false;
            var column = attrs["editable"];
            render();

            element.on('click', function() {
                if (edit) return;
                edit = true;
                element.html('<input type="text" class="form-control input-sm">');

                var $input = element.find('input');

                $input.on('focus', function() {
                    setTimeout(function() {
                        $input.select();
                    }, 0);
                });

                $input.val(scope.entry[column]);
                $input.focus();

                function moveTo(direction) {
                    // left 0
                    // up 1
                    // right 2
                    // down 3
                    var tgt = null;
                    switch(direction) {
                    case 0:
                        tgt = element.prev('td[editable]');
                        break;
                    case 2:
                        tgt = element.next('td[editable]');
                        break;
                    case 1:
                        tgt = element.parent().prevAll('tr.entry').first().find('td[editable="' + column + '"]');
                        break;
                    case 3:
                        tgt = element.parent().nextAll('tr.entry').first().find('td[editable="' + column + '"]');
                        break;
                    }

                    if (tgt) {
                        tgt.click();
                    }
                }

                $input.on('keydown', function(e) {
                    // down 40
                    // up 38
                    // left 37
                    // right 39
                    // esc 27
                    switch (e.keyCode) {
                    case 13:
                    case 37:
                    case 38:
                    case 39:
                    case 40:
                        // submit the data
                        edit = false;
                        scope.entry[column] = parseInt($input.val(), 10);
                        render();
                        scope.$emit('editable', scope.entry, column);

                        if (e.keyCode >= 37 && e.keyCode <= 40) {
                            moveTo(e.keyCode - 37);
                        }

                        break;
                    case 27:
                        edit = false;
                        render();
                        break;
                    }
                });
            });

            scope.$watch('entry.' + column, function(newVal) {
                edit = false;
                scope.value = numeral(newVal).format('0,0 CR');
                scope.compareValue = scope.compare(scope.category, scope.entry);
                render();
            });

            function render() {
                scope.extra = null;
                if (column === 'Sell' && scope.compareValue && scope.compareValue.sell) {
                    scope.extra = scope.compareValue.sell;
                    scope.cls = scope.extra > 0 ? 'good' : 'bad';
                }
                if (column === 'Buy' && scope.compareValue && scope.compareValue.buy) {
                    scope.extra = scope.compareValue.buy;
                    scope.cls = scope.extra > 0 ? 'bad' : 'good';
                }
                element.html('{{value}} <span class="{{cls}}" ng-show="extra">({{extra}})</span>');
                $compile(element.contents())(scope)
            }
        }
    };
    return directive;
};