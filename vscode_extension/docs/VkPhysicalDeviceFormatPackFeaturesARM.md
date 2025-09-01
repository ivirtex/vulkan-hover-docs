# VkPhysicalDeviceFormatPackFeaturesARM(3) Manual Page

## Name

VkPhysicalDeviceFormatPackFeaturesARM - Structure describing whether the additional formats feature is supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceFormatPackFeaturesARM` structure is defined as:

```c++
// Provided by VK_ARM_format_pack
typedef struct VkPhysicalDeviceFormatPackFeaturesARM {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           formatPack;
} VkPhysicalDeviceFormatPackFeaturesARM;
```

## [](#_members)Members

This structure describes the following feature:

- []()`formatPack` indicates that the implementation **must** support using a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `VK_FORMAT_R10X6_UINT_PACK16_ARM`, `VK_FORMAT_R10X6G10X6_UINT_2PACK16_ARM`, `VK_FORMAT_R10X6G10X6B10X6A10X6_UINT_4PACK16_ARM`, `VK_FORMAT_R12X4_UINT_PACK16_ARM`, `VK_FORMAT_R12X4G12X4_UINT_2PACK16_ARM`, `VK_FORMAT_R12X4G12X4B12X4A12X4_UINT_4PACK16_ARM`, `VK_FORMAT_R14X2_UINT_PACK16_ARM`, `VK_FORMAT_R14X2G14X2_UINT_2PACK16_ARM`, and `VK_FORMAT_R14X2G14X2B14X2A14X2_UINT_4PACK16_ARM`, with at least the following [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html):
  
  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT`
  - `VK_FORMAT_FEATURE_TRANSFER_SRC_BIT`
  - `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`

## [](#_description)Description

If the `VkPhysicalDeviceFormatPackFeaturesARM` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceFormatPackFeaturesARM`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceFormatPackFeaturesARM-sType-sType)VUID-VkPhysicalDeviceFormatPackFeaturesARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FORMAT_PACK_FEATURES_ARM`

## [](#_see_also)See Also

[VK\_ARM\_format\_pack](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_format_pack.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceFormatPackFeaturesARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0