# VkPhysicalDevice4444FormatsFeaturesEXT(3) Manual Page

## Name

VkPhysicalDevice4444FormatsFeaturesEXT - Structure describing additional 4444 formats supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDevice4444FormatsFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_4444_formats
typedef struct VkPhysicalDevice4444FormatsFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           formatA4R4G4B4;
    VkBool32           formatA4B4G4R4;
} VkPhysicalDevice4444FormatsFeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`formatA4R4G4B4` indicates that the implementation **must** support using a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `VK_FORMAT_A4R4G4B4_UNORM_PACK16_EXT` with at least the following [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html):
  
  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT`
  - `VK_FORMAT_FEATURE_BLIT_SRC_BIT`
  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`
- []()`formatA4B4G4R4` indicates that the implementation **must** support using a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `VK_FORMAT_A4B4G4R4_UNORM_PACK16_EXT` with at least the following [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html):
  
  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT`
  - `VK_FORMAT_FEATURE_BLIT_SRC_BIT`
  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`

## [](#_description)Description

If the `VkPhysicalDevice4444FormatsFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDevice4444FormatsFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDevice4444FormatsFeaturesEXT-sType-sType)VUID-VkPhysicalDevice4444FormatsFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_4444_FORMATS_FEATURES_EXT`

Note

Although the formats defined by the `VK_EXT_4444_formats` extension were promoted to Vulkan 1.3 as optional formats, the [VkPhysicalDevice4444FormatsFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice4444FormatsFeaturesEXT.html) structure was not promoted to Vulkan 1.3.

## [](#_see_also)See Also

[VK\_EXT\_4444\_formats](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_4444_formats.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDevice4444FormatsFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0