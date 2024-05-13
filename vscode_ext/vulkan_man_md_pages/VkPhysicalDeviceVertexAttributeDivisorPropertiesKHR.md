# VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR - Structure
describing max value of vertex attribute divisor that can be supported
by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR` structure is
defined as:

``` c
// Provided by VK_KHR_vertex_attribute_divisor
typedef struct VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxVertexAttribDivisor;
    VkBool32           supportsNonZeroFirstInstance;
} VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-maxVertexAttribDivisor"></span>
  `maxVertexAttribDivisor` is the maximum value of the number of
  instances that will repeat the value of vertex attribute data when
  instanced rendering is enabled.

- <span id="limits-supportsNonZeroFirstInstance"></span>
  `supportsNonZeroFirstInstance` specifies whether a non-zero value for
  the `firstInstance` parameter of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing"
  target="_blank" rel="noopener">drawing commands</a> is supported when
  [VkVertexInputBindingDivisorDescriptionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDivisorDescriptionKHR.html)::`divisor`
  is not `1`.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_PROPERTIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_vertex_attribute_divisor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_vertex_attribute_divisor.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
