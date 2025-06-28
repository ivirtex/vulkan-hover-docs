# VkDepthBiasInfoEXT(3) Manual Page

## Name

VkDepthBiasInfoEXT - Structure specifying depth bias parameters



## [](#_c_specification)C Specification

The `VkDepthBiasInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_depth_bias_control
typedef struct VkDepthBiasInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    float              depthBiasConstantFactor;
    float              depthBiasClamp;
    float              depthBiasSlopeFactor;
} VkDepthBiasInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `depthBiasConstantFactor` is a scalar factor controlling the constant depth value added to each fragment.
- `depthBiasClamp` is the maximum (or minimum) depth bias of a fragment.
- `depthBiasSlopeFactor` is a scalar factor applied to a fragment’s slope in depth bias calculations.

## [](#_description)Description

If `pNext` does not contain a [VkDepthBiasRepresentationInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthBiasRepresentationInfoEXT.html) structure, then this command is equivalent to including a [VkDepthBiasRepresentationInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthBiasRepresentationInfoEXT.html) with `depthBiasExact` set to `VK_FALSE` and `depthBiasRepresentation` set to `VK_DEPTH_BIAS_REPRESENTATION_LEAST_REPRESENTABLE_VALUE_FORMAT_EXT`.

Valid Usage

- [](#VUID-VkDepthBiasInfoEXT-depthBiasClamp-08950)VUID-VkDepthBiasInfoEXT-depthBiasClamp-08950  
  If the [`depthBiasClamp`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-depthBiasClamp) feature is not enabled, `depthBiasClamp` **must** be `0.0`

Valid Usage (Implicit)

- [](#VUID-VkDepthBiasInfoEXT-sType-sType)VUID-VkDepthBiasInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEPTH_BIAS_INFO_EXT`
- [](#VUID-VkDepthBiasInfoEXT-pNext-pNext)VUID-VkDepthBiasInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkDepthBiasRepresentationInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthBiasRepresentationInfoEXT.html)
- [](#VUID-VkDepthBiasInfoEXT-sType-unique)VUID-VkDepthBiasInfoEXT-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_EXT\_depth\_bias\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_bias_control.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdSetDepthBias2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthBias2EXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDepthBiasInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0