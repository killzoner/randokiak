import { Environment } from './environment.interface';

export const environment: Environment = {
  production: true,
  elasticHttpEndpoint: '/elasticsearch',
  searchIndex: 'randokiak-index',
  rdkapiHttpEndpoint: '/rdkapi'
};
