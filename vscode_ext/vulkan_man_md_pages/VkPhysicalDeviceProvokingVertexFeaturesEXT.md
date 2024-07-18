# VkPhysicalDeviceProvokingVertexFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceProvokingVertexFeaturesEXT - Structure describing the
provoking vertex features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceProvokingVertexFeaturesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_provoking_vertex
typedef struct VkPhysicalDeviceProvokingVertexFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           provokingVertexLast;
    VkBool32           transformFeedbackPreservesProvokingVertex;
} VkPhysicalDeviceProvokingVertexFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-provokingVertexLast"></span> `provokingVertexLast`
  indicates whether the implementation supports the
  `VK_PROVOKING_VERTEX_MODE_LAST_VERTEX_EXT`
  <a href="VkProvokingVertexModeEXT.html" target="_blank"
  rel="noopener">provoking vertex mode</a> for flat shading.

- <span id="features-transformFeedbackPreservesProvokingVertex"></span>
  `transformFeedbackPreservesProvokingVertex` indicates that the order
  of vertices within each primitive written by transform feedback will
  preserve the provoking vertex. This does not apply to triangle fan
  primitives when <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-transformFeedbackPreservesTriangleFanProvokingVertex"
  target="_blank"
  rel="noopener"><code>transformFeedbackPreservesTriangleFanProvokingVertex</code></a>
  is `VK_FALSE`. `transformFeedbackPreservesProvokingVertex` **must** be
  `VK_FALSE` when the
  [`VK_EXT_transform_feedback`](VK_EXT_transform_feedback.html)
  extension is not supported.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceProvokingVertexFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceProvokingVertexFeaturesEXT` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

When `VkPhysicalDeviceProvokingVertexFeaturesEXT` is in the `pNext`
chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) but the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-transformFeedback"
target="_blank" rel="noopener"><code>transformFeedback</code></a>
feature is not enabled, the value of
`transformFeedbackPreservesProvokingVertex` is ignored.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceProvokingVertexFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceProvokingVertexFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceProvokingVertexFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROVOKING_VERTEX_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_provoking_vertex](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_provoking_vertex.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceProvokingVertexFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
