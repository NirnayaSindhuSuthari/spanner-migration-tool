import { Component, Inject, OnInit, inject } from '@angular/core'
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog'
import { ConversionService } from 'src/app/services/conversion/conversion.service'

interface IDialogProps {
    message: string
    type: 'warning' | 'error' | 'success'
    title: string
}
@Component({
    selector: 'app-pop-up-warning',
    templateUrl: './pop-up-warning.component.html',
    styleUrls: ['./pop-up-warning.component.scss'],
})
export class PopUpWarningComponent implements OnInit {
    conversionService = inject(ConversionService)
    constructor(
        public dialogRef: MatDialogRef<PopUpWarningComponent>,
        @Inject(MAT_DIALOG_DATA) public data: IDialogProps,
        // private dialog: MatDialog,
    ) {
        if (data.title === undefined) {
            data.title = 'Update can not be saved'
        }
    }
    ngOnInit(): void {}

    onConfirm(): void {
        // Close the dialog, return true
        this.dialogRef.close(true)
    }

    onDismiss(): void {
        // Close the dialog, return false
        this.dialogRef.close(false)
    }

    getIconFromMessageType() {
        switch (this.data.type) {
            case 'warning':
                return 'warning'
            case 'error':
                return 'error'
            case 'success':
                return 'check_circle'
            default:
                return 'message'
        }
    }
}
