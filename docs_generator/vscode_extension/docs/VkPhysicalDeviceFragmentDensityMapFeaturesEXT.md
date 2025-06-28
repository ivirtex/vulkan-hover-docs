# VkPhysicalDeviceFragmentDensityMapFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceFragmentDensityMapFeaturesEXT - Structure describing fragment density map features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceFragmentDensityMapFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_fragment_density_map
typedef struct VkPhysicalDeviceFragmentDensityMapFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           fragmentDensityMap;
    VkBool32           fragmentDensityMapDynamic;
    VkBool32           fragmentDensityMapNonSubsampledImages;
} VkPhysicalDeviceFragmentDensityMapFeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`fragmentDensityMap` specifies whether the implementation supports render passes with a fragment density map attachment. If this feature is not enabled and the `pNext` chain of [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html) includes a [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html) structure, `fragmentDensityMapAttachment` **must** be `VK_ATTACHMENT_UNUSED`.
- []()`fragmentDensityMapDynamic` specifies whether the implementation supports dynamic fragment density map image views. If this feature is not enabled, `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT` **must** not be included in [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html)::`flags`.
- []()`fragmentDensityMapNonSubsampledImages` specifies whether the implementation supports regular non-subsampled image attachments with fragment density map render passes. If this feature is not enabled, render passes with a [fragment density map attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-fragmentdensitymapattachment) **must** only have [subsampled attachments](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-subsamplesampler) bound.

## [](#_description)Description

If the `VkPhysicalDeviceFragmentDensityMapFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceFragmentDensityMapFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceFragmentDensityMapFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceFragmentDensityMapFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_fragment\_density\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_fragment_density_map.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceFragmentDensityMapFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0