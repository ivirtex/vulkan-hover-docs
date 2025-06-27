# VkFormatProperties(3) Manual Page

## Name

VkFormatProperties - Structure specifying image format properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkFormatProperties` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkFormatProperties {
    VkFormatFeatureFlags    linearTilingFeatures;
    VkFormatFeatureFlags    optimalTilingFeatures;
    VkFormatFeatureFlags    bufferFeatures;
} VkFormatProperties;
```

## <a href="#_members" class="anchor"></a>Members

- `linearTilingFeatures` is a bitmask of
  [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html) specifying
  features supported by images created with a `tiling` parameter of
  `VK_IMAGE_TILING_LINEAR`.

- `optimalTilingFeatures` is a bitmask of
  [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html) specifying
  features supported by images created with a `tiling` parameter of
  `VK_IMAGE_TILING_OPTIMAL`.

- `bufferFeatures` is a bitmask of
  [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html) specifying
  features supported by buffers.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>If no format feature flags are supported, the format itself is not
supported, and images of that format cannot be created.</p></td>
</tr>
</tbody>
</table>

If `format` is a block-compressed format, then `bufferFeatures` **must**
not support any features for the format.

If `format` is not a multi-plane format then `linearTilingFeatures` and
`optimalTilingFeatures` **must** not contain
`VK_FORMAT_FEATURE_DISJOINT_BIT`.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkFormatFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlags.html),
[VkFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2.html),
[vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFormatProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
