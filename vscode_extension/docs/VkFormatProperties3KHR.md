# VkFormatProperties3(3) Manual Page

## Name

VkFormatProperties3 - Structure specifying image format properties



## [](#_c_specification)C Specification

To query supported format extended features which are properties of the physical device, add [VkFormatProperties3](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties3.html) structure to the `pNext` chain of [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties2.html).

The [VkFormatProperties3](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties3.html) structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkFormatProperties3 {
    VkStructureType          sType;
    void*                    pNext;
    VkFormatFeatureFlags2    linearTilingFeatures;
    VkFormatFeatureFlags2    optimalTilingFeatures;
    VkFormatFeatureFlags2    bufferFeatures;
} VkFormatProperties3;
```

or the equivalent

```c++
// Provided by VK_KHR_format_feature_flags2
typedef VkFormatProperties3 VkFormatProperties3KHR;
```

## [](#_members)Members

- `linearTilingFeatures` is a bitmask of [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html) specifying features supported by images created with a `tiling` parameter of `VK_IMAGE_TILING_LINEAR`.
- `optimalTilingFeatures` is a bitmask of [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html) specifying features supported by images created with a `tiling` parameter of `VK_IMAGE_TILING_OPTIMAL`.
- `bufferFeatures` is a bitmask of [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html) specifying features supported by buffers.

## [](#_description)Description

The bits reported in `linearTilingFeatures`, `optimalTilingFeatures` and `bufferFeatures` **must** include the bits reported in the corresponding fields of `VkFormatProperties2`::`formatProperties`.

Valid Usage (Implicit)

- [](#VUID-VkFormatProperties3-sType-sType)VUID-VkFormatProperties3-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_3`

## [](#_see_also)See Also

[VK\_KHR\_format\_feature\_flags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_format_feature_flags2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkFormatFeatureFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlags2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFormatProperties3).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0