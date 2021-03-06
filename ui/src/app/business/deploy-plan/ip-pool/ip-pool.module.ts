import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {IpPoolComponent} from './ip-pool.component';
import { IpPoolCreateComponent } from './ip-pool-create/ip-pool-create.component';
import { IpPoolDeleteComponent } from './ip-pool-delete/ip-pool-delete.component';
import { IpPoolListComponent } from './ip-pool-list/ip-pool-list.component';
import {CoreModule} from '../../../core/core.module';
import {SharedModule} from '../../../shared/shared.module';


@NgModule({
    declarations: [IpPoolComponent, IpPoolCreateComponent, IpPoolDeleteComponent, IpPoolListComponent],
    imports: [
        CommonModule,
        CoreModule,
        SharedModule
    ]
})
export class IpPoolModule {
}
