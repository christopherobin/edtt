/**
 * Systems Controller
 */
angular.module('Dashboard').controller('CreateSystemCtrl', ['$scope', '$modalInstance', 'Systems', CreateSystemCtrl]);

function CreateSystemCtrl($scope, $modalInstance, Systems) {
    $scope.economies = [
        'None',
        'Extraction',
        'Agriculture',
        'Industrial',
        'Service',
        'High-Tech',
        'Refinery',
        'Military',
        'Terraforming',
        'Tourism'
    ];

    $scope.allegiances = [
        'None',
        'Empire',
        'Independant',
        'Federation',
        'Alliance'
    ];

    $scope.system = {
        economy: $scope.economies[0],
        allegiance: $scope.allegiances[0]
    };

    $scope.ok = function () {
        $modalInstance.close($scope.system);
    };

    $scope.cancel = function () {
        $modalInstance.dismiss('cancel');
    };
}