angular.module('Dashboard').factory('Systems', ['$http', SystemsService]);

function SystemsService($http) {
	return {
		list: function(success, error) {
			error = error || Function();
			$http.get('/api/systems')
				.success(function(systems) {
					systems.sort(function(a, b) {
						return a.name.localeCompare(b.name);
					});
					success(systems);
				})
				.error(error);
		},
		system: function(name, success, error) {
			error = error || Function();
			$http.get('/api/systems/' + encodeURIComponent(name))
				.success(function(system) {
					system.stations = system.stations || [];
					system.stations.sort(function(a, b) {
						return a.localeCompare(b);
					});
					success(system);
				})
				.error(error);
		},
		station: function(system, station, success, error) {
			error = error || Function();
			$http.get('/api/systems/' + encodeURIComponent(system) + '/' + encodeURIComponent(station))
				.success(success)
				.error(error);
		},
		setEntry: function(system, station, entry, data, success, error) {
			error = error || Function();
			$http.post('/api/systems/' + encodeURIComponent(system) + '/' + encodeURIComponent(station) + '/market/' + encodeURIComponent(entry), data)
				.success(success)
				.error(error);
		},
		setServices: function(system, station, data, success, error) {
			error = error || Function();
			$http.post('/api/systems/' + encodeURIComponent(system) + '/' + encodeURIComponent(station) + '/services', data)
				.success(success)
				.error(error);
		},
		createSystem: function(data, success, error) {
			error = error || Function();
			$http.post('/api/systems', data)
				.success(success)
				.error(error);
		},
		createStation: function(system, data, success, error) {
			error = error || Function();
			$http.post('/api/systems/' + encodeURIComponent(system), data)
				.success(success)
				.error(error);
		}
	}
}
