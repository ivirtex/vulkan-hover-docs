# VkFormatProperties(3) Manual Page

## Name

VkFormatProperties - Structure specifying image format properties



## [](#_c_specification)C Specification

The `VkFormatProperties` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkFormatProperties {
    VkFormatFeatureFlags    linearTilingFeatures;
    VkFormatFeatureFlags    optimalTilingFeatures;
    VkFormatFeatureFlags    bufferFeatures;
} VkFormatProperties;
```

## [](#_members)Members

- `linearTilingFeatures` is a bitmask of [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html) specifying features supported by images created with a `tiling` parameter of `VK_IMAGE_TILING_LINEAR`.
- `optimalTilingFeatures` is a bitmask of [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html) specifying features supported by images created with a `tiling` parameter of `VK_IMAGE_TILING_OPTIMAL`.
- `bufferFeatures` is a bitmask of [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html) specifying features supported by buffers.

## [](#_description)Description

Note

If no format feature flags are supported, the format itself is not supported, and images of that format cannot be created.

If `format` is block-compressed, [requires sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion), or is a depth/stencil format then `bufferFeatures` **must** not support any features for the format.

If `format` is not a multi-plane format then `linearTilingFeatures` and `optimalTilingFeatures` **must** not contain `VK_FORMAT_FEATURE_DISJOINT_BIT`.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkFormatFeatureFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlags.html), [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties2.html), [vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFormatProperties.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFormatProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0