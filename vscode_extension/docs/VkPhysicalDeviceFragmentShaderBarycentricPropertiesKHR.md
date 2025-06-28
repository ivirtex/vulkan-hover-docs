# VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR - Structure describing fragment shader barycentric limits of an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_fragment_shader_barycentric
typedef struct VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           triStripVertexOrderIndependentOfProvokingVertex;
} VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR;
```

## [](#_members)Members

- []()When the [provoking vertex mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vertexpostproc-flatshading) is `VK_PROVOKING_VERTEX_MODE_LAST_VERTEX_EXT`, and the primitive order is odd in a triangle strip, the ordering of vertices is defined in [last vertex table](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-barycentric-order-table-last-vertex). `triStripVertexOrderIndependentOfProvokingVertex` equal to `VK_TRUE` indicates that the implementation ignores this and uses the vertex order defined by `VK_PROVOKING_VERTEX_MODE_FIRST_VERTEX_EXT` instead.

## [](#_description)Description

If the `VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR-sType-sType)VUID-VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_BARYCENTRIC_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_fragment\_shader\_barycentric](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_fragment_shader_barycentric.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0