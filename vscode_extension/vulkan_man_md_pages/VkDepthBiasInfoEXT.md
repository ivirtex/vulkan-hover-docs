# VkDepthBiasInfoEXT(3) Manual Page

## Name

VkDepthBiasInfoEXT - Structure specifying depth bias parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDepthBiasInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_depth_bias_control
typedef struct VkDepthBiasInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    float              depthBiasConstantFactor;
    float              depthBiasClamp;
    float              depthBiasSlopeFactor;
} VkDepthBiasInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `depthBiasConstantFactor` is a scalar factor controlling the constant
  depth value added to each fragment.

- `depthBiasClamp` is the maximum (or minimum) depth bias of a fragment.

- `depthBiasSlopeFactor` is a scalar factor applied to a fragment’s
  slope in depth bias calculations.

## <a href="#_description" class="anchor"></a>Description

If `pNext` does not contain a
[VkDepthBiasRepresentationInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDepthBiasRepresentationInfoEXT.html)
structure, then this command is equivalent to including a
[VkDepthBiasRepresentationInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDepthBiasRepresentationInfoEXT.html)
with `depthBiasExact` set to `VK_FALSE` and `depthBiasRepresentation`
set to
`VK_DEPTH_BIAS_REPRESENTATION_LEAST_REPRESENTABLE_VALUE_FORMAT_EXT`.

Valid Usage

- <a href="#VUID-VkDepthBiasInfoEXT-depthBiasClamp-08950"
  id="VUID-VkDepthBiasInfoEXT-depthBiasClamp-08950"></a>
  VUID-VkDepthBiasInfoEXT-depthBiasClamp-08950  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-depthBiasClamp"
  target="_blank" rel="noopener"><code>depthBiasClamp</code></a> feature
  is not enabled, `depthBiasClamp` **must** be `0.0`

Valid Usage (Implicit)

- <a href="#VUID-VkDepthBiasInfoEXT-sType-sType"
  id="VUID-VkDepthBiasInfoEXT-sType-sType"></a>
  VUID-VkDepthBiasInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEPTH_BIAS_INFO_EXT`

- <a href="#VUID-VkDepthBiasInfoEXT-pNext-pNext"
  id="VUID-VkDepthBiasInfoEXT-pNext-pNext"></a>
  VUID-VkDepthBiasInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkDepthBiasRepresentationInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDepthBiasRepresentationInfoEXT.html)

- <a href="#VUID-VkDepthBiasInfoEXT-sType-unique"
  id="VUID-VkDepthBiasInfoEXT-sType-unique"></a>
  VUID-VkDepthBiasInfoEXT-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_depth_bias_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_depth_bias_control.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdSetDepthBias2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBias2EXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDepthBiasInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
