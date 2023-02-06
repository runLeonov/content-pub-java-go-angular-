import {Injectable, NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';

import {AppComponent} from './app.component';
import {HeadercComponent} from './headerc/headerc.component';
import {
  HTTP_INTERCEPTORS,
  HttpClientModule,
  HttpEvent,
  HttpHandler,
  HttpInterceptor,
  HttpRequest
} from "@angular/common/http";
import {RouterModule, Routes} from "@angular/router";
import {TitlesComponent} from './titles/titles.component';
import {ProfileComponent} from './profile/profile.component';
import {PublishComponent} from './publish/publish.component';
import {FormsModule} from "@angular/forms";
import {TitleViewComponent} from './titleview/title-view.component';
import {TitleViewRandomComponent} from './title-view-random/title-view-random.component';

import {AuthPageComponent} from './auth-page/auth-page.component';
import {Observable} from "rxjs";

const appRoutes: Routes = [
  {path: 'titles', component: TitlesComponent},
  {path: 'profile', component: ProfileComponent},
  {path: 'title/create', component: PublishComponent},
  {path: 'random', component: TitleViewRandomComponent},
  {path: 'titles/:id', component: TitleViewComponent},
  {path: 'sing-up', component: AuthPageComponent},
]

@Injectable()
export class CustomInterceptor implements HttpInterceptor {

  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {

    request = request.clone({
      withCredentials: true
    });

    return next.handle(request);
  }
}

@NgModule({
  declarations: [
    AppComponent,
    HeadercComponent,
    TitlesComponent,
    ProfileComponent,
    PublishComponent,
    TitleViewComponent,
    TitleViewRandomComponent,
    AuthPageComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    RouterModule.forRoot(appRoutes),
    FormsModule,
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: CustomInterceptor,
      multi: true
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule {

  public getUser() {

  }
}
