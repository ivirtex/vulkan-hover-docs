# VkPhysicalDeviceProvokingVertexFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceProvokingVertexFeaturesEXT - Structure describing the provoking vertex features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceProvokingVertexFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_provoking_vertex
typedef struct VkPhysicalDeviceProvokingVertexFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           provokingVertexLast;
    VkBool32           transformFeedbackPreservesProvokingVertex;
} VkPhysicalDeviceProvokingVertexFeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`provokingVertexLast` indicates whether the implementation supports the `VK_PROVOKING_VERTEX_MODE_LAST_VERTEX_EXT` [provoking vertex mode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkProvokingVertexModeEXT.html) for flat shading.
- []()`transformFeedbackPreservesProvokingVertex` indicates that the order of vertices within each primitive written by transform feedback will preserve the provoking vertex. This does not apply to triangle fan primitives when [`transformFeedbackPreservesTriangleFanProvokingVertex`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-transformFeedbackPreservesTriangleFanProvokingVertex) is `VK_FALSE`. `transformFeedbackPreservesProvokingVertex` **must** be `VK_FALSE` when the `VK_EXT_transform_feedback` extension is not supported.

## [](#_description)Description

If the `VkPhysicalDeviceProvokingVertexFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceProvokingVertexFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

When `VkPhysicalDeviceProvokingVertexFeaturesEXT` is in the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) but the [`transformFeedback`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-transformFeedback) feature is not enabled, the value of `transformFeedbackPreservesProvokingVertex` is ignored.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceProvokingVertexFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceProvokingVertexFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROVOKING_VERTEX_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_provoking\_vertex](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_provoking_vertex.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceProvokingVertexFeaturesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0