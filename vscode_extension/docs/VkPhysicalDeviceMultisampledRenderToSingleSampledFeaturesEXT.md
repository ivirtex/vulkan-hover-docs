# VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT - Structure describing whether multisampled rendering to single-sampled attachments is supported



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_multisampled_render_to_single_sampled
typedef struct VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           multisampledRenderToSingleSampled;
} VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`multisampledRenderToSingleSampled` indicates that the implementation supports multisampled rendering to single-sampled render pass attachments.

## [](#_description)Description

If the `VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_multisampled\_render\_to\_single\_sampled](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_multisampled_render_to_single_sampled.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0