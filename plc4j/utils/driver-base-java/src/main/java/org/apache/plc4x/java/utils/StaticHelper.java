/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.java.utils;

import java.util.Collection;

public class StaticHelper {

    public static int COUNT(Object obj) {
        if(obj.getClass().isArray()) {
            Object[] arr = (Object[]) obj;
            return arr.length;
        } else if(obj instanceof Collection) {
            Collection col = (Collection) obj;
            return col.size();
        }
        throw new RuntimeException("Unable to count object of type " + obj.getClass().getName());
    }

    public static <T> T CAST(Object obj, Class<T> clazz) {
        try {
            return clazz.cast(obj);
        } catch(ClassCastException e) {
            throw new RuntimeException("Unable to cast object of type " + obj.getClass().getName() + " to " + clazz.getName());
        }
    }

    public static int CEIL(double value) {
        return (int) Math.ceil(value);
    }

}