import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';

import {AppComponent} from './app.component';
import {HeadercComponent} from './headerc/headerc.component';
import {HttpClientModule} from "@angular/common/http";
import {RouterModule, Routes} from "@angular/router";
import {TitlesComponent} from './titles/titles.component';
import { ProfileComponent } from './profile/profile.component';

const appRoutes: Routes = [
  {path: 'titles', component: TitlesComponent},
  {path: 'profile', component: ProfileComponent},
]

@NgModule({
  declarations: [
    AppComponent,
    HeadercComponent,
    TitlesComponent,
    ProfileComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    RouterModule.forRoot(appRoutes)
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {
}
