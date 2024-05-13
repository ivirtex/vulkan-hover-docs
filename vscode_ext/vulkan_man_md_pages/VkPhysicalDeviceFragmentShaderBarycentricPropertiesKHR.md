# VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR - Structure
describing fragment shader barycentric limits of an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR` structure
is defined as:

``` c
// Provided by VK_KHR_fragment_shader_barycentric
typedef struct VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           triStripVertexOrderIndependentOfProvokingVertex;
} VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- <span id="limits-triStripVertexOrderIndependentOfProvokingVertex"></span>
  When the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vertexpostproc-flatshading"
  target="_blank" rel="noopener">provoking vertex mode</a> is
  `VK_PROVOKING_VERTEX_MODE_LAST_VERTEX_EXT`, and the primitive order is
  odd in a triangle strip, the ordering of vertices is defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-barycentric-order-table-last-vertex"
  target="_blank" rel="noopener">last vertex table</a>.
  `triStripVertexOrderIndependentOfProvokingVertex` equal to `VK_TRUE`
  indicates that the implementation ignores this and uses the vertex
  order defined by `VK_PROVOKING_VERTEX_MODE_FIRST_VERTEX_EXT` instead.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR`
structure is included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_BARYCENTRIC_PROPERTIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_fragment_shader_barycentric](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_fragment_shader_barycentric.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
