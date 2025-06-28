# VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR(3) Manual Page

## Name

VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR - Structure describing whether pipeline executable properties are available



## [](#_c_specification)C Specification

The `VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_pipeline_executable_properties
typedef struct VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           pipelineExecutableInfo;
} VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`pipelineExecutableInfo` indicates that the implementation supports reporting properties and statistics about the pipeline executables associated with a compiled pipeline.

## [](#_description)Description

If the `VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR-sType-sType)VUID-VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_EXECUTABLE_PROPERTIES_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_executable\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_executable_properties.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0