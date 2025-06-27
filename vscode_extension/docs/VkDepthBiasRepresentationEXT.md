# VkDepthBiasRepresentationEXT(3) Manual Page

## Name

VkDepthBiasRepresentationEXT - Specify the depth bias representation



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkDepthBiasRepresentationInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDepthBiasRepresentationInfoEXT.html)::`depthBiasRepresentation`,
specifying the depth bias representation are:

``` c
// Provided by VK_EXT_depth_bias_control
typedef enum VkDepthBiasRepresentationEXT {
    VK_DEPTH_BIAS_REPRESENTATION_LEAST_REPRESENTABLE_VALUE_FORMAT_EXT = 0,
    VK_DEPTH_BIAS_REPRESENTATION_LEAST_REPRESENTABLE_VALUE_FORCE_UNORM_EXT = 1,
    VK_DEPTH_BIAS_REPRESENTATION_FLOAT_EXT = 2,
} VkDepthBiasRepresentationEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DEPTH_BIAS_REPRESENTATION_LEAST_REPRESENTABLE_VALUE_FORMAT_EXT`
  specifies that the depth bias representation is a factor of the
  format’s r as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-depthbias-computation"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-depthbias-computation</a>.

- `VK_DEPTH_BIAS_REPRESENTATION_LEAST_REPRESENTABLE_VALUE_FORCE_UNORM_EXT`
  specifies that the depth bias representation is a factor of a constant
  r defined by the bit-size or mantissa of the format as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-depthbias-computation"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-depthbias-computation</a>.

- `VK_DEPTH_BIAS_REPRESENTATION_FLOAT_EXT` specifies that the depth bias
  representation is a factor of constant r equal to 1.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_depth_bias_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_depth_bias_control.html),
[VkDepthBiasRepresentationInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDepthBiasRepresentationInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDepthBiasRepresentationEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
