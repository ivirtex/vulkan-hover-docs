# VkPhysicalDevicePipelineOpacityMicromapFeaturesARM(3) Manual Page

## Name

VkPhysicalDevicePipelineOpacityMicromapFeaturesARM - Structure describing features supported by VK\_ARM\_pipeline\_opacity\_micromap



## [](#_c_specification)C Specification

The [VkPhysicalDevicePipelineOpacityMicromapFeaturesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePipelineOpacityMicromapFeaturesARM.html) structure is defined as:

```c++
// Provided by VK_ARM_pipeline_opacity_micromap
typedef struct VkPhysicalDevicePipelineOpacityMicromapFeaturesARM {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           pipelineOpacityMicromap;
} VkPhysicalDevicePipelineOpacityMicromapFeaturesARM;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`pipelineOpacityMicromap` indicates if a pipeline **can** declare if it can be used with an acceleration structure referencing an opacity micromap, or not.

If the `VkPhysicalDevicePipelineOpacityMicromapFeaturesARM` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDevicePipelineOpacityMicromapFeaturesARM`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDevicePipelineOpacityMicromapFeaturesARM-sType-sType)VUID-VkPhysicalDevicePipelineOpacityMicromapFeaturesARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_OPACITY_MICROMAP_FEATURES_ARM`

## [](#_see_also)See Also

[VK\_ARM\_pipeline\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_pipeline_opacity_micromap.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDevicePipelineOpacityMicromapFeaturesARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0