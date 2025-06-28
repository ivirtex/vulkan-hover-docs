# VkPhysicalDeviceBorderColorSwizzleFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceBorderColorSwizzleFeaturesEXT - Structure describing whether samplers with custom border colors require the component swizzle specified in order to have defined behavior



## [](#_c_specification)C Specification

The `VkPhysicalDeviceBorderColorSwizzleFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_border_color_swizzle
typedef struct VkPhysicalDeviceBorderColorSwizzleFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           borderColorSwizzle;
    VkBool32           borderColorSwizzleFromImage;
} VkPhysicalDeviceBorderColorSwizzleFeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`borderColorSwizzle` indicates that defined values are returned by sampled image operations when used with a sampler that uses a `VK_BORDER_COLOR_FLOAT_OPAQUE_BLACK`, `VK_BORDER_COLOR_INT_OPAQUE_BLACK`, `VK_BORDER_COLOR_FLOAT_CUSTOM_EXT`, or `VK_BORDER_COLOR_INT_CUSTOM_EXT` `borderColor` and an image view that uses a non-[identity component mapping](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views-identity-mappings), when either `borderColorSwizzleFromImage` is enabled or the [VkSamplerBorderColorComponentMappingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerBorderColorComponentMappingCreateInfoEXT.html) is specified.
- []()`borderColorSwizzleFromImage` indicates that the implementation will return the correct border color values from sampled image operations under the conditions expressed above, without the application having to specify the border color component mapping when creating the sampler object. If this feature bit is not set, applications **can** chain a [VkSamplerBorderColorComponentMappingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerBorderColorComponentMappingCreateInfoEXT.html) structure when creating samplers for use with image views that do not have an [identity swizzle](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views-identity-mappings) and, when those samplers are combined with image views using the same component mapping, sampled image operations that use opaque black or custom border colors will return the correct border color values.

## [](#_description)Description

If the `VkPhysicalDeviceBorderColorSwizzleFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceBorderColorSwizzleFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceBorderColorSwizzleFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceBorderColorSwizzleFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BORDER_COLOR_SWIZZLE_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_border\_color\_swizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_border_color_swizzle.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceBorderColorSwizzleFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0