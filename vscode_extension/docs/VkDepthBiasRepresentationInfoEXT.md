# VkDepthBiasRepresentationInfoEXT(3) Manual Page

## Name

VkDepthBiasRepresentationInfoEXT - Structure specifying depth bias parameters



## [](#_c_specification)C Specification

The `VkDepthBiasRepresentationInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_depth_bias_control
typedef struct VkDepthBiasRepresentationInfoEXT {
    VkStructureType                 sType;
    const void*                     pNext;
    VkDepthBiasRepresentationEXT    depthBiasRepresentation;
    VkBool32                        depthBiasExact;
} VkDepthBiasRepresentationInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `depthBiasRepresentation` is a [VkDepthBiasRepresentationEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthBiasRepresentationEXT.html) value specifying the depth bias representation.
- `depthBiasExact` specifies that the implementation is not allowed to scale the depth bias value to ensure a minimum resolvable distance.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDepthBiasRepresentationInfoEXT-leastRepresentableValueForceUnormRepresentation-08947)VUID-VkDepthBiasRepresentationInfoEXT-leastRepresentableValueForceUnormRepresentation-08947  
  If the [`leastRepresentableValueForceUnormRepresentation`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-leastRepresentableValueForceUnormRepresentation) feature is not enabled, `depthBiasRepresentation` **must** not be `VK_DEPTH_BIAS_REPRESENTATION_LEAST_REPRESENTABLE_VALUE_FORCE_UNORM_EXT`
- [](#VUID-VkDepthBiasRepresentationInfoEXT-floatRepresentation-08948)VUID-VkDepthBiasRepresentationInfoEXT-floatRepresentation-08948  
  If the [`floatRepresentation`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-floatRepresentation) feature is not enabled, `depthBiasRepresentation` **must** not be `VK_DEPTH_BIAS_REPRESENTATION_FLOAT_EXT`
- [](#VUID-VkDepthBiasRepresentationInfoEXT-depthBiasExact-08949)VUID-VkDepthBiasRepresentationInfoEXT-depthBiasExact-08949  
  If the [`depthBiasExact`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-depthBiasExact) feature is not enabled, `depthBiasExact` **must** be `VK_FALSE`

Valid Usage (Implicit)

- [](#VUID-VkDepthBiasRepresentationInfoEXT-sType-sType)VUID-VkDepthBiasRepresentationInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEPTH_BIAS_REPRESENTATION_INFO_EXT`
- [](#VUID-VkDepthBiasRepresentationInfoEXT-depthBiasRepresentation-parameter)VUID-VkDepthBiasRepresentationInfoEXT-depthBiasRepresentation-parameter  
  `depthBiasRepresentation` **must** be a valid [VkDepthBiasRepresentationEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthBiasRepresentationEXT.html) value

## [](#_see_also)See Also

[VK\_EXT\_depth\_bias\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_bias_control.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkDepthBiasRepresentationEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthBiasRepresentationEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDepthBiasRepresentationInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0