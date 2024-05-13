# VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV - Structure
describing feature to allow descriptor pool overallocation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV` structure
is defined as:

``` c
// Provided by VK_NV_descriptor_pool_overallocation
typedef struct VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           descriptorPoolOverallocation;
} VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-descriptorPoolOverallocation"></span>
  `descriptorPoolOverallocation` indicates that the implementation
  allows the application to opt into descriptor pool overallocation by
  creating the descriptor pool with
  `VK_DESCRIPTOR_POOL_CREATE_ALLOW_OVERALLOCATION_SETS_BIT_NV` and/or
  `VK_DESCRIPTOR_POOL_CREATE_ALLOW_OVERALLOCATION_POOLS_BIT_NV` flags.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV`
structure is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV` **can** also be
used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV-sType-sType"
  id="VUID-VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_POOL_OVERALLOCATION_FEATURES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_descriptor_pool_overallocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_descriptor_pool_overallocation.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
