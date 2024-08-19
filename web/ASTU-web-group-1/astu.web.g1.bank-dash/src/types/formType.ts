// Form Types - Input Group ( Label and INput Fiels )
export interface InputGroupType {
  id: string;
  label: string;
  inputType: string;
  registerName: string;
  register: any;
  placeholder: string;
  errorMessage?: string;
}

export interface ToggleInputType extends InputGroupType {
  currentState: boolean
}