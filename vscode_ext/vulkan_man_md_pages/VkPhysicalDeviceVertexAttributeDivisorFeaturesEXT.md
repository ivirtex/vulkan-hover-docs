# VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR - Structure describing
if fetching of vertex attribute may be repeated for instanced rendering



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR` structure is
defined as:

``` c
// Provided by VK_KHR_vertex_attribute_divisor
typedef struct VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           vertexAttributeInstanceRateDivisor;
    VkBool32           vertexAttributeInstanceRateZeroDivisor;
} VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR;
```

or the equivalent

``` c
// Provided by VK_EXT_vertex_attribute_divisor
typedef VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR VkPhysicalDeviceVertexAttributeDivisorFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-vertexAttributeInstanceRateDivisor"></span>
  `vertexAttributeInstanceRateDivisor` specifies whether vertex
  attribute fetching may be repeated in the case of instanced rendering.

- <span id="features-vertexAttributeInstanceRateZeroDivisor"></span>
  `vertexAttributeInstanceRateZeroDivisor` specifies whether a zero
  value for
  [VkVertexInputBindingDivisorDescriptionEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDivisorDescriptionEXT.html)::`divisor`
  is supported.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR` **can** also be used
in the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_FEATURES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_vertex_attribute_divisor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_vertex_attribute_divisor.html),
[VK_KHR_vertex_attribute_divisor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_vertex_attribute_divisor.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
