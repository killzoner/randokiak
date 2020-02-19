import { Component, OnDestroy } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Environment } from 'src/environments/environment.interface';
import { Subscription } from 'rxjs';
import { map, catchError } from 'rxjs/operators';
import { ProfileReply } from './models/protos/random';
import { Profile } from './models/protos/randomprofile';
import { EsRootDoc } from './models/elastic';

@Component({
  selector: 'app-root',
  styleUrls: ['app.component.sass'],
  templateUrl: 'app.component.html',
})
export class AppComponent implements OnDestroy {

  private configuration: Environment;
  private subscription: Subscription = new Subscription();

  displayedColumns: string[] = ['gender', 'firstName', 'lastName', 'nat'];
  dataSource: Profile[] = [];

  constructor(private http: HttpClient) {
    this.configuration = environment;

    const baseUrl = this.configuration.elasticHttpEndpoint;
    const indexName = this.configuration.searchIndex;

    this.http.get<EsRootDoc>(
      `${baseUrl}/${indexName}/_search`)
      .pipe(
        map((value) =>
          value.hits.hits.map(
            doc => ProfileReply.fromJSON(doc._source).profile
          )
        ),
        catchError(() => [])
      )
      .subscribe((values) => {
        console.log(values);
        this.dataSource = values;
      });
  }

  getMorePeople() {
    const baseUrl = this.configuration.rdkapiHttpEndpoint;
    this.http.get(`${baseUrl}/api/v1/dumb`).subscribe();
  }

  ngOnDestroy() {
    this.subscription.unsubscribe();
  }
}
