# VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures(3) Manual Page

## Name

VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures - Structure
describing support for zero initialization of workgroup memory by an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures` structure is
defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderZeroInitializeWorkgroupMemory;
} VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures;
```

or the equivalent

``` c
// Provided by VK_KHR_zero_initialize_workgroup_memory
typedef VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-shaderZeroInitializeWorkgroupMemory"></span>
  `shaderZeroInitializeWorkgroupMemory` specifies whether the
  implementation supports initializing a variable in Workgroup storage
  class.

If the `VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures` **can** also be
used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures-sType-sType"
  id="VUID-VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures-sType-sType"></a>
  VUID-VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ZERO_INITIALIZE_WORKGROUP_MEMORY_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_zero_initialize_workgroup_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_zero_initialize_workgroup_memory.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
