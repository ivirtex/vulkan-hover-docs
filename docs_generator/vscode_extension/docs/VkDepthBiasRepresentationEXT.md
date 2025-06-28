# VkDepthBiasRepresentationEXT(3) Manual Page

## Name

VkDepthBiasRepresentationEXT - Specify the depth bias representation



## [](#_c_specification)C Specification

Possible values of [VkDepthBiasRepresentationInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthBiasRepresentationInfoEXT.html)::`depthBiasRepresentation`, specifying the depth bias representation are:

```c++
// Provided by VK_EXT_depth_bias_control
typedef enum VkDepthBiasRepresentationEXT {
    VK_DEPTH_BIAS_REPRESENTATION_LEAST_REPRESENTABLE_VALUE_FORMAT_EXT = 0,
    VK_DEPTH_BIAS_REPRESENTATION_LEAST_REPRESENTABLE_VALUE_FORCE_UNORM_EXT = 1,
    VK_DEPTH_BIAS_REPRESENTATION_FLOAT_EXT = 2,
} VkDepthBiasRepresentationEXT;
```

## [](#_description)Description

- `VK_DEPTH_BIAS_REPRESENTATION_LEAST_REPRESENTABLE_VALUE_FORMAT_EXT` specifies that the depth bias representation is a factor of the format’s r as described in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-depthbias-computation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-depthbias-computation).
- `VK_DEPTH_BIAS_REPRESENTATION_LEAST_REPRESENTABLE_VALUE_FORCE_UNORM_EXT` specifies that the depth bias representation is a factor of a constant r defined by the bit-size or mantissa of the format as described in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-depthbias-computation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-depthbias-computation).
- `VK_DEPTH_BIAS_REPRESENTATION_FLOAT_EXT` specifies that the depth bias representation is a factor of constant r equal to 1.

## [](#_see_also)See Also

[VK\_EXT\_depth\_bias\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_bias_control.html), [VkDepthBiasRepresentationInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthBiasRepresentationInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDepthBiasRepresentationEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0