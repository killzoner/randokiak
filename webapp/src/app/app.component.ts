import { Component, OnDestroy } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Environment } from 'src/environments/environment.interface';
import { Subscription, Subject, forkJoin, of } from 'rxjs';
import { map, catchError, first, switchMap, delay } from 'rxjs/operators';
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
  nbResults = new Subject<number>();

  constructor(private http: HttpClient) {
    this.configuration = environment;

    this.esSearch()
      .pipe(
        map((value) => {
          this.nbResults.next(value.hits.total.value);
          return value.hits.hits.map(
            doc => ProfileReply.fromJSON(doc._source).profile
          );
        }
        ),
        catchError(() => []),
        first()
      )
      .subscribe((values) => {
        this.dataSource = values;
      });
  }

  getMorePeople() {
    const baseUrl = this.configuration.rdkapiHttpEndpoint;
    this.http.get<{ profiles: { profile: Profile }[] }>(`${baseUrl}/api/v1/profiles`)
      .pipe(
        delay(1500), // delay call to ES to give time to process new items
        switchMap((values) => {
          const profiles = values.profiles.map((p) => p.profile);
          return forkJoin([
            this.esSearch(),
            of(profiles)
          ]);
        }),
      ).subscribe(([rootDoc, profiles]) => {
        this.nbResults.next(rootDoc.hits.total.value);
        this.dataSource = profiles;
      });
  }

  ngOnDestroy() {
    this.subscription.unsubscribe();
  }

  private esSearch() {
    const baseUrl = this.configuration.elasticHttpEndpoint;
    const indexName = this.configuration.searchIndex;

    return this.http.get<EsRootDoc>(
      `${baseUrl}/${indexName}/_search`);
  }
}
