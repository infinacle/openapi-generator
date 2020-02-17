package org.openapitools.codegen.languages;

public class AndroidInfiniClientCodegen extends AndroidClientCodegen {

    public AndroidInfiniClientCodegen() {
        super();
        embeddedTemplateDir = templateDir = "android-infini";
    }

    @Override
    public String getName() {
        return "android-infini";
    }
}