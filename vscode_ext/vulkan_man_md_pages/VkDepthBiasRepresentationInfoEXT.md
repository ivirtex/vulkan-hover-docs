# VkDepthBiasRepresentationInfoEXT(3) Manual Page

## Name

VkDepthBiasRepresentationInfoEXT - Structure specifying depth bias
parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDepthBiasRepresentationInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_depth_bias_control
typedef struct VkDepthBiasRepresentationInfoEXT {
    VkStructureType                 sType;
    const void*                     pNext;
    VkDepthBiasRepresentationEXT    depthBiasRepresentation;
    VkBool32                        depthBiasExact;
} VkDepthBiasRepresentationInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `depthBiasRepresentation` is a
  [VkDepthBiasRepresentationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDepthBiasRepresentationEXT.html)
  value specifying the depth bias representation.

- `depthBiasExact` specifies that the implementation is not allowed to
  scale the depth bias value to ensure a minimum resolvable distance.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkDepthBiasRepresentationInfoEXT-leastRepresentableValueForceUnormRepresentation-08947"
  id="VUID-VkDepthBiasRepresentationInfoEXT-leastRepresentableValueForceUnormRepresentation-08947"></a>
  VUID-VkDepthBiasRepresentationInfoEXT-leastRepresentableValueForceUnormRepresentation-08947  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-leastRepresentableValueForceUnormRepresentation"
  target="_blank"
  rel="noopener"><code>leastRepresentableValueForceUnormRepresentation</code></a>
  feature is not enabled, `depthBiasRepresentation` **must** not be
  `VK_DEPTH_BIAS_REPRESENTATION_LEAST_REPRESENTABLE_VALUE_FORCE_UNORM_EXT`

- <a
  href="#VUID-VkDepthBiasRepresentationInfoEXT-floatRepresentation-08948"
  id="VUID-VkDepthBiasRepresentationInfoEXT-floatRepresentation-08948"></a>
  VUID-VkDepthBiasRepresentationInfoEXT-floatRepresentation-08948  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-floatRepresentation"
  target="_blank" rel="noopener"><code>floatRepresentation</code></a>
  feature is not enabled, `depthBiasRepresentation` **must** not be
  `VK_DEPTH_BIAS_REPRESENTATION_FLOAT_EXT`

- <a href="#VUID-VkDepthBiasRepresentationInfoEXT-depthBiasExact-08949"
  id="VUID-VkDepthBiasRepresentationInfoEXT-depthBiasExact-08949"></a>
  VUID-VkDepthBiasRepresentationInfoEXT-depthBiasExact-08949  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-depthBiasExact"
  target="_blank" rel="noopener"><code>depthBiasExact</code></a> feature
  is not enabled, `depthBiasExact` **must** be `VK_FALSE`

Valid Usage (Implicit)

- <a href="#VUID-VkDepthBiasRepresentationInfoEXT-sType-sType"
  id="VUID-VkDepthBiasRepresentationInfoEXT-sType-sType"></a>
  VUID-VkDepthBiasRepresentationInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEPTH_BIAS_REPRESENTATION_INFO_EXT`

- <a
  href="#VUID-VkDepthBiasRepresentationInfoEXT-depthBiasRepresentation-parameter"
  id="VUID-VkDepthBiasRepresentationInfoEXT-depthBiasRepresentation-parameter"></a>
  VUID-VkDepthBiasRepresentationInfoEXT-depthBiasRepresentation-parameter  
  `depthBiasRepresentation` **must** be a valid
  [VkDepthBiasRepresentationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDepthBiasRepresentationEXT.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_depth_bias_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_depth_bias_control.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkDepthBiasRepresentationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDepthBiasRepresentationEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDepthBiasRepresentationInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
