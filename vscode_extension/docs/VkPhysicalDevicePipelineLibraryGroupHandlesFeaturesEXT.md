# VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT(3) Manual Page

## Name

VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT - Structure describing whether querying shader group handles from a pipeline library is supported by the implementation



## [](#_c_specification)C Specification

The `VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_pipeline_library_group_handles
typedef struct VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           pipelineLibraryGroupHandles;
} VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- []()`pipelineLibraryGroupHandles` indicates whether the implementation supports querying group handles directly from a ray tracing pipeline library, and guarantees bitwise identical group handles for such libraries when linked into other pipelines.

## [](#_description)Description

If the `VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT-sType-sType)VUID-VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_LIBRARY_GROUP_HANDLES_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_pipeline\_library\_group\_handles](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_library_group_handles.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0