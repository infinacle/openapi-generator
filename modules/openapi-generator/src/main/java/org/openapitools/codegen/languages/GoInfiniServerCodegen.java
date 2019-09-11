/*
 * Copyright 2018 OpenAPI-Generator Contributors (https://openapi-generator.tech)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package org.openapitools.codegen.languages;

import org.openapitools.codegen.CodegenConstants;
import org.openapitools.codegen.CodegenOperation;
import org.openapitools.codegen.CodegenType;
import org.openapitools.codegen.SupportingFile;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.File;
import java.util.Arrays;
import java.util.List;
import java.util.Map;

public class GoInfiniServerCodegen extends GoGinServerCodegen {

    private static final Logger LOGGER = LoggerFactory.getLogger(GoInfiniServerCodegen.class);

    protected String apiVersion = "1.0.0";
    protected int serverPort = 8080;
    protected String projectName = "openapi-server";
    protected String apiPath = "go";

    public GoInfiniServerCodegen() {
        super();

        /*
         * Template Location. This is the location which templates will be read from.
         * The generator will use the resource stream to attempt to read the templates.
         */
        embeddedTemplateDir = templateDir = "go-infini-server";
    }

    /**
     * Configures a friendly name for the generator. This will be used by the
     * generator to select the library with the -g flag.
     *
     * @return the friendly name for the generator
     */
    @Override
    public String getName() {
        return "go-infini-server";
    }

    /**
     * Returns human-friendly help for the generator. Provide the consumer with help
     * tips, parameters here
     *
     * @return A string value for the help message
     */
    @Override
    public String getHelp() {
        return "Generates a Go server library with the gin framework using OpenAPI-Generator."
                + "By default, it will also generate service classes.";
    }
}
