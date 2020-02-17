package org.openapitools.codegen.languages;

public class JavaInfiniClientCodegen extends JavaClientCodegen {
    public JavaInfiniClientCodegen() {
        super();
        embeddedTemplateDir = templateDir = "java-infini";
    }

    @Override
    public String getName() {
        return "java-infini";
    }
}